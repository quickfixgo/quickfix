package quickfix

import (
	"testing"

	"github.com/quickfixgo/quickfix/enum"
	"github.com/stretchr/testify/suite"
)

type InSessionTestSuite struct {
	SessionSuite
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

func (s *InSessionTestSuite) TestLogout() {
	s.mockApp.On("FromAdmin").Return(nil)
	s.mockApp.On("ToAdmin")
	s.session.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.State(latentState{})

	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_LOGOUT, s.mockApp.lastToAdmin)
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestLogoutResetOnLogout() {
	s.session.resetOnLogout = true

	s.mockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.mockApp.AssertExpectations(s.T())

	s.mockApp.On("FromAdmin").Return(nil)
	s.mockApp.On("ToAdmin")
	s.session.FixMsgIn(s.session, s.Logout())

	s.mockApp.AssertExpectations(s.T())
	s.State(latentState{})
	s.LastToAppMessageSent()
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_LOGOUT, s.mockApp.lastToAdmin)

	s.NextTargetMsgSeqNum(1)
	s.NextSenderMsgSeqNum(1)
	s.NoMessageQueued()
}

func (s *InSessionTestSuite) TestTimeoutNeedHeartbeat() {
	s.mockApp.On("ToAdmin").Return(nil)
	s.session.Timeout(s.session, needHeartbeat)

	s.mockApp.AssertExpectations(s.T())
	s.State(inSession{})
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_HEARTBEAT, s.mockApp.lastToAdmin)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestTimeoutPeerTimeout() {
	s.mockApp.On("ToAdmin").Return(nil)
	s.session.Timeout(s.session, peerTimeout)

	s.mockApp.AssertExpectations(s.T())
	s.State(pendingTimeout{inSession{}})
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_TEST_REQUEST, s.mockApp.lastToAdmin)
	s.NextSenderMsgSeqNum(2)
}
