package quickfix

import (
	"time"
)

const (
	operationSetNextSenderMsgSeqNum int = iota + 1
	operationSetNextTargetMsgSeqNum
	operationSaveMessage
	operationReset
	operationRefresh
)

type persisteMessage struct {
	operation int
	seqNum    int
	msg       []byte
}

type storeBackup struct {
	stop     chan struct{}
	messages chan *persisteMessage
	store    MessageStore
}

func newStoreBackup(store MessageStore) *storeBackup {
	return &storeBackup{stop: make(chan struct{}), messages: make(chan *persisteMessage, 500), store: store}
}

func (s *storeBackup) start() {
	go func() {
		for message := range s.messages {
			switch message.operation {
			case operationSetNextSenderMsgSeqNum:
				if err := s.store.SetNextSenderMsgSeqNum(message.seqNum); err != nil {
				}
			case operationSetNextTargetMsgSeqNum:
				if err := s.store.SetNextTargetMsgSeqNum(message.seqNum); err != nil {
				}
			case operationSaveMessage:
				if err := s.store.SaveMessage(message.seqNum, message.msg); err != nil {
				}
			case operationReset:
				if err := s.store.Reset(); err != nil {
				}
			case operationRefresh:
				if err := s.store.Refresh(); err != nil {
				}
			default:
				// log
			}
		}

		close(s.stop)
	}()
}

func (s *storeBackup) SetNextSenderMsgSeqNum(next int) {
	select {
	case s.messages <- &persisteMessage{operation: operationSetNextSenderMsgSeqNum, seqNum: next}:
	default:
		// log
	}
}

func (s *storeBackup) SetNextTargetMsgSeqNum(next int) {
	select {
	case s.messages <- &persisteMessage{operation: operationSetNextTargetMsgSeqNum, seqNum: next}:
	default:
		// log
	}
}

func (s *storeBackup) SaveMessage(seqNum int, msg []byte) {
	select {
	case s.messages <- &persisteMessage{operation: operationSaveMessage, seqNum: seqNum, msg: msg}:
	default:
		// log
	}
}

func (s *storeBackup) Reset() {
	select {
	case s.messages <- &persisteMessage{operation: operationReset}:
	default:
		// log
	}
}

func (s *storeBackup) Refresh() {
	select {
	case s.messages <- &persisteMessage{operation: operationRefresh}:
	default:
		// log
	}
}

func (s *storeBackup) Close() {
	close(s.messages)

	ticker := time.NewTicker(time.Second * 20)

	select {
	case <-s.stop:
	case <-ticker.C:
		// log
	}
}
