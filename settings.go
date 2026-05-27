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
	"io"
)

// The Settings type represents a collection of global and session settings.
type Settings struct {
	globalSettings  *SessionSettings
	sessionSettings map[SessionID]*SessionSettings
}

// Init initializes or resets a Settings instance.
func (s *Settings) Init() { _ = "STUB: not implemented"; return }

func (s *Settings) lazyInit() { _ = "STUB: not implemented"; return }

// NewSettings creates a Settings instance.
func NewSettings() *Settings { _ = "STUB: not implemented"; return nil }

func sessionIDFromSessionSettings(globalSettings *SessionSettings, sessionSettings *SessionSettings) SessionID {
	_ = "STUB: not implemented"
	return *new(SessionID)
}

// ParseSettings creates and initializes a Settings instance with config parsed from a Reader.
// Returns error if the config is has parse errors.
func ParseSettings(reader io.Reader) (*Settings, error) { _ = "STUB: not implemented"; return nil, nil }

// GlobalSettings are default setting inherited by all session settings.
func (s *Settings) GlobalSettings() *SessionSettings { _ = "STUB: not implemented"; return nil }

// SessionSettings return all session settings overlaying globalsettings.
func (s *Settings) SessionSettings() map[SessionID]*SessionSettings {
	_ = "STUB: not implemented"
	return nil
}

// AddSession adds Session Settings to Settings instance. Returns an error if session settings with duplicate sessionID has already been added.
func (s *Settings) AddSession(sessionSettings *SessionSettings) (SessionID, error) {
	_ = "STUB: not implemented"
	return *new(SessionID), nil
}
