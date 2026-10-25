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
	"math"
	"math/rand"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkFIXDecimalRead(b *testing.B) {
	var dec FIXDecimal
	byt := []byte("-124.3456")
	for i := 0; i < b.N; i++ {
		if err := dec.Read(byt); err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkFIXDecimalWrite(b *testing.B) {
	decValue, err := decimal.NewFromString("-124.3456")
	if err != nil {
		b.FailNow()
	}
	dec := FIXDecimal{Decimal: decValue, Scale: 5}
	for i := 0; i < b.N; i++ {
		if byt := dec.Write(); len(byt) == 0 {
			b.FailNow()
		}
	}
}

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

func TestFIXDecimalWriteMatchesStringFixed(t *testing.T) {
	coefs := []int64{
		0, 1, -1, 5, -5, 9, 10, 49, 50, 51, -49, -50, -51, 99, 100, 995, -995, 1243456, -1243456,
		math.MaxInt32, math.MinInt32, math.MaxInt64, math.MinInt64, math.MaxInt64 - 1, math.MinInt64 + 1,
		1e18, -1e18, 9999999999999999999 / 10,
	}
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 200; i++ {
		coefs = append(coefs, r.Int63()>>uint(r.Intn(63)), -(r.Int63() >> uint(r.Intn(63))))
	}

	for _, c := range coefs {
		for exp := int32(-25); exp <= 25; exp++ {
			for _, scale := range []int32{-2, 0, 1, 2, 4, 8, 12, 18, 19, 20} {
				dec := FIXDecimal{Decimal: decimal.New(c, exp), Scale: scale}
				expected := dec.Decimal.StringFixed(scale)
				assert.Equal(t, expected, string(dec.Write()), "coef=%d exp=%d scale=%d", c, exp, scale)
			}
		}
	}

	big := decimal.RequireFromString("123456789012345678901234567890.123456789")
	for _, scale := range []int32{0, 4, 9, 12} {
		dec := FIXDecimal{Decimal: big, Scale: scale}
		assert.Equal(t, big.StringFixed(scale), string(dec.Write()))
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
			assert.True(t, test.expected.Equal(field.Decimal), "Expected %s got %s", test.expected, field.Decimal)
		}
	}
}
