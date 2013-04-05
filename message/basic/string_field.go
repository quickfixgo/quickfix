package basic

import (
	"github.com/cbusbey/quickfixgo/message"
)

type StringField struct {
	FieldBase
}

func NewStringField(tag message.Tag, value string) *StringField {
	f := new(StringField)
	f.init(tag, value)
	return f
}
