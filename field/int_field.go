package field

import (
	"github.com/cbusbey/quickfixgo/tag"
	"strconv"
)

//Container for int, knows part of FieldConverter interface
type IntValue struct {
	Value int
}

func (f *IntValue) ConvertValueFromBytes(bytes []byte) (err error) {
	f.Value, err = strconv.Atoi(string(bytes))

	return
}

func (f IntValue) ConvertValueToBytes() []byte {
	return []byte(strconv.Itoa(f.Value))
}

//Generic int Field Type. Implements FieldConverter.
type IntField struct {
	FieldTag tag.Tag
	IntValue
}

func (f IntField) Tag() tag.Tag {
	return f.FieldTag
}

func NewIntField(tag tag.Tag, value int) *IntField {
	f := IntField{FieldTag: tag}
	f.Value = value
	return &f
}
