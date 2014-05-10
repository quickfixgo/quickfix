package fix

//StringValue is a container for string, implements FieldValue.
type StringValue struct {
	Value string
}

func (f *StringValue) Read(bytes []byte) error {
	f.Value = string(bytes)
	return nil
}

func (f StringValue) Write() []byte {
	return []byte(f.Value)
}

//StringField is a generic string Field Type. Implements Field.
type StringField struct {
	tagContainer
	StringValue
}

func NewStringField(tag Tag, value string) *StringField {
	var f StringField
	f.tag = tag
	f.Value = value

	return &f
}
