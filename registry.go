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
	"errors"
	"sync"
)

var sessionsLock sync.RWMutex
var sessions = make(map[SessionID]*session)
var errDuplicateSessionID = errors.New("Duplicate SessionID")
var errUnknownSession = errors.New("Unknown session")

// Messagable is a Message or something that can be converted to a Message.
type Messagable interface {
	ToMessage() *Message
}

// Send determines the session to send Messagable using header fields BeginString, TargetCompID, SenderCompID.
func Send(m Messagable) (err error) { _ = "STUB: not implemented"; return nil }

// SendToTarget sends a message based on the sessionID. Convenient for use in FromApp since it provides a session ID for incoming messages.
func SendToTarget(m Messagable, sessionID SessionID) error { _ = "STUB: not implemented"; return nil }

// ResetSession resets session's sequence numbers.
func ResetSession(sessionID SessionID) error { _ = "STUB: not implemented"; return nil }

// UnregisterSession removes a session from the set of known sessions.
func UnregisterSession(sessionID SessionID) error { _ = "STUB: not implemented"; return nil }

// SetNextTargetMsgSeqNum set the next expected target message sequence number for the session matching the session id.
func SetNextTargetMsgSeqNum(sessionID SessionID, seqNum int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetNextSenderMsgSeqNum sets the next outgoing message sequence number for the session matching the session id.
func SetNextSenderMsgSeqNum(sessionID SessionID, seqNum int) error {
	_ = "STUB: not implemented"
	return nil
}

// GetExpectedSenderNum retrieves the expected sender sequence number for the session matching the session id.
func GetExpectedSenderNum(sessionID SessionID) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetExpectedTargetNum retrieves the next target sequence number for the session matching the session id.
func GetExpectedTargetNum(sessionID SessionID) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetMessageStore returns the MessageStore interface for session matching the session id.
func GetMessageStore(sessionID SessionID) (MessageStore, error) {
	_ = "STUB: not implemented"
	return *new(MessageStore), nil
}

// GetLog returns the Log interface for session matching the session id.
func GetLog(sessionID SessionID) (Log, error) { _ = "STUB: not implemented"; return *new(Log), nil }

func registerSession(s *session) error { _ = "STUB: not implemented"; return nil }

func lookupSession(sessionID SessionID) (s *session, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
