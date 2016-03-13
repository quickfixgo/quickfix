package quickfix

import (
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/datadictionary"
)

type validateTest struct {
	TestName string
	*datadictionary.DataDictionary
	MessageBytes         []byte
	ExpectedRejectReason int
	ExpectedRefTagID     *Tag
	DoNotExpectReject    bool
}

func TestValidate(t *testing.T) {
	var tests = []validateTest{
		tcInvalidTagNumberHeader(),
		tcInvalidTagNumberBody(),
		tcInvalidTagNumberTrailer(),
		tcTagSpecifiedWithoutAValue(),
		tcInvalidMsgType(),
		tcValueIsIncorrect(),
		tcIncorrectDataFormatForValue(),
		tcTagSpecifiedOutOfRequiredOrder(),
		tcTagAppearsMoreThanOnce(),
		tcFloatValidation(),
		tcTagNotDefinedForMessage(),
		tcTagIsDefinedForMessage(),
		tcFieldNotFoundBody(),
		tcFieldNotFoundHeader(),
	}

	for _, test := range tests {
		msg, _ := parseMessage(test.MessageBytes)
		reject := validate(test.DataDictionary, msg)

		switch {
		case reject == nil && test.DoNotExpectReject:
			continue

		case reject != nil && test.DoNotExpectReject:
			t.Errorf("%v: Unexpected reject: %v", test.TestName, reject)
			continue

		case reject == nil:
			t.Errorf("%v: Expected reject", test.TestName)
			continue
		}

		if reject.RejectReason() != test.ExpectedRejectReason {
			t.Errorf("%v: Expected reason %v got %v", test.TestName, test.ExpectedRejectReason, reject.RejectReason())
		}

		switch {
		case reject.RefTagID() == nil && test.ExpectedRefTagID == nil:
		//ok, expected and actual ref tag not set
		case reject.RefTagID() != nil && test.ExpectedRefTagID == nil:
			t.Errorf("%v: Unexpected RefTag '%v'", test.TestName, *reject.RefTagID())
		case reject.RefTagID() == nil && test.ExpectedRefTagID != nil:
			t.Errorf("%v: Expected RefTag '%v'", test.TestName, *test.ExpectedRefTagID)
		case *reject.RefTagID() == *test.ExpectedRefTagID:
			//ok, tags equal
		default:
			t.Errorf("%v: Expected RefTag '%v' got '%v'", test.TestName, *test.ExpectedRefTagID, *reject.RefTagID())
		}
	}
}

func createFIX40NewOrderSingle() Message {
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

func createFIX43NewOrderSingle() Message {
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

func tcInvalidTagNumberHeader() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	invalidHeaderFieldMessage := createFIX40NewOrderSingle()
	tag := Tag(9999)
	invalidHeaderFieldMessage.Header.SetField(tag, FIXString("hello"))
	msgBytes, _ := invalidHeaderFieldMessage.Build()

	return validateTest{
		TestName:             "Invalid Tag Number Header",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}
func tcInvalidTagNumberBody() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	invalidBodyFieldMessage := createFIX40NewOrderSingle()
	tag := Tag(9999)
	invalidBodyFieldMessage.Body.SetField(tag, FIXString("hello"))
	msgBytes, _ := invalidBodyFieldMessage.Build()

	return validateTest{
		TestName:             "Invalid Tag Number Body",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcInvalidTagNumberTrailer() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	invalidTrailerFieldMessage := createFIX40NewOrderSingle()
	tag := Tag(9999)
	invalidTrailerFieldMessage.Trailer.SetField(tag, FIXString("hello"))
	msgBytes, _ := invalidTrailerFieldMessage.Build()

	return validateTest{
		TestName:             "Invalid Tag Number Trailer",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagNotDefinedForMessage() validateTest {
	dict40, _ := datadictionary.Parse("spec/FIX40.xml")

	invalidMsg := createFIX40NewOrderSingle()
	tag := Tag(41)
	invalidMsg.Body.SetField(tag, FIXString("hello"))
	msgBytes, _ := invalidMsg.Build()

	return validateTest{
		TestName:             "Tag Not Defined For Message",
		DataDictionary:       dict40,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagNotDefinedForThisMessageType,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagIsDefinedForMessage() validateTest {
	//compare to tcTagIsNotDefinedForMessage
	dict43, _ := datadictionary.Parse("spec/FIX43.xml")
	validMsg := createFIX43NewOrderSingle()
	msgBytes, _ := validMsg.Build()

	return validateTest{
		TestName:          "TagIsDefinedForMessage",
		DataDictionary:    dict43,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcFieldNotFoundBody() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	invalidMsg1 := NewMessage()
	invalidMsg1.Header.SetField(tagMsgType, FIXString("D")).
		SetField(tagBeginString, FIXString("FIX.4.0")).
		SetField(tagBodyLength, FIXString("0")).
		SetField(tagSenderCompID, FIXString("0")).
		SetField(tagTargetCompID, FIXString("0")).
		SetField(tagMsgSeqNum, FIXString("0")).
		SetField(tagSendingTime, FIXUTCTimestamp{Value: time.Now()})
	invalidMsg1.Trailer.SetField(tagCheckSum, FIXString("000"))

	invalidMsg1.Body.SetField(Tag(11), FIXString("A")).
		SetField(Tag(21), FIXString("A")).
		SetField(Tag(55), FIXString("A")).
		SetField(Tag(54), FIXString("A")).
		SetField(Tag(38), FIXString("A"))

	tag := Tag(40)
	//ord type is required
	//invalidMsg1.Body.SetField(Tag(40), "A"))

	msgBytes, _ := invalidMsg1.Build()

	return validateTest{
		TestName:             "FieldNotFoundBody",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonRequiredTagMissing,
		ExpectedRefTagID:     &tag,
	}
}

func tcFieldNotFoundHeader() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	invalidMsg2 := NewMessage()
	invalidMsg2.Trailer.SetField(tagCheckSum, FIXString("000"))
	invalidMsg2.Body.SetField(Tag(11), FIXString("A")).
		SetField(Tag(21), FIXString("A")).
		SetField(Tag(55), FIXString("A")).
		SetField(Tag(54), FIXString("A")).
		SetField(Tag(38), FIXString("A"))

	invalidMsg2.Header.SetField(tagMsgType, FIXString("D")).
		SetField(tagBeginString, FIXString("FIX.4.0")).
		SetField(tagBodyLength, FIXString("0")).
		SetField(tagSenderCompID, FIXString("0")).
		SetField(tagTargetCompID, FIXString("0")).
		SetField(tagMsgSeqNum, FIXString("0"))
	//sending time is required
	//invalidMsg2.Header.FieldMap.SetField(tag.SendingTime, "0"))

	tag := tagSendingTime
	msgBytes, _ := invalidMsg2.Build()

	return validateTest{
		TestName:             "FieldNotFoundHeader",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonRequiredTagMissing,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagSpecifiedWithoutAValue() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := createFIX40NewOrderSingle()

	bogusTag := Tag(109)
	builder.Body.SetField(bogusTag, FIXString(""))
	msgBytes, _ := builder.Build()

	return validateTest{
		TestName:             "Tag SpecifiedWithoutAValue",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedWithoutAValue,
		ExpectedRefTagID:     &bogusTag,
	}
}

func tcInvalidMsgType() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := createFIX40NewOrderSingle()
	builder.Header.SetField(tagMsgType, FIXString("z"))
	msgBytes, _ := builder.Build()

	return validateTest{
		TestName:             "Invalid MsgType",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidMsgType,
	}
}

func tcValueIsIncorrect() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	tag := Tag(21)
	builder := createFIX40NewOrderSingle()
	builder.Body.SetField(tag, FIXString("4"))
	msgBytes, _ := builder.Build()

	return validateTest{
		TestName:             "ValueIsIncorrect",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonValueIsIncorrect,
		ExpectedRefTagID:     &tag,
	}
}

func tcIncorrectDataFormatForValue() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := createFIX40NewOrderSingle()
	tag := Tag(38)
	builder.Body.SetField(tag, FIXString("+200.00"))
	msgBytes, _ := builder.Build()

	return validateTest{
		TestName:             "IncorrectDataFormatForValue",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonIncorrectDataFormatForValue,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagSpecifiedOutOfRequiredOrder() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := createFIX40NewOrderSingle()
	tag := tagOnBehalfOfCompID
	//should be in header
	builder.Body.SetField(tag, FIXString("CWB"))
	msgBytes, _ := builder.Build()

	return validateTest{
		TestName:             "Tag specified out of required order",
		DataDictionary:       dict,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedOutOfRequiredOrder,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagAppearsMoreThanOnce() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	tag := Tag(40)

	return validateTest{
		TestName:       "Tag appears more than once",
		DataDictionary: dict,
		MessageBytes: []byte("8=FIX.4.09=10735=D34=249=TW52=20060102-15:04:0556=ISLD11=ID21=140=140=254=138=20055=INTC60=20060102-15:04:0510=234"),
		ExpectedRejectReason: rejectReasonTagAppearsMoreThanOnce,
		ExpectedRefTagID:     &tag,
	}
}

func tcFloatValidation() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX42.xml")
	tag := Tag(38)
	return validateTest{
		TestName:       "FloatValidation",
		DataDictionary: dict,
		MessageBytes: []byte("8=FIX.4.29=10635=D34=249=TW52=20140329-22:38:4556=ISLD11=ID21=140=154=138=+200.0055=INTC60=20140329-22:38:4510=178"),
		ExpectedRejectReason: rejectReasonIncorrectDataFormatForValue,
		ExpectedRefTagID:     &tag,
	}
}

func TestValidateVisitField(t *testing.T) {
	fieldType := datadictionary.NewFieldType("myfield", 11, "STRING")
	fieldDef := &datadictionary.FieldDef{FieldType: fieldType}

	TagValues := make([]tagValue, 1)
	TagValues[0].init(Tag(11), []byte("value"))
	remFields, reject := validateVisitField(fieldDef, TagValues)
	if len(remFields) != 0 {
		t.Errorf("Expected len %v got %v", 0, len(remFields))
	}

	if reject != nil {
		t.Errorf("Unexpected Reject %v", reject)
	}
}

func TestValidateVisitFieldGroup(t *testing.T) {
	fieldType1 := datadictionary.NewFieldType("myfield", 2, "STRING")
	fieldDef1 := &datadictionary.FieldDef{FieldType: fieldType1, ChildFields: []*datadictionary.FieldDef{}}

	fieldType2 := datadictionary.NewFieldType("myfield", 3, "STRING")
	fieldDef2 := &datadictionary.FieldDef{FieldType: fieldType2, ChildFields: []*datadictionary.FieldDef{}}

	groupFieldType := datadictionary.NewFieldType("mygroupfield", 1, "INT")
	groupFieldDef := &datadictionary.FieldDef{FieldType: groupFieldType, ChildFields: []*datadictionary.FieldDef{fieldDef1, fieldDef2}}

	var repField1 tagValue
	var repField2 tagValue
	repField1.init(Tag(2), []byte("a"))
	repField2.init(Tag(3), []byte("a"))

	var groupID tagValue
	groupID.init(Tag(1), []byte("1"))

	var otherField tagValue
	otherField.init(Tag(500), []byte("blah"))

	var groupID2 tagValue
	groupID2.init(Tag(1), []byte("2"))

	var groupID3 tagValue
	groupID3.init(Tag(1), []byte("3"))

	var tests = []struct {
		fields               tagValues
		expectedRemFields    int
		expectReject         bool
		expectedRejectReason int
	}{
		//non-repeating
		{fields: tagValues{groupID, repField1},
			expectedRemFields: 0},
		{fields: tagValues{groupID, repField1, repField2},
			expectedRemFields: 0},
		//test with trailing tag not in group
		{fields: tagValues{groupID, repField1, repField2, otherField},
			expectedRemFields: 1},
		//repeats
		{fields: tagValues{groupID2, repField1, repField2, repField1, repField2, otherField},
			expectedRemFields: 1},
		//REJECT: group size declared > actual group size
		{expectReject: true,
			fields:               tagValues{groupID3, repField1, repField2, repField1, repField2, otherField},
			expectedRejectReason: rejectReasonIncorrectNumInGroupCountForRepeatingGroup,
		},
		{expectReject: true,
			fields:               tagValues{groupID3, repField1, repField1, otherField},
			expectedRejectReason: rejectReasonIncorrectNumInGroupCountForRepeatingGroup,
		},
		//REJECT: group size declared < actual group size
		{expectReject: true,
			fields:               tagValues{groupID, repField1, repField2, repField1, repField2, otherField},
			expectedRejectReason: rejectReasonIncorrectNumInGroupCountForRepeatingGroup,
		},
	}

	for _, test := range tests {
		remFields, reject := validateVisitGroupField(groupFieldDef, test.fields)

		if test.expectReject {
			if reject == nil {
				t.Error("Expected Reject")
			}

			if reject.RejectReason() != test.expectedRejectReason {
				t.Errorf("Expected reject reason %v got %v", test.expectedRejectReason, reject.RejectReason())
			}
			continue
		}

		if reject != nil {
			t.Errorf("Unexpected reject: %v", reject)
		}

		if len(remFields) != test.expectedRemFields {
			t.Errorf("Expected len %v got %v", test.expectedRemFields, len(remFields))
		}
	}
}
