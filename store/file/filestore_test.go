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
