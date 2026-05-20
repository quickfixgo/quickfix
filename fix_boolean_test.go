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

func TestBooleanWrite(t *testing.T) {
	var tests = []struct {
		val      FIXBoolean
		expected []byte
	}{
		{FIXBoolean(true), []byte("Y")},
		{FIXBoolean(false), []byte("N")},
	}

	for _, test := range tests {
		b := test.val.Write()
		assert.True(t, bytes.Equal(b, test.expected), fmt.Sprintf("got %v; want %v", b, test.expected))
	}
}

func TestFIXBooleanRead(t *testing.T) {
	var tests = []struct {
		bytes       []byte
		expected    bool
		expectError bool
	}{
		{[]byte("Y"), true, false},
		{[]byte("N"), false, false},
		{[]byte("blah"), false, true},
	}

	for _, test := range tests {
		var val FIXBoolean
		err := val.Read(test.bytes)

		assert.Equal(t, test.expectError, err != nil)
		assert.Equal(t, test.expected, val.Bool())
	}
}
