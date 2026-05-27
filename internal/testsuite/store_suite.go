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

package testsuite

import (
	"github.com/quickfixgo/quickfix"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
	MsgStore quickfix.MessageStore
}

func (s *StoreTestSuite) TestMessageStoreSetNextMsgSeqNumRefreshIncrNextMsgSeqNum() {
	_ = "STUB: not implemented"
	// Given a MessageStore with the following sender and target seqnums
	return
}

// When the store is refreshed from its backing store

// Then the sender and target seqnums should still be

// When the sender and target seqnums are incremented

// Then the sender and target seqnums should be

// When the store is refreshed from its backing store

// Then the sender and target seqnums should still be

func (s *StoreTestSuite) TestMessageStoreReset() {
	_ = "STUB: not implemented"
	// Given a MessageStore with the following sender and target seqnums
	return
}

// When the store is reset

// Then the sender and target seqnums should be

// When the store is refreshed from its backing store

// Then the sender and target seqnums should still be

func (s *StoreTestSuite) fetchMessages(beginSeqNum, endSeqNum int) (msgs [][]byte) {
	_ = "STUB: not implemented"

	// Fetch messages from the new iterator
	return nil
}

// Fetch messages from the old getter

// Ensure the output is the same

func (s *StoreTestSuite) TestMessageStoreSaveMessageGetMessage() {
	_ = "STUB: not implemented"
	// Given the following saved messages
	return
}

// When the messages are retrieved from the MessageStore

// Then the messages should be

// When the store is refreshed from its backing store

// And the messages are retrieved from the MessageStore

// Then the messages should still be

func (s *StoreTestSuite) TestMessageStoreSaveMessageAndIncrementGetMessage() {
	_ = "STUB: not implemented"
	return
}

// Given the following saved messages

// When the messages are retrieved from the MessageStore

// Then the messages should be

// When the store is refreshed from its backing store

// And the messages are retrieved from the MessageStore

// Then the messages should still be

func (s *StoreTestSuite) TestMessageStoreGetMessagesEmptyStore() {
	_ = "STUB: not implemented"
	// When messages are retrieved from an empty store
	return
}

// Then no messages should be returned

func (s *StoreTestSuite) TestMessageStoreGetMessagesVariousRanges() {
	_ = "STUB: not implemented"

	// Given the following saved messages
	return
}

// When the following requests are made to the store

// Then the returned messages should be

func (s *StoreTestSuite) TestMessageStoreCreationTime() { _ = "STUB: not implemented"; return }
