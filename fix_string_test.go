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
