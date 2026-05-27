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

type routeKey struct {
	FIXVersion string
	MsgType    string
}

// FIX ApplVerID string values.
const (
	ApplVerIDFIX27    = "0"
	ApplVerIDFIX30    = "1"
	ApplVerIDFIX40    = "2"
	ApplVerIDFIX41    = "3"
	ApplVerIDFIX42    = "4"
	ApplVerIDFIX43    = "5"
	ApplVerIDFIX44    = "6"
	ApplVerIDFIX50    = "7"
	ApplVerIDFIX50SP1 = "8"
	ApplVerIDFIX50SP2 = "9"
)

// A MessageRoute is a function that can process a fromApp/fromAdmin callback.
type MessageRoute func(msg *Message, sessionID SessionID) MessageRejectError

// A MessageRouter is a mutex for MessageRoutes.
type MessageRouter struct {
	routes map[routeKey]MessageRoute
}

// NewMessageRouter returns an initialized MessageRouter instance.
func NewMessageRouter() *MessageRouter { _ = "STUB: not implemented"; return nil }

// AddRoute adds a route to the MessageRouter instance keyed to begin string and msgType.
func (c MessageRouter) AddRoute(beginString string, msgType string, router MessageRoute) {
	_ = "STUB: not implemented"
	return
}

// Route may be called from the fromApp/fromAdmin callbacks. Messages that cannot be routed will be rejected with UnsupportedMessageType.
func (c MessageRouter) Route(msg *Message, sessionID SessionID) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func (c MessageRouter) tryRoute(beginString string, msgType string, msg *Message, sessionID SessionID) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}
