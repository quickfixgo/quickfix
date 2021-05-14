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

// SpannerStoreTestSuite runs all tests in the MessageStoreTestSuite against the SpannerStore implementation
type SpannerStoreTestSuite struct {
	MessageStoreTestSuite
}

func (suite *SpannerStoreTestSuite) SetupTest() {
	spannerDataSource := os.Getenv("SPANNER_TEST_DATA_SOURCE")
	if len(spannerDataSource) <= 0 {
		log.Println("SPANNER_TEST_DATA_SOURCE environment arg is not provided, skipping...")
		suite.T().SkipNow()
	}

	// create settings
	sessionID := SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	settings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`
[DEFAULT]
SpannerDataSource=%s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, spannerDataSource, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	// create store
	suite.msgStore, err = NewSpannerStoreFactory(settings).Create(sessionID)
	require.Nil(suite.T(), err)
	err = suite.msgStore.Reset()
	require.Nil(suite.T(), err)
}

func (suite *SpannerStoreTestSuite) TearDownTest() {
	suite.msgStore.Close()
}

func TestSpannerStoreTestSuite(t *testing.T) {
	suite.Run(t, new(SpannerStoreTestSuite))
}
