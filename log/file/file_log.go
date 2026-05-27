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
	"log"

	"github.com/quickfixgo/quickfix"
)

type fileLog struct {
	eventLogger   *log.Logger
	messageLogger *log.Logger
}

func (l fileLog) OnIncoming(msg []byte) { _ = "STUB: not implemented"; return }

func (l fileLog) OnOutgoing(msg []byte) { _ = "STUB: not implemented"; return }

func (l fileLog) OnEvent(msg string) { _ = "STUB: not implemented"; return }

func (l fileLog) OnEventf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

type fileLogFactory struct {
	globalLogPath   string
	sessionLogPaths map[quickfix.SessionID]string
}

// NewLogFactory creates an instance of LogFactory that writes messages and events to file.
// The location of global and session log files is configured via FileLogPath.
func NewLogFactory(settings *quickfix.Settings) (quickfix.LogFactory, error) {
	_ = "STUB: not implemented"
	return *new(quickfix.LogFactory), nil
}

func newFileLog(prefix string, logPath string) (fileLog, error) {
	_ = "STUB: not implemented"
	return *new(fileLog), nil
}

func (f fileLogFactory) Create() (quickfix.Log, error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}

func (f fileLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (quickfix.Log, error) {
	_ = "STUB: not implemented"
	return *new(quickfix.Log), nil
}
