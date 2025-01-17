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

var refPerGlobalNgEnbID = "00000000  00 21 22 23 00 d4 bc 30  d4 bc c0 d4 bc 08        |.!\"#...0......|"

func createGlobalNgEnbID() (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x30},
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0xc0},
		Len:   18,
	}
	longMacroEnbID := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   21,
	}
	var gnbDuID int64 = 42

	enbIDchoiceMacro, err := pdubuilder.CreateEnbIDchoiceMacro(&bs)
	if err != nil {
		return nil, err
	}
	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(enbIDchoiceMacro, plmnID, &shortMacroEnbID, &longMacroEnbID)
	if err != nil {
		return nil, err
	}
	ngeNbID.GetNgENb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	return ngeNbID, nil
}

func Test_perEncodingGlobalNgEnbID(t *testing.T) {

	ngeNbID, err := createGlobalNgEnbID()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ngeNbID.GetNgENb().GetGlobalNgENbId(), "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalNgEnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalngeNbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GlobalNgEnbID PER - decoded\n%v", &result)
	assert.DeepEqual(t, ngeNbID.GetNgENb().GetGlobalNgENbId().GetPlmnId().GetValue(), result.GetPlmnId().GetValue())
	assert.DeepEqual(t, ngeNbID.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue(), result.GetEnbId().GetEnbIdMacro().GetValue())
	assert.Equal(t, ngeNbID.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen(), result.GetEnbId().GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, ngeNbID.GetNgENb().GetGlobalNgENbId().GetShortMacroENbId().GetValue(), result.GetShortMacroENbId().GetValue())
	assert.Equal(t, ngeNbID.GetNgENb().GetGlobalNgENbId().GetShortMacroENbId().GetLen(), result.GetShortMacroENbId().GetLen())
	assert.DeepEqual(t, ngeNbID.GetNgENb().GetGlobalNgENbId().GetLongMacroENbId().GetValue(), result.GetLongMacroENbId().GetValue())
	assert.Equal(t, ngeNbID.GetNgENb().GetGlobalNgENbId().GetLongMacroENbId().GetLen(), result.GetLongMacroENbId().GetLen())
}

func Test_perGlobalNgEnbIDCompareBytes(t *testing.T) {

	ngeNbID, err := createGlobalNgEnbID()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ngeNbID.GetNgENb().GetGlobalNgENbId(), "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalNgEnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalNgEnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
