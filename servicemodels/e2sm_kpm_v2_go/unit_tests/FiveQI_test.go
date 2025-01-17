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

var refPerFiveQI = "00000000  00 0c                                             |..|"

func createFiveQi() *e2sm_kpm_v2_go.FiveQi {

	return &e2sm_kpm_v2_go.FiveQi{
		Value: 12,
	}
}

func Test_perEncodingFiveQi(t *testing.T) {

	fqi := createFiveQi()

	per, err := aper.Marshal(fqi, nil, nil)
	assert.NilError(t, err)
	t.Logf("FiveQI PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.FiveQi{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("FiveQI PER - decoded\n%v", &result)
	assert.Equal(t, fqi.GetValue(), result.GetValue())
}

func Test_perFiveQiCompareBytes(t *testing.T) {

	fqi := createFiveQi()

	per, err := aper.Marshal(fqi, nil, nil)
	assert.NilError(t, err)
	t.Logf("FiveQI PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerFiveQI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
