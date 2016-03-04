package quickfix

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/quickfixgo/quickfix/config"
)

type seqNumFile struct {
	fname string
	fd    *os.File
}

type fileStore struct {
	cache            MessageStore
	sessionID        SessionID
	senderSeqNumFile *seqNumFile
	targetSeqNumFile *seqNumFile
}

type fileStoreFactory struct {
	settings *SessionSettings
}

func newFileStore(sessionID SessionID, dirname string) (store *fileStore, err error) {
	cache, err := NewMemoryStoreFactory().Create(sessionID)
	if err != nil {
		return nil, err
	}

	sessionPrefix := sessionIDFilenamePrefix(sessionID)
	store = &fileStore{
		cache:            cache,
		sessionID:        sessionID,
		senderSeqNumFile: newSeqNumFile(dirname, sessionPrefix, ".senderseqnums"),
		targetSeqNumFile: newSeqNumFile(dirname, sessionPrefix, ".targetseqnums"),
	}

	if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
		return nil, err
	}

	if err := store.Refresh(); err != nil {
		return nil, err
	}

	return store, nil
}

func (f fileStoreFactory) Create(sessionID SessionID) (msgStore MessageStore, err error) {
	dirname, err := f.settings.Setting(config.FileStorePath)
	if err != nil {
		return nil, err
	}
	return newFileStore(sessionID, dirname)
}

// NewFileStoreFactory returns a file-based implementation of MessageStoreFactory
func NewFileStoreFactory(settings *SessionSettings) MessageStoreFactory {
	return fileStoreFactory{settings: settings}
}

// NextSenderMsgSeqNum returns the next MsgSeqNum that will sent
func (store *fileStore) NextSenderMsgSeqNum() int {
	return store.cache.NextSenderMsgSeqNum()
}

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received
func (store *fileStore) NextTargetMsgSeqNum() int {
	return store.cache.NextTargetMsgSeqNum()
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent
func (store *fileStore) IncrNextSenderMsgSeqNum() error {
	store.cache.IncrNextSenderMsgSeqNum()
	return store.senderSeqNumFile.Write(store.cache.NextSenderMsgSeqNum())
}

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (store *fileStore) IncrNextTargetMsgSeqNum() error {
	store.cache.IncrNextTargetMsgSeqNum()
	return store.targetSeqNumFile.Write(store.cache.NextTargetMsgSeqNum())
}

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent
func (store *fileStore) SetNextSenderMsgSeqNum(next int) error {
	store.cache.SetNextSenderMsgSeqNum(next)
	return store.senderSeqNumFile.Write(next)
}

// SetNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (store *fileStore) SetNextTargetMsgSeqNum(next int) error {
	store.cache.SetNextTargetMsgSeqNum(next)
	return store.targetSeqNumFile.Write(next)
}

// CreationTime returns the creation time of the store
func (store *fileStore) CreationTime() time.Time {
	return store.cache.CreationTime()
}

func (store *fileStore) SaveMessage(seqNum int, msg []byte) {
	store.cache.SaveMessage(seqNum, msg)
}

func (store *fileStore) GetMessages(beginSeqNum, endSeqNum int) chan []byte {
	return store.cache.GetMessages(beginSeqNum, endSeqNum)
}

// Refresh closes the store files and then reloads from them
func (store *fileStore) Refresh() (err error) {
	store.Close()

	if err := store.senderSeqNumFile.Open(); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := store.targetSeqNumFile.Open(); err != nil {
		log.Println(err.Error())
		return err
	}

	store.cache.Reset()

	senderSeqNum, err := store.senderSeqNumFile.Read()
	if err != nil {
		return err
	}
	store.cache.SetNextSenderMsgSeqNum(senderSeqNum)

	targetSeqNum, err := store.targetSeqNumFile.Read()
	if err != nil {
		return err
	}
	store.cache.SetNextTargetMsgSeqNum(targetSeqNum)

	return nil
}

// Reset deletes the store files and sets the seqnums back to 1
func (store *fileStore) Reset() error {
	if err := store.Close(); err != nil {
		return err
	}
	if err := store.senderSeqNumFile.Remove(); err != nil {
		return err
	}
	if err := store.targetSeqNumFile.Remove(); err != nil {
		return err
	}
	return store.Refresh()
}

// Close closes the store's files
func (store *fileStore) Close() error {
	if err := store.cache.Close(); err != nil {
		return err
	}
	if err := store.senderSeqNumFile.Close(); err != nil {
		return err
	}
	if err := store.targetSeqNumFile.Close(); err != nil {
		return err
	}
	return nil
}

func newSeqNumFile(dirname string, basename string, extension string) *seqNumFile {
	return &seqNumFile{
		fname: path.Join(dirname, fmt.Sprintf("%s.%s", basename, extension)),
		fd:    nil,
	}
}

func (f *seqNumFile) Read() (int, error) {
	if _, err := f.fd.Seek(0, 0); err != nil {
		return 0, err
	}

	seqNumAsBytes, err := ioutil.ReadAll(f.fd)
	if err != nil {
		return 0, err
	}

	// will be empty if this file was newly created
	if len(seqNumAsBytes) < 1 {
		return 1, nil
	}
	return strconv.Atoi(string(seqNumAsBytes))
}

func (f *seqNumFile) Write(seqNum int) error {
	if err := f.fd.Truncate(0); err != nil {
		return err
	}

	if _, err := f.fd.Seek(0, 0); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(f.fd, "%d", seqNum); err != nil {
		return err
	}

	return nil
}

func (f *seqNumFile) Open() (err error) {
	f.fd, err = os.OpenFile(f.fname, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return fmt.Errorf("error opening store file: %s: %s\n", f.fname, err.Error())
	}
	return nil
}

func (f *seqNumFile) Close() error {
	if f == nil || f.fd == nil {
		return nil
	}
	if err := f.fd.Close(); err != nil {
		return err
	}
	f.fd = nil
	return nil
}

func (f *seqNumFile) Remove() error {
	if f == nil {
		return nil
	}
	if f.fd != nil {
		if err := f.fd.Close(); err != nil {
			return err
		}
	}
	if err := os.Remove(f.fname); err != nil {
		return err
	}
	return nil
}
