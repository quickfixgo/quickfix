package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type logoutStateTestSuite struct {
	suite.Suite

	*messageFactory
	*mockApp
	*session
	receiver *mockSessionReceiver
}

func TestLogoutStateTestSuite(t *testing.T) {
	suite.Run(t, new(logoutStateTestSuite))
}

func (s *logoutStateTestSuite) SetupTest() {
	s.mockApp = new(mockApp)
	s.messageFactory = new(messageFactory)
	s.receiver = newMockSessionReceiver()
	s.session = &session{
		sessionState: logoutState{},
		sessionID:    SessionID{BeginString: "FIX.4.2", TargetCompID: "TW", SenderCompID: "ISLD"},
		store:        new(memoryStore),
		application:  s.mockApp,
		log:          nullLog{},
		messageOut:   s.receiver.sendChannel,
	}
}

func (s *logoutStateTestSuite) TestIsLoggedOn() {
	s.False(s.session.IsLoggedOn())
}

func (s *logoutStateTestSuite) TestTimeoutLogoutTimeout() {
	nextState := s.Timeout(s.session, logoutTimeout)
	s.IsType(latentState{}, nextState)
}

func (s *logoutStateTestSuite) TestTimeoutNotLogoutTimeout() {
	tests := []event{peerTimeout, needHeartbeat, logonTimeout}

	for _, test := range tests {
		nextState := s.Timeout(s.session, test)
		s.IsType(logoutState{}, nextState)
	}
}

func (s *logoutStateTestSuite) TestFixMsgInNotLogout() {
	s.mockApp.On("FromApp").Return(nil)
	nextState := s.FixMsgIn(s.session, s.NewOrderSingle())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(logoutState{}, nextState)
	s.Equal(2, s.store.NextTargetMsgSeqNum())
}

func (s *logoutStateTestSuite) TestFixMsgInNotLogoutReject() {
	s.mockApp.On("FromApp").Return(ConditionallyRequiredFieldMissing(Tag(11)))
	s.mockApp.On("ToApp").Return(nil)
	nextState := s.FixMsgIn(s.session, s.NewOrderSingle())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(logoutState{}, nextState)
	s.Equal(2, s.store.NextTargetMsgSeqNum())
	s.Equal(2, s.store.NextSenderMsgSeqNum())

	s.Nil(s.receiver.LastMessage(), "nothing should be sent in logout state")
}

func (s *logoutStateTestSuite) TestFixMsgInLogout() {
	s.mockApp.On("FromAdmin").Return(nil)
	nextState := s.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(latentState{}, nextState)
	s.Equal(2, s.session.store.NextTargetMsgSeqNum())
	s.Equal(1, s.session.store.NextSenderMsgSeqNum())

	s.Nil(s.receiver.LastMessage(), "nothing should be sent in logout state")
}

func (s *logoutStateTestSuite) TestFixMsgInLogoutResetOnLogout() {
	s.session.resetOnLogout = true

	s.mockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.mockApp.AssertExpectations(s.T())

	s.mockApp.On("FromAdmin").Return(nil)
	nextState := s.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(latentState{}, nextState)
	s.Equal(1, s.session.store.NextTargetMsgSeqNum())
	s.Equal(1, s.session.store.NextSenderMsgSeqNum())

	s.Nil(s.receiver.LastMessage(), "nothing should be sent in logout state")
	s.Empty(s.session.toSend)
}
