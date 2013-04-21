package basic

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&IntFieldTests{})

type IntFieldTests struct{ field *IntField }

func (s *IntFieldTests) SetUpTest(c *C) {
	s.field = NewIntField(tag.Tag(1), 5)
}

func (s *IntFieldTests) TestNewField(c *C) {
	c.Check(s.field.Tag(), Equals, tag.Tag(1))
	c.Check(s.field.Value(), Equals, "5")
	c.Check(s.field.IntValue(), Equals, 5)
}

func (s *IntFieldTests) TestToIntField(c *C) {
	f := FieldBase{}
	f.init(tag.Tag(1), "5")
	intField, err := ToIntField(f)

	c.Check(err, IsNil)
	c.Check(intField.Tag(), Equals, tag.Tag(1))
	c.Check(intField.Value(), Equals, "5")
	c.Check(intField.IntValue(), Equals, 5)
}

func (s *IntFieldTests) TestToIntFieldInvalid(c *C) {
	f := FieldBase{}
	f.init(tag.Tag(1), "blah")
	_, err := ToIntField(f)

	c.Check(err, Not(IsNil))
}
