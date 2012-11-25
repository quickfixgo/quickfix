//fix session specific package 
package session

import(
    "quickfixgo/log"
    "quickfixgo/settings"
    "quickfixgo/message"
    "quickfixgo/message/basic"
    )

type Session struct {
  ID
  MessageOut chan message.Buffer
  log.Log
}

func Create(dict settings.Dictionary, logFactory log.LogFactory) error {
  session:=new(Session)

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

  Sessions.newSession<-session

  return nil
}

func (s * Session) onDisconnect() {
  s.Log.OnEvent("Disconnected")
}

func (s * Session) fixMsgIn(msgBytes []byte) (disconnect bool) {
  s.Log.OnIncoming(string(msgBytes))
  _,err:=basic.MessageFromParsedBytes(msgBytes)

  if err == nil {

  } else {
    s.Log.OnEventf("Msg Parse Error: %s, %s", err.Error(), msgBytes)
  }

  return true
}
