package quickfix

import "testing"

func TestSessionState_IsLoggedOn(t *testing.T) {

	var tests = []struct {
		sessionState
		expectIsLoggedOn bool
	}{
		{latentState{}, false},
		{logonState{}, false},
		{logoutState{}, false},
		{inSession{}, true},
		{resendState{}, true},
	}

	for _, test := range tests {
		switch {
		case test.expectIsLoggedOn && !test.IsLoggedOn():
			t.Errorf("'%v' should be LoggedOn", test.sessionState)
		case !test.expectIsLoggedOn && test.IsLoggedOn():
			t.Errorf("'%v' should not be LoggedOn", test.sessionState)
		}
	}
}
