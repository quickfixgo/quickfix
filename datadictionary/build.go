package datadictionary

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
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

	b.dict = new(DataDictionary)
	b.buildFieldTypes()

	if err := b.buildComponents(); err != nil {
		return nil, err
	}

	if err := b.buildMessages(); err != nil {
		return nil, err
	}

	var err error
	if b.dict.Header, err = b.buildMessage(b.doc.Header); err != nil {
		return nil, err
	}
	if b.dict.Trailer, err = b.buildMessage(b.doc.Trailer); err != nil {
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
	c := &Component{Fields: make([]Field, 0)}

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
			if field, err := b.buildField(member); err != nil {
				return nil, err
			} else {
				c.Fields = append(c.Fields, field)
			}
		}
	}

	return c, nil
}

func (b builder) buildComponents() error {
	b.dict.Components = make(map[string]*Component)

	for _, c := range b.doc.Components {
		if _, allReadyBuilt := b.dict.Components[c.Name]; !allReadyBuilt {
			if built, err := b.buildComponent(c); err != nil {
				return err
			} else {
				b.dict.Components[c.Name] = built
			}
		}
	}

	return nil
}

func (b builder) buildMessages() error {
	b.dict.Messages = make(map[string]*Message)

	var err error
	for _, m := range b.doc.Messages {
		if b.dict.Messages[m.MsgType], err = b.buildMessage(m); err != nil {
			break
		}
	}

	return err
}

func (b builder) buildMessage(xmlMessage *XMLComponent) (*Message, error) {
	m := new(Message)
	m.Fields = make(map[tag.Tag]Field)
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
				m.Fields[f.tag()] = f
			}
		} else {
			var field Field
			var err error
			if field, err = b.buildField(member); err != nil {
				return nil, err
			}
			m.Fields[field.tag()] = field
		}
	}

	for _, f := range m.Fields {
		m.Tags.Add(f.tag())
		for _, t := range f.memberTags() {
			m.Tags.Add(t)
		}

		if f.required() {
			m.RequiredTags.Add(f.tag())
			for _, t := range f.requiredMemberTags() {
				m.RequiredTags.Add(t)
			}
		}
	}

	return m, nil
}

func (b builder) buildGroupField(xmlField *XMLComponentMember, groupFieldType *FieldType) (*GroupField, error) {
	fields := make([]Field, 0, len(xmlField.Members))

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
			if f, err := b.buildField(member); err != nil {
				return nil, err
			} else {
				fields = append(fields, f)
			}
		}
	}

	return &GroupField{FieldType: groupFieldType, Required: (xmlField.Required == "Y"), Fields: fields}, nil
}

func (b builder) buildField(xmlField *XMLComponentMember) (Field, error) {
	var fieldType *FieldType
	var ok bool

	if fieldType, ok = b.dict.FieldTypeByName[xmlField.Name]; !ok {
		return nil, newUnknownField(xmlField.Name)
	}

	if xmlField.XMLName.Local == "group" {
		f, err := b.buildGroupField(xmlField, fieldType)
		return f, err
	}

	return &MessageField{FieldType: fieldType, Required: (xmlField.Required == "Y")}, nil
}

func (b builder) buildFieldTypes() {
	b.dict.FieldTypeByTag = make(map[tag.Tag]*FieldType)
	b.dict.FieldTypeByName = make(map[string]*FieldType)
	for _, f := range b.doc.Fields {
		field := buildFieldType(f)
		b.dict.FieldTypeByTag[field.Tag] = field
		b.dict.FieldTypeByName[field.Name] = field
	}
}

func buildFieldType(xmlField *XMLField) *FieldType {
	field := FieldType{Name: xmlField.Name, Tag: tag.Tag(xmlField.Number), Type: xmlField.Type}

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
