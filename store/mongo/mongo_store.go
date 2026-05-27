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

type mongoStoreFactory struct {
	settings           *quickfix.Settings
	messagesCollection string
	sessionsCollection string
}

type mongoStore struct {
	sessionID          quickfix.SessionID
	cache              quickfix.MessageStore
	mongoURL           string
	mongoDatabase      string
	db                 *mongo.Client
	messagesCollection string
	sessionsCollection string
	allowTransactions  bool
}

// NewStoreFactory returns a mongo-based implementation of MessageStoreFactory.
func NewStoreFactory(settings *quickfix.Settings) quickfix.MessageStoreFactory {
	_ = "STUB: not implemented"
	return *new(quickfix.MessageStoreFactory)
}

// NewStoreFactoryPrefixed returns a mongo-based implementation of MessageStoreFactory, with prefix on collections.
func NewStoreFactoryPrefixed(settings *quickfix.Settings, collectionsPrefix string) quickfix.MessageStoreFactory {
	_ = "STUB: not implemented"
	return *new(quickfix.MessageStoreFactory)
}

// Create creates a new MongoStore implementation of the MessageStore interface.
func (f mongoStoreFactory) Create(sessionID quickfix.SessionID) (msgStore quickfix.MessageStore, err error) {
	_ = "STUB: not implemented"
	return *new(quickfix.MessageStore), nil
}

// Optional.

func newMongoStore(sessionID quickfix.SessionID, mongoURL, mongoDatabase, mongoReplicaSet, messagesCollection, sessionsCollection string) (store *mongoStore, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateMessageFilter(s *quickfix.SessionID) (messageFilter *mongoQuickFixEntryData) {
	_ = "STUB: not implemented"
	return nil
}

type mongoQuickFixEntryData struct {
	// Message specific data.
	Msgseq  int    `bson:"msgseq,omitempty"`
	Message []byte `bson:"message,omitempty"`
	// Session specific data.
	CreationTime   time.Time `bson:"creation_time,omitempty"`
	IncomingSeqNum int       `bson:"incoming_seq_num,omitempty"`
	OutgoingSeqNum int       `bson:"outgoing_seq_num,omitempty"`
	// Indexed data.
	BeginString      string `bson:"begin_string"`
	SessionQualifier string `bson:"session_qualifier"`
	SenderCompID     string `bson:"sender_comp_id"`
	SenderSubID      string `bson:"sender_sub_id"`
	SenderLocID      string `bson:"sender_loc_id"`
	TargetCompID     string `bson:"target_comp_id"`
	TargetSubID      string `bson:"target_sub_id"`
	TargetLocID      string `bson:"target_loc_id"`
}

// Reset deletes the store records and sets the seqnums back to 1.
func (store *mongoStore) Reset() error { _ = "STUB: not implemented"; return nil }

// Refresh reloads the store from the database.
func (store *mongoStore) Refresh() error { _ = "STUB: not implemented"; return nil }

func (store *mongoStore) populateCache() error { _ = "STUB: not implemented"; return nil }

// session record found, load it

// session record not found, create it

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent.
func (store *mongoStore) NextSenderMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received.
func (store *mongoStore) NextTargetMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent.
func (store *mongoStore) SetNextSenderMsgSeqNum(next int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received.
func (store *mongoStore) SetNextTargetMsgSeqNum(next int) error {
	_ = "STUB: not implemented"
	return nil
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent.
func (store *mongoStore) IncrNextSenderMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received.
func (store *mongoStore) IncrNextTargetMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

// CreationTime returns the creation time of the store.
func (store *mongoStore) CreationTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// SetCreationTime is a no-op for MongoStore.
func (store *mongoStore) SetCreationTime(_ time.Time) { _ = "STUB: not implemented"; return }

func (store *mongoStore) SaveMessage(seqNum int, msg []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (store *mongoStore) SaveMessageAndIncrNextSenderMsgSeqNum(seqNum int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// If the mongodb supports replicasets, perform this operation as a transaction instead-

func (store *mongoStore) IterateMessages(beginSeqNum, endSeqNum int, cb func([]byte) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Marshal into database form.

// Modify the query to use a range for the sequence filter.

func (store *mongoStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the store's database connection.
func (store *mongoStore) Close() error { _ = "STUB: not implemented"; return nil }
