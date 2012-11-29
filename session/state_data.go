package session

import(
    "time"
    "quickfixgo/fix"
    "quickfixgo/reject"
    "quickfixgo/message"
    )

type stateData struct {
  expectedSeqNum int
  seqNum int
  heartBeatTimeout time.Duration
}

func (s * stateData) checkTargetTooHigh(msg message.Message) reject.MessageReject {
  switch seqNum,err:=msg.Header().IntField(fix.MsgSeqNum); {
    case err!=nil:
      return reject.NewRequiredTagMissing(msg, fix.MsgSeqNum)
    case seqNum.IntValue() > s.expectedSeqNum:
      return reject.NewTargetTooHigh(msg, seqNum.IntValue(), s.expectedSeqNum)
  }

  return nil
}
