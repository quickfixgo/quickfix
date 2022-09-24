package quickfix

// Log is a generic interface for logging FIX messages and events.
type Log interface {
	//OnIncoming log incoming fix message
	OnIncoming([]byte)

	//OnOutgoing log outgoing fix message
	OnOutgoing([]byte)

	//OnEvent log fix event
	OnEvent(string)

	//OnEventf log fix event according to format specifier
	OnEventf(string, ...interface{})
}

// The LogFactory interface creates global and session specific Log instances
type LogFactory interface {
	//Create global log
	Create() (Log, error)

	//CreateSessionLog session specific log
	CreateSessionLog(sessionID SessionID) (Log, error)
}
