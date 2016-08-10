package quickfix

import (
	"testing"

	"github.com/quickfixgo/quickfix/internal"
	"github.com/stretchr/testify/suite"
)

type PendingTimeoutTestSuite struct {
	SessionSuite
}

func TestPendingTimeoutTestSuite(t *testing.T) {
	suite.Run(t, new(PendingTimeoutTestSuite))
}

func (s *PendingTimeoutTestSuite) SetupTest() {
	s.Init()
}

func (s *PendingTimeoutTestSuite) TestIsConnectedIsLoggedOn() {
	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	for _, state := range tests {
		s.session.State = state

		s.True(s.session.IsConnected())
		s.True(s.session.IsLoggedOn())
	}
}

func (s *PendingTimeoutTestSuite) TestSessionTimeout() {
	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	for _, state := range tests {
		s.session.State = state

		s.mockApp.On("OnLogout").Return(nil)
		s.session.Timeout(s.session, internal.PeerTimeout)

		s.mockApp.AssertExpectations(s.T())
		s.State(latentState{})
	}
}

func (s *PendingTimeoutTestSuite) TestTimeoutUnchangedState() {
	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	testEvents := []internal.Event{internal.NeedHeartbeat, internal.LogonTimeout, internal.LogoutTimeout}

	for _, state := range tests {
		s.session.State = state

		for _, event := range testEvents {
			s.session.Timeout(s.session, event)
			s.State(state)
		}
	}
}

func (s *PendingTimeoutTestSuite) TestDisconnected() {
	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	for _, state := range tests {
		s.SetupTest()
		s.session.State = state

		s.mockApp.On("OnLogout").Return(nil)
		s.session.Disconnected(s.session)

		s.mockApp.AssertExpectations(s.T())
		s.State(latentState{})
	}
}
