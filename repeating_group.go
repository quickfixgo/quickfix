package quickfix

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

//GroupItem interface is used to construct repeating group templates
type GroupItem interface {
	//Tag returns the tag identifying this GroupItem
	Tag() Tag

	//Parameter to Read is TagValues.  For most fields, only the first TagValue will be required.
	//The length of the slice extends from the TagValue mapped to the field to be read through the
	//following fields. This can be useful for GroupItems made up of repeating groups.
	//
	//The Read function returns the remaining TagValues not processed by the GroupItem. If there was a
	//problem reading the field, an error may be returned
	Read(TagValues) (TagValues, error)
}

type protoGroupElement struct {
	tagMethod func() Tag
}

func (t protoGroupElement) Tag() Tag { return t.tagMethod() }
func (t protoGroupElement) Read(tv TagValues) (TagValues, error) {
	if tv[0].Tag == t.tagMethod() {
		return tv[1:], nil
	}

	return tv, nil
}

//GroupElement returns a GroupItem made up of a single field
func GroupElement(tag Tag) GroupItem {
	t := struct{ protoGroupElement }{}
	t.tagMethod = func() Tag { return tag }
	return t
}

type GroupTemplate []GroupItem
type Group struct{ FieldMap }

type RepeatingGroup struct {
	GroupTemplate
	Groups []Group
}

//Add appends a new group to the RepeatingGroup and returns the new Group
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

func (f RepeatingGroup) findItemInGroupTemplate(t Tag) (item GroupItem, ok bool) {
	for _, templateField := range f.GroupTemplate {
		if t == templateField.Tag() {
			ok = true
			item = templateField
			break
		}
	}

	return
}

func (f RepeatingGroup) groupTagOrder() tagOrder {
	tagMap := make(map[Tag]int)
	for i, f := range f.GroupTemplate {
		tagMap[f.Tag()] = i
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
	return t == f.GroupTemplate[0].Tag()
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
		field, ok := f.findItemInGroupTemplate(tv[0].Tag)
		if !ok {
			break
		}

		tvRange := tv
		if tv, err = field.Read(tv); err != nil {
			return tv, err
		}

		if f.isDelimiter(field.Tag()) {
			group = Group{}
			group.init(tagOrdering)

			f.Groups = append(f.Groups, group)
		}

		group.tagLookup[tvRange[0].Tag] = tvRange
	}

	if len(f.Groups) != expectedGroupSize {
		return tv, fmt.Errorf("Only found %v instead of %v expected groups, is template wrong?", len(f.Groups), expectedGroupSize)
	}
	return tv, err
}
