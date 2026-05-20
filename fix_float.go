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
	"fmt"
	"strconv"
)

// FIXFloat is a FIX Float Value, implements FieldValue.
type FIXFloat float64

// Float64 converts the FIXFloat value to float64.
func (f FIXFloat) Float64() float64 { return float64(f) }

func (f *FIXFloat) Read(bytes []byte) error {
	val, err := strconv.ParseFloat(string(bytes), 64)
	if err != nil {
		return err
	}

	// `strconv` allows values like "+100.00", which is not allowed for FIX float types.
	for _, b := range bytes {
		if b != '.' && b != '-' && !isDecimal(b) {
			return fmt.Errorf("invalid value %v", string(bytes))
		}
	}

	*f = FIXFloat(val)

	return err
}

func (f FIXFloat) Write() []byte {
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64))
}

func isDecimal(b byte) bool {
	return '0' <= b && b <= '9'
}
