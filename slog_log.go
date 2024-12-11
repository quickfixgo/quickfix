package quickfix

import (
	"fmt"
	"log/slog"
	"strings"
)

type SlogLog struct {
	Name string
	*slog.Logger
}

var _ LogFactory = &SlogLog{}
var _ Log = &SlogLog{}

func (l *SlogLog) OnIncoming(s []byte) {
	l.logMessage("<-"+l.Name, s)
}

func (l *SlogLog) OnOutgoing(s []byte) {
	l.logMessage("->"+l.Name, s)
}

func (l *SlogLog) OnEvent(s string) {
	slog.Info("->"+l.Name+":event", "event", s)
}

func (l *SlogLog) OnEventf(format string, a ...interface{}) {
	l.OnEvent(fmt.Sprintf(format, a...))
}

// NewLogger creates an instance of LogFactory that writes messages and events to stdout.
func NewSlogLogger() *SlogLog {
	return &SlogLog{Logger: slog.Default()}
}

func (l *SlogLog) Create() (Log, error) {
	return l, nil
}

func (l *SlogLog) CreateSessionLog(sessionID SessionID) (Log, error) {
	return &SlogLog{Name: l.Name, Logger: l.Logger.With("session", sessionID.String())}, nil
}

func (l *SlogLog) logMessage(msg string, data []byte) {
	slog.Info(msg, "message", ToValues(data))
}

func ToValues(s []byte) slog.Value {
	fields := AsFields(s)
	var attrs []slog.Attr
	for _, f := range fields {
		attrs = append(attrs, slog.String(f.Tag, f.Value))
	}
	return slog.GroupValue(attrs...)
}

const sohCode = "\001"

func SplitParts(msg []byte) []string {
	return strings.Split(string(msg), sohCode)
}
func ToFields(splitMsg []string) []LField {
	var fields []LField
	for _, part := range splitMsg {
		parts := strings.Split(part, "=")
		if len(parts) == 2 {
			fields = append(fields, LField{Tag: parts[0], Value: parts[1]})
		}
	}
	return fields
}

type LField struct {
	Tag   string
	Value string
}

func AsFields(s []byte) []LField {
	return ToFields(SplitParts(s))
}
