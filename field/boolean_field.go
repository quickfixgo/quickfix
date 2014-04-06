package field

import (
	"errors"
	"github.com/quickfixgo/quickfix/tag"
)

//Container for bool, implements FieldValue
type BooleanValue struct {
	Value bool
}

func (f *BooleanValue) Read(bytes []byte) error {
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

//Generic boolean Field Type. Implements Field
type BooleanField struct {
	tagContainer
	BooleanValue
}

func NewBooleanField(tag tag.Tag, value bool) *BooleanField {
	var f BooleanField
	f.tag = tag
	f.Value = value

	return &f
}
