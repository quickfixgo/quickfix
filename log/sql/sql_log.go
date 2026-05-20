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

package sql

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/config"
)

type sqlLogFactory struct {
	settings *quickfix.Settings
}

type sqlLog struct {
	sessionID          quickfix.SessionID
	sqlDriver          string
	sqlDataSourceName  string
	sqlConnMaxLifetime time.Duration
	db                 *sql.DB
	placeholder        placeholderFunc
}

type placeholderFunc func(int) string

var rePlaceholder = regexp.MustCompile(`\?`)

func sqlString(raw string, placeholder placeholderFunc) string {
	if placeholder == nil {
		return raw
	}
	idx := 0
	return rePlaceholder.ReplaceAllStringFunc(raw, func(_ string) string {
		p := placeholder(idx)
		idx++
		return p
	})
}

func postgresPlaceholder(i int) string {
	return fmt.Sprintf("$%d", i+1)
}

// NewLogFactory returns a sql-based implementation of LogFactory.
func NewLogFactory(settings *quickfix.Settings) quickfix.LogFactory {
	return sqlLogFactory{settings: settings}
}

// Create creates a new SQLLog implementation of the Log interface.
func (f sqlLogFactory) Create() (log quickfix.Log, err error) {
	globalSettings := f.settings.GlobalSettings()

	sqlDriver, err := globalSettings.Setting(config.SQLLogDriver)
	if err != nil {
		return nil, err
	}
	sqlDataSourceName, err := globalSettings.Setting(config.SQLLogDataSourceName)
	if err != nil {
		return nil, err
	}
	sqlConnMaxLifetime := 0 * time.Second
	if globalSettings.HasSetting(config.SQLLogConnMaxLifetime) {
		sqlConnMaxLifetime, err = globalSettings.DurationSetting(config.SQLLogConnMaxLifetime)
		if err != nil {
			return nil, err
		}
	}

	return newSQLLog(quickfix.SessionID{}, sqlDriver, sqlDataSourceName, sqlConnMaxLifetime)
}

// CreateSessionLog creates a new SQLLog implementation of the Log interface.
func (f sqlLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (log quickfix.Log, err error) {
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

	sqlDriver, err := sessionSettings.Setting(config.SQLLogDriver)
	if err != nil {
		return nil, err
	}
	sqlDataSourceName, err := sessionSettings.Setting(config.SQLLogDataSourceName)
	if err != nil {
		return nil, err
	}
	sqlConnMaxLifetime := 0 * time.Second
	if sessionSettings.HasSetting(config.SQLLogConnMaxLifetime) {
		sqlConnMaxLifetime, err = sessionSettings.DurationSetting(config.SQLLogConnMaxLifetime)
		if err != nil {
			return nil, err
		}
	}
	return newSQLLog(sessionID, sqlDriver, sqlDataSourceName, sqlConnMaxLifetime)
}

func newSQLLog(sessionID quickfix.SessionID, driver string, dataSourceName string, connMaxLifetime time.Duration) (l *sqlLog, err error) {
	l = &sqlLog{
		sessionID:          sessionID,
		sqlDriver:          driver,
		sqlDataSourceName:  dataSourceName,
		sqlConnMaxLifetime: connMaxLifetime,
	}

	if l.sqlDriver == "postgres" || l.sqlDriver == "pgx" {
		l.placeholder = postgresPlaceholder
	}

	if l.db, err = sql.Open(l.sqlDriver, l.sqlDataSourceName); err != nil {
		return nil, err
	}
	l.db.SetConnMaxLifetime(l.sqlConnMaxLifetime)

	if err = l.db.Ping(); err != nil { // ensure immediate connection
		return nil, err
	}

	return l, nil
}

func (l sqlLog) OnIncoming(msg []byte) {
	l.insert("messages_log", string(msg))
}

func (l sqlLog) OnOutgoing(msg []byte) {
	l.insert("messages_log", string(msg))
}

func (l sqlLog) OnEvent(msg string) {
	l.insert("event_log", msg)
}

func (l sqlLog) OnEventf(format string, v ...interface{}) {
	l.insert("event_log", fmt.Sprintf(format, v...))
}

func (l sqlLog) insert(table string, value string) {
	s := l.sessionID

	_, err := l.db.Exec(sqlString(`INSERT INTO `+table+` (
			time,
			beginstring, session_qualifier,
			sendercompid, sendersubid, senderlocid,
			targetcompid, targetsubid, targetlocid, 
			text)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, l.placeholder),
		time.Now(),
		s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID,
		value,
	)
	if err != nil {
		log.Println(err)
	}
}

func (l *sqlLog) iterate(table string, cb func(string) error) error {
	s := l.sessionID
	rows, err := l.db.Query(sqlString(`SELECT text FROM `+table+`
		WHERE beginstring=? AND session_qualifier=?
		AND sendercompid=? AND sendersubid=? AND senderlocid=?
		AND targetcompid=? AND targetsubid=? AND targetlocid=?`,
		l.placeholder),
		s.BeginString, s.Qualifier,
		s.SenderCompID, s.SenderSubID, s.SenderLocationID,
		s.TargetCompID, s.TargetSubID, s.TargetLocationID)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var txt string
		if err = rows.Scan(&txt); err != nil {
			return err
		} else if err = cb(txt); err != nil {
			return err
		}
	}

	return rows.Err()
}

func (l *sqlLog) getEntries(table string) ([]string, error) {
	var msgs []string
	err := l.iterate(table, func(msg string) error {
		msgs = append(msgs, msg)
		return nil
	})
	return msgs, err
}

// Close closes the log's database connection.
func (l *sqlLog) close() error {
	if l.db != nil {
		l.db.Close()
		l.db = nil
	}
	return nil
}
