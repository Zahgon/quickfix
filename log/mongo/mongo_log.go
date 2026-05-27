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
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/quickfixgo/quickfix"
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
	_ = "STUB: not implemented"
	return *new(quickfix.LogFactory)
}

// NewLogFactoryPrefixed returns a mongo-based implementation of LogFactory, with prefix on collections.
func NewLogFactoryPrefixed(settings *quickfix.Settings, collectionsPrefix string) quickfix.LogFactory {
	_ = "STUB: not implemented"
	return *new(quickfix.LogFactory)
}

// Create creates a new mongo implementation of the Log interface.
func (f mongoLogFactory) Create() (l quickfix.Log, err error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}

// Optional.

// CreateSessionLog creates a new mongo implementation of the Log interface.
func (f mongoLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (l quickfix.Log, err error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}

// Optional.

func newmongoLog(sessionID quickfix.SessionID, mongoURL, mongoDatabase, mongoReplicaSet, messagesLogCollection, eventLogCollection string) (l *mongoLog, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l mongoLog) OnIncoming(msg []byte) { _ = "STUB: not implemented"; return }

func (l mongoLog) OnOutgoing(msg []byte) { _ = "STUB: not implemented"; return }

func (l mongoLog) OnEvent(msg string) { _ = "STUB: not implemented"; return }

func (l mongoLog) OnEventf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func generateEntry(s *quickfix.SessionID) (entry *entryData) { _ = "STUB: not implemented"; return nil }

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

func (l *mongoLog) insert(collection string, text []byte) { _ = "STUB: not implemented"; return }

func (l *mongoLog) iterate(coll string, cb func(string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *mongoLog) getEntries(coll string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// close closes the l's database connection.
func (l *mongoLog) close() error { _ = "STUB: not implemented"; return nil }
