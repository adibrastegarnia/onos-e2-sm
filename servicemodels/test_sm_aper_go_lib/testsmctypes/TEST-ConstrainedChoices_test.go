// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func createTestConstrainedChoicesMsg() (*test_sm_ies.TestConstrainedChoices, error) {

	testConstrainedChoices := test_sm_ies.TestConstrainedChoices{
		OtherCattr: "nil",
		ConstrainedChoice1: &test_sm_ies.ConstrainedChoice1{
			ConstrainedChoice1: &test_sm_ies.ConstrainedChoice1_ConstrainedChoice1A{
				ConstrainedChoice1A: 129,
			},
		},
		ConstrainedChoice2: &test_sm_ies.ConstrainedChoice2{
			ConstrainedChoice2: &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2B{
				ConstrainedChoice2B: 4294967296,
			},
		},
		ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3{
			ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3D{
				ConstrainedChoice3D: 2,
			},
		},
		ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4{
			ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A{
				ConstrainedChoice4A: 130,
			},
		},
	}

	return &testConstrainedChoices, nil
}

func Test_xerEncodingTestConstrainedChoices(t *testing.T) {

	testConstrainedChoices, err := createTestConstrainedChoicesMsg()
	assert.NilError(t, err, "Error creating TestConstrainedChoices PDU")

	xer, err := xerEncodeTestConstrainedChoices(testConstrainedChoices)
	assert.NilError(t, err)
	t.Logf("TestConstrainedChoices XER\n%s", string(xer))

	result, err := xerDecodeTestConstrainedChoices(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedChoices XER - decoded\n%v", result)
	assert.Equal(t, testConstrainedChoices.GetOtherCattr(), result.GetOtherCattr())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice1().GetConstrainedChoice1A(), result.GetConstrainedChoice1().GetConstrainedChoice1A())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice2().GetConstrainedChoice2B(), result.GetConstrainedChoice2().GetConstrainedChoice2B())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice3().GetConstrainedChoice3D(), result.GetConstrainedChoice3().GetConstrainedChoice3D())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice4().GetConstrainedChoice4A(), result.GetConstrainedChoice4().GetConstrainedChoice4A())

}

func Test_perEncodingTestConstrainedChoices(t *testing.T) {

	testConstrainedChoices, err := createTestConstrainedChoicesMsg()
	assert.NilError(t, err, "Error creating TestConstrainedChoices PDU")

	per, err := PerEncodeTestConstrainedChoices(testConstrainedChoices)
	assert.NilError(t, err)
	t.Logf("TestConstrainedChoices PER\n%v", hex.Dump(per))

	// Setting ChoiceMap to enable encoding with Go APER library (necessary prerequisite)
	aper.ChoiceMap = test_sm_ies.Choicemap
	// Generating APER bytes with Go APER lib
	perNew, err := aper.Marshal(testConstrainedChoices)
	assert.NilError(t, err)

	//Comparing bytes against each other
	assert.DeepEqual(t, per, perNew)

	result, err := PerDecodeTestConstrainedChoices(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestConstrainedChoices PER - decoded\n%v", result)
	assert.Equal(t, testConstrainedChoices.GetOtherCattr(), result.GetOtherCattr())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice1().GetConstrainedChoice1A(), result.GetConstrainedChoice1().GetConstrainedChoice1A())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice2().GetConstrainedChoice2B(), result.GetConstrainedChoice2().GetConstrainedChoice2B())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice3().GetConstrainedChoice3D(), result.GetConstrainedChoice3().GetConstrainedChoice3D())
	assert.Equal(t, testConstrainedChoices.GetConstrainedChoice4().GetConstrainedChoice4A(), result.GetConstrainedChoice4().GetConstrainedChoice4A())
}
