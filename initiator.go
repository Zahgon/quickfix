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
	"crypto/tls"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

// Initiator initiates connections and processes messages for all sessions.
type Initiator struct {
	app             Application
	settings        *Settings
	sessionSettings map[SessionID]*SessionSettings
	storeFactory    MessageStoreFactory
	logFactory      LogFactory
	globalLog       Log
	stopChan        chan interface{}
	wg              sync.WaitGroup
	sessions        map[SessionID]*session
	sessionFactory
}

// Start Initiator.
func (i *Initiator) Start() (err error) { _ = "STUB: not implemented"; return nil }

// TODO: move into session factory.

// Stop Initiator.
func (i *Initiator) Stop() { _ = "STUB: not implemented"; return }

// Closed already.

// NewInitiator creates and initializes a new Initiator.
func NewInitiator(app Application, storeFactory MessageStoreFactory, appSettings *Settings, logFactory LogFactory) (*Initiator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// waitForInSessionTime returns true if the session is in session, false if the handler should stop.
func (i *Initiator) waitForInSessionTime(session *session) bool {
	_ = "STUB: not implemented"
	return false
}

// waitForReconnectInterval returns true if a reconnect should be re-attempted, false if handler should stop.
func (i *Initiator) waitForReconnectInterval(reconnectInterval time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (i *Initiator) handleConnection(session *session, tlsConfig *tls.Config, dialer proxy.ContextDialer) {
	_ = "STUB: not implemented"
	return
}

// We start a goroutine in order to be able to cancel the dialer mid-connection
// on receiving a stop signal to stop the initiator.

// Unless InsecureSkipVerify is true, server name config is required for TLS
// to verify the received certificate

// This ensures we properly cleanup the goroutine and context used for
// dial cancelation after successful connection.
