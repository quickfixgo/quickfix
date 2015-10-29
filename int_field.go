package quickfix

import (
	"errors"
	"strconv"
)

const (
	//-
	asciiMinus = 45

	//ascii numbers 0-9
	ascii0 = 48
	ascii9 = 57
)

//atoi is similar to the function in strconv, but is tuned for ints appearing in FIX field types.
func atoi(d []byte) (int, error) {
	if d[0] == asciiMinus {
		n, err := parseUInt(d[1:])
		return (-1) * n, err
	}

	return parseUInt(d)
}

//parseUInt is similar to the function in strconv, but is tuned for ints appearing in FIX field types.
func parseUInt(d []byte) (n int, err error) {
	if len(d) == 0 {
		err = errors.New("empty bytes")
		return
	}

	for _, dec := range d {
		if dec < ascii0 || dec > ascii9 {
			err = errors.New("invalid format")
			return
		}

		n = n*10 + (int(dec) - ascii0)
	}

	return
}

//IntValue is a Container for int, implements FieldValue
type IntValue struct {
	Value int
}

func (f *IntValue) Read(tv TagValues) (TagValues, error) {
	var err error
	if f.Value, err = atoi(tv[0].Value); err != nil {
		return tv, err
	}

	return tv[1:], nil
}

func (f IntValue) Write() []byte {
	return []byte(strconv.Itoa(f.Value))
}

func (f IntValue) Clone() FieldValue {
	return &IntValue{f.Value}
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
