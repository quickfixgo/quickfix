package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&StringFieldTests{})

type StringFieldTests struct{}

func (s *StringFieldTests) SetUpTest(c *C) {
}

func (s *StringFieldTests) TestNewField(c *C) {
	field := NewStringField(1, "CWB")
	c.Check(field.Tag(), Equals, tag.Tag(1))
	c.Check(field.Value, Equals, "CWB")
}

func (s *StringFieldTests) TestConvertValueToBytes(c *C) {
	field := NewStringField(1, "CWB")
	bytes := field.ConvertValueToBytes()
	c.Check(string(bytes), Equals, "CWB")
}

func (s *StringFieldTests) TestConvertValueFromBytes(c *C) {
	field := new(StringField)
	err := field.ConvertValueFromBytes([]byte("blah"))

	c.Check(err, IsNil)
	c.Check(field.Value, Equals, "blah")
}
