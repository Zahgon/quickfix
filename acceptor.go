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
	"bytes"
	"crypto/tls"
	"net"
	"sync"
)

// Acceptor accepts connections from FIX clients and manages the associated sessions.
type Acceptor struct {
	app                   Application
	settings              *Settings
	logFactory            LogFactory
	storeFactory          MessageStoreFactory
	globalLog             Log
	sessions              map[SessionID]*session
	sessionGroup          sync.WaitGroup
	listenerShutdown      sync.WaitGroup
	dynamicSessions       bool
	dynamicQualifier      bool
	dynamicQualifierCount int
	dynamicSessionChan    chan *session
	sessionAddr           sync.Map
	sessionHostPort       map[SessionID]int
	listeners             map[string]net.Listener
	connectionValidator   ConnectionValidator
	tlsConfig             *tls.Config
	newListenerCallback   NewListenerCallback
	sessionFactory
}

// ConnectionValidator is an interface allowing to implement a custom authentication logic.
type ConnectionValidator interface {
	// Validate the connection for validity. This can be a part of authentication process.
	// For example, you may tie up a SenderCompID to an IP range, or to a specific TLS certificate as a part of mTLS.
	Validate(netConn net.Conn, session SessionID) error
}

// NewListenerCallback is a function that returns a net.Listener for the given address and tls.Config struct.
type NewListenerCallback func(address string, tlsConfig *tls.Config) (net.Listener, error)

// Start accepting connections.
func (a *Acceptor) Start() (err error) { _ = "STUB: not implemented"; return nil }

// Stop logs out existing sessions, close their connections, and stop accepting new connections.
func (a *Acceptor) Stop() { _ = "STUB: not implemented"; return }

// suppress sending on closed channel error

// RemoteAddr gets remote IP address for a given session.
func (a *Acceptor) RemoteAddr(sessionID SessionID) (net.Addr, bool) {
	_ = "STUB: not implemented"
	return *new(net.Addr), false
}

// NewAcceptor creates and initializes a new Acceptor.
func NewAcceptor(app Application, storeFactory MessageStoreFactory, settings *Settings, logFactory LogFactory) (a *Acceptor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Acceptor) listenForConnections(listener net.Listener) { _ = "STUB: not implemented"; return }

func (a *Acceptor) invalidMessage(msg *bytes.Buffer, err error) { _ = "STUB: not implemented"; return }

func (a *Acceptor) handleConnection(netConn net.Conn) { _ = "STUB: not implemented"; return }

// We have a session ID and a network connection. This seems to be a good place for any custom authentication logic.

func (a *Acceptor) dynamicSessionsLoop() { _ = "STUB: not implemented"; return }

// SetConnectionValidator sets an optional connection validator.
// Use it when you need a custom authentication logic that includes lower level interactions,
// like mTLS auth or IP whitelistening.
// To remove a previously set validator call it with a nil value:
//
//	a.SetConnectionValidator(nil)
func (a *Acceptor) SetConnectionValidator(validator ConnectionValidator) {
	_ = "STUB: not implemented"
	return
}

// SetTLSConfig allows the creator of the Acceptor to specify a fully customizable tls.Config of their choice,
// which will be used in the Start() method.
//
// Note: when the caller explicitly provides a tls.Config with this function,
// it takes precendent over TLS settings specified in the acceptor's settings.GlobalSettings(),
// meaning that the `settings.GlobalSettings()` object is not inspected or used for the creation of the tls.Config.
func (a *Acceptor) SetTLSConfig(tlsConfig *tls.Config) { _ = "STUB: not implemented"; return }

// SetNewListenerCallback allows the creator of the Acceptor to specify the callback used to create each net.Listener
// which will be used in the Start() method.
func (a *Acceptor) SetNewListenerCallback(cb NewListenerCallback) {
	_ = "STUB: not implemented"
	return
}
