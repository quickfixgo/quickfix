package quickfixgo

import (
	"quickfixgo/message"
	"quickfixgo/reject"
	"quickfixgo/session"
)

type Application interface {
	//Notification of a session begin created. 
	OnCreate(sessionID session.ID)

	//Notification of a session successfully logging on. 
	OnLogon(sessionID session.ID)

	//Notification of a session logging off or disconnecting. 
	OnLogout(sessionID session.ID)

	//Notification of admin message being sent to target. 
	ToAdmin(msgBuilder message.Builder, sessionID session.ID)

	//Notification of app message being sent to target. 
	ToApp(msgBuilder message.Builder, sessionID session.ID) error

	//Notification of admin message being received from target. 
	FromAdmin(msg message.Message, sessionID session.ID) reject.MessageReject

	//Notification of app message being received from target. 
	FromApp(msg message.Message, sessionID session.ID) reject.MessageReject
}
