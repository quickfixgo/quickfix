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

package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/quickfixgo/quickfix"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MongoLogTestSuite runs tests for the MongoLog impl of Log.
type MongoLogTestSuite struct {
	suite.Suite
	log       *mongoLog
	settings  *quickfix.Settings
	sessionID quickfix.SessionID
}

func (suite *MongoLogTestSuite) SetupTest() {
	mongoDbCxn := os.Getenv("MONGODB_TEST_CXN")
	if len(mongoDbCxn) <= 0 {
		log.Println("MONGODB_TEST_CXN environment arg is not provided, skipping...")
		suite.T().SkipNow()
	}
	mongoDatabase := "automated_testing_database"
	mongoReplicaSet := "replicaset"

	// create settings
	sessionID := quickfix.SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	settings, err := quickfix.ParseSettings(strings.NewReader(fmt.Sprintf(`
[DEFAULT]
MongoLogConnection=%s
MongoLogDatabase=%s
MongoLogReplicaSet=%s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, mongoDbCxn, mongoDatabase, mongoReplicaSet, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	suite.sessionID = sessionID
	suite.settings = settings
}

func (suite *MongoLogTestSuite) TestMongoLogNoSession() {
	// create log
	log, err := NewLogFactory(suite.settings).Create()
	require.Nil(suite.T(), err)
	suite.log = log.(*mongoLog)

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

func (suite *MongoLogTestSuite) TestMongoLog() {
	// create log
	log, err := NewLogFactory(suite.settings).CreateSessionLog(suite.sessionID)
	require.Nil(suite.T(), err)
	suite.log = log.(*mongoLog)

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

func (suite *MongoLogTestSuite) TearDownTest() {
	entry := generateEntry(&suite.log.sessionID)
	_, err := suite.log.db.Database(suite.log.mongoDatabase).Collection(suite.log.messagesLogCollection).DeleteMany(context.Background(), entry)
	require.Nil(suite.T(), err)

	entry2 := generateEntry(&suite.log.sessionID)
	_, err = suite.log.db.Database(suite.log.mongoDatabase).Collection(suite.log.eventLogCollection).DeleteMany(context.Background(), entry2)
	require.Nil(suite.T(), err)

	err = suite.log.close()
	require.Nil(suite.T(), err)
}

func TestMongoLogTestSuite(t *testing.T) {
	suite.Run(t, new(MongoLogTestSuite))
}
