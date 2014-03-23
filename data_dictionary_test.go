package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/field"
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&DataDictionaryTests{})

type DataDictionaryTests struct{}

func (s *DataDictionaryTests) TestNewDataDictionary(c *C) {
	dict, err := NewDataDictionary("../spec/bogus.xml")
	c.Check(err, NotNil)

	dict, err = NewDataDictionary("spec/FIX40.xml")
	c.Check(err, IsNil)
	c.Check(dict, NotNil)
}

func (s *DataDictionaryTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := NewMessage()
	msg.Header.FieldMap.SetField(field.NewStringField(9999, "hello"))
	err := dict.validate(*msg)
	c.Check(err, NotNil)
	c.Check(err.(InvalidTagNumberError).Tag, Equals, tag.Tag(9999))

	msg = NewMessage()
	msg.Trailer.FieldMap.SetField(field.NewStringField(9999, "hello"))
	err = dict.validate(*msg)
	c.Check(err, NotNil)
	c.Check(err.(InvalidTagNumberError).Tag, Equals, tag.Tag(9999))

	msg = NewMessage()
	msg.Body.SetField(field.NewStringField(9999, "hello"))
	err = dict.validate(*msg)
	c.Check(err, NotNil)
	c.Check(err.(InvalidTagNumberError).Tag, Equals, tag.Tag(9999))
}
