package quickfix

// Application interface should be implemented by FIX Applications.
// This is the primary interface for processing messages from a FIX Session.
type Application interface {
	//OnCreate notification of a session begin created.
	OnCreate(sessionID SessionID)

	//OnLogon notification of a session successfully logging on.
	OnLogon(sessionID SessionID)

	//OnLogout notification of a session logging off or disconnecting.
	OnLogout(sessionID SessionID)

	//ToAdmin notification of admin message being sent to target.
	ToAdmin(message *Message, sessionID SessionID)

	//ToApp notification of app message being sent to target.
	ToApp(message *Message, sessionID SessionID) error

	//FromAdmin notification of admin message being received from target.
	FromAdmin(message *Message, sessionID SessionID) MessageRejectError

	//FromApp notification of app message being received from target.
	FromApp(message *Message, sessionID SessionID) MessageRejectError
}
