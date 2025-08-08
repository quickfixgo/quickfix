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
	"errors"
	"strconv"
)

// parseUInt is similar to the function in strconv, but is tuned for uint64s appearing in FIX field types.
func parseUInt64(d []byte) (n uint64, err error) {
	if len(d) == 0 {
		err = errors.New("empty bytes")
		return
	}

	for _, dec := range d {
		if dec < ascii0 || dec > ascii9 {
			err = errors.New("invalid format")
			return
		}

		n = n*10 + (uint64(dec) - ascii0)
	}

	return
}

// FIXUInt64 is a FIX Uint64 Value, implements FieldValue.
type FIXUint64 uint64

// Uint64 converts the FIXUint64 value to uint64.
func (f FIXUint64) Uint64() uint64 { return uint64(f) }

func (f *FIXUint64) Read(bytes []byte) error {
	i, err := parseUInt64(bytes)
	if err != nil {
		return err
	}
	*f = FIXUint64(i)
	return nil
}

func (f FIXUint64) Write() []byte {
	return strconv.AppendInt(nil, int64(f), 10)
}
