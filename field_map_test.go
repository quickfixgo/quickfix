package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	. "gopkg.in/check.v1"
)

var _ = Suite(&FieldMapTests{})

type FieldMapTests struct{ fieldMap FieldMap }

func (s *FieldMapTests) SetUpTest(c *C) {
	s.fieldMap.init(normalFieldOrder)
}

func (s *FieldMapTests) TestSetAndGet(c *C) {
	field1 := fix.NewStringField(1, "hello")
	field2 := fix.NewStringField(2, "world")

	s.fieldMap.Set(field1)
	s.fieldMap.Set(field2)

	testField := fix.NewStringField(1, "")
	err := s.fieldMap.Get(testField)
	c.Check(err, IsNil)
	c.Check(testField.Value, Equals, "hello")

	testField = fix.NewStringField(2, "")
	err = s.fieldMap.Get(testField)
	c.Check(err, IsNil)
	c.Check(testField.Value, Equals, "world")

	testField = fix.NewStringField(44, "")
	err = s.fieldMap.Get(testField)
	c.Check(err, NotNil)
}

func (s *FieldMapTests) TestLength(c *C) {
	f1 := fix.NewStringField(1, "hello")
	f2 := fix.NewStringField(2, "world")

	beginString := fix.NewStringField(8, "FIX.4.4")
	bodyLength := fix.NewIntField(9, 100)
	checkSum := fix.NewStringField(10, "100")

	s.fieldMap.Set(f1)
	s.fieldMap.Set(f2)
	s.fieldMap.Set(beginString)
	s.fieldMap.Set(bodyLength)
	s.fieldMap.Set(checkSum)

	c.Check(s.fieldMap.length(), Equals, 16,
		Commentf("Length includes all fields but beginString, bodyLength, and checkSum"))
}

func (s *FieldMapTests) TestTotal(c *C) {
	f1 := fix.NewStringField(1, "hello")
	f2 := fix.NewStringField(2, "world")

	beginString := fix.NewStringField(8, "FIX.4.4")
	bodyLength := fix.NewIntField(9, 100)
	checkSum := fix.NewStringField(10, "100")

	s.fieldMap.Set(f1)
	s.fieldMap.Set(f2)
	s.fieldMap.Set(beginString)
	s.fieldMap.Set(bodyLength)
	s.fieldMap.Set(checkSum)

	c.Check(s.fieldMap.total(), Equals, 2116,
		Commentf("Total includes all fields but checkSum"))
}
