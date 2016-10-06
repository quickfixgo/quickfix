package quickfix

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
	ToAdmin(message *Message, sessionID SessionID)

	//Notification of app message being sent to target.
	ToApp(message *Message, sessionID SessionID) error

	//Notification of admin message being received from target.
	FromAdmin(message *Message, sessionID SessionID) MessageRejectError

	//Notification of app message being received from target.
	FromApp(message *Message, sessionID SessionID) MessageRejectError
}
