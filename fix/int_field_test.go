package fix

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&IntFieldTests{})

type IntFieldTests struct{}

func (s *IntFieldTests) TestNewField(c *C) {
	field := NewIntField(Tag(1), 5)
	c.Check(field.Tag(), Equals, Tag(1))
	c.Check(field.Value, Equals, 5)
}

func (s *IntFieldTests) TestWrite(c *C) {
	field := NewIntField(Tag(1), 5)
	bytes := field.Write()
	c.Check(string(bytes), Equals, "5")
}

func (s *IntFieldTests) TestRead(c *C) {
	field := new(IntField)
	err := field.Read([]byte("15"))
	c.Check(err, IsNil)
	c.Check(field.Value, Equals, 15)

	err = field.Read([]byte("blah"))
	c.Check(err, NotNil)
}
