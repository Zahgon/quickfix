package internal

import (
	"github.com/quickfixgo/quickfix/datadictionary"
)

type fieldTypeMap map[string]*datadictionary.FieldType

var (
	globalFieldTypesLookup fieldTypeMap
	GlobalFieldTypes       []*datadictionary.FieldType
)

// Sort fieldtypes by name.
type byFieldName []*datadictionary.FieldType

func (n byFieldName) Len() int           { _ = "STUB: not implemented"; return 0 }
func (n byFieldName) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (n byFieldName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func getGlobalFieldType(f *datadictionary.FieldDef) (t *datadictionary.FieldType, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BuildGlobalFieldTypes(specs []*datadictionary.DataDictionary) {
	_ = "STUB: not implemented"
	return
}

// Merge old enums with new.

// Verify an existing enum doesn't have the same description. Keep newer enum.
