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

func (s *DataDictionaryTests) createNewOrderSingle() *Message {
	msg := NewMessage()
	msg.Header.FieldMap.SetField(field.NewStringField(tag.MsgType, "D"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.BeginString, "FIX.4.0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.BodyLength, "0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.SenderCompID, "0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.TargetCompID, "0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.SendingTime, "0"))

	msg.Body.SetField(field.NewStringField(tag.ClOrdID, "A"))
	msg.Body.SetField(field.NewStringField(tag.HandlInst, "A"))
	msg.Body.SetField(field.NewStringField(tag.Symbol, "A"))
	msg.Body.SetField(field.NewStringField(tag.Side, "A"))
	msg.Body.SetField(field.NewStringField(tag.OrderQty, "A"))
	msg.Body.SetField(field.NewStringField(tag.OrdType, "A"))

	msg.Trailer.FieldMap.SetField(field.NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *DataDictionaryTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := s.createNewOrderSingle()
	msg.Header.FieldMap.SetField(field.NewStringField(9999, "hello"))
	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	msg = s.createNewOrderSingle()
	msg.Trailer.FieldMap.SetField(field.NewStringField(9999, "hello"))
	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	msg = NewMessage()
	msg = s.createNewOrderSingle()
	msg.Body.SetField(field.NewStringField(9999, "hello"))
	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))
}

func (s *DataDictionaryTests) TestValidateTagNotDefinedForMessage(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := s.createNewOrderSingle()
	msg.Body.SetField(field.NewStringField(41, "hello"))

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))

	msg = s.createNewOrderSingle()
	msg.Header.SetField(field.NewStringField(41, "hello"))

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))

	msg = s.createNewOrderSingle()
	msg.Trailer.SetField(field.NewStringField(41, "hello"))

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))
}

func (s *DataDictionaryTests) TestValidateFieldNotFound(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := s.createNewOrderSingle()
	msg.Body.fields = make(map[tag.Tag]*fieldValue)
	msg.Body.SetField(field.NewStringField(tag.ClOrdID, "A"))
	msg.Body.SetField(field.NewStringField(tag.HandlInst, "A"))
	msg.Body.SetField(field.NewStringField(tag.Symbol, "A"))
	msg.Body.SetField(field.NewStringField(tag.Side, "A"))
	msg.Body.SetField(field.NewStringField(tag.OrderQty, "A"))

	//ord type is required
	//msg.Body.SetField(field.NewStringField(tag.OrdType, "A"))

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.OrdType)

	msg = s.createNewOrderSingle()
	msg.Header.init()
	msg.Header.FieldMap.SetField(field.NewStringField(tag.MsgType, "D"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.BeginString, "FIX.4.0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.BodyLength, "0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.SenderCompID, "0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.TargetCompID, "0"))
	msg.Header.FieldMap.SetField(field.NewStringField(tag.MsgSeqNum, "0"))
	//sending time is required
	//msg.Header.FieldMap.SetField(field.NewStringField(tag.SendingTime, "0"))

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.SendingTime)

	msg = s.createNewOrderSingle()
	msg.Trailer.init()

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.CheckSum)
}

func (s *DataDictionaryTests) TestValidateTagSpecifiedWithoutAValue(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")
	msg := s.createNewOrderSingle()
	msg.Body.SetField(field.NewStringField(tag.ClientID, ""))
	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagSpecifiedWithoutAValue)
	c.Check(*reject.RefTagID(), Equals, tag.ClientID)
}
