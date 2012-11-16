package quickfixgo

type sessionRegistry struct {

}

var sessions *sessionRegistry

func init() {
  sessions=new(sessionRegistry)
}

func SendToTarget(msgBuilder MessageBuilder, sessionID SessionID) (bool, SessionNotFound) {
  return false, nil
}
