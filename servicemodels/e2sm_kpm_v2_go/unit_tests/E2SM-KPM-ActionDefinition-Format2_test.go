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

var refPerAD2 = "00000000  00 06 53 6f 6d 65 55 45  00 00 03 6f 6e 66 00 00  |..SomeUE...onf..|\n" +
	"00000010  40 20 74 72 69 61 6c 01  3f ff e0 21 22 23 40 40  |@ trial.?..!\"#@@|\n" +
	"00000020  01 02 03 00 0a 7c 0f 00  0f 00 01 70 00 00 fa 00  |.....|.....p....|\n" +
	"00000030  00 04 00 00 7a 00 01 c7  00 03 14 00 14 40 30 38  |....z........@08|"

func createActionDefinitionFormat2() (*e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat2, error) {

	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qfi int32 = 62
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	if err != nil {
		return nil, err
	}

	labelInfoList := e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	if err != nil {
		return nil, err
	}
	measInfoItem := pdubuilder.CreateMeasurementInfoItem(measName).SetLabelInfoList(&labelInfoList)

	measInfoList := &e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, err := pdubuilder.CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)
	if err != nil {
		return nil, err
	}
	ueID := "SomeUE"
	actionDefinitionFormat2, err := pdubuilder.CreateActionDefinitionFormat2([]byte(ueID), actionDefinitionFormat1)
	if err != nil {
		return nil, err
	}
	//if err := actionDefinitionFormat2.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat3 %s", err.Error())
	//}
	return actionDefinitionFormat2, nil
}

func Test_perEncodingE2SmKpmActionDefinitionFormat2(t *testing.T) {

	actionDefFormat2, err := createActionDefinitionFormat2()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(actionDefFormat2, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition-Format2 PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat2{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-ActionDefinition-Format2 PER - decoded\n%v", &result)
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetCellObjId().GetValue(), result.GetSubscriptInfo().GetCellObjId().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetGranulPeriod().GetValue(), result.GetSubscriptInfo().GetGranulPeriod().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetSubscriptId().GetValue(), result.GetSubscriptInfo().GetSubscriptId().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.DeepEqual(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, actionDefFormat2.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetSubscriptInfo().GetMeasInfoList().GetValue()[0].GetLabelInfoList().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.DeepEqual(t, actionDefFormat2.GetUeId().GetValue(), result.GetUeId().GetValue())
}

func Test_perE2SmKpmActionDefinitionFormat2CompareBytes(t *testing.T) {

	actionDefFormat2, err := createActionDefinitionFormat2()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(actionDefFormat2, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-ActionDefinition-Format2 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerAD2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
