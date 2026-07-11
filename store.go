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
	NextSenderMsgSeqNum() uint64
	NextTargetMsgSeqNum() uint64

	IncrNextSenderMsgSeqNum() error
	IncrNextTargetMsgSeqNum() error

	SetNextSenderMsgSeqNum(next uint64) error
	SetNextTargetMsgSeqNum(next uint64) error

	CreationTime() time.Time
	SetCreationTime(time.Time)

	SaveMessage(seqNum uint64, msg []byte) error
	SaveMessageAndIncrNextSenderMsgSeqNum(seqNum uint64, msg []byte) error
	GetMessages(beginSeqNum, endSeqNum uint64) ([][]byte, error)
	IterateMessages(beginSeqNum, endSeqNum uint64, cb func([]byte) error) error

	Refresh() error
	Reset() error

	Close() error
}

// The MessageStoreFactory interface is used by session to create a session specific message store.
type MessageStoreFactory interface {
	Create(sessionID SessionID) (MessageStore, error)
}
