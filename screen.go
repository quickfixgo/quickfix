package quickfix

import (
	"fmt"
	"time"
)

type screenLog struct {
	prefix string
}

//ScreenLogFactory creates Log instances that write to screen.
//Implements LogFactory interface.
type ScreenLogFactory struct{}

//Create returns a logger that writes to stdout with prefix "GLOBAL"
func (ScreenLogFactory) Create() Log {
	log := screenLog{"GLOBAL"}
	return log
}

//CreateSessionLog returns a logger that writes to stdout with a prefix equal to SessionID.String()
func (ScreenLogFactory) CreateSessionLog(sessionID SessionID) Log {
	log := screenLog{sessionID.String()}
	return log
}

func (l screenLog) OnIncoming(s string) {
	logTime := time.Now().UTC()
	fmt.Printf("<%v, %s, incoming>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnOutgoing(s string) {
	logTime := time.Now().UTC()
	fmt.Printf("<%v, %s, outgoing>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnEvent(s string) {
	logTime := time.Now().UTC()
	fmt.Printf("<%v, %s, event>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnEventf(format string, a ...interface{}) {
	l.OnEvent(fmt.Sprintf(format, a...))
}
