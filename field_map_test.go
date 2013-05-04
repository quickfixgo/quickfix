package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/field"
	. "launchpad.net/gocheck"
)

var _ = Suite(&FieldMapTests{})

type FieldMapTests struct{ fieldMap FieldMap }

func (s *FieldMapTests) SetUpTest(c *C) {
	s.fieldMap.init(normalFieldOrder)
}

func (s *FieldMapTests) TestSetAndGet(c *C) {
	field1 := field.NewStringField(1, "hello")
	field2 := field.NewStringField(2, "world")

	s.fieldMap.SetField(field1)
	s.fieldMap.SetField(field2)

	testField := field.NewStringField(1, "")
	err := s.fieldMap.Get(testField)
	c.Check(err, IsNil)
	c.Check(testField.Value, Equals, "hello")

	testField = field.NewStringField(2, "")
	err = s.fieldMap.Get(testField)
	c.Check(err, IsNil)
	c.Check(testField.Value, Equals, "world")

	testField = field.NewStringField(44, "")
	err = s.fieldMap.Get(testField)
	c.Check(err, NotNil)
}

func (s *FieldMapTests) TestLength(c *C) {
	f1 := field.NewStringField(1, "hello")
	f2 := field.NewStringField(2, "world")

	beginString := field.NewStringField(8, "FIX.4.4")
	bodyLength := field.NewIntField(9, 100)
	checkSum := field.NewStringField(10, "100")

	s.fieldMap.SetField(f1)
	s.fieldMap.SetField(f2)
	s.fieldMap.SetField(beginString)
	s.fieldMap.SetField(bodyLength)
	s.fieldMap.SetField(checkSum)

	c.Check(s.fieldMap.length(), Equals, 16,
		Commentf("Length includes all fields but beginString, bodyLength, and checkSum"))
}

func (s *FieldMapTests) TestTotal(c *C) {
	f1 := field.NewStringField(1, "hello")
	f2 := field.NewStringField(2, "world")

	beginString := field.NewStringField(8, "FIX.4.4")
	bodyLength := field.NewIntField(9, 100)
	checkSum := field.NewStringField(10, "100")

	s.fieldMap.SetField(f1)
	s.fieldMap.SetField(f2)
	s.fieldMap.SetField(beginString)
	s.fieldMap.SetField(bodyLength)
	s.fieldMap.SetField(checkSum)

	c.Check(s.fieldMap.total(), Equals, 2116,
		Commentf("Total includes all fields but checkSum"))
}
