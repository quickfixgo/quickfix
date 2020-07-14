package quickfix

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MessageStoreTestSuite is the suite of all tests that should be run against all MessageStore implementations
type MessageStoreTestSuite struct {
	suite.Suite
	msgStore MessageStore
}

// MemoryStoreTestSuite runs all tests in the MessageStoreTestSuite against the MemoryStore implementation
type MemoryStoreTestSuite struct {
	MessageStoreTestSuite
}

func (suite *MemoryStoreTestSuite) SetupTest() {
	var err error
	suite.msgStore, err = NewMemoryStoreFactory().Create(SessionID{})
	require.Nil(suite.T(), err)
}

func TestMemoryStoreTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryStoreTestSuite))
}

func (s *MessageStoreTestSuite) TestMessageStore_SetNextMsgSeqNum_Refresh_IncrNextMsgSeqNum() {
	// Given a MessageStore with the following sender and target seqnums
	s.Require().Nil(s.msgStore.SetNextSenderMsgSeqNum(867))
	s.Require().Nil(s.msgStore.SetNextTargetMsgSeqNum(5309))

	// When the store is refreshed from its backing store
	s.Require().Nil(s.msgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(867, s.msgStore.NextSenderMsgSeqNum())
	s.Equal(5309, s.msgStore.NextTargetMsgSeqNum())

	// When the sender and target seqnums are incremented
	s.Require().Nil(s.msgStore.IncrNextSenderMsgSeqNum())
	s.Require().Nil(s.msgStore.IncrNextTargetMsgSeqNum())

	// Then the sender and target seqnums should be
	s.Equal(868, s.msgStore.NextSenderMsgSeqNum())
	s.Equal(5310, s.msgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	s.Require().Nil(s.msgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(868, s.msgStore.NextSenderMsgSeqNum())
	s.Equal(5310, s.msgStore.NextTargetMsgSeqNum())
}

func (s *MessageStoreTestSuite) TestMessageStore_Reset() {
	// Given a MessageStore with the following sender and target seqnums
	s.Require().Nil(s.msgStore.SetNextSenderMsgSeqNum(1234))
	s.Require().Nil(s.msgStore.SetNextTargetMsgSeqNum(5678))

	// When the store is reset
	s.Require().Nil(s.msgStore.Reset())

	// Then the sender and target seqnums should be
	s.Equal(1, s.msgStore.NextSenderMsgSeqNum())
	s.Equal(1, s.msgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	s.Require().Nil(s.msgStore.Refresh())

	// Then the sender and target seqnums should still be
	s.Equal(1, s.msgStore.NextSenderMsgSeqNum())
	s.Equal(1, s.msgStore.NextTargetMsgSeqNum())
}

func (s *MessageStoreTestSuite) TestMessageStore_SaveMessage_GetMessage() {
	// Given the following saved messages
	expectedMsgsBySeqNum := map[int]string{
		1: "In the frozen land of Nador",
		2: "they were forced to eat Robin's minstrels",
		3: "and there was much rejoicing",
	}
	for seqNum, msg := range expectedMsgsBySeqNum {
		s.Require().Nil(s.msgStore.SaveMessage(seqNum, []byte(msg)))
	}

	// When the messages are retrieved from the MessageStore
	actualMsgs, err := s.msgStore.GetMessages(1, 3)
	s.Require().Nil(err)

	// Then the messages should be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))

	// When the store is refreshed from its backing store
	s.Require().Nil(s.msgStore.Refresh())

	// And the messages are retrieved from the MessageStore
	actualMsgs, err = s.msgStore.GetMessages(1, 3)
	s.Require().Nil(err)

	// Then the messages should still be
	s.Require().Len(actualMsgs, 3)
	s.Equal(expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	s.Equal(expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	s.Equal(expectedMsgsBySeqNum[3], string(actualMsgs[2]))
}

func (suite *MessageStoreTestSuite) TestMessageStore_GetMessages_EmptyStore() {
	// When messages are retrieved from an empty store
	messages, err := suite.msgStore.GetMessages(1, 2)
	require.Nil(suite.T(), err)

	// Then no messages should be returned
	require.Empty(suite.T(), messages, "Did not expect messages from empty store")
}

func (suite *MessageStoreTestSuite) TestMessageStore_GetMessages_VariousRanges() {
	t := suite.T()

	// Given the following saved messages
	require.Nil(t, suite.msgStore.SaveMessage(1, []byte("hello")))
	require.Nil(t, suite.msgStore.SaveMessage(2, []byte("cruel")))
	require.Nil(t, suite.msgStore.SaveMessage(3, []byte("world")))

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
		actualMsgs, err := suite.msgStore.GetMessages(tc.beginSeqNo, tc.endSeqNo)
		require.Nil(t, err)
		require.Len(t, actualMsgs, len(tc.expectedBytes))
		for i, expectedMsg := range tc.expectedBytes {
			assert.Equal(t, string(expectedMsg), string(actualMsgs[i]))
		}
	}
}

func (s *MessageStoreTestSuite) TestMessageStore_CreationTime() {
	s.False(s.msgStore.CreationTime().IsZero())

	t0 := time.Now()
	s.Require().Nil(s.msgStore.Reset())
	t1 := time.Now()
	s.Require().True(s.msgStore.CreationTime().After(t0))
	s.Require().True(s.msgStore.CreationTime().Before(t1))
}
