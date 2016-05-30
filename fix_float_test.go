package quickfix

import (
	"bytes"
	"testing"
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

		if !bytes.Equal(b, test.val) {
			t.Errorf("got %v; want %v", b, test.val)
		}
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
		if err := field.Read(test.bytes); err != nil {
			if !test.expectError {
				t.Errorf("UnExpected '%v'", err)
			}
		} else if test.expectError {
			t.Errorf("Expected error for %v", test.bytes)
		} else if float64(field) != test.value {
			t.Errorf("got %v want %v", field, test.value)
		}
	}
}

func BenchmarkFloatRead(b *testing.B) {
	val := []byte("15.1234")
	for i := 0; i < b.N; i++ {
		var field FIXFloat
		field.Read(val)
	}
}
