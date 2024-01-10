package quickfix

type EventType int

const (
	EventTypeLogon EventType = iota
)

type EventLogon struct {
	Addr string
	TS   int64
}
