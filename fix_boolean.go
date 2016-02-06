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

func (f *FIXBoolean) Read(bytes []byte) error {
	switch string(bytes) {
	case "Y":
		*f = true
	case "N":
		*f = false
	default:
		return errors.New("Invalid Value for bool: " + string(bytes))
	}

	return nil
}

func (f FIXBoolean) Write() []byte {
	if f {
		return []byte("Y")
	}

	return []byte("N")
}
