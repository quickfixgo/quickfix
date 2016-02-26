package quickfix

import (
	"fmt"
)

//Marshaler Marshals self to quickfix.Message type
type Marshaler interface {
	Marshal() Message
}

//Send determines the session to send Marshaler using header fields BeginString, TargetCompID, SenderCompID
func Send(m Marshaler) (err error) {
	msg := m.Marshal()
	var beginString FIXString
	if err := msg.Header.GetField(tagBeginString, &beginString); err != nil {
		return err
	}

	var targetCompID FIXString
	if err := msg.Header.GetField(tagTargetCompID, &targetCompID); err != nil {
		return err
	}

	var senderCompID FIXString
	if err := msg.Header.GetField(tagSenderCompID, &senderCompID); err != nil {

		return nil
	}

	sessionID := SessionID{BeginString: string(beginString), TargetCompID: string(targetCompID), SenderCompID: string(senderCompID)}

	return SendToTarget(msg, sessionID)
}

//SendToTarget sends a message based on the sessionID. Convenient for use in FromApp since it provides a session ID for incoming messages
func SendToTarget(m Marshaler, sessionID SessionID) error {
	msg := m.Marshal()
	session, err := lookupSession(sessionID)
	if err != nil {
		return err
	}

	//NOTE: must queue for send here. otherwise, if not executed in same goroutine as session run loop,
	//message may be sent on closed channel or sent outside of valid state
	session.queueForSend(msg)

	return nil
}

type sessionActivate struct {
	SessionID
	reply chan *session
}

type sessionResource struct {
	session *session
	active  bool
}

type sessionLookupResponse struct {
	session *session
	err     error
}

type sessionLookup struct {
	SessionID
	reply chan sessionLookupResponse
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

//lookupSession returns the Session associated with the sessionID.
func lookupSession(sessionID SessionID) (*session, error) {
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
