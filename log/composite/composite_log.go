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

package composite

import (
	"github.com/quickfixgo/quickfix"
)

type compositeLog struct {
	logs []quickfix.Log
}

func (l compositeLog) OnIncoming(s []byte) {
	for _, log := range l.logs {
		log.OnIncoming(s)
	}
}

func (l compositeLog) OnOutgoing(s []byte) {
	for _, log := range l.logs {
		log.OnOutgoing(s)
	}
}

func (l compositeLog) OnEvent(s string) {
	for _, log := range l.logs {
		log.OnEvent(s)
	}
}

func (l compositeLog) OnEventf(format string, a ...interface{}) {
	for _, log := range l.logs {
		log.OnEventf(format, a)
	}
}

type compositeLogFactory struct {
	logFactories []quickfix.LogFactory
}

func (clf compositeLogFactory) Create() (quickfix.Log, error) {
	logs := []quickfix.Log{}
	for _, lf := range clf.logFactories {
		log, err := lf.Create()
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return compositeLog{logs}, nil
}

func (clf compositeLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (quickfix.Log, error) {
	logs := []quickfix.Log{}
	for _, lf := range clf.logFactories {
		log, err := lf.CreateSessionLog(sessionID)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return compositeLog{logs}, nil
}

// NewLogFactory creates an instance of LogFactory that writes messages and events to stdout.
func NewLogFactory(logfactories []quickfix.LogFactory) quickfix.LogFactory {
	return compositeLogFactory{logfactories}
}
