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
	s.session.messageStash = make(map[int]Message)
}

func (s *InSessionTestSuite) TestPreliminary() {
	s.True(s.session.IsLoggedOn())
	s.True(s.session.IsConnected())
	s.True(s.session.IsSessionTime())
}

func (s *InSessionTestSuite) TestLogout() {
	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.fixMsgIn(s.session, s.Logout())

	s.MockApp.AssertExpectations(s.T())
	s.State(latentState{})

	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_LOGOUT, s.MockApp.lastToAdmin)
	s.NextTargetMsgSeqNum(2)
	s.NextSenderMsgSeqNum(2)
}

func (s *InSessionTestSuite) TestLogoutResetOnLogout() {
	s.session.ResetOnLogout = true

	s.MockApp.On("ToApp").Return(nil)
	s.Nil(s.queueForSend(s.NewOrderSingle()))
	s.MockApp.AssertExpectations(s.T())

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("OnLogout")
	s.session.fixMsgIn(s.session, s.Logout())

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

func (s *InSessionTestSuite) TestFIXMsgInTargetTooHigh() {
	s.MessageFactory.seqNum = 5

	s.MockApp.On("ToAdmin")
	msgSeqNumTooHigh := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNumTooHigh)

	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_RESEND_REQUEST, s.MockApp.lastToAdmin)
	s.FieldEquals(tagBeginSeqNo, 1, s.MockApp.lastToAdmin.Body)

	s.State(resendState{})
	s.NextTargetMsgSeqNum(1)

	stashedMsg, ok := s.session.messageStash[6]
	s.True(ok)

	rawMsg, _ := msgSeqNumTooHigh.Build()
	stashedRawMsg, _ := stashedMsg.Build()
	s.Equal(string(rawMsg), string(stashedRawMsg))
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestAllAdminExpectGapFill() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 3)
	s.NextSenderMsgSeqNum(4)

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.fixMsgIn(s.session, s.ResendRequest(1))

	s.MockApp.AssertExpectations(s.T())
	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_SEQUENCE_RESET, s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagNewSeqNo, 4, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagGapFillFlag, true, s.MockApp.lastToAdmin.Body)

	s.NextSenderMsgSeqNum(4)
	s.State(inSession{})
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestAllAdminThenApp() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.On("ToApp").Return(nil)
	s.session.send(s.NewOrderSingle())
	s.LastToAppMessageSent()

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 2)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)
	s.NextSenderMsgSeqNum(4)

	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToAdmin")
	s.MockApp.On("ToApp").Return(nil)
	s.fixMsgIn(s.session, s.ResendRequest(1))

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 3)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 2)

	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_SEQUENCE_RESET, s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagNewSeqNo, 3, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagGapFillFlag, true, s.MockApp.lastToAdmin.Body)

	s.LastToAppMessageSent()
	s.MessageType(enum.MsgType_ORDER_SINGLE, s.MockApp.lastToApp)
	s.FieldEquals(tagMsgSeqNum, 3, s.MockApp.lastToApp.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToApp.Header)

	s.NextSenderMsgSeqNum(4)
	s.State(inSession{})
}

func (s *InSessionTestSuite) TestFIXMsgInResendRequestDoNotSendApp() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.On("ToApp").Return(nil)
	s.session.send(s.NewOrderSingle())
	s.LastToAppMessageSent()

	s.session.Timeout(s.session, internal.NeedHeartbeat)
	s.LastToAdminMessageSent()

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 2)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)
	s.NextSenderMsgSeqNum(4)

	//NOTE: a cheat here, need to reset mock
	s.MockApp = MockApp{}
	s.MockApp.On("FromAdmin").Return(nil)
	s.MockApp.On("ToApp").Return(ErrDoNotSend)
	s.MockApp.On("ToAdmin")
	s.fixMsgIn(s.session, s.ResendRequest(1))

	s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 1)
	s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 1)

	s.LastToAdminMessageSent()
	s.MessageType(enum.MsgType_SEQUENCE_RESET, s.MockApp.lastToAdmin)
	s.FieldEquals(tagMsgSeqNum, 1, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagPossDupFlag, true, s.MockApp.lastToAdmin.Header)
	s.FieldEquals(tagNewSeqNo, 4, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagGapFillFlag, true, s.MockApp.lastToAdmin.Body)

	s.NoMessageSent()

	s.NextSenderMsgSeqNum(4)
	s.State(inSession{})
}
