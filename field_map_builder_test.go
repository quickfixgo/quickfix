package quickfixgo

import (
	. "launchpad.net/gocheck"
)

var _ = Suite(&FieldMapBuilderTests{})

type FieldMapBuilderTests struct{ fieldMap FieldMapBuilder }

func (s *FieldMapBuilderTests) SetUpTest(c *C) {
	s.fieldMap.init(normalFieldOrder)
}

func (s *FieldMapBuilderTests) TestSetAndGet(c *C) {
	field1 := NewStringField(1, "hello")
	field2 := NewStringField(2, "world")

	s.fieldMap.SetField(field1)
	s.fieldMap.SetField(field2)

	testField := NewStringField(1, "")
	err := s.fieldMap.Get(testField)
	c.Check(err, IsNil)
	c.Check(testField.Value, Equals, "hello")

	testField = NewStringField(2, "")
	err = s.fieldMap.Get(testField)
	c.Check(err, IsNil)
	c.Check(testField.Value, Equals, "world")

	testField = NewStringField(44, "")
	err = s.fieldMap.Get(testField)
	c.Check(err, NotNil)
}

func (s *FieldMapBuilderTests) TestLength(c *C) {
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

	c.Check(s.fieldMap.length(), Equals, 16,
		Commentf("Length includes all fields but beginString, bodyLength, and checkSum"))
}

func (s *FieldMapBuilderTests) TestTotal(c *C) {
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

	c.Check(s.fieldMap.total(), Equals, 2116,
		Commentf("Total includes all fields but checkSum"))
}
