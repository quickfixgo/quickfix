package quickfix

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/quickfixgo/quickfix/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type helper struct {
	SessionID         SessionID
	FileStore         MessageStore
	fileStoreRootPath string
}

func newFileStoreTestHelper(t *testing.T, sessionQualifier string) *helper {
	sessionID := SessionID{
		BeginString:  "FIX.4.4",
		SenderCompID: "SENDER",
		TargetCompID: "TARGET",
		Qualifier:    sessionQualifier,
	}

	settings := NewSessionSettings()
	fileStoreRootPath := path.Join(os.TempDir(), fmt.Sprintf("TestFileStore-%d", os.Getpid()))
	settings.Set(config.FileStorePath, path.Join(fileStoreRootPath, sessionID.Qualifier))

	fileStore, err := NewFileStoreFactory(settings).Create(sessionID)
	require.Nil(t, err)

	return &helper{
		SessionID:         sessionID,
		FileStore:         fileStore,
		fileStoreRootPath: fileStoreRootPath,
	}
}

func (helper *helper) Close() {
	helper.FileStore.Close()
	os.RemoveAll(helper.fileStoreRootPath)
}

func TestFileStore_SetNextMsgSeqNum_Refresh_IncrNextMsgSeqNum(t *testing.T) {
	// Given a FileStore with the following sender and target seqnums
	helper := newFileStoreTestHelper(t, "TestFileStore_SetNextMsgSeqNum_Refresh_IncrNextMsgSeqNum")
	defer helper.Close()
	helper.FileStore.SetNextSenderMsgSeqNum(867)
	helper.FileStore.SetNextTargetMsgSeqNum(5309)

	// When the store is closed and re-opened
	helper.FileStore.Close()
	helper = newFileStoreTestHelper(t, "TestFileStore_SetNextMsgSeqNum_Refresh_IncrNextMsgSeqNum")
	defer helper.Close()

	// Then the sender and target seqnums should still be
	assert.Equal(t, 867, helper.FileStore.NextSenderMsgSeqNum())
	assert.Equal(t, 5309, helper.FileStore.NextTargetMsgSeqNum())

	// When the sender and target seqnums are incremented
	require.Nil(t, helper.FileStore.IncrNextSenderMsgSeqNum())
	require.Nil(t, helper.FileStore.IncrNextTargetMsgSeqNum())

	// Then the sender and target seqnums should be
	assert.Equal(t, 868, helper.FileStore.NextSenderMsgSeqNum())
	assert.Equal(t, 5310, helper.FileStore.NextTargetMsgSeqNum())

	// When the store is closed and re-opened
	helper.FileStore.Close()
	helper = newFileStoreTestHelper(t, "TestFileStore_SetNextMsgSeqNum_Refresh_IncrNextMsgSeqNum")
	defer helper.Close()

	// Then the sender and target seqnums should still be
	assert.Equal(t, 868, helper.FileStore.NextSenderMsgSeqNum())
	assert.Equal(t, 5310, helper.FileStore.NextTargetMsgSeqNum())
}

func TestFileStore_Reset(t *testing.T) {
	// Given a FileStore with the following sender and target seqnums
	helper := newFileStoreTestHelper(t, "TestFileStore_Reset")
	defer helper.Close()
	helper.FileStore.SetNextSenderMsgSeqNum(1234)
	helper.FileStore.SetNextTargetMsgSeqNum(5678)

	// When the store is reset
	require.Nil(t, helper.FileStore.Reset())

	// Then the sender and target seqnums should be
	assert.Equal(t, 1, helper.FileStore.NextSenderMsgSeqNum())
	assert.Equal(t, 1, helper.FileStore.NextTargetMsgSeqNum())

	// When the store is closed and re-opened
	helper.FileStore.Close()
	helper = newFileStoreTestHelper(t, "TestFileStore_Reset")
	defer helper.Close()

	// Then the sender and target seqnums should still be
	assert.Equal(t, 1, helper.FileStore.NextSenderMsgSeqNum())
	assert.Equal(t, 1, helper.FileStore.NextTargetMsgSeqNum())
}
