package quickfix

import (
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/tag"
	. "launchpad.net/gocheck"
	"time"
)

var _ = Suite(&ValidationTests{})

type ValidationTests struct{}

func (s *ValidationTests) createFIX40NewOrderSingle() *MessageBuilder {
	msg := NewMessageBuilder()
	msg.Header.Set(NewStringField(tag.MsgType, "D"))
	msg.Header.Set(NewStringField(tag.BeginString, "FIX.4.0"))
	msg.Header.Set(NewStringField(tag.BodyLength, "0"))
	msg.Header.Set(NewStringField(tag.SenderCompID, "0"))
	msg.Header.Set(NewStringField(tag.TargetCompID, "0"))
	msg.Header.Set(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.Set(NewUTCTimestampField(tag.SendingTime, time.Now()))

	msg.Body.Set(NewStringField(tag.ClOrdID, "A"))
	msg.Body.Set(NewStringField(tag.HandlInst, "1"))
	msg.Body.Set(NewStringField(tag.Symbol, "A"))
	msg.Body.Set(NewStringField(tag.Side, "1"))
	msg.Body.Set(NewStringField(tag.OrdType, "1"))
	msg.Body.Set(NewIntField(tag.OrderQty, 5))

	msg.Trailer.Set(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *ValidationTests) createFIX43NewOrderSingle() *MessageBuilder {
	msg := NewMessageBuilder()
	msg.Header.Set(NewStringField(tag.MsgType, "D"))
	msg.Header.Set(NewStringField(tag.BeginString, "FIX.4.3"))
	msg.Header.Set(NewStringField(tag.BodyLength, "0"))
	msg.Header.Set(NewStringField(tag.SenderCompID, "0"))
	msg.Header.Set(NewStringField(tag.TargetCompID, "0"))
	msg.Header.Set(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.Set(NewUTCTimestampField(tag.SendingTime, time.Now()))

	msg.Body.Set(NewStringField(tag.ClOrdID, "A"))
	msg.Body.Set(NewStringField(tag.HandlInst, "1"))
	msg.Body.Set(NewStringField(tag.Symbol, "A"))
	msg.Body.Set(NewStringField(tag.Side, "1"))
	msg.Body.Set(NewIntField(tag.OrderQty, 5))
	msg.Body.Set(NewStringField(tag.OrdType, "1"))
	msg.Body.Set(NewUTCTimestampField(tag.TransactTime, time.Now()))

	msg.Trailer.Set(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *ValidationTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.Set(NewStringField(9999, "hello"))
	msg, _ := builder.Build()
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Trailer.Set(NewStringField(9999, "hello"))
	msg, _ = builder.Build()
	reject = Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(9999, "hello"))
	msg, _ = builder.Build()
	reject = Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))
}

func (s *ValidationTests) TestValidateTagNotDefinedForMessage(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(41, "hello"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))
}

func (s *ValidationTests) TestValidateTagNotDefinedForMessageComponent(c *C) {
	dict, err := datadictionary.Parse("spec/FIX43.xml")
	c.Check(err, IsNil)
	builder := s.createFIX43NewOrderSingle()
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, IsNil)
}

func (s *ValidationTests) TestValidateFieldNotFound(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Body.init(normalFieldOrder)
	builder.Body.Set(NewStringField(tag.ClOrdID, "A"))
	builder.Body.Set(NewStringField(tag.HandlInst, "A"))
	builder.Body.Set(NewStringField(tag.Symbol, "A"))
	builder.Body.Set(NewStringField(tag.Side, "A"))
	builder.Body.Set(NewStringField(tag.OrderQty, "A"))

	//ord type is required
	//builder.Body.Set(NewStringField(tag.OrdType, "A"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.OrdType)

	builder = s.createFIX40NewOrderSingle()
	builder.Header.init()
	builder.Header.Set(NewStringField(tag.MsgType, "D"))
	builder.Header.Set(NewStringField(tag.BeginString, "FIX.4.0"))
	builder.Header.Set(NewStringField(tag.BodyLength, "0"))
	builder.Header.Set(NewStringField(tag.SenderCompID, "0"))
	builder.Header.Set(NewStringField(tag.TargetCompID, "0"))
	builder.Header.Set(NewStringField(tag.MsgSeqNum, "0"))
	//sending time is required
	//msg.Header.FieldMap.Set(NewStringField(tag.SendingTime, "0"))

	msg, _ = builder.Build()

	reject = Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.SendingTime)
}

func (s *ValidationTests) TestValidateTagSpecifiedWithoutAValue(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(tag.ClientID, ""))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagSpecifiedWithoutAValue)
	c.Check(*reject.RefTagID(), Equals, tag.ClientID)
}

func (s *ValidationTests) TestValidateInvalidMsgType(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.Set(NewStringField(tag.MsgType, "z"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidMsgType)
}

func (s *ValidationTests) TestValidateValueIsIncorrect(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(tag.HandlInst, "4"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, ValueIsIncorrect)
	c.Check(*reject.RefTagID(), Equals, tag.HandlInst)
}

func (s *ValidationTests) TestValidateIncorrectDataFormatForValue(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(tag.OrderQty, "+200.00"))
	msg, _ := builder.Build()
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, IncorrectDataFormatForValue)
	c.Check(*reject.RefTagID(), Equals, tag.OrderQty)
}

func (s *ValidationTests) TestValidateTagSpecifiedOutOfRequiredOrder(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	//should be in header
	builder.Body.Set(NewStringField(tag.OnBehalfOfCompID, "CWB"))
	msg, _ := builder.Build()
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagSpecifiedOutOfRequiredOrder)
	c.Check(*reject.RefTagID(), Equals, tag.OnBehalfOfCompID)
}

func (s *ValidationTests) TestValidateTagAppearsMoreThanOnce(c *C) {

	msg, err := MessageFromParsedBytes([]byte("8=FIX.4.09=10735=D34=249=TW52=20060102-15:04:0556=ISLD11=ID21=140=140=254=138=20055=INTC60=20060102-15:04:0510=234"))
	c.Check(err, IsNil)

	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagAppearsMoreThanOnce)
	c.Check(*reject.RefTagID(), Equals, tag.OrdType)
}

func (s *ValidationTests) TestFloatValidation(c *C) {
	msg, err := MessageFromParsedBytes([]byte("8=FIX.4.29=10635=D34=249=TW52=20140329-22:38:4556=ISLD11=ID21=140=154=138=+200.0055=INTC60=20140329-22:38:4510=178"))
	c.Check(err, IsNil)

	dict, _ := datadictionary.Parse("spec/FIX42.xml")
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, IncorrectDataFormatForValue)
}

func (s *ValidationTests) TestValidateVisitField(c *C) {
	fieldType := &datadictionary.FieldType{Name: "myfield", Tag: tag.ClOrdID, Type: "STRING"}
	fieldDef := &datadictionary.FieldDef{FieldType: fieldType}

	fields := []*fieldBytes{newFieldBytes(tag.ClOrdID, []byte("value"))}
	remFields, reject := validateVisitField(fieldDef, fields, Message{})
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)
}

func (s *ValidationTests) TestValidateVisitFieldGroup(c *C) {
	fieldType1 := &datadictionary.FieldType{Name: "myfield", Tag: tag.Tag(2), Type: "STRING"}
	fieldDef1 := &datadictionary.FieldDef{FieldType: fieldType1, ChildFields: []*datadictionary.FieldDef{}}

	fieldType2 := &datadictionary.FieldType{Name: "myfield", Tag: tag.Tag(3), Type: "STRING"}
	fieldDef2 := &datadictionary.FieldDef{FieldType: fieldType2, ChildFields: []*datadictionary.FieldDef{}}

	groupFieldType := &datadictionary.FieldType{Name: "mygroupfield", Tag: tag.Tag(1), Type: "INT"}
	groupFieldDef := &datadictionary.FieldDef{FieldType: groupFieldType, ChildFields: []*datadictionary.FieldDef{fieldDef1, fieldDef2}}

	repField1 := newFieldBytes(tag.Tag(2), []byte("a"))
	repField2 := newFieldBytes(tag.Tag(3), []byte("a"))

	//non-repeating
	groupID := newFieldBytes(tag.Tag(1), []byte("1"))
	fields := []*fieldBytes{groupID, repField1}
	remFields, reject := validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)

	fields = []*fieldBytes{groupID, repField1, repField2}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)

	//test with trailing tag not in group
	otherField := newFieldBytes(tag.Tag(500), []byte("blah"))
	fields = []*fieldBytes{groupID, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	//repeats
	groupID = newFieldBytes(tag.Tag(1), []byte("2"))
	fields = []*fieldBytes{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	groupID = newFieldBytes(tag.Tag(1), []byte("3"))
	fields = []*fieldBytes{groupID, repField1, repField2, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	//REJECT: group size declared > actual group size
	groupID = newFieldBytes(tag.Tag(1), []byte("3"))
	fields = []*fieldBytes{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, IncorrectNumInGroupCountForRepeatingGroup)

	groupID = newFieldBytes(tag.Tag(1), []byte("3"))
	fields = []*fieldBytes{groupID, repField1, repField1, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, IncorrectNumInGroupCountForRepeatingGroup)

	//REJECT: group size declared < actual group size
	groupID = newFieldBytes(tag.Tag(1), []byte("1"))
	fields = []*fieldBytes{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields, Message{})
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, IncorrectNumInGroupCountForRepeatingGroup)
}
