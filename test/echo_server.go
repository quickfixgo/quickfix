package main

import(
    "fmt"
    "os"
    "os/signal"
    "github.com/cbusbey/quickfixgo"
    "github.com/cbusbey/quickfixgo/fix"
    "github.com/cbusbey/quickfixgo/log"
    "github.com/cbusbey/quickfixgo/settings"
    "github.com/cbusbey/quickfixgo/reject"
    "github.com/cbusbey/quickfixgo/message"
    "github.com/cbusbey/quickfixgo/message/basic"
    "github.com/cbusbey/quickfixgo/session"
    )

type EchoApplication struct {

}

func (e EchoApplication) OnCreate(sessionID session.ID) {
  fmt.Printf("OnCreate %v\n", sessionID.String())
}
func (e EchoApplication) OnLogon(sessionID session.ID) {
  fmt.Printf("OnLogon %v\n", sessionID.String())
}
func (e EchoApplication) OnLogout(sessionID session.ID) {
  fmt.Printf("OnLogout %v\n", sessionID.String())
}
func (e EchoApplication) ToAdmin(msgBuilder message.Builder, sessionID session.ID) {}

func (e EchoApplication) ToApp(msgBuilder message.Builder, sessionID session.ID) (err error) {
  return
}

func (e EchoApplication) FromAdmin(msg message.Message, sessionID session.ID) (reject reject.MessageReject) {
  return
}
func (e EchoApplication) FromApp(msg message.Message, sessionID session.ID) (reject reject.MessageReject) {
  fmt.Println("Got Message ", msg)
  reply:=basic.NewMessage()
  msgType,_:=msg.Header().Get(fix.MsgType)
  reply.MsgHeader.Set(msgType)

  for _,tag:=range msg.Body().Tags() {
    if field,ok:=msg.Body().Get(tag); ok {
      reply.MsgBody.Set(field)
    }
  }

  session.SendToTarget(reply, sessionID)

  return
}

func main() {
  app:=new(EchoApplication)

  globalSettings:=settings.NewDictionary()
  globalSettings.SetInt(settings.SocketAcceptPort, 5001)
  globalSettings.SetString(settings.SenderCompID, "ISLD")
  globalSettings.SetString(settings.TargetCompID, "TW")

  appSettings:=settings.NewApplicationSettings(globalSettings)
  appSettings.AddSession("FIX40", settings.NewDictionary().SetString(settings.BeginString, "FIX.4.0"))
  appSettings.AddSession("FIX41", settings.NewDictionary().SetString(settings.BeginString, "FIX.4.1"))
  appSettings.AddSession("FIX42", settings.NewDictionary().SetString(settings.BeginString, "FIX.4.2"))
  appSettings.AddSession("FIX43", settings.NewDictionary().SetString(settings.BeginString, "FIX.4.3"))
  appSettings.AddSession("FIX44", settings.NewDictionary().SetString(settings.BeginString, "FIX.4.4"))
  appSettings.AddSession("FIX50", settings.NewDictionary().SetString(settings.BeginString, "FIXT.1.1").SetString(settings.DefaultApplVerID, "7"))
  appSettings.AddSession("FIX50SP1", settings.NewDictionary().SetString(settings.BeginString, "FIXT.1.1").SetString(settings.DefaultApplVerID, "8"))
  appSettings.AddSession("FIX50SP2", settings.NewDictionary().SetString(settings.BeginString, "FIXT.1.1").SetString(settings.DefaultApplVerID, "9"))

  acceptor,err:=quickfixgo.NewAcceptor(app,appSettings,log.ScreenLogFactory{})
  if err!=nil {
    fmt.Printf("Unable to create Acceptor: ", err)
    return
  }

  if err=acceptor.Start(); err!=nil {
    fmt.Printf("Unable to start Acceptor: ", err)
    return
  }

  interrupt:=make(chan os.Signal)
  signal.Notify(interrupt)
  <-interrupt

  acceptor.Stop()
}
