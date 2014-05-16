package quickfix

import (
	"github.com/quickfixgo/quickfix/errors"
)

//The Application interface should be implemented by FIX Applications.
//This is the primary interface for processing messages from a FIX Session.
type Application interface {
	//Notification of a session begin created.
	OnCreate(sessionID SessionID)

	//Notification of a session successfully logging on.
	OnLogon(sessionID SessionID)

	//Notification of a session logging off or disconnecting.
	OnLogout(sessionID SessionID)

	//Notification of admin message being sent to target.
	ToAdmin(msgBuilder MessageBuilder, sessionID SessionID)

	//Notification of app message being sent to target.
	ToApp(msgBuilder MessageBuilder, sessionID SessionID) error

	//Notification of admin message being received from target.
	FromAdmin(msg Message, sessionID SessionID) errors.MessageRejectError

	//Notification of app message being received from target.
	FromApp(msg Message, sessionID SessionID) errors.MessageRejectError
}
