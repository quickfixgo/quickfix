package quickfix

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// FileStoreTestSuite runs all tests in the MessageStoreTestSuite against the FileStore implementation
type FileStoreTestSuite struct {
	MessageStoreTestSuite
	fileStoreRootPath string
}

func (suite *FileStoreTestSuite) SetupTest() {
	suite.fileStoreRootPath = path.Join(os.TempDir(), fmt.Sprintf("FileStoreTestSuite-%d", os.Getpid()))
	fileStorePath := path.Join(suite.fileStoreRootPath, fmt.Sprintf("%d", time.Now().UnixNano()))
	sessionID := SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}

	// create settings
	settings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`
[DEFAULT]
FileStorePath=%s

[SESSION]
BeginString=%s
SenderCompID=%s
TargetCompID=%s`, fileStorePath, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	require.Nil(suite.T(), err)

	// create store
	suite.msgStore, err = NewFileStoreFactory(settings).Create(sessionID)
	require.Nil(suite.T(), err)
}

func (suite *FileStoreTestSuite) TearDownTest() {
	suite.msgStore.Close()
	os.RemoveAll(suite.fileStoreRootPath)
}

func TestFileStoreTestSuite(t *testing.T) {
	suite.Run(t, new(FileStoreTestSuite))
}
