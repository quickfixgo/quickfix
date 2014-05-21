package quickfix

type messageStore interface {
	SaveMessage(seqNum int, msg []byte)
	GetMessages(beginSeqNum, endSeqNum int) chan []byte
}

type memoryStore struct {
	messageMap map[int][]byte
}

func newMemoryStore() memoryStore {
	store := memoryStore{}
	store.messageMap = make(map[int][]byte)
	return store
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
