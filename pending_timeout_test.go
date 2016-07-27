package quickfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPendingTimeout_SessionTimeout(t *testing.T) {
	session := &session{
		log: nullLog{},
	}

	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	for _, state := range tests {
		nextState := state.Timeout(session, peerTimeout)
		assert.IsType(t, latentState{}, nextState)
	}
}

func TestPendingTimeout_TimeoutUnchangedState(t *testing.T) {
	session := &session{
		log: nullLog{},
	}

	tests := []pendingTimeout{
		pendingTimeout{inSession{}},
		pendingTimeout{resendState{}},
	}

	testEvents := []event{needHeartbeat, logonTimeout, logoutTimeout}

	for _, state := range tests {
		for _, event := range testEvents {
			nextState := state.Timeout(session, event)
			assert.Equal(t, state, nextState)
		}
	}
}
