package message

import (
	"bytes"
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

//FieldMap maps tags to Fields. This interface is used by the Message type.
type FieldMap interface {
	Tags() []tag.Tag
	Length() int
	Total() int

	//ok will be 'true' if field exists in map
	Field(tag tag.Tag) (field Field, ok bool)

	//ok will be 'true' if field exists in map
	StringValue(tag tag.Tag) (stringField string, ok bool)

	//check error for FieldNotFoundError, FieldConvertError
	IntValue(tag tag.Tag) (intField int, err error)

	//check error for FieldNotFoundError, FieldConvertError
	UTCTimestampValue(tag tag.Tag) (timeField time.Time, err error)

	//check error for FieldNotFoundError, FieldConvertError
	BooleanValue(tag tag.Tag) (boolField bool, err error)

	Write(b *bytes.Buffer)
}
