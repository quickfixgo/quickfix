package quickfix

type sessionState interface {
	FixMsgIn(*session, Message) (nextState sessionState)
	Timeout(*session, event) (nextState sessionState)
}
