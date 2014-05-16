package quickfix

type nullLog struct{}

func (l nullLog) OnIncoming(s string)                      {}
func (l nullLog) OnOutgoing(s string)                      {}
func (l nullLog) OnEvent(s string)                         {}
func (l nullLog) OnEventf(format string, a ...interface{}) {}

type nullLogFactory struct{}

func (nullLogFactory) Create() (Log, error) {
	return nullLog{}, nil
}
func (nullLogFactory) CreateSessionLog(sessionID SessionID) (Log, error) {
	return nullLog{}, nil
}

//NewNullLogFactory creates an instance of LogFactory that returns no-op loggers.
func NewNullLogFactory() LogFactory {
	return nullLogFactory{}
}
