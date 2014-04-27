package message

import (
	"github.com/quickfixgo/quickfix/fix"
	. "gopkg.in/check.v1"
)

var _ = Suite(&StringFieldTests{})

type StringFieldTests struct{}

func (s *StringFieldTests) SetUpTest(c *C) {
}

func (s *StringFieldTests) TestNewField(c *C) {
	field := NewStringField(1, "CWB")
	c.Check(field.Tag(), Equals, fix.Tag(1))
	c.Check(field.Value, Equals, "CWB")
}

func (s *StringFieldTests) TestWrite(c *C) {
	field := NewStringField(1, "CWB")
	bytes := field.Write()
	c.Check(string(bytes), Equals, "CWB")
}

func (s *StringFieldTests) TestRead(c *C) {
	field := new(StringField)
	err := field.Read([]byte("blah"))

	c.Check(err, IsNil)
	c.Check(field.Value, Equals, "blah")
}
