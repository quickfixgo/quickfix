package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
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

//NewInvalidTagNumber creates a reject for tag in message.
func NewInvalidTagNumber(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Invalid tag number", rejectReason: InvalidTagNumber,
		refTagID: &tag}
}

//NewTagNotDefinedForThisMessageType creates a reject for tag in message.
func NewTagNotDefinedForThisMessageType(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag not defined for this message type", rejectReason: TagNotDefinedForThisMessageType,
		refTagID: &tag}
}

//NewRequiredTagMissing creates a reject for tag in message.
func NewRequiredTagMissing(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Required tag missing", rejectReason: RequiredTagMissing,
		refTagID: &tag}
}

//NewTagSpecifiedWithoutAValue creates a reject for tag in message.
func NewTagSpecifiedWithoutAValue(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag specified without a value", rejectReason: TagSpecifiedWithoutAValue,
		refTagID: &tag}
}

//NewValueIsIncorrect creates reject for tag in message.
//FIXME: to be compliant with legacy tests, for certain value issues, do not include reftag? (11c_NewSeqNoLess)
func NewValueIsIncorrect(msg Message, tag *tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Value is incorrect (out of range) for this tag", rejectReason: ValueIsIncorrect, refTagID: tag}
}

//NewIncorrectDataFormatForValue creates reject for tag in message.
func NewIncorrectDataFormatForValue(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Incorrect data format for value", rejectReason: IncorrectDataFormatForValue, refTagID: &tag}
}

//NewTagSpecifiedOutOfRequiredOrder creates reject for tag in message.
func NewTagSpecifiedOutOfRequiredOrder(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag specified out of required order", rejectReason: TagSpecifiedOutOfRequiredOrder, refTagID: &tag}
}

//NewTagAppearsMoreThanOnce creates reject for tag in message.
func NewTagAppearsMoreThanOnce(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag appears more than once", rejectReason: TagAppearsMoreThanOnce, refTagID: &tag}
}

//NewIncorrectNumInGroupCountForRepeatingGroup creates reject for tag in message.
func NewIncorrectNumInGroupCountForRepeatingGroup(msg Message, tag tag.Tag) MessageReject {
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

//IncorrectBeginString is a message reject specific to incorrect begin strings.
type IncorrectBeginString struct{ MessageReject }

//NewIncorrectBeginString creates an IncorrectBeginString reject for msg.
func NewIncorrectBeginString(msg Message) IncorrectBeginString {
	return IncorrectBeginString{messageRejectBase{rejectedMessage: msg, text: "Incorrect BeginString"}}
}

//NewCompIDProblem creates a reject for msg where msg has invalid comp id values.
func NewCompIDProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "CompID problem", rejectReason: CompIDProblem}
}

//NewSendingTimeAccuracyProblem creates a reject for a msg with stale or invalid sending time.
func NewSendingTimeAccuracyProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "SendingTime accuracy problem", rejectReason: SendingTimeAccuracyProblem}
}

//TargetTooHigh is a MessageReject where the sequence number is larger than expected.
type TargetTooHigh struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

//TargetTooLow is a MessageReject where the sequence number is less than expected.
type TargetTooLow struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

//NewTargetTooHigh creates a TargetTooHigh Instance with the specified expected and actual target numbers.
func NewTargetTooHigh(msg Message, receivedTarget, expectedTarget int) TargetTooHigh {
	return TargetTooHigh{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}

//NewTargetTooLow creates a TargetTooLow Instance with the specified expected and actual target numbers.
func NewTargetTooLow(msg Message, receivedTarget, expectedTarget int) TargetTooLow {
	return TargetTooLow{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too low, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}
