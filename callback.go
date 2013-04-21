package quickfixgo

//Callback interface for receiving session events
type Callback interface {
	OnCreate(sessionID SessionID)
	OnLogon(sessionID SessionID)
	OnLogout(sessionID SessionID)

	FromAdmin(msg Message, sessionID SessionID) MessageReject
	FromApp(msg Message, sessionID SessionID) MessageReject
}
