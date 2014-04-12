package datadictionary

import (
	"fmt"
	"github.com/quickfixgo/quickfix/fix"
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

func (b builder) findOrBuildComponent(xmlMember *XMLComponentMember) (*Component, error) {
	if comp, preBuilt := b.dict.Components[xmlMember.Name]; preBuilt {
		return comp, nil
	}

	var xmlComp *XMLComponent
	var ok bool

	if xmlComp, ok = b.componentByName[xmlMember.Name]; !ok {
		return nil, newUnknownComponent(xmlMember.Name)
	}

	var comp *Component
	var err error

	if comp, err = b.buildComponent(xmlComp); err != nil {
		return nil, err
	}

	b.dict.Components[xmlMember.Name] = comp

	return comp, nil
}

func (b builder) buildComponent(xmlComponent *XMLComponent) (*Component, error) {
	c := &Component{Fields: make([]*FieldDef, 0)}

	for _, member := range xmlComponent.Members {
		if member.XMLName.Local == "component" {
			var childComponent *Component
			var err error
			if childComponent, err = b.findOrBuildComponent(member); err != nil {
				return nil, err
			}

			for _, childField := range childComponent.Fields {
				c.Fields = append(c.Fields, childField)
			}
		} else {
			var field *FieldDef
			var err error
			if field, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}
			c.Fields = append(c.Fields, field)
		}
	}

	return c, nil
}

func (b builder) buildComponents() error {
	b.dict.Components = make(map[string]*Component)

	for _, c := range b.doc.Components {
		if _, allReadyBuilt := b.dict.Components[c.Name]; !allReadyBuilt {
			var builtComponent *Component
			var err error
			if builtComponent, err = b.buildComponent(c); err != nil {
				return err
			}
			b.dict.Components[c.Name] = builtComponent
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
	m.Fields = make(map[fix.Tag]*FieldDef)
	m.RequiredTags = make(TagSet)
	m.Tags = make(TagSet)

	for _, member := range xmlMessage.Members {
		if member.XMLName.Local == "component" {
			var ok bool
			var comp *Component
			if comp, ok = b.dict.Components[member.Name]; !ok {
				return nil, newUnknownComponent(member.Name)
			}
			for _, f := range comp.Fields {
				m.Fields[f.Tag] = f
			}
		} else {
			var field *FieldDef
			var err error
			if field, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}
			m.Fields[field.Tag] = field
		}
	}

	for _, f := range m.Fields {
		m.Tags.Add(f.Tag)
		for _, t := range f.childTags() {
			m.Tags.Add(t)
		}

		if f.Required {
			m.RequiredTags.Add(f.Tag)
			for _, t := range f.requiredChildTags() {
				m.RequiredTags.Add(t)
			}
		}
	}

	return m, nil
}

func (b builder) buildGroupFieldDef(xmlField *XMLComponentMember, groupFieldType *FieldType) (*FieldDef, error) {
	fields := make([]*FieldDef, 0, len(xmlField.Members))

	for _, member := range xmlField.Members {
		if member.XMLName.Local == "component" {
			var component *Component
			var err error

			if component, err = b.findOrBuildComponent(member); err != nil {
				return nil, err
			}

			for _, f := range component.Fields {
				fields = append(fields, f)
			}
		} else {
			var f *FieldDef
			var err error
			if f, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}
			fields = append(fields, f)
		}
	}

	return &FieldDef{FieldType: groupFieldType, Required: (xmlField.Required == "Y"), ChildFields: fields}, nil
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

	return &FieldDef{FieldType: fieldType, Required: (xmlField.Required == "Y"), ChildFields: make([]*FieldDef, 0)}, nil
}

func (b builder) buildFieldTypes() {
	b.dict.FieldTypeByTag = make(map[fix.Tag]*FieldType)
	b.dict.FieldTypeByName = make(map[string]*FieldType)
	for _, f := range b.doc.Fields {
		field := buildFieldType(f)
		b.dict.FieldTypeByTag[field.Tag] = field
		b.dict.FieldTypeByName[field.Name] = field
	}
}

func buildFieldType(xmlField *XMLField) *FieldType {
	field := FieldType{Name: xmlField.Name, Tag: fix.Tag(xmlField.Number), Type: xmlField.Type}

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
