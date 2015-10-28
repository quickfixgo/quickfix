package quickfix

import (
	"bytes"
	"testing"
)

func TestNewStringField(t *testing.T) {
	var tests = []struct {
		tag Tag
		val string
	}{
		{Tag(1), "CWB"},
	}

	for _, test := range tests {
		field := NewStringField(test.tag, test.val)

		if field.Tag() != test.tag {
			t.Errorf("got tag %v; want %v", field.Tag(), test.tag)
		}

		if field.Value != test.val {
			t.Errorf("got val %v; want %v", field.Value, test.val)
		}
	}
}

func TestStringFieldWrite(t *testing.T) {
	var tests = []struct {
		field *StringField
		val   []byte
	}{
		{NewStringField(1, "CWB"), []byte("CWB")},
	}

	for _, test := range tests {
		b := test.field.Write()

		if !bytes.Equal(b, test.val) {
			t.Errorf("got %v; want %v", b, test.val)
		}
	}
}

func TestStringFieldRead(t *testing.T) {
	var tests = []struct {
		bytes       []byte
		value       string
		expectError bool
	}{
		{[]byte("blah"), "blah", false},
	}

	for _, test := range tests {
		field := new(StringField)
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
