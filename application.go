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

// Application interface should be implemented by FIX Applications.
// This is the primary interface for processing messages from a FIX Session.
type Application interface {
	// OnCreate notification of a session begin created.
	OnCreate(sessionID SessionID)

	// OnLogon notification of a session successfully logging on.
	OnLogon(sessionID SessionID)

	// OnLogout notification of a session logging off or disconnecting.
	OnLogout(sessionID SessionID)

	// ToAdmin notification of admin message being sent to target.
	ToAdmin(message *Message, sessionID SessionID)

	// ToApp notification of app message being sent to target.
	ToApp(message *Message, sessionID SessionID) error

	// FromAdmin notification of admin message being received from target.
	FromAdmin(message *Message, sessionID SessionID) MessageRejectError

	// FromApp notification of app message being received from target.
	FromApp(message *Message, sessionID SessionID) MessageRejectError
}
