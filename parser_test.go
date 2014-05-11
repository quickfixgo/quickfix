package quickfix

import (
	"strings"
	"testing"
)

func TestParser_findStart(t *testing.T) {
	var testCases = []struct {
		stream        string
		expectError   bool
		expectedStart int
	}{
		{stream: "", expectError: true},
		{stream: "nostarthere", expectError: true},
		{stream: "hello8=FIX.4.0", expectedStart: 5},
	}
	for _, tc := range testCases {
		reader := strings.NewReader(tc.stream)
		parser := newParser(reader)

		start, err := parser.findStart()

		if err != nil {
			if !tc.expectError {
				t.Error("unxpected error", err)
			}
			continue
		}

		if err == nil && tc.expectError {
			t.Error("expected error")
		}

		if start != tc.expectedStart {
			t.Errorf("Expected start to be %v, but was %v", tc.expectedStart, start)
		}
	}

}

func TestParser_ReadEOF(t *testing.T) {

	var testCases = []struct {
		stream string
	}{
		{""},
		{"hello8=FIX.4.0"},
	}

	for _, tc := range testCases {
		reader := strings.NewReader(tc.stream)
		parser := newParser(reader)
		bytes, err := parser.ReadMessage()
		if len(bytes) != 0 {
			t.Error("Length of bytes should be 0")
		}

		if err == nil || err.Error() != "EOF" {
			t.Error("Expected EOF")
		}
	}
}

func TestParser_ReadMessage(t *testing.T) {
	stream := "hello8=FIX.4.09=5blah10=1038=FIX.4.09=4foo10=103"
	reader := strings.NewReader(stream)
	parser := newParser(reader)

	var testCases = []struct {
		expectedBytes     string
		expectedBufferLen int
		expectedBufferCap int
	}{
		{expectedBytes: "8=FIX.4.09=5blah10=103", expectedBufferCap: defaultBufSize - 31, expectedBufferLen: len(stream) - 31},
		{expectedBytes: "8=FIX.4.09=4foo10=103", expectedBufferCap: defaultBufSize - 31 - 25, expectedBufferLen: 0},
	}

	for _, tc := range testCases {
		msg, err := parser.ReadMessage()

		if err != nil {
			t.Error("Unexpected error", err)
		}

		if tc.expectedBytes != string(msg) {
			t.Errorf("Expected %v got %v", tc.expectedBytes, string(msg))
		}

		if cap(parser.buffer) != tc.expectedBufferCap {
			t.Errorf("Expected capacity %v got %v", tc.expectedBufferCap, cap(parser.buffer))
		}

		if len(parser.buffer) != tc.expectedBufferLen {
			t.Errorf("Expected len %v got %v", tc.expectedBufferLen, len(parser.buffer))
		}

	}
}

func TestParser_ReadMessageGrowBuffer(t *testing.T) {
	stream := "hello8=FIX.4.09=5blah10=1038=FIX.4.09=4foo10=103"

	var testCases = []struct {
		initialBufCap     int
		expectedBytes     string
		expectedBufferLen int
		expectedBufferCap int
	}{
		{initialBufCap: 0, expectedBufferCap: defaultBufSize - 31, expectedBufferLen: len(stream) - 31},
		{initialBufCap: 4, expectedBufferCap: defaultBufSize - 27, expectedBufferLen: len(stream) - 31},
		{initialBufCap: 8, expectedBufferCap: defaultBufSize - 23, expectedBufferLen: len(stream) - 31},
		{initialBufCap: 14, expectedBufferCap: defaultBufSize - 17, expectedBufferLen: len(stream) - 31},
		{initialBufCap: 16, expectedBufferCap: defaultBufSize - 15, expectedBufferLen: len(stream) - 31},
		{initialBufCap: 23, expectedBufferCap: defaultBufSize - 8, expectedBufferLen: len(stream) - 31},
		{initialBufCap: 30, expectedBufferCap: defaultBufSize - 1, expectedBufferLen: len(stream) - 31},
		{initialBufCap: 31, expectedBufferCap: 0, expectedBufferLen: 0},
		{initialBufCap: 40, expectedBufferCap: 9, expectedBufferLen: 9},
		{initialBufCap: 60, expectedBufferCap: 29, expectedBufferLen: 25},
		{initialBufCap: 80, expectedBufferCap: 49, expectedBufferLen: 25},
	}

	for _, tc := range testCases {
		parser := newParser(strings.NewReader(stream))
		parser.buffer = make([]byte, 0, tc.initialBufCap)
		msg, err := parser.ReadMessage()

		if err != nil {
			t.Fatal("unexpected err: ", err)
		}

		if string(msg) != "8=FIX.4.09=5blah10=103" {
			t.Error("Didn't get expected message, got ", string(msg))
		}

		if cap(parser.buffer) != tc.expectedBufferCap {
			t.Errorf("expected cap %v, got %v", tc.expectedBufferCap, cap(parser.buffer))
		}

		if len(parser.buffer) != tc.expectedBufferLen {
			t.Errorf("expected len %v, got %v", tc.expectedBufferLen, len(parser.buffer))
		}
	}
}
