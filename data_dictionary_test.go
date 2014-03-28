package quickfixgo

import (
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

func (s *DataDictionaryTests) createFIX40NewOrderSingle() *Message {
	msg := NewMessage()
	msg.Header.FieldMap.SetField(NewStringField(tag.MsgType, "D"))
	msg.Header.FieldMap.SetField(NewStringField(tag.BeginString, "FIX.4.0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.BodyLength, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.SenderCompID, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.TargetCompID, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.SendingTime, "0"))

	msg.Body.SetField(NewStringField(tag.ClOrdID, "A"))
	msg.Body.SetField(NewStringField(tag.HandlInst, "1"))
	msg.Body.SetField(NewStringField(tag.Symbol, "A"))
	msg.Body.SetField(NewStringField(tag.Side, "1"))
	msg.Body.SetField(NewStringField(tag.OrderQty, "A"))
	msg.Body.SetField(NewStringField(tag.OrdType, "1"))

	msg.Trailer.FieldMap.SetField(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *DataDictionaryTests) createFIX43NewOrderSingle() *Message {
	msg := NewMessage()
	msg.Header.FieldMap.SetField(NewStringField(tag.MsgType, "D"))
	msg.Header.FieldMap.SetField(NewStringField(tag.BeginString, "FIX.4.3"))
	msg.Header.FieldMap.SetField(NewStringField(tag.BodyLength, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.SenderCompID, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.TargetCompID, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.SendingTime, "0"))

	msg.Body.SetField(NewStringField(tag.ClOrdID, "A"))
	msg.Body.SetField(NewStringField(tag.HandlInst, "1"))
	msg.Body.SetField(NewStringField(tag.Symbol, "A"))
	msg.Body.SetField(NewStringField(tag.Side, "1"))
	msg.Body.SetField(NewStringField(tag.OrderQty, "A"))
	msg.Body.SetField(NewStringField(tag.OrdType, "1"))
	msg.Body.SetField(NewStringField(tag.TransactTime, "1"))

	msg.Trailer.FieldMap.SetField(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *DataDictionaryTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := s.createFIX40NewOrderSingle()
	msg.Header.FieldMap.SetField(NewStringField(9999, "hello"))
	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	msg = s.createFIX40NewOrderSingle()
	msg.Trailer.FieldMap.SetField(NewStringField(9999, "hello"))
	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	msg = NewMessage()
	msg = s.createFIX40NewOrderSingle()
	msg.Body.SetField(NewStringField(9999, "hello"))
	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))
}

func (s *DataDictionaryTests) TestValidateTagNotDefinedForMessage(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := s.createFIX40NewOrderSingle()
	msg.Body.SetField(NewStringField(41, "hello"))

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))

	msg = s.createFIX40NewOrderSingle()
	msg.Header.SetField(NewStringField(41, "hello"))

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))

	msg = s.createFIX40NewOrderSingle()
	msg.Trailer.SetField(NewStringField(41, "hello"))

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))
}

func (s *DataDictionaryTests) TestValidateTagNotDefinedForMessageComponent(c *C) {
	dict, _ := NewDataDictionary("spec/FIX43.xml")
	msg := s.createFIX43NewOrderSingle()

	reject := dict.validate(*msg)
	c.Check(reject, IsNil)
}

func (s *DataDictionaryTests) TestValidateFieldNotFound(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := s.createFIX40NewOrderSingle()
	msg.Body.fields = make(map[tag.Tag]*fieldValue)
	msg.Body.SetField(NewStringField(tag.ClOrdID, "A"))
	msg.Body.SetField(NewStringField(tag.HandlInst, "A"))
	msg.Body.SetField(NewStringField(tag.Symbol, "A"))
	msg.Body.SetField(NewStringField(tag.Side, "A"))
	msg.Body.SetField(NewStringField(tag.OrderQty, "A"))

	//ord type is required
	//msg.Body.SetField(NewStringField(tag.OrdType, "A"))

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.OrdType)

	msg = s.createFIX40NewOrderSingle()
	msg.Header.init()
	msg.Header.FieldMap.SetField(NewStringField(tag.MsgType, "D"))
	msg.Header.FieldMap.SetField(NewStringField(tag.BeginString, "FIX.4.0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.BodyLength, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.SenderCompID, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.TargetCompID, "0"))
	msg.Header.FieldMap.SetField(NewStringField(tag.MsgSeqNum, "0"))
	//sending time is required
	//msg.Header.FieldMap.SetField(NewStringField(tag.SendingTime, "0"))

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.SendingTime)

	msg = s.createFIX40NewOrderSingle()
	msg.Trailer.init()

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.CheckSum)
}

func (s *DataDictionaryTests) TestValidateTagSpecifiedWithoutAValue(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")
	msg := s.createFIX40NewOrderSingle()
	msg.Body.SetField(NewStringField(tag.ClientID, ""))
	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagSpecifiedWithoutAValue)
	c.Check(*reject.RefTagID(), Equals, tag.ClientID)
}

func (s *DataDictionaryTests) TestValidateInvalidMsgType(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	msg := s.createFIX40NewOrderSingle()
	msg.Header.SetField(NewStringField(tag.MsgType, "z"))

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidMsgType)
}

func (s *DataDictionaryTests) TestValidateValueIsIncorrect(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")
	msg := s.createFIX40NewOrderSingle()
	msg.Body.SetField(NewStringField(tag.HandlInst, "4"))
	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, ValueIsIncorrect)
	c.Check(*reject.RefTagID(), Equals, tag.HandlInst)
}
