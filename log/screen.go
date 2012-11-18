package log

import (
    "fmt"
    "time"
)

type screenLog struct {
  prefix string
}

//Screen log factory creates log implementations that write to screen.
//Implements LogFactory interface.
type ScreenLogFactory struct {}

func (ScreenLogFactory) Create() Log {
  log:=screenLog{"GLOBAL"}
  return log
}

func (ScreenLogFactory) CreateSessionLog(prefix string) Log {
  log:=screenLog{prefix}
  return log
}

func (l screenLog) OnIncoming(s string) {
  logTime:=time.Now().UTC()
  fmt.Printf("<%v, %s, incoming>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnOutgoing(s string) {
  logTime:=time.Now().UTC()
  fmt.Printf("<%v, %s, outgoing>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnEvent(s string) {
  logTime:=time.Now().UTC()
  fmt.Printf("<%v, %s, event>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnEventf(format string, a ...interface{}) {
  l.OnEvent(fmt.Sprintf(format, a))
}
