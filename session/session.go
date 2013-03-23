//fix session specific package 
package session

import(
    "time"
    "quickfixgo/fix"
    "quickfixgo/log"
    "quickfixgo/reject"
    "quickfixgo/settings"
    "quickfixgo/message"
    "quickfixgo/message/basic"
    )

type session struct {
  stateData

  log log.Log
  ID

  messageOut chan message.Buffer
  toSend chan message.Builder

  sessionEvent chan event
  callback Callback
  currentState state
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

  if session.ID.BeginString == "FIXT.1.1" {
    if defaultApplVerID,ok:=dict.GetString(settings.DefaultApplVerID); ok {
      session.ID.DefaultApplVerID=defaultApplVerID
    } else {
      return settings.RequiredConfigurationMissing(settings.DefaultApplVerID)
    }
  }

  session.toSend=make(chan message.Builder)

  session.sessionEvent=make(chan event)
  session.log=logFactory.CreateSessionLog(session.ID.String())
  session.callback=callback
  session.stateTimer=eventTimer{Task:func(){session.sessionEvent<-needHeartbeat}}

  callback.OnCreate(session.ID)
  sessions.newSession<-session

  return nil
}

func (s * session) accept() (chan message.Buffer, error) {
  s.currentState=logonState{}
  s.messageOut=make(chan message.Buffer)

  return s.messageOut, nil
}

func (s * session) onDisconnect() {
  s.log.OnEvent("Disconnected")
}

func (s * session) initiateLogout(reason string) {
  s.generateLogoutWithReason(reason)
  time.AfterFunc(time.Duration(2)*time.Second, func() {s.sessionEvent<-logoutTimeout})
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
  s.messageOut <-buf

  s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
  s.seqNum++
}

func (s *session) DoTargetTooHigh(reject reject.TargetTooHigh) {
  resend:=basic.NewMessage()
  resend.MsgHeader.Set(basic.NewStringField(fix.MsgType, "2"))
  resend.MsgBody.Set(basic.NewIntField(fix.BeginSeqNo,reject.ExpectedTarget))

  var endSeqNum=0
  if s.ID.BeginString < "FIX.4.2" {
    endSeqNum=999999
  }
  resend.MsgBody.Set(basic.NewIntField(fix.EndSeqNo,endSeqNum))

  s.send(resend)
}

func (s * session) verify(msg message.Message) reject.MessageReject {
  return s.verifySelect(msg, true)
}

func (s * session) verifyIgnoreSeqNumTooHigh(msg message.Message) reject.MessageReject {
  return s.verifySelect(msg, false)
}

func (s * session) verifySelect(msg message.Message, checkTooHigh bool) reject.MessageReject {
  if err:=s.checkSendingTime(msg); err!=nil {return err}
  if err:=s.checkTargetTooLow(msg); err!=nil {return err}

  if checkTooHigh {
    if err:=s.checkTargetTooHigh(msg); err!=nil {return err}
  }

  return s.fromCallback(msg)
}

func (s *session) fromCallback(msg message.Message) (reject.MessageReject) {
  msgType, _:=msg.Header().StringField(fix.MsgType)
  switch(msgType.Value()) {
    case "0","A","1","2","3","4","5":
      return s.callback.FromAdmin(msg, s.ID)
  }
  return s.callback.FromApp(msg, s.ID)
}

func (s * session) checkTargetTooLow(msg message.Message) reject.MessageReject {
  switch seqNum,err:=msg.Header().IntField(fix.MsgSeqNum); {
    case err!=nil:
      return reject.NewRequiredTagMissing(msg, fix.MsgSeqNum)
    case seqNum.IntValue() < s.expectedSeqNum:
      return reject.NewTargetTooLow(msg, seqNum.IntValue(), s.expectedSeqNum)
  }

  return nil
}

func (s * session) checkTargetTooHigh(msg message.Message) reject.MessageReject {
  switch seqNum,err:=msg.Header().IntField(fix.MsgSeqNum); {
    case err!=nil:
      return reject.NewRequiredTagMissing(msg, fix.MsgSeqNum)
    case seqNum.IntValue() > s.expectedSeqNum:
      return reject.NewTargetTooHigh(msg, seqNum.IntValue(), s.expectedSeqNum)
  }

  return nil
}

func (s * session) checkSendingTime(msg message.Message) reject.MessageReject {
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

func (s * session) run(msgIn chan[]byte) {
  defer func() {
    close(s.messageOut)
  }()

  for {

    switch s.currentState.(type) {
      case latentState:
        return
    }

    select {
      case msgBytes,ok:=<-msgIn:
        if ok {
          s.log.OnIncoming(string(msgBytes))
          if msg,err:=basic.MessageFromParsedBytes(msgBytes); err!=nil {
            s.log.OnEventf("Msg Parse Error: %s, %s", err.Error(), msgBytes)
          } else {
            s.currentState=s.currentState.FixMsgIn(s,msg)
          }
        } else {
          s.onDisconnect()
          return
        }

      case msg:=<-s.toSend:
        s.send(msg)

      case evt:=<-s.sessionEvent:
        s.currentState=s.currentState.Timeout(s,evt)
    }
  }
}
