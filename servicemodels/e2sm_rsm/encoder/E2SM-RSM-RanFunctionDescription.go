// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmRsmRanFunctionDescription(rfd *e2sm_rsm_ies.E2SmRsmRanfunctionDescription) ([]byte, error) {

	log.Debugf("Obtained E2SM-RSM-RanFunctionDescription message is\n%v", rfd)

	per, err := aper.MarshalWithParams(rfd, "valueExt", e2sm_rsm_ies.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RSM-RanFunctionDescription PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRsmRanFunctionDescription(per []byte) (*e2sm_rsm_ies.E2SmRsmRanfunctionDescription, error) {

	log.Debugf("Obtained E2SM-RSM-RanFunctionDescription PER bytes are\n%v", hex.Dump(per))

	result := e2sm_rsm_ies.E2SmRsmRanfunctionDescription{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_rsm_ies.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RSM-RanFunctionDescription from PER is\n%v", &result)

	return &result, nil
}
