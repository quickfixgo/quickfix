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

package testsuite

import (
	"sort"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
	MsgStore quickfix.MessageStore
}

func (s *StoreTestSuite) TestMessageStoreSetNextMsgSeqNumRefreshIncrNextMsgSeqNum() {
	// Given a MessageStore with the following sender and target seqnums
	s.Require().Nil(s.MsgStore.SetNextSenderMsgSeqNum(867))
	s.Require().Nil(s.MsgStore.SetNextTargetMsgSeqNum(5309))

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(uint64(867), s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(uint64(5309), s.MsgStore.NextTargetMsgSeqNum())

	// When the sender and target seqnums are incremented
	s.Require().Nil(s.MsgStore.IncrNextSenderMsgSeqNum())
	s.Require().Nil(s.MsgStore.IncrNextTargetMsgSeqNum())

	// Then the sender and target seqnums should be
	s.Equal(uint64(868), s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(uint64(5310), s.MsgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(uint64(868), s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(uint64(5310), s.MsgStore.NextTargetMsgSeqNum())
}

func (s *StoreTestSuite) TestMessageStoreReset() {
	// Given a MessageStore with the following sender and target seqnums
	s.Require().Nil(s.MsgStore.SetNextSenderMsgSeqNum(1234))
	s.Require().Nil(s.MsgStore.SetNextTargetMsgSeqNum(5678))

	// When the store is reset
	s.Require().Nil(s.MsgStore.Reset())

	// Then the sender and target seqnums should be
	s.Equal(uint64(1), s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(uint64(1), s.MsgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(uint64(1), s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(uint64(1), s.MsgStore.NextTargetMsgSeqNum())
}

func (s *StoreTestSuite) fetchMessages(beginSeqNum, endSeqNum uint64) (msgs [][]byte) {
	s.T().Helper()

	// Fetch messages from the new iterator
	err := s.MsgStore.IterateMessages(beginSeqNum, endSeqNum, func(msg []byte) error {
		msgs = append(msgs, msg)
		return nil
	})
	s.Require().Nil(err)

	// Fetch messages from the old getter
	oldMsgs, err := s.MsgStore.GetMessages(beginSeqNum, endSeqNum)
	s.Require().Nil(err)

	// Ensure the output is the same
	s.Require().Len(msgs, len(oldMsgs))
	for idx, msg := range msgs {
		s.Require().EqualValues(msg, oldMsgs[idx])
	}
	return
}

func (s *StoreTestSuite) TestMessageStoreSaveMessageGetMessage() {
	// Given the following saved messages
	expectedMsgsBySeqNum := map[uint64]string{
		1: "In the frozen land of Nador",
		2: "they were forced to eat Robin's minstrels",
		3: "and there was much rejoicing",
	}
	var seqNums []uint64
	for seqNum := range expectedMsgsBySeqNum {
		seqNums = append(seqNums, seqNum)
	}
	sort.Slice(seqNums, func(i, j int) bool { return seqNums[i] < seqNums[j] })
	for _, seqNum := range seqNums {
		s.Require().Nil(s.MsgStore.SaveMessage(seqNum, []byte(expectedMsgsBySeqNum[seqNum])))
	}

	// When the messages are retrieved from the MessageStore
	actualMsgs := s.fetchMessages(1, 3)

	// Then the messages should be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// And the messages are retrieved from the MessageStore
	actualMsgs = s.fetchMessages(1, 3)

	// Then the messages should still be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))
}

func (s *StoreTestSuite) TestMessageStoreSaveMessageAndIncrementGetMessage() {
	s.Require().Nil(s.MsgStore.SetNextSenderMsgSeqNum(420))

	// Given the following saved messages
	expectedMsgsBySeqNum := map[uint64]string{
		1: "In the frozen land of Nador",
		2: "they were forced to eat Robin's minstrels",
		3: "and there was much rejoicing",
	}
	var seqNums []uint64
	for seqNum := range expectedMsgsBySeqNum {
		seqNums = append(seqNums, seqNum)
	}
	sort.Slice(seqNums, func(i, j int) bool { return seqNums[i] < seqNums[j] })
	for _, seqNum := range seqNums {
		s.Require().Nil(s.MsgStore.SaveMessageAndIncrNextSenderMsgSeqNum(seqNum, []byte(expectedMsgsBySeqNum[seqNum])))
	}
	s.Equal(uint64(423), s.MsgStore.NextSenderMsgSeqNum())

	// When the messages are retrieved from the MessageStore
	actualMsgs := s.fetchMessages(1, 3)

	// Then the messages should be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// And the messages are retrieved from the MessageStore
	actualMsgs = s.fetchMessages(1, 3)

	s.Equal(uint64(423), s.MsgStore.NextSenderMsgSeqNum())

	// Then the messages should still be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))
}

func (s *StoreTestSuite) TestMessageStoreGetMessagesEmptyStore() {
	// When messages are retrieved from an empty store
	messages := s.fetchMessages(1, 2)

	// Then no messages should be returned
	require.Empty(s.T(), messages, "Did not expect messages from empty store")
}

func (s *StoreTestSuite) TestMessageStoreGetMessagesVariousRanges() {
	t := s.T()

	// Given the following saved messages
	require.Nil(t, s.MsgStore.SaveMessage(1, []byte("hello")))
	require.Nil(t, s.MsgStore.SaveMessage(2, []byte("cruel")))
	require.Nil(t, s.MsgStore.SaveMessage(3, []byte("world")))

	// When the following requests are made to the store
	var testCases = []struct {
		beginSeqNo, endSeqNo uint64
		expectedBytes        [][]byte
	}{
		{beginSeqNo: 1, endSeqNo: 1, expectedBytes: [][]byte{[]byte("hello")}},
		{beginSeqNo: 1, endSeqNo: 2, expectedBytes: [][]byte{[]byte("hello"), []byte("cruel")}},
		{beginSeqNo: 1, endSeqNo: 3, expectedBytes: [][]byte{[]byte("hello"), []byte("cruel"), []byte("world")}},
		{beginSeqNo: 1, endSeqNo: 4, expectedBytes: [][]byte{[]byte("hello"), []byte("cruel"), []byte("world")}},
		{beginSeqNo: 2, endSeqNo: 3, expectedBytes: [][]byte{[]byte("cruel"), []byte("world")}},
		{beginSeqNo: 3, endSeqNo: 3, expectedBytes: [][]byte{[]byte("world")}},
		{beginSeqNo: 3, endSeqNo: 4, expectedBytes: [][]byte{[]byte("world")}},
		{beginSeqNo: 4, endSeqNo: 4, expectedBytes: [][]byte{}},
		{beginSeqNo: 4, endSeqNo: 10, expectedBytes: [][]byte{}},
	}

	// Then the returned messages should be
	for _, tc := range testCases {
		actualMsgs := s.fetchMessages(tc.beginSeqNo, tc.endSeqNo)
		require.Len(t, actualMsgs, len(tc.expectedBytes))
		for i, expectedMsg := range tc.expectedBytes {
			assert.Equal(t, string(expectedMsg), string(actualMsgs[i]))
		}
	}
}

func (s *StoreTestSuite) TestMessageStoreCreationTime() {
	s.False(s.MsgStore.CreationTime().IsZero())

	t0 := time.Now()
	s.Require().Nil(s.MsgStore.Reset())
	t1 := time.Now()
	s.Require().True(s.MsgStore.CreationTime().After(t0))
	s.Require().True(s.MsgStore.CreationTime().Before(t1))
}
