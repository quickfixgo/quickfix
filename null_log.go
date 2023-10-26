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

type nullLog struct{}

func (l nullLog) OnIncoming([]byte)                        {}
func (l nullLog) OnOutgoing([]byte)                        {}
func (l nullLog) OnEvent(string)                           {}
func (l nullLog) OnEventf(format string, a ...interface{}) {}

type nullLogFactory struct{}

func (nullLogFactory) Create() (Log, error) {
	return nullLog{}, nil
}
func (nullLogFactory) CreateSessionLog(sessionID SessionID) (Log, error) {
	return nullLog{}, nil
}

// NewNullLogFactory creates an instance of LogFactory that returns no-op loggers.
func NewNullLogFactory() LogFactory {
	return nullLogFactory{}
}
