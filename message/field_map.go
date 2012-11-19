package message

import(
  "bytes"
)

//FieldMap maps tags to Fields. This interface is used by the Message type.
type FieldMap interface {
  Tags() []Tag
  Length() int
  Total() int

  //ok will be 'true' if field exists in map
  Get(tag Tag) (field Field, ok bool)

  //check error for FieldNotFoundError, FieldConvertError
  IntField(tag Tag) (intField IntField, err error)

  //check error for FieldNotFoundError, FieldConvertError
  StringField(tag Tag) (stringField StringField, err error)

  //check error for FieldNotFoundError, FieldConvertError
  UTCTimestampField(tag Tag) (timeField UTCTimestampField, err error)

  Write(b *bytes.Buffer)
}
