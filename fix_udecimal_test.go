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

	decimal "github.com/quagmt/udecimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkFIXUDecimalRead(b *testing.B) {
	var dec FIXUDecimal
	byt := []byte("-124.3456")
	for i := 0; i < b.N; i++ {
		if err := dec.Read(byt); err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkFIXUDecimalWrite(b *testing.B) {
	dec := FIXUDecimal{Decimal: decimal.MustParse("-124.3456"), Scale: 5}
	for i := 0; i < b.N; i++ {
		if byt := dec.Write(); len(byt) == 0 {
			b.FailNow()
		}
	}
}

func TestFIXUDecimalWrite(t *testing.T) {
	var tests = []struct {
		decimal  FIXUDecimal
		expected string
	}{
		{decimal: FIXUDecimal{Decimal: decimal.MustParse("-124.3456"), Scale: 4}, expected: "-124.3456"},
		{decimal: FIXUDecimal{Decimal: decimal.MustParse("-124.3456"), Scale: 5}, expected: "-124.34560"},
		{decimal: FIXUDecimal{Decimal: decimal.MustParse("-124.3456"), Scale: 0}, expected: "-124"},
	}

	for _, test := range tests {
		b := test.decimal.Write()
		assert.Equal(t, test.expected, string(b))
	}
}

func TestFIXUDecimalRead(t *testing.T) {
	var tests = []struct {
		bytes       string
		expected    decimal.Decimal
		expectError bool
	}{
		{bytes: "15", expected: decimal.MustParse("15")},
		{bytes: "15.000", expected: decimal.MustParse("15")},
		{bytes: "15.001", expected: decimal.MustParse("15.001")},
		{bytes: "-15.001", expected: decimal.MustParse("-15.001")},
		{bytes: "blah", expectError: true},
		{bytes: "+200.00", expected: decimal.MustParse("200")},
	}

	for _, test := range tests {
		var field FIXUDecimal

		err := field.Read([]byte(test.bytes))
		require.Equal(t, test.expectError, err != nil)

		if !test.expectError {
			assert.True(t, test.expected.Equal(field.Decimal), "Expected %s got %s", test.expected, field.Decimal)
		}
	}
}
