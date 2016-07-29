package quickfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResendState_IsLoggedOn(t *testing.T) {
	assert.True(t, resendState{}.IsLoggedOn())
}

func TestResendState_TimeoutPeerTimeout(t *testing.T) {
	otherEnd := make(chan []byte)
	go func() {
		<-otherEnd
	}()

	state := resendState{}
	session := &session{
		store:        new(memoryStore),
		application:  new(TestClient),
		messageOut:   otherEnd,
		log:          nullLog{},
		sessionState: state,
	}

	nextState := state.Timeout(session, peerTimeout)
	assert.Equal(t, pendingTimeout{state}, nextState)
}

func TestResendState_TimeoutUnchanged(t *testing.T) {
	otherEnd := make(chan []byte)
	go func() {
		<-otherEnd
	}()

	state := resendState{}
	session := &session{
		store:        new(memoryStore),
		application:  new(TestClient),
		messageOut:   otherEnd,
		log:          nullLog{},
		sessionState: state,
	}

	tests := []event{needHeartbeat, logonTimeout, logoutTimeout}

	for _, event := range tests {
		nextState := state.Timeout(session, event)
		assert.Equal(t, state, nextState)
	}
}
