package quickfix

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewMessageRejectError(t *testing.T) {
	var (
		expectedErrorString          = "Custom error"
		expectedRejectReason         = 5
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = false
	)
	msgRej := NewMessageRejectError(expectedErrorString, expectedRejectReason, &expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}

func TestNewBusinessMessageRejectError(t *testing.T) {
	var (
		expectedErrorString          = "Custom error"
		expectedRejectReason         = 5
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = true
	)
	msgRej := NewBusinessMessageRejectError(expectedErrorString, expectedRejectReason, &expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be true\n")
	}
}

func TestNewBusinessMessageRejectErrorWithRefID(t *testing.T) {
	var (
		expectedErrorString             = "Custom error"
		expectedRejectReason            = 5
		expectedbusinessRejectRefID     = "1"
		expectedRefTagID            Tag = 44
		expectedIsBusinessReject        = true
	)
	msgRej := NewBusinessMessageRejectErrorWithRefID(expectedErrorString, expectedRejectReason, expectedbusinessRejectRefID, &expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if strings.Compare(msgRej.BusinessRejectRefID(), expectedbusinessRejectRefID) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedbusinessRejectRefID, msgRej.BusinessRejectRefID())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be true\n")
	}
}

func TestIncorrectDataFormatForValue(t *testing.T) {
	var (
		expectedErrorString          = "Incorrect data format for value"
		expectedRejectReason         = 6
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = false
	)
	msgRej := IncorrectDataFormatForValue(expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}

func TestValueIsIncorrect(t *testing.T) {
	var (
		expectedErrorString          = "Value is incorrect (out of range) for this tag"
		expectedRejectReason         = 5
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = false
	)
	msgRej := ValueIsIncorrect(expectedRefTagID)
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}

func TestConditionallyRequiredFieldMissing(t *testing.T) {
	var (
		expectedRejectReason         = 5
		expectedRefTagID         Tag = 44
		expectedErrorString          = fmt.Sprintf("Conditionally Required Field Missing (%d)", expectedRefTagID)
		expectedIsBusinessReject     = true
	)
	msgRej := ConditionallyRequiredFieldMissing(expectedRefTagID)
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be true\n")
	}
}

func TestInvalidMessageType(t *testing.T) {
	var (
		expectedErrorString      = "Invalid MsgType"
		expectedRejectReason     = 11
		expectedIsBusinessReject = false
	)
	msgRej := InvalidMessageType()

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if msgRej.RefTagID() != nil {
		t.Errorf("expected: nil, got: %d\n", msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}

func TestUnsupportedMessageType(t *testing.T) {
	var (
		expectedRejectReason     = 3
		expectedErrorString      = "Unsupported Message Type"
		expectedIsBusinessReject = true
	)
	msgRej := UnsupportedMessageType()
	if msgRej.RefTagID() != nil {
		t.Errorf("expected: nil, got: %d\n", msgRej.RefTagID())
	}
	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be true\n")
	}
}

func TestTagNotDefinedForThisMessageType(t *testing.T) {
	var (
		expectedErrorString          = "Tag not defined for this message type"
		expectedRejectReason         = 2
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = false
	)
	msgRej := TagNotDefinedForThisMessageType(expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}

func TestRequiredTagMissing(t *testing.T) {
	var (
		expectedErrorString          = "Required tag missing"
		expectedRejectReason         = 1
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = false
	)
	msgRej := RequiredTagMissing(expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}

func TestTagSpecifiedWithoutAValue(t *testing.T) {
	var (
		expectedErrorString          = "Tag specified without a value"
		expectedRejectReason         = 4
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = false
	)
	msgRej := TagSpecifiedWithoutAValue(expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}

func TestInvalidTagNumber(t *testing.T) {
	var (
		expectedErrorString          = "Invalid tag number"
		expectedRejectReason         = 0
		expectedRefTagID         Tag = 44
		expectedIsBusinessReject     = false
	)
	msgRej := InvalidTagNumber(expectedRefTagID)

	if strings.Compare(msgRej.Error(), expectedErrorString) != 0 {
		t.Errorf("expected: %s, got: %s\n", expectedErrorString, msgRej.Error())
	}
	if msgRej.RejectReason() != expectedRejectReason {
		t.Errorf("expected: %d, got: %d\n", expectedRejectReason, msgRej.RejectReason())
	}
	if *msgRej.RefTagID() != expectedRefTagID {
		t.Errorf("expected: %d, got: %d\n", expectedRefTagID, msgRej.RefTagID())
	}
	if msgRej.IsBusinessReject() != expectedIsBusinessReject {
		t.Error("Expected IsBusinessReject to be false\n")
	}
}
