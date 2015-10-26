package quickfix

import (
	"bytes"
	"math"
	"sort"
)

//FieldMap is a collection of fix fields that make up a fix message.
type FieldMap struct {
	fieldLookup map[Tag]*fieldBytes
	fieldOrder
}

// fieldOrder true if tag i should occur before tag j
type fieldOrder func(i, j Tag) bool

type fieldSort struct {
	tags    []Tag
	compare fieldOrder
}

func (t fieldSort) Len() int           { return len(t.tags) }
func (t fieldSort) Swap(i, j int)      { t.tags[i], t.tags[j] = t.tags[j], t.tags[i] }
func (t fieldSort) Less(i, j int) bool { return t.compare(t.tags[i], t.tags[j]) }

//in the message header, the first 3 tags in the message header must be 8,9,35
func headerFieldOrder(i, j Tag) bool {
	var ordering = func(t Tag) uint32 {
		switch t {
		case tagBeginString:
			return 1
		case tagBodyLength:
			return 2
		case tagMsgType:
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
func normalFieldOrder(i, j Tag) bool { return i < j }

// In the trailer, CheckSum (tag 10) must be last
func trailerFieldOrder(i, j Tag) bool {
	switch {
	case i == tagCheckSum:
		return false
	case j == tagCheckSum:
		return true
	}

	return i < j
}

func (m *FieldMap) init(ordering fieldOrder) {
	m.fieldLookup = make(map[Tag]*fieldBytes)
	m.fieldOrder = ordering
}

func (m FieldMap) Tags() []Tag {
	tags := make([]Tag, 0, len(m.fieldLookup))
	for t := range m.fieldLookup {
		tags = append(tags, t)
	}

	return tags
}

func (m FieldMap) Get(parser Field) MessageRejectError {
	return m.GetField(parser.Tag(), parser)
}

func (m FieldMap) Has(tag Tag) bool {
	_, ok := m.fieldLookup[tag]
	return ok
}

func (m FieldMap) GetField(tag Tag, parser FieldValue) MessageRejectError {
	field, ok := m.fieldLookup[tag]

	if !ok {
		return conditionallyRequiredFieldMissing(tag)
	}

	if err := parser.Read(field.Value); err != nil {
		return incorrectDataFormatForValue(tag)
	}

	return nil
}

func (m FieldMap) SetField(tag Tag, field FieldValue) {
	m.fieldLookup[tag] = newFieldBytes(tag, field.Write())
}

func (m FieldMap) Set(field Field) {
	m.fieldLookup[field.Tag()] = newFieldBytes(field.Tag(), field.Write())
}

func (m FieldMap) sortedTags() []Tag {
	sortedTags := make([]Tag, len(m.fieldLookup))
	for tag := range m.fieldLookup {
		sortedTags = append(sortedTags, tag)
	}

	sort.Sort(fieldSort{sortedTags, m.fieldOrder})
	return sortedTags
}

func (m FieldMap) write(buffer *bytes.Buffer) {
	tags := m.sortedTags()

	for _, tag := range tags {
		if field, ok := m.fieldLookup[tag]; ok {
			buffer.Write(field.Data)
		}
	}
}

func (m FieldMap) total() int {
	total := 0
	for t, field := range m.fieldLookup {
		switch t {
		case tagCheckSum: //tag does not contribute to total
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
		case tagBeginString, tagBodyLength, tagCheckSum: //tags do not contribute to length
		default:
			length += m.fieldLookup[t].Length()
		}
	}

	return length
}
