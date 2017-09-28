package datadictionary

import (
	"errors"
	"fmt"
	"strconv"
)

type builder struct {
	doc             *XMLDoc
	dict            *DataDictionary
	componentByName map[string]*XMLComponent
}

func (b *builder) build(doc *XMLDoc) (*DataDictionary, error) {
	if doc.Type != "FIX" && doc.Type != "FIXT" {
		return nil, errors.New("type attribute must be FIX or FIXT")
	}

	b.doc = doc
	b.dict = &DataDictionary{FIXType: doc.Type, ServicePack: doc.ServicePack}

	var err error
	if b.dict.Major, err = strconv.Atoi(doc.Major); err != nil {
		return nil, errors.New("major attribute not valid on <fix>")
	}

	if b.dict.Minor, err = strconv.Atoi(doc.Minor); err != nil {
		return nil, errors.New("minor attribute not valid on <fix>")
	}

	b.componentByName = make(map[string]*XMLComponent)
	for _, c := range doc.Components {
		b.componentByName[c.Name] = c
	}

	b.buildFieldTypes()

	if err := b.buildComponents(); err != nil {
		return nil, err
	}

	if err := b.buildMessageDefs(); err != nil {
		return nil, err
	}

	if b.doc.Header != nil {
		if b.dict.Header, err = b.buildMessageDef(b.doc.Header); err != nil {
			return nil, err
		}
	}

	if b.doc.Trailer != nil {
		if b.dict.Trailer, err = b.buildMessageDef(b.doc.Trailer); err != nil {
			return nil, err
		}
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
	var parts []MessagePart

	for _, member := range xmlComponent.Members {
		if member.isComponent() {
			var childComponentType *ComponentType
			var err error
			if childComponentType, err = b.findOrBuildComponentType(member); err != nil {
				return nil, err
			}

			childComponent := Component{
				ComponentType: childComponentType,
				required:      member.isRequired(),
			}

			parts = append(parts, childComponent)
		} else {
			var field *FieldDef
			var err error
			if field, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}

			parts = append(parts, field)
		}
	}

	return NewComponentType(xmlComponent.Name, parts), nil
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
	var parts []MessagePart
	for _, member := range xmlMessage.Members {
		if member.isComponent() {
			var ok bool
			var comp *ComponentType
			if comp, ok = b.dict.ComponentTypes[member.Name]; !ok {
				return nil, newUnknownComponent(member.Name)
			}

			parts = append(parts, Component{ComponentType: comp, required: member.isRequired()})
		} else {
			var field *FieldDef
			var err error
			if field, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}
			parts = append(parts, field)
		}
	}

	return NewMessageDef(xmlMessage.Name, xmlMessage.MsgType, parts), nil
}

func (b builder) buildGroupFieldDef(xmlField *XMLComponentMember, groupFieldType *FieldType) (*FieldDef, error) {
	var parts []MessagePart

	for _, member := range xmlField.Members {
		if member.isComponent() {
			var err error
			var compType *ComponentType
			if compType, err = b.findOrBuildComponentType(member); err != nil {
				return nil, err
			}

			comp := Component{
				ComponentType: compType,
				required:      member.isRequired(),
			}

			parts = append(parts, comp)
		} else {
			var f *FieldDef
			var err error
			if f, err = b.buildFieldDef(member); err != nil {
				return nil, err
			}
			parts = append(parts, f)
		}
	}

	return NewGroupFieldDef(groupFieldType, xmlField.isRequired(), parts), nil
}

func (b builder) buildFieldDef(xmlField *XMLComponentMember) (*FieldDef, error) {
	var fieldType *FieldType
	var ok bool

	if fieldType, ok = b.dict.FieldTypeByName[xmlField.Name]; !ok {
		return nil, newUnknownField(xmlField.Name)
	}

	if xmlField.isGroup() {
		f, err := b.buildGroupFieldDef(xmlField, fieldType)
		return f, err
	}

	return NewFieldDef(fieldType, xmlField.isRequired()), nil
}

func (b builder) buildFieldTypes() {
	b.dict.FieldTypeByTag = make(map[int]*FieldType)
	b.dict.FieldTypeByName = make(map[string]*FieldType)
	for _, f := range b.doc.Fields {
		field := buildFieldType(f)
		b.dict.FieldTypeByTag[field.Tag()] = field
		b.dict.FieldTypeByName[field.Name()] = field
	}
}

func buildFieldType(xmlField *XMLField) *FieldType {
	field := NewFieldType(xmlField.Name, xmlField.Number, xmlField.Type)

	if len(xmlField.Values) > 0 {
		field.Enums = make(map[string]Enum)

		for _, enum := range xmlField.Values {
			field.Enums[enum.Enum] = Enum{Value: enum.Enum, Description: enum.Description}
		}
	}

	return field
}

func newUnknownComponent(name string) error {
	return fmt.Errorf("unknown component %v", name)
}

func newUnknownField(name string) error {
	return fmt.Errorf("unknown field %v", name)
}
