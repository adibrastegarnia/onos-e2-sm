// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMCI1 = "00000000  1f ff f0 21 22 23 40 40  01 02 03 00 17 68 18 00  |...!\"#@@.....h..|\n" +
	"00000010  1e 00 01 70 00 00 18 00  00 00 00 00 7a 00 01 c7  |...p........z...|\n" +
	"00000020  00 03 14 20                                       |... |"
var refPerMCI2 = "00000000  42 10 01 15                                       |B...|"

func createMatchingCondItem1() (*e2sm_kpm_v2_go.MatchingCondItem, error) {

	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 23
	var qfi int32 = 52
	var qci int32 = 24
	var qciMin int32 = 1
	var qciMax int32 = 30
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 25
	var layerMuMimo int32 = 1
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	if err != nil {
		return nil, err
	}

	mci, err := pdubuilder.CreateMatchingCondItemMeasLabel(labelInfoItem.GetMeasLabel())
	if err != nil {
		return nil, err
	}
	//if err := mci.Validate(); err != nil {
	//	return nil, err
	//}
	return mci, nil
}

func createMatchingCondItem2() (*e2sm_kpm_v2_go.MatchingCondItem, error) {

	testCondInfo := &e2sm_kpm_v2_go.TestCondInfo{
		TestType: &e2sm_kpm_v2_go.TestCondType{
			TestCondType: &e2sm_kpm_v2_go.TestCondType_AMbr{
				AMbr: e2sm_kpm_v2_go.AMBR_AMBR_TRUE,
			},
		},
		TestExpr: e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
		TestValue: &e2sm_kpm_v2_go.TestCondValue{
			TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueInt{
				ValueInt: 21,
			},
		},
	}

	mci, err := pdubuilder.CreateMatchingCondItemTestCondInfo(testCondInfo)
	if err != nil {
		return nil, err
	}
	//if err := mci.Validate(); err != nil {
	//	return nil, err
	//}
	return mci, nil
}

func Test_perEncodingMatchingCondItem1(t *testing.T) {

	mci, err := createMatchingCondItem1()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MatchingCondItem (MeasLabel) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MatchingCondItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MatchingCondItem (MeasLabel) PER - decoded\n%v", &result)
	assert.DeepEqual(t, mci.GetMeasLabel().GetPlmnId().GetValue(), result.GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, mci.GetMeasLabel().GetSliceId().GetSD(), result.GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, mci.GetMeasLabel().GetSliceId().GetSSt(), result.GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, mci.GetMeasLabel().GetFiveQi().GetValue(), result.GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, mci.GetMeasLabel().GetQFi().GetValue(), result.GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, mci.GetMeasLabel().GetQCi().GetValue(), result.GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, mci.GetMeasLabel().GetQCimax().GetValue(), result.GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, mci.GetMeasLabel().GetQCimin().GetValue(), result.GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, mci.GetMeasLabel().GetARpmax().GetValue(), result.GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, mci.GetMeasLabel().GetARpmin().GetValue(), result.GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, mci.GetMeasLabel().GetBitrateRange(), result.GetMeasLabel().GetBitrateRange())
	assert.Equal(t, mci.GetMeasLabel().GetLayerMuMimo(), result.GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, mci.GetMeasLabel().GetSUm().Number(), result.GetMeasLabel().GetSUm().Number())
	assert.Equal(t, mci.GetMeasLabel().GetDistBinX(), result.GetMeasLabel().GetDistBinX())
	assert.Equal(t, mci.GetMeasLabel().GetDistBinY(), result.GetMeasLabel().GetDistBinY())
	assert.Equal(t, mci.GetMeasLabel().GetDistBinZ(), result.GetMeasLabel().GetDistBinZ())
	assert.Equal(t, mci.GetMeasLabel().GetPreLabelOverride().Number(), result.GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, mci.GetMeasLabel().GetStartEndInd().Number(), result.GetMeasLabel().GetStartEndInd().Number())
}

func Test_perMatchingCondItem1CompareBytes(t *testing.T) {

	mci, err := createMatchingCondItem1()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MatchingCondItem (MeasLabel) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMCI1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingMatchingCondItem2(t *testing.T) {

	mci, err := createMatchingCondItem2()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MatchingCondItem (TestCondInfo) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MatchingCondItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MatchingCondItem (TestCondInfo) PER - decoded\n%v", &result)
	assert.Equal(t, mci.GetTestCondInfo().GetTestValue().GetValueInt(), result.GetTestCondInfo().GetTestValue().GetValueInt())
	assert.Equal(t, mci.GetTestCondInfo().GetTestType().GetAMbr().Number(), result.GetTestCondInfo().GetTestType().GetAMbr().Number())
	assert.Equal(t, mci.GetTestCondInfo().GetTestExpr().Number(), result.GetTestCondInfo().GetTestExpr().Number())
}

func Test_perMatchingCondItem2CompareBytes(t *testing.T) {

	mci, err := createMatchingCondItem2()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MatchingCondItem (TestCondInfo) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMCI2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
