//message package contains types and functions specific to generic fix messages
package message

type Message interface {
  Header() FieldMap
  Trailer() FieldMap
  Body() FieldMap
  Length() int
  String() string
}
