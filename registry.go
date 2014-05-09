package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Send determines the session to send msgBuilder using header fields BeginString, TargetCompID, SenderCompID
func Send(msg message.MessageBuilder) (err error) {
	var beginString field.BeginStringField
	if err := msg.Header.Get(&beginString); err != nil {
		return err
	}

	var targetCompID field.TargetCompIDField
	if err := msg.Header.Get(&targetCompID); err != nil {
		return err
	}

	var senderCompID field.SenderCompIDField
	if err := msg.Header.Get(&senderCompID); err != nil {

		return nil
	}

	sessionID := SessionID{BeginString: beginString.Value, TargetCompID: targetCompID.Value, SenderCompID: senderCompID.Value}

	return SendToTarget(msg, sessionID)
}

func SendToTarget(msgBuilder message.MessageBuilder, sessionID SessionID) error {
	session, err := LookupSession(sessionID)
	if err != nil {
		return err
	}

	session.send(msgBuilder)

	return nil
}

type sessionActivate struct {
	SessionID
	reply chan *Session
}

type sessionResource struct {
	session *Session
	active  bool
}

type sessionLookupResponse struct {
	session *Session
	err     error
}

type sessionLookup struct {
	SessionID
	reply chan sessionLookupResponse
}

type registry struct {
	newSession chan *Session
	activate   chan sessionActivate
	deactivate chan SessionID
	lookup     chan sessionLookup
}

var sessions *registry

func init() {
	sessions = new(registry)
	sessions.newSession = make(chan *Session)
	sessions.activate = make(chan sessionActivate)
	sessions.deactivate = make(chan SessionID)
	sessions.lookup = make(chan sessionLookup)

	go sessions.sessionResourceServerLoop()
}

func activate(sessionID SessionID) *Session {
	response := make(chan *Session)
	sessions.activate <- sessionActivate{sessionID, response}
	return <-response
}

func deactivate(sessionID SessionID) {
	sessions.deactivate <- sessionID
}

//LookupSession returns the Session associated with the sessionID.
func LookupSession(sessionID SessionID) (*Session, error) {
	responseChannel := make(chan sessionLookupResponse)
	sessions.lookup <- sessionLookup{sessionID, responseChannel}

	response := <-responseChannel
	return response.session, response.err
}

func (r *registry) sessionResourceServerLoop() {
	sessions := make(map[SessionID]*sessionResource)

	for {
		select {
		case session := <-r.newSession:
			sessions[session.sessionID] = &sessionResource{session, false}

		case deactivatedID := <-r.deactivate:
			if resource, ok := sessions[deactivatedID]; ok {
				resource.active = false
			}

		case lookup := <-r.lookup:
			if resource, ok := sessions[lookup.SessionID]; ok {
				lookup.reply <- sessionLookupResponse{resource.session, nil}
			} else {
				lookup.reply <- sessionLookupResponse{nil, fmt.Errorf("session not found")}
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
