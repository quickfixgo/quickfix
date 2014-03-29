package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
)

const (
	InvalidTagNumber                RejectReason = 0
	RequiredTagMissing              RejectReason = 1
	TagNotDefinedForThisMessageType RejectReason = 2
	UnsupportedMessageType          RejectReason = 3
	TagSpecifiedWithoutAValue       RejectReason = 4
	ValueIsIncorrect                RejectReason = 5
	IncorrectDataFormatForValue     RejectReason = 6
	CompIDProblem                   RejectReason = 9
	SendingTimeAccuracyProblem      RejectReason = 10
	InvalidMsgType                  RejectReason = 11
)

func NewInvalidTagNumber(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Invalid tag number", rejectReason: InvalidTagNumber,
		refTagID: &tag}
}

func NewTagNotDefinedForThisMessageType(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag not defined for this message type", rejectReason: TagNotDefinedForThisMessageType,
		refTagID: &tag}
}

func NewRequiredTagMissing(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Required tag missing", rejectReason: RequiredTagMissing,
		refTagID: &tag}
}

func NewTagSpecifiedWithoutAValue(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Tag specified without a value", rejectReason: TagSpecifiedWithoutAValue,
		refTagID: &tag}
}

//FIXME: to be compliant with legacy tests, for certain value issues, do not include reftag? (11c_NewSeqNoLess)
func NewValueIsIncorrect(msg Message, tag *tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Value is incorrect (out of range) for this tag", rejectReason: ValueIsIncorrect, refTagID: tag}
}

func NewIncorrectDataFormatForValue(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Incorrect data format for value", rejectReason: IncorrectDataFormatForValue, refTagID: &tag}
}

func NewInvalidMessageType(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Invalid MsgType", rejectReason: InvalidMsgType}
}

func NewUnsupportedMessageType(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Unsupported Message Type", rejectReason: UnsupportedMessageType, businessReject: true}
}

type IncorrectBeginString struct{ MessageReject }

func NewIncorrectBeginString(msg Message) IncorrectBeginString {
	return IncorrectBeginString{messageRejectBase{rejectedMessage: msg, text: "Incorrect BeginString"}}
}

func NewCompIDProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "CompID problem", rejectReason: CompIDProblem}
}

func NewSendingTimeAccuracyProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "SendingTime accuracy problem", rejectReason: SendingTimeAccuracyProblem}
}

type TargetTooHigh struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

type TargetTooLow struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

func NewTargetTooHigh(msg Message, receivedTarget, expectedTarget int) TargetTooHigh {
	return TargetTooHigh{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}

func NewTargetTooLow(msg Message, receivedTarget, expectedTarget int) TargetTooLow {
	return TargetTooLow{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too low, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}
