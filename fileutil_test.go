package quickfix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSessionIDFilename_MinimallyQualifiedSessionID(t *testing.T) {
	// When the session ID is
	sessionID := SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}

	// Then the filename should be
	require.Equal(t, "FIX.4.4-SENDER-TARGET", sessionIDFilenamePrefix(sessionID))
}

func TestSessionIDFilename_FullyQualifiedSessionID(t *testing.T) {
	// When the session ID is
	sessionID := SessionID{
		BeginString:      "FIX.4.4",
		SenderCompID:     "A",
		SenderSubID:      "B",
		SenderLocationID: "C",
		TargetCompID:     "D",
		TargetSubID:      "E",
		TargetLocationID: "F",
		Qualifier:        "G",
	}

	// Then the filename should be
	require.Equal(t, "FIX.4.4-A_B_C-D_E_F-G", sessionIDFilenamePrefix(sessionID))
}
