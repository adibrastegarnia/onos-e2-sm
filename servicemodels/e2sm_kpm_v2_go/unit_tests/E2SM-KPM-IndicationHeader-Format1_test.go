// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerE2SmKpmIndicationHeaderFormat1 = "00000000  7c 21 22 23 24 18 74 78  74 00 00 03 4f 4e 46 40  ||!\"#$.txt...ONF@|\n" +
	"00000010  73 6f 6d 65 54 79 70 65  06 6f 6e 66 0c 21 22 23  |someType.onf.!\"#|\n" +
	"00000020  30 d4 bc 09 00 15 00 20                           |0...... |"

func createE2SmKpmIndicationHeaderFormat1() (*e2sm_kpm_v2_go.E2SmKpmIndicationHeaderFormat1, error) {

	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x09, 0x00},
		Len:   28,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 21
	var gnbDuID int64 = 32
	var fileFormatVersion = "txt"
	var senderName = "ONF"
	var senderType = "someType"
	var vendorName = "onf"

	globalKpmNodeID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	if err != nil {
		return nil, err
	}
	globalKpmNodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmNodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}
	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(timeStamp)
	if err != nil {
		return nil, err
	}
	newE2SmKpmPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetSenderType(senderType).SetVendorName(vendorName).SetGlobalKPMnodeID(globalKpmNodeID)

	return newE2SmKpmPdu.GetIndicationHeaderFormats().GetIndicationHeaderFormat1(), nil
}

func Test_perEncodingE2SmKpmIndicationHeaderFormat1(t *testing.T) {

	ihf1, err := createE2SmKpmIndicationHeaderFormat1()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ihf1, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader-Format1 PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmIndicationHeaderFormat1{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-IndicationHeader-Format1 PER - decoded\n%v", &result)
	assert.Equal(t, ihf1.GetFileFormatversion(), result.GetFileFormatversion())
	assert.Equal(t, ihf1.GetKpmNodeId().GetGNb().GetGNbDuId().GetValue(), result.GetKpmNodeId().GetGNb().GetGNbDuId().GetValue())
	assert.Equal(t, ihf1.GetKpmNodeId().GetGNb().GetGNbCuUpId().GetValue(), result.GetKpmNodeId().GetGNb().GetGNbCuUpId().GetValue())
	assert.DeepEqual(t, ihf1.GetKpmNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue(), result.GetKpmNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
	assert.DeepEqual(t, ihf1.GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), result.GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())
	assert.Equal(t, ihf1.GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), result.GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	assert.Equal(t, ihf1.GetSenderName(), result.GetSenderName())
	assert.Equal(t, ihf1.GetSenderType(), result.GetSenderType())
	assert.DeepEqual(t, ihf1.GetColletStartTime().GetValue(), result.GetColletStartTime().GetValue())
	assert.Equal(t, ihf1.GetVendorName(), result.GetVendorName())
}

func Test_perE2SmKpmIndicationHeaderFormat1CompareBytes(t *testing.T) {

	ihf1, err := createE2SmKpmIndicationHeaderFormat1()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ihf1, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader-Format1 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmIndicationHeaderFormat1)
	t.Logf("E2SM-KPM-IndicationHeader-Format1 PER\n%v", hex.Dump(perRefBytes))
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
