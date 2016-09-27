package quickfix

//Log is a generic interface for logging FIX messages and events.
type Log interface {
	//log incoming fix message
	OnIncoming([]byte)

	//log outgoing fix message
	OnOutgoing([]byte)

	//log fix event
	OnEvent(string)

	//log fix event according to format specifier
	OnEventf(string, ...interface{})
}

//The LogFactory interface creates global and session specific Log instances
type LogFactory interface {
	//global log
	Create() (Log, error)

	//session specific log
	CreateSessionLog(sessionID SessionID) (Log, error)
}
