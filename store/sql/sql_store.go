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
	"regexp"
	"time"

	"github.com/quickfixgo/quickfix"
)

const (
	defaultMessagesTable = "messages"
	defaultSessionsTable = "sessions"
)

type sqlStoreFactory struct {
	settings *quickfix.Settings
}

type sqlStore struct {
	sessionID          quickfix.SessionID
	cache              quickfix.MessageStore
	sqlDriver          string
	sqlDataSourceName  string
	sqlConnMaxLifetime time.Duration
	db                 *sql.DB
	placeholder        placeholderFunc
	messagesTable      string
	sessionsTable      string

	sqlUpdateSeqNums      string
	sqlInsertSession      string
	sqlGetSeqNums         string
	sqlUpdateMessage      string
	sqlInsertMessage      string
	sqlGetMessages        string
	sqlUpdateSession      string
	sqlUpdateSenderSeqNum string
	sqlUpdateTargetSeqNum string
	sqlDeleteMessages     string
}

type placeholderFunc func(int) string

var rePlaceholder = regexp.MustCompile(`\?`)

func sqlString(raw string, placeholder placeholderFunc) string {
	_ = "STUB: not implemented"
	return ""
}

func postgresPlaceholder(i int) string { _ = "STUB: not implemented"; return "" }

// NewStoreFactory returns a sql-based implementation of MessageStoreFactory.
func NewStoreFactory(settings *quickfix.Settings) quickfix.MessageStoreFactory {
	_ = "STUB: not implemented"
	return *new(quickfix.MessageStoreFactory)
}

// Create creates a new SQLStore implementation of the MessageStore interface.
func (f sqlStoreFactory) Create(sessionID quickfix.SessionID) (msgStore quickfix.MessageStore, err error) {
	_ = "STUB: not implemented"
	return *new(quickfix.MessageStore), nil
}

func newSQLStore(sessionID quickfix.SessionID, driver, dataSourceName, messagesTableName, sessionsTableName string, connMaxLifetime time.Duration) (store *sqlStore, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ensure immediate connection

func (store *sqlStore) setSQLStatements() { _ = "STUB: not implemented"; return }

// Reset deletes the store records and sets the seqnums back to 1.
func (store *sqlStore) Reset() error { _ = "STUB: not implemented"; return nil }

// Refresh reloads the store from the database.
func (store *sqlStore) Refresh() error { _ = "STUB: not implemented"; return nil }

func (store *sqlStore) populateCache() error { _ = "STUB: not implemented"; return nil }

// session record found, load it

// fatal error, give up

// session record not found, create it

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent.
func (store *sqlStore) NextSenderMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received.
func (store *sqlStore) NextTargetMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent.
func (store *sqlStore) SetNextSenderMsgSeqNum(next int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received.
func (store *sqlStore) SetNextTargetMsgSeqNum(next int) error {
	_ = "STUB: not implemented"
	return nil
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent.
func (store *sqlStore) IncrNextSenderMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received.
func (store *sqlStore) IncrNextTargetMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

// CreationTime returns the creation time of the store.
func (store *sqlStore) CreationTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// SetCreationTime is a no-op for SQLStore.
func (store *sqlStore) SetCreationTime(_ time.Time) { _ = "STUB: not implemented"; return }

func (store *sqlStore) SaveMessage(seqNum int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *sqlStore) SaveMessageAndIncrNextSenderMsgSeqNum(seqNum int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *sqlStore) IterateMessages(beginSeqNum, endSeqNum int, cb func([]byte) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *sqlStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the store's database connection.
func (store *sqlStore) Close() error { _ = "STUB: not implemented"; return nil }
