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

package file

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/internal/testsuite"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// FileStoreTestSuite runs all tests in the MessageStoreTestSuite against the FileStore implementation.
type FileStoreTestSuite struct {
	testsuite.StoreTestSuite
	fileStoreRootPath string
}

func (suite *FileStoreTestSuite) SetupTest() {
	suite.fileStoreRootPath = path.Join(os.TempDir(), fmt.Sprintf("FileStoreTestSuite-%d", os.Getpid()))
	fileStorePath := path.Join(suite.fileStoreRootPath, fmt.Sprintf("%d", time.Now().UnixNano()))
	sessionID := quickfix.SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}

	// create settings
	settings, err := quickfix.ParseSettings(strings.NewReader(fmt.Sprintf(`
[DEFAULT]
FileStorePath=%s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, fileStorePath, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	// create store
	suite.MsgStore, err = NewStoreFactory(settings).Create(sessionID)
	require.Nil(suite.T(), err)
}

func (suite *FileStoreTestSuite) TearDownTest() {
	suite.MsgStore.Close()
	os.RemoveAll(suite.fileStoreRootPath)
}

func TestFileStoreTestSuite(t *testing.T) {
	suite.Run(t, new(FileStoreTestSuite))
}

func TestStringParse(t *testing.T) {
	assert := assert2.New(t)
	i, err := strconv.Atoi(strings.Trim("00005\n", "\r\n"))
	assert.Nil(err)
	assert.Equal(5, i)

	i, err = strconv.Atoi(strings.Trim("00006\r", "\r\n"))
	assert.Nil(err)
	assert.Equal(6, i)
}

func TestIterateMessagesEOF(t *testing.T) {
	newStore := func(t *testing.T) *fileStore {
		t.Helper()
		store, err := newFileStore(quickfix.SessionID{
			BeginString:  "FIX.4.4",
			SenderCompID: "SENDER",
			TargetCompID: "TARGET",
		}, t.TempDir(), true)
		require.NoError(t, err)
		t.Cleanup(func() { require.NoError(t, store.Close()) })
		return store
	}

	t.Run("empty", func(t *testing.T) {
		store := newStore(t)
		visited := 0
		err := store.IterateMessages(7, 10, func([]byte) error {
			visited++
			return nil
		})
		require.NoError(t, err)
		require.Zero(t, visited)
	})

	t.Run("complete records", func(t *testing.T) {
		store := newStore(t)
		for seqNum := 7; seqNum <= 10; seqNum++ {
			require.NoError(t, store.SaveMessage(seqNum, []byte(fmt.Sprintf("stored-%d", seqNum))))
		}
		var messages [][]byte
		err := store.IterateMessages(7, 10, func(message []byte) error {
			messages = append(messages, message)
			return nil
		})
		require.NoError(t, err)
		require.Equal(t, [][]byte{
			[]byte("stored-7"),
			[]byte("stored-8"),
			[]byte("stored-9"),
			[]byte("stored-10"),
		}, messages)
	})

	for _, completed := range []int{0, 2} {
		t.Run(fmt.Sprintf("truncated after %d records", completed), func(t *testing.T) {
			store := newStore(t)
			for seqNum := 7; seqNum <= 10; seqNum++ {
				require.NoError(t, store.SaveMessage(seqNum, []byte(fmt.Sprintf("stored-%d", seqNum))))
			}

			header, err := os.ReadFile(store.headerFname)
			require.NoError(t, err)
			body, err := os.ReadFile(store.bodyFname)
			require.NoError(t, err)
			lines := bytes.Split(header, []byte{'\n'})
			prefix := bytes.Join(lines[:completed], []byte{'\n'})
			if completed > 0 {
				prefix = append(prefix, '\n')
			}
			partial := lines[completed][:bytes.LastIndexByte(lines[completed], ',')+1]
			require.NoError(t, os.WriteFile(store.headerFname, append(prefix, partial...), 0600))

			visited := 0
			err = store.IterateMessages(7, 10, func([]byte) error {
				visited++
				return nil
			})
			require.Error(t, err)
			require.Equal(t, completed, visited)
			after, readErr := os.ReadFile(store.bodyFname)
			require.NoError(t, readErr)
			require.Equal(t, body, after)
		})
	}
}
