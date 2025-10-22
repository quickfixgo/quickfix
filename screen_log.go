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
	"fmt"
	"time"
)

type screenLog struct {
	prefix string
}

func (l screenLog) OnIncoming(s []byte) {
	logTime := time.Now().UTC()
	fmt.Printf("<%v, %s, incoming>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnOutgoing(s []byte) {
	logTime := time.Now().UTC()
	fmt.Printf("<%v, %s, outgoing>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnEvent(s string) {
	logTime := time.Now().UTC()
	fmt.Printf("<%v, %s, event>\n  (%s)\n", logTime, l.prefix, s)
}

func (l screenLog) OnEventf(format string, a ...interface{}) {
	l.OnEvent(fmt.Sprintf(format, a...))
}

type screenLogFactory struct{}

func (screenLogFactory) Create() (Log, error) {
	log := screenLog{"GLOBAL"}
	return log, nil
}

func (screenLogFactory) CreateSessionLog(sessionID SessionID) (Log, error) {
	log := screenLog{sessionID.String()}
	return log, nil
}

// NewScreenLogFactory creates an instance of LogFactory that writes messages and events to stdout.
func NewScreenLogFactory() LogFactory {
	return screenLogFactory{}
}
