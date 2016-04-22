package quickfix

import (
	"fmt"
)

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
	RefTagID() *Tag
	IsBusinessReject() bool
}

type messageRejectError struct {
	rejectReason     int
	text             string
	refTagID         *Tag
	isBusinessReject bool
}

func (e messageRejectError) Error() string          { return e.text }
func (e messageRejectError) RefTagID() *Tag         { return e.refTagID }
func (e messageRejectError) RejectReason() int      { return e.rejectReason }
func (e messageRejectError) IsBusinessReject() bool { return e.isBusinessReject }

//NewMessageRejectError returns a MessageRejectError with the given error message, reject reason, and optional reftagid
func NewMessageRejectError(err string, rejectReason int, refTagID *Tag) MessageRejectError {
	return messageRejectError{text: err, rejectReason: rejectReason, refTagID: refTagID}
}

//NewBusinessMessageRejectError returns a MessageRejectError with the given error mesage, reject reason, and optional reftagid.
//Reject is treated as a business level reject
func NewBusinessMessageRejectError(err string, rejectReason int, refTagID *Tag) MessageRejectError {
	return messageRejectError{text: err, rejectReason: rejectReason, refTagID: refTagID, isBusinessReject: true}
}

//incorrectDataFormatForValue returns an error indicating a field that cannot be parsed as the type required.
func incorrectDataFormatForValue(tag Tag) MessageRejectError {
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
	return NewBusinessMessageRejectError("Conditionally required field missing", rejectReasonConditionallyRequiredFieldMissing, &tag)
}

//valueIsIncorrectNoTag returns an error indicating a field with value that is not valid.
//FIXME: to be compliant with legacy tests, for certain value issues, do not include reftag? (11c_NewSeqNoLess)
func valueIsIncorrectNoTag() MessageRejectError {
	return NewMessageRejectError("Value is incorrect (out of range) for this tag", rejectReasonValueIsIncorrect, nil)
}

//invalidMessageType returns an error to indicate an invalid message type
func invalidMessageType() MessageRejectError {
	return NewMessageRejectError("Invalid MsgType", rejectReasonInvalidMsgType, nil)
}

//unsupportedMessageType returns an error to indicate an unhandled message.
func unsupportedMessageType() MessageRejectError {
	return NewBusinessMessageRejectError("Unsupported Message Type", rejectReasonUnsupportedMessageType, nil)
}

//tagNotDefinedForThisMessageType returns an error for an invalid tag appearing in a message.
func tagNotDefinedForThisMessageType(tag Tag) MessageRejectError {
	return NewMessageRejectError("Tag not defined for this message type", rejectReasonTagNotDefinedForThisMessageType, &tag)
}

//tagAppearsMoreThanOnce return an error for multiple tags in a message not detected as a repeating group.
func tagAppearsMoreThanOnce(tag Tag) MessageRejectError {
	return NewMessageRejectError("Tag appears more than once", rejectReasonTagAppearsMoreThanOnce, &tag)
}

//requiredTagMissing returns a validation error when a required field cannot be found in a message.
func requiredTagMissing(tag Tag) MessageRejectError {
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

//tagSpecifiedWithoutAValue returns a validation error for when a field has no value.
func tagSpecifiedWithoutAValue(tag Tag) MessageRejectError {
	return NewMessageRejectError("Tag specified without a value", rejectReasonTagSpecifiedWithoutAValue, &tag)
}

//invalidTagNumber returns a validation error for messages with invalid tags.
func invalidTagNumber(tag Tag) MessageRejectError {
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

//requiredConfigurationMissing indicates a missing required conditional configuration option.
func requiredConfigurationMissing(setting string) error {
	return fmt.Errorf("missing configuration: %v", setting)
}
