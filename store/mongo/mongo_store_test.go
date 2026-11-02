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
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/internal/testsuite"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// MongoStoreTestSuite runs all tests in the message.StoreTestSuite against the MongoStore implementation.
type MongoStoreTestSuite struct {
	testsuite.StoreTestSuite
}

func (suite *MongoStoreTestSuite) SetupTest() {
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
MongoStoreConnection=%s
MongoStoreDatabase=%s
MongoStoreReplicaSet=%s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, mongoDbCxn, mongoDatabase, mongoReplicaSet, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	// create store
	suite.MsgStore, err = NewStoreFactory(settings).Create(sessionID)
	require.Nil(suite.T(), err)
	err = suite.MsgStore.Reset()
	require.Nil(suite.T(), err)
}

func (suite *MongoStoreTestSuite) TearDownTest() {
	if suite.MsgStore != nil {
		err := suite.MsgStore.Close()
		require.Nil(suite.T(), err)
	}
}

func TestMongoStoreTestSuite(t *testing.T) {
	suite.Run(t, new(MongoStoreTestSuite))
}

func TestIterateMessagesCursorError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	for _, count := range []int{0, 2} {
		mt.Run(fmt.Sprintf("after_%d_messages", count), func(mt *mtest.T) {
			store := &mongoStore{
				db:                 mt.Client,
				mongoDatabase:      mt.DB.Name(),
				messagesCollection: mt.Coll.Name(),
			}
			batch := make([]bson.D, count)
			for i := range batch {
				batch[i] = bson.D{{Key: "msgseq", Value: 7 + i}, {Key: "message", Value: []byte("stored message")}}
			}
			mt.AddMockResponses(
				mtest.CreateCursorResponse(123, mt.DB.Name()+"."+mt.Coll.Name(), mtest.FirstBatch, batch...),
				mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 43, Name: "CursorNotFound", Message: "injected cursor failure"}),
				mtest.CreateSuccessResponse(),
			)

			visited := 0
			err := store.IterateMessages(7, 10, func([]byte) error {
				visited++
				return nil
			})
			require.Equal(mt, count, visited)
			require.ErrorContains(mt, err, "injected cursor failure")
		})
	}
}

func TestIterateMessagesCursorExhausted(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("success", func(mt *mtest.T) {
		store := &mongoStore{
			db:                 mt.Client,
			mongoDatabase:      mt.DB.Name(),
			messagesCollection: mt.Coll.Name(),
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			mt.DB.Name()+"."+mt.Coll.Name(),
			mtest.FirstBatch,
			bson.D{{Key: "msgseq", Value: 7}, {Key: "message", Value: []byte("stored message")}},
		))

		visited := 0
		err := store.IterateMessages(7, 10, func([]byte) error {
			visited++
			return nil
		})
		require.NoError(mt, err)
		require.Equal(mt, 1, visited)
	})
}
