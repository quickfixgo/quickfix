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

// Log is a generic interface for logging FIX messages and events.
type Log interface {
	// OnIncoming log incoming fix message.
	OnIncoming([]byte)

	// OnOutgoing log outgoing fix message.
	OnOutgoing([]byte)

	// OnEvent log fix event.
	OnEvent(string)

	// OnEventf log fix event according to format specifier.
	OnEventf(string, ...interface{})
}

// The LogFactory interface creates global and session specific Log instances.
type LogFactory interface {
	// Create global log.
	Create() (Log, error)

	// CreateSessionLog session specific log.
	CreateSessionLog(sessionID SessionID) (Log, error)
}
