// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/config"
)

type mongoLogFactory struct {
	settings              *quickfix.Settings
	messagesLogCollection string
	eventLogCollection    string
}

type mongoLog struct {
	sessionID             quickfix.SessionID
	mongoURL              string
	mongoDatabase         string
	db                    *mongo.Client
	messagesLogCollection string
	eventLogCollection    string
	allowTransactions     bool
}

// NewLogFactory returns a mongo-based implementation of LogFactory.
func NewLogFactory(settings *quickfix.Settings) quickfix.LogFactory {
	return NewLogFactoryPrefixed(settings, "")
}

// NewLogFactoryPrefixed returns a mongo-based implementation of LogFactory, with prefix on collections.
func NewLogFactoryPrefixed(settings *quickfix.Settings, collectionsPrefix string) quickfix.LogFactory {
	return mongoLogFactory{
		settings:              settings,
		messagesLogCollection: collectionsPrefix + "messages_log",
		eventLogCollection:    collectionsPrefix + "event_log",
	}
}

// Create creates a new mongo implementation of the Log interface.
func (f mongoLogFactory) Create() (l quickfix.Log, err error) {
	globalSettings := f.settings.GlobalSettings()

	mongoConnectionURL, err := globalSettings.Setting(config.MongoLogConnection)
	if err != nil {
		return nil, err
	}
	mongoDatabase, err := globalSettings.Setting(config.MongoLogDatabase)
	if err != nil {
		return nil, err
	}

	// Optional.
	mongoReplicaSet, _ := globalSettings.Setting(config.MongoLogReplicaSet)

	return newmongoLog(quickfix.SessionID{}, mongoConnectionURL, mongoDatabase, mongoReplicaSet, f.messagesLogCollection, f.eventLogCollection)
}

// CreateSessionLog creates a new mongo implementation of the Log interface.
func (f mongoLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (l quickfix.Log, err error) {
	globalSettings := f.settings.GlobalSettings()
	dynamicSessions, _ := globalSettings.BoolSetting(config.DynamicSessions)

	sessionSettings, ok := f.settings.SessionSettings()[sessionID]
	if !ok {
		if dynamicSessions {
			sessionSettings = globalSettings
		} else {
			return nil, fmt.Errorf("unknown session: %v", sessionID)
		}
	}
	mongoConnectionURL, err := sessionSettings.Setting(config.MongoLogConnection)
	if err != nil {
		return nil, err
	}
	mongoDatabase, err := sessionSettings.Setting(config.MongoLogDatabase)
	if err != nil {
		return nil, err
	}

	// Optional.
	mongoReplicaSet, _ := sessionSettings.Setting(config.MongoLogReplicaSet)

	return newmongoLog(sessionID, mongoConnectionURL, mongoDatabase, mongoReplicaSet, f.messagesLogCollection, f.eventLogCollection)
}

func newmongoLog(sessionID quickfix.SessionID, mongoURL, mongoDatabase, mongoReplicaSet, messagesLogCollection, eventLogCollection string) (l *mongoLog, err error) {

	allowTransactions := len(mongoReplicaSet) > 0
	l = &mongoLog{
		sessionID:             sessionID,
		mongoURL:              mongoURL,
		mongoDatabase:         mongoDatabase,
		messagesLogCollection: messagesLogCollection,
		eventLogCollection:    eventLogCollection,
		allowTransactions:     allowTransactions,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l.db, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURL).SetDirect(len(mongoReplicaSet) == 0).SetReplicaSet(mongoReplicaSet))
	if err != nil {
		return
	}

	return
}

func (l mongoLog) OnIncoming(msg []byte) {
	l.insert(l.messagesLogCollection, msg)
}

func (l mongoLog) OnOutgoing(msg []byte) {
	l.insert(l.messagesLogCollection, msg)
}

func (l mongoLog) OnEvent(msg string) {
	l.insert(l.eventLogCollection, []byte(msg))
}

func (l mongoLog) OnEventf(format string, v ...interface{}) {
	l.insert(l.eventLogCollection, []byte(fmt.Sprintf(format, v...)))
}

func generateEntry(s *quickfix.SessionID) (entry *entryData) {
	entry = &entryData{
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

type entryData struct {
	Time             time.Time `bson:"time,omitempty"`
	BeginString      string    `bson:"begin_string"`
	SenderCompID     string    `bson:"sender_comp_id"`
	SenderSubID      string    `bson:"sender_sub_id"`
	SenderLocID      string    `bson:"sender_loc_id"`
	TargetCompID     string    `bson:"target_comp_id"`
	TargetSubID      string    `bson:"target_sub_id"`
	TargetLocID      string    `bson:"target_loc_id"`
	SessionQualifier string    `bson:"session_qualifier"`
	Text             []byte    `bson:"text,omitempty"`
}

func (l *mongoLog) insert(collection string, text []byte) {
	entry := generateEntry(&l.sessionID)
	entry.Text = text
	entry.Time = time.Now()
	_, err := l.db.Database(l.mongoDatabase).Collection(collection).InsertOne(context.Background(), entry)
	if err != nil {
		log.Println(err)
	}
}

func (l *mongoLog) iterate(coll string, cb func(string) error) error {
	entry := generateEntry(&l.sessionID)

	cursor, err := l.db.Database(l.mongoDatabase).Collection(coll).Find(context.Background(), bson.D{}, nil)
	if err != nil {
		return err
	}
	defer func() { _ = cursor.Close(context.Background()) }()
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&entry); err != nil {
			return err
		} else if err = cb(string(entry.Text)); err != nil {
			return err
		}
	}
	return nil
}

func (l *mongoLog) getEntries(coll string) ([]string, error) {
	var txts []string
	err := l.iterate(coll, func(text string) error {
		txts = append(txts, text)
		return nil
	})
	return txts, err
}

// close closes the l's database connection.
func (l *mongoLog) close() error {
	if l.db != nil {
		err := l.db.Disconnect(context.Background())
		if err != nil {
			return errors.Wrap(err, "error disconnecting from database")
		}
		l.db = nil
	}
	return nil
}
