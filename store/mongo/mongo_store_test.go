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
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/internal/testsuite"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func newConcurrentMongoStore(t *testing.T, monitor *event.CommandMonitor) *mongoStore {
	t.Helper()
	uri := os.Getenv("MONGODB_TEST_CXN")
	if uri == "" {
		t.Skip("MONGODB_TEST_CXN is not set")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri).
		SetReplicaSet("replicaset").SetMonitor(monitor).SetTimeout(10*time.Second))
	require.NoError(t, err)
	sessionID := quickfix.SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	cache, err := quickfix.NewMemoryStoreFactory().Create(sessionID)
	require.NoError(t, err)
	store := &mongoStore{
		sessionID:          sessionID,
		cache:              cache,
		db:                 client,
		mongoDatabase:      "quickfix_concurrency_" + primitive.NewObjectID().Hex(),
		messagesCollection: "messages",
		sessionsCollection: "sessions",
		allowTransactions:  true,
	}
	t.Cleanup(func() {
		require.NoError(t, client.Database(store.mongoDatabase).Drop(context.Background()))
		require.NoError(t, store.Close())
	})
	require.NoError(t, store.populateCache())
	require.NoError(t, client.Database(store.mongoDatabase).CreateCollection(context.Background(), store.messagesCollection))
	return store
}

func waitForMongoOperation(t *testing.T, result <-chan error) error {
	t.Helper()
	select {
	case err := <-result:
		return err
	case <-time.After(10 * time.Second):
		t.Fatal("MongoStore operation did not complete")
		return nil
	}
}

func TestMongoStoreConcurrentSendAndReceive(t *testing.T) {
	for _, transactions := range []bool{true, false} {
		t.Run(fmt.Sprintf("transactions=%t", transactions), func(t *testing.T) {
			testMongoStoreConcurrentSend(t, transactions, (*mongoStore).IncrNextTargetMsgSeqNum, 2, 2)
		})
	}
}

func TestMongoStoreConcurrentResetAndRefresh(t *testing.T) {
	t.Run("reset", func(t *testing.T) {
		testMongoStoreConcurrentSend(t, true, (*mongoStore).Reset, 1, 1)
	})
	t.Run("refresh", func(t *testing.T) {
		testMongoStoreConcurrentSend(t, true, (*mongoStore).Refresh, 2, 1)
	})
}

func testMongoStoreConcurrentSend(t *testing.T, transactions bool, update func(*mongoStore) error, wantSender, wantTarget int) {
	t.Helper()
	inserted := make(chan struct{})
	resume := make(chan struct{})
	release := sync.OnceFunc(func() { close(resume) })
	defer release()
	var pauseInsert atomic.Bool
	store := newConcurrentMongoStore(t, &event.CommandMonitor{
		Succeeded: func(_ context.Context, ev *event.CommandSucceededEvent) {
			if ev.CommandName == "insert" && pauseInsert.CompareAndSwap(true, false) {
				close(inserted)
				<-resume
			}
		},
	})
	store.allowTransactions = transactions
	pauseInsert.Store(true)
	sent := make(chan error, 1)
	go func() { sent <- store.SaveMessageAndIncrNextSenderMsgSeqNum(1, []byte("message")) }()
	select {
	case <-inserted:
	case <-time.After(10 * time.Second):
		t.Fatal("message insert did not complete")
	}

	// With transactions, the insert establishes the sender's snapshot.
	// Other session updates must wait until the entire send completes.
	received := make(chan error, 1)
	started := make(chan struct{})
	go func() {
		close(started)
		received <- update(store)
	}()
	<-started
	var receiveErr error
	var receivedBeforeRelease bool
	select {
	case receiveErr = <-received:
		receivedBeforeRelease = true
	case <-time.After(100 * time.Millisecond):
	}
	release()
	sendErr := waitForMongoOperation(t, sent)
	if !receivedBeforeRelease {
		receiveErr = waitForMongoOperation(t, received)
	}
	require.NoError(t, sendErr)
	require.NoError(t, receiveErr)
	require.False(t, receivedBeforeRelease, "session update completed before the send")
	require.Equal(t, wantSender, store.NextSenderMsgSeqNum())
	require.Equal(t, wantTarget, store.NextTargetMsgSeqNum())
	require.NoError(t, store.Refresh())
	require.Equal(t, wantSender, store.NextSenderMsgSeqNum())
	require.Equal(t, wantTarget, store.NextTargetMsgSeqNum())
	messages, err := store.GetMessages(1, 1)
	require.NoError(t, err)
	if wantSender == 1 {
		require.Empty(t, messages)
	} else {
		require.Equal(t, [][]byte{[]byte("message")}, messages)
	}
}

func TestMongoStoreConcurrentSequenceNumbers(t *testing.T) {
	for _, transactions := range []bool{true, false} {
		t.Run(fmt.Sprintf("transactions=%t", transactions), func(t *testing.T) {
			store := newConcurrentMongoStore(t, nil)
			store.allowTransactions = transactions
			const count = 32
			messages := make([][]byte, count)
			for i := range messages {
				messages[i] = []byte(fmt.Sprintf("message %d", i+1))
			}
			results := make(chan error, 2)
			go func() {
				for i, msg := range messages {
					if err := store.SaveMessageAndIncrNextSenderMsgSeqNum(i+1, msg); err != nil {
						results <- err
						return
					}
				}
				results <- nil
			}()
			go func() {
				for i := 0; i < count; i++ {
					if err := store.IncrNextTargetMsgSeqNum(); err != nil {
						results <- err
						return
					}
				}
				results <- nil
			}()
			stopReading := make(chan struct{})
			readerDone := make(chan struct{})
			stop := sync.OnceFunc(func() { close(stopReading) })
			defer stop()
			go func() {
				defer close(readerDone)
				for {
					select {
					case <-stopReading:
						return
					default:
						store.NextSenderMsgSeqNum()
						store.NextTargetMsgSeqNum()
						store.CreationTime()
					}
				}
			}()
			firstErr := waitForMongoOperation(t, results)
			secondErr := waitForMongoOperation(t, results)
			stop()
			<-readerDone
			require.NoError(t, firstErr)
			require.NoError(t, secondErr)
			require.Equal(t, count+1, store.NextSenderMsgSeqNum())
			require.Equal(t, count+1, store.NextTargetMsgSeqNum())
			require.NoError(t, store.Refresh())
			require.Equal(t, count+1, store.NextSenderMsgSeqNum())
			require.Equal(t, count+1, store.NextTargetMsgSeqNum())
			actual, err := store.GetMessages(1, count)
			require.NoError(t, err)
			require.Equal(t, messages, actual)
		})
	}
}

func TestMongoStoreSendFailureUnlocks(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	for step, command := range []string{"insert", "update", "commitTransaction"} {
		mt.Run(command, func(mt *mtest.T) {
			cache, err := quickfix.NewMemoryStoreFactory().Create(quickfix.SessionID{})
			require.NoError(mt, err)
			store := &mongoStore{
				cache: cache, db: mt.Client, mongoDatabase: mt.DB.Name(),
				messagesCollection: "messages", sessionsCollection: "sessions", allowTransactions: true,
			}
			for i := 0; i < step; i++ {
				mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "n", Value: 1}))
			}
			mt.AddMockResponses(
				mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 42, Message: "send failed"}),
				mtest.CreateSuccessResponse(), // Abort the failed transaction, if necessary.
			)
			err = store.SaveMessageAndIncrNextSenderMsgSeqNum(1, []byte("message"))
			var commandErr mongo.CommandError
			require.ErrorAs(mt, err, &commandErr)
			require.EqualValues(mt, 42, commandErr.Code)
			mt.ClearMockResponses()
			mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "n", Value: 1}))
			result := make(chan error, 1)
			go func() { result <- store.IncrNextTargetMsgSeqNum() }()
			require.NoError(mt, waitForMongoOperation(mt.T, result))
			require.Equal(mt, 1, store.NextSenderMsgSeqNum())
			require.Equal(mt, 2, store.NextTargetMsgSeqNum())
		})
	}
}
