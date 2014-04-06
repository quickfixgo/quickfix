package quickfix

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
	FromAdmin(msg Message, sessionID SessionID) MessageReject

	//Notification of app message being received from target.
	FromApp(msg Message, sessionID SessionID) MessageReject
}
