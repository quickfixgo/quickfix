package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&StringFieldTests{})

type StringFieldTests struct{ field *StringField }

func (s *StringFieldTests) SetUpTest(c *C) {
	s.field = NewStringField(1, "CWB")
}

func (s *StringFieldTests) TestTagValue(c *C) {
	c.Check(s.field.Tag(), Equals, tag.Tag(1))
	c.Check(s.field.Value(), Equals, "CWB")
}
