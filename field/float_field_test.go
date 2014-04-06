package field

import (
	"github.com/quickfixgo/quickfix/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&FloatFieldTests{})

type FloatFieldTests struct{}

func (s *FloatFieldTests) TestNewField(c *C) {
	field := NewFloatField(tag.Tag(1), 5.0)
	c.Check(field.Tag(), Equals, tag.Tag(1))
	c.Check(field.Value, Equals, 5.0)
}

func (s *FloatFieldTests) TestWrite(c *C) {
	field := NewFloatField(tag.Tag(1), 5.0)
	bytes := field.Write()
	c.Check(string(bytes), Equals, "5")
}

func (s *FloatFieldTests) TestRead(c *C) {
	field := new(FloatField)
	err := field.Read([]byte("15"))
	c.Check(err, IsNil)
	c.Check(field.Value, Equals, 15.0)

	err = field.Read([]byte("blah"))
	c.Check(err, NotNil)

	err = field.Read([]byte("+200.00"))
	c.Check(err, NotNil)
}
