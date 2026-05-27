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

type resendState struct {
	loggedOn
	messageStash          map[int]*Message
	currentResendRangeEnd int
	resendRangeEnd        int
}

func (s resendState) String() string { _ = "STUB: not implemented"; return "" }

func (s resendState) Timeout(session *session, event internal.Event) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}

// Wrap pendingTimeout in resend. prevents us falling back to inSession if recovering
// from pendingTimeout.

func (s resendState) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	_ = "STUB: not implemented"
	return *new(sessionState)
}
