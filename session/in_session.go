package session

import(
    "quickfixgo/fix"
    "quickfixgo/message"
    )

type inSession struct {

}

func (state inSession) OnFixMsgIn(session *session, msg message.Message) (nextState state) {
  if msgType, err:=msg.Header().StringField(fix.MsgType); err==nil {
    switch msgType.Value() {
      //logout
      case "5":
        return state.handleLogout(session, msg)
    }
  }

  session.expectedSeqNum++

  return state
}

func (state inSession) handleLogout(session *session, msg message.Message) (nextState state){
  session.generateLogout()
  return latentState{}
}
