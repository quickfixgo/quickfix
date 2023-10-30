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

import "github.com/shopspring/decimal"

// FIXDecimal is a FIX Float Value that implements an arbitrary precision fixed-point decimal.  Implements FieldValue.
type FIXDecimal struct {
	decimal.Decimal

	// Scale is the number of digits after the decimal point when Writing the field value as a FIX value.
	Scale int32
}

func (d FIXDecimal) Write() []byte {
	return []byte(d.Decimal.StringFixed(d.Scale))
}

func (d *FIXDecimal) Read(bytes []byte) (err error) {
	d.Decimal, err = decimal.NewFromString(string(bytes))
	return
}
