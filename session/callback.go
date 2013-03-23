package session

import(
    "quickfixgo/reject"
    "quickfixgo/message"
    )

//Callback interface for receiving session events
type Callback interface {
  OnCreate(sessionID ID)
  OnLogon(sessionID ID)
  OnLogout(sessionID ID)

  FromAdmin(msg message.Message, sessionID ID) reject.MessageReject
  FromApp(msg message.Message, sessionID ID) reject.MessageReject
}
