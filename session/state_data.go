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

func (s * stateData) verifyIgnoreSeqNumTooHigh(msg message.Message) reject.MessageReject {
  return s.verifySelect(msg, false)
}

func (s * stateData) verifySelect(msg message.Message, checkTooHigh bool) reject.MessageReject {
  if err:=s.checkSendingTime(msg); err!=nil {return err}

  if checkTooHigh {
    if err:=s.checkTargetTooHigh(msg); err!=nil {return err}
  }

  return nil
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

func (s * stateData) checkSendingTime(msg message.Message) reject.MessageReject {
  if sendingTime,err:=msg.Header().UTCTimestampField(fix.SendingTime); err!=nil {
    return reject.NewRequiredTagMissing(msg, fix.SendingTime)
  } else {
    if delta:=time.Since(sendingTime.UTCTimestampValue());
      delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
      return reject.NewSendingTimeAccuracyProblem(msg)
    }
  }

  return nil
}
