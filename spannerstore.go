package quickfix

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/pkg/errors"
	"github.com/quickfixgo/quickfix/config"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
)

type spannerStoreFactory struct {
	settings *Settings
}

type spannerStore struct {
	sessionID SessionID
	cache     *memoryStore
	client    *spanner.Client
}

// NewSpannerStoreFactory returns a Spanner-based implementation of MessageStoreFactory
func NewSpannerStoreFactory(settings *Settings) MessageStoreFactory {
	return spannerStoreFactory{settings: settings}
}

// Create creates a new SpannerStore implementation of the MessageStore interface
func (f spannerStoreFactory) Create(sessionID SessionID) (msgStore MessageStore, err error) {
	sessionSettings, ok := f.settings.SessionSettings()[sessionID]
	if !ok {
		return nil, fmt.Errorf("unknown session: %v", sessionID)
	}
	spannerDataSourceName, err := sessionSettings.Setting(config.SpannerDataSource)
	if err != nil {
		return nil, err
	}
	return newSpannerStore(sessionID, spannerDataSourceName)
}

func newSpannerStore(sessionID SessionID, database string) (store *spannerStore, err error) {
	store = &spannerStore{
		sessionID: sessionID,
		cache:     &memoryStore{},
	}
	if err = store.cache.Reset(); err != nil {
		err = errors.Wrap(err, "cache reset")
		return
	}
	client, err := spanner.NewClient(context.Background(), database)
	if err != nil {
		return nil, err
	}
	store.client = client
	if err = store.populateCache(); err != nil {
		return nil, err
	}
	return store, nil
}

// Reset deletes the store records and sets the seqnums back to 1
func (store *spannerStore) Reset() error {
	s := store.sessionID
	_, err := store.client.ReadWriteTransaction(context.Background(), func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `DELETE FROM messages
				WHERE beginstring=@beginstring
				AND sendercompid=@sendercompid
				AND sendersubid=@sendersubid
				AND senderlocid=@senderlocid
				AND targetcompid=@targetcompid
				AND targetsubid=@targetsubid
				AND targetlocid=@targetlocid
				AND session_qualifier=@session_qualifier`,
			Params: map[string]interface{}{
				"beginstring":       s.BeginString,
				"sendercompid":      s.SenderCompID,
				"sendersubid":       s.SenderSubID,
				"senderlocid":       s.SenderLocationID,
				"targetcompid":      s.TargetCompID,
				"targetsubid":       s.TargetSubID,
				"targetlocid":       s.TargetLocationID,
				"session_qualifier": s.Qualifier,
			},
		}
		_, err := tx.Update(ctx, stmt)
		if err != nil {
			return err
		}
		err = store.cache.Reset()
		if err != nil {
			return err
		}
		stmt = spanner.Statement{
			SQL: `UPDATE sessions SET
				creation_time=@creation_time,
				incoming_seqnum=@incoming_seqnum,
				outgoing_seqnum=@outgoing_seqnum
				WHERE beginstring=@beginstring
				AND sendercompid=@sendercompid
				AND sendersubid=@sendersubid
				AND senderlocid=@senderlocid
				AND targetcompid=@targetcompid
				AND targetsubid=@targetsubid
				AND targetlocid?@targetlocid
				AND session_qualifier=@session_qualifier`,
			Params: map[string]interface{}{
				"creation_time":     store.cache.CreationTime(),
				"incoming_seqnum":   store.cache.NextTargetMsgSeqNum(),
				"outgoing_seqnum":   store.cache.NextSenderMsgSeqNum(),
				"beginstring":       s.BeginString,
				"sendercompid":      s.SenderCompID,
				"sendersubid":       s.SenderSubID,
				"senderlocid":       s.SenderLocationID,
				"targetcompid":      s.TargetCompID,
				"targetsubid":       s.TargetSubID,
				"targetlocid":       s.TargetLocationID,
				"session_qualifier": s.Qualifier,
			},
		}
		_, err = tx.Update(ctx, stmt)
		return err
	})
	return err
}

// Refresh reloads the store from the database
func (store *spannerStore) Refresh() error {
	if err := store.cache.Reset(); err != nil {
		return err
	}
	return store.populateCache()
}

type Session struct {
	Creation         time.Time `spanner:"creation_time"`
	IncomingSequence int64     `spanner:"incoming_seqnum"`
	OutgoingSequence int64     `spanner:"outgoing_seqnum"`
}

func (store *spannerStore) populateCache() error {
	s := store.sessionID
	_, err := store.client.ReadWriteTransaction(context.Background(), func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		key := spanner.Key{s.BeginString, s.SenderCompID, s.SenderSubID, s.SenderLocationID, s.TargetCompID, s.TargetSubID, s.TargetLocationID, s.Qualifier}
		col := []string{"creation_time", "incoming_seqnum", "outgoing_seqnum"}
		row, err := tx.ReadRow(ctx, "sessions", key, col)
		if err == nil {
			var session Session
			err = row.ToStruct(&session)
			if err != nil {
				return err
			}
			store.cache.creationTime = session.Creation
			if err = store.cache.SetNextTargetMsgSeqNum(int(session.IncomingSequence)); err != nil {
				return errors.Wrap(err, "cache set next target")
			}
			if err = store.cache.SetNextSenderMsgSeqNum(int(session.OutgoingSequence)); err != nil {
				return errors.Wrap(err, "cache set next sender")
			}
			return nil
		}
		if spanner.ErrCode(err) != codes.NotFound {
			return err
		}
		stmt := spanner.Statement{
			SQL: `INSERT sessions (
					beginstring,
					sendercompid, sendersubid, senderlocid,
					targetcompid, targetsubid, targetlocid,
					session_qualifier,
					creation_time,
					incoming_seqnum, outgoing_seqnum)
				VALUES (
					@beginstring,
					@sendercompid, @sendersubid, @senderlocid,
					@targetcompid, @targetsubid, @targetlocid,
					@session_qualifier,
					@creation_time,
					@incoming_seqnum, @outgoing_seqnum)`,
			Params: map[string]interface{}{
				"beginstring":       s.BeginString,
				"sendercompid":      s.SenderCompID,
				"sendersubid":       s.SenderSubID,
				"senderlocid":       s.SenderLocationID,
				"targetcompid":      s.TargetCompID,
				"targetsubid":       s.TargetSubID,
				"targetlocid":       s.TargetLocationID,
				"session_qualifier": s.Qualifier,
				"creation_time":     store.cache.creationTime,
				"incoming_seqnum":   store.cache.NextTargetMsgSeqNum(),
				"outgoing_seqnum":   store.cache.NextSenderMsgSeqNum(),
			},
		}
		_, err = tx.Update(ctx, stmt)
		return err
	})
	return err
}

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent
func (store *spannerStore) NextSenderMsgSeqNum() int {
	return store.cache.NextSenderMsgSeqNum()
}

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received
func (store *spannerStore) NextTargetMsgSeqNum() int {
	return store.cache.NextTargetMsgSeqNum()
}

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent
func (store *spannerStore) SetNextSenderMsgSeqNum(next int) error {
	s := store.sessionID
	cols := []string{
		"beginstring",
		"sendercompid", "sendersubid", "senderlocid",
		"targetcompid", "targetsubid", "targetlocid",
		"session_qualifier",
		"outgoing_seqnum",
	}
	vals := []interface{}{
		s.BeginString,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID,
		s.Qualifier,
		next,
	}
	m := spanner.Update("sessions", cols, vals)
	_, err := store.client.Apply(context.Background(), []*spanner.Mutation{m})
	return err
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received
func (store *spannerStore) SetNextTargetMsgSeqNum(next int) error {
	s := store.sessionID
	cols := []string{
		"beginstring",
		"sendercompid", "sendersubid", "senderlocid",
		"targetcompid", "targetsubid", "targetlocid",
		"session_qualifier",
		"incoming_seqnum",
	}
	vals := []interface{}{
		s.BeginString,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID,
		s.Qualifier,
		next,
	}
	m := spanner.Update("sessions", cols, vals)
	_, err := store.client.Apply(context.Background(), []*spanner.Mutation{m})
	return err

}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent
func (store *spannerStore) IncrNextSenderMsgSeqNum() error {
	if err := store.cache.IncrNextSenderMsgSeqNum(); err != nil {
		return errors.Wrap(err, "cache incr next")
	}
	return store.SetNextSenderMsgSeqNum(store.cache.NextSenderMsgSeqNum())
}

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (store *spannerStore) IncrNextTargetMsgSeqNum() error {
	if err := store.cache.IncrNextTargetMsgSeqNum(); err != nil {
		return errors.Wrap(err, "cache incr next")
	}
	return store.SetNextTargetMsgSeqNum(store.cache.NextTargetMsgSeqNum())
}

// CreationTime returns the creation time of the store
func (store *spannerStore) CreationTime() time.Time {
	return store.cache.CreationTime()
}

func (store *spannerStore) SaveMessage(seqNum int, msg []byte) error {
	s := store.sessionID
	cols := []string{
		"beginstring",
		"sendercompid", "sendersubid", "senderlocid",
		"targetcompid", "targetsubid", "targetlocid",
		"session_qualifier", "msgseqnum",
		"message",
	}
	vals := []interface{}{
		s.BeginString,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID,
		s.Qualifier, seqNum,
		msg,
	}
	m := spanner.Insert("messages", cols, vals)
	_, err := store.client.Apply(context.Background(), []*spanner.Mutation{m})
	return err
}

func (store *spannerStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	s := store.sessionID
	tx := store.client.ReadOnlyTransaction()
	stmt := spanner.Statement{
		SQL: `SELECT message FROM messages
				WHERE beginstring=@beginstring
				AND sendercompid=@sendercompid
				AND sendersubid=@sendersubid
				AND senderlocid=@senderlocid
				AND targetcompid=@targetcompid
				AND targetsubid=@targetsubid
				AND targetlocid=@targetlocid
				AND session_qualifier=@session_qualifier
				AND msgseqnum>=@beginseq AND msgseqnum<=@endseq`,
		Params: map[string]interface{}{
			"beginstring":       s.BeginString,
			"sendercompid":      s.SenderCompID,
			"sendersubid":       s.SenderSubID,
			"senderlocid":       s.SenderLocationID,
			"targetcompid":      s.TargetCompID,
			"targetsubid":       s.TargetSubID,
			"targetlocid":       s.TargetLocationID,
			"session_qualifier": s.Qualifier,
			"beginseq":          beginSeqNum,
			"endseq":            endSeqNum,
		},
	}
	iter := tx.Query(context.Background(), stmt)
	defer iter.Stop()
	var msgs [][]byte
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return msgs, nil
		}
		if err != nil {
			return msgs, err
		}
		var msg []byte
		if err := row.Columns(&msg); err != nil {
			return msgs, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

// Close closes the store's database connection
func (store *spannerStore) Close() error {
	if store.client != nil {
		store.client.Close()
		store.client = nil
	}
	return nil
}
