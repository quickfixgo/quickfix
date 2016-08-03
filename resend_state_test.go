package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type resendStateTestSuite struct {
	suite.Suite

	*mockApp
	state       resendState
	sendChannel chan []byte
	session     *session
}

func TestResendStateTestSuite(t *testing.T) {
	suite.Run(t, new(resendStateTestSuite))
}

func (s *resendStateTestSuite) SetupTest() {
	s.mockApp = new(mockApp)
	s.sendChannel = make(chan []byte, 10)
	s.session = &session{
		store:        new(memoryStore),
		application:  s.mockApp,
		messageOut:   s.sendChannel,
		log:          nullLog{},
		sessionState: s.state,
	}
}

func (s *resendStateTestSuite) TestIsLoggedOn() {
	s.True(s.state.IsLoggedOn())
}

func (s *resendStateTestSuite) TestTimeoutPeerTimeout() {
	s.mockApp.On("ToAdmin")
	nextState := s.state.Timeout(s.session, peerTimeout)

	s.mockApp.AssertExpectations(s.T())
	s.Equal(pendingTimeout{s.state}, nextState)
}

func (s *resendStateTestSuite) TestTimeoutUnchangedIgnoreLogonLogoutTimeout() {
	tests := []event{logonTimeout, logoutTimeout}

	for _, event := range tests {
		nextState := s.state.Timeout(s.session, event)
		s.Equal(s.state, nextState)
	}
}

func (s *resendStateTestSuite) TestTimeoutUnchangedNeedHeartbeat() {
	s.mockApp.On("ToAdmin")
	nextState := s.state.Timeout(s.session, needHeartbeat)

	s.mockApp.AssertExpectations(s.T())
	s.Equal(s.state, nextState)
}
