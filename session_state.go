package quickfix

import (
	"github.com/quickfixgo/quickfix/message"
)

type sessionState interface {
	FixMsgIn(*session, message.Message) (nextState sessionState)
	Timeout(*session, event) (nextState sessionState)
}
