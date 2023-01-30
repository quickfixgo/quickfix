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

import "github.com/quickfixgo/quickfix/internal"

type logoutState struct{ connectedNotLoggedOn }

func (state logoutState) String() string { return "Logout State" }

func (state logoutState) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	nextState = inSession{}.FixMsgIn(session, msg)
	if nextState, ok := nextState.(latentState); ok {
		return nextState
	}

	return state
}

func (state logoutState) Timeout(session *session, event internal.Event) (nextState sessionState) {
	switch event {
	case internal.LogoutTimeout:
		session.log.OnEvent("Timed out waiting for logout response")
		return latentState{}
	}

	return state
}

func (state logoutState) Stop(session *session) (nextstate sessionState) {
	return state
}
