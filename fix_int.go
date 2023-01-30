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

const (
	// ASCII - char.
	asciiMinus = 45

	// ASCII numbers 0-9.
	ascii0 = 48
	ascii9 = 57
)

// atoi is similar to the function in strconv, but is tuned for ints appearing in FIX field types.
func atoi(d []byte) (int, error) {
	if d[0] == asciiMinus {
		n, err := parseUInt(d[1:])
		return (-1) * n, err
	}

	return parseUInt(d)
}

// parseUInt is similar to the function in strconv, but is tuned for ints appearing in FIX field types.
func parseUInt(d []byte) (n int, err error) {
	if len(d) == 0 {
		err = errors.New("empty bytes")
		return
	}

	for _, dec := range d {
		if dec < ascii0 || dec > ascii9 {
			err = errors.New("invalid format")
			return
		}

		n = n*10 + (int(dec) - ascii0)
	}

	return
}

// FIXInt is a FIX Int Value, implements FieldValue.
type FIXInt int

// Int converts the FIXInt value to int.
func (f FIXInt) Int() int { return int(f) }

func (f *FIXInt) Read(bytes []byte) error {
	i, err := atoi(bytes)
	if err != nil {
		return err
	}
	*f = FIXInt(i)
	return nil
}

func (f FIXInt) Write() []byte {
	return strconv.AppendInt(nil, int64(f), 10)
}
