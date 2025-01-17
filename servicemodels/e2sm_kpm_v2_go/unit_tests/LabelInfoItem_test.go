// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerLabelInfoItem = "00000000  3f ff e0 01 02 03 40 40  01 02 03 00 17 7c 18 00  |?.....@@.....|..|\n" +
	"00000010  1e 00 01 70 00 00 18 00  00 00 00 00 7a 00 01 c7  |...p........z...|\n" +
	"00000020  00 03 14 20                                       |... |"

func createLabelInfoItem() *e2sm_kpm_v2_go.LabelInfoItem {

	var br int32 = 25
	var lmm int32 = 1
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	return &e2sm_kpm_v2_go.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2_go.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			SliceId: &e2sm_kpm_v2_go.Snssai{
				SD:  []byte{0x01, 0x02, 0x03},
				SSt: []byte{0x01},
			},
			FiveQi: &e2sm_kpm_v2_go.FiveQi{
				Value: 23,
			},
			QFi: &e2sm_kpm_v2_go.Qfi{
				Value: 62,
			},
			QCi: &e2sm_kpm_v2_go.Qci{
				Value: 24,
			},
			QCimax: &e2sm_kpm_v2_go.Qci{
				Value: 30,
			},
			QCimin: &e2sm_kpm_v2_go.Qci{
				Value: 1,
			},
			ARpmax: &e2sm_kpm_v2_go.Arp{
				Value: 15,
			},
			ARpmin: &e2sm_kpm_v2_go.Arp{
				Value: 1,
			},
			BitrateRange:     &br,
			LayerMuMimo:      &lmm,
			SUm:              &sum,
			DistBinX:         &dbx,
			DistBinY:         &dby,
			DistBinZ:         &dbz,
			PreLabelOverride: &plo,
			StartEndInd:      &seind,
		},
	}
}

func Test_perEncodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	per, err := aper.MarshalWithParams(lii, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("LabelInfoItem PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.LabelInfoItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("LabelInfoItem PER - decoded\n%v", &result)
	assert.DeepEqual(t, lii.GetMeasLabel().GetPlmnId().GetValue(), result.GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, lii.GetMeasLabel().GetSliceId().GetSD(), result.GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, lii.GetMeasLabel().GetSliceId().GetSSt(), result.GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, lii.GetMeasLabel().GetFiveQi().GetValue(), result.GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, lii.GetMeasLabel().GetQFi().GetValue(), result.GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, lii.GetMeasLabel().GetQCi().GetValue(), result.GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, lii.GetMeasLabel().GetQCimax().GetValue(), result.GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, lii.GetMeasLabel().GetQCimin().GetValue(), result.GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, lii.GetMeasLabel().GetARpmax().GetValue(), result.GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, lii.GetMeasLabel().GetARpmin().GetValue(), result.GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, lii.GetMeasLabel().GetBitrateRange(), result.GetMeasLabel().GetBitrateRange())
	assert.Equal(t, lii.GetMeasLabel().GetLayerMuMimo(), result.GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, lii.GetMeasLabel().GetSUm().Number(), result.GetMeasLabel().GetSUm().Number())
	assert.Equal(t, lii.GetMeasLabel().GetDistBinX(), result.GetMeasLabel().GetDistBinX())
	assert.Equal(t, lii.GetMeasLabel().GetDistBinY(), result.GetMeasLabel().GetDistBinY())
	assert.Equal(t, lii.GetMeasLabel().GetDistBinZ(), result.GetMeasLabel().GetDistBinZ())
	assert.Equal(t, lii.GetMeasLabel().GetPreLabelOverride().Number(), result.GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, lii.GetMeasLabel().GetStartEndInd().Number(), result.GetMeasLabel().GetStartEndInd().Number())
}

func Test_perLabelInfoItemCompareBytes(t *testing.T) {

	lii := createLabelInfoItem()

	per, err := aper.MarshalWithParams(lii, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("LabelInfoItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerLabelInfoItem)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
