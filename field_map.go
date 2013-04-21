package quickfixgo

import (
	"bytes"
	"github.com/cbusbey/quickfixgo/tag"
	"sort"
	"time"
)

type FieldMap struct {
	fields map[tag.Tag]Field
	fieldOrder
}

type fieldSort struct {
	tags    []tag.Tag
	compare fieldOrder
}

// fieldOrder true if tag i should occur before tag j
type fieldOrder func(i, j tag.Tag) bool

// Normal FieldOrder (ascending tags)
func normalFieldOrder(i, j tag.Tag) bool { return i < j }

func (t fieldSort) Len() int           { return len(t.tags) }
func (t fieldSort) Swap(i, j int)      { t.tags[i], t.tags[j] = t.tags[j], t.tags[i] }
func (t fieldSort) Less(i, j int) bool { return t.compare(t.tags[i], t.tags[j]) }

func (fieldMap *FieldMap) init(ordering fieldOrder) {
	fieldMap.fields = make(map[tag.Tag]Field)
	fieldMap.fieldOrder = ordering
}

func (m *FieldMap) Tags() []tag.Tag {
	tags := make([]tag.Tag, 0, len(m.fields))
	for tag := range m.fields {
		tags = append(tags, tag)
	}

	return tags
}

func (m *FieldMap) SetField(field Field) {
	m.fields[field.Tag()] = field
}

func (m FieldMap) Field(tag tag.Tag) (field Field, ok bool) {
	field, ok = m.fields[tag]

	return
}

func (m *FieldMap) Remove(tag tag.Tag) {
	delete(m.fields, tag)
}

func (m FieldMap) StringValue(tag tag.Tag) (string, bool) {
	message_field, ok := m.Field(tag)
	if !ok {
		return "", false
	}

	return message_field.Value(), true
}

func (m FieldMap) IntValue(tag tag.Tag) (int, error) {
	message_field, ok := m.Field(tag)
	if !ok {
		return 0, FieldNotFoundError{tag}
	}

	switch typeField := message_field.(type) {
	case IntField:
		return typeField.IntValue(), nil
	}

	intField, err := ToIntField(message_field)
	if err != nil {
		return 0, err
	}

	return intField.IntValue(), nil
}

func (m FieldMap) UTCTimestampValue(tag tag.Tag) (time.Time, error) {
	message_field, ok := m.Field(tag)
	if !ok {
		return time.Time{}, FieldNotFoundError{tag}
	}

	switch typeField := message_field.(type) {
	case UTCTimestampField:
		return typeField.UTCTimestampValue(), nil
	}

	utcTimestampField, err := ToUTCTimestampField(message_field)
	if err != nil {
		return time.Time{}, FieldConvertError{tag, message_field.Value()}
	}

	return utcTimestampField.UTCTimestampValue(), nil
}

func (m FieldMap) BooleanValue(tag tag.Tag) (bool, error) {

	field, ok := m.Field(tag)
	if !ok {
		return false, FieldNotFoundError{tag}
	}

	switch typeField := field.(type) {
	case BooleanField:
		return typeField.BooleanValue(), nil
	}

	boolField, err := ToBooleanField(field)
	if err != nil {
		return false, FieldConvertError{tag, field.Value()}
	}

	return boolField.BooleanValue(), nil
}

func (m FieldMap) sortedTags() []tag.Tag {
	tags := m.Tags()
	sort.Sort(fieldSort{tags, m.fieldOrder})
	return tags
}

func (m *FieldMap) Length() int {
	length := 0
	for _, field := range m.fields {
		switch field.Tag() {
		case 8, 9, 10: //tags do not contribute to length
		default:
			length += field.Length()
		}
	}

	return length
}

func (m *FieldMap) Total() int {
	total := 0
	for _, field := range m.fields {
		switch field.Tag() {
		case 10: //tag does not contribute to total
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
			b.WriteString(field.String())
		}
	}
}
