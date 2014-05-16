package quickfix

type sessionState interface {
	FixMsgIn(*Session, Message) (nextState sessionState)
	Timeout(*Session, event) (nextState sessionState)
}
