package quickfix

import (
	"time"

	"git.5th.im/lb-public/gear/log"
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

type backupStore struct {
	stop     chan struct{}
	messages chan *persisteMessage
	store    MessageStore
}

func newBackupStore(store MessageStore) *backupStore {
	return &backupStore{stop: make(chan struct{}), messages: make(chan *persisteMessage, 500), store: store}
}

func (s *backupStore) start() {
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
				log.Errorf("backup store: unsupported operation(%v)\n", message.operation)
			}
		}

		close(s.stop)
	}()
}

func (s *backupStore) SetNextSenderMsgSeqNum(next int) {
	select {
	case s.messages <- &persisteMessage{operation: operationSetNextSenderMsgSeqNum, seqNum: next}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SetNextSenderMsgSeqNum operation")
	}
}

func (s *backupStore) SetNextTargetMsgSeqNum(next int) {
	select {
	case s.messages <- &persisteMessage{operation: operationSetNextTargetMsgSeqNum, seqNum: next}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SetNextTargetMsgSeqNum operation")
	}
}

func (s *backupStore) SaveMessage(seqNum int, msg []byte) {
	select {
	case s.messages <- &persisteMessage{operation: operationSaveMessage, seqNum: seqNum, msg: msg}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SaveMessage operation")
	}
}

func (s *backupStore) Reset() {
	select {
	case s.messages <- &persisteMessage{operation: operationReset}:
	default:
		log.Warn("encountering a large amount of traffic, drop the Reset operation")
	}
}

func (s *backupStore) Refresh() {
	select {
	case s.messages <- &persisteMessage{operation: operationRefresh}:
	default:
		log.Warn("encountering a large amount of traffic, drop the Refresh operation")
	}
}

func (s *backupStore) Close() {
	close(s.messages)

	ticker := time.NewTicker(time.Second * 10)

	select {
	case <-s.stop:
	case <-ticker.C:
		// log
	}
}
