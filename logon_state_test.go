package quickfix

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type LogonStateTestSuite struct {
	SessionSuite
}

func TestLogonStateTestSuite(t *testing.T) {
	suite.Run(t, new(LogonStateTestSuite))
}

func (s *LogonStateTestSuite) SetupTest() {
	s.Init()
	s.session.sessionState = logonState{}
}

func (s *LogonStateTestSuite) TestIsLoggedOn() {
	s.False(s.session.IsLoggedOn())
}

func (s *LogonStateTestSuite) TestTimeoutLogonTimeout() {
	nextState := s.Timeout(s.session, logonTimeout)
	s.Equal(latentState{}, nextState)
}

func (s *LogonStateTestSuite) TestTimeoutNotLogonTimeout() {
	tests := []event{peerTimeout, needHeartbeat, logoutTimeout}

	for _, test := range tests {
		nextState := s.Timeout(s.session, test)
		s.Equal(logonState{}, nextState)
	}
}

func (s *LogonStateTestSuite) TestFixMsgInNotLogon() {
	nextState := s.FixMsgIn(s.session, s.NewOrderSingle())

	s.mockApp.AssertExpectations(s.T())
	s.Equal(latentState{}, nextState)
	s.NextTargetMsgSeqNum(1)
}

func (s *LogonStateTestSuite) TestFixMsgInLogon() {
	s.store.IncrNextSenderMsgSeqNum()
	s.messageFactory.seqNum = 1
	s.store.IncrNextTargetMsgSeqNum()

	logon := s.Logon()
	logon.Body.SetField(tagHeartBtInt, FIXInt(32))

	s.mockApp.On("FromAdmin").Return(nil)
	s.mockApp.On("OnLogon")
	s.mockApp.On("ToAdmin")
	nextState := s.FixMsgIn(s.session, logon)

	s.mockApp.AssertExpectations(s.T())

	s.Equal(inSession{}, nextState)
	s.Equal(32*time.Second, s.session.heartBeatTimeout)

	s.LastToAdminMessageSent()
	s.FieldEquals(tagMsgType, "A", s.mockApp.lastToAdmin.Header)
	s.FieldEquals(tagHeartBtInt, 32, s.mockApp.lastToAdmin.Body)

	s.NextTargetMsgSeqNum(3)
	s.NextSenderMsgSeqNum(3)
}

func (s *LogonStateTestSuite) TestFixMsgInLogonInitiateLogon() {
	s.session.initiateLogon = true
	s.store.IncrNextSenderMsgSeqNum()
	s.messageFactory.seqNum = 1
	s.store.IncrNextTargetMsgSeqNum()

	logon := s.Logon()
	logon.Body.SetField(tagHeartBtInt, FIXInt(32))

	s.mockApp.On("FromAdmin").Return(nil)
	s.mockApp.On("OnLogon")
	nextState := s.FixMsgIn(s.session, logon)

	s.mockApp.AssertExpectations(s.T())
	s.Equal(inSession{}, nextState)

	s.NextTargetMsgSeqNum(3)
	s.NextSenderMsgSeqNum(2)
}
