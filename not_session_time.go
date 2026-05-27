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

import "github.com/quickfixgo/quickfix/internal"

type notSessionTime struct{ latentState }

func (notSessionTime) String() string      { _ = "STUB: not implemented"; return "" }
func (notSessionTime) IsSessionTime() bool { _ = "STUB: not implemented"; return false }

func (state notSessionTime) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state notSessionTime) Timeout(*session, internal.Event) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

func (state notSessionTime) Stop(*session) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}
