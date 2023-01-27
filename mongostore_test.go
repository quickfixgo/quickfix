package quickfix

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MongoStoreTestSuite runs all tests in the MessageStoreTestSuite against the MongoStore implementation
type MongoStoreTestSuite struct {
	MessageStoreTestSuite
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
	sessionID := SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	settings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`
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
	suite.msgStore, err = NewMongoStoreFactory(settings).Create(sessionID)
	require.Nil(suite.T(), err)
	err = suite.msgStore.Reset()
	require.Nil(suite.T(), err)
}

func (suite *MongoStoreTestSuite) TearDownTest() {
	if suite.msgStore != nil {
		err := suite.msgStore.Close()
		require.Nil(suite.T(), err)
	}
}

func TestMongoStoreTestSuite(t *testing.T) {
	suite.Run(t, new(MongoStoreTestSuite))
}
