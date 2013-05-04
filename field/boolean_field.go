package field

import (
	"errors"
	"github.com/cbusbey/quickfixgo/tag"
)

//Container for bool, knows part of FieldConverter interface
type BooleanValue struct {
	Value bool
}

func (f *BooleanValue) ConvertValueFromBytes(bytes []byte) error {
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

func (f BooleanValue) ConvertValueToBytes() []byte {
	if f.Value {
		return []byte("Y")
	}

	return []byte("N")
}

//Generic boolean Field Type. Implements FieldConverter
type BooleanField struct {
	FieldTag tag.Tag
	BooleanValue
}

func (f BooleanField) Tag() tag.Tag {
	return f.FieldTag
}

func NewBooleanField(tag tag.Tag, value bool) *BooleanField {
	f := BooleanField{FieldTag: tag}
	f.Value = value

	return &f
}
