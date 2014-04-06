package quickfix

import (
	"github.com/quickfixgo/quickfix/tag"
)

//FieldMap is a collection of fix fields that make up a fix message.
type FieldMap struct {
	lookup map[tag.Tag]*fieldBytes

	fields []*fieldBytes
}

func (fieldMap *FieldMap) init() {
	fieldMap.lookup = make(map[tag.Tag]*fieldBytes)
	fieldMap.fields = make([]*fieldBytes, 0)
}

func (fieldMap *FieldMap) append(field *fieldBytes) {
	fieldMap.fields = append(fieldMap.fields, field)
	fieldMap.lookup[field.Tag] = field
}

//Tags are a copy of all tags in this field map.
func (fieldMap FieldMap) Tags() []tag.Tag {
	tags := make([]tag.Tag, len(fieldMap.fields))
	for i, f := range fieldMap.fields {
		tags[i] = f.Tag
	}

	return tags
}

func (fieldMap FieldMap) Get(parser Field) error {
	field, ok := fieldMap.lookup[parser.Tag()]

	if !ok {
		return FieldNotFoundError{parser.Tag()}
	}

	return parser.Read(field.Value)
}

func (fieldMap FieldMap) GetField(tag tag.Tag, field FieldValue) error {
	if f, ok := fieldMap.lookup[tag]; ok {
		return field.Read(f.Value)
	}
	return FieldNotFoundError{tag}
}

func (fieldMap FieldMap) length() int {
	length := 0
	for i, f := range fieldMap.fields {
		switch f.Tag {
		case tag.BeginString, tag.BodyLength, tag.CheckSum: //tags do not contribute to length
		default:
			length += fieldMap.fields[i].Length()
		}
	}

	return length
}
