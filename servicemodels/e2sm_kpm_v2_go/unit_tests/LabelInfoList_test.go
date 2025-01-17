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

var refPerLabelInfoList = "00000000  02 3f ff e0 01 02 03 40  40 01 02 03 00 17 7c 18  |.?.....@@.....|.|\n" +
	"00000010  00 1e 00 01 70 00 00 18  00 00 00 00 00 7a 00 01  |....p........z..|\n" +
	"00000020  c7 00 03 14 27 ff fc 01  02 03 40 40 01 02 03 00  |....'.....@@....|\n" +
	"00000030  17 7c 18 00 1e 00 01 70  00 00 18 00 00 00 00 00  |.|.....p........|\n" +
	"00000040  7a 00 01 c7 00 03 14 20                           |z...... |"

func createLabelInfoList() (*e2sm_kpm_v2_go.LabelInfoList, error) {

	labelInfoList := &e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}

	var br int32 = 25
	var lmm int32 = 1
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	item := &e2sm_kpm_v2_go.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2_go.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
				Value: []byte{0x1, 0x2, 0x3},
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
	labelInfoList.Value = append(labelInfoList.Value, item)
	labelInfoList.Value = append(labelInfoList.Value, item)

	//if err := labelInfoList.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating labelInfoList %s", err.Error())
	//}
	return labelInfoList, nil
}

func Test_perEncodingLabelInfoList(t *testing.T) {

	lil, err := createLabelInfoList()
	assert.NilError(t, err)

	per, err := aper.MarshalWithParams(lil, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("LabelInfoList PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.LabelInfoList{}
	err = aper.UnmarshalWithParams(per, &result, "", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("LabelInfoList PER - decoded\n%v", &result)
	assert.DeepEqual(t, lil.GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, lil.GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, lil.GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, lil.GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.DeepEqual(t, lil.GetValue()[1].GetMeasLabel().GetPlmnId().GetValue(), result.GetValue()[1].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, lil.GetValue()[1].GetMeasLabel().GetSliceId().GetSD(), result.GetValue()[1].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, lil.GetValue()[1].GetMeasLabel().GetSliceId().GetSSt(), result.GetValue()[1].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetFiveQi().GetValue(), result.GetValue()[1].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetQFi().GetValue(), result.GetValue()[1].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetQCi().GetValue(), result.GetValue()[1].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetQCimax().GetValue(), result.GetValue()[1].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetQCimin().GetValue(), result.GetValue()[1].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetARpmax().GetValue(), result.GetValue()[1].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetARpmin().GetValue(), result.GetValue()[1].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetBitrateRange(), result.GetValue()[1].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetLayerMuMimo(), result.GetValue()[1].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetSUm().Number(), result.GetValue()[1].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetDistBinX(), result.GetValue()[1].GetMeasLabel().GetDistBinX())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetDistBinY(), result.GetValue()[1].GetMeasLabel().GetDistBinY())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetDistBinZ(), result.GetValue()[1].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetPreLabelOverride().Number(), result.GetValue()[1].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, lil.GetValue()[1].GetMeasLabel().GetStartEndInd().Number(), result.GetValue()[1].GetMeasLabel().GetStartEndInd().Number())
}

func Test_perLabelInfoListCompareBytes(t *testing.T) {

	lil, err := createLabelInfoList()
	assert.NilError(t, err)

	per, err := aper.MarshalWithParams(lil, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("LabelInfoList PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerLabelInfoList)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
