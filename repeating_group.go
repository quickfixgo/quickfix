package quickfix

import (
	"fmt"
	"math"
	"strconv"
)

//GroupItem interface is used to construct repeating group templates
type GroupItem interface {
	//Tag returns the tag identifying this GroupItem
	Tag() Tag

	//Parameter to Read is tagValues.  For most fields, only the first tagValue will be required.
	//The length of the slice extends from the tagValue mapped to the field to be read through the
	//following fields. This can be useful for GroupItems made up of repeating groups.
	//
	//The Read function returns the remaining tagValues not processed by the GroupItem. If there was a
	//problem reading the field, an error may be returned
	read(tagValues) (tagValues, error)
}

type protoGroupElement struct {
	tag Tag
}

func (t protoGroupElement) Tag() Tag { return t.tag }
func (t protoGroupElement) read(tv tagValues) (tagValues, error) {
	if tv[0].Tag == t.tag {
		return tv[1:], nil
	}

	return tv, nil
}

//GroupElement returns a GroupItem made up of a single field
func GroupElement(tag Tag) GroupItem {
	return protoGroupElement{tag: tag}
}

//GroupTemplate specifies the group item order for a RepeatingGroup
type GroupTemplate []GroupItem

type Group struct{ FieldMap }

//RepeatingGroup is a FIX Repeating Group type
type RepeatingGroup struct {
	Tag
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

//tagValues returns tagValues for all Items in the repeating group ordered by
//Group sequence and Group template order
func (f RepeatingGroup) tagValues() tagValues {
	tvs := make(tagValues, 1, 1)
	tvs[0].init(f.Tag, []byte(strconv.Itoa(len(f.Groups))))

	for _, group := range f.Groups {
		tags := group.sortedTags()

		for _, tag := range tags {
			if fields, ok := group.tagLookup[tag]; ok {
				tvs = append(tvs, fields...)
			}
		}
	}

	return tvs
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

func (f RepeatingGroup) delimiter() Tag {
	return f.GroupTemplate[0].Tag()
}

func (f RepeatingGroup) isDelimiter(t Tag) bool {
	return t == f.delimiter()
}

func (f *RepeatingGroup) read(tv tagValues) (tagValues, error) {
	expectedGroupSize, err := atoi(tv[0].Value)
	if err != nil {
		return tv, err
	}

	if expectedGroupSize == 0 {
		return tv[1:], nil
	}

	tv = tv[1:cap(tv)]
	tagOrdering := f.groupTagOrder()
	var group Group
	group.init(tagOrdering)
	for len(tv) > 0 {
		field, ok := f.findItemInGroupTemplate(tv[0].Tag)
		if !ok {
			break
		}

		tvRange := tv
		if tv, err = field.read(tv); err != nil {
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
		return tv, repeatingGroupFieldsOutOfOrder(f.Tag, fmt.Sprintf("group %v: template is wrong or delimiter %v not found: expected %v groups, but found %v", f.Tag, f.delimiter(), expectedGroupSize, len(f.Groups)))
	}

	return tv, err
}
