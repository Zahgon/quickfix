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

import "bytes"

// SessionID is a unique identifier of a Session.
type SessionID struct {
	BeginString, TargetCompID, TargetSubID, TargetLocationID, SenderCompID, SenderSubID, SenderLocationID, Qualifier string
}

// IsFIXT returns true if the SessionID has a FIXT BeginString.
func (s SessionID) IsFIXT() bool { _ = "STUB: not implemented"; return false }

func appendOptional(b *bytes.Buffer, delim, v string) { _ = "STUB: not implemented"; return }

func (s SessionID) String() string { _ = "STUB: not implemented"; return "" }
