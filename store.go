package quickfixgo

type MessageStore interface {
	SaveMessage(seqNum int, buf Buffer)
	GetMessages(beginSeqNum, endSeqNum int) chan Buffer
}

type MemoryStore struct {
	messageMap map[int]Buffer
}

func NewMemoryStore() MemoryStore {
	store := MemoryStore{}
	store.messageMap = make(map[int]Buffer)
	return store
}

func (store MemoryStore) SaveMessage(seqNum int, buf Buffer) {
	store.messageMap[seqNum] = buf
}

func (store MemoryStore) GetMessages(beginSeqNum, endSeqNum int) chan Buffer {
	buffers := make(chan Buffer)

	go func() {
		for seqNum := beginSeqNum; seqNum <= endSeqNum; seqNum++ {
			if buffer, ok := store.messageMap[seqNum]; ok {
				buffers <- buffer
			}
		}
		close(buffers)
	}()

	return buffers
}
