package quickfix

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"github.com/quickfixgo/quickfix/config"
)

type mongoStoreFactory struct {
	settings           *Settings
	messagesCollection string
	sessionsCollection string
}

type mongoStore struct {
	sessionID          SessionID
	cache              *memoryStore
	mongoURL           string
	mongoDatabase      string
	db                 *mgo.Session
	messagesCollection string
	sessionsCollection string
}

// NewMongoStoreFactory returns a mongo-based implementation of MessageStoreFactory
func NewMongoStoreFactory(settings *Settings) MessageStoreFactory {
	return NewMongoStoreFactoryPrefixed(settings, "")
}

// NewMongoStoreFactoryPrefixed returns a mongo-based implementation of MessageStoreFactory, with prefix on collections
func NewMongoStoreFactoryPrefixed(settings *Settings, collectionsPrefix string) MessageStoreFactory {
	return mongoStoreFactory{
		settings:           settings,
		messagesCollection: collectionsPrefix + "messages",
		sessionsCollection: collectionsPrefix + "sessions",
	}
}

// Create creates a new MongoStore implementation of the MessageStore interface
func (f mongoStoreFactory) Create(sessionID SessionID) (msgStore MessageStore, err error) {
	sessionSettings, ok := f.settings.SessionSettings()[sessionID]
	if !ok {
		return nil, fmt.Errorf("unknown session: %v", sessionID)
	}
	mongoConnectionURL, err := sessionSettings.Setting(config.MongoStoreConnection)
	if err != nil {
		return nil, err
	}
	mongoDatabase, err := sessionSettings.Setting(config.MongoStoreDatabase)
	if err != nil {
		return nil, err
	}
	return newMongoStore(sessionID, mongoConnectionURL, mongoDatabase, f.messagesCollection, f.sessionsCollection)
}

func newMongoStore(sessionID SessionID, mongoURL string, mongoDatabase string, messagesCollection string, sessionsCollection string) (store *mongoStore, err error) {
	store = &mongoStore{
		sessionID:          sessionID,
		cache:              &memoryStore{},
		mongoURL:           mongoURL,
		mongoDatabase:      mongoDatabase,
		messagesCollection: messagesCollection,
		sessionsCollection: sessionsCollection,
	}

	if err = store.cache.Reset(); err != nil {
		err = errors.Wrap(err, "cache reset")
		return
	}

	if store.db, err = mgo.Dial(mongoURL); err != nil {
		return
	}
	err = store.populateCache()

	return
}

func generateMessageFilter(s *SessionID) (messageFilter *mongoQuickFixEntryData) {
	messageFilter = &mongoQuickFixEntryData{
		BeginString:      s.BeginString,
		SessionQualifier: s.Qualifier,
		SenderCompID:     s.SenderCompID,
		SenderSubID:      s.SenderSubID,
		SenderLocID:      s.SenderLocationID,
		TargetCompID:     s.TargetCompID,
		TargetSubID:      s.TargetSubID,
		TargetLocID:      s.TargetLocationID,
	}
	return
}

type mongoQuickFixEntryData struct {
	//Message specific data
	Msgseq  int    `bson:"msgseq,omitempty"`
	Message []byte `bson:"message,omitempty"`
	//Session specific data
	CreationTime   time.Time `bson:"creation_time,omitempty"`
	IncomingSeqNum int       `bson:"incoming_seq_num,omitempty"`
	OutgoingSeqNum int       `bson:"outgoing_seq_num,omitempty"`
	//Indexed data
	BeginString      string `bson:"begin_string"`
	SessionQualifier string `bson:"session_qualifier"`
	SenderCompID     string `bson:"sender_comp_id"`
	SenderSubID      string `bson:"sender_sub_id"`
	SenderLocID      string `bson:"sender_loc_id"`
	TargetCompID     string `bson:"target_comp_id"`
	TargetSubID      string `bson:"target_sub_id"`
	TargetLocID      string `bson:"target_loc_id"`
}

// Reset deletes the store records and sets the seqnums back to 1
func (store *mongoStore) Reset() error {
	msgFilter := generateMessageFilter(&store.sessionID)
	_, err := store.db.DB(store.mongoDatabase).C(store.messagesCollection).RemoveAll(msgFilter)

	if err != nil {
		return err
	}

	if err = store.cache.Reset(); err != nil {
		return err
	}

	sessionUpdate := generateMessageFilter(&store.sessionID)
	sessionUpdate.CreationTime = store.cache.CreationTime()
	sessionUpdate.IncomingSeqNum = store.cache.NextTargetMsgSeqNum()
	sessionUpdate.OutgoingSeqNum = store.cache.NextSenderMsgSeqNum()
	err = store.db.DB(store.mongoDatabase).C(store.sessionsCollection).Update(msgFilter, sessionUpdate)

	return err
}

// Refresh reloads the store from the database
func (store *mongoStore) Refresh() error {
	if err := store.cache.Reset(); err != nil {
		return err
	}
	return store.populateCache()
}

func (store *mongoStore) populateCache() error {
	msgFilter := generateMessageFilter(&store.sessionID)
	query := store.db.DB(store.mongoDatabase).C(store.sessionsCollection).Find(msgFilter)

	cnt, err := query.Count()
	if err != nil {
		return errors.Wrap(err, "count")
	}

	if cnt > 0 {
		// session record found, load it
		sessionData := &mongoQuickFixEntryData{}
		if err = query.One(&sessionData); err != nil {
			return errors.Wrap(err, "query one")
		}

		store.cache.creationTime = sessionData.CreationTime
		if err = store.cache.SetNextTargetMsgSeqNum(sessionData.IncomingSeqNum); err != nil {
			return errors.Wrap(err, "cache set next target")
		}

		if err = store.cache.SetNextSenderMsgSeqNum(sessionData.OutgoingSeqNum); err != nil {
			return errors.Wrap(err, "cache set next sender")
		}

		return nil
	}

	// session record not found, create it
	msgFilter.CreationTime = store.cache.creationTime
	msgFilter.IncomingSeqNum = store.cache.NextTargetMsgSeqNum()
	msgFilter.OutgoingSeqNum = store.cache.NextSenderMsgSeqNum()

	if err = store.db.DB(store.mongoDatabase).C(store.sessionsCollection).Insert(msgFilter); err != nil {
		return errors.Wrap(err, "insert")
	}
	return nil
}

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent
func (store *mongoStore) NextSenderMsgSeqNum() int {
	return store.cache.NextSenderMsgSeqNum()
}

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received
func (store *mongoStore) NextTargetMsgSeqNum() int {
	return store.cache.NextTargetMsgSeqNum()
}

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent
func (store *mongoStore) SetNextSenderMsgSeqNum(next int) error {
	msgFilter := generateMessageFilter(&store.sessionID)
	sessionUpdate := generateMessageFilter(&store.sessionID)
	sessionUpdate.IncomingSeqNum = store.cache.NextTargetMsgSeqNum()
	sessionUpdate.OutgoingSeqNum = next
	sessionUpdate.CreationTime = store.cache.CreationTime()
	if err := store.db.DB(store.mongoDatabase).C(store.sessionsCollection).Update(msgFilter, sessionUpdate); err != nil {
		return err
	}
	return store.cache.SetNextSenderMsgSeqNum(next)
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received
func (store *mongoStore) SetNextTargetMsgSeqNum(next int) error {
	msgFilter := generateMessageFilter(&store.sessionID)
	sessionUpdate := generateMessageFilter(&store.sessionID)
	sessionUpdate.IncomingSeqNum = next
	sessionUpdate.OutgoingSeqNum = store.cache.NextSenderMsgSeqNum()
	sessionUpdate.CreationTime = store.cache.CreationTime()
	if err := store.db.DB(store.mongoDatabase).C(store.sessionsCollection).Update(msgFilter, sessionUpdate); err != nil {
		return err
	}
	return store.cache.SetNextTargetMsgSeqNum(next)
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent
func (store *mongoStore) IncrNextSenderMsgSeqNum() error {
	if err := store.cache.IncrNextSenderMsgSeqNum(); err != nil {
		return errors.Wrap(err, "cache incr")
	}
	return store.SetNextSenderMsgSeqNum(store.cache.NextSenderMsgSeqNum())
}

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (store *mongoStore) IncrNextTargetMsgSeqNum() error {
	if err := store.cache.IncrNextTargetMsgSeqNum(); err != nil {
		return errors.Wrap(err, "cache incr")
	}
	return store.SetNextTargetMsgSeqNum(store.cache.NextTargetMsgSeqNum())
}

// CreationTime returns the creation time of the store
func (store *mongoStore) CreationTime() time.Time {
	return store.cache.CreationTime()
}

func (store *mongoStore) SaveMessage(seqNum int, msg []byte) (err error) {
	msgFilter := generateMessageFilter(&store.sessionID)
	msgFilter.Msgseq = seqNum
	msgFilter.Message = msg
	err = store.db.DB(store.mongoDatabase).C(store.messagesCollection).Insert(msgFilter)
	return
}

func (store *mongoStore) GetMessages(beginSeqNum, endSeqNum int) (msgs [][]byte, err error) {
	msgFilter := generateMessageFilter(&store.sessionID)
	//Marshal into database form
	msgFilterBytes, err := bson.Marshal(msgFilter)
	if err != nil {
		return
	}
	seqFilter := bson.M{}
	err = bson.Unmarshal(msgFilterBytes, &seqFilter)
	if err != nil {
		return
	}
	//Modify the query to use a range for the sequence filter
	seqFilter["msgseq"] = bson.M{
		"$gte": beginSeqNum,
		"$lte": endSeqNum,
	}

	iter := store.db.DB(store.mongoDatabase).C(store.messagesCollection).Find(seqFilter).Sort("msgseq").Iter()
	for iter.Next(msgFilter) {
		msgs = append(msgs, msgFilter.Message)
	}
	err = iter.Close()
	return
}

// Close closes the store's database connection
func (store *mongoStore) Close() error {
	if store.db != nil {
		store.db.Close()
		store.db = nil
	}
	return nil
}
