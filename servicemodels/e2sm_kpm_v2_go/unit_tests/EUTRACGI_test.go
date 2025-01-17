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

var refPerEutraCGI = "00000000  00 21 22 23 d4 bc 09 00                           |.!\"#....|"

func createEutracgi() *e2sm_kpm_v2_go.Eutracgi {

	return &e2sm_kpm_v2_go.Eutracgi{
		PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		EUtracellIdentity: &e2sm_kpm_v2_go.EutracellIdentity{
			Value: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x09, 0x00},
				Len:   28,
			},
		},
	}
}

func Test_perEncodingEutraCGI(t *testing.T) {

	eCgi := createEutracgi()

	per, err := aper.MarshalWithParams(eCgi, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("EUTRACGI PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.Eutracgi{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("EUTRACGI PER - decoded\n%v", &result)
	assert.DeepEqual(t, eCgi.GetPLmnIdentity().GetValue(), result.GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, eCgi.GetEUtracellIdentity().GetValue().GetValue(), result.GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, eCgi.GetEUtracellIdentity().GetValue().GetLen(), result.GetEUtracellIdentity().GetValue().GetLen())
}

func Test_perEutraCGICompareBytes(t *testing.T) {

	eCgi := createEutracgi()

	per, err := aper.MarshalWithParams(eCgi, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("EUTRACGI PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEutraCGI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
