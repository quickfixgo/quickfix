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
	FieldMap
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
	fieldMap.FieldMap.init()
	fieldMap.fields = make(map[tag.Tag]*fieldBytes)
	fieldMap.fieldOrder = ordering
}

func (m FieldMapBuilder) SetField(tag tag.Tag, field FieldValue) {
	m.fields[tag] = newFieldBytes(tag, field.Write())
}

func (m FieldMapBuilder) Set(field Field) {
	m.fields[field.Tag()] = newFieldBytes(field.Tag(), field.Write())
}

func (m FieldMapBuilder) sortedTags() []tag.Tag {
	sortedTags := make([]tag.Tag, len(m.fields))
	for tag := range m.fields {
		sortedTags = append(sortedTags, tag)
	}

	sort.Sort(fieldSort{sortedTags, m.fieldOrder})
	return sortedTags
}

func (m FieldMapBuilder) Write(b *bytes.Buffer) {
	tags := m.sortedTags()

	for _, tag := range tags {
		if field, ok := m.fields[tag]; ok {
			b.Write(field.Data)
		}
	}
}

func (m FieldMapBuilder) Remove(tag tag.Tag) {
	delete(m.fields, tag)
}

func (m FieldMapBuilder) total() int {
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
