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

func PerEncodeE2SmRsmIndicationMessage(im *e2sm_rsm_ies.E2SmRsmIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-RSM-IndicationMessage message is\n%v", im)

	per, err := aper.MarshalWithParams(im, "choiceExt", e2sm_rsm_ies.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RSM-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRsmIndicationMessage(per []byte) (*e2sm_rsm_ies.E2SmRsmIndicationMessage, error) {

	log.Debugf("Obtained E2SM-RSM-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2sm_rsm_ies.E2SmRsmIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_rsm_ies.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RSM-IndicationMessage from PER is\n%v", &result)

	return &result, nil
}
