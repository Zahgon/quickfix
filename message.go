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
	"time"

	"github.com/quickfixgo/quickfix/datadictionary"
)

// Header is first section of a FIX Message.
type Header struct{ FieldMap }

// msgparser contains message parsing vars needed to parse a string into a message.
type msgParser struct {
	msg                     *Message
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary
	rawBytes                []byte
	fieldIndex              int
	parsedFieldBytes        *TagValue
	trailerBytes            []byte
	foundBody               bool
	foundTrailer            bool
}

// in the message header, the first 3 tags in the message header must be 8,9,35.
func headerFieldOrdering(i, j Tag) bool { _ = "STUB: not implemented"; return false }

// Init initializes the Header instance.
func (h *Header) Init() { _ = "STUB: not implemented"; return }

// Body is the primary application section of a FIX message.
type Body struct{ FieldMap }

// Init initializes the FIX message.
func (b *Body) Init() {
	_ = "STUB: not implemented"

	// Trailer is the last section of a FIX message.
	return
}

type Trailer struct{ FieldMap }

// In the trailer, CheckSum (tag 10) must be last.
func trailerFieldOrdering(i, j Tag) bool { _ = "STUB: not implemented"; return false }

// Init initializes the FIX message.
func (t *Trailer) Init() { _ = "STUB: not implemented"; return }

// Message is a FIX Message abstraction.
type Message struct {
	Header  Header
	Trailer Trailer
	Body    Body

	// ReceiveTime is the time that this message was read from the socket connection.
	ReceiveTime time.Time

	rawMessage *bytes.Buffer

	// Slice of Bytes corresponding to the message body.
	bodyBytes []byte

	// Field bytes as they appear in the raw message.
	fields []TagValue
}

// ToMessage returns the message itself.
func (m *Message) ToMessage() *Message {
	_ = "STUB: not implemented"

	// parseError is returned when bytes cannot be parsed as a FIX message.
	return nil
}

type parseError struct {
	OrigError string
}

func (e parseError) Error() string { _ = "STUB: not implemented"; return "" }

// NewMessage returns a newly initialized Message instance.
func NewMessage() *Message { _ = "STUB: not implemented"; return nil }

// CopyInto erases the dest messages and copies the currency message content
// into it.
func (m *Message) CopyInto(to *Message) { _ = "STUB: not implemented"; return }

// ParseMessage constructs a Message from a byte slice wrapping a FIX message.
func ParseMessage(msg *Message, rawMessage *bytes.Buffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ParseMessageWithDataDictionary constructs a Message from a byte slice wrapping a FIX message using an optional session and application DataDictionary for reference.
func ParseMessageWithDataDictionary(
	msg *Message,
	rawMessage *bytes.Buffer,
	transportDataDictionary *datadictionary.DataDictionary,
	appDataDictionary *datadictionary.DataDictionary,
) (err error) {
	_ = "STUB: not implemented"
	// Create msgparser before we go any further.
	return nil
}

// doParsing executes the message parsing process.
func doParsing(mp *msgParser) (err error) { _ = "STUB: not implemented"; return nil }

// Initialize for parsing.

// Allocate expected message fields in one chunk.

// Message must start with begin string, body length, msg type.
// Get begin string.

// Get body length.

// Get msg type.

// Start parsing.

// This will happen if there are no fields in the body

// Body length would only be larger than trailer if fields out of order.

// Tags do not contribute to length.

// parseGroup iterates through a repeating group to maintain correct order of those fields.
func parseGroup(mp *msgParser, tags []Tag) { _ = "STUB: not implemented"; return }

// Is this field a member for the group.

// Is this field a nested repeating group.

// Add the field member to the group.

// Found a header tag for some reason..

// Found the trailer at the end of the message.

// Found a body field outside the group.

// Is this a new group not inside the existing group.

// Add the current repeating group.

// Cycle again with the new group.

// Did this tag occur after a nested group and belongs to the parent group.

// Add the field member to the group.

// Continue parsing the parent group.

// Add the repeating group.

// Add the next body field.

// isNumInGroupField evaluates if this tag is the start of a repeating group.
// tags slice will contain multiple tags if the tag in question is found while processing a group already.
func isNumInGroupField(msg *Message, tags []Tag, appDataDictionary *datadictionary.DataDictionary) bool {
	_ = "STUB: not implemented"
	return false
}

// Map nested fields.

// getGroupFields gets the relevant fields for parsing a repeating group if this tag is the start of a repeating group.
// tags slice will contain multiple tags if the tag in question is found while processing a group already.
func getGroupFields(msg *Message, tags []Tag, appDataDictionary *datadictionary.DataDictionary) (fields []*datadictionary.FieldDef) {
	_ = "STUB: not implemented"
	return nil
}

// Map nested fields.

// isGroupMember evaluates if this tag belongs to a repeating group.
func isGroupMember(tag Tag, fields []*datadictionary.FieldDef) bool {
	_ = "STUB: not implemented"
	return false
}

func isHeaderField(tag Tag, dataDict *datadictionary.DataDictionary) bool {
	_ = "STUB: not implemented"
	return false
}

func isTrailerField(tag Tag, dataDict *datadictionary.DataDictionary) bool {
	_ = "STUB: not implemented"
	return false
}

// MsgType returns MsgType (tag 35) field's value.
func (m *Message) MsgType() (string, MessageRejectError) {
	_ = "STUB: not implemented"
	return "", *new(MessageRejectError)
}

func (m *Message) msgTypeNoLock() (string, MessageRejectError) {
	_ = "STUB: not implemented"
	return "", *new(MessageRejectError)
}

// IsMsgTypeOf returns true if the Header contains MsgType (tag 35) field and its value is the specified one.
func (m *Message) IsMsgTypeOf(msgType string) bool { _ = "STUB: not implemented"; return false }

// reverseRoute returns a message builder with routing header fields initialized as the reverse of this message.
func (m *Message) reverseRoute() *Message { _ = "STUB: not implemented"; return nil }

// Tags added in 4.1.

func extractSpecificField(field *TagValue, expectedTag Tag, buffer []byte) (remBuffer []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractXMLDataField(parsedFieldBytes *TagValue, buffer []byte, dataLen int) (remBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractField(parsedFieldBytes *TagValue, buffer []byte) (remBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Message) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) String() string { _ = "STUB: not implemented"; return "" }

func formatCheckSum(value int) string { _ = "STUB: not implemented"; return "" }

// Build constructs a []byte from a Message instance.
func (m *Message) build() []byte { _ = "STUB: not implemented"; return nil }

// Constructs a []byte from a Message instance, using the given bodyBytes.
// This is a workaround for the fact that we currently rely on the generated Message types to properly serialize/deserialize RepeatingGroups.
// In other words, we cannot go from bytes to a Message then back to bytes, which is exactly what we need to do in the case of a Resend.
// This func lets us pull the Message from the Store, parse it, update the Header, and then build it back into bytes using the original Body.
// Note: The only standard non-Body group is NoHops.  If that is used in the Header, this workaround may fail.
func (m *Message) buildWithBodyBytes(bodyBytes []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (m *Message) cook(bodyLen, bodyTotal int) { _ = "STUB: not implemented"; return }
