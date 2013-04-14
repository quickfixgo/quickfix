package basic

import (
	"github.com/cbusbey/quickfixgo/message"
)

type BooleanField struct {
	FieldBase
	boolValue bool
}

func (f *BooleanField) BooleanValue() bool { return f.boolValue }

func NewBooleanField(tag message.Tag, value bool) *BooleanField {
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

	return nil, message.FieldConvertError{f.Tag(), f.Value()}
}
