package quickfix

import (
	"reflect"

	"git.5th.im/lb-public/gear/log"
)

const (
	operationSetNextSenderMsgSeqNum int = iota + 1
	operationSetNextTargetMsgSeqNum
	operationSaveMessage
	operationReset
)

type persisteMessage struct {
	operation int
	seqNum    int
	msg       []byte
}

type backupStore struct {
	messages chan *persisteMessage
	store    MessageStore
}

func newBackupStore(store MessageStore) *backupStore {
	if store == nil {
		return nil
	}

	if reflect.ValueOf(store).IsNil() {
		return nil
	}

	backup := &backupStore{messages: make(chan *persisteMessage, 500), store: store}

	backup.start()

	return backup
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
			default:
				log.Errorf("backup store: unsupported operation(%v)\n", message.operation)
			}
		}
	}()
}

func (s *backupStore) SetNextSenderMsgSeqNum(next int) {
	if s == nil {
		return
	}

	select {
	case s.messages <- &persisteMessage{operation: operationSetNextSenderMsgSeqNum, seqNum: next}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SetNextSenderMsgSeqNum operation")
	}
}

func (s *backupStore) SetNextTargetMsgSeqNum(next int) {
	if s == nil {
		return
	}

	select {
	case s.messages <- &persisteMessage{operation: operationSetNextTargetMsgSeqNum, seqNum: next}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SetNextTargetMsgSeqNum operation")
	}
}

func (s *backupStore) SaveMessage(seqNum int, msg []byte) {
	if s == nil {
		return
	}

	select {
	case s.messages <- &persisteMessage{operation: operationSaveMessage, seqNum: seqNum, msg: msg}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SaveMessage operation")
	}
}

func (s *backupStore) Reset() {
	if s == nil {
		return
	}

	select {
	case s.messages <- &persisteMessage{operation: operationReset}:
	default:
		log.Warn("encountering a large amount of traffic, drop the Reset operation")
	}
}
