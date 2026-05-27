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

var dayLookup = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,

	"Sun": time.Sunday,
	"Mon": time.Monday,
	"Tue": time.Tuesday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
}

var applVerIDLookup = map[string]string{
	BeginStringFIX40: "2",
	BeginStringFIX41: "3",
	BeginStringFIX42: "4",
	BeginStringFIX43: "5",
	BeginStringFIX44: "6",
	"FIX.5.0":        "7",
	"FIX.5.0SP1":     "8",
	"FIX.5.0SP2":     "9",
}

type sessionFactory struct {
	// True if building sessions that initiate logon.
	BuildInitiators bool
}

const shortForm = "15:04:05"

// Creates Session, associates with internal session registry.
func (f sessionFactory) createSession(
	sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings,
	logFactory LogFactory, application Application,
) (session *session, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f sessionFactory) newSession(
	sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings, logFactory LogFactory,
	application Application) (s *session, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Always use a default message validator without data dictionaries

// If the transport or app data dictionary setting is set, the other also needs to be set.

// Default to 1 buffered message per channel

func (f sessionFactory) buildAcceptorSettings(session *session, settings *SessionSettings) error {
	_ = "STUB: not implemented"
	return nil
}

func (f sessionFactory) buildInitiatorSettings(session *session, settings *SessionSettings) error {
	_ = "STUB: not implemented"
	return nil
}

func (f sessionFactory) configureSocketConnectAddress(session *session, settings *SessionSettings) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (f sessionFactory) buildHeartBtIntSettings(session *session, settings *SessionSettings, mustProvide bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}
