// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/pdubuilder"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var mhoTestSm MhoServiceModel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	cellID := asn1.BitString{
		Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
		Len:   28, //uint32
	}
	cgi, err := pdubuilder.CreateCellGlobalIDEutraCGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoIndicationHeader(cgi)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmMhoPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmMhoPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmMhoIndicationHeader to bytes")
	//assert.Equal(t, 24, len(protoBytes))

	asn1Bytes, err := mhoTestSm.IndicationHeaderProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for MHO-IndicationHeader are \n%v", hex.Dump(asn1Bytes))
	//assert.Equal(t, 8, len(asn1Bytes))
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	indicationHeaderAsn1Bytes := []byte{0x10, 0x12, 0xf4, 0x10, 0xab, 0xd4, 0xbc, 0x00}
	protoBytes, err := mhoTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	//assert.Equal(t, 24, len(protoBytes))
	testIH := &e2sm_mho_go.E2SmMhoIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded MHO-IndicationHeader is \n%v", testIH)
	assert.DeepEqual(t, []byte{0x12, 0xf4, 0x10}, testIH.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().GetValue())
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	ueID := &e2sm_mho_go.UeIdentity{
		Value: []byte("1234"),
	}
	cgi, err := pdubuilder.CreateCellGlobalIDNrCGI([]byte{0xAA, 0xFD, 0xD4}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x40, 0x00},
		Len:   36,
	})
	assert.NilError(t, err)
	rsrp := &e2sm_mho_go.Rsrp{
		Value: 1234,
	}
	measItem, err := pdubuilder.CreateMeasurementRecordItem(cgi, rsrp)
	assert.NilError(t, err)

	measReport := make([]*e2sm_mho_go.E2SmMhoMeasurementReportItem, 0)
	measReport = append(measReport, measItem)

	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoIndicationMsgFormat1(ueID, measReport)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmMhoPdu != nil)

	err = newE2SmMhoPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmMhoPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmMhoIndicationMessage to bytes")

	//assert.Equal(t, 39, len(protoBytes))

	asn1Bytes, err := mhoTestSm.IndicationMessageProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for MHO-IndicationMessage (Format1) are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	indicationMessageAsn1 := []byte{
		0x00, 0x04, 0x31, 0x32, 0x33, 0x34, 0x00, 0x80, 0xaa, 0xfd, 0xd4, 0x00, 0x00, 0x00, 0x40, 0x04,
		0x01, 0x04, 0xd2, 0x00, 0x15,
	}
	protoBytes, err := mhoTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	//assert.Equal(t, 39, len(protoBytes))
	testIM := &e2sm_mho_go.E2SmMhoIndicationMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded MHO-IndicationMessage is \n%v", testIM)

	assert.DeepEqual(t, []byte("1234"), testIM.GetIndicationMessageFormat1().GetUeId().Value)
	//assert.Equal(t, 1234, int(testIM.GetIndicationMessageFormat1().GetRsrp().Value))
}

func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 3
	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42
	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56

	retsl := make([]*e2sm_mho_go.RicEventTriggerStyleList, 0)
	retsi, err := pdubuilder.CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	assert.NilError(t, err)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2sm_mho_go.RicReportStyleList, 0)
	rrsi, err := pdubuilder.CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	assert.NilError(t, err)
	rrsl = append(rrsl, rrsi)
	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, newE2SmMhoPdu != nil)
	newE2SmMhoPdu.SetRicReportStyleList(rrsl).SetRicEventTriggerStyleList(retsl).GetRanFunctionName().SetRanFunctionInstance(ranFunctionInstance)

	err = newE2SmMhoPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmMhoPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmMhoRanfunctionDescription to bytes")
	//assert.Equal(t, 79, len(protoBytes))

	asn1Bytes, err := mhoTestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	//assert.Equal(t, 59, len(asn1Bytes))
	t.Logf("ASN1 bytes for MHO-RanFunctionDescription are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
	// This message is taken as an output from the function above
	ranFuncDescriptionAsn1 := []byte{
		0x20, 0x20, 0x4f, 0x4e, 0x46, 0x00, 0x00, 0x02, 0x4f, 0x69, 0x64, 0x06, 0x80, 0x4f, 0x70, 0x65,
		0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x01, 0x03, 0x60, 0x06, 0x81,
		0xc0, 0x4f, 0x4e, 0x46, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x00, 0x2a, 0x00, 0x30, 0x10, 0x4f, 0x4e,
		0x46, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x00, 0x15, 0x00, 0x38,
	}

	protoBytes, err := mhoTestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	//assert.Equal(t, 79, len(protoBytes))
	testRFD := &e2sm_mho_go.E2SmMhoRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, testRFD)
	t.Logf("Decoded MHO-RanFunctionDescription is \n%v", testRFD)
	assert.NilError(t, err)
	assert.Equal(t, "Oid", testRFD.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, int32(3), testRFD.GetRanFunctionName().GetRanFunctionInstance())
}

func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
	var rtPeriod int32 = 12
	e2SmMhoEventTriggerDefinition, err := pdubuilder.CreateE2SmMhoEventTriggerDefinitionPeriodic(rtPeriod)
	assert.NilError(t, err, "error creating E2SmPDU")
	assert.Assert(t, e2SmMhoEventTriggerDefinition != nil, "Created E2SmPDU is nil")

	err = e2SmMhoEventTriggerDefinition.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(e2SmMhoEventTriggerDefinition)
	assert.NilError(t, err, "unexpected error marshalling E2SmMhoEventTriggerDefinition to bytes")
	//assert.Equal(t, 4, len(protoBytes))

	asn1Bytes, err := mhoTestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	//assert.Equal(t, 3, len(asn1Bytes))
	t.Logf("ASN1 bytes for MHO-EventTriggerDefinition are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	eventTriggerDefinitionAsn1 := []byte{0x14, 0x01, 0x0c}
	protoBytes, err := mhoTestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	//assert.Equal(t, 6, len(protoBytes))
	testIM := &e2sm_mho_go.E2SmMhoEventTriggerDefinition{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded MHO-EventTriggerDefinition is \n%v", testIM)
	assert.Equal(t, int32(12), testIM.GetEventDefinitionFormats().GetEventDefinitionFormat1().GetReportingPeriodMs())
}

func TestServicemodel_ControlHeaderProtoToASN1(t *testing.T) {

	var controlMessagePriority int32 = 1

	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoControlHeader(controlMessagePriority)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmMhoPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmMhoPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmMhoControlHeader to bytes")
	assert.Equal(t, 6, len(protoBytes))

	asn1Bytes, err := mhoTestSm.ControlHeaderProtoToASN1(protoBytes)

	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 2, len(asn1Bytes))
	t.Logf("ASN1 bytes for MHO-ControlHeader are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlHeaderASN1toProto(t *testing.T) {
	ControlHeaderAsn1Bytes := []byte{0x20, 0x01, 0x01}

	protoBytes, err := mhoTestSm.ControlHeaderASN1toProto(ControlHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 6, len(protoBytes))
	testCH := &e2sm_mho_go.E2SmMhoControlHeader{}
	err = proto.Unmarshal(protoBytes, testCH)
	assert.NilError(t, err)
	t.Logf("Decoded MHO-ControlHeader is \n%v", testCH)
	assert.Equal(t, 1, int(testCH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))
	assert.DeepEqual(t, 1, int(testCH.GetControlHeaderFormat1().GetRicControlMessagePriority().GetValue()))

}

func TestServicemodel_ControlMessageProtoToASN1(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	servingCgi := &e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho_go.Eutracgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_mho_go.EutracellIdentity{
					Value: &asn1.BitString{
						Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
						Len:   28, //uint32
					},
				},
			},
		},
	}
	targetCgi := &e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho_go.Eutracgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_mho_go.EutracellIdentity{
					Value: &asn1.BitString{
						Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
						Len:   28, //uint32
					},
				},
			},
		},
	}
	ueID := &e2sm_mho_go.UeIdentity{
		Value: []byte("1234"),
	}
	newE2SmMhoPdu, err := pdubuilder.CreateE2SmMhoControlMessage(servingCgi, ueID, targetCgi)
	assert.NilError(t, err, "error creating E2SmPDU")

	//err = newE2SmMhoPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")
	protoBytes, err := proto.Marshal(newE2SmMhoPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmMhoControlMessage to bytes")

	//assert.Equal(t, 54, len(protoBytes))
	asn1Bytes, err := mhoTestSm.ControlMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("ASN1 bytes for MHO-ControlMessage are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlMessageASN1toProto(t *testing.T) {
	ControlMessageAsn1 := []byte{
		0x10, 0x12, 0xf4, 0x10, 0xab, 0xd4, 0xbc, 0x00, 0x04, 0x31, 0x32, 0x33, 0x34, 0x40, 0x12, 0xf4,
		0x10, 0xab, 0xd4, 0xbc, 0x00,
	}

	protoBytes, err := mhoTestSm.ControlMessageASN1toProto(ControlMessageAsn1)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	//assert.Equal(t, 54, len(protoBytes))
	testIM := &e2sm_mho_go.E2SmMhoControlMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded MHO-ControlMessage is \n%v", testIM)
	assert.DeepEqual(t, []byte("1234"), testIM.GetControlMessageFormat1().UedId.GetValue())
}
