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

var refPerGnbCuUpID = "00000000  20 04 d2                                          | ..|"

func createGnbCuUpID() *e2sm_kpm_v2_go.GnbCuUpId {

	return &e2sm_kpm_v2_go.GnbCuUpId{
		Value: 1234,
	}
}

func Test_perEncodingGnbCuUpID(t *testing.T) {

	gnbCuUpID := createGnbCuUpID()

	per, err := aper.Marshal(gnbCuUpID, nil, nil)
	assert.NilError(t, err)
	t.Logf("GnbCuUpID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GnbCuUpId{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("GnbCuUpID PER - decoded\n%v", result.GetValue())
	assert.Equal(t, gnbCuUpID.GetValue(), result.GetValue())
}

func Test_perGnbCuUpIDCompareBytes(t *testing.T) {

	gnbCuUpID := createGnbCuUpID()

	per, err := aper.Marshal(gnbCuUpID, nil, nil)
	assert.NilError(t, err)
	t.Logf("GnbCuUpID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGnbCuUpID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
