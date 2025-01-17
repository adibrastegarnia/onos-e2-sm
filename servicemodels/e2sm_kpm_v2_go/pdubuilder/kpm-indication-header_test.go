// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmIndicationHeader(t *testing.T) {
	bs := asn1.BitString{
		Value: []byte{0x9b, 0xcd, 0x40},
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	var fileFormatVersion = "txt"
	var senderName = "ONF"
	var senderType = "someType"
	var vendorName = "onf"

	globalKpmNodeID, err := CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	globalKpmNodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmNodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmIndicationHeader(timeStamp)
	assert.NilError(t, err)
	newE2SmKpmPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetSenderType(senderType).SetVendorName(vendorName).SetGlobalKPMnodeID(globalKpmNodeID)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
