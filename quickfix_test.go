package quickfix

import (
	"time"

	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/internal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type QuickFIXSuite struct {
	suite.Suite
}

type KnowsFieldMap interface {
	Has(Tag) bool
	GetString(Tag) (string, MessageRejectError)
	GetInt(Tag) (int, MessageRejectError)
}

func (s *QuickFIXSuite) MessageType(msgType string, msg Message) {
	s.FieldEquals(tagMsgType, msgType, msg.Header)
}

func (s *QuickFIXSuite) FieldEquals(tag Tag, expectedValue interface{}, fieldMap KnowsFieldMap) {
	s.Require().True(fieldMap.Has(tag), "Tag %v not set", tag)

	switch expected := expectedValue.(type) {
	case string:
		val, err := fieldMap.GetString(tag)
		s.Nil(err)
		s.Equal(expected, val)
	case int:
		val, err := fieldMap.GetInt(tag)
		s.Nil(err)
		s.Equal(expected, val)
	default:
		s.FailNow("Field type not handled")
	}
}

func (s *QuickFIXSuite) MessageEqualsBytes(msgBytes []byte, msg Message) {
	_, err := msg.Build()
	s.Require().Nil(err)
	s.Equal(string(msg.rawMessage), string(msgBytes))
}

type MockApp struct {
	mock.Mock

	lastToAdmin Message
	lastToApp   Message
}

func (e *MockApp) OnCreate(sessionID SessionID) {
}

func (e *MockApp) OnLogon(sessionID SessionID) {
	e.Called()
}

func (e *MockApp) OnLogout(sessionID SessionID) {
	e.Called()
}

func (e *MockApp) FromAdmin(msg Message, sessionID SessionID) (reject MessageRejectError) {
	if err, ok := e.Called().Get(0).(MessageRejectError); ok {
		return err
	}

	return nil
}

func (e *MockApp) ToAdmin(msg Message, sessionID SessionID) {
	e.Called()
	e.lastToAdmin = msg
}

func (e *MockApp) ToApp(msg Message, sessionID SessionID) (err error) {
	e.lastToApp = msg
	return e.Called().Error(0)
}

func (e *MockApp) FromApp(msg Message, sessionID SessionID) (reject MessageRejectError) {
	if err, ok := e.Called().Get(0).(MessageRejectError); ok {
		return err
	}

	return nil
}

type MessageFactory struct {
	seqNum int
}

func (m *MessageFactory) buildMessage(msgType string) Message {
	m.seqNum++
	msg := NewMessage()
	msg.Header.
		SetField(tagBeginString, FIXString(enum.BeginStringFIX42)).
		SetField(tagSenderCompID, FIXString("TW")).
		SetField(tagTargetCompID, FIXString("ISLD")).
		SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()}).
		SetField(tagMsgSeqNum, FIXInt(m.seqNum)).
		SetField(tagMsgType, FIXString(msgType))
	return msg
}

func (m *MessageFactory) Logout() Message {
	return m.buildMessage("5")
}

func (m *MessageFactory) NewOrderSingle() Message {
	return m.buildMessage("D")
}

func (m *MessageFactory) Heartbeat() Message {
	return m.buildMessage("0")
}

func (m *MessageFactory) Logon() Message {
	return m.buildMessage("A")
}

type MockSessionReceiver struct {
	sendChannel chan []byte
}

func newMockSessionReceiver() MockSessionReceiver {
	return MockSessionReceiver{
		sendChannel: make(chan []byte, 10),
	}
}

func (p *MockSessionReceiver) LastMessage() (msg []byte, ok bool) {
	select {
	case msg, ok = <-p.sendChannel:
	default:
		ok = true
	}

	return
}

type SessionSuiteRig struct {
	QuickFIXSuite
	MessageFactory
	MockApp MockApp
	*session
	Receiver MockSessionReceiver
}

func (s *SessionSuiteRig) Init() {
	s.MockApp = MockApp{}
	s.MessageFactory = MessageFactory{}
	s.Receiver = newMockSessionReceiver()
	s.session = &session{
		sessionID:    SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"},
		store:        new(memoryStore),
		application:  &s.MockApp,
		log:          nullLog{},
		messageOut:   s.Receiver.sendChannel,
		sessionEvent: make(chan internal.Event),
	}
}

func (s *SessionSuiteRig) State(state sessionState) {
	s.Equal(state, s.session.State, "session state should be %v", state)
}

func (s *SessionSuiteRig) MessageSentEquals(msg Message) {
	msgBytes, ok := s.Receiver.LastMessage()
	s.True(ok, "Should be connected")
	s.NotNil(msgBytes, "Message should have been sent")
	s.MessageEqualsBytes(msgBytes, msg)
}

func (s *SessionSuiteRig) LastToAppMessageSent() {
	s.MessageSentEquals(s.MockApp.lastToApp)
}

func (s *SessionSuiteRig) LastToAdminMessageSent() {
	s.MessageSentEquals(s.MockApp.lastToAdmin)
}

func (s *SessionSuiteRig) NotStopped() {
	s.False(s.session.Stopped(), "session should not be stopped")
}

func (s *SessionSuiteRig) Stopped() {
	s.True(s.session.Stopped(), "session should be stopped")
}

func (s *SessionSuiteRig) Disconnected() {
	msg, ok := s.Receiver.LastMessage()
	s.Nil(msg, "Expect disconnect, not message")
	s.False(ok, "Expect disconnect")
}

func (s *SessionSuiteRig) NoMessageSent() {
	msg, _ := s.Receiver.LastMessage()
	s.Nil(msg, "no message should be sent")
}

func (s *SessionSuiteRig) NoMessageQueued() {
	s.Empty(s.session.toSend, "no messages should be queueud")
}

func (s *SessionSuiteRig) ExpectStoreReset() {
	s.NextSenderMsgSeqNum(1)
	s.NextTargetMsgSeqNum(1)
}

func (s *SessionSuiteRig) NextTargetMsgSeqNum(expected int) {
	s.Equal(expected, s.session.store.NextTargetMsgSeqNum(), "NextTargetMsgSeqNum should be %v ", expected)
}

func (s *SessionSuiteRig) NextSenderMsgSeqNum(expected int) {
	s.Equal(expected, s.session.store.NextSenderMsgSeqNum(), "NextSenderMsgSeqNum should be %v", expected)
}

func (s *SessionSuiteRig) NoMessagePersisted(seqNum int) {
	persistedMessages, err := s.session.store.GetMessages(seqNum, seqNum)
	s.Nil(err)
	s.Empty(persistedMessages, "The message should not be persisted")
}

func (s *SessionSuiteRig) MessagePersisted(msg Message) {
	var err error
	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	s.Nil(err, "message should have seq num")

	persistedMessages, err := s.session.store.GetMessages(seqNum, seqNum)
	s.Nil(err)
	s.Len(persistedMessages, 1, "a message should be stored at %v", seqNum)
	s.MessageEqualsBytes(persistedMessages[0], msg)
}
