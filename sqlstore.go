package quickfix

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/omni3x/quickfix/config"
)

var isPostgres bool

type sqlStoreFactory struct {
	settings                     *Settings
	targetSeqNumDebounceInterval time.Duration
	db                           *sql.DB
}

type sqlStore struct {
	sessionID                    SessionID
	cache                        *memoryStore
	sqlDriver                    string
	sqlDataSourceName            string
	sqlConnMaxLifetime           time.Duration
	db                           *sql.DB
	targetSeqNumDebounceInterval time.Duration
	timeOfLastTargetSeqNumUpdate time.Time
}

// NewSQLStoreFactory returns a sql-based implementation of MessageStoreFactory. targetSeqNumDebounceInterval controls
// how often the target's seqnum/incoming_seqnum is updated in the DB
func NewSQLStoreFactory(settings *Settings, db *sql.DB, targetSeqNumDebounceInterval time.Duration) MessageStoreFactory {
	return sqlStoreFactory{settings: settings, db: db, targetSeqNumDebounceInterval: targetSeqNumDebounceInterval}
}

// Create creates a new SQLStore implementation of the MessageStore interface
func (f sqlStoreFactory) Create(sessionID SessionID) (msgStore MessageStore, err error) {
	sessionSettings, ok := f.settings.SessionSettings()[sessionID]
	if !ok {
		return nil, fmt.Errorf("unknown session: %v", sessionID)
	}

	sqlConnMaxLifetime := 0 * time.Second
	if f.db != nil {
		return newSQLStore(sessionID, "", "", sqlConnMaxLifetime, f.db, f.targetSeqNumDebounceInterval)
	}

	sqlDriver, err := sessionSettings.Setting(config.SQLStoreDriver)
	if err != nil {
		return nil, err
	}
	sqlDataSourceName, err := sessionSettings.Setting(config.SQLStoreDataSourceName)
	if err != nil {
		return nil, err
	}
	if sessionSettings.HasSetting(config.SQLStoreConnMaxLifetime) {
		sqlConnMaxLifetime, err = sessionSettings.DurationSetting(config.SQLStoreConnMaxLifetime)
		if err != nil {
			return nil, err
		}
	}
	return newSQLStore(sessionID, sqlDriver, sqlDataSourceName, sqlConnMaxLifetime, nil, f.targetSeqNumDebounceInterval)
}

func newSQLStore(sessionID SessionID, driver string, dataSourceName string, connMaxLifetime time.Duration, db *sql.DB, targetSeqNumDebounceInterval time.Duration) (store *sqlStore, err error) {
	isPostgres = driver == "postgres" || db != nil
	store = &sqlStore{
		sessionID:                    sessionID,
		cache:                        &memoryStore{},
		sqlDriver:                    driver,
		sqlDataSourceName:            dataSourceName,
		sqlConnMaxLifetime:           connMaxLifetime,
		targetSeqNumDebounceInterval: targetSeqNumDebounceInterval,
		timeOfLastTargetSeqNumUpdate: time.Now(),
	}
	store.cache.Reset()

	if db != nil {
		store.db = db
	} else {
		if store.db, err = sql.Open(store.sqlDriver, store.sqlDataSourceName); err != nil {
			return nil, err
		}
		store.db.SetConnMaxLifetime(store.sqlConnMaxLifetime)
	}

	if err = store.db.Ping(); err != nil { // ensure immediate connection
		return nil, err
	}
	if err = store.populateCache(); err != nil {
		return nil, err
	}

	return store, nil
}

// Reset deletes the store records and sets the seqnums back to 1
func (store *sqlStore) Reset() error {
	s := store.sessionID
	// TODO: we don't use the messages table
	// _, err := store.db.Exec(`DELETE FROM messages
	// 	WHERE beginstring=? AND session_qualifier=?
	// 	AND sendercompid=? AND sendersubid=? AND senderlocid=?
	// 	AND targetcompid=? AND targetsubid=? AND targetlocid=?`,
	// 	s.BeginString, s.Qualifier,
	// 	s.SenderCompID, s.SenderSubID, s.SenderLocationID,
	// 	s.TargetCompID, s.TargetSubID, s.TargetLocationID)
	// if err != nil {
	// 	return err
	// }

	var err error
	if err = store.cache.Reset(); err != nil {
		return err
	}

	queryStr := `UPDATE sessions
		SET creation_time=$1, incoming_seqnum=$2, outgoing_seqnum=$3
		WHERE beginstring=$4 AND session_qualifier=$5
		AND sendercompid=$6 AND sendersubid=$7 AND senderlocid=$8
		AND targetcompid=$9 AND targetsubid=$10 AND targetlocid=$11`
	if !isPostgres {
		queryStr = `UPDATE sessions
		SET creation_time=?, incoming_seqnum=?, outgoing_seqnum=?
		WHERE beginstring=? AND session_qualifier=?
		AND sendercompid=? AND sendersubid=? AND senderlocid=?
		AND targetcompid=? AND targetsubid=? AND targetlocid=?`
	}

	_, err = store.db.Exec(queryStr,
		store.cache.CreationTime(), store.cache.NextTargetMsgSeqNum(), store.cache.NextSenderMsgSeqNum(),
		s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID)

	return err
}

// Refresh reloads the store from the database
func (store *sqlStore) Refresh() error {
	if err := store.cache.Reset(); err != nil {
		return err
	}
	return store.populateCache()
}

func (store *sqlStore) populateCache() (err error) {
	s := store.sessionID
	var creationTime time.Time
	var incomingSeqNum, outgoingSeqNum int
	queryStr := `SELECT creation_time, incoming_seqnum, outgoing_seqnum
	  FROM sessions
		WHERE beginstring=$1 AND session_qualifier=$2
		AND sendercompid=$3 AND sendersubid=$4 AND senderlocid=$5
		AND targetcompid=$6 AND targetsubid=$7 AND targetlocid=$8`
	if !isPostgres {
		queryStr = `SELECT creation_time, incoming_seqnum, outgoing_seqnum
	  FROM sessions
		WHERE beginstring=? AND session_qualifier=?
		AND sendercompid=? AND sendersubid=? AND senderlocid=?
		AND targetcompid=? AND targetsubid=? AND targetlocid=?`
	}
	row := store.db.QueryRow(queryStr,
		s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID)

	err = row.Scan(&creationTime, &incomingSeqNum, &outgoingSeqNum)

	// session record found, load it
	if err == nil {
		store.cache.creationTime = creationTime
		store.cache.SetNextTargetMsgSeqNum(incomingSeqNum)
		store.cache.SetNextSenderMsgSeqNum(outgoingSeqNum)
		return nil
	}

	// fatal error, give up
	if err != sql.ErrNoRows {
		return err
	}

	// session record not found, create it
	queryStr = `INSERT INTO sessions (
				creation_time, incoming_seqnum, outgoing_seqnum,
				beginstring, session_qualifier,
				sendercompid, sendersubid, senderlocid,
				targetcompid, targetsubid, targetlocid)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	if !isPostgres {
		queryStr = `INSERT INTO sessions (
				creation_time, incoming_seqnum, outgoing_seqnum,
				beginstring, session_qualifier,
				sendercompid, sendersubid, senderlocid,
				targetcompid, targetsubid, targetlocid)
				VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	}
	_, err = store.db.Exec(queryStr,
		store.cache.creationTime,
		store.cache.NextTargetMsgSeqNum(),
		store.cache.NextSenderMsgSeqNum(),
		s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID)

	return err
}

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent
func (store *sqlStore) NextSenderMsgSeqNum() int {
	return store.cache.NextSenderMsgSeqNum()
}

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received
func (store *sqlStore) NextTargetMsgSeqNum() int {
	return store.cache.NextTargetMsgSeqNum()
}

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent
func (store *sqlStore) SetNextSenderMsgSeqNum(next int) error {
	s := store.sessionID

	queryStr := `UPDATE sessions SET outgoing_seqnum = $1
		WHERE beginstring=$2 AND session_qualifier=$3
		AND sendercompid=$4 AND sendersubid=$5 AND senderlocid=$6
		AND targetcompid=$7 AND targetsubid=$8 AND targetlocid=$9`

	if !isPostgres {
		queryStr = `UPDATE sessions SET outgoing_seqnum = ?
		WHERE beginstring=? AND session_qualifier=?
		AND sendercompid=? AND sendersubid=? AND senderlocid=?
		AND targetcompid=? AND targetsubid=? AND targetlocid=?`
	}
	_, err := store.db.Exec(queryStr,
		next, s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID)

	if err != nil {
		return err
	}
	return store.cache.SetNextSenderMsgSeqNum(next)
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received
func (store *sqlStore) SetNextTargetMsgSeqNum(next int) error {
	s := store.sessionID
	if time.Since(store.timeOfLastTargetSeqNumUpdate) > store.targetSeqNumDebounceInterval {
		queryStr := `UPDATE sessions SET incoming_seqnum = $1
		WHERE beginstring=$2 AND session_qualifier=$3
		AND sendercompid=$4 AND sendersubid=$5 AND senderlocid=$6
		AND targetcompid=$7 AND targetsubid=$8 AND targetlocid=$9`
		if !isPostgres {
			queryStr = `UPDATE sessions SET incoming_seqnum = ?
		WHERE beginstring=? AND session_qualifier=?
		AND sendercompid=? AND sendersubid=? AND senderlocid=?
		AND targetcompid=? AND targetsubid=? AND targetlocid=?`
		}
		_, err := store.db.Exec(queryStr,
			next, s.BeginString, s.Qualifier,
			s.SenderCompID, s.SenderSubID, s.SenderLocationID,
			s.TargetCompID, s.TargetSubID, s.TargetLocationID)

		if err != nil {
			return err
		}
		store.timeOfLastTargetSeqNumUpdate = time.Now()
	}
	return store.cache.SetNextTargetMsgSeqNum(next)
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent
func (store *sqlStore) IncrNextSenderMsgSeqNum() error {
	store.cache.IncrNextSenderMsgSeqNum()
	return store.SetNextSenderMsgSeqNum(store.cache.NextSenderMsgSeqNum())
}

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received
func (store *sqlStore) IncrNextTargetMsgSeqNum() error {
	store.cache.IncrNextTargetMsgSeqNum()
	return store.SetNextTargetMsgSeqNum(store.cache.NextTargetMsgSeqNum())
}

// CreationTime returns the creation time of the store
func (store *sqlStore) CreationTime() time.Time {
	return store.cache.CreationTime()
}

func (store *sqlStore) SaveMessage(seqNum int, msg []byte) error {
	s := store.sessionID

	_, err := store.db.Exec(`INSERT INTO messages (
			msgseqnum, message,
			beginstring, session_qualifier,
			sendercompid, sendersubid, senderlocid,
			targetcompid, targetsubid, targetlocid)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		seqNum, string(msg),
		s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID)

	return err
}

func (store *sqlStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	s := store.sessionID
	var msgs [][]byte
	rows, err := store.db.Query(`SELECT message FROM messages
		WHERE beginstring=? AND session_qualifier=?
		AND sendercompid=? AND sendersubid=? AND senderlocid=?
		AND targetcompid=? AND targetsubid=? AND targetlocid=?
		AND msgseqnum>=? AND msgseqnum<=?
		ORDER BY msgseqnum`,
		s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID,
		beginSeqNum, endSeqNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			return nil, err
		}
		msgs = append(msgs, []byte(message))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}

// Close closes the store's database connection
func (store *sqlStore) Close() error {
	if store.db != nil {
		store.db.Close()
		store.db = nil
	}
	return nil
}
