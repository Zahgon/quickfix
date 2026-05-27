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
	"time"
)

// SessionSettings maps session settings to values with typed accessors.
type SessionSettings struct {
	settings map[string][]byte
}

// ConditionallyRequiredSetting indicates a missing setting.
type ConditionallyRequiredSetting struct {
	Setting string
}

func (e ConditionallyRequiredSetting) Error() string { _ = "STUB: not implemented"; return "" }

// IncorrectFormatForSetting indicates a setting that is incorrectly formatted.
type IncorrectFormatForSetting struct {
	Setting string
	Value   []byte
	Err     error
}

func (e IncorrectFormatForSetting) Error() string { _ = "STUB: not implemented"; return "" }

// Init initializes or resets SessionSettings.
func (s *SessionSettings) Init() { _ = "STUB: not implemented"; return }

// NewSessionSettings returns a newly initialized SessionSettings instance.
func NewSessionSettings() *SessionSettings { _ = "STUB: not implemented"; return nil }

// SetRaw assigns a value to a setting on SessionSettings.
func (s *SessionSettings) SetRaw(setting string, val []byte) {
	_ = "STUB: not implemented"
	// Lazy init.
	return
}

// Set assigns a string value to a setting on SessionSettings.
func (s *SessionSettings) Set(setting string, val string) {
	_ = "STUB: not implemented"
	// Lazy init
	return
}

// HasSetting returns true if a setting is set, false if not.
func (s *SessionSettings) HasSetting(setting string) bool { _ = "STUB: not implemented"; return false }

// RawSetting is a settings accessor that returns the raw byte slice value of
// the setting. Returns an error if the setting is missing.
func (s *SessionSettings) RawSetting(setting string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setting is a settings string accessor. Returns an error if the setting is missing.
func (s *SessionSettings) Setting(setting string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IntSetting returns the requested setting parsed as an int.  Returns an errror if the setting is not set or cannot be parsed as an int.
func (s *SessionSettings) IntSetting(setting string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DurationSetting returns the requested setting parsed as a time.Duration.
// Returns an error if the setting is not set or cannot be parsed as a time.Duration.
func (s *SessionSettings) DurationSetting(setting string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// BoolSetting returns the requested setting parsed as a boolean.  Returns an error if the setting is not set or cannot be parsed as a bool.
func (s SessionSettings) BoolSetting(setting string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *SessionSettings) overlay(overlay *SessionSettings) { _ = "STUB: not implemented"; return }

func (s *SessionSettings) clone() *SessionSettings { _ = "STUB: not implemented"; return nil }
