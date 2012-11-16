package quickfixgo

type Message interface {
  Header() FieldMap
  Trailer() FieldMap
  Body() FieldMap
  Length() int
  String() string
}
