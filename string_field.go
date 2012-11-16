package quickfixgo

// StringField is the common interface implemented by all string fields
type StringField interface { 
  Field 
}

/*func NewStringField(tag Tag, value string) StringField {
  field:=New(tag, value)
  return field
}
*/
