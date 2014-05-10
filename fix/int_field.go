package fix

import (
	"strconv"
)

//IntValue is a Container for int, implements FieldValue
type IntValue struct {
	Value int
}

func (f *IntValue) Read(bytes []byte) (err error) {
	f.Value, err = strconv.Atoi(string(bytes))

	return
}

func (f IntValue) Write() []byte {
	return []byte(strconv.Itoa(f.Value))
}

//IntField is a generic int Field Type, implements Field
type IntField struct {
	tagContainer
	IntValue
}

func NewIntField(tag Tag, value int) *IntField {
	var f IntField
	f.tag = tag
	f.Value = value
	return &f
}
