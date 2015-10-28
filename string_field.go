package quickfix

//StringValue is a container for string, implements FieldValue.
type StringValue struct {
	Value string
}

func (f *StringValue) Read(tv TagValues) (TagValues, error) {
	f.Value = string(tv[0].Value)
	return tv[1:], nil
}

func (f StringValue) Write() []byte {
	return []byte(f.Value)
}

func (f StringValue) Clone() FieldValue {
	return &StringValue{f.Value}
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
