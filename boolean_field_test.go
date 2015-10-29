package quickfix

import (
	"bytes"
	"testing"
)

func TestNewBooleanField(t *testing.T) {
	var tests = []struct {
		tag Tag
		val bool
	}{
		{Tag(1), true},
		{Tag(2), false},
	}

	for _, test := range tests {
		field := NewBooleanField(test.tag, test.val)

		if field.Tag() != test.tag {
			t.Errorf("got tag %v; want %v", field.Tag(), test.tag)
		}

		if field.Value != test.val {
			t.Errorf("got val %v; want %v", field.Value, test.val)
		}
	}
}

func TestBooleanFieldWrite(t *testing.T) {

	var tests = []struct {
		field *BooleanField
		val   []byte
	}{
		{NewBooleanField(1, true), []byte("Y")},
		{NewBooleanField(1, false), []byte("N")},
	}

	for _, test := range tests {
		b := test.field.Write()

		if !bytes.Equal(b, test.val) {
			t.Errorf("got %v; want %v", b, test.val)
		}
	}
}

func TestBooleanFieldRead(t *testing.T) {
	var tests = []struct {
		bytes       []byte
		value       bool
		expectError bool
	}{
		{[]byte("Y"), true, false},
		{[]byte("N"), false, false},
		{[]byte("blah"), false, true},
	}

	for _, test := range tests {
		field := NewBooleanField(Tag(1), false)
		_, err := field.Read([]TagValue{TagValue{Value: test.bytes}})

		if test.expectError && err == nil {
			t.Errorf("Expected error for %v", test.bytes)
		} else if !test.expectError && err != nil {
			t.Errorf("UnExpected '%v'", err)
		}

		if field.Value != test.value {
			t.Errorf("got %v want %v", field.Value, test.value)
		}
	}
}
