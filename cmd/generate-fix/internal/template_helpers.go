package internal

import (
	"github.com/quickfixgo/quickfix/datadictionary"
)

func isDecimalType(quickfixType string) bool { _ = "STUB: not implemented"; return false }

func checkIfDecimalImportRequiredForFields(fTypes []*datadictionary.FieldType) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkIfTimeImportRequiredForFields(fTypes []*datadictionary.FieldType) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkFieldDecimalRequired(f *datadictionary.FieldDef) (required bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkFieldTimeRequired(f *datadictionary.FieldDef) (required bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func collectStandardImports(m *datadictionary.MessageDef) (imports []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectExtraImports(m *datadictionary.MessageDef) (imports []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkIfEnumImportRequired(m *datadictionary.MessageDef) (required bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkFieldEnumRequired(f *datadictionary.FieldDef) (required bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func quickfixValueType(quickfixType string) (goType string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func quickfixType(field *datadictionary.FieldType) (quickfixType string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func requiredFields(m *datadictionary.MessageDef) (required []*datadictionary.FieldDef) {
	_ = "STUB: not implemented"
	return nil
}

func beginString(spec *datadictionary.DataDictionary) string { _ = "STUB: not implemented"; return "" }

func routerBeginString(spec *datadictionary.DataDictionary) (routerBeginString string) {
	_ = "STUB: not implemented"
	return ""
}

// ApplVerID enums.
