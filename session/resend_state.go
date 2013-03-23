package session

import(
  "quickfixgo/message"
    )

type resendState struct {
  inSession
}


func (state resendState) FixMsgIn(session *session, msg message.Message) (nextState state) {
  return state.inSession.FixMsgIn(session, msg)
}

