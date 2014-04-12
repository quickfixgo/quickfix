package message

import (
	"fmt"
	"github.com/quickfixgo/quickfix/fix"
	"regexp"
	"strconv"
)

//FloatValue is a Container for float, implements FieldValue
type FloatValue struct {
	Value float64
}

func (f *FloatValue) Read(bytes []byte) (err error) {
	if f.Value, err = strconv.ParseFloat(string(bytes), 64); err != nil {
		return
	}

	//strconv allows values like "+100.00", which is not allowed for FIX float types
	if valid, _ := regexp.MatchString("^[0-9.-]+$", string(bytes)); !valid {
		return fmt.Errorf("invalid value %v", string(bytes))
	}

	return
}

func (f FloatValue) Write() []byte {
	return []byte(strconv.FormatFloat(f.Value, 'f', -1, 64))
}

type FloatField struct {
	tagContainer
	FloatValue
}

func NewFloatField(tag fix.Tag, value float64) *FloatField {
	var f FloatField
	f.tag = tag
	f.Value = value
	return &f
}
