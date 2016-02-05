package quickfix

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type RepeatingGroupField struct {
	Tag
	FieldValue
}

type GroupTemplate []RepeatingGroupField
type Group struct{ FieldMap }

type RepeatingGroup struct {
	GroupTemplate
	Groups []Group
}

func (f *RepeatingGroup) Add() Group {
	var g Group
	g.init(f.groupTagOrder())

	f.Groups = append(f.Groups, g)
	return g
}

func (f RepeatingGroup) Write() []byte {
	buf := bytes.NewBufferString(strconv.Itoa(len(f.Groups)))
	buf.WriteString("")

	for _, group := range f.Groups {
		group.write(buf)
	}

	//remove the last soh char
	bytes := buf.Bytes()
	return bytes[:len(bytes)-1]
}

func (f RepeatingGroup) findFieldInGroupTemplate(t Tag) (field RepeatingGroupField, ok bool) {
	for _, templateField := range f.GroupTemplate {
		if t == templateField.Tag {
			ok = true
			field.Tag = templateField.Tag
			field.FieldValue = templateField.Clone()
			break
		}
	}

	return
}

func (f RepeatingGroup) groupTagOrder() tagOrder {
	tagMap := make(map[Tag]int)
	for i, f := range f.GroupTemplate {
		tagMap[f.Tag] = i
	}

	return func(i, j Tag) bool {
		orderi := math.MaxInt64
		orderj := math.MaxInt64

		if iIndex, ok := tagMap[i]; ok {
			orderi = iIndex
		}

		if jIndex, ok := tagMap[j]; ok {
			orderj = jIndex
		}

		return orderi < orderj
	}
}

func (f RepeatingGroup) isDelimiter(t Tag) bool {
	return t == f.GroupTemplate[0].Tag
}

func (f *RepeatingGroup) Read(tv TagValues) (TagValues, error) {
	var err error
	expectedGroupSize, err := atoi(tv[0].Value)
	if err != nil {
		return tv, err
	}

	if expectedGroupSize == 0 {
		return tv[1:], nil
	}

	tv = tv[1:]
	tagOrdering := f.groupTagOrder()
	var group Group
	group.init(tagOrdering)
	for len(tv) > 0 {
		field, ok := f.findFieldInGroupTemplate(tv[0].Tag)
		if !ok {
			break
		}

		if tv, err = field.Read(tv); err != nil {
			return tv, err
		}

		if f.isDelimiter(field.Tag) {
			group = Group{}
			group.init(tagOrdering)

			f.Groups = append(f.Groups, group)
		}

		group.SetField(field.Tag, field)
	}

	if len(f.Groups) != expectedGroupSize {
		return tv, fmt.Errorf("Only found %v instead of %v expected groups, is template wrong?", len(f.Groups), expectedGroupSize)
	}
	return tv, err
}

func (f RepeatingGroup) Clone() FieldValue {
	var clone RepeatingGroup
	clone.GroupTemplate = make(GroupTemplate, len(f.GroupTemplate))
	clone.Groups = make([]Group, len(f.Groups))

	for i, field := range f.GroupTemplate {
		clone.GroupTemplate[i] = RepeatingGroupField{field.Tag, field.FieldValue.Clone()}
	}

	for i, group := range f.Groups {
		clone.Groups[i].init(group.tagOrder)
	}

	return &clone
}
