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

package sql

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/internal/testsuite"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// SqlStoreTestSuite runs all tests in the MessageStoreTestSuite against the SqlStore implementation.
type SQLStoreTestSuite struct {
	testsuite.StoreTestSuite
	sqlStoreRootPath string
}

func (suite *SQLStoreTestSuite) SetupTest() {
	suite.sqlStoreRootPath = path.Join(os.TempDir(), fmt.Sprintf("SqlStoreTestSuite-%d", os.Getpid()))
	err := os.MkdirAll(suite.sqlStoreRootPath, os.ModePerm)
	require.Nil(suite.T(), err)
	sqlDriver := "sqlite3"
	sqlDsn := path.Join(suite.sqlStoreRootPath, fmt.Sprintf("%d.db", time.Now().UnixNano()))

	// create tables
	db, err := sql.Open(sqlDriver, sqlDsn)
	require.Nil(suite.T(), err)
	ddlFnames, err := filepath.Glob(fmt.Sprintf("../../_sql/%s/*.sql", sqlDriver))
	require.Nil(suite.T(), err)
	for _, fname := range ddlFnames {
		sqlBytes, err := os.ReadFile(fname)
		require.Nil(suite.T(), err)
		_, err = db.Exec(string(sqlBytes))
		require.Nil(suite.T(), err)
	}

	// create settings
	sessionID := quickfix.SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	settings, err := quickfix.ParseSettings(strings.NewReader(fmt.Sprintf(`
[DEFAULT]
SQLStoreDriver=%s
SQLStoreDataSourceName=%s
SQLStoreConnMaxLifetime=14400s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, sqlDriver, sqlDsn, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	// create store
	suite.MsgStore, err = NewStoreFactory(settings).Create(sessionID)
	require.Nil(suite.T(), err)
}

func (suite *SQLStoreTestSuite) TestSqlPlaceholderReplacement() {
	got := sqlString("A ? B ? C ?", postgresPlaceholder)
	suite.Equal("A $1 B $2 C $3", got)
}

func (suite *SQLStoreTestSuite) TestStoreTableRenameOverride() {
	sqlDriver := "sqlite3"
	sqlDsn := path.Join(suite.sqlStoreRootPath, fmt.Sprintf("rename-override-%d.db", time.Now().UnixNano()))

	// Create DB with original schema
	db, err := sql.Open(sqlDriver, sqlDsn)
	require.NoError(suite.T(), err)

	ddlFnames, err := filepath.Glob(fmt.Sprintf("../../_sql/%s/*.sql", sqlDriver))
	require.NoError(suite.T(), err)
	for _, fname := range ddlFnames {
		sqlBytes, err := os.ReadFile(fname)
		require.NoError(suite.T(), err)
		_, err = db.Exec(string(sqlBytes))
		require.NoError(suite.T(), err)
	}

	// Rename tables
	_, err = db.Exec(`ALTER TABLE sessions RENAME TO renamed_sessions`)
	require.NoError(suite.T(), err)
	_, err = db.Exec(`ALTER TABLE messages RENAME TO renamed_messages`)
	require.NoError(suite.T(), err)

	// Set config to use renamed tables
	sessionID := quickfix.SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	settings, err := quickfix.ParseSettings(strings.NewReader(fmt.Sprintf(`
[DEFAULT]
SQLStoreDriver=%s
SQLStoreDataSourceName=%s
SQLStoreSessionsTableName=renamed_sessions
SQLStoreMessagesTableName=renamed_messages

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s
`, sqlDriver, sqlDsn, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.NoError(suite.T(), err)

	// Create store with renamed table config
	store, err := NewStoreFactory(settings).Create(sessionID)
	require.NoError(suite.T(), err)
	defer store.Close()

	// SaveMessage + SetNextSenderMsgSeqNum
	msg := []byte("8=FIX.4.4\x019=12\x0135=0\x01")
	require.NoError(suite.T(), store.SaveMessage(1, msg))
	require.NoError(suite.T(), store.SetNextSenderMsgSeqNum(2))
	require.NoError(suite.T(), store.SetNextTargetMsgSeqNum(2))

	// SaveMessageAndIncrNextSenderMsgSeqNum
	require.NoError(suite.T(), store.SaveMessageAndIncrNextSenderMsgSeqNum(2, msg))

	// Get and check sequence numbers
	nextSender := store.NextSenderMsgSeqNum()
	suite.Equal(3, nextSender)
	nextTarget := store.NextTargetMsgSeqNum()
	suite.Equal(2, nextTarget)

	// IterateMessages
	count := 0
	err = store.IterateMessages(1, 2, func(_ []byte) error {
		count++
		return nil
	})
	require.NoError(suite.T(), err)
	suite.Equal(2, count)

	// Reset
	require.NoError(suite.T(), store.Reset())

	// After reset, sequence numbers should be 1
	nextSender = store.NextSenderMsgSeqNum()
	suite.Equal(1, nextSender)
	nextTarget = store.NextTargetMsgSeqNum()
	suite.Equal(1, nextTarget)
}

func (suite *SQLStoreTestSuite) TearDownTest() {
	suite.MsgStore.Close()
	os.RemoveAll(suite.sqlStoreRootPath)
}

func TestSqlStoreTestSuite(t *testing.T) {
	suite.Run(t, new(SQLStoreTestSuite))
}
