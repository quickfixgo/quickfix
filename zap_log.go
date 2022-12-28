package quickfix

import (
	"bytes"
	"go.uber.org/zap"
)

// ZAP logger for quickfix
// Uses ZAP (https://pkg.go.dev/go.uber.org/zap), currently the lowest latency logger for golang
// Supports screen logging, json logging and a fair amount of additional customization
// This wrapper is pretty straightforward passing the quickfix logging directly to zap at info level
// It sets a named logger for each session with Prefix+Sender:Target as the label.
// Setting prefix to "" results in the message being labeled with just the sender and target.
// There is also an option to enable or disable detailed message logging.
// If message logging is enabled, EOM field separators are replaced with | for readability.
//
// Possible future improvements:
//   - Take advantage of ZAP field based logging to put the message body in its own json field
//   - Allow different log levels for messages and events
//   - Pre-check info is enabled - which would save some formatting time if info is disabled (unlikely?)
//   - Support split logging of events to ZAP and messages to a file
//

type ZapLog struct {
	logger      *zap.SugaredLogger
	rawMessages bool
	name        string
}

type ZapLogFactory struct {
	prefix   string
	messages bool
}

func (l ZapLog) OnIncoming(s []byte) {
	if l.rawMessages {
		sNew := ReplaceEos(s)
		l.logger.Infof("incoming <-:  %s", sNew)
	}
}

func (l ZapLog) OnOutgoing(s []byte) {
	if l.rawMessages {
		sNew := ReplaceEos(s)
		l.logger.Infof("outgoing ->: %s", sNew)
	}
}

func (l ZapLog) OnEvent(s string) {
	l.logger.Infof("event:  %s", s)
}

func (l ZapLog) OnEventf(format string, a ...interface{}) {
	l.logger.Infof(format, a...)
}

func (f ZapLogFactory) Create() (Log, error) {
	name := f.prefix + "GLOBAL"
	log := ZapLog{
		name:        name,
		logger:      zap.S().Named(name).WithOptions(zap.AddCallerSkip(1)),
		rawMessages: f.messages,
	}
	return log, nil
}

func (f ZapLogFactory) CreateSessionLog(sessionID SessionID) (Log, error) {
	name := f.prefix + sessionID.SenderCompID + ":" + sessionID.TargetCompID
	log := ZapLog{
		name:        name,
		logger:      zap.S().Named(name).WithOptions(zap.AddCallerSkip(1)),
		rawMessages: f.messages,
	}
	return log, nil
}

// NewZapLogFactory Factory for quickfix to use zap logging
func NewZapLogFactory(prefix string, logMessages bool) ZapLogFactory {
	return ZapLogFactory{
		prefix:   prefix,
		messages: logMessages,
	}
}

// ReplaceEos
// Format message for display, replace EOS (ascii 001) with |
func ReplaceEos(s []byte) []byte {
	sNew := bytes.Replace(s, []byte("\u0001"), []byte("|"), 9999)
	return sNew
}
