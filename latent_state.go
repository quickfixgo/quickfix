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

type latentState struct{ inSessionTime }

func (state latentState) String() string    { return "Latent State" }
func (state latentState) IsLoggedOn() bool  { return false }
func (state latentState) IsConnected() bool { return false }

func (state latentState) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	session.log.OnEventf("Invalid Session State: Unexpected Msg %v while in Latent state", msg)
	return state
}

func (state latentState) Timeout(*session, internal.Event) (nextState sessionState) {
	return state
}

func (state latentState) ShutdownNow(*session) {}
func (state latentState) Stop(*session) (nextState sessionState) {
	return state
}
