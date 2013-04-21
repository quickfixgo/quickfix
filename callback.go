package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
)

//Callback interface for receiving session events
type Callback interface {
	OnCreate(sessionID SessionID)
	OnLogon(sessionID SessionID)
	OnLogout(sessionID SessionID)

	FromAdmin(msg message.Message, sessionID SessionID) reject.MessageReject
	FromApp(msg message.Message, sessionID SessionID) reject.MessageReject
}
