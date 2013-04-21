package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/message"
)

type state interface {
	FixMsgIn(*session, message.Message) (nextState state)
	Timeout(*session, event) (nextState state)
}
