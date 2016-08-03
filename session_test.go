package quickfix

import (
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/enum"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func newFIXString(val string) *FIXString {
	s := FIXString(val)
	return &s
}

func TestSession_CheckCorrectCompID(t *testing.T) {
	session := session{}
	session.sessionID.TargetCompID = "TAR"
	session.sessionID.SenderCompID = "SND"

	var testCases = []struct {
		senderCompID *FIXString
		targetCompID *FIXString
		returnsError bool
		rejectReason int
	}{
		{returnsError: true, rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: newFIXString("TAR"),
			returnsError: true,
			rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: newFIXString("TAR"),
			targetCompID: newFIXString("JCD"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: newFIXString("JCD"),
			targetCompID: newFIXString("SND"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: newFIXString("TAR"),
			targetCompID: newFIXString("SND"),
			returnsError: false},
	}

	for _, tc := range testCases {
		msg := NewMessage()

		if tc.senderCompID != nil {
			msg.Header.SetField(tagSenderCompID, tc.senderCompID)
		}

		if tc.targetCompID != nil {
			msg.Header.SetField(tagTargetCompID, tc.targetCompID)
		}

		rej := session.checkCompID(msg)

		if !tc.returnsError {
			require.Nil(t, rej)
			continue
		}

		require.NotNil(t, rej)
		assert.Equal(t, tc.rejectReason, rej.RejectReason())
	}
}

func TestSession_CheckBeginString(t *testing.T) {
	session := session{
		sessionID: SessionID{BeginString: "FIX.4.2"},
	}

	msg := NewMessage()

	//wrong value
	msg.Header.SetField(tagBeginString, FIXString("FIX.4.4"))
	err := session.checkBeginString(msg)
	require.NotNil(t, err, "wrong begin string should return error")
	assert.IsType(t, incorrectBeginString{}, err)

	msg.Header.SetField(tagBeginString, FIXString(session.sessionID.BeginString))
	err = session.checkBeginString(msg)
	assert.Nil(t, err)
}

func TestSession_CheckTargetTooHigh(t *testing.T) {
	store := new(memoryStore)
	session := session{store: store}

	msg := NewMessage()
	require.Nil(t, store.SetNextTargetMsgSeqNum(45))

	err := session.checkTargetTooHigh(msg)
	require.NotNil(t, err, "missing sequence number should return error")
	assert.Equal(t, rejectReasonRequiredTagMissing, err.RejectReason())

	msg.Header.SetField(tagMsgSeqNum, FIXInt(47))
	err = session.checkTargetTooHigh(msg)
	require.NotNil(t, err, "sequence number too high should return an error")
	assert.IsType(t, targetTooHigh{}, err)

	//spot on
	msg.Header.SetField(tagMsgSeqNum, FIXInt(45))
	err = session.checkTargetTooHigh(msg)
	assert.Nil(t, err)
}

func TestSession_CheckSendingTime(t *testing.T) {
	session := session{}
	msg := NewMessage()

	err := session.checkSendingTime(msg)
	require.NotNil(t, err, "sending time is a required field")
	assert.Equal(t, rejectReasonRequiredTagMissing, err.RejectReason())

	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = session.checkSendingTime(msg)
	require.NotNil(t, err, "sending time too late should give error")
	assert.Equal(t, rejectReasonSendingTimeAccuracyProblem, err.RejectReason())

	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = session.checkSendingTime(msg)
	require.NotNil(t, err, "future sending time should give error")
	assert.Equal(t, rejectReasonSendingTimeAccuracyProblem, err.RejectReason())

	sendingTime = time.Now()
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})

	err = session.checkSendingTime(msg)
	assert.Nil(t, err, "sending time should be ok")
}

func TestSession_CheckTargetTooLow(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	session := session{store: store}

	msg := NewMessage()
	require.Nil(t, store.SetNextTargetMsgSeqNum(45))

	err := session.checkTargetTooLow(msg)
	require.NotNil(t, err, "sequence number is required")
	assert.Equal(t, rejectReasonRequiredTagMissing, err.RejectReason())

	//too low
	msg.Header.SetField(tagMsgSeqNum, FIXInt(43))
	err = session.checkTargetTooLow(msg)
	require.NotNil(t, err, "sequence number too low should return error")
	assert.IsType(t, targetTooLow{}, err)

	//spot on
	msg.Header.SetField(tagMsgSeqNum, FIXInt(45))
	err = session.checkTargetTooLow(msg)
	assert.Nil(t, err)
}

type SessionSendTestSuite struct {
	suite.Suite

	*messageFactory
	*mockApp
	receiver *mockSessionReceiver
	*session
}

func (suite *SessionSendTestSuite) SetupTest() {
	suite.mockApp = new(mockApp)
	suite.messageFactory = new(messageFactory)
	suite.receiver = newMockSessionReceiver()
	suite.session = &session{
		store:        new(memoryStore),
		application:  suite,
		log:          nullLog{},
		messageOut:   suite.receiver.sendChannel,
		sessionState: inSession{},
		sessionID:    SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"},
	}
}

func TestSessionSendTestSuite(t *testing.T) {
	suite.Run(t, new(SessionSendTestSuite))
}

func (suite *SessionSendTestSuite) shouldNotPersistMessage() {
	suite.Equal(1, suite.store.NextSenderMsgSeqNum(), "The next sender sequence number should not be incremented")
	persistedMessages, err := suite.store.GetMessages(1, 1)
	suite.Nil(err)
	suite.Len(persistedMessages, 0, "The message should not be persisted")
}

func (suite *SessionSendTestSuite) shouldPersistMessage() {
	suite.Equal(2, suite.store.NextSenderMsgSeqNum(), "The next sender sequence number should be incremented")
	persistedMessages, err := suite.store.GetMessages(1, 1)
	suite.Nil(err)
	suite.Len(persistedMessages, 1, "The message should be persisted")
}

func (suite *SessionSendTestSuite) shouldPersistMessageWithSequenceNum(expectedSeqNum int) {
	suite.Equal(expectedSeqNum+1, suite.store.NextSenderMsgSeqNum(), "The next sender sequence number should be incremented")
	persistedMessages, err := suite.store.GetMessages(expectedSeqNum, expectedSeqNum)
	suite.Nil(err)
	suite.Len(persistedMessages, 1, "The message should be persisted")
}

func (suite *SessionSendTestSuite) shouldBeType(msg Message, expectedMsgType string) {
	actual, err := msg.Header.GetString(tagMsgType)
	suite.Nil(err, "Message doesn't have message type")
	suite.Equal(expectedMsgType, actual)
}

func (suite *SessionSendTestSuite) shouldSendMessage() {
	suite.NotNil(suite.receiver.LastMessage(), "The message should have been sent")
}
func (suite *SessionSendTestSuite) shouldNotSendMessage() {
	suite.Nil(suite.receiver.LastMessage(), "The message should not have been sent")
}

func (suite *SessionSendTestSuite) shouldSendMessages(cnt int) {
	for i := 0; i < cnt; i++ {
		suite.shouldSendMessage()
	}

	suite.shouldNotSendMessage()
}

func (suite *SessionSendTestSuite) sentMessageShouldBe(sentMsg []byte, msg Message) {
	expectedBytes, _ := msg.Build()
	suite.Equal(string(expectedBytes), string(sentMsg))
}

func (suite *SessionSendTestSuite) TestQueueForSendAppMessage() {
	suite.mockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldPersistMessage()
	suite.shouldNotSendMessage()
}

func (suite *SessionSendTestSuite) TestQueueForSendDoNotSendAppMessage() {
	suite.mockApp.On("ToApp").Return(DoNotSend)
	suite.Equal(DoNotSend, suite.queueForSend(suite.NewOrderSingle()))

	suite.AssertExpectations(suite.T())
	suite.shouldNotPersistMessage()
	suite.shouldNotSendMessage()

	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.send(suite.Heartbeat()))
	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldSendMessages(1)
}

func (suite *SessionSendTestSuite) TestQueueForSendAdminMessage() {
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldPersistMessage()
	suite.shouldNotSendMessage()
}

func (suite *SessionSendTestSuite) TestSendAppMessage() {
	suite.mockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldPersistMessage()
	suite.shouldSendMessage()
}

func (suite *SessionSendTestSuite) TestSendAppDoNotSendMessage() {
	suite.mockApp.On("ToApp").Return(DoNotSend)
	suite.Equal(DoNotSend, suite.send(suite.NewOrderSingle()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldNotPersistMessage()
	suite.shouldNotSendMessage()
}

func (suite *SessionSendTestSuite) TestSendAdminMessage() {
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.send(suite.Heartbeat()))
	suite.mockApp.AssertExpectations(suite.T())

	suite.shouldPersistMessage()
	suite.shouldSendMessage()
}

func (suite *SessionSendTestSuite) TestSendFlushesQueue() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldNotSendMessage()

	suite.mockApp.On("ToApp").Return(nil)
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldSendMessages(3)
}

func (suite *SessionSendTestSuite) TestSendNotLoggedOn() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldNotSendMessage()

	var tests = []sessionState{logoutState{}, latentState{}, logonState{}}

	for _, test := range tests {
		suite.mockApp.On("ToApp").Return(nil)
		suite.sessionState = test
		require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
		suite.mockApp.AssertExpectations(suite.T())
		suite.shouldNotSendMessage()
	}
}

func (suite *SessionSendTestSuite) TestDropAndSendAdminMessage() {
	suite.mockApp.On("ToAdmin")
	suite.Require().Nil(suite.dropAndSend(suite.Heartbeat(), false))
	suite.mockApp.AssertExpectations(suite.T())

	suite.shouldPersistMessage()
	suite.shouldSendMessage()
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueue() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.mockApp.AssertExpectations(suite.T())

	suite.shouldNotSendMessage()

	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.dropAndSend(suite.Logon(), false))
	suite.mockApp.AssertExpectations(suite.T())

	msg := suite.mockApp.lastToAdmin
	suite.shouldBeType(msg, enum.MsgType_LOGON)

	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	suite.Nil(err)
	suite.Equal(3, seqNum)

	//only one message sent
	sentMsgBytes := suite.receiver.LastMessage()
	suite.NotNil(sentMsgBytes)
	suite.shouldNotSendMessage()

	suite.sentMessageShouldBe(sentMsgBytes, msg)
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueueWithReset() {
	suite.mockApp.On("ToApp").Return(nil)
	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.mockApp.AssertExpectations(suite.T())
	suite.shouldNotSendMessage()

	suite.mockApp.On("ToAdmin")
	require.Nil(suite.T(), suite.dropAndSend(suite.Logon(), true))
	suite.mockApp.AssertExpectations(suite.T())
	msg := suite.mockApp.lastToAdmin

	suite.shouldBeType(msg, enum.MsgType_LOGON)
	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	suite.Nil(err)
	suite.Equal(1, seqNum)

	//only one message sent
	sentMsgBytes := suite.receiver.LastMessage()
	suite.NotNil(sentMsgBytes)
	suite.shouldNotSendMessage()

	suite.sentMessageShouldBe(sentMsgBytes, msg)
}
