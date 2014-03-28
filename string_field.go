package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

//Container for int, knows part of FieldConverter interface
type StringValue struct {
	Value string
}

func (f *StringValue) ConvertValueFromBytes(bytes []byte) error {
	f.Value = string(bytes)
	return nil
}

func (f StringValue) ConvertValueToBytes() []byte {
	return []byte(f.Value)
}

//Generic string Field Type. Implements FieldConverter.
type StringField struct {
	FieldTag tag.Tag
	StringValue
}

func (f StringField) Tag() tag.Tag {
	return f.FieldTag
}

func NewStringField(tag tag.Tag, value string) *StringField {
	f := StringField{FieldTag: tag}
	f.Value = value

	return &f
}
