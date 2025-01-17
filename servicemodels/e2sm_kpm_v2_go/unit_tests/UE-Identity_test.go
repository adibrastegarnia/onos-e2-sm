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

var refPerUeID = "00000000  06 53 6f 6d 65 55 45                              |.SomeUE|"

func Test_perEncodeUeIdentity(t *testing.T) {

	ueIdentity := &e2sm_kpm_v2_go.UeIdentity{
		Value: []byte("SomeUE"),
	}

	per, err := aper.Marshal(ueIdentity, nil, nil)
	assert.NilError(t, err)
	t.Logf("UE-Identity PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.UeIdentity{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("UE-Identity PER - decoded\n%v", &result)
	assert.DeepEqual(t, ueIdentity.GetValue(), result.GetValue())
}

func Test_perUeIdentityCompareBytes(t *testing.T) {

	ueIdentity := &e2sm_kpm_v2_go.UeIdentity{
		Value: []byte("SomeUE"),
	}

	per, err := aper.Marshal(ueIdentity, nil, nil)
	assert.NilError(t, err)
	t.Logf("UE-Identity PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerUeID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
