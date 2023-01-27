package quickfix

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func requireNotFileExists(t *testing.T, fname string) {
	_, err := os.Stat(fname)
	require.NotNil(t, err)
	require.True(t, os.IsNotExist(err))
}

func requireFileExists(t *testing.T, fname string) {
	_, err := os.Stat(fname)
	require.Nil(t, err)
}

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

func TestOpenOrCreateFile(t *testing.T) {
	// When the file doesn't exist yet
	fname := path.Join(os.TempDir(), fmt.Sprintf("TestOpenOrCreateFile-%d", os.Getpid()))
	requireNotFileExists(t, fname)
	defer os.Remove(fname)

	// Then it should be created
	f, err := openOrCreateFile(fname, 0664)
	require.Nil(t, err)
	requireFileExists(t, fname)

	// When the file already exists
	f.Close()

	// Then it should be opened
	f, err = openOrCreateFile(fname, 0664)
	require.Nil(t, err)
	require.Nil(t, f.Close())
}
