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

package composite

import (
	"github.com/quickfixgo/quickfix"
)

type compositeLog struct {
	logs []quickfix.Log
}

func (l compositeLog) OnIncoming(s []byte) { _ = "STUB: not implemented"; return }

func (l compositeLog) OnOutgoing(s []byte) { _ = "STUB: not implemented"; return }

func (l compositeLog) OnEvent(s string) { _ = "STUB: not implemented"; return }

func (l compositeLog) OnEventf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

type compositeLogFactory struct {
	logFactories []quickfix.LogFactory
}

func (clf compositeLogFactory) Create() (quickfix.Log, error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}

func (clf compositeLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (quickfix.Log, error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}

// NewLogFactory creates an instance of LogFactory that writes messages and events to stdout.
func NewLogFactory(logfactories []quickfix.LogFactory) quickfix.LogFactory {
	_ = "STUB: not implemented"
	return *new(quickfix.LogFactory)
}
