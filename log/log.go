//Package log implements logging of FIX messages and events.
package log

type Log interface {
	//log incoming fix message
	OnIncoming(string)

	//log outgoing fix message
	OnOutgoing(string)

	//log fix event
	OnEvent(string)

	//log fix event according to format specifier
	OnEventf(string, ...interface{})
}

type LogFactory interface {
	//global log
	Create() Log

	//session specific log
	CreateSessionLog(prefix string) Log
}
