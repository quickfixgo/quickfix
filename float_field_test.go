package quickfix

import (
	"bytes"
	"testing"
)

func TestNewFloatField(t *testing.T) {
	var tests = []struct {
		tag Tag
		val float64
	}{
		{Tag(1), 5.0},
	}

	for _, test := range tests {
		field := NewFloatField(test.tag, test.val)

		if field.Tag() != test.tag {
			t.Errorf("got tag %v; want %v", field.Tag(), test.tag)
		}

		if field.Value != test.val {
			t.Errorf("got val %v; want %v", field.Value, test.val)
		}
	}
}

func TestFloatFieldWrite(t *testing.T) {
	var tests = []struct {
		field *FloatField
		val   []byte
	}{
		{NewFloatField(1, 5.0), []byte("5")},
	}

	for _, test := range tests {
		b := test.field.Write()

		if !bytes.Equal(b, test.val) {
			t.Errorf("got %v; want %v", b, test.val)
		}
	}
}

func TestFloatFieldRead(t *testing.T) {
	var tests = []struct {
		bytes       []byte
		value       float64
		expectError bool
	}{
		{[]byte("15"), 15.0, false},
		{[]byte("blah"), 0.0, true},
		{[]byte("+200.00"), 0.0, true},
	}

	for _, test := range tests {
		field := new(FloatField)
		if _, err := field.Read([]TagValue{TagValue{Value: test.bytes}}); err != nil {
			if !test.expectError {
				t.Errorf("UnExpected '%v'", err)
			}
		} else if test.expectError {
			t.Errorf("Expected error for %v", test.bytes)
		} else if field.Value != test.value {
			t.Errorf("got %v want %v", field.Value, test.value)
		}
	}
}
