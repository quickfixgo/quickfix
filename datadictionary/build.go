package datadictionary

import (
	"fmt"
)

type builder struct {
	doc             *XMLDoc
	dict            *DataDictionary
	componentByName map[string]*XMLComponent
}

func (b *builder) build(doc *XMLDoc) (*DataDictionary, error) {
	b.doc = doc

	b.componentByName = make(map[string]*XMLComponent)
	for _, c := range doc.Components {
		b.componentByName[c.Name] = c
	}

	b.dict = &DataDictionary{FIXType: doc.Type, Major: doc.Major, Minor: doc.Minor, ServicePack: doc.ServicePack}
	b.buildFieldTypes()

	if err := b.buildComponents(); err != nil {
		return nil, err
	}

	if err := b.buildMessageDefs(); err != nil {
		return nil, err
	}

	var err error
	if b.dict.Header, err = b.buildMessageDef(b.doc.Header); err != nil {
		return nil, err
	}
	if b.dict.Trailer, err = b.buildMessageDef(b.doc.Trailer); err != nil {
		return nil, err
	}

	return b.dict, nil
}

func (b builder) findOrBuildComponentType(xmlMember *XMLComponentMember) (*ComponentType, error) {
	if comp, preBuilt := b.dict.ComponentTypes[xmlMember.Name]; preBuilt {
		return comp, nil
	}

	var xmlComp *XMLComponent
	var ok bool

	if xmlComp, ok = b.componentByName[xmlMember.Name]; !ok {
		return nil, newUnknownComponent(xmlMember.Name)
	}

	var comp *ComponentType
	var err error

	if comp, err = b.buildComponentType(xmlComp); err != nil {
		return nil, err
	}

	b.dict.ComponentTypes[xmlMember.Name] = comp

	return comp, nil
}

func (b builder) buildComponentType(xmlComponent *XMLComponent) (*ComponentType, error) {
	c := &ComponentType{name: xmlComponent.Name}

	for _, member := range xmlComponent.Members {
		if member.XMLName.Local == "component" {
			var childComponentType *ComponentType
			var err error
			if childComponentType, err = b.findOrBuildComponentType(member); err != nil {
				return nil, err
			}

			childComponent := Component{
				ComponentType: childComponentType,
				required:      (member.Required == "Y"),
			}

			c.Parts = append(c.Parts, childComponent)
			c.Fields = append(c.Fields, childComponentType.Fields...)
		} else {
			var field *FieldDef
			var err error
			if field, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}

			c.Fields = append(c.Fields, field)
			c.Parts = append(c.Parts, field)
		}
	}

	return c, nil
}

func (b builder) buildComponents() error {
	b.dict.ComponentTypes = make(map[string]*ComponentType)

	for _, c := range b.doc.Components {
		if _, allReadyBuilt := b.dict.ComponentTypes[c.Name]; !allReadyBuilt {
			var builtComponent *ComponentType
			var err error
			if builtComponent, err = b.buildComponentType(c); err != nil {
				return err
			}
			b.dict.ComponentTypes[c.Name] = builtComponent
		}
	}

	return nil
}

func (b builder) buildMessageDefs() error {
	b.dict.Messages = make(map[string]*MessageDef)

	var err error
	for _, m := range b.doc.Messages {
		if b.dict.Messages[m.MsgType], err = b.buildMessageDef(m); err != nil {
			break
		}
	}

	return err
}

func (b builder) buildMessageDef(xmlMessage *XMLComponent) (*MessageDef, error) {
	m := &MessageDef{Name: xmlMessage.Name, MsgType: xmlMessage.MsgType}
	m.Fields = make(map[int]*FieldDef)
	m.RequiredTags = make(TagSet)
	m.Tags = make(TagSet)

	for _, member := range xmlMessage.Members {
		if member.XMLName.Local == "component" {
			var ok bool
			var comp *ComponentType
			if comp, ok = b.dict.ComponentTypes[member.Name]; !ok {
				return nil, newUnknownComponent(member.Name)
			}

			for _, f := range comp.Fields {
				m.Fields[f.Tag] = f
			}

			m.Parts = append(m.Parts, Component{ComponentType: comp, required: (member.Required == "Y")})
		} else {
			var field *FieldDef
			var err error
			if field, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}
			m.Fields[field.Tag] = field
			m.Parts = append(m.Parts, field)
		}
	}

	for _, f := range m.Fields {
		m.Tags.Add(f.Tag)
		for _, t := range f.childTags() {
			m.Tags.Add(t)
		}

		if f.Required() {
			m.RequiredTags.Add(f.Tag)
			for _, t := range f.requiredChildTags() {
				m.RequiredTags.Add(t)
			}
		}
	}

	return m, nil
}

func (b builder) buildGroupFieldDef(xmlField *XMLComponentMember, groupFieldType *FieldType) (*FieldDef, error) {
	var parts []MessagePart
	var fields []*FieldDef

	for _, member := range xmlField.Members {
		if member.XMLName.Local == "component" {
			var err error
			var comp *ComponentType
			if comp, err = b.findOrBuildComponentType(member); err != nil {
				return nil, err
			}

			parts = append(parts, Component{ComponentType: comp, required: (member.Required == "Y")})
			//FIXME: set fields
			for _, f := range comp.Parts {
				switch field := f.(type) {
				case *FieldDef:
					fields = append(fields, field)
				}
			}

		} else {
			var f *FieldDef
			var err error
			if f, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}
			parts = append(parts, f)
			fields = append(fields, f)
		}
	}

	return &FieldDef{FieldType: groupFieldType, required: (xmlField.Required == "Y"), Parts: parts, ChildFields: fields}, nil
}

func (b builder) buildFieldDef(xmlField *XMLComponentMember) (*FieldDef, error) {
	var fieldType *FieldType
	var ok bool

	if fieldType, ok = b.dict.FieldTypeByName[xmlField.Name]; !ok {
		return nil, newUnknownField(xmlField.Name)
	}

	if xmlField.XMLName.Local == "group" {
		f, err := b.buildGroupFieldDef(xmlField, fieldType)
		return f, err
	}

	return &FieldDef{FieldType: fieldType, required: (xmlField.Required == "Y")}, nil
}

func (b builder) buildFieldTypes() {
	b.dict.FieldTypeByTag = make(map[int]*FieldType)
	b.dict.FieldTypeByName = make(map[string]*FieldType)
	for _, f := range b.doc.Fields {
		field := buildFieldType(f)
		b.dict.FieldTypeByTag[field.Tag] = field
		b.dict.FieldTypeByName[field.Name()] = field
	}
}

func buildFieldType(xmlField *XMLField) *FieldType {
	field := FieldType{name: xmlField.Name, Tag: xmlField.Number, Type: xmlField.Type}

	if len(xmlField.Values) > 0 {
		field.Enums = make(map[string]Enum)

		for _, enum := range xmlField.Values {
			field.Enums[enum.Enum] = Enum{Value: enum.Enum, Description: enum.Description}
		}
	}

	return &field
}

func newUnknownComponent(name string) error {
	return fmt.Errorf("unknown component %v", name)
}

func newUnknownField(name string) error {
	return fmt.Errorf("unknown field %v", name)
}
