package quickfix

import (
	"github.com/quickfixgo/quickfix/message"
)

type sessionState interface {
	FixMsgIn(*Session, message.Message) (nextState sessionState)
	Timeout(*Session, event) (nextState sessionState)
}
