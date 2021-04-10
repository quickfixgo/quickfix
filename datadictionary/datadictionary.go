//Package datadictionary provides support for parsing and organizing FIX Data Dictionaries
package datadictionary

import (
	"encoding/xml"
	"io"
	"os"
)

//DataDictionary models FIX messages, components, and fields.
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

//MessagePart can represent a Field, Repeating Group, or Component
type MessagePart interface {
	Name() string
	Required() bool
}

//messagePartWithFields is a MessagePart with multiple Fields
type messagePartWithFields interface {
	MessagePart
	Fields() []*FieldDef
	RequiredFields() []*FieldDef
}

//ComponentType is a grouping of fields.
type ComponentType struct {
	name           string
	parts          []MessagePart
	fields         []*FieldDef
	requiredFields []*FieldDef
	requiredParts  []MessagePart
}

//NewComponentType returns an initialized component type
func NewComponentType(name string, parts []MessagePart) *ComponentType {
	comp := ComponentType{
		name:  name,
		parts: parts,
	}

	for _, part := range parts {

		if part.Required() {
			comp.requiredParts = append(comp.requiredParts, part)
		}

		switch f := part.(type) {
		case messagePartWithFields:
			comp.fields = append(comp.fields, f.Fields()...)

			if f.Required() {
				comp.requiredFields = append(comp.requiredFields, f.RequiredFields()...)
			}
		case *FieldDef:
			comp.fields = append(comp.fields, f)

			if f.Required() {
				comp.requiredFields = append(comp.requiredFields, f)
			}
		}
	}

	return &comp
}

//Name returns the name of this component type
func (c ComponentType) Name() string { return c.name }

//Fields returns all fields contained in this component. Includes fields
//encapsulated in components of this component
func (c ComponentType) Fields() []*FieldDef { return c.fields }

//RequiredFields returns those fields that are required for this component
func (c ComponentType) RequiredFields() []*FieldDef { return c.requiredFields }

//RequiredParts returns those parts that are required for this component
func (c ComponentType) RequiredParts() []MessagePart { return c.requiredParts }

// Parts returns all parts in declaration order contained in this component
func (c ComponentType) Parts() []MessagePart { return c.parts }

//TagSet is set for tags.
type TagSet map[int]struct{}

//Add adds a tag to the tagset.
func (t TagSet) Add(tag int) {
	t[tag] = struct{}{}
}

//Component is a Component as it appears in a given MessageDef
type Component struct {
	*ComponentType
	required bool
}

//NewComponent returns an initialized Component instance
func NewComponent(ct *ComponentType, required bool) *Component {
	return &Component{
		ComponentType: ct,
		required:      required,
	}
}

//Required returns true if this component is required for the containing
//MessageDef
func (c Component) Required() bool { return c.required }

//Field models a field or repeating group in a message
type Field interface {
	Tag() int
}

//FieldDef models a field belonging to a message.
type FieldDef struct {
	*FieldType
	required bool

	Parts          []MessagePart
	Fields         []*FieldDef
	requiredParts  []MessagePart
	requiredFields []*FieldDef
}

//NewFieldDef returns an initialized FieldDef
func NewFieldDef(fieldType *FieldType, required bool) *FieldDef {
	return &FieldDef{
		FieldType: fieldType,
		required:  required,
	}
}

//NewGroupFieldDef returns an initialized FieldDef for a repeating group
func NewGroupFieldDef(fieldType *FieldType, required bool, parts []MessagePart) *FieldDef {
	field := FieldDef{
		FieldType: fieldType,
		required:  required,
		Parts:     parts,
	}

	for _, part := range parts {
		if part.Required() {
			field.requiredParts = append(field.requiredParts, part)
		}

		if comp, ok := part.(Component); ok {
			field.Fields = append(field.Fields, comp.Fields()...)

			if comp.required {
				field.requiredFields = append(field.requiredFields, comp.requiredFields...)
			}
		} else {
			if child, ok := part.(*FieldDef); ok {
				field.Fields = append(field.Fields, child)

				if child.required {
					field.requiredFields = append(field.requiredFields, child)
				}
			} else {
				panic("unknown part")
			}
		}
	}

	return &field
}

//Required returns true if this FieldDef is required for the containing
//MessageDef
func (f FieldDef) Required() bool { return f.required }

//IsGroup is true if the field is a repeating group.
func (f FieldDef) IsGroup() bool {
	return len(f.Fields) > 0
}

//RequiredParts returns those parts that are required for this FieldDef. IsGroup
//must return true
func (f FieldDef) RequiredParts() []MessagePart { return f.requiredParts }

//RequiredFields returns those fields that are required for this FieldDef. IsGroup
//must return true
func (f FieldDef) RequiredFields() []*FieldDef { return f.requiredFields }

func (f FieldDef) childTags() []int {
	tags := make([]int, 0, len(f.Fields))

	for _, f := range f.Fields {
		tags = append(tags, f.Tag())
		tags = append(tags, f.childTags()...)
	}

	return tags
}

//FieldType holds information relating to a field.  Includes Tag, type, and enums, if defined.
type FieldType struct {
	name  string
	tag   int
	Type  string
	Enums map[string]Enum
}

//NewFieldType returns a pointer to an initialized FieldType
func NewFieldType(name string, tag int, fixType string) *FieldType {
	return &FieldType{
		name: name,
		tag:  tag,
		Type: fixType,
	}
}

//Name returns the name for this FieldType
func (f FieldType) Name() string { return f.name }

//Tag returns the tag for this fieldType
func (f FieldType) Tag() int { return f.tag }

//Enum is a container for value and description.
type Enum struct {
	Value       string
	Description string
}

//MessageDef can apply to header, trailer, or body of a FIX Message.
type MessageDef struct {
	Name    string
	MsgType string
	Fields  map[int]*FieldDef
	//Parts are the MessageParts of contained in this MessageDef in declaration
	//order
	Parts         []MessagePart
	requiredParts []MessagePart

	RequiredTags TagSet
	Tags         TagSet
}

//RequiredParts returns those parts that are required for this Message
func (m MessageDef) RequiredParts() []MessagePart { return m.requiredParts }

//NewMessageDef returns a pointer to an initialized MessageDef
func NewMessageDef(name, msgType string, parts []MessagePart) *MessageDef {
	msg := MessageDef{
		Name:         name,
		MsgType:      msgType,
		Fields:       make(map[int]*FieldDef),
		RequiredTags: make(TagSet),
		Tags:         make(TagSet),
		Parts:        parts,
	}

	processField := func(field *FieldDef, allowRequired bool) {
		msg.Fields[field.Tag()] = field
		msg.Tags.Add(field.Tag())
		for _, t := range field.childTags() {
			msg.Tags.Add(t)
		}

		if allowRequired && field.Required() {
			msg.RequiredTags.Add(field.Tag())
		}
	}

	for _, part := range parts {
		if part.Required() {
			msg.requiredParts = append(msg.requiredParts, part)
		}

		switch pType := part.(type) {
		case messagePartWithFields:
			for _, f := range pType.Fields() {
				//field if required in component is required in message only if
				//component is required
				processField(f, pType.Required())
			}

		case *FieldDef:
			processField(pType, true)

		default:
			panic("Unknown Part")
		}
	}

	return &msg
}

//Parse loads and and build a datadictionary instance from an xml file.
func Parse(path string) (*DataDictionary, error) {
	var xmlFile *os.File
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	return ParseSrc(xmlFile)
}

//ParseSrc loads and and build a datadictionary instance from an xml source.
func ParseSrc(xmlSrc io.Reader) (*DataDictionary, error) {
	doc := new(XMLDoc)
	decoder := xml.NewDecoder(xmlSrc)
	if err := decoder.Decode(doc); err != nil {
		return nil, err
	}

	b := new(builder)
	dict, err := b.build(doc)
	if err != nil {
		return nil, err
	}

	return dict, nil
}
