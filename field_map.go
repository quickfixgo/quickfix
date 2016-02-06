package quickfix

import (
	"bytes"
	"math"
	"sort"
)

//FieldMap is a collection of fix fields that make up a fix message.
type FieldMap struct {
	tagLookup map[Tag][]TagValue
	tagOrder
}

// tagOrder true if tag i should occur before tag j
type tagOrder func(i, j Tag) bool

type tagSort struct {
	tags    []Tag
	compare tagOrder
}

func (t tagSort) Len() int           { return len(t.tags) }
func (t tagSort) Swap(i, j int)      { t.tags[i], t.tags[j] = t.tags[j], t.tags[i] }
func (t tagSort) Less(i, j int) bool { return t.compare(t.tags[i], t.tags[j]) }

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

func (m *FieldMap) init(ordering tagOrder) {
	m.tagLookup = make(map[Tag][]TagValue)
	m.tagOrder = ordering
}

func (m FieldMap) Tags() []Tag {
	tags := make([]Tag, 0, len(m.tagLookup))
	for t := range m.tagLookup {
		tags = append(tags, t)
	}

	return tags
}

func (m FieldMap) Get(parser Field) MessageRejectError {
	return m.GetField(parser.Tag(), parser)
}

func (m FieldMap) Has(tag Tag) bool {
	_, ok := m.tagLookup[tag]
	return ok
}

func (m FieldMap) GetField(tag Tag, parser FieldValueReader) MessageRejectError {
	tagValues, ok := m.tagLookup[tag]
	if !ok {
		return conditionallyRequiredFieldMissing(tag)
	}

	if err := parser.Read(tagValues[0].Value); err != nil {
		return incorrectDataFormatForValue(tag)
	}

	return nil
}

func (m FieldMap) GetGroup(parser *RepeatingGroup) MessageRejectError {
	tagValues, ok := m.tagLookup[parser.Tag]
	if !ok {
		return conditionallyRequiredFieldMissing(parser.Tag)
	}

	if _, err := parser.Read(tagValues); err != nil {
		return incorrectDataFormatForValue(parser.Tag)
	}

	return nil
}

func (m FieldMap) SetField(tag Tag, field FieldValueWriter) FieldMap {
	tValues := make([]TagValue, 1)
	tValues[0].init(tag, field.Write())
	m.tagLookup[tag] = tValues
	return m
}

func (m FieldMap) Set(field FieldWriter) FieldMap {
	tValues := make([]TagValue, 1)
	tValues[0].init(field.Tag(), field.Write())
	m.tagLookup[field.Tag()] = tValues
	return m
}

func (m FieldMap) SetGroup(field RepeatingGroup) FieldMap {
	m.tagLookup[field.Tag] = field.TagValues()
	return m
}

func (m FieldMap) sortedTags() []Tag {
	sortedTags := make([]Tag, len(m.tagLookup))
	for tag := range m.tagLookup {
		sortedTags = append(sortedTags, tag)
	}

	sort.Sort(tagSort{sortedTags, m.tagOrder})
	return sortedTags
}

func (m FieldMap) write(buffer *bytes.Buffer) {
	tags := m.sortedTags()

	for _, tag := range tags {
		if fields, ok := m.tagLookup[tag]; ok {
			for _, tv := range fields {
				buffer.Write(tv.bytes)
			}
		}
	}
}

func (m FieldMap) total() int {
	total := 0
	for _, fields := range m.tagLookup {
		for _, tv := range fields {
			switch tv.Tag {
			case tagCheckSum: //tag does not contribute to total
			default:
				total += tv.total()
			}
		}
	}

	return total
}

func (m FieldMap) length() int {
	length := 0
	for _, fields := range m.tagLookup {
		for _, tv := range fields {
			switch tv.Tag {
			case tagBeginString, tagBodyLength, tagCheckSum: //tags do not contribute to length
			default:
				length += tv.length()
			}
		}
	}

	return length
}
