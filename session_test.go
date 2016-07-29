package quickfix

import (
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/enum"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func buildMessage() Message {
	builder := NewMessage()
	builder.Header.SetField(tagBeginString, FIXString(enum.BeginStringFIX40))
	builder.Header.SetField(tagMsgType, FIXString("D"))
	return builder
}

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
		builder := buildMessage()

		if tc.senderCompID != nil {
			builder.Header.SetField(tagSenderCompID, tc.senderCompID)
		}

		if tc.targetCompID != nil {
			builder.Header.SetField(tagTargetCompID, tc.targetCompID)
		}

		msgBytes, _ := builder.Build()
		msg, _ := ParseMessage(msgBytes)
		err := session.checkCompID(msg)

		if err == nil {
			if tc.returnsError {
				t.Error("expected error")
			}

			return
		}

		if !tc.returnsError {
			t.Fatal("unexpected error", err)
		}

		if err.RejectReason() != tc.rejectReason {
			t.Errorf("expected %v got %v", tc.rejectReason, err.RejectReason())
		}
	}
}

func TestSession_CheckBeginString(t *testing.T) {
	session := session{
		sessionID: SessionID{BeginString: "FIX.4.2"},
	}

	builder := buildMessage()

	//wrong value
	builder.Header.SetField(tagBeginString, FIXString("FIX.4.4"))
	msgBytes, _ := builder.Build()
	msg, _ := ParseMessage(msgBytes)

	err := session.checkBeginString(msg)
	if err == nil {
		t.Error("Expected Error")
	}
	_ = err.(incorrectBeginString)

	builder.Header.SetField(tagBeginString, FIXString(session.sessionID.BeginString))
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)

	err = session.checkBeginString(msg)

	if err != nil {
		t.Error("Unexpected error", err)
	}
}

func TestSession_CheckTargetTooHigh(t *testing.T) {
	store := new(memoryStore)
	session := session{store: store}
	builder := buildMessage()
	msgBytes, _ := builder.Build()
	msg, _ := ParseMessage(msgBytes)

	store.SetNextTargetMsgSeqNum(45)

	//missing seq number
	err := session.checkTargetTooHigh(msg)

	if err == nil {
		t.Error("Expected error")
	}

	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Expected required tag missing, got", err.RejectReason())
	}

	//too low
	builder.Header.SetField(tagMsgSeqNum, FIXInt(47))
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)
	err = session.checkTargetTooHigh(msg)

	if err == nil {
		t.Error("Expected error")
	}
	_ = err.(targetTooHigh)

	//spot on
	builder.Header.SetField(tagMsgSeqNum, FIXInt(45))
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)

	err = session.checkTargetTooHigh(msg)
	if err != nil {
		t.Error("Unexpected error", err)
	}
}

func TestSession_CheckSendingTime(t *testing.T) {
	session := session{}
	builder := buildMessage()
	msgBytes, _ := builder.Build()
	msg, _ := ParseMessage(msgBytes)

	//missing sending time
	err := session.checkSendingTime(msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Reject reason not expected, got ", err.RejectReason())
	}

	//sending time too late
	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	builder.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)

	err = session.checkSendingTime(msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonSendingTimeAccuracyProblem {
		t.Error("Reject reason not expected, got ", err.RejectReason())
	}

	//future sending time
	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	builder.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)

	err = session.checkSendingTime(msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonSendingTimeAccuracyProblem {
		t.Error("Reject reason not expected, got ", err.RejectReason())
	}

	//sending time ok
	sendingTime = time.Now()
	builder.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)

	err = session.checkSendingTime(msg)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestSession_CheckTargetTooLow(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	session := session{store: store}

	builder := buildMessage()
	msgBytes, _ := builder.Build()
	msg, _ := ParseMessage(msgBytes)

	store.SetNextTargetMsgSeqNum(45)

	//missing seq number
	err := session.checkTargetTooLow(msg)
	if err == nil {
		t.Error("Expected error")
	}

	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Unexpected reject reason", err.RejectReason())
	}

	//too low
	builder.Header.SetField(tagMsgSeqNum, FIXInt(43))
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)

	err = session.checkTargetTooLow(msg)
	if err == nil {
		t.Error("Expected error")
	}
	_ = err.(targetTooLow)

	//spot on
	builder.Header.SetField(tagMsgSeqNum, FIXInt(45))
	msgBytes, _ = builder.Build()
	msg, _ = ParseMessage(msgBytes)

	err = session.checkTargetTooLow(msg)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

type TestClient struct {
	adminMessages []Message
	appMessages   []Message
}

func (e *TestClient) OnCreate(sessionID SessionID) {
}

func (e *TestClient) OnLogon(sessionID SessionID) {
}

func (e *TestClient) OnLogout(sessionID SessionID) {
}

func (e *TestClient) FromAdmin(msg Message, sessionID SessionID) (reject MessageRejectError) {
	return nil
}

func (e *TestClient) ToAdmin(msg Message, sessionID SessionID) {
	e.adminMessages = append(e.adminMessages, msg)
}

func (e *TestClient) ToApp(msg Message, sessionID SessionID) (err error) {
	e.appMessages = append(e.appMessages, msg)
	return nil
}

func (e *TestClient) FromApp(msg Message, sessionID SessionID) (reject MessageRejectError) {
	return nil
}

type SessionSendTestSuite struct {
	suite.Suite

	*TestClient
	*session
	sendChannel chan []byte
}

func (suite *SessionSendTestSuite) SetupTest() {
	suite.sendChannel = make(chan []byte, 10)
	suite.TestClient = new(TestClient)
	suite.session = &session{
		store:       new(memoryStore),
		application: suite,
		log:         nullLog{},
		messageOut:  suite.sendChannel,
	}
}

func (suite *SessionSendTestSuite) sentMessage() (msg []byte) {
	select {
	case msg = <-suite.sendChannel:
	default:
	}
	return
}

func TestSessionSendTestSuite(t *testing.T) {
	suite.Run(t, new(SessionSendTestSuite))
}

func (suite *SessionSendTestSuite) NewOrderSingle() Message {
	msg := buildMessage()
	msg.Header.SetField(tagMsgType, FIXString("D"))
	return msg
}

func (suite *SessionSendTestSuite) Heartbeat() Message {
	msg := buildMessage()
	msg.Header.SetField(tagMsgType, FIXString("0"))
	return msg
}

func (suite *SessionSendTestSuite) Logon() Message {
	msg := buildMessage()
	msg.Header.SetField(tagMsgType, FIXString("A"))
	return msg
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

func (suite *SessionSendTestSuite) shouldCallToApp() {
	suite.Len(suite.appMessages, 1, "ToApp should be called")
	suite.appMessages = []Message{}
}

func (suite *SessionSendTestSuite) shouldCallToAdmin() {
	suite.Len(suite.adminMessages, 1, "ToAdmin should be called")
	suite.adminMessages = []Message{}
}

func (suite *SessionSendTestSuite) shouldBeType(msg Message, expectedMsgType string) {
	actual, err := msg.Header.GetString(tagMsgType)
	suite.Nil(err, "Message doesn't have message type")
	suite.Equal(expectedMsgType, actual)
}

func (suite *SessionSendTestSuite) lastToApp() Message {
	suite.NotEmpty(suite.appMessages, "ToApp should be called")
	return suite.appMessages[len(suite.appMessages)-1]
}

func (suite *SessionSendTestSuite) lastToAdmin() Message {
	suite.NotEmpty(suite.appMessages, "ToAdmin should be called")
	return suite.adminMessages[len(suite.adminMessages)-1]
}

func (suite *SessionSendTestSuite) shouldSendMessage() {
	suite.NotNil(suite.sentMessage(), "The message should have been sent")
}
func (suite *SessionSendTestSuite) shouldNotSendMessage() {
	suite.Nil(suite.sentMessage(), "The message should not have been sent")
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
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))

	suite.shouldCallToApp()
	suite.shouldPersistMessage()
	suite.shouldNotSendMessage()
}

func (suite *SessionSendTestSuite) TestQueueForSendAdminMessage() {
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))

	suite.shouldCallToAdmin()
	suite.shouldPersistMessage()
	suite.shouldNotSendMessage()
}

func (suite *SessionSendTestSuite) TestSendAppMessage() {
	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))

	suite.shouldCallToApp()
	suite.shouldPersistMessage()
	suite.shouldSendMessage()
}

func (suite *SessionSendTestSuite) TestSendAdminMessage() {
	require.Nil(suite.T(), suite.send(suite.Heartbeat()))

	suite.shouldCallToAdmin()
	suite.shouldPersistMessage()
	suite.shouldSendMessage()
}

func (suite *SessionSendTestSuite) TestSendFlushesQueue() {
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.shouldNotSendMessage()

	require.Nil(suite.T(), suite.send(suite.NewOrderSingle()))
	suite.shouldSendMessages(3)
}

func (suite *SessionSendTestSuite) TestDropAndSendAdminMessage() {
	require.Nil(suite.T(), suite.dropAndSend(suite.Heartbeat(), false))

	suite.shouldCallToAdmin()
	suite.shouldPersistMessage()
	suite.shouldSendMessage()
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueue() {
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.shouldNotSendMessage()

	require.Nil(suite.T(), suite.dropAndSend(suite.Logon(), false))
	msg := suite.lastToAdmin()
	suite.shouldBeType(msg, enum.MsgType_LOGON)

	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	suite.Nil(err)
	suite.Equal(3, seqNum)

	//only one message sent
	sentMsgBytes := suite.sentMessage()
	suite.NotNil(sentMsgBytes)
	suite.shouldNotSendMessage()

	suite.sentMessageShouldBe(sentMsgBytes, msg)
}

func (suite *SessionSendTestSuite) TestDropAndSendDropsQueueWithReset() {
	require.Nil(suite.T(), suite.queueForSend(suite.NewOrderSingle()))
	require.Nil(suite.T(), suite.queueForSend(suite.Heartbeat()))
	suite.shouldNotSendMessage()

	require.Nil(suite.T(), suite.dropAndSend(suite.Logon(), true))
	msg := suite.lastToAdmin()

	suite.shouldBeType(msg, enum.MsgType_LOGON)
	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	suite.Nil(err)
	suite.Equal(1, seqNum)

	//only one message sent
	sentMsgBytes := suite.sentMessage()
	suite.NotNil(sentMsgBytes)
	suite.shouldNotSendMessage()

	suite.sentMessageShouldBe(sentMsgBytes, msg)
}
