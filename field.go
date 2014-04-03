package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

//The FieldValue interface is used to write/extract typed field values to/from raw bytes
type FieldValue interface {
	Write() []byte
	Read([]byte) error
}

//Field is the interface implemented by all typed Fields in a Message
type Field interface {
	Tag() tag.Tag
	FieldValue
}

type tagContainer struct {
	tag tag.Tag
}

func (c tagContainer) Tag() tag.Tag {
	return c.tag
}
