package quickfix

import (
	"errors"
)

//FIXBoolean is a FIX Boolean value, implements FieldValue.
type FIXBoolean bool

//NewFIXBoolean returns an initialized FIXBoolean
func NewFIXBoolean(val bool) *FIXBoolean {
	b := FIXBoolean(val)
	return &b
}

func (f *FIXBoolean) Read(tv TagValues) (TagValues, error) {
	bytes := tv[0].Value
	switch string(bytes) {
	case "Y":
		*f = true
	case "N":
		*f = false
	default:
		return tv, errors.New("Invalid Value for bool: " + string(bytes))
	}

	return tv[1:], nil
}

func (f FIXBoolean) Write() []byte {
	if f {
		return []byte("Y")
	}

	return []byte("N")
}

func (f FIXBoolean) Clone() FieldValue {
	clone := f
	return &clone
}
