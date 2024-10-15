package quickfix

import (
	"fmt"
)

const (
	OperationSetNextSenderMsgSeqNum int = iota + 1
	OperationSetNextTargetMsgSeqNum
	OperationSaveMessage
	OperationReset
)

type BackupMessage struct {
	Operation int
	SeqNum    int
	Msg       []byte
}

type backupStoreFactory struct {
	messagesQueue chan *BackupMessage
	backupFactory MessageStoreFactory
}

type backupStore struct {
	messagesQueue chan *BackupMessage
	store         MessageStore
}

func NewBackupStoreFactory(messagesQueue chan *BackupMessage, backupFactory MessageStoreFactory) *backupStoreFactory {
	return &backupStoreFactory{messagesQueue: messagesQueue, backupFactory: backupFactory}
}

func (f backupStoreFactory) Create(sessionID SessionID) (msgStore *backupStore, err error) {
	backupStore, err := f.backupFactory.Create(sessionID)
	if err != nil {
		return nil, err
	}

	return newBackupStore(backupStore, f.messagesQueue), nil
}

func newBackupStore(store MessageStore, messagesQueue chan *BackupMessage) *backupStore {
	backup := &backupStore{messagesQueue: messagesQueue, store: store}

	backup.start()

	return backup
}

func (s *backupStore) start() {
	if s == nil {
		return
	}

	go func() {
		for message := range s.messagesQueue {
			switch message.Operation {
			case OperationSetNextSenderMsgSeqNum:
				if err := s.store.SetNextSenderMsgSeqNum(message.SeqNum); err != nil {
				}
			case OperationSetNextTargetMsgSeqNum:
				if err := s.store.SetNextTargetMsgSeqNum(message.SeqNum); err != nil {
				}
			case OperationSaveMessage:
				if err := s.store.SaveMessage(message.SeqNum, message.Msg); err != nil {
				}
			case OperationReset:
				if err := s.store.Reset(); err != nil {
				}
			default:
				fmt.Printf("backup store: unsupported operation(%v)\n", message.Operation)
			}
		}
	}()
}

func (s *backupStore) SetNextSenderMsgSeqNum(next int) {
	if s == nil {
		return
	}

	select {
	case s.messagesQueue <- &BackupMessage{Operation: OperationSetNextSenderMsgSeqNum, SeqNum: next}:
	default:
		fmt.Println("encountering a large amount of traffic, drop the SetNextSenderMsgSeqNum operation")
	}
}

func (s *backupStore) SetNextTargetMsgSeqNum(next int) {
	if s == nil {
		return
	}

	select {
	case s.messagesQueue <- &BackupMessage{Operation: OperationSetNextTargetMsgSeqNum, SeqNum: next}:
	default:
		fmt.Println("encountering a large amount of traffic, drop the SetNextTargetMsgSeqNum operation")
	}
}

func (s *backupStore) SaveMessage(seqNum int, msg []byte) {
	if s == nil {
		return
	}

	select {
	case s.messagesQueue <- &BackupMessage{Operation: OperationSaveMessage, SeqNum: seqNum, Msg: msg}:
	default:
		fmt.Println("encountering a large amount of traffic, drop the SaveMessage operation")
	}
}

func (s *backupStore) Reset() {
	if s == nil {
		return
	}

	select {
	case s.messagesQueue <- &BackupMessage{Operation: OperationReset}:
	default:
		fmt.Println("encountering a large amount of traffic, drop the Reset operation")
	}
}
