package basic

import (
	"errors"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/tag"
)

type BooleanField struct {
	FieldBase
	boolValue bool
}

func (f *BooleanField) BooleanValue() bool { return f.boolValue }

func NewBooleanField(tag tag.Tag, value bool) *BooleanField {
	f := new(BooleanField)
	if value {
		f.init(tag, "Y")
	} else {
		f.init(tag, "N")
	}

	f.boolValue = value
	return f
}

//Converts a generic field to a BooleanField.
//Check error for convert errors.
func ToBooleanField(f message.Field) (*BooleanField, error) {

	switch f.Value() {
	case "Y":
		return NewBooleanField(f.Tag(), true), nil
	case "N":
		return NewBooleanField(f.Tag(), false), nil
	}

	return nil, errors.New("Invalid Value for bool: " + f.Value())
}
