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

var refPerArpLB = "00000000  00                                                |.|"
var refPerArpUB = "00000000  70                                                |p|"
var refPerArpExt = "00000000  80 01 79                                          |..y|"

func Test_perEncodingArpLB(t *testing.T) {

	arp := &e2sm_kpm_v2_go.Arp{
		Value: 15,
	}

	per, err := aper.Marshal(arp, nil, nil)
	assert.NilError(t, err)
	t.Logf("ARP PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.Arp{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	t.Logf("ARP PER - decoded\n%v", &result)
	assert.Equal(t, arp.GetValue(), result.GetValue())
}

func Test_perArpCompareBytesLB(t *testing.T) {

	arp := &e2sm_kpm_v2_go.Arp{
		Value: 1,
	}

	per, err := aper.Marshal(arp, nil, nil)
	assert.NilError(t, err)
	t.Logf("ARP PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerArpLB)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingArpUB(t *testing.T) {

	arp := &e2sm_kpm_v2_go.Arp{
		Value: 15,
	}

	per, err := aper.Marshal(arp, nil, nil)
	assert.NilError(t, err)
	t.Logf("ARP PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.Arp{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("ARP PER - decoded\n%v", &result)
	assert.Equal(t, arp.GetValue(), result.GetValue())
}

func Test_perArpCompareBytesUB(t *testing.T) {

	arp := &e2sm_kpm_v2_go.Arp{
		Value: 15,
	}

	per, err := aper.Marshal(arp, nil, nil)
	assert.NilError(t, err)
	t.Logf("ARP PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerArpUB)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingArpExt(t *testing.T) {

	arp := &e2sm_kpm_v2_go.Arp{
		Value: 121,
	}

	per, err := aper.Marshal(arp, nil, nil)
	assert.NilError(t, err)
	t.Logf("ARP PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.Arp{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("ARP PER - decoded\n%v", &result)
	assert.Equal(t, arp.GetValue(), result.GetValue())
}

func Test_perArpCompareBytesExt(t *testing.T) {

	arp := &e2sm_kpm_v2_go.Arp{
		Value: 121,
	}

	per, err := aper.Marshal(arp, nil, nil)
	assert.NilError(t, err)
	t.Logf("ARP PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerArpExt)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
