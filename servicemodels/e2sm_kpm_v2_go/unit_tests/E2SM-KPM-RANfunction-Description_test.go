// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerE2SmKpmRanFunctionDescriptionFull = "00000000  74 04 6f 6e 66 00 00 05  6f 69 64 31 32 33 07 00  |t.onf...oid123..|\n" +
	"00000010  73 6f 6d 65 44 65 73 63  72 69 70 74 69 6f 6e 00  |someDescription.|\n" +
	"00000020  15 00 00 43 00 21 22 23  00 d4 bc 08 80 30 39 20  |...C.!\"#.....09 |\n" +
	"00000030  1a 85 00 00 00 00 03 4f  4e 46 00 21 22 23 12 f0  |.......ONF.!\"#..|\n" +
	"00000040  de bc 50 00 0b 01 00 6f  6e 66 00 0f 00 0b 01 00  |..P....onf......|\n" +
	"00000050  6f 6e 66 00 0f 00 00 41  a0 4f 70 65 6e 4e 65 74  |onf....A.OpenNet|\n" +
	"00000060  77 6f 72 6b 69 6e 67 00  00 17 00 02 00 01        |working.......|"

var refPerE2SmKpmRanFunctionDescriptionNodeListOnly = "00000000  44 04 6f 6e 66 00 00 05  6f 69 64 31 32 33 07 00  |D.onf...oid123..|\n" +
	"00000010  73 6f 6d 65 44 65 73 63  72 69 70 74 69 6f 6e 00  |someDescription.|\n" +
	"00000020  15 00 00 43 00 21 22 23  00 d4 bc 08 80 30 39 20  |...C.!\"#.....09 |\n" +
	"00000030  1a 85 00 00 00 00 03 4f  4e 46 00 21 22 23 12 f0  |.......ONF.!\"#..|\n" +
	"00000040  de bc 50                                          |..P|"

var refPerE2SmKpmRanFunctionDescriptionEventListOnly = "00000000  24 04 6f 6e 66 00 00 05  6f 69 64 31 32 33 07 00  |$.onf...oid123..|\n" +
	"00000010  73 6f 6d 65 44 65 73 63  72 69 70 74 69 6f 6e 00  |someDescription.|\n" +
	"00000020  15 00 0b 0c 00 53 6f 6d  65 52 65 61 6c 6c 79 43  |.....SomeReallyC|\n" +
	"00000030  6f 6f 6c 44 65 73 63 72  69 70 74 69 6f 6e 00 0f  |oolDescription..|"

var refPerE2SmKpmRanFunctionDescriptionReportListOnly = "00000000  14 04 6f 6e 66 00 00 05  6f 69 64 31 32 33 07 00  |..onf...oid123..|\n" +
	"00000010  73 6f 6d 65 44 65 73 63  72 69 70 74 69 6f 6e 00  |someDescription.|\n" +
	"00000020  15 00 0b 01 00 6f 6e 66  00 0f 00 00 41 a0 4f 70  |.....onf....A.Op|\n" +
	"00000030  65 6e 4e 65 74 77 6f 72  6b 69 6e 67 00 00 17 00  |enNetworking....|\n" +
	"00000040  02 00 01                                          |...|"

var refPerE2SmKpmRanFunctionDescriptionMndtOnly = "00000000  00 04 6f 6e 66 00 00 05  6f 69 64 31 32 33 07 00  |..onf...oid123..|\n" +
	"00000010  73 6f 6d 65 44 65 73 63  72 69 70 74 69 6f 6e     |someDescription|"

func createE2SmKpmRanFunctionDescription() (*e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription, error) {

	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	cellIDbits := []byte{0x12, 0xF0, 0xDE, 0xBC, 0x50}
	cellGlobalID, err := pdubuilder.CreateCellGlobalIDNRCGI(plmnID, cellIDbits) // 36 bits
	if err != nil {
		return nil, err
	}

	var cellObjID = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	if err != nil {
		return nil, err
	}
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2sm_kpm_v2_go.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := pdubuilder.CreateRicKpmnodeItem(globalKpmnodeID).SetCellMeasurementObjectList(cmol)

	rknl := make([]*e2sm_kpm_v2_go.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName = "onf"
	var ricFormatType int32 = 15
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2_go.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	var measTypeName = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := pdubuilder.CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2_go.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 1
	var indHdrFormat int32 = 2
	rrsi := pdubuilder.CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2_go.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd).SetRanFunctionInstance(rfi).SetRicKpmNodeList(rknl).SetRicReportStyleList(rrsl).SetRicEventTriggerStyleList(retsl)
	fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func createE2SmKpmRanFunctionDescriptionReportList() (*e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription, error) {

	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"
	var rfi int32 = 21

	var ricStyleType int32 = 11
	var ricStyleName = "onf"
	var ricFormatType int32 = 15

	measInfoActionList := e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	var measTypeName = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := pdubuilder.CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2_go.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 1
	var indHdrFormat int32 = 2
	rrsi := pdubuilder.CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2_go.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd).SetRanFunctionInstance(rfi).SetRicReportStyleList(rrsl)
	fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func createE2SmKpmRanFunctionDescriptionEventList() (*e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription, error) {

	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"
	var rfi int32 = 21

	var ricStyleType int32 = 11
	var ricStyleName = "SomeReallyCoolDescription"
	var ricFormatType int32 = 15
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2_go.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd).SetRanFunctionInstance(rfi).SetRicEventTriggerStyleList(retsl)
	fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func createE2SmKpmRanFunctionDescriptionNodeList() (*e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription, error) {

	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	cellIDbits := []byte{0x12, 0xF0, 0xDE, 0xBC, 0x50}
	cellGlobalID, err := pdubuilder.CreateCellGlobalIDNRCGI(plmnID, cellIDbits) // 36 bits
	if err != nil {
		return nil, err
	}

	var cellObjID = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	if err != nil {
		return nil, err
	}
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2sm_kpm_v2_go.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := pdubuilder.CreateRicKpmnodeItem(globalKpmnodeID).SetCellMeasurementObjectList(cmol)

	rknl := make([]*e2sm_kpm_v2_go.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd).SetRanFunctionInstance(rfi).SetRicKpmNodeList(rknl)
	fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func createE2SmKpmRanFunctionDescriptionMndtOnly() (*e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription, error) {

	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd)
	fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func Test_perEncodingE2SmKpmRanFunctionDescription(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescription()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription PER - decoded\n%v", result)
}

func Test_perE2SmKpmRanFunctionDescriptionCompareBytes(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescription()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmRanFunctionDescriptionFull)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)

	result, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription PER - decoded\n%v", result)
}

func Test_perE2SmKpmRanFunctionDescriptionNodeListCompareBytes(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescriptionNodeList()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmRanFunctionDescriptionNodeListOnly)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)

	result, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription PER - decoded\n%v", result)
}

func Test_perE2SmKpmRanFunctionDescriptionEventListCompareBytes(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescriptionEventList()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmRanFunctionDescriptionEventListOnly)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)

	result, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription PER - decoded\n%v", result)
}

func Test_perE2SmKpmRanFunctionDescriptionReportListCompareBytes(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescriptionReportList()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmRanFunctionDescriptionReportListOnly)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)

	result, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription PER - decoded\n%v", result)
}

func Test_perEncodingE2SmKpmRanFunctionDescriptionMndtOnly(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescriptionMndtOnly()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription (mandatory part only) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription (mandatory part only) PER - decoded\n%v", result)
}

func Test_perE2SmKpmRanFunctionDescriptionMndtOnlyCompareBytes(t *testing.T) {

	rfd, err := createE2SmKpmRanFunctionDescriptionMndtOnly()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-RANfunctionDescription (mandatory part only) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmRanFunctionDescriptionMndtOnly)
	assert.NilError(t, err)
	t.Logf("Extracted bytes are \n%v", hex.Dump(perRefBytes))
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perDecodeRadysisBytes(t *testing.T) {

	radisysBytesRanFunctionDefinition := []byte{0x70, 0x18, 0x4F, 0x52, 0x41, 0x4E, 0x2D, 0x45, 0x32, 0x53, 0x4D, 0x2D, 0x4B, 0x50, 0x4D, 0x00,
		0x00, 0x18, 0x31, 0x2E, 0x33, 0x2E, 0x36, 0x2E, 0x31, 0x2E, 0x34, 0x2E, 0x31, 0x2E, 0x35, 0x33,
		0x31, 0x34, 0x38, 0x2E, 0x31, 0x2E, 0x32, 0x2E, 0x32, 0x2E, 0x32, 0x05, 0x00, 0x4B, 0x50, 0x4D,
		0x20, 0x6D, 0x6F, 0x6E, 0x69, 0x74, 0x6F, 0x72, 0x00, 0x00, 0x40, 0x00, 0x13, 0xF1, 0x84, 0x50,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x30, 0x00, 0x13, 0xF1, 0x84, 0x00, 0x00,
		0x00, 0x00, 0x10, 0x00, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6F, 0x64, 0x69, 0x63, 0x20,
		0x72, 0x65, 0x70, 0x6F, 0x72, 0x74, 0x00, 0x01, 0x00, 0x03, 0x09, 0x00, 0x45, 0x32, 0x20, 0x4E,
		0x6F, 0x64, 0x65, 0x20, 0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x00,
		0x01, 0x00, 0x07, 0x42, 0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74,
		0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x53, 0x75, 0x6D, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52,
		0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2E,
		0x53, 0x75, 0x6D, 0x00, 0x00, 0x01, 0x42, 0xA0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E,
		0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x53, 0x75, 0x6D, 0x00, 0x00,
		0x02, 0x43, 0xC0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45, 0x73, 0x74,
		0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x72, 0x65, 0x63, 0x6F, 0x6E, 0x66, 0x69, 0x67, 0x46, 0x61,
		0x69, 0x6C, 0x00, 0x00, 0x03, 0x43, 0x00, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52,
		0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x48, 0x4F, 0x46, 0x61, 0x69, 0x6C,
		0x00, 0x00, 0x04, 0x42, 0xE0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45,
		0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x4F, 0x74, 0x68, 0x65, 0x72, 0x00, 0x00, 0x05,
		0x41, 0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x4D, 0x65, 0x61, 0x6E, 0x00, 0x00,
		0x06, 0x41, 0x40, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x4D, 0x61, 0x78, 0x00, 0x00,
		0x07, 0x00, 0x01, 0x00, 0x01}
	t.Logf("Radysis E2SmKpmRanfunctionDescription PER\n%v", hex.Dump(radisysBytesRanFunctionDefinition))

	result, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(radisysBytesRanFunctionDefinition)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-RANfunctionDescription (Radisys) PER - decoded\n%v", result)
}
