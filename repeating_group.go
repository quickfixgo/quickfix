package quickfix

import (
	"bytes"
	"strconv"
)

type RepeatingGroupField struct {
	Tag
	FieldValue
}

type Group []RepeatingGroupField

type RepeatingGroup struct {
	GroupTemplate Group
	Groups        []Group
}

func (f *RepeatingGroup) AddGroup(group Group) {
	f.Groups = append(f.Groups, group)
}

func (f RepeatingGroup) Write() []byte {
	buf := bytes.NewBufferString(strconv.Itoa(len(f.Groups)))

	for _, group := range f.Groups {
		for _, field := range group {
			buf.WriteString("")
			buf.WriteString(strconv.Itoa(int(field.Tag)))
			buf.WriteString("=")
			buf.Write(field.Write())
		}
	}

	return buf.Bytes()
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
	var group Group
	for len(tv) > 0 {
		field, ok := f.findFieldInGroupTemplate(tv[0].Tag)
		if !ok {
			break
		}

		if tv, err = field.Read(tv); err != nil {
			return tv, err
		}

		if f.isDelimiter(field.Tag) {
			group = Group{field}
			f.Groups = append(f.Groups, group)
		} else {
			if len(group) == 0 {
				// didn't get initial delimiter
				break
			}

			group = append(group, field)
			f.Groups[len(f.Groups)-1] = group
		}
	}

	return tv, err
}

func (f RepeatingGroup) Clone() FieldValue {
	var clone RepeatingGroup

	clone.GroupTemplate = make(Group, len(f.GroupTemplate))
	clone.Groups = make([]Group, len(f.Groups))

	for i, field := range f.GroupTemplate {
		clone.GroupTemplate[i] = RepeatingGroupField{field.Tag, field.FieldValue.Clone()}
	}

	for i, group := range f.Groups {
		clone.Groups[i] = make(Group, len(group))

		for j, field := range group {
			clone.Groups[i][j] = RepeatingGroupField{field.Tag, field.FieldValue.Clone()}
		}
	}

	return &clone
}
