package quickfixgo

import(
    "quickfixgo/message"
    )

type sessionRegistry struct {

}

var sessions *sessionRegistry

func init() {
  sessions=new(sessionRegistry)
}

func SendToTarget(msgBuilder message.Builder, sessionID SessionID) (bool, SessionNotFound) {
  return false, nil
}
