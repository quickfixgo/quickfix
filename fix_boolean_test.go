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
