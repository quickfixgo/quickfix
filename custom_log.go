package quickfix

import (
	"bytes"
	"fmt"
)

func makeReadable(s []byte) string {
	return string(bytes.Replace(s, []byte("\x01"), []byte("|"), -1))
}

type customLog struct {
	sessionPrefix string
	logFunc       func(msg string, keysAndValues ...interface{})
}

func (l customLog) OnIncoming(s []byte) {
	msg := fmt.Sprintf("FIX incoming [%v |%v]", l.sessionPrefix, makeReadable(s))
	l.logFunc(msg)
}

func (l customLog) OnOutgoing(s []byte) {
	msg := fmt.Sprintf("FIX outgoing [%v - |%v]", l.sessionPrefix, makeReadable(s))
	l.logFunc(msg)
}

func (l customLog) OnEvent(s string) {
	msg := fmt.Sprintf("FIX event [%v - %v]", l.sessionPrefix, s)
	l.logFunc(msg)
}

func (l customLog) OnEventf(format string, a ...interface{}) {
	l.OnEvent(fmt.Sprintf(format, a...))
}

type customLogFactory struct {
	logFunc func(msg string, keysAndValues ...interface{})
}

func (f customLogFactory) Create() (Log, error) {
	log := customLog{"GLOBAL", f.logFunc}
	return log, nil
}

func (f customLogFactory) CreateSessionLog(sessionID SessionID) (Log, error) {
	log := customLog{sessionID.String(), f.logFunc}
	return log, nil
}

// NewCustomLogFactory creates an instance of LogFactory that
// logs messages and events using the provided log function
func NewCustomLogFactory(logFunc func(msg string, keysAndValues ...interface{})) LogFactory {
	return customLogFactory{logFunc: logFunc}
}
