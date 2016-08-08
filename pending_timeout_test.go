package quickfix

import (
	"testing"

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

func (s *PendingTimeoutTestSuite) TestSessionTimeout() {
	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	for _, state := range tests {
		s.session.State = state

		s.session.Timeout(s.session, peerTimeout)
		s.State(latentState{})
	}
}

func (s *PendingTimeoutTestSuite) TestTimeoutUnchangedState() {
	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	testEvents := []event{needHeartbeat, logonTimeout, logoutTimeout}

	for _, state := range tests {
		s.session.State = state

		for _, event := range testEvents {
			s.session.Timeout(s.session, event)
			s.State(state)
		}
	}
}
