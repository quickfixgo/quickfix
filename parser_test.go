package quickfixgo

import (
	. "launchpad.net/gocheck"
	"strings"
)

var _ = Suite(&MessageParserTests{})

type MessageParserTests struct{}

func (s *MessageParserTests) TestReadEOF(c *C) {
	reader := strings.NewReader("")

	parser := newParser(reader)
	bytes, err := parser.readMessage()
	c.Check(len(bytes), Equals, 0)
	c.Check(err.Error(), Equals, "EOF")

	reader = strings.NewReader("hello8=FIX.4.0")
	parser = newParser(reader)
	bytes, err = parser.readMessage()
	c.Check(len(bytes), Equals, 0)
	c.Check(err.Error(), Equals, "EOF")
}

func (s *MessageParserTests) TestReadSingleMessage(c *C) {
	reader := strings.NewReader("hello8=FIX.4.09=5blah10=1038=FIX.4.2...")
	parser := newParser(reader)
	bytes, _ := parser.readMessage()
	c.Check(string(bytes), Equals, "8=FIX.4.09=5blah10=103")
}

func (s *MessageParserTests) TestReadMultiMessage(c *C) {
	reader := strings.NewReader("hello8=FIX.4.09=5blah10=1038=FIX.4.09=4foo10=103")
	parser := newParser(reader)
	msg1, _ := parser.readMessage()
	c.Check(string(msg1), Equals, "8=FIX.4.09=5blah10=103")

	msg2, _ := parser.readMessage()
	c.Check(string(msg2), Equals, "8=FIX.4.09=4foo10=103")
}

func (s *MessageParserTests) TestReadMessageBufferExactSize(c *C) {
	fixStr := "hello8=FIX.4.09=5blah10=1038=FIX.4.2..."
	reader := strings.NewReader(fixStr)
	parser := newParserSize(reader, len(fixStr))

	bytes, _ := parser.readMessage()
	c.Check(string(bytes), Equals, "8=FIX.4.09=5blah10=103")
}

func (s *MessageParserTests) TestReadMessageBufferExpandMissingBegin(c *C) {
	fixStr := "hello8=FIX.4.09=5blah10=1038=FIX.4.2..."
	reader := strings.NewReader(fixStr)
	parser := newParserSize(reader, 2)

	bytes, _ := parser.readMessage()
	c.Check(string(bytes), Equals, "8=FIX.4.09=5blah10=103")
}

func (s *MessageParserTests) TestReadMessageBufferExpandMissingLengthAfterTag(c *C) {
	fixStr := "8=FIX.4.09=5blah10=1038=FIX.4.2..."
	reader := strings.NewReader(fixStr)
	parser := newParserSize(reader, 11)

	bytes, _ := parser.readMessage()
	c.Check(string(bytes), Equals, "8=FIX.4.09=5blah10=103")
}

func (s *MessageParserTests) TestReadMessageBufferExpandMissingLengthAtEquals(c *C) {
	fixStr := "8=FIX.4.09=5blah10=1038=FIX.4.2..."
	reader := strings.NewReader(fixStr)
	parser := newParserSize(reader, 12)

	bytes, _ := parser.readMessage()
	c.Check(string(bytes), Equals, "8=FIX.4.09=5blah10=103")
}

func (s *MessageParserTests) TestReadMessageBigLength(c *C) {
	reader := strings.NewReader("8=FIX.4.09=10blah3jvde10=1038=FIX.4.2...")
	parser := newParserSize(reader, 20)

	bytes, _ := parser.readMessage()
	c.Check(string(bytes), Equals, "8=FIX.4.09=10blah3jvde10=103")
}
