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
	"math/bits"

	"github.com/shopspring/decimal"
)

// FIXDecimal is a FIX Float Value that implements an arbitrary precision fixed-point decimal.  Implements FieldValue.
type FIXDecimal struct {
	decimal.Decimal

	// Scale is the number of digits after the decimal point when Writing the field value as a FIX value.
	Scale int32
}

// pow10 holds 10^0 through 10^19, the powers of ten that fit in a uint64.
var pow10 = [...]uint64{
	1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}

func (d FIXDecimal) Write() []byte {
	if b, ok := d.appendFast(); ok {
		return b
	}
	return []byte(d.Decimal.StringFixed(d.Scale))
}

// appendFast formats the value with Scale decimal places, rounding half away from zero
// as decimal.StringFixed does, without going through big.Int arithmetic or intermediate strings.
// It reports false when the value or scale is outside the range it handles, and the caller must fall back.
func (d FIXDecimal) appendFast() ([]byte, bool) {
	scale := d.Scale
	if scale < 0 || int(scale) >= len(pow10) {
		return nil, false
	}

	coef := d.Decimal.Coefficient()
	if !coef.IsInt64() {
		return nil, false
	}
	c := coef.Int64()
	neg := c < 0
	mag := uint64(c)
	if neg {
		mag = -mag
	}

	// The value is mag * 10^exp; shift it so that the last digit of mag is the last decimal place.
	shift := int64(d.Decimal.Exponent()) + int64(scale)
	switch {
	case shift > 0:
		if shift >= int64(len(pow10)) {
			if mag == 0 {
				break
			}
			return nil, false
		}
		hi, lo := bits.Mul64(mag, pow10[shift])
		if hi != 0 {
			return nil, false
		}
		mag = lo
	case shift < 0:
		if -shift >= int64(len(pow10)) {
			// Everything is rounded away.
			mag = 0
			break
		}
		div := pow10[-shift]
		q, r := mag/div, mag%div
		if r >= div-r { // r >= div/2, without overflowing for div > 1<<63
			q++
		}
		mag = q
	}

	var digits [20]byte
	n := len(digits)
	for mag > 0 || n > len(digits)-int(scale)-1 {
		n--
		digits[n] = byte('0' + mag%10)
		mag /= 10
	}
	intLen := len(digits) - n - int(scale)
	isZero := true
	for _, ch := range digits[n:] {
		if ch != '0' {
			isZero = false
			break
		}
	}

	size := len(digits) - n
	if scale > 0 {
		size++
	}
	if neg && !isZero {
		size++
	}
	b := make([]byte, 0, size)
	if neg && !isZero {
		b = append(b, '-')
	}
	b = append(b, digits[n:n+intLen]...)
	if scale > 0 {
		b = append(b, '.')
		b = append(b, digits[n+intLen:]...)
	}
	return b, true
}

func (d *FIXDecimal) Read(bytes []byte) (err error) {
	d.Decimal, err = decimal.NewFromString(string(bytes))
	return
}
