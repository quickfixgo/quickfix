// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"errors"
	"sync"
)

var sessionsLock sync.RWMutex
var sessions = make(map[SessionID]*session)
var errDuplicateSessionID = errors.New("Duplicate SessionID")
var errUnknownSession = errors.New("Unknown session")

// Messagable is a Message or something that can be converted to a Message.
type Messagable interface {
	ToMessage() *Message
}

// Send determines the session to send Messagable using header fields BeginString, TargetCompID, SenderCompID.
func Send(m Messagable) (err error) {
	msg := m.ToMessage()
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
		return err
	}

	sessionID := SessionID{BeginString: string(beginString), TargetCompID: string(targetCompID), SenderCompID: string(senderCompID)}

	return SendToTarget(msg, sessionID)
}

// SendToTarget sends a message based on the sessionID. Convenient for use in FromApp since it provides a session ID for incoming messages.
func SendToTarget(m Messagable, sessionID SessionID) error {
	msg := m.ToMessage()
	session, ok := lookupSession(sessionID)
	if !ok {
		return errUnknownSession
	}

	return session.queueForSend(msg)
}

// ResetSession resets session's sequence numbers.
func ResetSession(sessionID SessionID) error {
	session, ok := lookupSession(sessionID)
	if !ok {
		return errUnknownSession
	}
	session.log.OnEvent("Session reset")
	session.State.ShutdownNow(session)
	if err := session.dropAndReset(); err != nil {
		session.logError(err)
		return err
	}

	return nil
}

// UnregisterSession removes a session from the set of known sessions.
func UnregisterSession(sessionID SessionID) error {
	sessionsLock.Lock()
	defer sessionsLock.Unlock()

	if _, ok := sessions[sessionID]; ok {
		delete(sessions, sessionID)
		return nil
	}

	return errUnknownSession
}

// SetNextTargetMsgSeqNum set the next expected target message sequence number for the session matching the session id.
func SetNextTargetMsgSeqNum(sessionID SessionID, seqNum int) error {
	session, ok := lookupSession(sessionID)
	if !ok {
		return errUnknownSession
	}
	return session.store.SetNextTargetMsgSeqNum(seqNum)
}

// SetNextSenderMsgSeqNum sets the next outgoing message sequence number for the session matching the session id.
func SetNextSenderMsgSeqNum(sessionID SessionID, seqNum int) error {
	session, ok := lookupSession(sessionID)
	if !ok {
		return errUnknownSession
	}
	return session.store.SetNextSenderMsgSeqNum(seqNum)
}

// GetExpectedSenderNum retrieves the expected sender sequence number for the session matching the session id.
func GetExpectedSenderNum(sessionID SessionID) (int, error) {
	session, ok := lookupSession(sessionID)
	if !ok {
		return 0, errUnknownSession
	}
	return session.store.NextSenderMsgSeqNum(), nil
}

// GetExpectedTargetNum retrieves the next target sequence number for the session matching the session id.
func GetExpectedTargetNum(sessionID SessionID) (int, error) {
	session, ok := lookupSession(sessionID)
	if !ok {
		return 0, errUnknownSession
	}
	return session.store.NextTargetMsgSeqNum(), nil
}

// GetMessageStore returns the MessageStore interface for session matching the session id.
func GetMessageStore(sessionID SessionID) (MessageStore, error) {
	session, ok := lookupSession(sessionID)
	if !ok {
		return nil, errUnknownSession
	}
	return session.store, nil
}

// GetLog returns the Log interface for session matching the session id.
func GetLog(sessionID SessionID) (Log, error) {
	session, ok := lookupSession(sessionID)
	if !ok {
		return nil, errUnknownSession
	}
	return session.log, nil
}

func registerSession(s *session) error {
	sessionsLock.Lock()
	defer sessionsLock.Unlock()

	if _, ok := sessions[s.sessionID]; ok {
		return errDuplicateSessionID
	}

	sessions[s.sessionID] = s
	return nil
}

func lookupSession(sessionID SessionID) (s *session, ok bool) {
	sessionsLock.RLock()
	defer sessionsLock.RUnlock()

	s, ok = sessions[sessionID]
	return
}
