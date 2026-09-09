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
	"testing"
)

func TestPR514(t *testing.T) {
	sessionID := SessionID{
		BeginString:  BeginStringFIX42,
		TargetCompID: "BigCorp",
		TargetSubID:  "acceptor",
		SenderCompID: "SmallCorp",
		SenderSubID:  "initiator",
	}
	storeFactory := NewMemoryStoreFactory()
	store, err := storeFactory.Create(sessionID)
	if err != nil {
		t.Error(err)
	}
	s := &session{
		sessionID: sessionID,
		store:     store,
	}
	if err := registerSession(s); err != nil {
		t.Error(err)
	}

	msg := NewMessage()
	msg.Header.SetString(tagBeginString, sessionID.BeginString)
	msg.Header.SetString(tagTargetCompID, sessionID.TargetCompID)
	msg.Header.SetString(tagTargetSubID, sessionID.TargetSubID)
	msg.Header.SetString(tagSenderCompID, sessionID.SenderCompID)
	msg.Header.SetString(tagSenderSubID, sessionID.SenderSubID)

	if err := Send(msg); err != nil {
		if errors.Is(err, errUnknownSession) {
			t.Errorf("Unable to find registered session: %s", err)
		}
	}
}
