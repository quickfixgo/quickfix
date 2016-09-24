package quickfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionID_String(t *testing.T) {
	var testCases = []struct {
		sessionID      SessionID
		expectedString string
	}{
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", TargetCompID: "TAR"}, "FIX.4.2:SND->TAR"},
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", TargetCompID: "TAR", Qualifier: "BLAH"}, "FIX.4.2:SND->TAR:BLAH"},
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", SenderSubID: "SSUB", SenderLocationID: "SLOC",
			TargetCompID: "TAR", TargetSubID: "TSUB", TargetLocationID: "TLOC",
			Qualifier: "BLAH"}, "FIX.4.2:SND/SSUB/SLOC->TAR/TSUB/TLOC:BLAH"},
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", SenderLocationID: "SLOC",
			TargetCompID: "TAR", TargetSubID: "TSUB", TargetLocationID: "TLOC",
		}, "FIX.4.2:SND/SLOC->TAR/TSUB/TLOC"},
	}

	for _, tc := range testCases {
		actual := tc.sessionID.String()
		assert.Equal(t, tc.expectedString, actual)
	}
}
