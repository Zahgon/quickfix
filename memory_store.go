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

package quickfix

import (
	"time"
)

type memoryStore struct {
	senderMsgSeqNum, targetMsgSeqNum int
	creationTime                     time.Time
	messageMap                       map[int][]byte
}

func (store *memoryStore) NextSenderMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

func (store *memoryStore) NextTargetMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

func (store *memoryStore) IncrNextSenderMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

func (store *memoryStore) IncrNextTargetMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

func (store *memoryStore) SetNextSenderMsgSeqNum(nextSeqNum int) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *memoryStore) SetNextTargetMsgSeqNum(nextSeqNum int) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *memoryStore) CreationTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (store *memoryStore) SetCreationTime(t time.Time) { _ = "STUB: not implemented"; return }

func (store *memoryStore) Reset() error { _ = "STUB: not implemented"; return nil }

func (store *memoryStore) Refresh() error {
	_ = "STUB: not implemented"
	// NOP, nothing to refresh.
	return nil
}

func (store *memoryStore) Close() error {
	_ = "STUB: not implemented"
	// NOP, nothing to close.
	return nil
}

func (store *memoryStore) SaveMessage(seqNum int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *memoryStore) SaveMessageAndIncrNextSenderMsgSeqNum(seqNum int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *memoryStore) IterateMessages(beginSeqNum, endSeqNum int, cb func([]byte) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *memoryStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type memoryStoreFactory struct{}

func (f memoryStoreFactory) Create(_ SessionID) (MessageStore, error) {
	_ = "STUB: not implemented"
	return *new(MessageStore), nil
}

// NewMemoryStoreFactory returns a MessageStoreFactory instance that created in-memory MessageStores.
func NewMemoryStoreFactory() MessageStoreFactory {
	_ = "STUB: not implemented"
	return *new(MessageStoreFactory)
}
