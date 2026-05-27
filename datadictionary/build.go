package datadictionary

type builder struct {
	doc             *XMLDoc
	dict            *DataDictionary
	componentByName map[string]*XMLComponent
}

func (b *builder) build(doc *XMLDoc) (*DataDictionary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b builder) findOrBuildComponentType(xmlMember *XMLComponentMember) (*ComponentType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b builder) buildComponentType(xmlComponent *XMLComponent) (*ComponentType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b builder) buildComponents() error { _ = "STUB: not implemented"; return nil }

func (b builder) buildMessageDefs() error { _ = "STUB: not implemented"; return nil }

func (b builder) buildMessageDef(xmlMessage *XMLComponent) (*MessageDef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b builder) buildGroupFieldDef(xmlField *XMLComponentMember, groupFieldType *FieldType) (*FieldDef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b builder) buildFieldDef(xmlField *XMLComponentMember) (*FieldDef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b builder) buildFieldTypes() { _ = "STUB: not implemented"; return }

func buildFieldType(xmlField *XMLField) *FieldType { _ = "STUB: not implemented"; return nil }

func newUnknownComponent(name string) error { _ = "STUB: not implemented"; return nil }

func newUnknownField(name string) error { _ = "STUB: not implemented"; return nil }
