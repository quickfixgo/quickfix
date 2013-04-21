package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

type StringField struct {
	FieldBase
}

func NewStringField(tag tag.Tag, value string) *StringField {
	f := new(StringField)
	f.init(tag, value)
	return f
}
