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
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// SQLLogTestSuite runs tests for the SQLLog impl of Log.
type SQLLogTestSuite struct {
	suite.Suite
	sqlLogRootPath string
	log            *sqlLog

	settings  *quickfix.Settings
	sessionID quickfix.SessionID
}

func (suite *SQLLogTestSuite) SetupTest() {
	suite.sqlLogRootPath = path.Join(os.TempDir(), fmt.Sprintf("SQLLogTestSuite-%d", os.Getpid()))
	err := os.MkdirAll(suite.sqlLogRootPath, os.ModePerm)
	require.Nil(suite.T(), err)
	sqlDriver := "sqlite3"
	sqlDsn := path.Join(suite.sqlLogRootPath, fmt.Sprintf("%d.db", time.Now().UnixNano()))

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
SQLLogDriver=%s
SQLLogDataSourceName=%s
SQLLogConnMaxLifetime=14400s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, sqlDriver, sqlDsn, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	suite.sessionID = sessionID
	suite.settings = settings
}

func (suite *SQLLogTestSuite) TestSQLLogNoSession() {
	// create log
	log, err := NewLogFactory(suite.settings).Create()
	require.Nil(suite.T(), err)
	suite.log = log.(*sqlLog)

	suite.log.OnIncoming([]byte("Cool1"))
	suite.log.OnOutgoing([]byte("Cool2"))
	entries, err := suite.log.getEntries("messages_log")
	require.Nil(suite.T(), err)
	require.Len(suite.T(), entries, 2)
	require.Equal(suite.T(), "Cool1", entries[0])
	require.Equal(suite.T(), "Cool2", entries[1])

	suite.log.OnEvent("Cool3")
	suite.log.OnEvent("Cool4")
	entries, err = suite.log.getEntries("event_log")
	require.Nil(suite.T(), err)
	require.Len(suite.T(), entries, 2)
	require.Equal(suite.T(), "Cool3", entries[0])
	require.Equal(suite.T(), "Cool4", entries[1])
}

func (suite *SQLLogTestSuite) TestSQLLog() {
	// create log
	log, err := NewLogFactory(suite.settings).CreateSessionLog(suite.sessionID)
	require.Nil(suite.T(), err)
	suite.log = log.(*sqlLog)

	suite.log.OnIncoming([]byte("Cool1"))
	suite.log.OnOutgoing([]byte("Cool2"))
	entries, err := suite.log.getEntries("messages_log")
	require.Nil(suite.T(), err)
	require.Len(suite.T(), entries, 2)
	require.Equal(suite.T(), "Cool1", entries[0])
	require.Equal(suite.T(), "Cool2", entries[1])

	suite.log.OnEvent("Cool3")
	suite.log.OnEvent("Cool4")
	entries, err = suite.log.getEntries("event_log")
	require.Nil(suite.T(), err)
	require.Len(suite.T(), entries, 2)
	require.Equal(suite.T(), "Cool3", entries[0])
	require.Equal(suite.T(), "Cool4", entries[1])
}

func (suite *SQLLogTestSuite) TestSqlPlaceholderReplacement() {
	got := sqlString("A ? B ? C ?", postgresPlaceholder)
	suite.Equal("A $1 B $2 C $3", got)
}

func (suite *SQLLogTestSuite) TearDownTest() {
	suite.log.close()
	os.RemoveAll(suite.sqlLogRootPath)
}

func TestSQLLogTestSuite(t *testing.T) {
	suite.Run(t, new(SQLLogTestSuite))
}
