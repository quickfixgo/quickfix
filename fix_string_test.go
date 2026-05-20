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
	"testing"
)

func TestFIXStringWrite(t *testing.T) {
	var tests = []struct {
		field FIXString
		val   []byte
	}{
		{"CWB", []byte("CWB")},
	}

	for _, test := range tests {
		b := test.field.Write()

		if !bytes.Equal(b, test.val) {
			t.Errorf("got %v; want %v", b, test.val)
		}
	}
}

func TestFIXStringRead(t *testing.T) {
	var tests = []struct {
		bytes       []byte
		value       string
		expectError bool
	}{
		{[]byte("blah"), "blah", false},
		{nil, "", false},
		{[]byte(""), "", false},
	}

	for _, test := range tests {
		var field FIXString
		err := field.Read(test.bytes)

		if test.expectError && err == nil {
			t.Errorf("Expected error for %v", test.bytes)
		} else if !test.expectError && err != nil {
			t.Errorf("UnExpected '%v'", err)
		}

		if string(field) != test.value {
			t.Errorf("got %v want %v", field, test.value)
		}
	}
}
