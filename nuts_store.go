package quickfix

import (
	"encoding/binary"
	"time"

	"github.com/nutsdb/nutsdb"
	"github.com/pkg/errors"
)

var (
	keyMessages       = []byte("messages")
	keyOutgoingSeqnum = []byte("outgoing_seqnum")
	keyIncomingSeqnum = []byte("incoming_seqnum")
)

type nutsDbStoreFactory struct {
	db *nutsdb.DB
}

func NewNutsDbStoreFactory(db *nutsdb.DB) MessageStoreFactory {
	return nutsDbStoreFactory{db: db}
}

func (f nutsDbStoreFactory) Create(sessionID SessionID) (msgStore MessageStore, err error) {
	sessionPrefix := SessionIDFilenamePrefix(sessionID)
	store := &nutsDbStore{
		db:     f.db,
		cache:  &memoryStore{},
		bucket: sessionPrefix,
	}
	if err = store.cache.Reset(); err != nil {
		err = errors.Wrap(err, "cache reset")
		return
	}
	if err = store.populateCache(); err != nil {
		return nil, err
	}
	return store, nil
}

type nutsDbStore struct {
	cache *memoryStore

	bucket string
	db     *nutsdb.DB
}

func (store *nutsDbStore) NextSenderMsgSeqNum() int {
	return store.cache.NextSenderMsgSeqNum()
}

func (store *nutsDbStore) NextTargetMsgSeqNum() int {
	return store.cache.NextTargetMsgSeqNum()
}

func (store *nutsDbStore) IncrNextSenderMsgSeqNum() error {
	if err := store.cache.IncrNextSenderMsgSeqNum(); err != nil {
		return errors.Wrap(err, "cache incr next")
	}
	return store.SetNextSenderMsgSeqNum(store.cache.NextSenderMsgSeqNum())
}

func (store *nutsDbStore) IncrNextTargetMsgSeqNum() error {
	if err := store.cache.IncrNextTargetMsgSeqNum(); err != nil {
		return errors.Wrap(err, "cache incr next")
	}
	return store.SetNextTargetMsgSeqNum(store.cache.NextTargetMsgSeqNum())
}

func (store *nutsDbStore) SetNextSenderMsgSeqNum(next int) error {

	err := store.db.Update(func(tx *nutsdb.Tx) error {
		nextBuf := make([]byte, 8)
		binary.BigEndian.PutUint64(nextBuf, uint64(next))
		return tx.Put(store.bucket, keyOutgoingSeqnum, nextBuf, nutsdb.Persistent)
	})

	if err != nil {
		return err
	}
	return store.cache.SetNextSenderMsgSeqNum(next)
}
func (store *nutsDbStore) SetNextTargetMsgSeqNum(next int) error {

	err := store.db.Update(func(tx *nutsdb.Tx) error {
		nextBuf := make([]byte, 8)
		binary.BigEndian.PutUint64(nextBuf, uint64(next))
		return tx.Put(store.bucket, keyIncomingSeqnum, nextBuf, nutsdb.Persistent)
	})
	if err != nil {
		return err
	}
	return store.cache.SetNextTargetMsgSeqNum(next)
}

func (store *nutsDbStore) CreationTime() time.Time {
	return store.cache.CreationTime()
}

func (store *nutsDbStore) Reset() error {
	return store.db.Update(func(tx *nutsdb.Tx) error {
		// err := tx.Delete(store.bucket, keyMessages)
		// if err != nil {
		// 	return err
		// }
		err := tx.DeleteBucket(nutsdb.DataStructureList, store.bucket)
		if err = store.cache.Reset(); err != nil {
			return err
		}
		nextBuf := make([]byte, 8)
		binary.BigEndian.PutUint64(nextBuf, uint64(store.cache.NextSenderMsgSeqNum()))
		err = tx.Put(store.bucket, keyOutgoingSeqnum, nextBuf, nutsdb.Persistent)
		if err != nil {
			return err
		}
		binary.BigEndian.PutUint64(nextBuf, uint64(store.cache.NextTargetMsgSeqNum()))
		return tx.Put(store.bucket, keyIncomingSeqnum, nextBuf, nutsdb.Persistent)
	})
}

func (store *nutsDbStore) Refresh() error {
	if err := store.cache.Reset(); err != nil {
		return err
	}
	return store.populateCache()

}
func (store *nutsDbStore) populateCache() error {
	return store.db.Update(func(tx *nutsdb.Tx) error {
		var incomingSeqNum, outgoingSeqNum int

		if e, err := tx.Get(store.bucket, keyOutgoingSeqnum); err == nil {
			outgoingSeqNum = int(binary.BigEndian.Uint64(e.Value))
			if err := store.cache.SetNextSenderMsgSeqNum(outgoingSeqNum); err != nil {
				return errors.Wrap(err, "cache set next sender")
			}
		} else {
			nextBuf := make([]byte, 8)
			binary.BigEndian.PutUint64(nextBuf, uint64(store.cache.NextSenderMsgSeqNum()))
			if err := tx.Put(store.bucket, keyOutgoingSeqnum, nextBuf, nutsdb.Persistent); err != nil {
				return err
			}

		}

		if e, err := tx.Get(store.bucket, keyIncomingSeqnum); err == nil {
			incomingSeqNum = int(binary.BigEndian.Uint64(e.Value))
			if err := store.cache.SetNextTargetMsgSeqNum(incomingSeqNum); err != nil {
				return errors.Wrap(err, "cache set next target")
			}
		} else {
			nextBuf := make([]byte, 8)
			binary.BigEndian.PutUint64(nextBuf, uint64(store.cache.NextTargetMsgSeqNum()))
			if err := tx.Put(store.bucket, keyIncomingSeqnum, nextBuf, nutsdb.Persistent); err != nil {
				return err
			}
		}

		return nil
	})
}

func (store *nutsDbStore) SaveMessage(seqNum int, msg []byte) error {
	return store.db.Update(func(tx *nutsdb.Tx) error {
		return tx.RPush(store.bucket, keyMessages, msg)
	})
}

func (store *nutsDbStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	var msgs [][]byte
	store.db.View(func(tx *nutsdb.Tx) error {
		msgs, _ = tx.LRange(store.bucket, keyMessages, beginSeqNum, endSeqNum)
		return nil
	})
	return msgs, nil
}

func (store *nutsDbStore) Close() error {
	if store.db != nil {
		return store.db.Close()
	}
	return nil
}
