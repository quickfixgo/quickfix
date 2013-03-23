package session

import (
	"quickfixgo/message"
)

type state interface {
	FixMsgIn(*session, message.Message) (nextState state)
	Timeout(*session, event) (nextState state)
}
