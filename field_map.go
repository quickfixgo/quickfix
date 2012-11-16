package quickfixgo

import(
  "bytes"
)

//FieldMap maps tags to Fields. This interface is used by the Message type.
type FieldMap interface {
  Tags() []Tag
  Length() int
  Total() int
  Get(tag Tag) (field Field, ok bool)

  IntField(tag Tag) (intField IntField, ok bool)
  StringField(tag Tag) (stringField StringField, ok bool)
  UTCTimestampField(tag Tag) (timeField UTCTimestampField, ok bool)
  Write(b *bytes.Buffer)
}
