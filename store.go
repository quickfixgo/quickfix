package quickfix

import "time"

//The MessageStore interface provides methods to record and retrieve messages for resend purposes
type MessageStore interface {
	NextSenderMsgSeqNum() int
	NextTargetMsgSeqNum() int

	IncrNextSenderMsgSeqNum() error
	IncrNextTargetMsgSeqNum() error

	SetNextSenderMsgSeqNum(next int) error
	SetNextTargetMsgSeqNum(next int) error

	CreationTime() time.Time

	SaveMessage(seqNum int, msg []byte)
	GetMessages(beginSeqNum, endSeqNum int) chan []byte

	Refresh() error
	Reset() error

	Close() error
}

//The MessageStoreFactory interface is used by session to create a session specific message store
type MessageStoreFactory interface {
	Create(sessionID SessionID) (MessageStore, error)
}

type memoryStore struct {
	senderMsgSeqNum, targetMsgSeqNum int
	creationTime                     time.Time
	messageMap                       map[int][]byte
}

func (store memoryStore) NextSenderMsgSeqNum() int {
	return store.senderMsgSeqNum + 1
}

func (store memoryStore) NextTargetMsgSeqNum() int {
	return store.targetMsgSeqNum + 1
}

func (store *memoryStore) IncrNextSenderMsgSeqNum() error {
	store.senderMsgSeqNum++
	return nil
}

func (store *memoryStore) IncrNextTargetMsgSeqNum() error {
	store.targetMsgSeqNum++
	return nil
}

func (store *memoryStore) SetNextSenderMsgSeqNum(nextSeqNum int) error {
	store.senderMsgSeqNum = nextSeqNum - 1
	return nil
}
func (store *memoryStore) SetNextTargetMsgSeqNum(nextSeqNum int) error {
	store.targetMsgSeqNum = nextSeqNum - 1
	return nil
}

func (store memoryStore) CreationTime() time.Time {
	return store.creationTime
}

func (store *memoryStore) Reset() error {
	store.senderMsgSeqNum = 0
	store.targetMsgSeqNum = 0
	store.creationTime = time.Now()
	store.messageMap = nil
	return nil
}

func (store *memoryStore) Refresh() error {
	//nop, nothing to refresh
	return nil
}

func (store *memoryStore) Close() error {
	//nop, nothing to close
	return nil
}

func (store *memoryStore) SaveMessage(seqNum int, msg []byte) {
	if store.messageMap == nil {
		store.messageMap = make(map[int][]byte)
	}

	store.messageMap[seqNum] = msg
}

func (store memoryStore) GetMessages(beginSeqNum, endSeqNum int) chan []byte {
	msgs := make(chan []byte)

	go func() {
		for seqNum := beginSeqNum; seqNum <= endSeqNum; seqNum++ {
			if msg, ok := store.messageMap[seqNum]; ok {
				msgs <- msg
			}
		}
		close(msgs)
	}()

	return msgs
}

type memoryStoreFactory struct{}

func (f memoryStoreFactory) Create(sessionID SessionID) (MessageStore, error) {
	var m memoryStore
	m.Reset()
	return &m, nil
}

//NewMemoryStoreFactory returns a MessageStoreFactory instance that created in-memory MessageStores
func NewMemoryStoreFactory() MessageStoreFactory { return memoryStoreFactory{} }
