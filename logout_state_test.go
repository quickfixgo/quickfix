package quickfix

import (
	"testing"

	"github.com/quickfixgo/quickfix/internal"
	"github.com/stretchr/testify/suite"
)

type LogoutStateTestSuite struct {
	SessionSuiteRig
}

func TestLogoutStateTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutStateTestSuite))
}

func (s *LogoutStateTestSuite) SetupTest() {
	s.Init()
	s.session.State = logoutState{}
}

func (s *LogoutStateTestSuite) TestPreliminary() {
	s.False(s.session.IsLoggedOn())
	s.True(s.session.IsConnected())
	s.True(s.session.IsSessionTime())
}

func (s *LogoutStateTestSuite) TestTimeoutLogoutTimeout() {
	s.MockApp.On("OnLogout").Return(nil)
	s.Timeout(s.session, internal.LogoutTimeout)

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
}

func (s *LogoutStateTestSuite) TestTimeoutNotLogoutTimeout() {
	tests := []internal.Event{internal.PeerTimeout, internal.NeedHeartbeat, internal.LogonTimeout}

	for _, test := range tests {
		s.Timeout(s.session, test)
		s.State(logoutState{})
	}
}

func (s *LogoutStateTestSuite) TestDisconnected() {
	s.MockApp.On("OnLogout").Return(nil)
	s.session.Disconnected(s.session)

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
}

func (s *LogoutStateTestSuite) TestFixMsgInNotLogout() {
	s.MockApp.On("FromApp").Return(nil)
	s.fixMsgIn(s.session, s.NewOrderSingle())

	s.MockApp.AssertExpectations(s.T())
	s.State(logoutState{})
	s.NextTargetMsgSeqNum(2)
}

func (s *LogoutStateTestSuite) TestFixMsgInNotLogoutReject() {
	s.MockApp.On("FromApp").Return(ConditionallyRequiredFieldMissing(Tag(11)))
	s.MockApp.On("ToApp").Return(nil)
	s.fixMsgIn(s.session, s.NewOrderSingle())

	s.MockApp.AssertExpectations(s.T())
	s.State(logoutState{})
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(2)

	s.NoMessageSent()
}

func (s *LogoutStateTestSuite) TestFixMsgInLogout() {
	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("OnLogout").Return(nil)
	s.fixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(1)
	s.NoMessageSent()
}

func (s *LogoutStateTestSuite) TestFixMsgInLogoutResetOnLogout() {
	s.session.ResetOnLogout = true

	s.MockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.MockApp.AssertExpectations(s.T())

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("OnLogout").Return(nil)
	s.fixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
	s.NextTargetMsgSeqNum(1)
	s.NextSenderMsgSeqNum(1)

	s.NoMessageSent()
	s.NoMessageQueued()
}

func (s *LogoutStateTestSuite) TestStop() {
	s.session.Stop(s.session)
	s.State(logoutState{})
	s.NotStopped()
}
