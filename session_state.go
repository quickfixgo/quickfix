package quickfix

import (
	"fmt"

	"github.com/quickfixgo/quickfix/internal"
)

type stateMachine struct {
	State sessionState
}

func (sm *stateMachine) Disconnected(session *session) {
	sm.setState(session, latentState{})
}

func (sm *stateMachine) FixMsgIn(session *session, m Message) {
	if sm.State == nil {
		return
	}
	sm.setState(session, sm.State.FixMsgIn(session, m))
}

func (sm *stateMachine) Timeout(session *session, e internal.Event) {
	if sm.State == nil {
		return
	}

	sm.setState(session, sm.State.Timeout(session, e))
}

func (sm *stateMachine) setState(session *session, nextState sessionState) {
	if sm.IsConnected() && !nextState.IsConnected() {
		sm.handleDisconnectState(session)
	}

	sm.State = nextState
}

func (sm *stateMachine) handleDisconnectState(s *session) {
	doOnLogout := s.IsLoggedOn()

	switch s.State.(type) {
	case logoutState:
		s.application.OnLogout(s.sessionID)
	case logonState:
		if s.initiateLogon {
			s.application.OnLogout(s.sessionID)
		}
	}

	if doOnLogout {
		s.application.OnLogout(s.sessionID)
	}

	s.onDisconnect()
}

func (sm *stateMachine) IsLoggedOn() bool {
	return sm.State.IsLoggedOn()
}

func (sm *stateMachine) IsConnected() bool {
	return sm.State.IsConnected()
}

func handleStateError(s *session, err error) sessionState {
	s.logError(err)
	return latentState{}
}

//sessionState is the current state of the session state machine. The session state determines how the session responds to
//incoming messages, timeouts, and requests to send application messages.
type sessionState interface {
	//FixMsgIn is called by the session on incoming messages from the counter party.  The return type is the next session state
	//following message processing
	FixMsgIn(*session, Message) (nextState sessionState)

	//Timeout is called by the session on a timeout event.
	Timeout(*session, internal.Event) (nextState sessionState)

	//IsLoggedOn returns true if state is logged on an in session, false otherwise
	IsLoggedOn() bool

	//IsConnected returns true if the state is connected
	IsConnected() bool

	//debugging convenience
	fmt.Stringer
}
