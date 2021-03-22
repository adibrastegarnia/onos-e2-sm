// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createActionDefinitionFormat1() (*e2sm_kpm_v2.E2SmKpmActionDefinitionFormat1, error) {

	var cellObjID string = "onf"
	var granularity int32 = 21
	var subscriptionID int64 = 12345
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
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	startEndIndication := e2sm_kpm_v2.StartEndInd_START_END_IND_START
	var measurementName string = "trial"

	labelInfoItem, _ := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, fiveQI, qfi,
		qci, qciMax, qciMin, arpMax, arpMin, bitrateRange, layerMuMimo,
		distX, distY, distZ, startEndIndication)

	labelInfoList := e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	measInfoItem, _ := pdubuilder.CreateMeasurementInfoItem(measName, &labelInfoList)

	measInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinitionFormat1, _ := pdubuilder.CreateActionDefinitionFormat1(cellObjID, &measInfoList, granularity, subscriptionID)
	if err := actionDefinitionFormat1.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat1 %s", err.Error())
	}
	return actionDefinitionFormat1, nil
}

func Test_xerEncodeE2SmKpmActionDefinitionFormat1(t *testing.T) {

	actionDefFormat1, err := createActionDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmKpmActionDefinitionFormat1(actionDefFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 1486, len(xer))
	t.Logf("E2SmKpmActionDefinitionFormat1 XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmActionDefinitionFormat1(t *testing.T) {

	actionDefFormat1, err := createActionDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmKpmActionDefinitionFormat1(actionDefFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 1486, len(xer))
	t.Logf("E2SmKpmActionDefinitionFormat1 XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmActionDefinitionFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmActionDefinitionFormat1 XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmActionDefinitionFormat1(t *testing.T) {

	actionDefFormat1, err := createActionDefinitionFormat1()
	assert.NilError(t, err)

	per, err := perEncodeE2SmKpmActionDefinitionFormat1(actionDefFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 57, len(per))
	t.Logf("E2SmKpmActionDefinitionFormat1 PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmActionDefinitionFormat1(t *testing.T) {

	actionDefFormat1, err := createActionDefinitionFormat1()
	assert.NilError(t, err)

	per, err := perEncodeE2SmKpmActionDefinitionFormat1(actionDefFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 57, len(per))
	t.Logf("E2SmKpmActionDefinitionFormat1 PER\n%s", string(per))

	result, err := perDecodeE2SmKpmActionDefinitionFormat1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmActionDefinitionFormat1 PER - decoded\n%s", result)
}