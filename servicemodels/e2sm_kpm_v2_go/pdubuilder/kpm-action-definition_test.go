// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmActionDefinitionFormat1(t *testing.T) {
	var ricStyleType int32 = 12
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
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)

	labelInfoList := e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem := CreateMeasurementInfoItem(measName).SetLabelInfoList(&labelInfoList)
	assert.NilError(t, err)

	measInfoList := e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinition, err := CreateActionDefinitionFormat1(cellObjID, &measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat1(ricStyleType, actionDefinition)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}

func TestE2SmKpmActionDefinitionFormat2(t *testing.T) {
	var ricStyleType int32 = 12
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
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)

	labelInfoList := e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem := CreateMeasurementInfoItem(measName).SetLabelInfoList(&labelInfoList)

	measInfoList := &e2sm_kpm_v2_go.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, err := CreateActionDefinitionFormat1(cellObjID, measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	ueID := "SomeUE"
	actionDefinitionFormat2, err := CreateActionDefinitionFormat2([]byte(ueID), actionDefinitionFormat1)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat2(ricStyleType, actionDefinitionFormat2)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}

func TestE2SmKpmActionDefinitionFormat3(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	var valEnum int64 = 201
	tce := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
	tci, err := CreateTestCondInfo(CreateTestCondTypeRSRP(), tce, CreateTestCondValueEnum(valEnum))
	assert.NilError(t, err)

	mci, err := CreateMatchingCondItemTestCondInfo(tci)
	assert.NilError(t, err)

	mcl := &e2sm_kpm_v2_go.MatchingCondList{
		Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci)

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measCondItem, err := CreateMeasurementCondItem(measName, mcl)
	assert.NilError(t, err)

	measCondList := e2sm_kpm_v2_go.MeasurementCondList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondItem, 0),
	}
	measCondList.Value = append(measCondList.Value, measCondItem)

	actionDefinitionFormat3, err := CreateActionDefinitionFormat3(cellObjID, &measCondList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinitionFormat3(ricStyleType, actionDefinitionFormat3)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
