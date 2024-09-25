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

package composite

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/log/file"
	"github.com/quickfixgo/quickfix/log/mongo"
	"github.com/quickfixgo/quickfix/log/screen"
	"github.com/quickfixgo/quickfix/log/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// CompositeLogTestSuite runs tests for the MongoLog impl of Log.
type CompositeLogTestSuite struct {
	suite.Suite
	sqlLogRootPath string
	settings       *quickfix.Settings
	sessionID      quickfix.SessionID
}

func (suite *CompositeLogTestSuite) SetupTest() {
	mongoDbCxn := os.Getenv("MONGODB_TEST_CXN")
	if len(mongoDbCxn) <= 0 {
		log.Println("MONGODB_TEST_CXN environment arg is not provided, skipping...")
		suite.T().SkipNow()
	}
	mongoDatabase := "automated_testing_database"
	mongoReplicaSet := "replicaset"

	// create settings
	sessionID := quickfix.SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	logPath := path.Join(os.TempDir(), fmt.Sprintf("TestLogStore-%d", os.Getpid()))
	suite.sqlLogRootPath = path.Join(os.TempDir(), fmt.Sprintf("SQLLogTestSuite-%d", os.Getpid()))
	err := os.MkdirAll(suite.sqlLogRootPath, os.ModePerm)
	require.Nil(suite.T(), err)
	sqlDriver := "sqlite3"
	sqlDsn := path.Join(suite.sqlLogRootPath, fmt.Sprintf("%d.db", time.Now().UnixNano()))

	settings, err := quickfix.ParseSettings(strings.NewReader(fmt.Sprintf(`
[DEFAULT]
MongoLogConnection=%s
MongoLogDatabase=%s
MongoLogReplicaSet=%s
FileLogPath=%s
SQLLogDriver=%s
SQLLogDataSourceName=%s
SQLLogConnMaxLifetime=14400s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, mongoDbCxn, mongoDatabase, mongoReplicaSet, logPath, sqlDriver, sqlDsn, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	suite.sessionID = sessionID
	suite.settings = settings
}

func (suite *CompositeLogTestSuite) TestCreateLogNoSession() {

	mngoLogFactory := mongo.NewLogFactory(suite.settings)
	sqlLogFactory := sql.NewLogFactory(suite.settings)
	// create log
	_, err := NewLogFactory([]quickfix.LogFactory{mngoLogFactory, sqlLogFactory}).Create()
	require.Nil(suite.T(), err)
}

func (suite *CompositeLogTestSuite) TestCreateLogSession() {

	screenLogFactory := screen.NewLogFactory()
	fileLogFactory, err := file.NewLogFactory(suite.settings)
	require.Nil(suite.T(), err)

	// create log
	_, err = NewLogFactory([]quickfix.LogFactory{screenLogFactory, fileLogFactory}).CreateSessionLog(suite.sessionID)
	require.Nil(suite.T(), err)
}

func (suite *CompositeLogTestSuite) TearDownTest() {
	os.RemoveAll(suite.sqlLogRootPath)
}

func TestCompositeLogTestSuite(t *testing.T) {
	suite.Run(t, new(CompositeLogTestSuite))
}
