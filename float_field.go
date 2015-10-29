package quickfix

import (
	"fmt"
	"regexp"
	"strconv"
)

//FloatValue is a Container for float, implements FieldValue
type FloatValue struct {
	Value float64
}

func (f *FloatValue) Read(tv TagValues) (TagValues, error) {
	var err error
	bytes := tv[0].Value
	if f.Value, err = strconv.ParseFloat(string(bytes), 64); err != nil {
		return tv, err
	}

	//strconv allows values like "+100.00", which is not allowed for FIX float types
	if valid, _ := regexp.MatchString("^[0-9.-]+$", string(bytes)); !valid {
		return tv, fmt.Errorf("invalid value %v", string(bytes))
	}

	return tv[1:], err
}

func (f FloatValue) Write() []byte {
	return []byte(strconv.FormatFloat(f.Value, 'f', -1, 64))
}

func (f FloatValue) Clone() FieldValue {
	return &FloatValue{f.Value}
}

type FloatField struct {
	tagContainer
	FloatValue
}

func NewFloatField(tag Tag, value float64) *FloatField {
	var f FloatField
	f.tag = tag
	f.Value = value
	return &f
}
