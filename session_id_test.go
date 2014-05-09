package quickfix

import "testing"

func TestSessionID_String(t *testing.T) {
	var testCases = []struct {
		sessionID      SessionID
		expectedString string
	}{
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", TargetCompID: "TAR"}, "FIX.4.2:SND->TAR"},
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", TargetCompID: "TAR", Qualifier: "BLAH"}, "FIX.4.2:SND->TAR:BLAH"},
	}

	for _, tc := range testCases {
		actual := tc.sessionID.String()

		if tc.expectedString != actual {
			t.Errorf("Expected %v got %v ", tc.expectedString, actual)
		}
	}
}
