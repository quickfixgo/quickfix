package quickfix

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Send determines the session to send msgBuilder using header fields BeginString, TargetCompID, SenderCompID, and DefaultApplVerID
func Send(msg message.MessageBuilder) (err error) {
	var beginString field.BeginString
	if err := msg.Header.Get(&beginString); err != nil {
		return err
	}

	var targetCompID field.TargetCompID
	if err := msg.Header.Get(&targetCompID); err != nil {
		return err
	}

	var senderCompID field.SenderCompID
	if err := msg.Header.Get(&senderCompID); err != nil {

		return nil
	}

	sessionID := SessionID{BeginString: beginString.Value, TargetCompID: targetCompID.Value, SenderCompID: senderCompID.Value}

	var defaultApplVerID field.DefaultApplVerID
	if err := msg.Header.Get(&defaultApplVerID); err == nil {
		sessionID.DefaultApplVerID = defaultApplVerID.Value
	}

	return SendToTarget(msg, sessionID)
}

func SendToTarget(msgBuilder message.MessageBuilder, sessionID SessionID) (err error) {
	session := lookup(sessionID)
	session.send(msgBuilder)

	return
}

type sessionActivate struct {
	SessionID
	reply chan *session
}

type sessionResource struct {
	session *session
	active  bool
}

type sessionLookup struct {
	SessionID
	reply chan *session
}

type registry struct {
	newSession chan *session
	activate   chan sessionActivate
	deactivate chan SessionID
	lookup     chan sessionLookup
}

var sessions *registry

func init() {
	sessions = new(registry)
	sessions.newSession = make(chan *session)
	sessions.activate = make(chan sessionActivate)
	sessions.deactivate = make(chan SessionID)
	sessions.lookup = make(chan sessionLookup)

	go sessions.sessionResourceServerLoop()
}

func activate(sessionID SessionID) *session {
	response := make(chan *session)
	sessions.activate <- sessionActivate{sessionID, response}
	return <-response
}

func deactivate(sessionID SessionID) {
	sessions.deactivate <- sessionID
}

func lookup(sessionID SessionID) *session {
	response := make(chan *session)
	sessions.lookup <- sessionLookup{sessionID, response}
	return <-response
}

func (r *registry) sessionResourceServerLoop() {
	sessions := make(map[SessionID]*sessionResource)

	for {
		select {
		case session := <-r.newSession:
			sessions[session.SessionID] = &sessionResource{session, false}

		case deactivatedID := <-r.deactivate:
			if resource, ok := sessions[deactivatedID]; ok {
				resource.active = false
			}

		case lookup := <-r.lookup:
			if resource, ok := sessions[lookup.SessionID]; ok {
				lookup.reply <- resource.session
			} else {
				lookup.reply <- nil
			}

		case request := <-r.activate:
			resource, ok := sessions[request.SessionID]

			switch {
			case !ok:
				request.reply <- nil

			case resource.active:
				request.reply <- nil

			default:
				resource.active = true
				request.reply <- resource.session
			}
		}
	}
}
