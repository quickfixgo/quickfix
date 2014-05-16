package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
)

//The FieldValue interface is used to write/extract typed field values to/from raw bytes
type FieldValue interface {
	Write() []byte
	Read([]byte) error
}

//Field is the interface implemented by all typed Fields in a Message
type Field interface {
	Tag() fix.Tag
	FieldValue
}
