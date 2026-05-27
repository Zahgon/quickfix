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
	"io"
	"time"
)

const (
	defaultBufSize = 4096
)

type parser struct {
	// Buffer is a slice of bigBuffer.
	bigBuffer, buffer []byte
	reader            io.Reader
	lastRead          time.Time
}

func newParser(reader io.Reader) *parser { _ = "STUB: not implemented"; return nil }

func (p *parser) readMore() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Initialize the parser.

// Shift buffer back to the start of bigBuffer.

// Reallocate big buffer with enough space to shift buffer.

func (p *parser) findIndex(delim []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *parser) findIndexAfterOffset(offset int, delim []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *parser) findStart() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *parser) findEndAfterOffset(offset int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *parser) jumpLength() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *parser) ReadMessage() (msgBytes *bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
