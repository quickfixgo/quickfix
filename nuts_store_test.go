package quickfix

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/nutsdb/nutsdb"
	"github.com/stretchr/testify/require"
)

func TestSetup(t *testing.T) {
	// create settings
	sessionID := SessionID{BeginString: "FIX.4.4", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	// 	settings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`
	// [DEFAULT]
	// TimeStampPrecision=MICROS

	// [SESSION]
	// Tenant=lb_hk
	// BeginString=%s
	// SenderCompID=%s
	// TargetCompID=%s`, sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID)))
	// require.Nil(t, err)

	file := path.Join(os.TempDir(), fmt.Sprintf("%d", os.Getpid()))
	file = path.Join(file, "/tmp/nutsdb")

	db, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithSyncEnable(false),
		nutsdb.WithRWMode(nutsdb.MMap),
		nutsdb.WithDir(file),
		nutsdb.WithEntryIdxMode(nutsdb.HintKeyValAndRAMIdxMode),
	)
	require.Nil(t, err)

	defer os.RemoveAll(file)

	messageStoreFactory := NewNutsDbStoreFactory(db)
	store, err := messageStoreFactory.Create(sessionID)
	require.Nil(t, err)
	require.NotNil(t, store)

	for i := 0; i < 10; i++ {
		err = store.SaveMessage(i, []byte(fmt.Sprintf("hello_%v", i)))
		require.Nil(t, err)
		err = store.IncrNextSenderMsgSeqNum()
		require.Nil(t, err)
	}

	bs, err := store.GetMessages(0, 1)
	require.Nil(t, err)
	require.Equal(t, len(bs), 2)

	bs, err = store.GetMessages(1, 10)
	require.Nil(t, err)
	require.Equal(t, len(bs), 9)

	err = store.Reset()
	require.Nil(t, err)

	bs, err = store.GetMessages(0, 1)
	require.Nil(t, err)
	require.Equal(t, len(bs), 0)

	for i := 0; i < 100; i++ {
		err = store.SaveMessage(i, []byte(fmt.Sprintf("hello_%v", i)))
		require.Nil(t, err)
		err = store.IncrNextSenderMsgSeqNum()
		require.Nil(t, err)
	}

	bs, err = store.GetMessages(0, 100)
	require.Nil(t, err)
	require.Equal(t, len(bs), 100)

	err = store.Close()
	require.Nil(t, err)
}
