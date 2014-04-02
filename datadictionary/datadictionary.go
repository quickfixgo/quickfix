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
	Fields []*Field
}

type TagSet map[tag.Tag]struct{}

func (t TagSet) Add(tag tag.Tag) {
	t[tag] = struct{}{}
}

type Field struct {
	*FieldType
	Required bool
}

type FieldType struct {
	Name string
	tag.Tag
	Type  string
	Enums map[string]Enum
}

type Enum struct {
	Value       string
	Description string
}

type Message struct {
	Fields []*Field

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
