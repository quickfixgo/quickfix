package quickfix

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"

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
	mongoUrl           string
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
	mongoConnectionUrl, err := sessionSettings.Setting(config.MongoStoreConnection)
	if err != nil {
		return nil, err
	}
	mongoDatabase, err := sessionSettings.Setting(config.MongoStoreDatabase)
	if err != nil {
		return nil, err
	}
	return newMongoStore(sessionID, mongoConnectionUrl, mongoDatabase, f.messagesCollection, f.sessionsCollection)
}

func newMongoStore(sessionID SessionID, mongoUrl string, mongoDatabase string, messagesCollection string, sessionsCollection string) (store *mongoStore, err error) {
	store = &mongoStore{
		sessionID:          sessionID,
		cache:              &memoryStore{},
		mongoUrl:           mongoUrl,
		mongoDatabase:      mongoDatabase,
		messagesCollection: messagesCollection,
		sessionsCollection: sessionsCollection,
	}
	store.cache.Reset()

	if store.db, err = mgo.Dial(mongoUrl); err != nil {
		return
	}
	err = store.populateCache()

	return
}

func generateMessageFilter(s *SessionID) (messageFilter *mongoQuickFixEntryData) {
	messageFilter = &mongoQuickFixEntryData{
		BeginString:      s.BeginString,
		SessionQualifier: s.Qualifier,
		SenderCompId:     s.SenderCompID,
		SenderSubId:      s.SenderSubID,
		SenderLocId:      s.SenderLocationID,
		TargetCompId:     s.TargetCompID,
		TargetSubId:      s.TargetSubID,
		TargetLocId:      s.TargetLocationID,
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
	SenderCompId     string `bson:"sender_comp_id"`
	SenderSubId      string `bson:"sender_sub_id"`
	SenderLocId      string `bson:"sender_loc_id"`
	TargetCompId     string `bson:"target_comp_id"`
	TargetSubId      string `bson:"target_sub_id"`
	TargetLocId      string `bson:"target_loc_id"`
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

func (store *mongoStore) populateCache() (err error) {
	msgFilter := generateMessageFilter(&store.sessionID)
	query := store.db.DB(store.mongoDatabase).C(store.sessionsCollection).Find(msgFilter)

	if cnt, err := query.Count(); err == nil && cnt > 0 {
		// session record found, load it
		sessionData := &mongoQuickFixEntryData{}
		err = query.One(&sessionData)
		if err == nil {
			store.cache.creationTime = sessionData.CreationTime
			store.cache.SetNextTargetMsgSeqNum(sessionData.IncomingSeqNum)
			store.cache.SetNextSenderMsgSeqNum(sessionData.OutgoingSeqNum)
		}
	} else if err == nil && cnt == 0 {
		// session record not found, create it
		msgFilter.CreationTime = store.cache.creationTime
		msgFilter.IncomingSeqNum = store.cache.NextTargetMsgSeqNum()
		msgFilter.OutgoingSeqNum = store.cache.NextSenderMsgSeqNum()
		err = store.db.DB(store.mongoDatabase).C(store.sessionsCollection).Insert(msgFilter)
	}
	return
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
	store.cache.IncrNextSenderMsgSeqNum()
	return store.SetNextSenderMsgSeqNum(store.cache.NextSenderMsgSeqNum())
}

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (store *mongoStore) IncrNextTargetMsgSeqNum() error {
	store.cache.IncrNextTargetMsgSeqNum()
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
