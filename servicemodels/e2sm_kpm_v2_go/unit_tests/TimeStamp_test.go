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

var refPerTimeStamp = "00000000  01 02 03 04                                       |....|"

func Test_perEncodingTimeStamp(t *testing.T) {

	stamp := []byte{0x01, 0x02, 0x03, 0x04}
	timeStamp := &e2sm_kpm_v2_go.TimeStamp{
		Value: stamp,
	}

	per, err := aper.Marshal(timeStamp, nil, nil)
	assert.NilError(t, err)
	t.Logf("TimeStamp PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TimeStamp{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TimeStamp PER - decoded\n%v", &result)
	assert.DeepEqual(t, timeStamp.GetValue(), result.GetValue())
}

func Test_perTimeStampCompareBytes(t *testing.T) {

	stamp := []byte{0x01, 0x02, 0x03, 0x04}
	timeStamp := &e2sm_kpm_v2_go.TimeStamp{
		Value: stamp,
	}

	per, err := aper.Marshal(timeStamp, nil, nil)
	assert.NilError(t, err)
	t.Logf("TimeStamp PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTimeStamp)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
