package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/tag"
)

//RejectReason (tag 373) enum values.
const (
	InvalidTagNumber                          RejectReason = 0
	RequiredTagMissing                        RejectReason = 1
	TagNotDefinedForThisMessageType           RejectReason = 2
	UnsupportedMessageType                    RejectReason = 3
	TagSpecifiedWithoutAValue                 RejectReason = 4
	ValueIsIncorrect                          RejectReason = 5
	IncorrectDataFormatForValue               RejectReason = 6
	CompIDProblem                             RejectReason = 9
	SendingTimeAccuracyProblem                RejectReason = 10
	InvalidMsgType                            RejectReason = 11
	TagAppearsMoreThanOnce                    RejectReason = 13
	TagSpecifiedOutOfRequiredOrder            RejectReason = 14
	IncorrectNumInGroupCountForRepeatingGroup RejectReason = 16
)

//newInvalidTagNumber creates a reject for tag in message.
func newInvalidTagNumber(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Invalid tag number", rejectReason: InvalidTagNumber,
		refTagID: &tag}
}

//newTagNotDefinedForThisMessageType creates a reject for tag in message.
func newTagNotDefinedForThisMessageType(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag not defined for this message type", rejectReason: TagNotDefinedForThisMessageType,
		refTagID: &tag}
}

//newRequiredTagMissing creates a reject for tag in message.
func newRequiredTagMissing(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Required tag missing", rejectReason: RequiredTagMissing,
		refTagID: &tag}
}

//newTagSpecifiedWithoutAValue creates a reject for tag in message.
func newTagSpecifiedWithoutAValue(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag specified without a value", rejectReason: TagSpecifiedWithoutAValue,
		refTagID: &tag}
}

//newValueIsIncorrect creates reject for tag in message.
//FIXME: to be compliant with legacy tests, for certain value issues, do not include reftag? (11c_NewSeqNoLess)
func newValueIsIncorrect(msg Message, tag *tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Value is incorrect (out of range) for this tag", rejectReason: ValueIsIncorrect, refTagID: tag}
}

//newIncorrectDataFormatForValue creates reject for tag in message.
func newIncorrectDataFormatForValue(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Incorrect data format for value", rejectReason: IncorrectDataFormatForValue, refTagID: &tag}
}

//newTagSpecifiedOutOfRequiredOrder creates reject for tag in message.
func newTagSpecifiedOutOfRequiredOrder(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag specified out of required order", rejectReason: TagSpecifiedOutOfRequiredOrder, refTagID: &tag}
}

//newTagAppearsMoreThanOnce creates reject for tag in message.
func newTagAppearsMoreThanOnce(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag appears more than once", rejectReason: TagAppearsMoreThanOnce, refTagID: &tag}
}

//newIncorrectNumInGroupCountForRepeatingGroup creates reject for tag in message.
func newIncorrectNumInGroupCountForRepeatingGroup(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Incorrect NumInGroup count for repeating group", rejectReason: IncorrectNumInGroupCountForRepeatingGroup, refTagID: &tag}
}

//NewInvalidMessageType creates reject for message of invalid type.
func NewInvalidMessageType(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Invalid MsgType", rejectReason: InvalidMsgType}
}

//NewUnsupportedMessageType creates reject for message of unsupported type.
func NewUnsupportedMessageType(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Unsupported Message Type", rejectReason: UnsupportedMessageType, businessReject: true}
}

//incorrectBeginString is a message reject specific to incorrect begin strings.
type incorrectBeginString struct{ MessageReject }

//newIncorrectBeginString creates an IncorrectBeginString reject for msg.
func newIncorrectBeginString(msg Message) incorrectBeginString {
	return incorrectBeginString{messageRejectBase{rejectedMessage: msg, text: "Incorrect BeginString"}}
}

//newCompIDProblem creates a reject for msg where msg has invalid comp id values.
func newCompIDProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "CompID problem", rejectReason: CompIDProblem}
}

//newSendingTimeAccuracyProblem creates a reject for a msg with stale or invalid sending time.
func newSendingTimeAccuracyProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "SendingTime accuracy problem", rejectReason: SendingTimeAccuracyProblem}
}

//targetTooHigh is a MessageReject where the sequence number is larger than expected.
type targetTooHigh struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

//targetTooLow is a MessageReject where the sequence number is less than expected.
type targetTooLow struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

//newTargetTooHigh creates a TargetTooHigh Instance with the specified expected and actual target numbers.
func newTargetTooHigh(msg Message, receivedTarget, expectedTarget int) targetTooHigh {
	return targetTooHigh{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}

//newTargetTooLow creates a TargetTooLow Instance with the specified expected and actual target numbers.
func newTargetTooLow(msg Message, receivedTarget, expectedTarget int) targetTooLow {
	return targetTooLow{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too low, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}
