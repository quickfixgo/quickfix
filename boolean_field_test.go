package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&BooleanFieldTests{})

type BooleanFieldTests struct{}

func (s *BooleanFieldTests) TestNewFieldTrue(c *C) {
	field := NewBooleanField(tag.Tag(1), true)
	c.Check(field.Tag(), Equals, tag.Tag(1))
	c.Check(field.Value(), Equals, "Y")
	c.Check(field.BooleanValue(), Equals, true)
}

func (s *BooleanFieldTests) TestNewFieldFalse(c *C) {
	field := NewBooleanField(tag.Tag(2), false)
	c.Check(field.Tag(), Equals, tag.Tag(2))
	c.Check(field.Value(), Equals, "N")
	c.Check(field.BooleanValue(), Equals, false)
}

func (s *BooleanFieldTests) TestToBooleanFieldTrue(c *C) {
	f := FieldBase{}
	f.init(tag.Tag(1), "Y")
	booleanField, err := ToBooleanField(f)

	c.Check(err, IsNil)
	c.Check(booleanField.Tag(), Equals, tag.Tag(1))
	c.Check(booleanField.Value(), Equals, "Y")
	c.Check(booleanField.BooleanValue(), Equals, true)
}

func (s *BooleanFieldTests) TestToBooleanFieldFalse(c *C) {
	f := FieldBase{}
	f.init(tag.Tag(1), "N")
	booleanField, err := ToBooleanField(f)

	c.Check(err, IsNil)
	c.Check(booleanField.Tag(), Equals, tag.Tag(1))
	c.Check(booleanField.Value(), Equals, "N")
	c.Check(booleanField.BooleanValue(), Equals, false)
}

func (s *BooleanFieldTests) TestToBooleanFieldInvalid(c *C) {
	f := FieldBase{}
	f.init(tag.Tag(1), "blah")
	_, err := ToBooleanField(f)

	c.Check(err, Not(IsNil))
}
