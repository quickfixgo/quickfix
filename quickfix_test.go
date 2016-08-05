package quickfix

import "github.com/stretchr/testify/suite"

type QuickFIXSuite struct {
	suite.Suite
}

type KnowsFieldMap interface {
	Has(Tag) bool
	GetString(Tag) (string, MessageRejectError)
	GetInt(Tag) (int, MessageRejectError)
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
		sessionID:   SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"},
		store:       new(memoryStore),
		application: s.mockApp,
		log:         nullLog{},
		messageOut:  s.Receiver.sendChannel,
	}
}

func (s *SessionSuite) LastToAdminMessageSent() {
	msgBytes := s.Receiver.LastMessage()
	s.NotNil(msgBytes)

	_, err := s.mockApp.lastToAdmin.Build()
	s.Nil(err)

	s.Equal(string(s.mockApp.lastToAdmin.rawMessage), string(msgBytes))
}

func (s *SessionSuite) NoMessageSent() {
	s.Nil(s.Receiver.LastMessage(), "no message should be sent")
}

func (s *SessionSuite) NoMessageQueued() {
	s.Empty(s.session.toSend, "no messages should be queueud")
}

func (s *SessionSuite) NextTargetMsgSeqNum(expected int) {
	s.Equal(expected, s.session.store.NextTargetMsgSeqNum(), "NextTargetMsgSeqNum should be ", expected)
}

func (s *SessionSuite) NextSenderMsgSeqNum(expected int) {
	s.Equal(expected, s.session.store.NextSenderMsgSeqNum(), "NextSenderMsgSeqNum should be ", expected)
}
