package session

import(
    "quickfixgo/message"
    )

type latentState struct {

}

func (state latentState) OnFixMsgIn(session *session, msg message.Message) (nextState state) {
  session.Log.OnEventf("Invalid Session State: Unexpected Msg %v", msg)
  return state
}
