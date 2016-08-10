package quickfix

import (
	"github.com/quickfixgo/quickfix/internal"
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

type SessionSuite struct {
	QuickFIXSuite
	*messageFactory
	*mockApp
	*session
	Receiver *mockSessionReceiver
}

func (s *SessionSuite) Init() {
	s.mockApp = new(mockApp)
	s.messageFactory = new(messageFactory)
	s.Receiver = newMockSessionReceiver()
	s.session = &session{
		sessionID:    SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"},
		store:        new(memoryStore),
		application:  s.mockApp,
		log:          nullLog{},
		messageOut:   s.Receiver.sendChannel,
		sessionEvent: make(chan internal.Event),
	}
}

func (s *SessionSuite) State(state sessionState) {
	s.Equal(state, s.session.State, "session state should be %v", state)
}

func (s *SessionSuite) MessageSentEquals(msg Message) {
	msgBytes, ok := s.Receiver.LastMessage()
	s.True(ok, "Should be connected")
	s.NotNil(msgBytes, "Message should have been sent")
	s.MessageEqualsBytes(msgBytes, msg)
}

func (s *SessionSuite) LastToAppMessageSent() {
	s.MessageSentEquals(s.mockApp.lastToApp)
}

func (s *SessionSuite) LastToAdminMessageSent() {
	s.MessageSentEquals(s.mockApp.lastToAdmin)
}

func (s *SessionSuite) Disconnected() {
	msg, ok := s.Receiver.LastMessage()
	s.Nil(msg, "Expect disconnect, not message")
	s.False(ok, "Expect disconnect")
}

func (s *SessionSuite) NoMessageSent() {
	msg, _ := s.Receiver.LastMessage()
	s.Nil(msg, "no message should be sent")
}

func (s *SessionSuite) NoMessageQueued() {
	s.Empty(s.session.toSend, "no messages should be queueud")
}

func (s *SessionSuite) ExpectStoreReset() {
	s.NextSenderMsgSeqNum(1)
	s.NextTargetMsgSeqNum(1)
}

func (s *SessionSuite) NextTargetMsgSeqNum(expected int) {
	s.Equal(expected, s.session.store.NextTargetMsgSeqNum(), "NextTargetMsgSeqNum should be %v ", expected)
}

func (s *SessionSuite) NextSenderMsgSeqNum(expected int) {
	s.Equal(expected, s.session.store.NextSenderMsgSeqNum(), "NextSenderMsgSeqNum should be %v", expected)
}

func (s *SessionSuite) NoMessagePersisted(seqNum int) {
	persistedMessages, err := s.session.store.GetMessages(seqNum, seqNum)
	s.Nil(err)
	s.Empty(persistedMessages, "The message should not be persisted")
}

func (s *SessionSuite) MessagePersisted(msg Message) {
	var err error
	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	s.Nil(err, "message should have seq num")

	persistedMessages, err := s.session.store.GetMessages(seqNum, seqNum)
	s.Nil(err)
	s.Len(persistedMessages, 1, "a message should be stored at %v", seqNum)
	s.MessageEqualsBytes(persistedMessages[0], msg)
}
