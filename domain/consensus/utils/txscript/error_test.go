// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import (
	"testing"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   ErrorCode
		want string
	}{
		{ErrInternal, "ErrInternal"},
		{ErrInvalidFlags, "ErrInvalidFlags"},
		{ErrInvalidIndex, "ErrInvalidIndex"},
		{ErrUnsupportedAddress, "ErrUnsupportedAddress"},
		{ErrTooManyRequiredSigs, "ErrTooManyRequiredSigs"},
		{ErrNotMultisigScript, "ErrNotMultisigScript"},
		{ErrEarlyReturn, "ErrEarlyReturn"},
		{ErrEmptyStack, "ErrEmptyStack"},
		{ErrEvalFalse, "ErrEvalFalse"},
		{ErrScriptUnfinished, "ErrScriptUnfinished"},
		{ErrInvalidProgramCounter, "ErrInvalidProgramCounter"},
		{ErrScriptTooBig, "ErrScriptTooBig"},
		{ErrElementTooBig, "ErrElementTooBig"},
		{ErrTooManyOperations, "ErrTooManyOperations"},
		{ErrStackOverflow, "ErrStackOverflow"},
		{ErrInvalidPubKeyCount, "ErrInvalidPubKeyCount"},
		{ErrInvalidSignatureCount, "ErrInvalidSignatureCount"},
		{ErrNumberTooBig, "ErrNumberTooBig"},
		{ErrVerify, "ErrVerify"},
		{ErrEqualVerify, "ErrEqualVerify"},
		{ErrNumEqualVerify, "ErrNumEqualVerify"},
		{ErrCheckSigVerify, "ErrCheckSigVerify"},
		{ErrCheckMultiSigVerify, "ErrCheckMultiSigVerify"},
		{ErrDisabledOpcode, "ErrDisabledOpcode"},
		{ErrReservedOpcode, "ErrReservedOpcode"},
		{ErrMalformedPush, "ErrMalformedPush"},
		{ErrInvalidStackOperation, "ErrInvalidStackOperation"},
		{ErrUnbalancedConditional, "ErrUnbalancedConditional"},
		{ErrMinimalData, "ErrMinimalData"},
		{ErrInvalidSigHashType, "ErrInvalidSigHashType"},
		{ErrSigLength, "ErrSigLength"},
		{ErrSigHighS, "ErrSigHighS"},
		{ErrNotPushOnly, "ErrNotPushOnly"},
		{ErrPubKeyFormat, "ErrPubKeyFormat"},
		{ErrCleanStack, "ErrCleanStack"},
		{ErrNullFail, "ErrNullFail"},
		{ErrNegativeLockTime, "ErrNegativeLockTime"},
		{ErrUnsatisfiedLockTime, "ErrUnsatisfiedLockTime"},
		{ErrMinimalIf, "ErrMinimalIf"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(numErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   Error
		want string
	}{
		{
			Error{Description: "some error"},
			"some error",
		},
		{
			Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

