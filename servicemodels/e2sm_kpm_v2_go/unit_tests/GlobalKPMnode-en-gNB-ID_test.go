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

var refPerGlobalKPMnodeEnGnbID = "00000000  60 21 22 23 50 d4 bc 09  00 00 2a 00 20           |`!\"#P.....*. |"

func createGlobalKpmnodeEnGnbID() (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	var bsValue = []byte{0xd4, 0xbc, 0x09, 0x00}
	var bsLen uint32 = 32
	plmnID := []byte{0x21, 0x22, 0x23}
	var gnbDuID int64 = 32
	var gnbCuUpID int64 = 42

	enbID, err := pdubuilder.CreateGlobalKpmnodeIDenGNbID(bsValue, bsLen, plmnID)
	if err != nil {
		return nil, err
	}
	enbID.GetEnGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	enbID.GetEnGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	return enbID, nil
}

func Test_perEncodeGlobalKpmnodeEnGnbID(t *testing.T) {

	enbID, err := createGlobalKpmnodeEnGnbID()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID.GetEnGNb(), "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeEnGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeEnGnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GlobalKPMnodeEnGnbID PER - decoded\n%v", &result)
	assert.DeepEqual(t, enbID.GetEnGNb().GetGlobalGNbId().GetPLmnIdentity().GetValue(), result.GetGlobalGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, enbID.GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetValue(), result.GetGlobalGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, enbID.GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetLen(), result.GetGlobalGNbId().GetGNbId().GetGNbId().GetLen())
	assert.Equal(t, enbID.GetEnGNb().GetGNbCuUpId().GetValue(), result.GetGNbCuUpId().GetValue())
	assert.Equal(t, enbID.GetEnGNb().GetGNbDuId().GetValue(), result.GetGNbDuId().GetValue())
}

func Test_perGlobalKpmnodeEnGnbIDCompareBytes(t *testing.T) {

	enbID, err := createGlobalKpmnodeEnGnbID()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID.GetEnGNb(), "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeEnGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMnodeEnGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
