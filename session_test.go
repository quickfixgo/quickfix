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
	adminCalled int
	appCalled   int
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
	e.adminCalled = e.adminCalled + 1
}

func (e *TestClient) ToApp(msg Message, sessionID SessionID) (err error) {
	e.appCalled = e.appCalled + 1
	return nil
}

func (e *TestClient) FromApp(msg Message, sessionID SessionID) (reject MessageRejectError) {
	return nil
}

type SessionSendTestSuite struct {
	suite.Suite

	*TestClient
	session
	sendChannel chan []byte
}

func (suite *SessionSendTestSuite) SetupTest() {
	suite.sendChannel = make(chan []byte, 1)
	suite.TestClient = new(TestClient)
	suite.session = session{
		store:       new(memoryStore),
		application: suite,
		log:         nullLog{},
		messageOut:  suite.sendChannel,
	}
}

func (suite *SessionSendTestSuite) SentMessage() (msg []byte) {
	select {
	case msg = <-suite.sendChannel:
	default:
	}
	return
}

func TestSessionSendTestSuite(t *testing.T) {
	suite.Run(t, new(SessionSendTestSuite))
}

func (suite *SessionSendTestSuite) SendNewOrderSingle() {
	msg := buildMessage()
	msg.Header.SetField(tagMsgType, FIXString("D"))
	require.Nil(suite.T(), suite.send(msg))
}

func (suite *SessionSendTestSuite) SendHeartbeat() {
	msg := buildMessage()
	msg.Header.SetField(tagMsgType, FIXString("0"))
	require.Nil(suite.T(), suite.send(msg))
}

func (suite *SessionSendTestSuite) SendLogon() {
	msg := buildMessage()
	msg.Header.SetField(tagMsgType, FIXString("A"))
	require.Nil(suite.T(), suite.send(msg))
}

func (suite *SessionSendTestSuite) MessagePersisted() {
	suite.Equal(2, suite.store.NextSenderMsgSeqNum(), "The next sender sequence number should be incremented")
	persistedMessages, err := suite.store.GetMessages(1, 1)
	suite.Nil(err)
	suite.Len(persistedMessages, 1, "The message should be persisted")
}

func (suite *SessionSendTestSuite) TestSendAppMessageLoggedOn() {
	var tests = []struct {
		sessionState
	}{
		{inSession{}},
		{resendState{}},
		{pendingTimeout{inSession{}}},
		{pendingTimeout{resendState{}}},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.session.sessionState = test.sessionState

		suite.SendNewOrderSingle()

		suite.Equal(1, suite.appCalled, "ToApp should be called")
		suite.MessagePersisted()
		suite.NotNil(suite.SentMessage(), "The message should have been sent")
	}
}

func (suite *SessionSendTestSuite) TestSendAdminMessageLoggedOn() {
	var tests = []struct {
		sessionState
	}{
		{inSession{}},
		{resendState{}},
		{pendingTimeout{inSession{}}},
		{pendingTimeout{resendState{}}},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.session.sessionState = test.sessionState

		suite.SendHeartbeat()

		suite.Equal(1, suite.adminCalled, "ToAdmin should be called")
		suite.MessagePersisted()
		suite.NotNil(suite.SentMessage(), "The message should have been sent")
	}
}

func (suite *SessionSendTestSuite) TestSendAppMessageNotLoggedOn() {
	var tests = []struct {
		sessionState
	}{
		{logonState{}},
		{logoutState{}},
		{latentState{}},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.session.sessionState = test.sessionState

		suite.SendNewOrderSingle()

		suite.Equal(1, suite.appCalled, "ToApp should be called even if the message is not sent")
		suite.MessagePersisted()
		suite.Nil(suite.SentMessage(), "The message should not be sent")
	}
}

func (suite *SessionSendTestSuite) TestSendAdminMessageNotLoggedOnNotSent() {
	var tests = []struct {
		sessionState
	}{
		{logonState{}},
		{logoutState{}},
		{latentState{}},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.session.sessionState = test.sessionState

		suite.SendHeartbeat()

		suite.Equal(1, suite.adminCalled, "ToAdmin should be called even if the message is not sent")
		suite.MessagePersisted()
		suite.Nil(suite.SentMessage(), "The message should not be sent")
	}
}

func (suite *SessionSendTestSuite) TestSendAdminMessageNotLoggedOnIsSent() {
	var tests = []struct {
		sessionState
	}{
		{logonState{}},
		{logoutState{}},
		{latentState{}},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.session.sessionState = test.sessionState

		suite.SendLogon()

		suite.Equal(1, suite.adminCalled, "ToAdmin should be called")
		suite.MessagePersisted()
		suite.NotNil(suite.SentMessage(), "The message should be sent")
	}
}
