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
	"strconv"
)

// TagValue is a low-level FIX field abstraction.
type TagValue struct {
	tag   Tag
	value []byte
	bytes []byte
}

func (tv *TagValue) init(tag Tag, value []byte) {
	tv.bytes = strconv.AppendInt(nil, int64(tag), 10)
	tv.bytes = append(tv.bytes, []byte("=")...)
	tv.bytes = append(tv.bytes, value...)
	tv.bytes = append(tv.bytes, []byte("")...)

	tv.tag = tag
	tv.value = value
}

func (tv *TagValue) parse(rawFieldBytes []byte) error {
	_ = "STUB: not implemented"

	// Most of the Fix tags are 4 or less characters long, so we can optimize
	// for that by checking the 5 first characters without looping over the
	// whole byte slice.
	return nil
}

func (tv TagValue) String() string { _ = "STUB: not implemented"; return "" }

func bytesTotal(bytes []byte) (total int) { _ = "STUB: not implemented"; return 0 }

func (tv TagValue) total() int { _ = "STUB: not implemented"; return 0 }

func (tv TagValue) length() int { _ = "STUB: not implemented"; return 0 }
