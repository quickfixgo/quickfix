package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SessionTestSuite struct {
	suite.Suite

	*messageFactory
	*mockApp
	*session
	receiver *mockSessionReceiver
}

func TestInSessionTestSuite(t *testing.T) {
	suite.Run(t, new(SessionTestSuite))
}

func (s *SessionTestSuite) TestIsLoggedOn() {
	s.True(s.session.IsLoggedOn())
}

func (s *SessionTestSuite) SetupTest() {
	s.mockApp = new(mockApp)
	s.messageFactory = new(messageFactory)
	s.receiver = newMockSessionReceiver()
	s.session = &session{
		sessionState: inSession{},
		sessionID:    SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"},
		store:        new(memoryStore),
		application:  s.mockApp,
		log:          nullLog{},
		messageOut:   s.receiver.sendChannel,
	}
}

func (s *SessionTestSuite) TestLogout() {
	s.mockApp.On("FromAdmin").Return(nil)
	s.mockApp.On("ToAdmin")
	nextState := s.session.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.Equal(latentState{}, nextState)
	s.NotNil(s.receiver.LastMessage())
	msgType, err := s.mockApp.lastToAdmin.Header.GetString(tagMsgType)
	s.Nil(err)
	s.Equal("5", msgType)

	s.Equal(2, s.session.store.NextSenderMsgSeqNum())
	s.Equal(2, s.session.store.NextTargetMsgSeqNum())
}
