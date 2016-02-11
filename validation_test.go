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
	msg.Header.SetField(tagMsgType, FIXString("D"))
	msg.Header.SetField(tagBeginString, FIXString("FIX.4.0"))
	msg.Header.SetField(tagBodyLength, FIXString("0"))
	msg.Header.SetField(tagSenderCompID, FIXString("0"))
	msg.Header.SetField(tagTargetCompID, FIXString("0"))
	msg.Header.SetField(tagMsgSeqNum, FIXString("0"))
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Value: time.Now()})

	msg.Body.SetField(Tag(11), FIXString("A"))
	msg.Body.SetField(Tag(21), FIXString("1"))
	msg.Body.SetField(Tag(55), FIXString("A"))
	msg.Body.SetField(Tag(54), FIXString("1"))
	msg.Body.SetField(Tag(40), FIXString("1"))
	msg.Body.SetField(Tag(38), FIXInt(5))

	msg.Trailer.SetField(tagCheckSum, FIXString("000"))

	return msg
}

func (s *ValidationTests) createFIX43NewOrderSingle() Message {
	msg := NewMessage()
	msg.Header.SetField(tagMsgType, FIXString("D"))
	msg.Header.SetField(tagBeginString, FIXString("FIX.4.3"))
	msg.Header.SetField(tagBodyLength, FIXString("0"))
	msg.Header.SetField(tagSenderCompID, FIXString("0"))
	msg.Header.SetField(tagTargetCompID, FIXString("0"))
	msg.Header.SetField(tagMsgSeqNum, FIXString("0"))
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Value: time.Now()})

	msg.Body.SetField(Tag(11), FIXString("A"))
	msg.Body.SetField(Tag(21), FIXString("1"))
	msg.Body.SetField(Tag(55), FIXString("A"))
	msg.Body.SetField(Tag(54), FIXString("1"))
	msg.Body.SetField(Tag(38), FIXInt(5))
	msg.Body.SetField(Tag(40), FIXString("1"))
	msg.Body.SetField(Tag(60), FIXUTCTimestamp{Value: time.Now()})

	msg.Trailer.SetField(tagCheckSum, FIXString("000"))

	return msg
}

func (s *ValidationTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.SetField(9999, FIXString("hello"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)
	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonInvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Trailer.SetField(9999, FIXString("hello"))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	reject = validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonInvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Body.SetField(9999, FIXString("hello"))
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
	builder.Body.SetField(41, FIXString("hello"))
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
	builder.Header.SetField(tagMsgType, FIXString("D"))
	builder.Header.SetField(tagBeginString, FIXString("FIX.4.0"))
	builder.Header.SetField(tagBodyLength, FIXString("0"))
	builder.Header.SetField(tagSenderCompID, FIXString("0"))
	builder.Header.SetField(tagTargetCompID, FIXString("0"))
	builder.Header.SetField(tagMsgSeqNum, FIXString("0"))
	builder.Header.SetField(tagSendingTime, FIXUTCTimestamp{Value: time.Now()})
	builder.Trailer.SetField(tagCheckSum, FIXString("000"))

	builder.Body.SetField(Tag(11), FIXString("A"))
	builder.Body.SetField(Tag(21), FIXString("A"))
	builder.Body.SetField(Tag(55), FIXString("A"))
	builder.Body.SetField(Tag(54), FIXString("A"))
	builder.Body.SetField(Tag(38), FIXString("A"))

	//ord type is required
	//builder.Body.SetField(Tag(40), "A"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonRequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, Tag(40))

	builder = NewMessage()
	builder.Trailer.SetField(tagCheckSum, FIXString("000"))
	builder.Body.SetField(Tag(11), FIXString("A"))
	builder.Body.SetField(Tag(21), FIXString("A"))
	builder.Body.SetField(Tag(55), FIXString("A"))
	builder.Body.SetField(Tag(54), FIXString("A"))
	builder.Body.SetField(Tag(38), FIXString("A"))

	builder.Header.SetField(tagMsgType, FIXString("D"))
	builder.Header.SetField(tagBeginString, FIXString("FIX.4.0"))
	builder.Header.SetField(tagBodyLength, FIXString("0"))
	builder.Header.SetField(tagSenderCompID, FIXString("0"))
	builder.Header.SetField(tagTargetCompID, FIXString("0"))
	builder.Header.SetField(tagMsgSeqNum, FIXString("0"))
	//sending time is required
	//msg.Header.FieldMap.SetField(tag.SendingTime, "0"))

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
	builder.Body.SetField(Tag(109), FIXString(""))
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
	builder.Header.SetField(tagMsgType, FIXString("z"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	reject := validate(dict, msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonInvalidMsgType)
}

func (s *ValidationTests) TestValidateValueIsIncorrect(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.SetField(Tag(21), FIXString("4"))
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
	builder.Body.SetField(Tag(38), FIXString("+200.00"))
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
	builder.Body.SetField(tagOnBehalfOfCompID, FIXString("CWB"))
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

	TagValues := make([]tagValue, 1)
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

	var repField1 tagValue
	var repField2 tagValue
	repField1.init(Tag(2), []byte("a"))
	repField2.init(Tag(3), []byte("a"))

	//non-repeating
	var groupID tagValue
	groupID.init(Tag(1), []byte("1"))
	fields := []tagValue{groupID, repField1}
	remFields, reject := validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)

	fields = []tagValue{groupID, repField1, repField2}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 0)
	c.Check(reject, IsNil)

	//test with trailing tag not in group
	var otherField tagValue
	otherField.init(Tag(500), []byte("blah"))
	fields = []tagValue{groupID, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	//repeats
	groupID.init(Tag(1), []byte("2"))
	fields = []tagValue{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	groupID.init(Tag(1), []byte("3"))
	fields = []tagValue{groupID, repField1, repField2, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(len(remFields), Equals, 1)
	c.Check(reject, IsNil)

	//REJECT: group size declared > actual group size
	groupID.init(Tag(1), []byte("3"))
	fields = []tagValue{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectNumInGroupCountForRepeatingGroup)

	groupID.init(Tag(1), []byte("3"))
	fields = []tagValue{groupID, repField1, repField1, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectNumInGroupCountForRepeatingGroup)

	//REJECT: group size declared < actual group size
	groupID.init(Tag(1), []byte("1"))
	fields = []tagValue{groupID, repField1, repField2, repField1, repField2, otherField}
	remFields, reject = validateVisitGroupField(groupFieldDef, fields)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, rejectReasonIncorrectNumInGroupCountForRepeatingGroup)
}
