//Package errors contain error types generated and consumed by quickFIX components.
package errors

import (
	"fmt"
	"github.com/quickfixgo/quickfix/fix"
)

//RejectReason (tag 373).
type RejectReason int

//RejectReason enum values.
const (
	RejectReasonInvalidTagNumber                          RejectReason = 0
	RejectReasonRequiredTagMissing                        RejectReason = 1
	RejectReasonTagNotDefinedForThisMessageType           RejectReason = 2
	RejectReasonUnsupportedMessageType                    RejectReason = 3
	RejectReasonTagSpecifiedWithoutAValue                 RejectReason = 4
	RejectReasonValueIsIncorrect                          RejectReason = 5
	RejectReasonConditionallyRequiredFieldMissing         RejectReason = 5
	RejectReasonIncorrectDataFormatForValue               RejectReason = 6
	RejectReasonCompIDProblem                             RejectReason = 9
	RejectReasonSendingTimeAccuracyProblem                RejectReason = 10
	RejectReasonInvalidMsgType                            RejectReason = 11
	RejectReasonTagAppearsMoreThanOnce                    RejectReason = 13
	RejectReasonTagSpecifiedOutOfRequiredOrder            RejectReason = 14
	RejectReasonIncorrectNumInGroupCountForRepeatingGroup RejectReason = 16
)

//MessageRejectError is a type of error that can correlate to a message reject.
type MessageRejectError interface {
	error
	RejectReason() RejectReason
	RefTagID() *fix.Tag
	IsBusinessReject() bool
}

type messageRejectError struct {
	rejectReason     RejectReason
	text             string
	refTagID         *fix.Tag
	isBusinessReject bool
}

func (e messageRejectError) Error() string              { return e.text }
func (e messageRejectError) RefTagID() *fix.Tag         { return e.refTagID }
func (e messageRejectError) RejectReason() RejectReason { return e.rejectReason }
func (e messageRejectError) IsBusinessReject() bool     { return e.isBusinessReject }

//IncorrectDataFormatForValue returns an error indicating a field that cannot be parsed as the type required.
func IncorrectDataFormatForValue(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Incorrect data format for value", rejectReason: RejectReasonIncorrectDataFormatForValue, refTagID: &tag}
}

//ValueIsIncorrect returns an error indicating a field with value that is not valid.
func ValueIsIncorrect(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Value is incorrect (out of range) for this tag", rejectReason: RejectReasonValueIsIncorrect, refTagID: &tag}
}

//ValueIsIncorrectNoTag returns an error indicating a field with value that is not valid.
//FIXME: to be compliant with legacy tests, for certain value issues, do not include reftag? (11c_NewSeqNoLess)
func ValueIsIncorrectNoTag() MessageRejectError {
	return messageRejectError{text: "Value is incorrect (out of range) for this tag", rejectReason: RejectReasonValueIsIncorrect}
}

//InvalidMessageType returns an error to indicate an invalid message type
func InvalidMessageType() MessageRejectError {
	return messageRejectError{text: "Invalid MsgType", rejectReason: RejectReasonInvalidMsgType}
}

//UnsupportedMessageType returns an error to indicate an unhandled message.
func UnsupportedMessageType() MessageRejectError {
	return messageRejectError{text: "Unsupported Message Type", rejectReason: RejectReasonUnsupportedMessageType, isBusinessReject: true}
}

//TagNotDefinedForThisMessageType returns an error for an invalid tag appearing in a message.
func TagNotDefinedForThisMessageType(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Tag not defined for this message type", rejectReason: RejectReasonTagNotDefinedForThisMessageType, refTagID: &tag}
}

//TagAppearsMoreThanOnce return an error for multiple tags in a message not detected as a repeating group.
func TagAppearsMoreThanOnce(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Tag appears more than once", rejectReason: RejectReasonTagAppearsMoreThanOnce, refTagID: &tag}
}

//RequiredTagMissing returns a validation error when a required field cannot be found in a message.
func RequiredTagMissing(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Required tag missing", rejectReason: RejectReasonRequiredTagMissing, refTagID: &tag}
}

//IncorrectNumInGroupCountForRepeatingGroup returns a validation error when the num in group value for a group does not match actual group size.
func IncorrectNumInGroupCountForRepeatingGroup(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Incorrect NumInGroup count for repeating group", rejectReason: RejectReasonIncorrectNumInGroupCountForRepeatingGroup, refTagID: &tag}
}

//TagSpecifiedOutOfRequiredOrder returns validation error when the group order does not match the spec.
func TagSpecifiedOutOfRequiredOrder(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Tag specified out of required order", rejectReason: RejectReasonTagSpecifiedOutOfRequiredOrder, refTagID: &tag}
}

//TagSpecifiedWithoutAValue returns a validation error for when a field has no value.
func TagSpecifiedWithoutAValue(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Tag specified without a value", rejectReason: RejectReasonTagSpecifiedWithoutAValue, refTagID: &tag}
}

//InvalidTagNumber returns a validation error for messages with invalid tags.
func InvalidTagNumber(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Invalid tag number", rejectReason: RejectReasonInvalidTagNumber, refTagID: &tag}
}

//ConditionallyRequiredFieldMissing indicates that the requested field could not be found in the FIX message.
func ConditionallyRequiredFieldMissing(tag fix.Tag) MessageRejectError {
	return messageRejectError{text: "Conditionally required field missing", rejectReason: RejectReasonConditionallyRequiredFieldMissing, refTagID: &tag, isBusinessReject: true}
}

//CompIDProblem creates a reject for msg where msg has invalid comp id values.
func CompIDProblem() MessageRejectError {
	return messageRejectError{text: "CompID problem", rejectReason: RejectReasonCompIDProblem}
}

//SendingTimeAccuracyProblem creates a reject for a msg with stale or invalid sending time.
func SendingTimeAccuracyProblem() MessageRejectError {
	return messageRejectError{text: "SendingTime accuracy problem", rejectReason: RejectReasonSendingTimeAccuracyProblem}
}

//ParseError is returned when bytes cannot be parsed as a FIX message.
type ParseError struct {
	OrigError string
}

func (e ParseError) Error() string { return fmt.Sprintf("error parsing message: %s", e.OrigError) }
