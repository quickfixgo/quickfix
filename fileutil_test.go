// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

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
