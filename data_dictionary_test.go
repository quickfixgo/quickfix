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

func (s *DataDictionaryTests) createFIX40NewOrderSingle() *MessageBuilder {
	msg := NewMessageBuilder()
	msg.Header.SetField(NewStringField(tag.MsgType, "D"))
	msg.Header.SetField(NewStringField(tag.BeginString, "FIX.4.0"))
	msg.Header.SetField(NewStringField(tag.BodyLength, "0"))
	msg.Header.SetField(NewStringField(tag.SenderCompID, "0"))
	msg.Header.SetField(NewStringField(tag.TargetCompID, "0"))
	msg.Header.SetField(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.SetField(NewStringField(tag.SendingTime, "0"))

	msg.Body.SetField(NewStringField(tag.ClOrdID, "A"))
	msg.Body.SetField(NewStringField(tag.HandlInst, "1"))
	msg.Body.SetField(NewStringField(tag.Symbol, "A"))
	msg.Body.SetField(NewStringField(tag.Side, "1"))
	msg.Body.SetField(NewStringField(tag.OrderQty, "A"))
	msg.Body.SetField(NewStringField(tag.OrdType, "1"))

	msg.Trailer.SetField(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *DataDictionaryTests) createFIX43NewOrderSingle() *MessageBuilder {
	msg := NewMessageBuilder()
	msg.Header.SetField(NewStringField(tag.MsgType, "D"))
	msg.Header.SetField(NewStringField(tag.BeginString, "FIX.4.3"))
	msg.Header.SetField(NewStringField(tag.BodyLength, "0"))
	msg.Header.SetField(NewStringField(tag.SenderCompID, "0"))
	msg.Header.SetField(NewStringField(tag.TargetCompID, "0"))
	msg.Header.SetField(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.SetField(NewStringField(tag.SendingTime, "0"))

	msg.Body.SetField(NewStringField(tag.ClOrdID, "A"))
	msg.Body.SetField(NewStringField(tag.HandlInst, "1"))
	msg.Body.SetField(NewStringField(tag.Symbol, "A"))
	msg.Body.SetField(NewStringField(tag.Side, "1"))
	msg.Body.SetField(NewStringField(tag.OrderQty, "A"))
	msg.Body.SetField(NewStringField(tag.OrdType, "1"))
	msg.Body.SetField(NewStringField(tag.TransactTime, "1"))

	msg.Trailer.SetField(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *DataDictionaryTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.SetField(NewStringField(9999, "hello"))
	msg, _ := builder.Build()
	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Trailer.SetField(NewStringField(9999, "hello"))
	msg, _ = builder.Build()
	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Body.SetField(NewStringField(9999, "hello"))
	msg, _ = builder.Build()
	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))
}

func (s *DataDictionaryTests) TestValidateTagNotDefinedForMessage(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Body.SetField(NewStringField(41, "hello"))
	msg, _ := builder.Build()

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))

	builder = s.createFIX40NewOrderSingle()
	builder.Header.SetField(NewStringField(41, "hello"))
	msg, _ = builder.Build()

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))

	builder = s.createFIX40NewOrderSingle()
	builder.Trailer.SetField(NewStringField(41, "hello"))
	msg, _ = builder.Build()

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))
}

func (s *DataDictionaryTests) TestValidateTagNotDefinedForMessageComponent(c *C) {
	dict, _ := NewDataDictionary("spec/FIX43.xml")
	builder := s.createFIX43NewOrderSingle()

	msg, _ := builder.Build()

	reject := dict.validate(*msg)
	c.Check(reject, IsNil)
}

func (s *DataDictionaryTests) TestValidateFieldNotFound(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Body.fields = make(map[tag.Tag]*fieldBytes)
	builder.Body.SetField(NewStringField(tag.ClOrdID, "A"))
	builder.Body.SetField(NewStringField(tag.HandlInst, "A"))
	builder.Body.SetField(NewStringField(tag.Symbol, "A"))
	builder.Body.SetField(NewStringField(tag.Side, "A"))
	builder.Body.SetField(NewStringField(tag.OrderQty, "A"))

	//ord type is required
	//builder.Body.SetField(NewStringField(tag.OrdType, "A"))
	msg, _ := builder.Build()

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.OrdType)

	builder = s.createFIX40NewOrderSingle()
	builder.Header.init()
	builder.Header.SetField(NewStringField(tag.MsgType, "D"))
	builder.Header.SetField(NewStringField(tag.BeginString, "FIX.4.0"))
	builder.Header.SetField(NewStringField(tag.BodyLength, "0"))
	builder.Header.SetField(NewStringField(tag.SenderCompID, "0"))
	builder.Header.SetField(NewStringField(tag.TargetCompID, "0"))
	builder.Header.SetField(NewStringField(tag.MsgSeqNum, "0"))
	//sending time is required
	//msg.Header.FieldMap.SetField(NewStringField(tag.SendingTime, "0"))

	msg, _ = builder.Build()

	reject = dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.SendingTime)
}

func (s *DataDictionaryTests) TestValidateTagSpecifiedWithoutAValue(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.SetField(NewStringField(tag.ClientID, ""))
	msg, _ := builder.Build()

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagSpecifiedWithoutAValue)
	c.Check(*reject.RefTagID(), Equals, tag.ClientID)
}

func (s *DataDictionaryTests) TestValidateInvalidMsgType(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.SetField(NewStringField(tag.MsgType, "z"))
	msg, _ := builder.Build()

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidMsgType)
}

func (s *DataDictionaryTests) TestValidateValueIsIncorrect(c *C) {
	dict, _ := NewDataDictionary("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.SetField(NewStringField(tag.HandlInst, "4"))
	msg, _ := builder.Build()

	reject := dict.validate(*msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, ValueIsIncorrect)
	c.Check(*reject.RefTagID(), Equals, tag.HandlInst)
}
