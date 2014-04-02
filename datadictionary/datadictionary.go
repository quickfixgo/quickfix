//Package datadictionary provides support for parsing and organizing FIX Data Dictionaries
package datadictionary

import (
	"encoding/xml"
	"github.com/cbusbey/quickfixgo/tag"
	"os"
)

type DataDictionary struct {
	FieldTypeByTag  map[tag.Tag]*FieldType
	FieldTypeByName map[string]*FieldType
	Messages        map[string]*Message
	Components      map[string]*Component
	Header          *Message
	Trailer         *Message
}

type Component struct {
	Fields []FieldInterface
}

type TagSet map[tag.Tag]struct{}

func (t TagSet) Add(tag tag.Tag) {
	t[tag] = struct{}{}
}

type FieldInterface interface {
	Tag() tag.Tag
	Required() bool
	MemberTags() []tag.Tag
	RequiredMemberTags() []tag.Tag
}

type Field struct {
	FieldType *FieldType
	required  bool
}

func (f Field) Tag() tag.Tag {
	return f.FieldType.Tag
}

func (f Field) Required() bool {
	return f.required
}

func (f Field) MemberTags() []tag.Tag {
	return []tag.Tag{}
}
func (f Field) RequiredMemberTags() []tag.Tag {
	return []tag.Tag{}
}

type GroupField struct {
	*FieldType
	required bool

	Fields []FieldInterface
}

func (f GroupField) Tag() tag.Tag {
	return f.FieldType.Tag
}

func (f GroupField) Required() bool {
	return f.required
}

func (f GroupField) MemberTags() []tag.Tag {
	tags := make([]tag.Tag, len(f.Fields))

	for _, f := range f.Fields {
		tags = append(tags, f.Tag())
		for _, t := range f.MemberTags() {
			tags = append(tags, t)
		}
	}

	return tags
}

func (f GroupField) RequiredMemberTags() []tag.Tag {
	tags := make([]tag.Tag, 0)

	for _, f := range f.Fields {
		if !f.Required() {
			continue
		}

		tags = append(tags, f.Tag())
		for _, t := range f.RequiredMemberTags() {
			tags = append(tags, t)
		}
	}

	return tags

	return []tag.Tag{}
}

type FieldType struct {
	Name  string
	Tag   tag.Tag
	Type  string
	Enums map[string]Enum
}

type Enum struct {
	Value       string
	Description string
}

type Message struct {
	Fields []FieldInterface

	RequiredTags TagSet
	Tags         TagSet
}

type Header Message
type Trailer Message

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

	return build(doc)
}
