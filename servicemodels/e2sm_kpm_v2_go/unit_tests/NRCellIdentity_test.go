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

var refPerNrCellID = "00000000  d4 bc 09 00 00                                    |.....|"

func createNrcellIdentity() *e2sm_kpm_v2_go.NrcellIdentity {
	return &e2sm_kpm_v2_go.NrcellIdentity{
		Value: &asn1.BitString{
			Value: []byte{0xd4, 0xbc, 0x09, 0x00, 0x00},
			Len:   36,
		},
	}
}

func Test_perEncodingNrCellIdentity(t *testing.T) {

	nrCgi := createNrcellIdentity()

	per, err := aper.Marshal(nrCgi, nil, nil)
	assert.NilError(t, err)
	t.Logf("NrCellIdentity PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.NrcellIdentity{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("NrCellIdentity PER - decoded\n%v", &result)
	assert.DeepEqual(t, nrCgi.GetValue().GetValue(), result.GetValue().GetValue())
	assert.Equal(t, nrCgi.GetValue().GetLen(), result.GetValue().GetLen())
}

func Test_perNrCellIdentityCompareBytes(t *testing.T) {

	nrCgi := createNrcellIdentity()

	per, err := aper.Marshal(nrCgi, nil, nil)
	assert.NilError(t, err)
	t.Logf("NrCellIdentity PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerNrCellID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
