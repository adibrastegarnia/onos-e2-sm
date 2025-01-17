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

func createGlobalKpmnodeID1() *e2sm_kpm_v2_go.GlobalKpmnodeId {

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_GNb{
			GNb: &e2sm_kpm_v2_go.GlobalKpmnodeGnbId{
				GlobalGNbId: &e2sm_kpm_v2_go.GlobalgNbId{
					GnbId: &e2sm_kpm_v2_go.GnbIdChoice{
						GnbIdChoice: &e2sm_kpm_v2_go.GnbIdChoice_GnbId{
							GnbId: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x09, 0x00},
								Len:   28,
							},
						},
					},
					PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
				},
				GNbCuUpId: &e2sm_kpm_v2_go.GnbCuUpId{
					Value: 21,
				},
				GNbDuId: &e2sm_kpm_v2_go.GnbDuId{
					Value: 32,
				},
			},
		},
	}
}

func createGlobalKpmnodeID2() *e2sm_kpm_v2_go.GlobalKpmnodeId {

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_EnGNb{
			EnGNb: &e2sm_kpm_v2_go.GlobalKpmnodeEnGnbId{
				GlobalGNbId: &e2sm_kpm_v2_go.GlobalenGnbId{
					GNbId: &e2sm_kpm_v2_go.EngnbId{
						EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
							GNbId: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x09, 0x00},
								Len:   28,
							},
						},
					},
					PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
				},
				GNbCuUpId: &e2sm_kpm_v2_go.GnbCuUpId{
					Value: 21,
				},
				GNbDuId: &e2sm_kpm_v2_go.GnbDuId{
					Value: 32,
				},
			},
		},
	}
}

func createGlobalKpmnodeID3() *e2sm_kpm_v2_go.GlobalKpmnodeId {

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_NgENb{
			NgENb: &e2sm_kpm_v2_go.GlobalKpmnodeNgEnbId{
				GlobalNgENbId: &e2sm_kpm_v2_go.GlobalngeNbId{
					EnbId: &e2sm_kpm_v2_go.EnbIdChoice{
						EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdMacro{
							EnbIdMacro: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x00},
								Len:   20,
							},
						},
					},
					PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
					ShortMacroENbId: &asn1.BitString{
						Value: []byte{0xd4, 0xbc, 0x00},
						Len:   18,
					},
					LongMacroENbId: &asn1.BitString{
						Value: []byte{0xd4, 0xbc, 0x08},
						Len:   21,
					},
				},
				GNbDuId: &e2sm_kpm_v2_go.GnbDuId{
					Value: 32,
				},
			},
		},
	}
}

func createGlobalKpmnodeID4() *e2sm_kpm_v2_go.GlobalKpmnodeId {

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_ENb{
			ENb: &e2sm_kpm_v2_go.GlobalKpmnodeEnbId{
				GlobalENbId: &e2sm_kpm_v2_go.GlobalEnbId{
					ENbId: &e2sm_kpm_v2_go.EnbId{
						EnbId: &e2sm_kpm_v2_go.EnbId_HomeENbId{
							HomeENbId: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x09, 0x00},
								Len:   28,
							},
						},
					},
					PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
				},
			},
		},
	}
}

var refPerGlobalKPMnodeIDgnb = "00000000  0c 21 22 23 30 d4 bc 09  00 15 00 20              |.!\"#0...... |"
var refPerGlobalKPMnodeIDenGnb = "00000000  2c 21 22 23 30 d4 bc 09  00 15 00 20              |,!\"#0...... |"
var refPerGlobalKPMnodeIDngEnb = "00000000  48 21 22 23 00 d4 bc 00  d4 bc 00 d4 bc 08 20     |H!\"#.......... |"
var refPerGlobalKPMnodeIDenb = "00000000  60 21 22 23 40 d4 bc 09  00                       |`!\"#@....|"

func Test_perEncodingGlobalKpmNodeIDgnb(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID1()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (GNb) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GlobalKpmnodeID (GNb) PER - decoded\n%v", &result)
	assert.DeepEqual(t, globalKpmnodeID.GetGNb().GetGlobalGNbId().GetPlmnId().GetValue(), result.GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
	assert.DeepEqual(t, globalKpmnodeID.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), result.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())
	assert.Equal(t, globalKpmnodeID.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), result.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	assert.Equal(t, globalKpmnodeID.GetGNb().GetGNbCuUpId().GetValue(), result.GetGNb().GetGNbCuUpId().GetValue())
	assert.Equal(t, globalKpmnodeID.GetGNb().GetGNbDuId().GetValue(), result.GetGNb().GetGNbDuId().GetValue())
}

func Test_perGlobalKpmNodeIDgnbCompareBytes(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID1()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (GNb) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMnodeIDgnb)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGlobalKpmNodeIDenGnb(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID2()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (enGNb) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GlobalKpmnodeID (enGNb) PER - decoded\n%v", &result)
	assert.DeepEqual(t, globalKpmnodeID.GetEnGNb().GetGlobalGNbId().GetPLmnIdentity().GetValue(), result.GetEnGNb().GetGlobalGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, globalKpmnodeID.GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetValue(), result.GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, globalKpmnodeID.GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetLen(), result.GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetLen())
	assert.Equal(t, globalKpmnodeID.GetEnGNb().GetGNbCuUpId().GetValue(), result.GetEnGNb().GetGNbCuUpId().GetValue())
	assert.Equal(t, globalKpmnodeID.GetEnGNb().GetGNbDuId().GetValue(), result.GetEnGNb().GetGNbDuId().GetValue())
}

func Test_perGlobalKpmNodeIDenGnbCompareBytes(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID2()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (enGNb) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMnodeIDenGnb)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGlobalKpmNodeIDngEnb(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID3()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (ngENb) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GlobalKpmnodeID (ngENb) PER - decoded\n%v", &result)
	assert.DeepEqual(t, globalKpmnodeID.GetNgENb().GetGlobalNgENbId().GetPlmnId().GetValue(), result.GetNgENb().GetGlobalNgENbId().GetPlmnId().GetValue())
	assert.DeepEqual(t, globalKpmnodeID.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue(), result.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue())
	assert.Equal(t, globalKpmnodeID.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen(), result.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, globalKpmnodeID.GetNgENb().GetGlobalNgENbId().GetShortMacroENbId().GetValue(), result.GetNgENb().GetGlobalNgENbId().GetShortMacroENbId().GetValue())
	assert.Equal(t, globalKpmnodeID.GetNgENb().GetGlobalNgENbId().GetShortMacroENbId().GetLen(), result.GetNgENb().GetGlobalNgENbId().GetShortMacroENbId().GetLen())
	assert.DeepEqual(t, globalKpmnodeID.GetNgENb().GetGlobalNgENbId().GetLongMacroENbId().GetValue(), result.GetNgENb().GetGlobalNgENbId().GetLongMacroENbId().GetValue())
	assert.Equal(t, globalKpmnodeID.GetNgENb().GetGlobalNgENbId().GetLongMacroENbId().GetLen(), result.GetNgENb().GetGlobalNgENbId().GetLongMacroENbId().GetLen())
	assert.Equal(t, globalKpmnodeID.GetNgENb().GetGNbDuId().GetValue(), result.GetNgENb().GetGNbDuId().GetValue())
}

func Test_perGlobalKpmNodeIDngEnbCompareBytes(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID3()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (ngENb) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMnodeIDngEnb)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGlobalKpmNodeIDenb(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID4()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (ENb) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GlobalKpmnodeID (ENb) PER - decoded\n%v", &result)
	assert.DeepEqual(t, globalKpmnodeID.GetENb().GetGlobalENbId().GetPLmnIdentity().GetValue(), result.GetENb().GetGlobalENbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, globalKpmnodeID.GetENb().GetGlobalENbId().GetENbId().GetHomeENbId().GetValue(), result.GetENb().GetGlobalENbId().GetENbId().GetHomeENbId().GetValue())
	assert.Equal(t, globalKpmnodeID.GetENb().GetGlobalENbId().GetENbId().GetHomeENbId().GetLen(), result.GetENb().GetGlobalENbId().GetENbId().GetHomeENbId().GetLen())
}

func Test_perGlobalKpmNodeIDenbCompareBytes(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID4()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(globalKpmnodeID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("GlobalKpmnodeID (ENb) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMnodeIDenb)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
