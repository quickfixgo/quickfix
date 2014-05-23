package quickfix

import "time"

//The MessageStore interface provides methods to record and retrieve messages for resend purposes
type MessageStore interface {
	NextSenderMsgSeqNum() int
	NextTargetMsgSeqNum() int

	IncrNextSenderMsgSeqNum()
	IncrNextTargetMsgSeqNum()

	SetNextSenderMsgSeqNum(next int)
	SetNextTargetMsgSeqNum(next int)

	CreationTime() time.Time

	SaveMessage(seqNum int, msg []byte)
	GetMessages(beginSeqNum, endSeqNum int) chan []byte

	Refresh()
	Reset()
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

func (store *memoryStore) IncrNextSenderMsgSeqNum() {
	store.senderMsgSeqNum++
}

func (store *memoryStore) IncrNextTargetMsgSeqNum() {
	store.targetMsgSeqNum++
}

func (store *memoryStore) SetNextSenderMsgSeqNum(nextSeqNum int) {
	store.senderMsgSeqNum = nextSeqNum - 1
}
func (store *memoryStore) SetNextTargetMsgSeqNum(nextSeqNum int) {
	store.targetMsgSeqNum = nextSeqNum - 1
}

func (store memoryStore) CreationTime() time.Time {
	return store.creationTime
}

func (store *memoryStore) Reset() {
	store.senderMsgSeqNum = 0
	store.targetMsgSeqNum = 0
	store.creationTime = time.Now()
	store.messageMap = make(map[int][]byte)
}

func (store *memoryStore) Refresh() {
	//nop, nothing to refresh
}

func (store memoryStore) SaveMessage(seqNum int, msg []byte) {
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
	return &memoryStore{messageMap: make(map[int][]byte)}, nil
}

//NewMemoryStoreFactory returns a MessageStoreFactory instance that created in-memory MessageStores
func NewMemoryStoreFactory() MessageStoreFactory { return memoryStoreFactory{} }
