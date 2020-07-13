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
	Read([]TagValue) ([]TagValue, error)

	//Clone makes a copy of this GroupItem
	Clone() GroupItem
}

type protoGroupElement struct {
	tag Tag
}

func (t protoGroupElement) Tag() Tag { return t.tag }
func (t protoGroupElement) Read(tv []TagValue) ([]TagValue, error) {
	if tv[0].tag == t.tag {
		return tv[1:], nil
	}

	return tv, nil
}

func (t protoGroupElement) Clone() GroupItem { return t }

//GroupElement returns a GroupItem made up of a single field
func GroupElement(tag Tag) GroupItem {
	return protoGroupElement{tag: tag}
}

//GroupTemplate specifies the group item order for a RepeatingGroup
type GroupTemplate []GroupItem

//Clone makes a copy of this GroupTemplate
func (gt GroupTemplate) Clone() GroupTemplate {
	clone := make(GroupTemplate, len(gt))
	for i := range gt {
		clone[i] = gt[i].Clone()
	}

	return clone
}

//Group is a group of fields occurring in a repeating group
type Group struct{ FieldMap }

//RepeatingGroup is a FIX Repeating Group type
type RepeatingGroup struct {
	tag      Tag
	template GroupTemplate
	groups   []*Group
}

//NewRepeatingGroup returns an initilized RepeatingGroup instance
func NewRepeatingGroup(tag Tag, template GroupTemplate) *RepeatingGroup {
	return &RepeatingGroup{
		tag:      tag,
		template: template,
	}
}

//Tag returns the Tag for this repeating Group
func (f RepeatingGroup) Tag() Tag {
	return f.tag
}

//Clone makes a copy of this RepeatingGroup (tag, template)
func (f RepeatingGroup) Clone() GroupItem {
	return &RepeatingGroup{
		tag:      f.tag,
		template: f.template.Clone(),
	}
}

//Len returns the number of Groups in this RepeatingGroup
func (f RepeatingGroup) Len() int {
	return len(f.groups)
}

//Get returns the ith group in this RepeatingGroup
func (f RepeatingGroup) Get(i int) *Group {
	return f.groups[i]
}

//Add appends a new group to the RepeatingGroup and returns the new Group
func (f *RepeatingGroup) Add() *Group {
	g := new(Group)
	g.initWithOrdering(f.groupTagOrder())

	f.groups = append(f.groups, g)
	return g
}

//Write returns tagValues for all Items in the repeating group ordered by
//Group sequence and Group template order
func (f RepeatingGroup) Write() []TagValue {
	tvs := make([]TagValue, 1)
	tvs[0].init(f.tag, []byte(strconv.Itoa(len(f.groups))))

	for _, group := range f.groups {
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
	for _, templateField := range f.template {
		if t == templateField.Tag() {
			ok = true
			item = templateField.Clone()
			break
		}
	}

	return
}

func (f RepeatingGroup) groupTagOrder() tagOrder {
	tagMap := make(map[Tag]int)
	for i, f := range f.template {
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
	return f.template[0].Tag()
}

func (f RepeatingGroup) isDelimiter(t Tag) bool {
	return t == f.delimiter()
}

func (f *RepeatingGroup) Read(tv []TagValue) ([]TagValue, error) {
	expectedGroupSize, err := atoi(tv[0].value)
	if err != nil {
		return tv, err
	}

	if expectedGroupSize == 0 {
		return tv[1:], nil
	}

	tv = tv[1:cap(tv)]
	tagOrdering := f.groupTagOrder()
	group := new(Group)
	group.initWithOrdering(tagOrdering)
	for len(tv) > 0 {
		gi, ok := f.findItemInGroupTemplate(tv[0].tag)
		if !ok {
			break
		}

		tvRange := tv
		if tv, err = gi.Read(tv); err != nil {
			return tv, err
		}

		if f.isDelimiter(gi.Tag()) {
			group = new(Group)
			group.initWithOrdering(tagOrdering)

			f.groups = append(f.groups, group)
		}

		group.tagLookup[tvRange[0].tag] = tvRange
	}

	if len(f.groups) != expectedGroupSize {
		return tv, repeatingGroupFieldsOutOfOrder(f.tag, fmt.Sprintf("group %v: template is wrong or delimiter %v not found: expected %v groups, but found %v", f.tag, f.delimiter(), expectedGroupSize, len(f.groups)))
	}

	return tv, err
}
