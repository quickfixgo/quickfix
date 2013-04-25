package quickfixgo

import (
	"bytes"
	"github.com/cbusbey/quickfixgo/tag"
	"sort"
)

// fieldOrder true if tag i should occur before tag j
type fieldOrder func(i, j tag.Tag) bool

// Normal FieldOrder (ascending tags)
func normalFieldOrder(i, j tag.Tag) bool { return i < j }

//Collection of fix fields that make up a fix message 
type FieldMap struct {
	fields map[tag.Tag]*field
	fieldOrder
}

type fieldSort struct {
	tags    []tag.Tag
	compare fieldOrder
}

func (t fieldSort) Len() int           { return len(t.tags) }
func (t fieldSort) Swap(i, j int)      { t.tags[i], t.tags[j] = t.tags[j], t.tags[i] }
func (t fieldSort) Less(i, j int) bool { return t.compare(t.tags[i], t.tags[j]) }

func (fieldMap *FieldMap) init(ordering fieldOrder) {
	fieldMap.fields = make(map[tag.Tag]*field)
	fieldMap.fieldOrder = ordering
}

func (m *FieldMap) Tags() []tag.Tag {
	tags := make([]tag.Tag, 0, len(m.fields))
	for tag := range m.fields {
		tags = append(tags, tag)
	}

	return tags
}

func (m *FieldMap) SetField(field FieldConverter) {
	m.fields[field.Tag()] = newField(field.Tag(), field.ConvertValueToBytes())
}

func (m FieldMap) Get(parser FieldConverter) error {
	field, ok := m.fields[parser.Tag()]

	if !ok {
		return FieldNotFoundError{parser.Tag()}
	}

	return parser.ConvertValueFromBytes(field.Value)
}

func (m *FieldMap) Remove(tag tag.Tag) {
	delete(m.fields, tag)
}

func (m FieldMap) sortedTags() []tag.Tag {
	tags := m.Tags()
	sort.Sort(fieldSort{tags, m.fieldOrder})
	return tags
}

func (m *FieldMap) length() int {
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

func (m *FieldMap) total() int {
	total := 0
	for t, field := range m.fields {
		switch t {
		case tag.CheckSum: //tag does not contribute to total
		default:
			total += field.Total()
		}
	}

	return total
}

func (m *FieldMap) Write(b *bytes.Buffer) {
	tags := m.sortedTags()

	for _, tag := range tags {
		if field, ok := m.fields[tag]; ok {
			b.Write(field.Data)
		}
	}
}
