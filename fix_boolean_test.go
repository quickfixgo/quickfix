package quickfix

import (
	"bytes"
	"testing"
)

func TestBooleanFieldWrite(t *testing.T) {

	var tests = []struct {
		val      FIXBoolean
		expected []byte
	}{
		{FIXBoolean(true), []byte("Y")},
		{FIXBoolean(false), []byte("N")},
	}

	for _, test := range tests {
		b := test.val.Write()

		if !bytes.Equal(b, test.expected) {
			t.Errorf("got %v; want %v", b, test.expected)
		}
	}
}

func TestFIXBooleanFieldRead(t *testing.T) {
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

		if test.expectError && err == nil {
			t.Errorf("Expected error for %v", test.bytes)
		} else if !test.expectError && err != nil {
			t.Errorf("UnExpected '%v'", err)
		}

		if val != FIXBoolean(test.expected) {
			t.Errorf("got %v want %v", val, test.expected)
		}
	}
}
