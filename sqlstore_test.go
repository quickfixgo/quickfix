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
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// SqlStoreTestSuite runs all tests in the MessageStoreTestSuite against the SqlStore implementation.
type SQLStoreTestSuite struct {
	MessageStoreTestSuite
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
	ddlFnames, err := filepath.Glob(fmt.Sprintf("_sql/%s/*.sql", sqlDriver))
	require.Nil(suite.T(), err)
	for _, fname := range ddlFnames {
		sqlBytes, err := ioutil.ReadFile(fname)
		require.Nil(suite.T(), err)
		_, err = db.Exec(string(sqlBytes))
		require.Nil(suite.T(), err)
	}

	// create settings
	sessionID := SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	settings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`
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
	suite.msgStore, err = NewSQLStoreFactory(settings).Create(sessionID)
	require.Nil(suite.T(), err)
}

func (suite *SQLStoreTestSuite) TestSqlPlaceholderReplacement() {
	got := sqlString("A ? B ? C ?", postgresPlaceholder)
	suite.Equal("A $1 B $2 C $3", got)
}

func (suite *SQLStoreTestSuite) TearDownTest() {
	suite.msgStore.Close()
	os.RemoveAll(suite.sqlStoreRootPath)
}

func TestSqlStoreTestSuite(t *testing.T) {
	suite.Run(t, new(SQLStoreTestSuite))
}
