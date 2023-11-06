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
	s.Equal(867, s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(5309, s.MsgStore.NextTargetMsgSeqNum())

	// When the sender and target seqnums are incremented
	s.Require().Nil(s.MsgStore.IncrNextSenderMsgSeqNum())
	s.Require().Nil(s.MsgStore.IncrNextTargetMsgSeqNum())

	// Then the sender and target seqnums should be
	s.Equal(868, s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(5310, s.MsgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(868, s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(5310, s.MsgStore.NextTargetMsgSeqNum())
}

func (s *StoreTestSuite) TestMessageStoreReset() {
	// Given a MessageStore with the following sender and target seqnums
	s.Require().Nil(s.MsgStore.SetNextSenderMsgSeqNum(1234))
	s.Require().Nil(s.MsgStore.SetNextTargetMsgSeqNum(5678))

	// When the store is reset
	s.Require().Nil(s.MsgStore.Reset())

	// Then the sender and target seqnums should be
	s.Equal(1, s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(1, s.MsgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(1, s.MsgStore.NextSenderMsgSeqNum())
	s.Equal(1, s.MsgStore.NextTargetMsgSeqNum())
}

func (s *StoreTestSuite) TestMessageStoreSaveMessageGetMessage() {
	// Given the following saved messages
	expectedMsgsBySeqNum := map[int]string{
		1: "In the frozen land of Nador",
		2: "they were forced to eat Robin's minstrels",
		3: "and there was much rejoicing",
	}
	for seqNum, msg := range expectedMsgsBySeqNum {
		s.Require().Nil(s.MsgStore.SaveMessage(seqNum, []byte(msg)))
	}

	// When the messages are retrieved from the MessageStore
	actualMsgs, err := s.MsgStore.GetMessages(1, 3)
	s.Require().Nil(err)

	// Then the messages should be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// And the messages are retrieved from the MessageStore
	actualMsgs, err = s.MsgStore.GetMessages(1, 3)
	s.Require().Nil(err)

	// Then the messages should still be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))
}

func (s *StoreTestSuite) TestMessageStoreSaveMessageAndIncrementGetMessage() {
	s.Require().Nil(s.MsgStore.SetNextSenderMsgSeqNum(420))

	// Given the following saved messages
	expectedMsgsBySeqNum := map[int]string{
		1: "In the frozen land of Nador",
		2: "they were forced to eat Robin's minstrels",
		3: "and there was much rejoicing",
	}
	for seqNum, msg := range expectedMsgsBySeqNum {
		s.Require().Nil(s.MsgStore.SaveMessageAndIncrNextSenderMsgSeqNum(seqNum, []byte(msg)))
	}
	s.Equal(423, s.MsgStore.NextSenderMsgSeqNum())

	// When the messages are retrieved from the MessageStore
	actualMsgs, err := s.MsgStore.GetMessages(1, 3)
	s.Require().Nil(err)

	// Then the messages should be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))

	// When the store is refreshed from its backing store
	s.Require().Nil(s.MsgStore.Refresh())

	// And the messages are retrieved from the MessageStore
	actualMsgs, err = s.MsgStore.GetMessages(1, 3)
	s.Require().Nil(err)

	s.Equal(423, s.MsgStore.NextSenderMsgSeqNum())

	// Then the messages should still be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))
}

func (s *StoreTestSuite) TestMessageStoreGetMessagesEmptyStore() {
	// When messages are retrieved from an empty store
	messages, err := s.MsgStore.GetMessages(1, 2)
	require.Nil(s.T(), err)

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
		beginSeqNo, endSeqNo int
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
		actualMsgs, err := s.MsgStore.GetMessages(tc.beginSeqNo, tc.endSeqNo)
		require.Nil(t, err)
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
