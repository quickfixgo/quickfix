package quickfix

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

func BenchmarkParser_ReadMessage(b *testing.B) {
	stream := "8=FIXT.1.19=11135=D34=449=TW52=20140511-23:10:3456=ISLD11=ID21=340=154=155=INTC60=20140511-23:10:3410=2348=FIXT.1.19=9535=D34=549=TW52=20140511-23:10:3456=ISLD11=ID21=340=154=155=INTC60=20140511-23:10:3410=198"

	for i := 0; i < b.N; i++ {
		reader := strings.NewReader(stream)
		parser := newParser(reader)
		_, _ = parser.ReadMessage()
	}
}

type ParserSuite struct {
	suite.Suite
	*parser
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

func (s *ParserSuite) SetupTest() {
	s.parser = new(parser)
}

func (s *ParserSuite) TestInvalidNilLength() {
	stream := "8=\x019=\x01"
	s.reader = strings.NewReader(stream)
	_, err := s.ReadMessage()
	s.NotNil(err)
}

func (s *ParserSuite) TestOverflowLength() {
	stream := "8=\x019=9300000000000000000\x01"
	s.reader = strings.NewReader(stream)
	_, err := s.ReadMessage()
	s.NotNil(err)
}

func (s *ParserSuite) TestJumpLength() {
	stream := "8=FIXT.1.19=11135=D34=449=TW52=20140511-23:10:3456=ISLD11=ID21=340=154=155=INTC60=20140511-23:10:3410=2348=FIXT.1.19=9535=D34=549=TW52=20140511-23:10:3456=ISLD11=ID21=340=154=155=INTC60=20140511-23:10:3410=198"

	s.reader = strings.NewReader(stream)
	index, err := s.parser.jumpLength()
	s.Nil(err)

	expectedIndex := 111 + 17 - 1
	s.Equal(expectedIndex, index)
}

func (s *ParserSuite) TestBadLength() {
	stream := "8=FIXT.1.19=11135=D34=449=TW52=20140511-23:10:3456=ISLD11=ID21=340=154=155=INTC60=20140511-23:10:3410=2348=FIXT.1.19=9535=D34=549=TW52=20140511-23:10:3456=ISLD11=ID21=340=154=155=INTC60=20140511-23:10:3410=198"

	s.reader = strings.NewReader(stream)
	bytes, _ := s.parser.ReadMessage()

	s.Equal(stream, bytes.String())
}

func (s *ParserSuite) TestFindStart() {
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
		s.SetupTest()

		s.reader = strings.NewReader(tc.stream)

		start, err := s.parser.findStart()
		if tc.expectError {
			s.NotNil(err)
			continue
		}

		s.Nil(err)
		s.Equal(tc.expectedStart, start)
	}
}

func (s *ParserSuite) TestReadEOF() {
	var testCases = []struct {
		stream string
	}{
		{""},
		{"hello8=FIX.4.0"},
	}

	for _, tc := range testCases {
		s.SetupTest()

		s.reader = strings.NewReader(tc.stream)
		bytes, err := s.parser.ReadMessage()
		s.Nil(bytes)

		s.NotNil(err)
		s.Equal("EOF", err.Error())
	}
}

func (s *ParserSuite) TestReadMessage() {
	stream := "hello8=FIX.4.09=5blah10=1038=FIX.4.09=4foo10=103"
	s.reader = strings.NewReader(stream)

	var testCases = []struct {
		expectedBytes     string
		expectedBufferLen int
		expectedBufferCap int
	}{
		{expectedBytes: "8=FIX.4.09=5blah10=103", expectedBufferCap: defaultBufSize - 31, expectedBufferLen: len(stream) - 31},
		{expectedBytes: "8=FIX.4.09=4foo10=103", expectedBufferCap: defaultBufSize - 31 - 25, expectedBufferLen: 0},
	}

	for _, tc := range testCases {
		msg, err := s.parser.ReadMessage()
		s.Nil(err)

		s.Equal(tc.expectedBytes, msg.String())
		s.Equal(tc.expectedBufferCap, cap(s.parser.buffer))
		s.Equal(tc.expectedBufferLen, len(s.parser.buffer))
	}
}

func (s *ParserSuite) TestReadMessageGrowBuffer() {
	stream := "hello8=FIX.4.09=5blah10=1038=FIX.4.09=4foo10=103"

	var testCases = []struct {
		initialBufCap        int
		expectedBytes        string
		expectedBufferLen    int
		expectedBufferCap    int
		expectedBigBufferLen int
	}{
		{initialBufCap: 0, expectedBufferCap: defaultBufSize - 31, expectedBufferLen: len(stream) - 31, expectedBigBufferLen: defaultBufSize},
		{initialBufCap: 4, expectedBufferCap: 6, expectedBufferLen: 6, expectedBigBufferLen: 32},
		{initialBufCap: 8, expectedBufferCap: 6, expectedBufferLen: 6, expectedBigBufferLen: 32},
		{initialBufCap: 14, expectedBufferCap: 10, expectedBufferLen: 10, expectedBigBufferLen: 36},
		{initialBufCap: 16, expectedBufferCap: 18, expectedBufferLen: 18, expectedBigBufferLen: 44},
		{initialBufCap: 23, expectedBufferCap: 10, expectedBufferLen: 10, expectedBigBufferLen: 36},
		{initialBufCap: 30, expectedBufferCap: 24, expectedBufferLen: 24, expectedBigBufferLen: 50},
		{initialBufCap: 31, expectedBufferCap: 0, expectedBufferLen: 0, expectedBigBufferLen: 31},
		{initialBufCap: 40, expectedBufferCap: 9, expectedBufferLen: 9, expectedBigBufferLen: 40},
		{initialBufCap: 60, expectedBufferCap: 29, expectedBufferLen: 25, expectedBigBufferLen: 60},
		{initialBufCap: 80, expectedBufferCap: 49, expectedBufferLen: 25, expectedBigBufferLen: 80},
	}

	for _, tc := range testCases {
		s.SetupTest()

		s.reader = strings.NewReader(stream)
		s.parser.bigBuffer = make([]byte, tc.initialBufCap)
		s.parser.buffer = s.parser.bigBuffer[0:0]
		msg, err := s.parser.ReadMessage()

		s.Nil(err)
		s.Equal("8=FIX.4.09=5blah10=103", msg.String())
		s.Equal(tc.expectedBigBufferLen, len(s.parser.bigBuffer))
		s.Equal(tc.expectedBufferCap, cap(s.parser.buffer))
		s.Equal(tc.expectedBufferLen, len(s.parser.buffer))
	}
}
