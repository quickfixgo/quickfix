package main

import(
    "fmt"
    "quickfixgo"
    "quickfixgo/log"
    "quickfixgo/message"
    )

type EchoApplication struct {

}

func (e EchoApplication) OnCreate(sessionID quickfixgo.SessionID) {}
func (e EchoApplication) OnLogon(sessionID quickfixgo.SessionID) {}
func (e EchoApplication) OnLogout(sessionID quickfixgo.SessionID) {}
func (e EchoApplication) ToAdmin(msgBuilder message.Builder, sessionID quickfixgo.SessionID) {}

func (e EchoApplication) ToApp(msgBuilder message.Builder, sessionID quickfixgo.SessionID) (err error) {
  return
}

func (e EchoApplication) FromAdmin(msg message.Message, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
  return
}
func (e EchoApplication) FromApp(msg message.Message, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
  return
}


func main() {
  app:=new(EchoApplication)

  settings:=quickfixgo.NewSessionSettings()
  settings.GetGlobalSettings().SetInt(quickfixgo.SocketAcceptPort, 5000)
  
  sessionDict:=quickfixgo.NewDictionary()
  sessionID:=quickfixgo.SessionID{BeginString: "FIX.4.2", SenderCompID: "ISLD", TargetCompID: "TW"}
  settings.SetSessionSettings(sessionID, sessionDict)

  acceptor,err:=quickfixgo.NewAcceptor(app,settings,log.ScreenLogFactory{})
  if acceptor == nil {
    fmt.Printf("Unable to create Acceptor: ", err)
    return
  }

  if err=acceptor.Start(); err!=nil {
    fmt.Printf("Unable to start Acceptor: ", err)
    return
  }

  var i int
  fmt.Scanf("%d", &i)

  acceptor.Stop()
}
