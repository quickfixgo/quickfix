package session

import(
    "quickfixgo/message"
    )

type inSession struct {

}

func (state inSession) OnFixMsgIn(session *session, msg message.Message) (nextState state) {

  return state
}
