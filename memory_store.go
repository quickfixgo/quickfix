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

	"github.com/pkg/errors"
)

type memoryStore struct {
	senderMsgSeqNum, targetMsgSeqNum uint64
	creationTime                     time.Time
	messageMap                       map[uint64][]byte
}

func (store *memoryStore) NextSenderMsgSeqNum() uint64 {
	return store.senderMsgSeqNum + 1
}

func (store *memoryStore) NextTargetMsgSeqNum() uint64 {
	return store.targetMsgSeqNum + 1
}

func (store *memoryStore) IncrNextSenderMsgSeqNum() error {
	store.senderMsgSeqNum++
	return nil
}

func (store *memoryStore) IncrNextTargetMsgSeqNum() error {
	store.targetMsgSeqNum++
	return nil
}

func (store *memoryStore) SetNextSenderMsgSeqNum(nextSeqNum uint64) error {
	store.senderMsgSeqNum = nextSeqNum - 1
	return nil
}
func (store *memoryStore) SetNextTargetMsgSeqNum(nextSeqNum uint64) error {
	store.targetMsgSeqNum = nextSeqNum - 1
	return nil
}

func (store *memoryStore) CreationTime() time.Time {
	return store.creationTime
}

func (store *memoryStore) SetCreationTime(t time.Time) {
	store.creationTime = t
}

func (store *memoryStore) Reset() error {
	store.senderMsgSeqNum = 0
	store.targetMsgSeqNum = 0
	store.creationTime = time.Now()
	store.messageMap = nil
	return nil
}

func (store *memoryStore) Refresh() error {
	// NOP, nothing to refresh.
	return nil
}

func (store *memoryStore) Close() error {
	// NOP, nothing to close.
	return nil
}

func (store *memoryStore) SaveMessage(seqNum uint64, msg []byte) error {
	if store.messageMap == nil {
		store.messageMap = make(map[uint64][]byte)
	}

	store.messageMap[seqNum] = msg
	return nil
}

func (store *memoryStore) SaveMessageAndIncrNextSenderMsgSeqNum(seqNum uint64, msg []byte) error {
	err := store.SaveMessage(seqNum, msg)
	if err != nil {
		return err
	}
	return store.IncrNextSenderMsgSeqNum()
}

func (store *memoryStore) IterateMessages(beginSeqNum, endSeqNum uint64, cb func([]byte) error) error {
	for seqNum := beginSeqNum; seqNum <= endSeqNum; seqNum++ {
		if m, ok := store.messageMap[seqNum]; ok {
			if err := cb(m); err != nil {
				return err
			}
		}
	}
	return nil
}

func (store *memoryStore) GetMessages(beginSeqNum, endSeqNum uint64) ([][]byte, error) {
	var msgs [][]byte
	err := store.IterateMessages(beginSeqNum, endSeqNum, func(m []byte) error {
		msgs = append(msgs, m)
		return nil
	})
	return msgs, err
}

type memoryStoreFactory struct{}

func (f memoryStoreFactory) Create(_ SessionID) (MessageStore, error) {
	m := new(memoryStore)
	if err := m.Reset(); err != nil {
		return m, errors.Wrap(err, "reset")
	}
	return m, nil
}

// NewMemoryStoreFactory returns a MessageStoreFactory instance that created in-memory MessageStores.
func NewMemoryStoreFactory() MessageStoreFactory { return memoryStoreFactory{} }
