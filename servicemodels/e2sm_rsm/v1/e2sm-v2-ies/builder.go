// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2sm_v2_ies

func (m *RanfunctionName) SetRanFunctionInstance(rfi int32) *RanfunctionName {
	m.RanFunctionInstance = &rfi
	return m
}
