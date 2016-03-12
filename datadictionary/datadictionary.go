//Package datadictionary provides support for parsing and organizing FIX Data Dictionaries
package datadictionary

import (
	"encoding/xml"
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

//ComponentType is a grouping of fields.
type ComponentType struct {
	name  string
	Parts []MessagePart
}

//Name returns the name of this component type
func (c ComponentType) Name() string { return c.name }

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

//Required returns true if this component is required for the containing
//MessageDef
func (c Component) Required() bool { return c.required }

//FieldDef models a field or component belonging to a message.
type FieldDef struct {
	*FieldType
	required bool

	Parts       []MessagePart
	ChildFields []*FieldDef
}

//Required returns true if this FieldDef is required for the containing
//MessageDef
func (f FieldDef) Required() bool { return f.required }

//IsGroup is true if the field is a repeating group.
func (f FieldDef) IsGroup() bool {
	return len(f.ChildFields) > 0
}

func (f FieldDef) childTags() []int {
	tags := make([]int, 0, len(f.ChildFields))

	for _, f := range f.ChildFields {
		tags = append(tags, f.Tag)
		for _, t := range f.childTags() {
			tags = append(tags, t)
		}
	}

	return tags
}

func (f FieldDef) requiredChildTags() []int {
	var tags []int

	for _, f := range f.ChildFields {
		if !f.Required() {
			continue
		}

		tags = append(tags, f.Tag)
		for _, t := range f.requiredChildTags() {
			tags = append(tags, t)
		}
	}

	return tags
}

//FieldType holds information relating to a field.  Includes Tag, type, and enums, if defined.
type FieldType struct {
	name  string
	Tag   int
	Type  string
	Enums map[string]Enum
}

//NewFieldType returns a pointer to an initialized FieldType
func NewFieldType(name string, tag int, fixType string) *FieldType {
	return &FieldType{
		name: name,
		Tag:  tag,
		Type: fixType,
	}
}

//Name returns the name for this FieldType
func (f FieldType) Name() string { return f.name }

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
	Parts []MessagePart

	RequiredTags TagSet
	Tags         TagSet
}

//Parse loads and and build a datadictionary instance from an xml file.
func Parse(path string) (*DataDictionary, error) {
	var xmlFile *os.File
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	doc := new(XMLDoc)
	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(doc); err != nil {
		return nil, err
	}

	b := new(builder)
	var dict *DataDictionary
	if dict, err = b.build(doc); err != nil {
		return nil, err
	}

	return dict, nil
}
