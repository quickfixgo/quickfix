package message

import (
	"github.com/quickfixgo/quickfix/fix"
	. "launchpad.net/gocheck"
)

var _ = Suite(&BooleanFieldTests{})

type BooleanFieldTests struct{}

func (s *BooleanFieldTests) TestNewField(c *C) {
	field := NewBooleanField(fix.Tag(1), true)
	c.Check(field.Tag(), Equals, fix.Tag(1))
	c.Check(field.Value, Equals, true)

	field = NewBooleanField(fix.Tag(2), false)
	c.Check(field.Tag(), Equals, fix.Tag(2))
	c.Check(field.Value, Equals, false)
}

func (s *BooleanFieldTests) TestWrite(c *C) {
	field := NewBooleanField(fix.Tag(1), true)
	bytes := field.Write()
	c.Check(string(bytes), Equals, "Y")

	field = NewBooleanField(fix.Tag(1), false)
	bytes = field.Write()

	c.Check(string(bytes), Equals, "N")
}

func (s *BooleanFieldTests) TestRead(c *C) {
	field := NewBooleanField(fix.Tag(1), false)
	err := field.Read([]byte("Y"))
	c.Check(err, IsNil)
	c.Check(field.Value, Equals, true)

	err = field.Read([]byte("N"))
	c.Check(err, IsNil)
	c.Check(field.Value, Equals, false)

	err = field.Read([]byte("blah"))
	c.Check(err, NotNil)
}
