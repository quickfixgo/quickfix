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
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatWrite(t *testing.T) {
	var tests = []struct {
		field FIXFloat
		val   []byte
	}{
		{FIXFloat(5.0), []byte("5")},
	}

	for _, test := range tests {
		b := test.field.Write()

		assert.True(t, bytes.Equal(b, test.val), fmt.Sprintf("got %v; want %v", b, test.val))
	}
}

func TestFloatRead(t *testing.T) {
	var tests = []struct {
		bytes       []byte
		value       float64
		expectError bool
	}{
		{[]byte("15"), 15.0, false},
		{[]byte("99.9"), 99.9, false},
		{[]byte("0.00"), 0.0, false},
		{[]byte("-99.9"), -99.9, false},
		{[]byte("-99.9.9"), 0.0, true},
		{[]byte("blah"), 0.0, true},
		{[]byte("1.a1"), 0.0, true},
		{[]byte("+200.00"), 0.0, true},
	}

	for _, test := range tests {
		var field FIXFloat
		err := field.Read(test.bytes)
		assert.Equal(t, test.expectError, err != nil)
		assert.Equal(t, test.value, field.Float64())
	}
}

func BenchmarkFloatRead(b *testing.B) {
	val := []byte("15.1234")
	for i := 0; i < b.N; i++ {
		var field FIXFloat
		_ = field.Read(val)
	}
}
