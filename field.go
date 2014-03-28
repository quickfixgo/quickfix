package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

type FieldValue interface {
	ConvertValueFromBytes([]byte) error
}

type Field interface {
	Tag() tag.Tag
	FieldValue
}
