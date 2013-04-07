package store

import (
	"github.com/cbusbey/quickfixgo/message"
)

type MessageStore interface {
	SaveMessage(seqNum int, buf message.Buffer)
	GetMessages(beginSeqNum, endSeqNum int) chan message.Buffer
}

type MemoryStore struct {
	messageMap map[int]message.Buffer
}

func NewMemoryStore() MemoryStore {
	store := MemoryStore{}
	store.messageMap = make(map[int]message.Buffer)
	return store
}

func (store MemoryStore) SaveMessage(seqNum int, buf message.Buffer) {
	store.messageMap[seqNum] = buf
}

func (store MemoryStore) GetMessages(beginSeqNum, endSeqNum int) chan message.Buffer {
	buffers := make(chan message.Buffer)

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
