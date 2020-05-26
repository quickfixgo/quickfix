package quickfix

import (
	"errors"
	"fmt"
)

//ErrDoNotSend is a convenience error to indicate a DoNotSend in ToApp
var ErrDoNotSend = errors.New("Do Not Send")

//rejectReason enum values.
const (
	rejectReasonInvalidTagNumber                          = 0
	rejectReasonRequiredTagMissing                        = 1
	rejectReasonTagNotDefinedForThisMessageType           = 2
	rejectReasonUnsupportedMessageType                    = 3
	rejectReasonTagSpecifiedWithoutAValue                 = 4
	rejectReasonValueIsIncorrect                          = 5
	rejectReasonConditionallyRequiredFieldMissing         = 5
	rejectReasonIncorrectDataFormatForValue               = 6
	rejectReasonCompIDProblem                             = 9
	rejectReasonSendingTimeAccuracyProblem                = 10
	rejectReasonInvalidMsgType                            = 11
	rejectReasonTagAppearsMoreThanOnce                    = 13
	rejectReasonTagSpecifiedOutOfRequiredOrder            = 14
	rejectReasonRepeatingGroupFieldsOutOfOrder            = 15
	rejectReasonIncorrectNumInGroupCountForRepeatingGroup = 16
)

//MessageRejectError is a type of error that can correlate to a message reject.
type MessageRejectError interface {
	error

	//RejectReason, tag 373 for session rejects, tag 380 for business rejects.
	RejectReason() int
	BusinessRejectRefID() string
	RefTagID() *Tag
	IsBusinessReject() bool
}

//RejectLogon indicates the application is rejecting permission to logon. Implements MessageRejectError
type RejectLogon struct {
	Text string
}

func (e RejectLogon) Error() string { return e.Text }

//RefTagID implements MessageRejectError
func (RejectLogon) RefTagID() *Tag { return nil }

//RejectReason implements MessageRejectError
func (RejectLogon) RejectReason() int { return 0 }

//BusinessRejectRefID implements MessageRejectError
func (RejectLogon) BusinessRejectRefID() string { return "" }

//IsBusinessReject implements MessageRejectError
func (RejectLogon) IsBusinessReject() bool { return false }

type messageRejectError struct {
	rejectReason        int
	text                string
	businessRejectRefID string
	refTagID            *Tag
	isBusinessReject    bool
}

func (e messageRejectError) Error() string               { return e.text }
func (e messageRejectError) RefTagID() *Tag              { return e.refTagID }
func (e messageRejectError) RejectReason() int           { return e.rejectReason }
func (e messageRejectError) BusinessRejectRefID() string { return e.businessRejectRefID }
func (e messageRejectError) IsBusinessReject() bool      { return e.isBusinessReject }

//NewMessageRejectError returns a MessageRejectError with the given error message, reject reason, and optional reftagid
func NewMessageRejectError(err string, rejectReason int, refTagID *Tag) MessageRejectError {
	return messageRejectError{text: err, rejectReason: rejectReason, refTagID: refTagID}
}

//NewBusinessMessageRejectError returns a MessageRejectError with the given error mesage, reject reason, and optional reftagid.
//Reject is treated as a business level reject
func NewBusinessMessageRejectError(err string, rejectReason int, refTagID *Tag) MessageRejectError {
	return messageRejectError{text: err, rejectReason: rejectReason, refTagID: refTagID, isBusinessReject: true}
}

//NewBusinessMessageRejectErrorWithRefID returns a MessageRejectError with the given error mesage, reject reason, refID, and optional reftagid.
//Reject is treated as a business level reject
func NewBusinessMessageRejectErrorWithRefID(err string, rejectReason int, businessRejectRefID string, refTagID *Tag) MessageRejectError {
	return messageRejectError{text: err, rejectReason: rejectReason, refTagID: refTagID, businessRejectRefID: businessRejectRefID, isBusinessReject: true}
}

//IncorrectDataFormatForValue returns an error indicating a field that cannot be parsed as the type required.
func IncorrectDataFormatForValue(tag Tag) MessageRejectError {
	return NewMessageRejectError("Incorrect data format for value", rejectReasonIncorrectDataFormatForValue, &tag)
}

//repeatingGroupFieldsOutOfOrder returns an error indicating a problem parsing repeating groups fields
func repeatingGroupFieldsOutOfOrder(tag Tag, reason string) MessageRejectError {
	if reason != "" {
		reason = fmt.Sprintf("Repeating group fields out of order (%s)", reason)
	} else {
		reason = "Repeating group fields out of order"
	}
	return NewMessageRejectError(reason, rejectReasonRepeatingGroupFieldsOutOfOrder, &tag)
}

//ValueIsIncorrect returns an error indicating a field with value that is not valid.
func ValueIsIncorrect(tag Tag) MessageRejectError {
	return NewMessageRejectError("Value is incorrect (out of range) for this tag", rejectReasonValueIsIncorrect, &tag)
}

//ConditionallyRequiredFieldMissing indicates that the requested field could not be found in the FIX message.
func ConditionallyRequiredFieldMissing(tag Tag) MessageRejectError {
	return NewBusinessMessageRejectError(fmt.Sprintf("Conditionally Required Field Missing (%d)", tag), rejectReasonConditionallyRequiredFieldMissing, &tag)
}

//valueIsIncorrectNoTag returns an error indicating a field with value that is not valid.
//FIXME: to be compliant with legacy tests, for certain value issues, do not include reftag? (11c_NewSeqNoLess)
func valueIsIncorrectNoTag() MessageRejectError {
	return NewMessageRejectError("Value is incorrect (out of range) for this tag", rejectReasonValueIsIncorrect, nil)
}

//InvalidMessageType returns an error to indicate an invalid message type
func InvalidMessageType() MessageRejectError {
	return NewMessageRejectError("Invalid MsgType", rejectReasonInvalidMsgType, nil)
}

//UnsupportedMessageType returns an error to indicate an unhandled message.
func UnsupportedMessageType() MessageRejectError {
	return NewBusinessMessageRejectError("Unsupported Message Type", rejectReasonUnsupportedMessageType, nil)
}

//TagNotDefinedForThisMessageType returns an error for an invalid tag appearing in a message.
func TagNotDefinedForThisMessageType(tag Tag) MessageRejectError {
	return NewMessageRejectError("Tag not defined for this message type", rejectReasonTagNotDefinedForThisMessageType, &tag)
}

//tagAppearsMoreThanOnce return an error for multiple tags in a message not detected as a repeating group.
func tagAppearsMoreThanOnce(tag Tag) MessageRejectError {
	return NewMessageRejectError("Tag appears more than once", rejectReasonTagAppearsMoreThanOnce, &tag)
}

//RequiredTagMissing returns a validation error when a required field cannot be found in a message.
func RequiredTagMissing(tag Tag) MessageRejectError {
	return NewMessageRejectError("Required tag missing", rejectReasonRequiredTagMissing, &tag)
}

//incorrectNumInGroupCountForRepeatingGroup returns a validation error when the num in group value for a group does not match actual group size.
func incorrectNumInGroupCountForRepeatingGroup(tag Tag) MessageRejectError {
	return NewMessageRejectError("Incorrect NumInGroup count for repeating group", rejectReasonIncorrectNumInGroupCountForRepeatingGroup, &tag)
}

//tagSpecifiedOutOfRequiredOrder returns validation error when the group order does not match the spec.
func tagSpecifiedOutOfRequiredOrder(tag Tag) MessageRejectError {
	return NewMessageRejectError("Tag specified out of required order", rejectReasonTagSpecifiedOutOfRequiredOrder, &tag)
}

//TagSpecifiedWithoutAValue returns a validation error for when a field has no value.
func TagSpecifiedWithoutAValue(tag Tag) MessageRejectError {
	return NewMessageRejectError("Tag specified without a value", rejectReasonTagSpecifiedWithoutAValue, &tag)
}

//InvalidTagNumber returns a validation error for messages with invalid tags.
func InvalidTagNumber(tag Tag) MessageRejectError {
	return NewMessageRejectError("Invalid tag number", rejectReasonInvalidTagNumber, &tag)
}

//compIDProblem creates a reject for msg where msg has invalid comp id values.
func compIDProblem() MessageRejectError {
	return NewMessageRejectError("CompID problem", rejectReasonCompIDProblem, nil)
}

//sendingTimeAccuracyProblem creates a reject for a msg with stale or invalid sending time.
func sendingTimeAccuracyProblem() MessageRejectError {
	return NewMessageRejectError("SendingTime accuracy problem", rejectReasonSendingTimeAccuracyProblem, nil)
}
