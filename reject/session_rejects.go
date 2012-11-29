package reject

import(
    "fmt"
    "quickfixgo/message"
    )

const(
  RequiredTagMissing RejectReason = 1
    )

func NewRequiredTagMissing(msg message.Message, tag message.Tag) MessageReject {
  return messageRejectBase{rejectedMessage: msg, text: "Required tag missing", rejectReason:RequiredTagMissing,
    refTagID: tag}
}

type TargetTooHigh struct {
  MessageReject
  ReceivedTarget int
  ExpectedTarget int
}

func NewTargetTooHigh(msg message.Message, receivedTarget, expectedTarget int) TargetTooHigh {
  return TargetTooHigh{
    MessageReject: messageRejectBase{
      rejectedMessage: msg,
      text: fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", expectedTarget, receivedTarget)},
    ReceivedTarget: receivedTarget,
    ExpectedTarget: expectedTarget}
}


