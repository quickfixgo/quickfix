package quickfix

import (
	"github.com/quickfixgo/quickfix/datadictionary"
	. "gopkg.in/check.v1"
	"time"
)

var _ = Suite(&ValidationTests{})

type ValidationTests struct{}

func (s *ValidationTests) createFIX40NewOrderSingle() Message {
	msg := NewMessage()
	msg.Header.Set(NewStringField(tagMsgType, "D"))
	msg.Header.Set(NewStringField(tagBeginString, "FIX.4.0"))
	msg.Header.Set(NewStringField(tagBodyLength, "0"))
	msg.Header.Set(NewStringField(tagSenderCompID, "0"))
	msg.Header.Set(NewStringField(tagTargetCompID, "0"))
	msg.Header.Set(NewStringField(tagMsgSeqNum, "0"))
	msg.Header.Set(NewUTCTimestampField(tagSendingTime, time.Now()))

	msg.Body.Set(NewStringField(Tag(11), "A"))
	msg.Body.Set(NewStringField(Tag(21), "1"))
	msg.Body.Set(NewStringField(Tag(55), "A"))
	msg.Body.Set(NewStringField(Tag(54), "1"))
	msg.Body.Set(NewStringField(Tag(40), "1"))
	msg.Body.Set(NewIntField(Tag(38), 5))

	msg.Trailer.Set(NewStringField(tagCheckSum, "000"))

	return msg
}

func (s *ValidationTests) createFIX43NewOrderSingle() Message {
	msg := NewMessage()
	msg.Header.Set(NewStringField(tagMsgType, "D"))
	msg.Header.Set(NewStringField(tagBeginString, "FIX.4.3"))
	msg.Header.Set(NewStringField(tagBodyLength, "0"))
	msg.Header.Set(NewStringField(tagSenderCompID, "0"))
	msg.Header.Set(NewStringField(tagTargetCompID, "0"))
	msg.Header.Set(NewStringField(tagMsgSeqNum, "0"))
	msg.Header.Set(NewUTCTimestampField(tagSendingTime, time.Now()))

	msg.Body.Set(NewStringField(Tag(11), "A"))
	msg.Body.Set(NewStringField(Tag(21), "1"))
	msg.Body.Set(NewStringField(Tag(55), "A"))
	msg.Body.Set(NewStringField(Tag(54), "1"))
	msg.Body.Set(NewIntField(Tag(38), 5))
	msg.Body.Set(NewStringField(Tag(40), "1"))
	msg.Body.Set(NewUTCTimestampField(Tag(60), time.Now()))

	msg.Trailer.Set(NewStringField(tagCheckSum, "000"))

	return msg
}

func (s *ValidationTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.Set(NewStringField(9999, "hello"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)
	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonInvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Trailer.Set(NewStringField(9999, "hello"))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	reject = validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonInvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(9999, "hello"))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	reject = validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonInvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, Tag(9999))
}

func (s *ValidationTests) TestValidateTagNotDefinedForMessage(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(41, "hello"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonTagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, Tag(41))
}

func (s *ValidationTests) TestValidateTagNotDefinedForMessageComponent(c *C) {
	dict, err := datadictionary.Parse("spec/FIX43.xml")
	c.Check(err, IsNil)
	builder := s.createFIX43NewOrderSingle()
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, IsNil)
}

func (s *ValidationTests) TestValidateFieldNotFound(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := NewMessage()
	builder.Header.Set(NewStringField(tagMsgType, "D"))
	builder.Header.Set(NewStringField(tagBeginString, "FIX.4.0"))
	builder.Header.Set(NewStringField(tagBodyLength, "0"))
	builder.Header.Set(NewStringField(tagSenderCompID, "0"))
	builder.Header.Set(NewStringField(tagTargetCompID, "0"))
	builder.Header.Set(NewStringField(tagMsgSeqNum, "0"))
	builder.Header.Set(NewUTCTimestampField(tagSendingTime, time.Now()))
	builder.Trailer.Set(NewStringField(tagCheckSum, "000"))

	builder.Body.Set(NewStringField(Tag(11), "A"))
	builder.Body.Set(NewStringField(Tag(21), "A"))
	builder.Body.Set(NewStringField(Tag(55), "A"))
	builder.Body.Set(NewStringField(Tag(54), "A"))
	builder.Body.Set(NewStringField(Tag(38), "A"))

	//ord type is required
	//builder.Body.Set(NewStringField(Tag(40), "A"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonRequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, Tag(40))

	builder = NewMessage()
	builder.Trailer.Set(NewStringField(tagCheckSum, "000"))
	builder.Body.Set(NewStringField(Tag(11), "A"))
	builder.Body.Set(NewStringField(Tag(21), "A"))
	builder.Body.Set(NewStringField(Tag(55), "A"))
	builder.Body.Set(NewStringField(Tag(54), "A"))
	builder.Body.Set(NewStringField(Tag(38), "A"))

	builder.Header.Set(NewStringField(tagMsgType, "D"))
	builder.Header.Set(NewStringField(tagBeginString, "FIX.4.0"))
	builder.Header.Set(NewStringField(tagBodyLength, "0"))
	builder.Header.Set(NewStringField(tagSenderCompID, "0"))
	builder.Header.Set(NewStringField(tagTargetCompID, "0"))
	builder.Header.Set(NewStringField(tagMsgSeqNum, "0"))
	//sending time is required
	//msg.Header.FieldMap.Set(NewStringField(tag.SendingTime, "0"))

	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	reject = validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonRequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tagSendingTime)
}

func (s *ValidationTests) TestValidateTagSpecifiedWithoutAValue(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(Tag(109), ""))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonTagSpecifiedWithoutAValue)
	c.Check(*reject.RefTagID(), Equals, Tag(109))
}

func (s *ValidationTests) TestValidateInvalidMsgType(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.Set(NewStringField(tagMsgType, "z"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonInvalidMsgType)
}

func (s *ValidationTests) TestValidateValueIsIncorrect(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(Tag(21), "4"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonValueIsIncorrect)
	c.Check(*reject.RefTagID(), Equals, Tag(21))
}

func (s *ValidationTests) TestValidateIncorrectDataFormatForValue(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(Tag(38), "+200.00"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectDataFormatForValue)
	c.Check(*reject.RefTagID(), Equals, Tag(38))
}

func (s *ValidationTests) TestValidateTagSpecifiedOutOfRequiredOrder(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	//should be in header
	builder.Body.Set(NewStringField(tagOnBehalfOfCompID, "CWB"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonTagSpecifiedOutOfRequiredOrder)
	c.Check(*reject.RefTagID(), Equals, tagOnBehalfOfCompID)
}

func (s *ValidationTests) TestValidateTagAppearsMoreThanOnce(c *C) {

	msg, err := parseMessage([]byte("8=FIX.4.09=10735=D34=249=TW52=20060102-15:04:0556=ISLD11=ID21=140=140=254=138=20055=INTC60=20060102-15:04:0510=234"))
	c.Check(err, IsNil)

	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonTagAppearsMoreThanOnce)
	c.Check(*reject.RefTagID(), Equals, Tag(40))
}

func (s *ValidationTests) TestFloatValidation(c *C) {
	msg, err := parseMessage([]byte("8=FIX.4.29=10635=D34=249=TW52=20140329-22:38:4556=ISLD11=ID21=140=154=138=+200.0055=INTC60=20140329-22:38:4510=178"))
	c.Check(err, IsNil)

	dict, _ := datadictionary.Parse("spec/FIX42.xml")
	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectDataFormatForValue)
}

func (s *ValidationTests) TestValidateVisitField(c *C) {
	fieldType := &datadictionary.FieldType{Name: "myfield", Tag: 11, Type: "STRING"}
	fieldDef := &datadictionary.FieldDef{FieldType: fieldType}

	TagValues := make([]TagValue, 1)
	TagValues[0].init(Tag(11), []byte("value"))
	remFields, reject := validateVisitField(fieldDef, TagValues)
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)
}

func (s *ValidationTests) TestValidateVisitFieldGroup(c *C) {
	fieldType1 := &datadictionary.FieldType{Name: "myfield", Tag: 2, Type: "STRING"}
	fieldDef1 := &datadictionary.FieldDef{FieldType: fieldType1, ChildFields: []*datadictionary.FieldDef{}}

	fieldType2 := &datadictionary.FieldType{Name: "myfield", Tag: 3, Type: "STRING"}
	fieldDef2 := &datadictionary.FieldDef{FieldType: fieldType2, ChildFields: []*datadictionary.FieldDef{}}

	groupFieldType := &datadictionary.FieldType{Name: "mygroupfield", Tag: 1, Type: "INT"}
	groupFieldDef := &datadictionary.FieldDef{FieldType: groupFieldType, ChildFields: []*datadictionary.FieldDef{fieldDef1, fieldDef2}}

	var repField1 TagValue
	var repField2 TagValue
	repField1.init(Tag(2), []byte("a"))
	repField2.init(Tag(3), []byte("a"))

	//non-repeating
	var groupID TagValue
	groupID.init(Tag(1), []byte("1"))
	fields := []TagValue{groupID, repField1}
	remFields, reject := validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)

	fields = []TagValue{groupID, repField1, repField2}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)

	//test with trailing tag not in group
	var otherField TagValue
	otherField.init(Tag(500), []byte("blah"))
	fields = []TagValue{groupID, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	//repeats
	groupID.init(Tag(1), []byte("2"))
	fields = []TagValue{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	groupID.init(Tag(1), []byte("3"))
	fields = []TagValue{groupID, repField1, repField2, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	//REJECT: group size declared > actual group size
	groupID.init(Tag(1), []byte("3"))
	fields = []TagValue{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectNumInGroupCountForRepeatingGroup)

	groupID.init(Tag(1), []byte("3"))
	fields = []TagValue{groupID, repField1, repField1, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectNumInGroupCountForRepeatingGroup)

	//REJECT: group size declared < actual group size
	groupID.init(Tag(1), []byte("1"))
	fields = []TagValue{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectNumInGroupCountForRepeatingGroup)
}
