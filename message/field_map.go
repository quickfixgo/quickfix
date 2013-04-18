package message

import (
	"bytes"
	"time"
)

//FieldMap maps tags to Fields. This interface is used by the Message type.
type FieldMap interface {
	Tags() []Tag
	Length() int
	Total() int

	//ok will be 'true' if field exists in map
	Field(tag Tag) (field Field, ok bool)

	//ok will be 'true' if field exists in map
	StringValue(tag Tag) (stringField string, ok bool)

	//check error for FieldNotFoundError, FieldConvertError
	IntValue(tag Tag) (intField int, err error)

	//check error for FieldNotFoundError, FieldConvertError
	UTCTimestampValue(tag Tag) (timeField time.Time, err error)

	//check error for FieldNotFoundError, FieldConvertError
	BooleanValue(tag Tag) (boolField bool, err error)

	Write(b *bytes.Buffer)
}
