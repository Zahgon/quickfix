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

// TimestampPrecision defines the precision used by FIXUTCTimestamp.
type TimestampPrecision int

// All TimestampPrecisions supported by FIX.
const (
	Millis TimestampPrecision = iota
	Seconds
	Micros
	Nanos
)

// FIXUTCTimestamp is a FIX UTC Timestamp value, implements FieldValue.
type FIXUTCTimestamp struct {
	time.Time
	Precision TimestampPrecision
}

const (
	utcTimestampMillisFormat  = "20060102-15:04:05.000"
	utcTimestampSecondsFormat = "20060102-15:04:05"
	utcTimestampMicrosFormat  = "20060102-15:04:05.000000"
	utcTimestampNanosFormat   = "20060102-15:04:05.000000000"
)

func (f *FIXUTCTimestamp) Read(bytes []byte) (err error) {
	_ = "STUB: not implemented"

	// Seconds.
	return nil
}

// Millis.

// micros

// nanos

func (f FIXUTCTimestamp) Write() []byte { _ = "STUB: not implemented"; return nil }
