package quickfix

import (
	"git.5th.im/lb-public/gear/log"
)

const (
	operationSetNextSenderMsgSeqNum int = iota + 1
	operationSetNextTargetMsgSeqNum
	operationSaveMessage
	operationReset
)

type persistentMessage struct {
	operation int
	seqNum    int
	msg       []byte
}

type backupStoreFactory struct {
	messagesQueue chan *persistentMessage
	backupFactory MessageStoreFactory
}

type backupStore struct {
	messagesQueue chan *persistentMessage
	store         MessageStore
}

func NewBackupStoreFactory(messagesQueue chan *persistentMessage, backupFactory MessageStoreFactory) *backupStoreFactory {
	return &backupStoreFactory{messagesQueue: messagesQueue, backupFactory: backupFactory}
}

func (f backupStoreFactory) Create(sessionID SessionID) (msgStore *backupStore, err error) {
	backupStore, err := f.backupFactory.Create(sessionID)
	if err != nil {
		return nil, err
	}

	return newBackupStore(backupStore, f.messagesQueue), nil
}

func newBackupStore(store MessageStore, messagesQueue chan *persistentMessage) *backupStore {
	backup := &backupStore{messagesQueue: messagesQueue, store: store}

	backup.start()

	return backup
}

func (s *backupStore) start() {
	go func() {
		for message := range s.messagesQueue {
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
	case s.messagesQueue <- &persistentMessage{operation: operationSetNextSenderMsgSeqNum, seqNum: next}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SetNextSenderMsgSeqNum operation")
	}
}

func (s *backupStore) SetNextTargetMsgSeqNum(next int) {
	if s == nil {
		return
	}

	select {
	case s.messagesQueue <- &persistentMessage{operation: operationSetNextTargetMsgSeqNum, seqNum: next}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SetNextTargetMsgSeqNum operation")
	}
}

func (s *backupStore) SaveMessage(seqNum int, msg []byte) {
	if s == nil {
		return
	}

	select {
	case s.messagesQueue <- &persistentMessage{operation: operationSaveMessage, seqNum: seqNum, msg: msg}:
	default:
		log.Warn("encountering a large amount of traffic, drop the SaveMessage operation")
	}
}

func (s *backupStore) Reset() {
	if s == nil {
		return
	}

	select {
	case s.messagesQueue <- &persistentMessage{operation: operationReset}:
	default:
		log.Warn("encountering a large amount of traffic, drop the Reset operation")
	}
}
