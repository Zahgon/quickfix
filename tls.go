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
)

func loadTLSConfig(settings *SessionSettings) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaultTLSConfig brought to you by https://github.com/gtank/cryptopasta/
func defaultTLSConfig() *tls.Config {
	_ = "STUB: not implemented"

	// Avoids most of the memorably-named TLS attacks
	return nil
}

// Causes servers to use Go's default ciphersuite preferences,
// which are tuned to avoid attacks. Does nothing on clients.

// Only use curves which have constant-time implementations

func setMinVersionExplicit(settings *SessionSettings, tlsConfig *tls.Config) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck // SA1019 min version ok
