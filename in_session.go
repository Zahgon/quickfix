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
	"github.com/quickfixgo/quickfix/internal"
)

type inSession struct{ loggedOn }

func (state inSession) String() string { _ = "STUB: not implemented"; return "" }

func (state inSession) FixMsgIn(session *session, msg *Message) sessionState {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state inSession) Timeout(session *session, event internal.Event) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state inSession) handleLogout(session *session, msg *Message) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state inSession) handleTestRequest(session *session, msg *Message) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state inSession) handleSequenceReset(session *session, msg *Message) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

// FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess).

func (state inSession) handleResendRequest(session *session, msg *Message) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state inSession) resendMessages(session *session, beginSeqNo, endSeqNo int, inReplyTo Message) error {
	_ = "STUB: not implemented"
	return nil
}

// resendMutex must always be locked before sendMutex to prevent a potential deadlock
// sendMutex is locked below in session.EnqueueBytesAndSend()

// We cant continue with a message that cant be parsed correctly.

// workaround for maintaining repeating group field order

// gapfill for catch-up

func (state inSession) processReject(session *session, msg *Message, rej MessageRejectError) sessionState {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

// Assumes target too high reject already sent.

func (state inSession) doTargetTooLow(session *session, msg *Message, rej targetTooLow) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state *inSession) generateSequenceReset(session *session, beginSeqNo int, endSeqNo int, inReplyTo Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}
