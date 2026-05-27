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
	"sync"
	"time"
)

// field stores a slice of TagValues.
type field []TagValue

func fieldTag(f field) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func initField(f field, tag Tag, value []byte) { _ = "STUB: not implemented"; return }

func writeField(f field, buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }

// tagOrder true if tag i should occur before tag j.
type tagOrder func(i, j Tag) bool

type tagSort struct {
	tags    []Tag
	compare tagOrder
}

func (t tagSort) Len() int           { _ = "STUB: not implemented"; return 0 }
func (t tagSort) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (t tagSort) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// FieldMap is a collection of fix fields that make up a fix message.
type FieldMap struct {
	tagLookup map[Tag]field
	tagSort
	rwLock *sync.RWMutex
}

// ascending tags.
func normalFieldOrder(i, j Tag) bool { _ = "STUB: not implemented"; return false }

func (m *FieldMap) init() {
	m.initWithOrdering(normalFieldOrder)
}

func (m *FieldMap) initWithOrdering(ordering tagOrder) { _ = "STUB: not implemented"; return }

// Tags returns all of the Field Tags in this FieldMap.
func (m FieldMap) Tags() []Tag { _ = "STUB: not implemented"; return nil }

// Get parses out a field in this FieldMap. Returned reject may indicate the field is not present, or the field value is invalid.
func (m FieldMap) Get(parser Field) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// Has returns true if the Tag is present in this FieldMap.
func (m FieldMap) Has(tag Tag) bool { _ = "STUB: not implemented"; return false }

// GetField parses of a field with Tag tag. Returned reject may indicate the field is not present, or the field value is invalid.
func (m FieldMap) GetField(tag Tag, parser FieldValueReader) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// GetField parses of a field with Tag tag. Returned reject may indicate the field is not present, or the field value is invalid.
func (m FieldMap) getFieldNoLock(tag Tag, parser FieldValueReader) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// GetBytes is a zero-copy GetField wrapper for []bytes fields.
func (m FieldMap) GetBytes(tag Tag) ([]byte, MessageRejectError) {
	_ = "STUB: not implemented"
	return nil, *new(MessageRejectError)
}

// getBytesNoLock is a lock free zero-copy GetField wrapper for []bytes fields.
func (m FieldMap) getBytesNoLock(tag Tag) ([]byte, MessageRejectError) {
	_ = "STUB: not implemented"
	return nil, *new(MessageRejectError)
}

// GetBool is a GetField wrapper for bool fields.
func (m FieldMap) GetBool(tag Tag) (bool, MessageRejectError) {
	_ = "STUB: not implemented"
	return false, *new(MessageRejectError)
}

// GetInt is a GetField wrapper for int fields.
func (m FieldMap) GetInt(tag Tag) (int, MessageRejectError) {
	_ = "STUB: not implemented"
	return 0, *new(MessageRejectError)
}

// GetInt is a lock free GetField wrapper for int fields.
func (m FieldMap) getIntNoLock(tag Tag) (int, MessageRejectError) {
	_ = "STUB: not implemented"
	return 0, *new(MessageRejectError)
}

// GetTime is a GetField wrapper for utc timestamp fields.
func (m FieldMap) GetTime(tag Tag) (t time.Time, err MessageRejectError) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(MessageRejectError)
}

// GetString is a GetField wrapper for string fields.
func (m FieldMap) GetString(tag Tag) (string, MessageRejectError) {
	_ = "STUB: not implemented"
	return "", *new(MessageRejectError)
}

// GetString is a GetField wrapper for string fields.
func (m FieldMap) getStringNoLock(tag Tag) (string, MessageRejectError) {
	_ = "STUB: not implemented"
	return "", *new(MessageRejectError)
}

// GetGroup is a Get function specific to Group Fields.
func (m FieldMap) GetGroup(parser FieldGroupReader) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// SetField sets the field with Tag tag.
func (m *FieldMap) SetField(tag Tag, field FieldValueWriter) *FieldMap {
	_ = "STUB: not implemented"
	return nil
}

// SetBytes sets bytes.
func (m *FieldMap) SetBytes(tag Tag, value []byte) *FieldMap { _ = "STUB: not implemented"; return nil }

// SetBool is a SetField wrapper for bool fields.
func (m *FieldMap) SetBool(tag Tag, value bool) *FieldMap { _ = "STUB: not implemented"; return nil }

// SetInt is a SetField wrapper for int fields.
func (m *FieldMap) SetInt(tag Tag, value int) *FieldMap { _ = "STUB: not implemented"; return nil }

// SetString is a SetField wrapper for string fields.
func (m *FieldMap) SetString(tag Tag, value string) *FieldMap {
	_ = "STUB: not implemented"
	return nil
}

// Remove removes a tag from field map.
func (m *FieldMap) Remove(tag Tag) { _ = "STUB: not implemented"; return }

// Clear purges all fields from field map.
func (m *FieldMap) Clear() { _ = "STUB: not implemented"; return }

func (m *FieldMap) clearNoLock() { _ = "STUB: not implemented"; return }

// CopyInto overwrites the given FieldMap with this one.
func (m *FieldMap) CopyInto(to *FieldMap) { _ = "STUB: not implemented"; return }

func (m *FieldMap) add(f field) { _ = "STUB: not implemented"; return }

func (m *FieldMap) getOrCreate(tag Tag) field { _ = "STUB: not implemented"; return *new(field) }

// Set is a setter for fields.
func (m *FieldMap) Set(field FieldWriter) *FieldMap { _ = "STUB: not implemented"; return nil }

// SetGroup is a setter specific to group fields.
func (m *FieldMap) SetGroup(field FieldGroupWriter) *FieldMap {
	_ = "STUB: not implemented"
	return nil
}

func (m *FieldMap) sortedTags() []Tag { _ = "STUB: not implemented"; return nil }

func (m FieldMap) write(buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (m FieldMap) total() int { _ = "STUB: not implemented"; return 0 }

// Tag does not contribute to total.

func (m FieldMap) length() int { _ = "STUB: not implemented"; return 0 }

// Tags do not contribute to length.
