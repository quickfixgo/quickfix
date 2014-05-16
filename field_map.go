package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	"math"
	"sort"
)

// fieldOrder true if tag i should occur before tag j
type fieldOrder func(i, j fix.Tag) bool

type fieldSort struct {
	tags    []fix.Tag
	compare fieldOrder
}

func (t fieldSort) Len() int           { return len(t.tags) }
func (t fieldSort) Swap(i, j int)      { t.tags[i], t.tags[j] = t.tags[j], t.tags[i] }
func (t fieldSort) Less(i, j int) bool { return t.compare(t.tags[i], t.tags[j]) }

//in the message header, the first 3 tags in the message header must be 8,9,35
func headerFieldOrder(i, j fix.Tag) bool {
	var ordering = func(t fix.Tag) uint32 {
		switch t {
		case tag.BeginString:
			return 1
		case tag.BodyLength:
			return 2
		case tag.MsgType:
			return 3
		}

		return math.MaxUint32
	}

	orderi := ordering(i)
	orderj := ordering(j)

	switch {
	case orderi < orderj:
		return true
	case orderi > orderj:
		return false
	}

	return i < j
}

// In the body, ascending tags
func normalFieldOrder(i, j fix.Tag) bool { return i < j }

// In the trailer, CheckSum (tag 10) must be last
func trailerFieldOrder(i, j fix.Tag) bool {
	switch {
	case i == tag.CheckSum:
		return false
	case j == tag.CheckSum:
		return true
	}

	return i < j
}

//FieldMap is a collection of fix fields that make up a fix message.
type FieldMap struct {
	fieldLookup map[fix.Tag]*fieldBytes
	fieldOrder
}

func (m *FieldMap) init(ordering fieldOrder) {
	m.fieldLookup = make(map[fix.Tag]*fieldBytes)
	m.fieldOrder = ordering
}

func (m *FieldMap) append(field *fieldBytes) {
	m.fieldLookup[field.Tag] = field
}

//Tags are a copy of all tags in this field map.
func (m FieldMap) Tags() []fix.Tag {
	tags := make([]fix.Tag, 0, len(m.fieldLookup))
	for t := range m.fieldLookup {
		tags = append(tags, t)
	}

	return tags
}

func (m FieldMap) Get(parser Field) MessageRejectError {
	return m.GetField(parser.Tag(), parser)
}

func (m FieldMap) Has(tag fix.Tag) bool {
	_, ok := m.fieldLookup[tag]
	return ok
}

func (m FieldMap) GetField(tag fix.Tag, parser FieldValue) MessageRejectError {
	field, ok := m.fieldLookup[tag]

	if !ok {
		return ConditionallyRequiredFieldMissing(tag)
	}

	if err := parser.Read(field.Value); err != nil {
		return IncorrectDataFormatForValue(tag)
	}

	return nil
}

func (m FieldMap) SetField(tag fix.Tag, field FieldValue) {
	m.fieldLookup[tag] = newFieldBytes(tag, field.Write())
}

func (m FieldMap) Set(field Field) {
	m.fieldLookup[field.Tag()] = newFieldBytes(field.Tag(), field.Write())
}

func (m FieldMap) sortedTags() []fix.Tag {
	sortedTags := make([]fix.Tag, len(m.fieldLookup))
	for tag := range m.fieldLookup {
		sortedTags = append(sortedTags, tag)
	}

	sort.Sort(fieldSort{sortedTags, m.fieldOrder})
	return sortedTags
}

func (m FieldMap) Write(buffer *bytes.Buffer) {
	tags := m.sortedTags()

	for _, tag := range tags {
		if field, ok := m.fieldLookup[tag]; ok {
			buffer.Write(field.Data)
		}
	}
}

/*

func (b FieldMapBuilder) Remove(tag fix.Tag) {
	delete(b.fields, tag)
}
*/

func (m FieldMap) total() int {
	total := 0
	for t, field := range m.fieldLookup {
		switch t {
		case tag.CheckSum: //tag does not contribute to total
		default:
			total += field.Total()
		}
	}

	return total
}

func (m FieldMap) length() int {
	length := 0
	for t := range m.fieldLookup {
		switch t {
		case tag.BeginString, tag.BodyLength, tag.CheckSum: //tags do not contribute to length
		default:
			length += m.fieldLookup[t].Length()
		}
	}

	return length
}
