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
