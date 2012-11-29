//fix session specific package 
package session

import(
    "time"
    "quickfixgo/fix"
    "quickfixgo/log"
    "quickfixgo/settings"
    "quickfixgo/message"
    "quickfixgo/message/basic"
    )

type session struct {
  stateData

  log log.Log
  ID

  MessageOut chan message.Buffer
  SessionEvent chan event
  callback Callback
  state
  stateTimer eventTimer
}

//Creates Session, associates with internal session registry
func Create(dict settings.Dictionary, logFactory log.LogFactory, callback Callback) error {
  session:=new(session)

  if beginString,ok:=dict.GetString(settings.BeginString); ok {
    session.ID.BeginString=beginString
  } else {
    return settings.RequiredConfigurationMissing(settings.BeginString)
  }

  if targetCompID,ok:=dict.GetString(settings.TargetCompID); ok {
    session.ID.TargetCompID=targetCompID
  } else {
    return settings.RequiredConfigurationMissing(settings.TargetCompID)
  }

  if senderCompID,ok:=dict.GetString(settings.SenderCompID); ok {
    session.ID.SenderCompID=senderCompID
  } else {
    return settings.RequiredConfigurationMissing(settings.SenderCompID)
  }

  session.MessageOut=make(chan message.Buffer)
  session.SessionEvent=make(chan event)
  session.log=logFactory.CreateSessionLog(session.ID.String())
  session.callback=callback
  session.state=logonState{}
  session.stateTimer=eventTimer{Task:func(){session.SessionEvent<-needHeartbeat}}

  callback.OnCreate(session.ID)
  sessions.newSession<-session

  return nil
}

func (s * session) onDisconnect() {
  s.log.OnEvent("Disconnected")
}

func (s * session) FixMsgIn(msgBytes []byte) (disconnect bool) {
  s.log.OnIncoming(string(msgBytes))
  msg,err:=basic.MessageFromParsedBytes(msgBytes)

  if err == nil {
    s.state = s.state.OnFixMsgIn(s, msg)
  } else {
    s.log.OnEventf("Msg Parse Error: %s, %s", err.Error(), msgBytes)
  }

  switch s.state.(type) {
    case latentState:
      return true
  }

  return false
}

func (s * session) OnTimeout(e event) (disconnect bool) {

  return
}

func (s * session) generateLogout() {
  s.generateLogoutWithReason("")
}

func (s *session) generateLogoutWithReason(reason string) {
  reply:=basic.NewMessage()
  reply.MsgHeader.Set(basic.NewStringField(fix.MsgType, "5"))
  reply.MsgHeader.Set(basic.NewStringField(fix.BeginString, s.BeginString))
  reply.MsgHeader.Set(basic.NewStringField(fix.TargetCompID, s.TargetCompID))
  reply.MsgHeader.Set(basic.NewStringField(fix.SenderCompID, s.SenderCompID))

  if reason !="" {
    reply.MsgBody.Set(basic.NewStringField(fix.Text, reason))
  }

  s.send(reply)
  s.log.OnEvent("Sending logout response")
}

func (s * session) send(builder message.Builder) {
  sendingTime:=time.Now().UTC()

  //FIXME
  builder.SetHeaderField(basic.NewIntField(fix.MsgSeqNum, s.seqNum))
  builder.SetHeaderField(basic.NewStringField(fix.BeginString, s.BeginString))
  builder.SetHeaderField(basic.NewStringField(fix.SenderCompID, s.SenderCompID))
  builder.SetHeaderField(basic.NewStringField(fix.TargetCompID, s.TargetCompID))
  if s.BeginString >= "FIX.4.2" {
    builder.SetHeaderField(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
  } else {
    builder.SetHeaderField(basic.NewUTCTimestampFieldNoMillis(fix.SendingTime, sendingTime))
  }

  //FIXME -log in message out receiver
  buf:=builder.Build()
  s.log.OnOutgoing(string(buf.Bytes()))
  s.MessageOut <-buf

  s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))

  s.seqNum++
}
