package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

//Collection of fix fields that make up a fix message
type FieldMap struct {
	fields map[tag.Tag]*fieldBytes
	tags   []tag.Tag
}

func (fieldMap *FieldMap) init() {
	fieldMap.fields = make(map[tag.Tag]*fieldBytes)
	fieldMap.tags = make([]tag.Tag, 0)
}

func (fieldMap *FieldMap) append(tag tag.Tag, field *fieldBytes) {
	fieldMap.tags = append(fieldMap.tags, tag)
	fieldMap.fields[tag] = field
}

func (m FieldMap) Tags() []tag.Tag {
	tagsCopy := make([]tag.Tag, len(m.tags))
	copy(tagsCopy, m.tags)

	return tagsCopy
}

func (m FieldMap) Get(parser Field) error {
	field, ok := m.fields[parser.Tag()]

	if !ok {
		return FieldNotFoundError{parser.Tag()}
	}

	return parser.Read(field.Value)
}

func (m FieldMap) GetField(tag tag.Tag, field FieldValue) error {
	if f, ok := m.fields[tag]; ok {
		return field.Read(f.Value)
	} else {
		return FieldNotFoundError{tag}
	}
}

func (m FieldMap) length() int {
	length := 0
	for t, field := range m.fields {
		switch t {
		case tag.BeginString, tag.BodyLength, tag.CheckSum: //tags do not contribute to length
		default:
			length += field.Length()
		}
	}

	return length
}
