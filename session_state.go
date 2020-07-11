package quickfix

import (
	"fmt"
	"time"

	"github.com/quickfixgo/quickfix/internal"
)

type stateMachine struct {
	State                 sessionState
	pendingStop, stopped  bool
	notifyOnInSessionTime chan interface{}
}

func (sm *stateMachine) Start(s *session) {
	sm.pendingStop = false
	sm.stopped = false

	sm.State = latentState{}
	sm.CheckSessionTime(s, time.Now())
}

func (sm *stateMachine) Connect(session *session) {
	// No special logon logic needed for FIX Acceptors.
	if !session.InitiateLogon {
		sm.setState(session, logonState{})
		return
	}

	if session.RefreshOnLogon {
		if err := session.store.Refresh(); err != nil {
			session.logError(err)
			return
		}
	}
	session.log.OnEvent("Sending logon request")
	if err := session.sendLogon(); err != nil {
		session.logError(err)
		return
	}

	sm.setState(session, logonState{})
	// Fire logon timeout event after the pre-configured delay period.
	time.AfterFunc(session.LogonTimeout, func() { session.sessionEvent <- internal.LogonTimeout })
}

func (sm *stateMachine) Stop(session *session) {
	sm.pendingStop = true
	sm.setState(session, sm.State.Stop(session))
}

func (sm *stateMachine) Stopped() bool {
	return sm.stopped
}

func (sm *stateMachine) Disconnected(session *session) {
	if sm.IsConnected() {
		sm.setState(session, latentState{})
	}
}

func (sm *stateMachine) Incoming(session *session, m fixIn) {
	sm.CheckSessionTime(session, time.Now())
	if !sm.IsConnected() {
		return
	}

	session.log.OnIncoming(m.bytes.Bytes())

	msg := session.messagePool.Get()
	if err := ParseMessageWithDataDictionary(msg, m.bytes, session.transportDataDictionary, session.appDataDictionary); err != nil {
		session.log.OnEventf("Msg Parse Error: %v, %q", err.Error(), m.bytes)
	} else {
		msg.ReceiveTime = m.receiveTime
		sm.fixMsgIn(session, msg)
	}

	if !msg.keepMessage {
		session.returnToPool(msg)
	}
	session.peerTimer.Reset(time.Duration(float64(1.2) * float64(session.HeartBtInt)))
}

func (sm *stateMachine) fixMsgIn(session *session, m *Message) {
	sm.setState(session, sm.State.FixMsgIn(session, m))
}

func (sm *stateMachine) SendAppMessages(session *session) {
	sm.CheckSessionTime(session, time.Now())

	session.sendMutex.Lock()
	defer session.sendMutex.Unlock()

	if session.IsLoggedOn() {
		session.sendQueued()
	} else {
		session.dropQueued()
	}
}

func (sm *stateMachine) Timeout(session *session, e internal.Event) {
	sm.CheckSessionTime(session, time.Now())
	sm.setState(session, sm.State.Timeout(session, e))
}

func (sm *stateMachine) CheckSessionTime(session *session, now time.Time) {
	if !session.SessionTime.IsInRange(now) {
		if sm.IsSessionTime() {
			session.log.OnEvent("Not in session")
		}

		sm.State.ShutdownNow(session)
		sm.setState(session, notSessionTime{})

		if sm.notifyOnInSessionTime == nil {
			sm.notifyOnInSessionTime = make(chan interface{})
		}
		return
	}

	if !sm.IsSessionTime() {
		session.log.OnEvent("In session")
		sm.notifyInSessionTime()
		sm.setState(session, latentState{})
	}

	if !session.SessionTime.IsInSameRange(session.store.CreationTime(), now) {
		session.log.OnEvent("Session reset")
		sm.State.ShutdownNow(session)
		if err := session.dropAndReset(); err != nil {
			session.logError(err)
		}
		sm.setState(session, latentState{})
	}
}

func (sm *stateMachine) setState(session *session, nextState sessionState) {
	if !nextState.IsConnected() {
		if sm.IsConnected() {
			sm.handleDisconnectState(session)
		}

		if sm.pendingStop {
			sm.stopped = true
			sm.notifyInSessionTime()
		}
	}

	sm.State = nextState
}

func (sm *stateMachine) notifyInSessionTime() {
	if sm.notifyOnInSessionTime != nil {
		close(sm.notifyOnInSessionTime)
	}
	sm.notifyOnInSessionTime = nil
}

func (sm *stateMachine) handleDisconnectState(s *session) {
	doOnLogout := s.IsLoggedOn()

	switch s.State.(type) {
	case logoutState:
		doOnLogout = true
	case logonState:
		if s.InitiateLogon {
			doOnLogout = true
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

func (sm *stateMachine) IsSessionTime() bool {
	return sm.State.IsSessionTime()
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
	FixMsgIn(*session, *Message) (nextState sessionState)

	//Timeout is called by the session on a timeout event.
	Timeout(*session, internal.Event) (nextState sessionState)

	//IsLoggedOn returns true if state is logged on an in session, false otherwise
	IsLoggedOn() bool

	//IsConnected returns true if the state is connected
	IsConnected() bool

	//IsSessionTime returns true if the state is in session time
	IsSessionTime() bool

	//ShutdownNow terminates the session state immediately
	ShutdownNow(*session)

	//Stop triggers a clean stop
	Stop(*session) (nextState sessionState)

	//debugging convenience
	fmt.Stringer
}

type inSessionTime struct{}

func (inSessionTime) IsSessionTime() bool { return true }

type connected struct{}

func (connected) IsConnected() bool   { return true }
func (connected) IsSessionTime() bool { return true }

type connectedNotLoggedOn struct{ connected }

func (connectedNotLoggedOn) IsLoggedOn() bool     { return false }
func (connectedNotLoggedOn) ShutdownNow(*session) {}

type loggedOn struct{ connected }

func (loggedOn) IsLoggedOn() bool { return true }
func (loggedOn) ShutdownNow(s *session) {
	if err := s.sendLogout(""); err != nil {
		s.logError(err)
	}
}

func (loggedOn) Stop(s *session) (nextState sessionState) {
	if err := s.initiateLogout(""); err != nil {
		return handleStateError(s, err)
	}

	return logoutState{}
}
