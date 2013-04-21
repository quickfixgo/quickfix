package quickfixgo

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
	"strconv"
)

//Basic impl of Field interface 
type FieldBase struct {
	tag   tag.Tag
	value string
	data  string
	total int
}

func (f FieldBase) Tag() tag.Tag   { return f.tag }
func (f FieldBase) Value() string  { return f.value }
func (f FieldBase) String() string { return f.data }
func (f FieldBase) Length() int    { return len(f.data) }
func (f FieldBase) Total() int     { return f.total }

func (field *FieldBase) computeTotal() {
	field.total = 0

	for _, b := range []byte(field.data) {
		field.total += int(b)
	}
}

func (f *FieldBase) init(tag tag.Tag, value string) {
	f.tag = tag
	f.value = value
	f.data = fmt.Sprint(tag, "=", value, "")
	f.computeTotal()
}

func ParseField(fieldBytes []byte) (*FieldBase, error) {
	sepIndex := bytes.IndexByte(fieldBytes, '=')

	if sepIndex == -1 {
		return nil, fmt.Errorf("field.Parse: No '=' in '%s'", fieldBytes)
	}

	parsed_tag, err := strconv.Atoi(string(fieldBytes[:sepIndex]))

	if err != nil {
		return nil, fmt.Errorf("field.Parse: %s", err.Error())
	}

	field := FieldBase{
		tag:   tag.Tag(parsed_tag),
		value: string(fieldBytes[(sepIndex + 1):(len(fieldBytes) - 1)]),
		data:  string(fieldBytes),
	}
	field.computeTotal()

	return &field, nil
}
