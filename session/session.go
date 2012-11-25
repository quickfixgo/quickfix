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
  ID
  MessageOut chan message.Buffer
  sessionEvent chan event
  log.Log
  Callback
  state
  expectedSeqNum int
  seqNum int
  heartBtInt time.Duration
  heartBeatTimeout time.Duration
  stateTimer eventTimer
}

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
  session.Log=logFactory.CreateSessionLog(session.ID.String())
  session.Callback=callback
  session.state=logonState{}
  session.sessionEvent=make(chan event)
  session.stateTimer=eventTimer{Task:func(){session.sessionEvent<-needHeartbeat}}

  callback.OnCreate(session.ID)
  sessions.newSession<-session

  return nil
}

func (s * session) onDisconnect() {
  s.Log.OnEvent("Disconnected")
}

func (s * session) fixMsgIn(msgBytes []byte) (disconnect bool) {
  s.Log.OnIncoming(string(msgBytes))
  msg,err:=basic.MessageFromParsedBytes(msgBytes)

  if err == nil {
    s.state = s.state.OnFixMsgIn(s, msg)
  } else {
    s.Log.OnEventf("Msg Parse Error: %s, %s", err.Error(), msgBytes)
  }

  switch s.state.(type) {
    case latentState:
      return true
  }

  return false
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
  s.Log.OnOutgoing(string(buf.Bytes()))
  s.MessageOut <-buf

  s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))

  s.seqNum++
}
