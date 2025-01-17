// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerGlobalKPMnodeGnbID = "00000000  60 21 22 23 00 d4 bc 08  00 1f 00 2a              |`!\"#.......*|"

func createGlobalKpmnodeGnbID() (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUpID int64 = 31
	var gnbDuID int64 = 42

	gNbID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	if err != nil {
		return nil, err
	}
	gNbID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gNbCuUpID,
	}
	gNbID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	return gNbID, nil
}

func Test_perEncodeGlobalKpmnodeGnbID(t *testing.T) {

	gNbID, err := createGlobalKpmnodeGnbID()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(gNbID.GetGNb(), "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeGnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GlobalKPMnodeGnbID PER - decoded\n%v", &result)
	assert.DeepEqual(t, gNbID.GetGNb().GetGlobalGNbId().GetPlmnId().GetValue(), result.GetGlobalGNbId().GetPlmnId().GetValue())
	assert.DeepEqual(t, gNbID.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), result.GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())
	assert.Equal(t, gNbID.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), result.GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	assert.Equal(t, gNbID.GetGNb().GetGNbCuUpId().GetValue(), result.GetGNbCuUpId().GetValue())
	assert.Equal(t, gNbID.GetGNb().GetGNbDuId().GetValue(), result.GetGNbDuId().GetValue())
}

func Test_perGlobalKpmnodeGnbIDCompareBytes(t *testing.T) {

	gNbID, err := createGlobalKpmnodeGnbID()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(gNbID.GetGNb(), "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMnodeGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
