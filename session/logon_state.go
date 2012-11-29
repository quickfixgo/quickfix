package session

import(
    "time"
    "quickfixgo/fix"
    "quickfixgo/message"
    "quickfixgo/message/basic"
    )

type logonState struct {}

func (s logonState) OnFixMsgIn(session *session, msg message.Message) (nextState state) {
  if msgType, err:=msg.Header().StringField(fix.MsgType); err==nil && msgType.Value() =="A" {
    if err=s.handleLogon(session, msg); err!=nil {
      session.log.OnEvent(err.Error())
      return latentState{}
    } else {
      return inSession{}
    }
  }

  session.log.OnEventf("Invalid Session State: Received Msg %v while waiting for Logon", msg)
  return latentState{}
}

func (s logonState) handleLogon(session *session, msg message.Message) error {
  //reset on logon
  session.expectedSeqNum=1
  session.seqNum=1

  reply:=basic.NewMessage()
  reply.MsgHeader.Set(basic.NewStringField(fix.MsgType, "A"))
  reply.MsgHeader.Set(basic.NewStringField(fix.BeginString, session.BeginString))
  reply.MsgHeader.Set(basic.NewStringField(fix.TargetCompID, session.TargetCompID))
  reply.MsgHeader.Set(basic.NewStringField(fix.SenderCompID, session.SenderCompID))
  reply.MsgBody.Set(basic.NewIntField(fix.EncryptMethod,0))

  if HeartBtInt,err:=msg.Body().IntField(fix.HeartBtInt); err==nil {
    session.heartBeatTimeout = time.Duration(HeartBtInt.IntValue()) * time.Second
    reply.MsgBody.Set(HeartBtInt)
  }

  if session.DefaultApplVerID!="" {
    reply.MsgTrailer.Set(basic.NewStringField(fix.DefaultApplVerID, session.DefaultApplVerID))
  }

  session.log.OnEvent("Received logon request")
  session.send(reply)
  session.log.OnEvent("Responding to logon request")

  session.callback.OnLogon(session.ID)

  return nil
}
