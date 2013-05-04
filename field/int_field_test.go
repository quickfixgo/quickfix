package field

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&IntFieldTests{})

type IntFieldTests struct{}

func (s *IntFieldTests) TestNewField(c *C) {
	field := NewIntField(tag.Tag(1), 5)
	c.Check(field.Tag(), Equals, tag.Tag(1))
	c.Check(field.Value, Equals, 5)
}

func (s *IntFieldTests) TestConvertValueToBytes(c *C) {
	field := NewIntField(tag.Tag(1), 5)
	bytes := field.ConvertValueToBytes()
	c.Check(string(bytes), Equals, "5")
}

func (s *IntFieldTests) TestConvertValueFromBytes(c *C) {
	field := new(IntField)
	err := field.ConvertValueFromBytes([]byte("15"))
	c.Check(err, IsNil)
	c.Check(field.Value, Equals, 15)

	err = field.ConvertValueFromBytes([]byte("blah"))
	c.Check(err, NotNil)
}
