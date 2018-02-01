package quickfix

import (
	"time"
)

// NewNoStoreFactory returns an in-memory sequence-only implementation of MessageStoreFactory which does not grow in memory size.
// The NoStore MessageStore implementation is useful for market data publishers.
func NewNoStoreFactory() MessageStoreFactory {
	return noStoreFactory{}
}

type noStoreFactory struct {
	// nope
}

// Create creates a new NoStore implementation of the MessageStore interface
func (nsf noStoreFactory) Create(sessionID SessionID) (MessageStore, error) {
	store := new(noStore)
	store.Reset()
	return store, nil
}

type noStore struct {
	senderMsgSeqNum, targetMsgSeqNum int
	creationTime                     time.Time
}

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent
func (n *noStore) NextSenderMsgSeqNum() int {
	return n.senderMsgSeqNum + 1
}

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received
func (n *noStore) NextTargetMsgSeqNum() int {
	return n.targetMsgSeqNum + 1
}

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (n *noStore) IncrNextTargetMsgSeqNum() error {
	n.targetMsgSeqNum++
	return nil
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent
func (n *noStore) IncrNextSenderMsgSeqNum() error {
	n.senderMsgSeqNum++
	return nil
}

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent
func (n *noStore) SetNextSenderMsgSeqNum(num int) error {
	n.senderMsgSeqNum = num - 1
	return nil
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received
func (n *noStore) SetNextTargetMsgSeqNum(num int) error {
	n.targetMsgSeqNum = num - 1
	return nil
}

// CreationTime returns the creation time of the store
func (n *noStore) CreationTime() time.Time {
	return n.creationTime
}

// SaveMessage is a no-op
func (n *noStore) SaveMessage(seqNum int, msg []byte) error {
	return nil
}

// GetMessages is a no-op, and will always return nil
func (n *noStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	return nil, nil
}

// Refresh is a no-op
func (n *noStore) Refresh() error {
	return nil
}

// Reset sets the seqnums back to 1
func (n *noStore) Reset() error {
	n.senderMsgSeqNum = 0
	n.targetMsgSeqNum = 0
	n.creationTime = time.Now()
	return nil
}

// Close is a no-op
func (n *noStore) Close() error {
	return nil
}
