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
	"github.com/quickfixgo/quickfix/datadictionary"
)

const (
	UserDefinedTagMin int = 5000
)

// Validator validates a FIX message.
type Validator interface {
	Validate(*Message) MessageRejectError
}

// ValidatorSettings describe validation behavior.
type ValidatorSettings struct {
	CheckFieldsOutOfOrder     bool
	RejectInvalidMessage      bool
	AllowUnknownMessageFields bool
	CheckUserDefinedFields    bool
	CheckFieldsHaveValues     bool
}

// Default configuration for message validation.
// See http://www.quickfixengine.org/quickfix/doc/html/configuration.html.
var defaultValidatorSettings = ValidatorSettings{
	CheckFieldsOutOfOrder:     true,
	CheckFieldsHaveValues:     true,
	RejectInvalidMessage:      true,
	AllowUnknownMessageFields: false,
	CheckUserDefinedFields:    true,
}

type fixValidator struct {
	dataDictionary *datadictionary.DataDictionary
	settings       ValidatorSettings
}

type fixtValidator struct {
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary
	settings                ValidatorSettings
}

// NewValidator creates a FIX message validator from the given data dictionaries.
func NewValidator(settings ValidatorSettings, appDataDictionary, transportDataDictionary *datadictionary.DataDictionary) Validator {
	_ = "STUB: not implemented"
	return *new(Validator)
}

// Validate tests the message against the provided data dictionary.
func (v *fixValidator) Validate(msg *Message) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// Validate tests the message against the provided transport and app data dictionaries.
// If the message is an admin message, it will be validated against the transport data dictionary.
func (v *fixtValidator) Validate(msg *Message) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func validateFIX(d *datadictionary.DataDictionary, settings ValidatorSettings, msgType string, msg *Message) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func validateFIXT(transportDD, appDD *datadictionary.DataDictionary, settings ValidatorSettings, msgType string, msg *Message) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func validateMsgType(d *datadictionary.DataDictionary, msgType string, _ *Message) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func validateWalk(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, settings ValidatorSettings, msgType string, msg *Message) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

// is body

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []TagValue) ([]TagValue, MessageRejectError) {
	_ = "STUB: not implemented"
	return nil, *new(MessageRejectError)
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []TagValue) ([]TagValue, MessageRejectError) {
	_ = "STUB: not implemented"
	return nil, *new(MessageRejectError)
}

// Start of repeating group.

// Group complete.

func validateFieldContent(msg *Message, checkFieldsHaveValues, checkFieldsOutOfOrder bool) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func validateRequired(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message *Message) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func validateRequiredFieldMap(_ *Message, requiredTags map[int]struct{}, fieldMap FieldMap) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func validateFields(transportDD *datadictionary.DataDictionary,
	appDD *datadictionary.DataDictionary,
	settings ValidatorSettings,
	msgType string,
	message *Message,
) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}

func getFieldType(d *datadictionary.DataDictionary, field int) (*datadictionary.FieldType, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func checkFieldNotDefined(settings ValidatorSettings, field Tag) bool {
	_ = "STUB: not implemented"
	return false
}

func validateField(d *datadictionary.DataDictionary,
	settings ValidatorSettings,
	_ datadictionary.TagSet,
	field TagValue,
) MessageRejectError {
	_ = "STUB: not implemented"
	return *new(MessageRejectError)
}
