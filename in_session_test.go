package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InSessionTestSuite struct {
	suite.Suite

	*messageFactory
	*mockApp
	*session
	receiver *mockSessionReceiver
}

func TestInSessionTestSuite(t *testing.T) {
	suite.Run(t, new(InSessionTestSuite))
}

func (s *InSessionTestSuite) TestIsLoggedOn() {
	s.True(s.session.IsLoggedOn())
}

func (s *InSessionTestSuite) SetupTest() {
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

func (s *InSessionTestSuite) TestLogout() {
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

func (s *InSessionTestSuite) TestLogoutResetOnLogout() {
	s.session.resetOnLogout = true

	s.mockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.mockApp.AssertExpectations(s.T())

	s.mockApp.On("FromAdmin").Return(nil)
	s.mockApp.On("ToAdmin")
	nextState := s.session.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.Equal(latentState{}, nextState)
	s.NotNil(s.receiver.LastMessage())
	msgType, err := s.mockApp.lastToAdmin.Header.GetString(tagMsgType)
	s.Nil(err)
	s.Equal("5", msgType)

	s.Equal(1, s.session.store.NextSenderMsgSeqNum())
	s.Equal(1, s.session.store.NextTargetMsgSeqNum())
	s.Empty(s.session.toSend)
}
