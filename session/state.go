package session

import(
    "quickfixgo/message"
    )

type state interface {
  OnFixMsgIn(*session, message.Message) (nextState state)
}
