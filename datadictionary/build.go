package datadictionary

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
)

func build(doc *XMLDoc) (*DataDictionary, error) {
	dict := new(DataDictionary)
	buildFieldTypes(doc.Fields, dict)

	if err := buildComponents(doc.Components, dict); err != nil {
		return nil, err
	}

	if err := buildMessages(doc.Messages, dict); err != nil {
		return nil, err
	}

	var err error
	if dict.Header, err = buildMessage(doc.Header, dict); err != nil {
		return nil, err
	}
	if dict.Trailer, err = buildMessage(doc.Trailer, dict); err != nil {
		return nil, err
	}

	return dict, nil
}

func findOrBuildComponent(xmlMember *XMLComponentMember, xmlCompByName map[string]*XMLComponent, dict *DataDictionary) (*Component, error) {
	if comp, preBuilt := dict.Components[xmlMember.Name]; preBuilt {
		return comp, nil
	}

	var xmlComp *XMLComponent
	var ok bool

	if xmlComp, ok = xmlCompByName[xmlMember.Name]; !ok {
		return nil, newUnknownComponent(xmlMember.Name)
	}

	var comp *Component
	var err error

	if comp, err = buildComponent(xmlComp, xmlCompByName, dict); err != nil {
		return nil, err
	}

	dict.Components[xmlMember.Name] = comp

	return comp, nil
}

func buildComponent(xmlComponent *XMLComponent, xmlCompByName map[string]*XMLComponent, dict *DataDictionary) (*Component, error) {
	c := &Component{Fields: make([]FieldInterface, 0)}

	for _, member := range xmlComponent.Members {
		if member.XMLName.Local == "component" {
			var childComponent *Component
			var err error
			if childComponent, err = findOrBuildComponent(member, xmlCompByName, dict); err != nil {
				return nil, err
			}

			for _, childField := range childComponent.Fields {
				c.Fields = append(c.Fields, childField)
			}
		} else {
			if field, err := buildFieldInterface(member, dict); err != nil {
				return nil, err
			} else {
				c.Fields = append(c.Fields, field)
			}
		}
	}

	return c, nil
}

func buildComponents(xmlComponents []*XMLComponent, dict *DataDictionary) error {
	dict.Components = make(map[string]*Component)

	xmlComponentsByName := make(map[string]*XMLComponent)
	for _, c := range xmlComponents {
		xmlComponentsByName[c.Name] = c
	}

	for _, c := range xmlComponents {
		if _, allReadyBuilt := dict.Components[c.Name]; !allReadyBuilt {
			if built, err := buildComponent(c, xmlComponentsByName, dict); err != nil {
				return err
			} else {
				dict.Components[c.Name] = built
			}
		}
	}

	return nil
}

func buildMessages(messages []*XMLComponent, dict *DataDictionary) error {
	dict.Messages = make(map[string]*Message)

	var err error
	for _, m := range messages {
		if dict.Messages[m.MsgType], err = buildMessage(m, dict); err != nil {
			break
		}
	}

	return err
}

func buildMessage(xmlMessage *XMLComponent, dict *DataDictionary) (*Message, error) {
	m := new(Message)
	m.Fields = make([]FieldInterface, 0)
	m.RequiredTags = make(TagSet)
	m.Tags = make(TagSet)

	for _, member := range xmlMessage.Members {
		if member.XMLName.Local == "component" {
			var ok bool
			var comp *Component
			if comp, ok = dict.Components[member.Name]; !ok {
				return nil, newUnknownComponent(member.Name)
			}
			for _, f := range comp.Fields {
				m.Fields = append(m.Fields, f)
			}
		} else {
			var field FieldInterface
			var err error
			if field, err = buildFieldInterface(member, dict); err != nil {
				return nil, err
			}
			m.Fields = append(m.Fields, field)
		}
	}

	for _, f := range m.Fields {
		m.Tags.Add(f.Tag())
		for _, t := range f.MemberTags() {
			m.Tags.Add(t)
		}

		if f.Required() {
			m.RequiredTags.Add(f.Tag())
			for _, t := range f.RequiredMemberTags() {
				m.RequiredTags.Add(t)
			}
		}
	}

	return m, nil
}

func buildGroupField(xmlField *XMLComponentMember, groupFieldType *FieldType, dict *DataDictionary) (*GroupField, error) {
	fields := make([]FieldInterface, 0, len(xmlField.Members))

	for _, member := range xmlField.Members {
		if member.XMLName.Local == "component" {
			var component *Component
			var ok bool

			if component, ok = dict.Components[member.Name]; !ok {
				return nil, newUnknownComponent(member.Name)
			}

			for _, f := range component.Fields {
				fields = append(fields, f)
			}
		} else {
			if f, err := buildFieldInterface(member, dict); err != nil {
				return nil, err
			} else {
				fields = append(fields, f)
			}
		}
	}

	return &GroupField{FieldType: groupFieldType, required: (xmlField.Required == "Y"), Fields: fields}, nil
}

func buildFieldInterface(xmlField *XMLComponentMember, dict *DataDictionary) (FieldInterface, error) {
	var fieldType *FieldType
	var ok bool

	if fieldType, ok = dict.FieldTypeByName[xmlField.Name]; !ok {
		return nil, newUnknownField(xmlField.Name)
	}

	if xmlField.XMLName.Local == "group" {
		f, err := buildGroupField(xmlField, fieldType, dict)
		return f, err
	}

	return &Field{FieldType: fieldType, required: (xmlField.Required == "Y")}, nil
}

func buildFieldTypes(xmlFields []*XMLField, dict *DataDictionary) {
	dict.FieldTypeByTag = make(map[tag.Tag]*FieldType)
	dict.FieldTypeByName = make(map[string]*FieldType)
	for _, f := range xmlFields {
		field := buildFieldType(f)
		dict.FieldTypeByTag[field.Tag] = field
		dict.FieldTypeByName[field.Name] = field
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
