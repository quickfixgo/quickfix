// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFIXDecimalWrite(t *testing.T) {
	var tests = []struct {
		decimal  FIXDecimal
		expected string
	}{
		{decimal: FIXDecimal{Decimal: decimal.New(-1243456, -4), Scale: 4}, expected: "-124.3456"},
		{decimal: FIXDecimal{Decimal: decimal.New(-1243456, -4), Scale: 5}, expected: "-124.34560"},
		{decimal: FIXDecimal{Decimal: decimal.New(-1243456, -4), Scale: 0}, expected: "-124"},
	}

	for _, test := range tests {
		b := test.decimal.Write()
		assert.Equal(t, test.expected, string(b))
	}
}

func TestFIXDecimalRead(t *testing.T) {
	var tests = []struct {
		bytes       string
		expected    decimal.Decimal
		expectError bool
	}{
		{bytes: "15", expected: decimal.New(15, 0)},
		{bytes: "15.000", expected: decimal.New(15, 0)},
		{bytes: "15.001", expected: decimal.New(15001, -3)},
		{bytes: "-15.001", expected: decimal.New(-15001, -3)},
		{bytes: "blah", expectError: true},
		{bytes: "+200.00", expected: decimal.New(200, 0)},
	}

	for _, test := range tests {
		var field FIXDecimal

		err := field.Read([]byte(test.bytes))
		require.Equal(t, test.expectError, err != nil)

		if !test.expectError {
			assert.True(t, test.expected.Equals(field.Decimal), "Expected %s got %s", test.expected, field.Decimal)
		}
	}
}
