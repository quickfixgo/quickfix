package quickfixgo

import (
	. "launchpad.net/gocheck"
)

var _ = Suite(&FieldMapTests{})

type FieldMapTests struct{ fieldMap FieldMap }

func (s *FieldMapTests) SetUpTest(c *C) {
	s.fieldMap.init(normalFieldOrder)
}

func (s *FieldMapTests) TestSetAndGet(c *C) {
	field1 := NewStringField(1, "hello")
	field2 := NewStringField(2, "world")

	s.fieldMap.SetField(field1)
	s.fieldMap.SetField(field2)

	retField, ok := s.fieldMap.Field(field1.Tag())
	c.Check(ok, Equals, true)
	c.Check(retField, Equals, field1)

	retField, ok = s.fieldMap.Field(field2.Tag())
	c.Check(ok, Equals, true)
	c.Check(retField, Equals, field2)

	retField, ok = s.fieldMap.Field(44)
	c.Check(ok, Equals, false)
}

func (s *FieldMapTests) TestLength(c *C) {
	f1 := NewStringField(1, "hello")
	f2 := NewStringField(2, "world")

	beginString := NewStringField(8, "FIX.4.4")
	bodyLength := NewIntField(9, 100)
	checkSum := NewStringField(10, "100")

	s.fieldMap.SetField(f1)
	s.fieldMap.SetField(f2)
	s.fieldMap.SetField(beginString)
	s.fieldMap.SetField(bodyLength)
	s.fieldMap.SetField(checkSum)

	c.Check(s.fieldMap.Length(), Equals, f1.Length()+f2.Length(),
		Commentf("Length includes all fields but beginString, bodyLength, and checkSum"))
}

func (s *FieldMapTests) TestTotal(c *C) {
	f1 := NewStringField(1, "hello")
	f2 := NewStringField(2, "world")

	beginString := NewStringField(8, "FIX.4.4")
	bodyLength := NewIntField(9, 100)
	checkSum := NewStringField(10, "100")

	s.fieldMap.SetField(f1)
	s.fieldMap.SetField(f2)
	s.fieldMap.SetField(beginString)
	s.fieldMap.SetField(bodyLength)
	s.fieldMap.SetField(checkSum)

	c.Check(s.fieldMap.Total(), Equals, f1.Total()+f2.Total()+beginString.Total()+bodyLength.Total(),
		Commentf("Total includes all fields but checkSum"))
}
