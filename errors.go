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
	"errors"
)

// ErrDoNotSend is a convenience error to indicate a DoNotSend in ToApp.
var ErrDoNotSend = errors.New("Do Not Send")

// rejectReason enum values.
const (
	rejectReasonInvalidTagNumber                          = 0
	rejectReasonRequiredTagMissing                        = 1
	rejectReasonTagNotDefinedForThisMessageType           = 2
	rejectReasonUnsupportedMessageType                    = 3
	rejectReasonTagSpecifiedWithoutAValue                 = 4
	rejectReasonValueIsIncorrect                          = 5
	rejectReasonConditionallyRequiredFieldMissing         = 5
	rejectReasonIncorrectDataFormatForValue               = 6
	rejectReasonCompIDProblem                             = 9
	rejectReasonSendingTimeAccuracyProblem                = 10
	rejectReasonInvalidMsgType                            = 11
	rejectReasonTagAppearsMoreThanOnce                    = 13
	rejectReasonTagSpecifiedOutOfRequiredOrder            = 14
	rejectReasonRepeatingGroupFieldsOutOfOrder            = 15
	rejectReasonIncorrectNumInGroupCountForRepeatingGroup = 16
)

// MessageRejectError is a type of error that can correlate to a message reject.
type MessageRejectError interface {
	error

	// RejectReason, tag 373 for session rejects, tag 380 for business rejects.
	RejectReason() int
	BusinessRejectRefID() string
	RefTagID() *Tag
	IsBusinessReject() bool
}

// RejectLogon indicates the application is rejecting permission to logon. Implements MessageRejectError.
type RejectLogon struct {
	Text string
}

func (e RejectLogon) Error() string {
	_ = "STUB: not implemented"

	// RefTagID implements MessageRejectError.
	return ""
}

func (RejectLogon) RefTagID() *Tag {
	_ = "STUB: not implemented"

	// RejectReason implements MessageRejectError.
	return nil
}

func (RejectLogon) RejectReason() int {
	_ = "STUB: not implemented"

	// BusinessRejectRefID implements MessageRejectError.
	return 0
}

func (RejectLogon) BusinessRejectRefID() string {
	_ = "STUB: not implemented"

	// IsBusinessReject implements MessageRejectError.
	return ""
}

func (RejectLogon) IsBusinessReject() bool { _ = "STUB: not implemented"; return false }

type messageRejectError struct {
	rejectReason        int
	text                string
	businessRejectRefID string
	refTagID            *Tag
	isBusinessReject    bool
}

func (e messageRejectError) Error() string               { _ = "STUB: not implemented"; return "" }
func (e messageRejectError) RefTagID() *Tag              { _ = "STUB: not implemented"; return nil }
func (e messageRejectError) RejectReason() int           { _ = "STUB: not implemented"; return 0 }
func (e messageRejectError) BusinessRejectRefID() string { _ = "STUB: not implemented"; return "" }
func (e messageRejectError) IsBusinessReject() bool      { _ = "STUB: not implemented"; return false }

// NewMessageRejectError returns a MessageRejectError with the given error message, reject reason, and optional reftagid.
func NewMessageRejectError(err string, rejectReason int, refTagID *Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// NewBusinessMessageRejectError returns a MessageRejectError with the given error mesage, reject reason, and optional reftagid.
// Reject is treated as a business level reject.
func NewBusinessMessageRejectError(err string, rejectReason int, refTagID *Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// NewBusinessMessageRejectErrorWithRefID returns a MessageRejectError with the given error mesage, reject reason, refID, and optional reftagid.
// Reject is treated as a business level reject.
func NewBusinessMessageRejectErrorWithRefID(err string, rejectReason int, businessRejectRefID string, refTagID *Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// IncorrectDataFormatForValue returns an error indicating a field that cannot be parsed as the type required.
func IncorrectDataFormatForValue(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// repeatingGroupFieldsOutOfOrder returns an error indicating a problem parsing repeating groups fields.
func repeatingGroupFieldsOutOfOrder(tag Tag, reason string) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// ValueIsIncorrect returns an error indicating a field with value that is not valid.
func ValueIsIncorrect(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// ConditionallyRequiredFieldMissing indicates that the requested field could not be found in the FIX message.
func ConditionallyRequiredFieldMissing(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// valueIsIncorrectNoTag returns an error indicating a field with value that is not valid.
// FIXME: to be compliant with legacy tests, for certain value issues, do not include reftag? (11c_NewSeqNoLess).
func valueIsIncorrectNoTag() MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// InvalidMessageType returns an error to indicate an invalid message type.
func InvalidMessageType() MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// UnsupportedMessageType returns an error to indicate an unhandled message.
func UnsupportedMessageType() MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// TagNotDefinedForThisMessageType returns an error for an invalid tag appearing in a message.
func TagNotDefinedForThisMessageType(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// tagAppearsMoreThanOnce return an error for multiple tags in a message not detected as a repeating group.
func tagAppearsMoreThanOnce(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// RequiredTagMissing returns a validation error when a required field cannot be found in a message.
func RequiredTagMissing(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// incorrectNumInGroupCountForRepeatingGroup returns a validation error when the num in group value for a group does not match actual group size.
func incorrectNumInGroupCountForRepeatingGroup(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// tagSpecifiedOutOfRequiredOrder returns validation error when the group order does not match the spec.
func tagSpecifiedOutOfRequiredOrder(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// TagSpecifiedWithoutAValue returns a validation error for when a field has no value.
func TagSpecifiedWithoutAValue(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// InvalidTagNumber returns a validation error for messages with invalid tags.
func InvalidTagNumber(tag Tag) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// compIDProblem creates a reject for msg where msg has invalid comp id values.
func compIDProblem() MessageRejectError { _ = "STUB: not implemented"; return *new(MessageRejectError) }

// sendingTimeAccuracyProblem creates a reject for a msg with stale or invalid sending time.
func sendingTimeAccuracyProblem() MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}
