package quickfix

import (
	"testing"

	"github.com/quickfixgo/quickfix/internal"
	"github.com/stretchr/testify/suite"
)

type resendStateTestSuite struct {
	SessionSuiteRig
}

func TestResendStateTestSuite(t *testing.T) {
	suite.Run(t, new(resendStateTestSuite))
}

func (s *resendStateTestSuite) SetupTest() {
	s.Init()
	s.session.State = resendState{}
}

func (s *resendStateTestSuite) TestIsLoggedOn() {
	s.True(s.session.IsLoggedOn())
}

func (s *resendStateTestSuite) TestIsConnected() {
	s.True(s.session.IsConnected())
}

func (s *resendStateTestSuite) TestIsSessionTime() {
	s.True(s.session.IsSessionTime())
}

func (s *resendStateTestSuite) TestTimeoutPeerTimeout() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.PeerTimeout)

	s.MockApp.AssertExpectations(s.T())
	s.State(pendingTimeout{resendState{}})
}

func (s *resendStateTestSuite) TestTimeoutUnchangedIgnoreLogonLogoutTimeout() {
	tests := []internal.Event{internal.LogonTimeout, internal.LogoutTimeout}

	for _, event := range tests {
		s.session.Timeout(s.session, event)
		s.State(resendState{})
	}
}

func (s *resendStateTestSuite) TestTimeoutUnchangedNeedHeartbeat() {
	s.MockApp.On("ToAdmin")
	s.session.Timeout(s.session, internal.NeedHeartbeat)

	s.MockApp.AssertExpectations(s.T())
	s.State(resendState{})
}

func (s *resendStateTestSuite) TestFixMsgIn() {
	s.session.State = inSession{}

	//in session expects seq number 1, send too high
	s.MessageFactory.SetNextSeqNum(2)
	s.MockApp.On("ToAdmin")

	msgSeqNum2 := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNum2)

	s.MockApp.AssertExpectations(s.T())
	s.State(resendState{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeResendRequest), s.MockApp.lastToAdmin)
	s.FieldEquals(tagBeginSeqNo, 1, s.MockApp.lastToAdmin.Body)
	s.NextTargetMsgSeqNum(1)

	msgSeqNum3 := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNum3)
	s.State(resendState{})
	s.NextTargetMsgSeqNum(1)

	msgSeqNum4 := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNum4)

	s.State(resendState{})
	s.NextTargetMsgSeqNum(1)

	s.MessageFactory.SetNextSeqNum(1)
	s.MockApp.On("FromApp").Return(nil)
	s.fixMsgIn(s.session, s.NewOrderSingle())

	s.MockApp.AssertNumberOfCalls(s.T(), "FromApp", 4)
	s.State(inSession{})
	s.NextTargetMsgSeqNum(5)
}

func (s *resendStateTestSuite) TestFixMsgInSequenceReset() {
	s.session.State = inSession{}

	//in session expects seq number 1, send too high
	s.MessageFactory.SetNextSeqNum(3)
	s.MockApp.On("ToAdmin")

	msgSeqNum3 := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNum3)

	s.MockApp.AssertExpectations(s.T())
	s.State(resendState{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeResendRequest), s.MockApp.lastToAdmin)
	s.FieldEquals(tagBeginSeqNo, 1, s.MockApp.lastToAdmin.Body)
	s.NextTargetMsgSeqNum(1)

	s.MessageFactory.SetNextSeqNum(1)
	s.MockApp.On("FromAdmin").Return(nil)
	s.fixMsgIn(s.session, s.SequenceReset(2))
	s.NextTargetMsgSeqNum(2)
	s.State(resendState{})

	s.MockApp.On("FromApp").Return(nil)
	s.fixMsgIn(s.session, s.NewOrderSingle())

	s.MockApp.AssertNumberOfCalls(s.T(), "FromApp", 2)
	s.NextTargetMsgSeqNum(4)
	s.State(inSession{})
}

func (s *resendStateTestSuite) TestFixMsgInResendChunk() {
	s.session.State = inSession{}
	s.ResendRequestChunkSize = 2

	//in session expects seq number 1, send too high
	s.MessageFactory.SetNextSeqNum(4)
	s.MockApp.On("ToAdmin")

	msgSeqNum4 := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNum4)

	s.MockApp.AssertExpectations(s.T())
	s.State(resendState{})
	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeResendRequest), s.MockApp.lastToAdmin)
	s.FieldEquals(tagBeginSeqNo, 1, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagEndSeqNo, 2, s.MockApp.lastToAdmin.Body)
	s.NextTargetMsgSeqNum(1)

	msgSeqNum5 := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNum5)
	s.State(resendState{})
	s.NextTargetMsgSeqNum(1)

	msgSeqNum6 := s.NewOrderSingle()
	s.fixMsgIn(s.session, msgSeqNum6)

	s.State(resendState{})
	s.NextTargetMsgSeqNum(1)

	s.MessageFactory.SetNextSeqNum(1)
	s.MockApp.On("FromApp").Return(nil)
	s.fixMsgIn(s.session, s.NewOrderSingle())

	s.MockApp.AssertNumberOfCalls(s.T(), "FromApp", 1)
	s.State(resendState{})
	s.NextTargetMsgSeqNum(2)

	s.fixMsgIn(s.session, s.NewOrderSingle())
	s.MockApp.AssertNumberOfCalls(s.T(), "FromApp", 2)
	s.State(resendState{})
	s.NextTargetMsgSeqNum(3)

	s.LastToAdminMessageSent()
	s.MessageType(string(msgTypeResendRequest), s.MockApp.lastToAdmin)
	s.FieldEquals(tagBeginSeqNo, 3, s.MockApp.lastToAdmin.Body)
	s.FieldEquals(tagEndSeqNo, 0, s.MockApp.lastToAdmin.Body)
}
