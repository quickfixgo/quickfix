package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
)

//The FieldConverter interface is used to write/extract typed field values to/from raw bytes
type FieldConverter interface {
	Tag() tag.Tag
	ConvertValueToBytes() []byte
	ConvertValueFromBytes([]byte) error
}
