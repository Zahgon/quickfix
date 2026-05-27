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

package file

import (
	"os"
	"sync"
	"time"

	"github.com/quickfixgo/quickfix"
)

type fileStoreFactory struct {
	settings *quickfix.Settings
}

type fileStore struct {
	sessionID          quickfix.SessionID
	cache              quickfix.MessageStore
	bodyFname          string
	headerFname        string
	sessionFname       string
	senderSeqNumsFname string
	targetSeqNumsFname string

	fileMu            sync.Mutex
	bodyFile          *os.File
	headerFile        *os.File
	sessionFile       *os.File
	senderSeqNumsFile *os.File
	targetSeqNumsFile *os.File
	fileSync          bool
}

// NewStoreFactory returns a file-based implementation of MessageStoreFactory.
func NewStoreFactory(settings *quickfix.Settings) quickfix.MessageStoreFactory {
	_ = "STUB: not implemented"
	return *new(quickfix.MessageStoreFactory)
}

// Create creates a new FileStore implementation of the MessageStore interface.
func (f fileStoreFactory) Create(sessionID quickfix.SessionID) (msgStore quickfix.MessageStore, err error) {
	_ = "STUB: not implemented"
	return *new(quickfix.MessageStore), nil
}

//existing behavior is to fsync writes

func newFileStore(sessionID quickfix.SessionID, dirname string, fileSync bool) (*fileStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset deletes the store files and sets the seqnums back to 1.
func (store *fileStore) Reset() error { _ = "STUB: not implemented"; return nil }

// Refresh closes the store files and then reloads from them.
func (store *fileStore) Refresh() (err error) { _ = "STUB: not implemented"; return nil }

func (store *fileStore) populateCache() (creationTimePopulated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (store *fileStore) setSession() error { _ = "STUB: not implemented"; return nil }

func (store *fileStore) setSeqNum(f *os.File, seqNum int) error {
	_ = "STUB: not implemented"
	return nil
}

// NextSenderMsgSeqNum returns the next MsgSeqNum that will be sent.
func (store *fileStore) NextSenderMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

// NextTargetMsgSeqNum returns the next MsgSeqNum that should be received.
func (store *fileStore) NextTargetMsgSeqNum() int { _ = "STUB: not implemented"; return 0 }

// SetNextSenderMsgSeqNum sets the next MsgSeqNum that will be sent.
func (store *fileStore) SetNextSenderMsgSeqNum(next int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetNextTargetMsgSeqNum sets the next MsgSeqNum that should be received.
func (store *fileStore) SetNextTargetMsgSeqNum(next int) error {
	_ = "STUB: not implemented"
	return nil
}

// IncrNextSenderMsgSeqNum increments the next MsgSeqNum that will be sent.
func (store *fileStore) IncrNextSenderMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

// IncrNextTargetMsgSeqNum increments the next MsgSeqNum that should be received.
func (store *fileStore) IncrNextTargetMsgSeqNum() error { _ = "STUB: not implemented"; return nil }

// CreationTime returns the creation time of the store.
func (store *fileStore) CreationTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// SetCreationTime is a no-op for FileStore.
func (store *fileStore) SetCreationTime(_ time.Time) { _ = "STUB: not implemented"; return }

func (store *fileStore) SaveMessage(seqNum int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *fileStore) SaveMessageAndIncrNextSenderMsgSeqNum(seqNum int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *fileStore) syncBodyAndHeaderFilesLocked() error { _ = "STUB: not implemented"; return nil }

func (store *fileStore) IterateMessages(beginSeqNum, endSeqNum int, cb func([]byte) error) error {
	_ = "STUB: not implemented"
	// Sync files
	return nil
}

// Open a read only view to body and header file

// Iterate over the header file

// If we have reached the end of possible iteration then break

// If we have not yet reached the starting sequence number then continue

// Otherwise process the file

func (store *fileStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the store's files.
func (store *fileStore) Close() error { _ = "STUB: not implemented"; return nil }
