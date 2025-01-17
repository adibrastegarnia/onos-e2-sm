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

var refPerRicKPMnodeItem = "00000000  43 00 6f 6e 66 00 d4 bc  0c 00 15 00 16 00 00 00  |C.onf...........|\n" +
	"00000010  00 03 31 32 33 40 4f 4e  46 09 bc d4 00           |..123@ONF....|"

func createRicKpmNodeItem() (*e2sm_kpm_v2_go.RicKpmnodeItem, error) {

	res := e2sm_kpm_v2_go.RicKpmnodeItem{
		RicKpmnodeType: &e2sm_kpm_v2_go.GlobalKpmnodeId{
			GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_GNb{
				GNb: &e2sm_kpm_v2_go.GlobalKpmnodeGnbId{
					GlobalGNbId: &e2sm_kpm_v2_go.GlobalgNbId{
						GnbId: &e2sm_kpm_v2_go.GnbIdChoice{
							GnbIdChoice: &e2sm_kpm_v2_go.GnbIdChoice_GnbId{
								GnbId: &asn1.BitString{
									Value: []byte{0xd4, 0xbc, 0x0c},
									Len:   22,
								},
							},
						},
						PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
							Value: []byte("onf"),
						},
					},
					GNbCuUpId: &e2sm_kpm_v2_go.GnbCuUpId{
						Value: 21,
					},
					GNbDuId: &e2sm_kpm_v2_go.GnbDuId{
						Value: 22,
					},
				},
			},
		},
		CellMeasurementObjectList: nil,
	}

	item := &e2sm_kpm_v2_go.CellMeasurementObjectItem{
		CellObjectId: &e2sm_kpm_v2_go.CellObjectId{
			Value: "123",
		},
		CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId{
			CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId_EUtraCgi{
				EUtraCgi: &e2sm_kpm_v2_go.Eutracgi{
					EUtracellIdentity: &e2sm_kpm_v2_go.EutracellIdentity{
						Value: &asn1.BitString{
							Value: []byte{0x09, 0xbc, 0xd4, 0x00},
							Len:   28,
						},
					},
					PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: []byte("ONF"),
					},
				},
			},
		},
	}

	res.CellMeasurementObjectList = append(res.CellMeasurementObjectList, item)

	//if err := res.Validate(); err != nil {
	//	return nil, err
	//}
	return &res, nil
}

func Test_perEncodingRicKpmNodeItem(t *testing.T) {

	rkni, err := createRicKpmNodeItem()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(rkni, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("RIC-KPMnodeItem PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicKpmnodeItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("RIC-KPMnodeItem PER - decoded\n%v", &result)
	assert.Equal(t, rkni.GetRicKpmnodeType().GetGNb().GetGNbDuId().GetValue(), result.GetRicKpmnodeType().GetGNb().GetGNbDuId().GetValue())
	assert.Equal(t, rkni.GetRicKpmnodeType().GetGNb().GetGNbCuUpId().GetValue(), result.GetRicKpmnodeType().GetGNb().GetGNbCuUpId().GetValue())
	assert.DeepEqual(t, rkni.GetRicKpmnodeType().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue(), result.GetRicKpmnodeType().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
	assert.DeepEqual(t, rkni.GetRicKpmnodeType().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), result.GetRicKpmnodeType().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())
	assert.Equal(t, rkni.GetRicKpmnodeType().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), result.GetRicKpmnodeType().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	assert.Equal(t, len(rkni.GetCellMeasurementObjectList()), 1)
	assert.Equal(t, rkni.GetCellMeasurementObjectList()[0].GetCellObjectId().GetValue(), result.GetCellMeasurementObjectList()[0].GetCellObjectId().GetValue())
	assert.DeepEqual(t, rkni.GetCellMeasurementObjectList()[0].GetCellGlobalId().GetEUtraCgi().GetPLmnIdentity().GetValue(), result.GetCellMeasurementObjectList()[0].GetCellGlobalId().GetEUtraCgi().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, rkni.GetCellMeasurementObjectList()[0].GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetCellMeasurementObjectList()[0].GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, rkni.GetCellMeasurementObjectList()[0].GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetCellMeasurementObjectList()[0].GetCellGlobalId().GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func Test_perRicKpmNodeItemCompareBytes(t *testing.T) {

	rkni, err := createRicKpmNodeItem()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(rkni, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("RIC-KPMnodeItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicKPMnodeItem)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
