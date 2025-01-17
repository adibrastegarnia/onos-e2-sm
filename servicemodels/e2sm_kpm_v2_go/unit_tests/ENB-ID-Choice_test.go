// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerEnbIDchoiceMacro = "00000000  00 d4 bc 00                                       |....|"
var refPerEnbIDchoiceShortMacro = "00000000  20 d4 bc 00                                       | ...|"
var refPerEnbIDchoiceLongMacro = "00000000  40 d4 bc 08                                       |@...|"

func createEnbIDChoiceMacro() *e2sm_kpm_v2_go.EnbIdChoice {

	return &e2sm_kpm_v2_go.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x00},
				Len:   20,
			},
		},
	}
}

func createEnbIDChoiceShortMacro() *e2sm_kpm_v2_go.EnbIdChoice {

	return &e2sm_kpm_v2_go.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x00},
				Len:   18,
			},
		},
	}
}

func createEnbIDChoiceLongMacro() *e2sm_kpm_v2_go.EnbIdChoice {

	return &e2sm_kpm_v2_go.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x08},
				Len:   21,
			},
		},
	}
}

func Test_perEncodingEnbIDchoiceMacro(t *testing.T) {

	enbID := createEnbIDChoiceMacro()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("EnbIDchoice (Macro) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.EnbIdChoice{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("EnbIDchoice (Macro) PER - decoded\n%v", &result)
	assert.DeepEqual(t, enbID.GetEnbIdMacro().GetValue(), result.GetEnbIdMacro().GetValue())
	assert.Equal(t, enbID.GetEnbIdMacro().GetLen(), result.GetEnbIdMacro().GetLen())
}

func Test_perEnbIDchoiceMacroCompareBytes(t *testing.T) {

	enbID := createEnbIDChoiceMacro()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("EnbIDchoice (Macro) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnbIDchoiceMacro)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingEnbIDchoiceShortMacro(t *testing.T) {

	enbID := createEnbIDChoiceShortMacro()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("EnbIDchoice (Short Macro) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.EnbIdChoice{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("EnbIDchoice (Short Macro) PER - decoded\n%v", &result)
	assert.DeepEqual(t, enbID.GetEnbIdShortmacro().GetValue(), result.GetEnbIdShortmacro().GetValue())
	assert.Equal(t, enbID.GetEnbIdShortmacro().GetLen(), result.GetEnbIdShortmacro().GetLen())
}

func Test_perEnbIDchoiceShortMacroCompareBytes(t *testing.T) {

	enbID := createEnbIDChoiceShortMacro()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("EnbIDchoice (Short Macro) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnbIDchoiceShortMacro)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingEnbIDchoiceLongMacro(t *testing.T) {

	enbID := createEnbIDChoiceLongMacro()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("EnbIDchoice (Long Macro) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.EnbIdChoice{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("EnbIDchoice (Long Macro) PER - decoded\n%v", &result)
	assert.DeepEqual(t, enbID.GetEnbIdShortmacro().GetValue(), result.GetEnbIdShortmacro().GetValue())
	assert.Equal(t, enbID.GetEnbIdShortmacro().GetLen(), result.GetEnbIdShortmacro().GetLen())
}

func Test_perEnbIDchoiceLongMacroCompareBytes(t *testing.T) {

	enbID := createEnbIDChoiceLongMacro()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("EnbIDchoice (Long Macro) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnbIDchoiceLongMacro)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
