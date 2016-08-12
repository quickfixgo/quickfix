package quickfix

import (
	"testing"

	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/internal"
	"github.com/stretchr/testify/suite"
)

type InSessionTestSuite struct {
	SessionSuiteRig
}

func TestInSessionTestSuite(t *testing.T) {
	suite.Run(t, new(InSessionTestSuite))
}

func (s *InSessionTestSuite) SetupTest() {
	s.Init()
	s.session.State = inSession{}
}

func (s *InSessionTestSuite) TestIsLoggedOn() {
	s.True(s.session.IsLoggedOn())
}

func (s *InSessionTestSuite) TestIsConnected() {
	s.True(s.session.IsConnected())
}

func (s *InSessionTestSuite) TestIsSessionTime() {
	s.True(s.session.IsSessionTime())
}

func (s *InSessionTestSuite) TestLogout() {
	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.FixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})

	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_LOGOUT, s.MockApp.lastToAdmin)
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestLogoutResetOnLogout() {
	s.session.resetOnLogout = true

	s.MockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.MockApp.AssertExpectations(s.T())

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.FixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
	s.LastToAppMessageSent()
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_LOGOUT, s.MockApp.lastToAdmin)

	s.NextTargetMsgSeqNum(1)
	s.NextSenderMsgSeqNum(1)
	s.NoMessageQueued()
}

func (s *InSessionTestSuite) TestTimeoutNeedHeartbeat() {
	s.MockApp.On("ToAdmin").Return(nil)
	s.session.Timeout(s.session, internal.NeedHeartbeat)

	s.MockApp.AssertExpectations(s.T())
	s.State(inSession{})
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_HEARTBEAT, s.MockApp.lastToAdmin)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestTimeoutPeerTimeout() {
	s.MockApp.On("ToAdmin").Return(nil)
	s.session.Timeout(s.session, internal.PeerTimeout)

	s.MockApp.AssertExpectations(s.T())
	s.State(pendingTimeout{inSession{}})
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_TEST_REQUEST, s.MockApp.lastToAdmin)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestDisconnected() {
	s.MockApp.On("OnLogout").Return(nil)
	s.session.Disconnected(s.session)
	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})
}

func (s *InSessionTestSuite) TestStop() {
	s.MockApp.On("ToAdmin")
	s.session.Stop(s.session)

	s.MockApp.AssertExpectations(s.T())
	s.State(logoutState{})
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_LOGOUT, s.MockApp.lastToAdmin)

	s.MockApp.On("OnLogout")
	s.session.Timeout(s.session, <-s.sessionEvent)
	s.MockApp.AssertExpectations(s.T())
	s.Stopped()
	s.Disconnected()
}
