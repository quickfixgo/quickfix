package basic

import (
	"github.com/cbusbey/quickfixgo/message"
	. "launchpad.net/gocheck"
)

var _ = Suite(&StringFieldTests{})

type StringFieldTests struct{ field *StringField }

func (s *StringFieldTests) SetUpTest(c *C) {
	s.field = NewStringField(1, "CWB")
}

func (s *StringFieldTests) TestTagValue(c *C) {
	c.Check(s.field.Tag(), Equals, message.Tag(1))
	c.Check(s.field.Value(), Equals, "CWB")
}
