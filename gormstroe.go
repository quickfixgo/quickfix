package quickfix

import (
	"time"

	"github.com/jinzhu/gorm"
)

type gormStoreFactory struct {
	settings *Settings
}

func NewGormStoreFactory(settings *Settings) MessageStoreFactory {
	return gormStoreFactory{settings: settings}
}

type gromStore struct {
	sessionID   SessionID
	cache       *memoryStore
	db          *gorm.DB
	placeholder placeholderFunc
}

func (f gormStoreFactory) Create(sessionID SessionID) (msgStore MessageStore, err error) {
	panic(`todo conn db`)
}

// Reset deletes the store records and sets the seqnums back to 1
func (store *gromStore) Reset() error {
	panic(`todo`)
}

// Refresh reloads the store from the database
func (store *gromStore) Refresh() error {
	if err := store.cache.Reset(); err != nil {
		return err
	}
	return store.populateCache()
}

func (store *gromStore) populateCache() error {
	panic(`todo`)
}

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent
func (store *gromStore) NextSenderMsgSeqNum() int {
	return store.cache.NextSenderMsgSeqNum()
}

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received
func (store *gromStore) NextTargetMsgSeqNum() int {
	return store.cache.NextTargetMsgSeqNum()
}

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent
func (store *gromStore) SetNextSenderMsgSeqNum(next int) error {
	panic(`todo`)
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received
func (store *gromStore) SetNextTargetMsgSeqNum(next int) error {
	panic(`todo`)
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent
func (store *gromStore) IncrNextSenderMsgSeqNum() error {
	panic(`todo`)
}

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (store *gromStore) IncrNextTargetMsgSeqNum() error {
	panic(`todo`)
}

// CreationTime returns the creation time of the store
func (store *gromStore) CreationTime() time.Time {
	return store.cache.CreationTime()
}

func (store *gromStore) SaveMessage(seqNum int, msg []byte) error {
	panic(`todo`)
}

func (store *gromStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	panic(`todo`)
}

// Close closes the store's database connection
func (store *gromStore) Close() error {
	if store.db != nil {
		store.db.Close()
		store.db = nil
	}
	return nil
}
