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

func (suite *MessageStoreTestSuite) TestMessageStore_SetNextMsgSeqNum_Refresh_IncrNextMsgSeqNum() {
	t := suite.T()

	// Given a MessageStore with the following sender and target seqnums
	suite.msgStore.SetNextSenderMsgSeqNum(867)
	suite.msgStore.SetNextTargetMsgSeqNum(5309)

	// When the store is refreshed from its backing store
	suite.msgStore.Refresh()

	// Then the sender and target seqnums should still be
	assert.Equal(t, 867, suite.msgStore.NextSenderMsgSeqNum())
	assert.Equal(t, 5309, suite.msgStore.NextTargetMsgSeqNum())

	// When the sender and target seqnums are incremented
	require.Nil(t, suite.msgStore.IncrNextSenderMsgSeqNum())
	require.Nil(t, suite.msgStore.IncrNextTargetMsgSeqNum())

	// Then the sender and target seqnums should be
	assert.Equal(t, 868, suite.msgStore.NextSenderMsgSeqNum())
	assert.Equal(t, 5310, suite.msgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	suite.msgStore.Refresh()

	// Then the sender and target seqnums should still be
	assert.Equal(t, 868, suite.msgStore.NextSenderMsgSeqNum())
	assert.Equal(t, 5310, suite.msgStore.NextTargetMsgSeqNum())
}

func (suite *MessageStoreTestSuite) TestMessageStore_Reset() {
	t := suite.T()

	// Given a MessageStore with the following sender and target seqnums
	suite.msgStore.SetNextSenderMsgSeqNum(1234)
	suite.msgStore.SetNextTargetMsgSeqNum(5678)

	// When the store is reset
	require.Nil(t, suite.msgStore.Reset())

	// Then the sender and target seqnums should be
	assert.Equal(t, 1, suite.msgStore.NextSenderMsgSeqNum())
	assert.Equal(t, 1, suite.msgStore.NextTargetMsgSeqNum())

	// When the store is refreshed from its backing store
	suite.msgStore.Refresh()

	// Then the sender and target seqnums should still be
	assert.Equal(t, 1, suite.msgStore.NextSenderMsgSeqNum())
	assert.Equal(t, 1, suite.msgStore.NextTargetMsgSeqNum())
}

func (suite *MessageStoreTestSuite) TestMessageStore_SaveMessage_GetMessage() {
	t := suite.T()

	// Given the following saved messages
	expectedMsgsBySeqNum := map[int]string{
		1: "In the frozen land of Nador",
		2: "they were forced to eat Robin's minstrels",
		3: "and there was much rejoicing",
	}
	for seqNum, msg := range expectedMsgsBySeqNum {
		require.Nil(t, suite.msgStore.SaveMessage(seqNum, []byte(msg)))
	}

	// When the messages are retrieved from the MessageStore
	actualMsgs, err := suite.msgStore.GetMessages(1, 3)
	require.Nil(t, err)

	// Then the messages should be
	require.Len(t, actualMsgs, 3)
	assert.Equal(t, expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	assert.Equal(t, expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	assert.Equal(t, expectedMsgsBySeqNum[3], string(actualMsgs[2]))

	// When the store is refreshed from its backing store
	suite.msgStore.Refresh()

	// And the messages are retrieved from the MessageStore
	actualMsgs, err = suite.msgStore.GetMessages(1, 3)
	require.Nil(t, err)

	// Then the messages should still be
	require.Len(t, actualMsgs, 3)
	assert.Equal(t, expectedMsgsBySeqNum[1], string(actualMsgs[0]))
	assert.Equal(t, expectedMsgsBySeqNum[2], string(actualMsgs[1]))
	assert.Equal(t, expectedMsgsBySeqNum[3], string(actualMsgs[2]))
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

func (suite *MessageStoreTestSuite) TestMessageStore_CreationTime() {
	assert.False(suite.T(), suite.msgStore.CreationTime().IsZero())

	t0 := time.Now()
	suite.msgStore.Reset()
	t1 := time.Now()
	require.True(suite.T(), suite.msgStore.CreationTime().After(t0))
	require.True(suite.T(), suite.msgStore.CreationTime().Before(t1))
}
