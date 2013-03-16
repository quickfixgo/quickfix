package session

import(
  "quickfixgo/message"
    )

type resendState struct {
  inSession
}


func (state resendState) OnFixMsgIn(session *session, msg message.Message) (nextState state) {
  return state.inSession.OnFixMsgIn(session, msg)
}

