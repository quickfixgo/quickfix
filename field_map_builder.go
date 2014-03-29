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

type FieldMapBuilder struct {
	fields map[tag.Tag]*fieldBytes
	fieldOrder
}

type fieldSort struct {
	tags    []tag.Tag
	compare fieldOrder
}

func (t fieldSort) Len() int           { return len(t.tags) }
func (t fieldSort) Swap(i, j int)      { t.tags[i], t.tags[j] = t.tags[j], t.tags[i] }
func (t fieldSort) Less(i, j int) bool { return t.compare(t.tags[i], t.tags[j]) }

func (fieldMap *FieldMapBuilder) init(ordering fieldOrder) {
	fieldMap.fields = make(map[tag.Tag]*fieldBytes)
	fieldMap.fieldOrder = ordering
}

func (b FieldMapBuilder) SetField(tag tag.Tag, field FieldValue) {
	b.fields[tag] = newFieldBytes(tag, field.Write())
}

func (b FieldMapBuilder) Set(field Field) {
	b.fields[field.Tag()] = newFieldBytes(field.Tag(), field.Write())
}

func (b FieldMapBuilder) Get(parser Field) error {
	field, ok := b.fields[parser.Tag()]

	if !ok {
		return FieldNotFoundError{parser.Tag()}
	}

	return parser.Read(field.Value)
}

func (b FieldMapBuilder) GetField(tag tag.Tag, field FieldValue) error {
	if f, ok := b.fields[tag]; ok {
		return field.Read(f.Value)
	} else {
		return FieldNotFoundError{tag}
	}
}

func (b FieldMapBuilder) sortedTags() []tag.Tag {
	sortedTags := make([]tag.Tag, len(b.fields))
	for tag := range b.fields {
		sortedTags = append(sortedTags, tag)
	}

	sort.Sort(fieldSort{sortedTags, b.fieldOrder})
	return sortedTags
}

func (b FieldMapBuilder) Write(buffer *bytes.Buffer) {
	tags := b.sortedTags()

	for _, tag := range tags {
		if field, ok := b.fields[tag]; ok {
			buffer.Write(field.Data)
		}
	}
}

func (b FieldMapBuilder) Remove(tag tag.Tag) {
	delete(b.fields, tag)
}

func (b FieldMapBuilder) total() int {
	total := 0
	for t, field := range b.fields {
		switch t {
		case tag.CheckSum: //tag does not contribute to total
		default:
			total += field.Total()
		}
	}

	return total
}

func (b FieldMapBuilder) length() int {
	length := 0
	for t := range b.fields {
		switch t {
		case tag.BeginString, tag.BodyLength, tag.CheckSum: //tags do not contribute to length
		default:
			length += b.fields[t].Length()
		}
	}

	return length
}
