// Package datadictionary provides support for parsing and organizing FIX Data Dictionaries
package datadictionary

import (
	"io"
)

// DataDictionary models FIX messages, components, and fields.
type DataDictionary struct {
	FIXType         string
	Major           int
	Minor           int
	ServicePack     int
	FieldTypeByTag  map[int]*FieldType
	FieldTypeByName map[string]*FieldType
	Messages        map[string]*MessageDef
	ComponentTypes  map[string]*ComponentType
	Header          *MessageDef
	Trailer         *MessageDef
}

// MessagePart can represent a Field, Repeating Group, or Component.
type MessagePart interface {
	Name() string
	Required() bool
}

// messagePartWithFields is a MessagePart with multiple Fields.
type messagePartWithFields interface {
	MessagePart
	Fields() []*FieldDef
	RequiredFields() []*FieldDef
}

// ComponentType is a grouping of fields.
type ComponentType struct {
	name           string
	parts          []MessagePart
	fields         []*FieldDef
	requiredFields []*FieldDef
	requiredParts  []MessagePart
}

// NewComponentType returns an initialized component type.
func NewComponentType(name string, parts []MessagePart) *ComponentType {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of this component type.
func (c ComponentType) Name() string {
	_ = "STUB: not implemented"

	// Fields returns all fields contained in this component. Includes fields
	// encapsulated in components of this component.
	return ""
}

func (c ComponentType) Fields() []*FieldDef {
	_ = "STUB: not implemented"

	// RequiredFields returns those fields that are required for this component.
	return nil
}

func (c ComponentType) RequiredFields() []*FieldDef { _ = "STUB: not implemented"; return nil }

// RequiredParts returns those parts that are required for this component.
func (c ComponentType) RequiredParts() []MessagePart { _ = "STUB: not implemented"; return nil }

// Parts returns all parts in declaration order contained in this component.
func (c ComponentType) Parts() []MessagePart {
	_ = "STUB: not implemented"

	// TagSet is set for tags.
	return nil
}

type TagSet map[int]struct{}

// Add adds a tag to the tagset.
func (t TagSet) Add(tag int) { _ = "STUB: not implemented"; return }

// Component is a Component as it appears in a given MessageDef.
type Component struct {
	*ComponentType
	required bool
}

// NewComponent returns an initialized Component instance.
func NewComponent(ct *ComponentType, required bool) *Component {
	_ = "STUB: not implemented"
	return nil
}

// Required returns true if this component is required for the containing
// MessageDef.
func (c Component) Required() bool {
	_ = "STUB: not implemented"

	// Field models a field or repeating group in a message.
	return false
}

type Field interface {
	Tag() int
}

// FieldDef models a field belonging to a message.
type FieldDef struct {
	*FieldType
	required bool

	Parts          []MessagePart
	Fields         []*FieldDef
	requiredParts  []MessagePart
	requiredFields []*FieldDef
}

// NewFieldDef returns an initialized FieldDef.
func NewFieldDef(fieldType *FieldType, required bool) *FieldDef {
	_ = "STUB: not implemented"
	return nil
}

// NewGroupFieldDef returns an initialized FieldDef for a repeating group.
func NewGroupFieldDef(fieldType *FieldType, required bool, parts []MessagePart) *FieldDef {
	_ = "STUB: not implemented"
	return nil
}

// Required returns true if this FieldDef is required for the containing
// MessageDef.
func (f FieldDef) Required() bool {
	_ = "STUB: not implemented"

	// IsGroup is true if the field is a repeating group.
	return false
}

func (f FieldDef) IsGroup() bool { _ = "STUB: not implemented"; return false }

// RequiredParts returns those parts that are required for this FieldDef. IsGroup
// must return true.
func (f FieldDef) RequiredParts() []MessagePart { _ = "STUB: not implemented"; return nil }

// RequiredFields returns those fields that are required for this FieldDef. IsGroup
// must return true.
func (f FieldDef) RequiredFields() []*FieldDef { _ = "STUB: not implemented"; return nil }

func (f FieldDef) childTags() []int { _ = "STUB: not implemented"; return nil }

// FieldType holds information relating to a field.  Includes Tag, type, and enums, if defined.
type FieldType struct {
	name  string
	tag   int
	Type  string
	Enums map[string]Enum
}

// NewFieldType returns a pointer to an initialized FieldType.
func NewFieldType(name string, tag int, fixType string) *FieldType {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name for this FieldType.
func (f FieldType) Name() string {
	_ = "STUB: not implemented"

	// Tag returns the tag for this fieldType.
	return ""
}

func (f FieldType) Tag() int {
	_ = "STUB: not implemented"

	// Enum is a container for value and description.
	return 0
}

type Enum struct {
	Value       string
	Description string
}

// MessageDef can apply to header, trailer, or body of a FIX Message.
type MessageDef struct {
	Name    string
	MsgType string
	Fields  map[int]*FieldDef
	// Parts are the MessageParts of contained in this MessageDef in declaration order.
	Parts         []MessagePart
	requiredParts []MessagePart

	RequiredTags TagSet
	Tags         TagSet
}

// RequiredParts returns those parts that are required for this Message.
func (m MessageDef) RequiredParts() []MessagePart { _ = "STUB: not implemented"; return nil }

// NewMessageDef returns a pointer to an initialized MessageDef.
func NewMessageDef(name, msgType string, parts []MessagePart) *MessageDef {
	_ = "STUB: not implemented"
	return nil
}

// Field if required in component is required in message only if
// component is required.

// Parse loads and build a datadictionary instance from an xml file.
func Parse(path string) (*DataDictionary, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseSrc loads and build a datadictionary instance from an xml source.
func ParseSrc(xmlSrc io.Reader) (*DataDictionary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
