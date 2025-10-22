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
	"time"
)

// The MessageStore interface provides methods to record and retrieve messages for resend purposes.
type MessageStore interface {
	NextSenderMsgSeqNum() int
	NextTargetMsgSeqNum() int

	IncrNextSenderMsgSeqNum() error
	IncrNextTargetMsgSeqNum() error

	SetNextSenderMsgSeqNum(next int) error
	SetNextTargetMsgSeqNum(next int) error

	CreationTime() time.Time
	SetCreationTime(time.Time)

	SaveMessage(seqNum int, msg []byte) error
	SaveMessageAndIncrNextSenderMsgSeqNum(seqNum int, msg []byte) error
	GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error)

	Refresh() error
	Reset() error

	Close() error
}

// The MessageStoreFactory interface is used by session to create a session specific message store.
type MessageStoreFactory interface {
	Create(sessionID SessionID) (MessageStore, error)
}
