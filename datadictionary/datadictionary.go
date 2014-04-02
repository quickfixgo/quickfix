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
	Fields []Field
}

type TagSet map[tag.Tag]struct{}

func (t TagSet) Add(tag tag.Tag) {
	t[tag] = struct{}{}
}

type Field interface {
	tag() tag.Tag
	required() bool
	memberTags() []tag.Tag
	requiredMemberTags() []tag.Tag
}

type MessageField struct {
	*FieldType
	Required bool
}

func (f MessageField) tag() tag.Tag {
	return f.FieldType.Tag
}

func (f MessageField) required() bool {
	return f.Required
}

func (f MessageField) memberTags() []tag.Tag {
	return []tag.Tag{}
}
func (f MessageField) requiredMemberTags() []tag.Tag {
	return []tag.Tag{}
}

type GroupField struct {
	*FieldType
	Required bool

	Fields []Field
}

func (f GroupField) tag() tag.Tag {
	return f.FieldType.Tag
}

func (f GroupField) required() bool {
	return f.Required
}

func (f GroupField) memberTags() []tag.Tag {
	tags := make([]tag.Tag, len(f.Fields))

	for _, f := range f.Fields {
		tags = append(tags, f.tag())
		for _, t := range f.memberTags() {
			tags = append(tags, t)
		}
	}

	return tags
}

func (f GroupField) requiredMemberTags() []tag.Tag {
	tags := make([]tag.Tag, 0)

	for _, f := range f.Fields {
		if !f.required() {
			continue
		}

		tags = append(tags, f.tag())
		for _, t := range f.requiredMemberTags() {
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
	Fields map[tag.Tag]Field

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

	b := new(builder)
	var dict *DataDictionary
	if dict, err = b.build(doc); err != nil {
		return nil, err
	}

	return dict, nil
}
