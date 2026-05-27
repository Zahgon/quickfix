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
	_ = "STUB: not implemented"
	return ""
}

func postgresPlaceholder(i int) string { _ = "STUB: not implemented"; return "" }

// NewLogFactory returns a sql-based implementation of LogFactory.
func NewLogFactory(settings *quickfix.Settings) quickfix.LogFactory {
	_ = "STUB: not implemented"
	return *new(quickfix.LogFactory)
}

// Create creates a new SQLLog implementation of the Log interface.
func (f sqlLogFactory) Create() (log quickfix.Log, err error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}

// CreateSessionLog creates a new SQLLog implementation of the Log interface.
func (f sqlLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (log quickfix.Log, err error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}

func newSQLLog(sessionID quickfix.SessionID, driver string, dataSourceName string, connMaxLifetime time.Duration) (l *sqlLog, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ensure immediate connection

func (l sqlLog) OnIncoming(msg []byte) { _ = "STUB: not implemented"; return }

func (l sqlLog) OnOutgoing(msg []byte) { _ = "STUB: not implemented"; return }

func (l sqlLog) OnEvent(msg string) { _ = "STUB: not implemented"; return }

func (l sqlLog) OnEventf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l sqlLog) insert(table string, value string) { _ = "STUB: not implemented"; return }

func (l *sqlLog) iterate(table string, cb func(string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *sqlLog) getEntries(table string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the log's database connection.
func (l *sqlLog) close() error { _ = "STUB: not implemented"; return nil }
