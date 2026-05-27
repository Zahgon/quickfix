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

// GroupItem interface is used to construct repeating group templates.
type GroupItem interface {
	// Tag returns the tag identifying this GroupItem.
	Tag() Tag

	// Read Parameter to Read is tagValues.  For most fields, only the first tagValue will be required.
	// The length of the slice extends from the tagValue mapped to the field to be read through the
	// following fields. This can be useful for GroupItems made up of repeating groups.
	//
	// The Read function returns the remaining tagValues not processed by the GroupItem. If there was a
	// problem reading the field, an error may be returned.
	Read([]TagValue) ([]TagValue, error)

	// Clone makes a copy of this GroupItem.
	Clone() GroupItem
}

type protoGroupElement struct {
	tag Tag
}

func (t protoGroupElement) Tag() Tag { _ = "STUB: not implemented"; return *new(Tag) }
func (t protoGroupElement) Read(tv []TagValue) ([]TagValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t protoGroupElement) Clone() GroupItem {
	_ = "STUB: not implemented"

	// GroupElement returns a GroupItem made up of a single field.
	return *new(GroupItem)
}

func GroupElement(tag Tag) GroupItem { _ = "STUB: not implemented"; return *new(GroupItem) }

// GroupTemplate specifies the group item order for a RepeatingGroup.
type GroupTemplate []GroupItem

// Clone makes a copy of this GroupTemplate.
func (gt GroupTemplate) Clone() GroupTemplate {
	_ = "STUB: not implemented"
	return *new(GroupTemplate)
}

// Group is a group of fields occurring in a repeating group.
type Group struct{ FieldMap }

// RepeatingGroup is a FIX Repeating Group type.
type RepeatingGroup struct {
	tag      Tag
	template GroupTemplate
	groups   []*Group
}

// NewRepeatingGroup returns an initilized RepeatingGroup instance.
func NewRepeatingGroup(tag Tag, template GroupTemplate) *RepeatingGroup {
	_ = "STUB: not implemented"
	return nil
}

// Tag returns the Tag for this repeating Group.
func (f RepeatingGroup) Tag() Tag {
	_ = "STUB: not implemented"

	// Clone makes a copy of this RepeatingGroup (tag, template).
	return *new(Tag)
}

func (f RepeatingGroup) Clone() GroupItem { _ = "STUB: not implemented"; return *new(GroupItem) }

// Len returns the number of Groups in this RepeatingGroup.
func (f RepeatingGroup) Len() int { _ = "STUB: not implemented"; return 0 }

// Get returns the ith group in this RepeatingGroup.
func (f RepeatingGroup) Get(i int) *Group {
	_ = "STUB: not implemented"

	// Add appends a new group to the RepeatingGroup and returns the new Group.
	return nil
}

func (f *RepeatingGroup) Add() *Group { _ = "STUB: not implemented"; return nil }

// Write returns tagValues for all Items in the repeating group ordered by
// Group sequence and Group template order.
func (f RepeatingGroup) Write() []TagValue { _ = "STUB: not implemented"; return nil }

func (f RepeatingGroup) findItemInGroupTemplate(t Tag) (item GroupItem, ok bool) {
	_ = "STUB: not implemented"
	return *new(GroupItem), false
}

func (f RepeatingGroup) groupTagOrder() tagOrder { _ = "STUB: not implemented"; return *new(tagOrder) }

func (f RepeatingGroup) delimiter() Tag { _ = "STUB: not implemented"; return *new(Tag) }

func (f RepeatingGroup) isDelimiter(t Tag) bool { _ = "STUB: not implemented"; return false }

func (f *RepeatingGroup) Read(tv []TagValue) ([]TagValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
