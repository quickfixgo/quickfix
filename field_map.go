package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

//Collection of fix fields that make up a fix message
type FieldMap struct {
	lookup map[tag.Tag]*fieldBytes

	tags   []tag.Tag
	fields []*fieldBytes
}

func (fieldMap *FieldMap) init() {
	fieldMap.lookup = make(map[tag.Tag]*fieldBytes)
	fieldMap.tags = make([]tag.Tag, 0)
	fieldMap.fields = make([]*fieldBytes, 0)
}

func (fieldMap *FieldMap) append(tag tag.Tag, field *fieldBytes) {
	fieldMap.tags = append(fieldMap.tags, tag)
	fieldMap.fields = append(fieldMap.fields, field)
	fieldMap.lookup[tag] = field
}

func (m FieldMap) Tags() []tag.Tag {
	tagsCopy := make([]tag.Tag, len(m.tags))
	copy(tagsCopy, m.tags)

	return tagsCopy
}

func (m FieldMap) Get(parser Field) error {
	field, ok := m.lookup[parser.Tag()]

	if !ok {
		return FieldNotFoundError{parser.Tag()}
	}

	return parser.Read(field.Value)
}

func (m FieldMap) GetField(tag tag.Tag, field FieldValue) error {
	if f, ok := m.lookup[tag]; ok {
		return field.Read(f.Value)
	} else {
		return FieldNotFoundError{tag}
	}
}

func (m FieldMap) length() int {
	length := 0
	for i, t := range m.tags {
		switch t {
		case tag.BeginString, tag.BodyLength, tag.CheckSum: //tags do not contribute to length
		default:
			length += m.fields[i].Length()
		}
	}

	return length
}
