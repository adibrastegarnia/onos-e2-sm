// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package generic

import (
	"bytes"
	"fmt"
	pgs "github.com/lyft/protoc-gen-star"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

const moduleName = "choice"

var templateDir string = os.Getenv("GOPATH")
var templates = template.Must(template.ParseGlob(filepath.Join(templateDir, "src/github.com/onosproject/onos-e2-sm/protoc-gen-choice/templates/choice.tpl")))

// Defines data structure to pass to enum template
type choiceStruct struct {
	ProtoFileName string
	Choices       []choiceMsg
}

type choiceMsg struct {
	MsgName string
	Items   []choiceItem
}

type choiceItem struct {
	Leafs      []leaf
	ChoiceName string
}

type leaf struct {
	Index    int
	LeafName string
}

// ReportModule creates a report of all the target messages generated by the
// protoc run, writing the file into the /tmp directory.
type reportModule struct {
	*pgs.ModuleBase
}

// NewModule configures the module with an instance of ModuleBase
func NewModule() pgs.Module {
	return &reportModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

// Name is the identifier used to identify the module. This value is
// automatically attached to the BuildContext associated with the ModuleBase.
func (m *reportModule) Name() string {
	return moduleName
}

// Execute is passed the target files as well as its dependencies in the pkgs
// map. The implementation should return a slice of Artifacts that represent
// the files to be generated. In this case, "/tmp/report.txt" will be created
// outside of the normal protoc flow.
func (m *reportModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	for _, f := range targets { // Input .proto files
		m.Push(f.Name().String()).Debug("reporting")

		choices := choiceStruct{
			ProtoFileName: adjustProtoFileName(extractProtoFileName(f.Name().Split()[0])),
			Choices:       make([]choiceMsg, 0),
		}
		fmt.Fprintf(buf, "ProtoFileName is %v\n", choices.ProtoFileName)

		for _, msg := range f.AllMessages() {
			fmt.Fprintf(buf, "OneOf list is %v\n", msg.OneOfs())
			if msg.OneOfs() != nil {
				chMsg := choiceMsg{
					MsgName: adjustOneOfStructName(msg.Name().String()),
					Items:   make([]choiceItem, 0),
				}
				for i, plg := range msg.OneOfs() {
					fmt.Fprintf(buf, "%v OneOf name is %v\n", i+1, plg.Name())
					chItem := choiceItem{
						ChoiceName: plg.Name().String(),
						Leafs:      make([]leaf, 0),
					}
					for j, field := range plg.Fields() {
						//plg.Fields()
						fmt.Fprintf(buf, "%v, OneOf field is %v\n", j+1, field.Name())
						lf := leaf{
							Index:    j + 1,
							LeafName: msg.Name().String() + "_" + adjustOneOfLeafName(field.Name().String()),
						}
						chItem.Leafs = append(chItem.Leafs, lf)
					}
					chMsg.Items = append(chMsg.Items, chItem)
				}

				//for j, oneof := range msg.OneOfFields() {
				//	oneofItem := choiceItem{
				//		Index:    j + 1,
				//		LeafName: msg.Name().String() + "_" + adjustOneOfLeafName(oneof.Name().String()),
				//	}
				//	chItem.Leafs = append(chItem.Leafs, oneofItem)
				//	fmt.Fprintf(buf, "Obtained OneOf leaf is \n%v\n", oneofItem)
				//}
				fmt.Fprintf(buf, "Obtained OneOf item is \n%v\n", chMsg)

				choices.Choices = append(choices.Choices, chMsg)
			}
		}

		//Generating new .go file
		m.OverwriteGeneratorTemplateFile("choiceOptions.go", templates.Lookup("choice.tpl"), choices)

		m.Pop()
		fmt.Fprintf(buf, "-----------------------------------------------------------------------------------------------\n")
	}
	m.OverwriteCustomFile(
		"/tmp/report.txt",
		buf.String(),
		0644,
	)

	return m.Artifacts()
}

/////////////////////////////////
/// Here is necessary tooling ///
/////////////////////////////////

func upperCaseFirstLetter(str string) string {

	for i, ch := range str {
		return string(unicode.ToUpper(ch)) + str[i+1:]
	}
	return ""
}

func adjustOneOfLeafName(leafName string) string {

	var res string
	for _, element := range strings.Split(leafName, "_") {
		res = res + upperCaseFirstLetter(element)
	}

	return res
}

func adjustOneOfStructName(msgName string) string {

	var res string
	for i, r := range msgName {
		if unicode.IsUpper(r) {
			if i > 1 {
				res = res + "_" + strings.ToLower(string(r))
			} else {
				res = strings.ToLower(string(r))
			}
		} else {
			res = res + string(r)
		}
	}

	return res
}

func extractProtoFileName(proto string) string {

	if strings.LastIndex(proto, "/") != 1 {
		return proto[strings.LastIndex(proto, "/")+1:]
	}
	return proto
}

func adjustProtoFileName(filename string) string {

	res := dashToUnderscore(filename)
	// space for future adjustments
	return res
}

func dashToUnderscore(str string) string {

	return strings.ReplaceAll(str, "-", "_")
}