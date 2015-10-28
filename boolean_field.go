package quickfix

import (
	"errors"
)

//BooleanValue is a container for bool, implements FieldValue.
type BooleanValue struct {
	Value bool
}

func (f *BooleanValue) Read(tv []TagValue) error {
	bytes := tv[0].Value
	switch string(bytes) {
	case "Y":
		f.Value = true
	case "N":
		f.Value = false
	default:
		return errors.New("Invalid Value for bool: " + string(bytes))
	}

	return nil
}

func (f BooleanValue) Write() []byte {
	if f.Value {
		return []byte("Y")
	}

	return []byte("N")
}

//BooleanField is a generic boolean Field Type, Implements Field.
type BooleanField struct {
	tagContainer
	BooleanValue
}

func NewBooleanField(tag Tag, value bool) *BooleanField {
	var f BooleanField
	f.tag = tag
	f.Value = value

	return &f
}
