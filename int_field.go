package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	"strconv"
)

//Container for int, implements FieldValue
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

//Generic int Field Type. Implements Field
type IntField struct {
	tagContainer
	IntValue
}

func NewIntField(tag tag.Tag, value int) *IntField {
	var f IntField
	f.tag = tag
	f.Value = value
	return &f
}
