package quickfix

type messageStore interface {
	SaveMessage(seqNum int, buf buffer)
	GetMessages(beginSeqNum, endSeqNum int) chan buffer
}

type memoryStore struct {
	messageMap map[int]buffer
}

func newMemoryStore() memoryStore {
	store := memoryStore{}
	store.messageMap = make(map[int]buffer)
	return store
}

func (store memoryStore) SaveMessage(seqNum int, buf buffer) {
	store.messageMap[seqNum] = buf
}

func (store memoryStore) GetMessages(beginSeqNum, endSeqNum int) chan buffer {
	buffers := make(chan buffer)

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
