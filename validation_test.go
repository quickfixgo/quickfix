// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/quickfixgo/quickfix/datadictionary"
)

type validateTest struct {
	TestName             string
	Validator            Validator
	MessageBytes         []byte
	ExpectedRejectReason int
	ExpectedRefTagID     *Tag
	DoNotExpectReject    bool
}

func TestValidate(t *testing.T) {
	var tests = []validateTest{
		tcInvalidTagNumberHeader(),
		tcInvalidTagNumberHeaderFixT(),
		tcInvalidTagNumberBody(),
		tcInvalidTagNumberBodyFixT(),
		tcInvalidTagNumberTrailer(),
		tcInvalidTagNumberTrailerFixT(),
		tcTagSpecifiedWithoutAValue(),
		tcTagSpecifiedWithoutAValueFixT(),
		tcInvalidMsgType(),
		tcInvalidMsgTypeFixT(),
		tcValueIsIncorrect(),
		tcValueIsIncorrectFixT(),
		tcIncorrectDataFormatForValue(),
		tcIncorrectDataFormatForValueFixT(),
		tcTagSpecifiedOutOfRequiredOrderHeader(),
		tcTagSpecifiedOutOfRequiredOrderHeaderFixT(),
		tcTagSpecifiedOutOfRequiredOrderTrailer(),
		tcTagSpecifiedOutOfRequiredOrderTrailerFixT(),
		tcTagSpecifiedOutOfRequiredOrderDisabledHeader(),
		tcTagSpecifiedOutOfRequiredOrderDisabledHeaderFixT(),
		tcTagSpecifiedOutOfRequiredOrderDisabledTrailer(),
		tcTagSpecifiedOutOfRequiredOrderDisabledTrailerFixT(),
		tcTagAppearsMoreThanOnce(),
		tcTagAppearsMoreThanOnceFixT(),
		tcFloatValidation(),
		tcFloatValidationFixT(),
		tcTagNotDefinedForMessage(),
		tcTagNotDefinedForMessageFixT(),
		tcTagIsDefinedForMessage(),
		tcTagIsDefinedForMessageFixT(),
		tcFieldNotFoundBody(),
		tcFieldNotFoundBodyFixT(),
		tcFieldNotFoundHeader(),
		tcFieldNotFoundHeaderFixT(),
		tcInvalidTagCheckDisabled(),
		tcInvalidTagCheckDisabledFixT(),
		tcInvalidTagCheckEnabled(),
		tcInvalidTagCheckEnabledFixT(),
	}

	msg := NewMessage()
	for _, test := range tests {
		assert.Nil(t, ParseMessage(msg, bytes.NewBuffer(test.MessageBytes)))
		reject := test.Validator.Validate(msg)

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
		// OK, expected and actual ref tag not set.
		case reject.RefTagID() != nil && test.ExpectedRefTagID == nil:
			t.Errorf("%v: Unexpected RefTag '%v'", test.TestName, *reject.RefTagID())
		case reject.RefTagID() == nil && test.ExpectedRefTagID != nil:
			t.Errorf("%v: Expected RefTag '%v'", test.TestName, *test.ExpectedRefTagID)
		case *reject.RefTagID() == *test.ExpectedRefTagID:
			// OK, tags equal.
		default:
			t.Errorf("%v: Expected RefTag '%v' got '%v'", test.TestName, *test.ExpectedRefTagID, *reject.RefTagID())
		}
	}
}

func createFIX40NewOrderSingle() *Message {
	msg := NewMessage()
	msg.Header.SetField(tagMsgType, FIXString("D"))
	msg.Header.SetField(tagBeginString, FIXString("FIX.4.0"))
	msg.Header.SetField(tagBodyLength, FIXString("0"))
	msg.Header.SetField(tagSenderCompID, FIXString("0"))
	msg.Header.SetField(tagTargetCompID, FIXString("0"))
	msg.Header.SetField(tagMsgSeqNum, FIXString("0"))
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()})

	msg.Body.SetField(Tag(11), FIXString("A"))
	msg.Body.SetField(Tag(21), FIXString("1"))
	msg.Body.SetField(Tag(55), FIXString("A"))
	msg.Body.SetField(Tag(54), FIXString("1"))
	msg.Body.SetField(Tag(40), FIXString("1"))
	msg.Body.SetField(Tag(38), FIXInt(5))
	msg.Body.SetField(Tag(100), FIXString("0"))

	msg.Trailer.SetField(tagCheckSum, FIXString("000"))

	return msg
}

func createFIX43NewOrderSingle() *Message {
	msg := NewMessage()
	msg.Header.SetField(tagMsgType, FIXString("D"))
	msg.Header.SetField(tagBeginString, FIXString("FIX.4.3"))
	msg.Header.SetField(tagBodyLength, FIXString("0"))
	msg.Header.SetField(tagSenderCompID, FIXString("0"))
	msg.Header.SetField(tagTargetCompID, FIXString("0"))
	msg.Header.SetField(tagMsgSeqNum, FIXString("0"))
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()})

	msg.Body.SetField(Tag(11), FIXString("A"))
	msg.Body.SetField(Tag(21), FIXString("1"))
	msg.Body.SetField(Tag(55), FIXString("A"))
	msg.Body.SetField(Tag(54), FIXString("1"))
	msg.Body.SetField(Tag(38), FIXInt(5))
	msg.Body.SetField(Tag(40), FIXString("1"))
	msg.Body.SetField(Tag(60), FIXUTCTimestamp{Time: time.Now()})

	msg.Trailer.SetField(tagCheckSum, FIXString("000"))

	return msg
}

func createFIX50SP2NewOrderSingle() *Message {
	msg := NewMessage()
	msg.Header.SetField(tagMsgType, FIXString("D"))
	msg.Header.SetField(tagBeginString, FIXString("FIXT.1.1"))
	msg.Header.SetField(tagBodyLength, FIXString("0"))
	msg.Header.SetField(tagSenderCompID, FIXString("0"))
	msg.Header.SetField(tagTargetCompID, FIXString("0"))
	msg.Header.SetField(tagMsgSeqNum, FIXString("0"))
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()})

	msg.Body.SetField(Tag(11), FIXString("A"))
	msg.Body.SetField(Tag(21), FIXString("1"))
	msg.Body.SetField(Tag(55), FIXString("A"))
	msg.Body.SetField(Tag(54), FIXString("1"))
	msg.Body.SetField(Tag(40), FIXString("1"))
	msg.Body.SetField(Tag(38), FIXInt(5))
	msg.Body.SetField(Tag(60), FIXUTCTimestamp{Time: time.Now(), Precision: Micros})
	msg.Body.SetField(Tag(100), FIXString("0"))

	msg.Trailer.SetField(tagCheckSum, FIXString("000"))

	return msg
}

func tcInvalidTagNumberHeader() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	invalidHeaderFieldMessage := createFIX40NewOrderSingle()
	tag := Tag(9999)
	invalidHeaderFieldMessage.Header.SetField(tag, FIXString("hello"))
	msgBytes := invalidHeaderFieldMessage.build()

	return validateTest{
		TestName:             "Invalid Tag Number Header",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcInvalidTagNumberHeaderFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	invalidHeaderFieldMessage := createFIX50SP2NewOrderSingle()
	tag := Tag(9999)
	invalidHeaderFieldMessage.Header.SetField(tag, FIXString("hello"))
	msgBytes := invalidHeaderFieldMessage.build()

	return validateTest{
		TestName:             "Invalid Tag Number Header FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcInvalidTagNumberBody() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	invalidBodyFieldMessage := createFIX40NewOrderSingle()
	tag := Tag(9999)
	invalidBodyFieldMessage.Body.SetField(tag, FIXString("hello"))
	msgBytes := invalidBodyFieldMessage.build()

	return validateTest{
		TestName:             "Invalid Tag Number Body",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcInvalidTagNumberBodyFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	invalidBodyFieldMessage := createFIX50SP2NewOrderSingle()
	tag := Tag(9999)
	invalidBodyFieldMessage.Body.SetField(tag, FIXString("hello"))
	msgBytes := invalidBodyFieldMessage.build()

	return validateTest{
		TestName:             "Invalid Tag Number Body FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcInvalidTagNumberTrailer() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	invalidTrailerFieldMessage := createFIX40NewOrderSingle()
	tag := Tag(9999)
	invalidTrailerFieldMessage.Trailer.SetField(tag, FIXString("hello"))
	msgBytes := invalidTrailerFieldMessage.build()

	return validateTest{
		TestName:             "Invalid Tag Number Trailer",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcInvalidTagNumberTrailerFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	invalidTrailerFieldMessage := createFIX50SP2NewOrderSingle()
	tag := Tag(9999)
	invalidTrailerFieldMessage.Trailer.SetField(tag, FIXString("hello"))
	msgBytes := invalidTrailerFieldMessage.build()

	return validateTest{
		TestName:             "Invalid Tag Number Trailer FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidTagNumber,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagNotDefinedForMessage() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	invalidMsg := createFIX40NewOrderSingle()
	tag := Tag(41)
	invalidMsg.Body.SetField(tag, FIXString("hello"))
	msgBytes := invalidMsg.build()

	return validateTest{
		TestName:             "Tag Not Defined For Message",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagNotDefinedForThisMessageType,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagNotDefinedForMessageFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	invalidMsg := createFIX50SP2NewOrderSingle()
	tag := Tag(41)
	invalidMsg.Body.SetField(tag, FIXString("hello"))
	msgBytes := invalidMsg.build()

	return validateTest{
		TestName:             "Tag Not Defined For Message FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagNotDefinedForThisMessageType,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagIsDefinedForMessage() validateTest {
	// Compare to `tcTagIsNotDefinedForMessage`.
	dict, _ := datadictionary.Parse("spec/FIX43.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	validMsg := createFIX43NewOrderSingle()
	msgBytes := validMsg.build()

	return validateTest{
		TestName:          "TagIsDefinedForMessage",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcTagIsDefinedForMessageFixT() validateTest {
	// Compare to `tcTagIsNotDefinedForMessage`.
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	validMsg := createFIX50SP2NewOrderSingle()
	msgBytes := validMsg.build()

	return validateTest{
		TestName:          "TagIsDefinedForMessage FIXT",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcFieldNotFoundBody() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	invalidMsg1 := NewMessage()
	invalidMsg1.Header.SetField(tagMsgType, FIXString("D")).
		SetField(tagBeginString, FIXString("FIX.4.0")).
		SetField(tagBodyLength, FIXString("0")).
		SetField(tagSenderCompID, FIXString("0")).
		SetField(tagTargetCompID, FIXString("0")).
		SetField(tagMsgSeqNum, FIXString("0")).
		SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()})
	invalidMsg1.Trailer.SetField(tagCheckSum, FIXString("000"))

	invalidMsg1.Body.SetField(Tag(11), FIXString("A")).
		SetField(Tag(21), FIXString("A")).
		SetField(Tag(55), FIXString("A")).
		SetField(Tag(54), FIXString("A")).
		SetField(Tag(38), FIXString("A"))

	tag := Tag(40)
	// Ord type is required. invalidMsg1.Body.SetField(Tag(40), "A")).
	msgBytes := invalidMsg1.build()

	return validateTest{
		TestName:             "FieldNotFoundBody",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonRequiredTagMissing,
		ExpectedRefTagID:     &tag,
	}
}

func tcFieldNotFoundBodyFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	invalidMsg1 := NewMessage()
	invalidMsg1.Header.SetField(tagMsgType, FIXString("D")).
		SetField(tagBeginString, FIXString("FIXT.1.1")).
		SetField(tagBodyLength, FIXString("0")).
		SetField(tagSenderCompID, FIXString("0")).
		SetField(tagTargetCompID, FIXString("0")).
		SetField(tagMsgSeqNum, FIXString("0")).
		SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()})
	invalidMsg1.Trailer.SetField(tagCheckSum, FIXString("000"))

	invalidMsg1.Body.SetField(Tag(11), FIXString("A")).
		SetField(Tag(21), FIXString("A")).
		SetField(Tag(55), FIXString("A")).
		SetField(Tag(54), FIXString("A")).
		SetField(Tag(38), FIXString("A")).
		SetField(Tag(60), FIXUTCTimestamp{Time: time.Now()})

	tag := Tag(40)
	// Ord type is required. invalidMsg1.Body.SetField(Tag(40), "A")).
	msgBytes := invalidMsg1.build()

	return validateTest{
		TestName:             "FieldNotFoundBody FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonRequiredTagMissing,
		ExpectedRefTagID:     &tag,
	}
}

func tcFieldNotFoundHeader() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)

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

	// Sending time is required. invalidMsg2.Header.FieldMap.SetField(tag.SendingTime, "0")).
	tag := tagSendingTime
	msgBytes := invalidMsg2.build()

	return validateTest{
		TestName:             "FieldNotFoundHeader",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonRequiredTagMissing,
		ExpectedRefTagID:     &tag,
	}
}

func tcFieldNotFoundHeaderFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)

	invalidMsg2 := NewMessage()
	invalidMsg2.Trailer.SetField(tagCheckSum, FIXString("000"))
	invalidMsg2.Body.SetField(Tag(11), FIXString("A")).
		SetField(Tag(21), FIXString("A")).
		SetField(Tag(55), FIXString("A")).
		SetField(Tag(54), FIXString("A")).
		SetField(Tag(38), FIXString("A"))

	invalidMsg2.Header.SetField(tagMsgType, FIXString("D")).
		SetField(tagBeginString, FIXString("FIXT.1.1")).
		SetField(tagBodyLength, FIXString("0")).
		SetField(tagSenderCompID, FIXString("0")).
		SetField(tagTargetCompID, FIXString("0")).
		SetField(tagMsgSeqNum, FIXString("0"))

	// Sending time is required. invalidMsg2.Header.FieldMap.SetField(tag.SendingTime, "0")).
	tag := tagSendingTime
	msgBytes := invalidMsg2.build()

	return validateTest{
		TestName:             "FieldNotFoundHeader FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonRequiredTagMissing,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagSpecifiedWithoutAValue() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	builder := createFIX40NewOrderSingle()

	bogusTag := Tag(109)
	builder.Body.SetField(bogusTag, FIXString(""))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "Tag SpecifiedWithoutAValue",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedWithoutAValue,
		ExpectedRefTagID:     &bogusTag,
	}
}

func tcTagSpecifiedWithoutAValueFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	builder := createFIX50SP2NewOrderSingle()

	bogusTag := Tag(109)
	builder.Body.SetField(bogusTag, FIXString(""))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "Tag SpecifiedWithoutAValue FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedWithoutAValue,
		ExpectedRefTagID:     &bogusTag,
	}
}

func tcInvalidMsgType() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	builder := createFIX40NewOrderSingle()
	builder.Header.SetField(tagMsgType, FIXString("z"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "Invalid MsgType",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidMsgType,
	}
}

func tcInvalidMsgTypeFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	builder := createFIX50SP2NewOrderSingle()
	builder.Header.SetField(tagMsgType, FIXString("zz"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "Invalid MsgType FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonInvalidMsgType,
	}
}

func tcValueIsIncorrect() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)

	tag := Tag(21)
	builder := createFIX40NewOrderSingle()
	builder.Body.SetField(tag, FIXString("4"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "ValueIsIncorrect",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonValueIsIncorrect,
		ExpectedRefTagID:     &tag,
	}
}

func tcValueIsIncorrectFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)

	tag := Tag(21)
	builder := createFIX50SP2NewOrderSingle()
	builder.Body.SetField(tag, FIXString("4"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "ValueIsIncorrect FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonValueIsIncorrect,
		ExpectedRefTagID:     &tag,
	}
}

func tcIncorrectDataFormatForValue() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	builder := createFIX40NewOrderSingle()
	tag := Tag(38)
	builder.Body.SetField(tag, FIXString("+200.00"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "IncorrectDataFormatForValue",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonIncorrectDataFormatForValue,
		ExpectedRefTagID:     &tag,
	}
}

func tcIncorrectDataFormatForValueFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	builder := createFIX50SP2NewOrderSingle()
	tag := Tag(38)
	builder.Body.SetField(tag, FIXString("+200.00"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "IncorrectDataFormatForValue FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonIncorrectDataFormatForValue,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagSpecifiedOutOfRequiredOrderHeader() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)

	builder := createFIX40NewOrderSingle()
	tag := tagOnBehalfOfCompID
	// Should be in header.
	builder.Body.SetField(tag, FIXString("CWB"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "Tag specified out of required order in Header",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedOutOfRequiredOrder,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagSpecifiedOutOfRequiredOrderHeaderFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)

	builder := createFIX50SP2NewOrderSingle()
	tag := tagOnBehalfOfCompID
	// Should be in header.
	builder.Body.SetField(tag, FIXString("CWB"))
	msgBytes := builder.build()

	return validateTest{
		TestName:             "Tag specified out of required order in Header FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedOutOfRequiredOrder,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagSpecifiedOutOfRequiredOrderTrailer() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)

	builder := createFIX40NewOrderSingle()
	tag := tagSignature
	// Should be in trailer.
	builder.Body.SetField(tag, FIXString("SIG"))
	msgBytes := builder.build()

	refTag := Tag(100)
	return validateTest{
		TestName:             "Tag specified out of required order in Trailer",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedOutOfRequiredOrder,
		ExpectedRefTagID:     &refTag,
	}
}

func tcTagSpecifiedOutOfRequiredOrderTrailerFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)

	builder := createFIX50SP2NewOrderSingle()
	tag := tagSignature
	// Should be in trailer.
	builder.Body.SetField(tag, FIXString("SIG"))
	msgBytes := builder.build()

	refTag := Tag(100)
	return validateTest{
		TestName:             "Tag specified out of required order in Trailer FIXT",
		Validator:            validator,
		MessageBytes:         msgBytes,
		ExpectedRejectReason: rejectReasonTagSpecifiedOutOfRequiredOrder,
		ExpectedRefTagID:     &refTag,
	}
}

func tcInvalidTagCheckDisabled() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.RejectInvalidMessage = false
	validator := NewValidator(customValidatorSettings, dict, nil)

	builder := createFIX40NewOrderSingle()
	tag := Tag(9999)
	builder.Body.SetField(tag, FIXString("hello"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Invalid Tag Check - Disabled",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcInvalidTagCheckDisabledFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.RejectInvalidMessage = false
	validator := NewValidator(customValidatorSettings, appDict, tDict)

	builder := createFIX50SP2NewOrderSingle()
	tag := Tag(9999)
	builder.Body.SetField(tag, FIXString("hello"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Invalid Tag Check - Disabled FIXT",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcInvalidTagCheckEnabled() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.RejectInvalidMessage = true
	validator := NewValidator(customValidatorSettings, dict, nil)

	builder := createFIX40NewOrderSingle()
	tag := Tag(9999)
	builder.Body.SetField(tag, FIXString("hello"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Invalid Tag Check - Enabled",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: false,
		ExpectedRefTagID:  &tag,
	}
}

func tcInvalidTagCheckEnabledFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.RejectInvalidMessage = true
	validator := NewValidator(customValidatorSettings, appDict, tDict)

	builder := createFIX50SP2NewOrderSingle()
	tag := Tag(9999)
	builder.Body.SetField(tag, FIXString("hello"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Invalid Tag Check - Enabled FIXT",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: false,
		ExpectedRefTagID:  &tag,
	}
}

func tcTagSpecifiedOutOfRequiredOrderDisabledHeader() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.CheckFieldsOutOfOrder = false
	validator := NewValidator(customValidatorSettings, dict, nil)

	builder := createFIX40NewOrderSingle()
	tag := tagOnBehalfOfCompID
	// Should be in header.
	builder.Body.SetField(tag, FIXString("CWB"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Tag specified out of required order in Header - Disabled",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcTagSpecifiedOutOfRequiredOrderDisabledHeaderFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.CheckFieldsOutOfOrder = false
	validator := NewValidator(customValidatorSettings, appDict, tDict)

	builder := createFIX50SP2NewOrderSingle()
	tag := tagOnBehalfOfCompID
	// Should be in header.
	builder.Body.SetField(tag, FIXString("CWB"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Tag specified out of required order in Header - Disabled FIXT",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcTagSpecifiedOutOfRequiredOrderDisabledTrailer() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.CheckFieldsOutOfOrder = false
	validator := NewValidator(customValidatorSettings, dict, nil)

	builder := createFIX40NewOrderSingle()
	tag := tagSignature
	// Should be in trailer.
	builder.Body.SetField(tag, FIXString("SIG"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Tag specified out of required order in Trailer - Disabled",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcTagSpecifiedOutOfRequiredOrderDisabledTrailerFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	customValidatorSettings := defaultValidatorSettings
	customValidatorSettings.CheckFieldsOutOfOrder = false
	validator := NewValidator(customValidatorSettings, appDict, tDict)

	builder := createFIX50SP2NewOrderSingle()
	tag := tagSignature
	// Should be in trailer.
	builder.Body.SetField(tag, FIXString("SIG"))
	msgBytes := builder.build()

	return validateTest{
		TestName:          "Tag specified out of required order in Trailer - Disabled FIXT",
		Validator:         validator,
		MessageBytes:      msgBytes,
		DoNotExpectReject: true,
	}
}

func tcTagAppearsMoreThanOnce() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	tag := Tag(40)

	return validateTest{
		TestName:             "Tag appears more than once",
		Validator:            validator,
		MessageBytes:         []byte("8=FIX.4.09=10735=D34=249=TW52=20060102-15:04:0556=ISLD11=ID21=140=140=254=138=20055=INTC60=20060102-15:04:0510=234"),
		ExpectedRejectReason: rejectReasonTagAppearsMoreThanOnce,
		ExpectedRefTagID:     &tag,
	}
}

func tcTagAppearsMoreThanOnceFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	tag := Tag(40)

	return validateTest{
		TestName:             "Tag appears more than once FIXT",
		Validator:            validator,
		MessageBytes:         []byte("8=FIXT.1.19=10735=D34=249=TW52=20060102-15:04:0556=ISLD11=ID21=140=140=254=138=20055=INTC60=20060102-15:04:0510=234"),
		ExpectedRejectReason: rejectReasonTagAppearsMoreThanOnce,
		ExpectedRefTagID:     &tag,
	}
}

func tcFloatValidation() validateTest {
	dict, _ := datadictionary.Parse("spec/FIX42.xml")
	validator := NewValidator(defaultValidatorSettings, dict, nil)
	tag := Tag(38)
	return validateTest{
		TestName:             "FloatValidation",
		Validator:            validator,
		MessageBytes:         []byte("8=FIX.4.29=10635=D34=249=TW52=20140329-22:38:4556=ISLD11=ID21=140=154=138=+200.0055=INTC60=20140329-22:38:4510=178"),
		ExpectedRejectReason: rejectReasonIncorrectDataFormatForValue,
		ExpectedRefTagID:     &tag,
	}
}

func tcFloatValidationFixT() validateTest {
	tDict, _ := datadictionary.Parse("spec/FIXT11.xml")
	appDict, _ := datadictionary.Parse("spec/FIX50SP2.xml")
	validator := NewValidator(defaultValidatorSettings, appDict, tDict)
	tag := Tag(38)
	return validateTest{
		TestName:             "FloatValidation FIXT",
		Validator:            validator,
		MessageBytes:         []byte("8=FIXT.1.19=10635=D34=249=TW52=20140329-22:38:4556=ISLD11=ID21=140=154=138=+200.0055=INTC60=20140329-22:38:4510=178"),
		ExpectedRejectReason: rejectReasonIncorrectDataFormatForValue,
		ExpectedRefTagID:     &tag,
	}
}

func TestValidateVisitField(t *testing.T) {
	fieldType0 := datadictionary.NewFieldType("myfield", 11, "STRING")
	fieldDef0 := &datadictionary.FieldDef{FieldType: fieldType0}

	fieldType1 := datadictionary.NewFieldType("myfield", 2, "STRING")
	fieldDef1 := &datadictionary.FieldDef{FieldType: fieldType1, Fields: []*datadictionary.FieldDef{}}

	fieldType2 := datadictionary.NewFieldType("myfield", 3, "STRING")
	fieldDef2 := &datadictionary.FieldDef{FieldType: fieldType2, Fields: []*datadictionary.FieldDef{}}

	groupFieldType := datadictionary.NewFieldType("mygroupfield", 1, "INT")
	groupFieldDef := &datadictionary.FieldDef{FieldType: groupFieldType, Fields: []*datadictionary.FieldDef{fieldDef1, fieldDef2}}

	var field TagValue
	field.init(Tag(11), []byte("value"))

	var repField1 TagValue
	var repField2 TagValue
	repField1.init(Tag(2), []byte("a"))
	repField2.init(Tag(3), []byte("a"))

	var groupID TagValue
	groupID.init(Tag(1), []byte("1"))

	var groupID2 TagValue
	groupID2.init(Tag(1), []byte("2"))

	var groupID3 TagValue
	groupID3.init(Tag(1), []byte("3"))

	var tests = []struct {
		fieldDef             *datadictionary.FieldDef
		fields               []TagValue
		expectedRemFields    int
		expectReject         bool
		expectedRejectReason int
	}{
		// Non-repeating.
		{expectedRemFields: 0,
			fieldDef: fieldDef0,
			fields:   []TagValue{field}},
		// Single field group.
		{expectedRemFields: 0,
			fieldDef: groupFieldDef,
			fields:   []TagValue{groupID, repField1}},
		// Multiple field group.
		{expectedRemFields: 0,
			fieldDef: groupFieldDef,
			fields:   []TagValue{groupID, repField1, repField2}},
		// Test with trailing tag not in group.
		{expectedRemFields: 1,
			fieldDef: groupFieldDef,
			fields:   []TagValue{groupID, repField1, repField2, field}},
		// Repeats.
		{expectedRemFields: 1,
			fieldDef: groupFieldDef,
			fields:   []TagValue{groupID2, repField1, repField2, repField1, repField2, field}},
		// REJECT: group size declared > actual group size.
		{expectReject: true,
			fieldDef:             groupFieldDef,
			fields:               []TagValue{groupID3, repField1, repField2, repField1, repField2, field},
			expectedRejectReason: rejectReasonIncorrectNumInGroupCountForRepeatingGroup,
		},
		{expectReject: true,
			fieldDef:             groupFieldDef,
			fields:               []TagValue{groupID3, repField1, repField1, field},
			expectedRejectReason: rejectReasonIncorrectNumInGroupCountForRepeatingGroup,
		},
		// REJECT: group size declared < actual group size.
		{expectReject: true,
			fieldDef:             groupFieldDef,
			fields:               []TagValue{groupID, repField1, repField2, repField1, repField2, field},
			expectedRejectReason: rejectReasonIncorrectNumInGroupCountForRepeatingGroup,
		},
	}

	for _, test := range tests {
		remFields, reject := validateVisitField(test.fieldDef, test.fields)

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
