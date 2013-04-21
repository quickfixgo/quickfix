package basic

import (
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/tag"
	"strconv"
)

type IntField struct {
	FieldBase
	intValue int
}

func (f *IntField) IntValue() int { return f.intValue }

func NewIntField(tag tag.Tag, value int) *IntField {
	f := new(IntField)
	f.init(tag, strconv.Itoa(value))
	f.intValue = value
	return f
}

//Converts a generic field to an IntField.
//Check error for convert errors.
func ToIntField(f message.Field) (*IntField, error) {
	value, err := strconv.Atoi(f.Value())

	if err != nil {
		return nil, err
	}

	return NewIntField(f.Tag(), value), nil
}
