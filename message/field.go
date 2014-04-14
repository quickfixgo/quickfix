package message

import (
	"github.com/quickfixgo/quickfix/fix"
)

//The FieldValue interface is used to write/extract typed field values to/from raw bytes
type FieldValue interface {
	Write() []byte
	Read([]byte) error
}

//FieldWriter is a read-only interface for writing field values. Should not require pointer receiver from interface impl.
type FieldWriter interface {
	Tag() fix.Tag
	Write() []byte
}

//Field is the interface implemented by all typed Fields in a Message
type Field interface {
	Tag() fix.Tag
	FieldValue
}
