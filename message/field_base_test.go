package message

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&FieldBaseTests{})

type FieldBaseTests struct{ field FieldBase }

func (s *FieldBaseTests) TestParseField(c *C) {
	stringField := "8=FIX.4.0"

	parsed_field, err := ParseField([]byte(stringField))
	c.Check(parsed_field, NotNil)
	c.Check(err, IsNil)
	c.Check(parsed_field.Tag(), Equals, tag.Tag(8))
	c.Check(parsed_field.Value(), Equals, "FIX.4.0")

	//no equals?!
	stringField = "uhoh"
	parsed_field, err = ParseField([]byte(stringField))
	c.Check(err, NotNil)

	//tag not valid?
	stringField = "not_an_int=uhoh"
	parsed_field, err = ParseField([]byte(stringField))
	c.Check(err, NotNil)
}

func (s *FieldBaseTests) SetUpTest(c *C) {
	s.field.init(1, "hello")
}

func (s *FieldBaseTests) TestTagValue(c *C) {
	c.Check(s.field.Tag(), Equals, tag.Tag(1))
	c.Check(s.field.Value(), Equals, "hello")
}

func (s *FieldBaseTests) TestString(c *C) {
	c.Check(s.field.String(), Equals, "1=hello")
}

func (s *FieldBaseTests) TestLength(c *C) {
	expected := len(s.field.String())
	c.Check(s.field.Length(), Equals, expected)
}

func (s *FieldBaseTests) TestTotal(c *C) {
	c.Check(s.field.Total(), Equals, 643, Commentf("Total is the summation of the ascii byte values of the field string"))
}
