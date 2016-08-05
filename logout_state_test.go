package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LogoutStateTestSuite struct {
	SessionSuite
}

func TestLogoutStateTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutStateTestSuite))
}

func (s *LogoutStateTestSuite) SetupTest() {
	s.Init()
	s.session.sessionState = logoutState{}
}

func (s *LogoutStateTestSuite) TestIsLoggedOn() {
	s.False(s.session.IsLoggedOn())
}

func (s *LogoutStateTestSuite) TestTimeoutLogoutTimeout() {
	nextState := s.Timeout(s.session, logoutTimeout)
	s.IsType(latentState{}, nextState)
}

func (s *LogoutStateTestSuite) TestTimeoutNotLogoutTimeout() {
	tests := []event{peerTimeout, needHeartbeat, logonTimeout}

	for _, test := range tests {
		nextState := s.Timeout(s.session, test)
		s.IsType(logoutState{}, nextState)
	}
}

func (s *LogoutStateTestSuite) TestFixMsgInNotLogout() {
	s.mockApp.On("FromApp").Return(nil)
	nextState := s.FixMsgIn(s.session, s.NewOrderSingle())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(logoutState{}, nextState)
	s.NextTargetMsgSeqNum(2)
}

func (s *LogoutStateTestSuite) TestFixMsgInNotLogoutReject() {
	s.mockApp.On("FromApp").Return(ConditionallyRequiredFieldMissing(Tag(11)))
	s.mockApp.On("ToApp").Return(nil)
	nextState := s.FixMsgIn(s.session, s.NewOrderSingle())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(logoutState{}, nextState)
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(2)

	s.NoMessageSent()
}

func (s *LogoutStateTestSuite) TestFixMsgInLogout() {
	s.mockApp.On("FromAdmin").Return(nil)
	nextState := s.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(latentState{}, nextState)
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(1)
	s.NoMessageSent()
}

func (s *LogoutStateTestSuite) TestFixMsgInLogoutResetOnLogout() {
	s.session.resetOnLogout = true

	s.mockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.mockApp.AssertExpectations(s.T())

	s.mockApp.On("FromAdmin").Return(nil)
	nextState := s.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.IsType(latentState{}, nextState)
	s.NextTargetMsgSeqNum(1)
	s.NextSenderMsgSeqNum(1)

	s.NoMessageSent()
	s.NoMessageQueued()
}
