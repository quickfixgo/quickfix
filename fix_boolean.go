package quickfix

import (
	"errors"
)

//FIXBoolean is a FIX Boolean value, implements FieldValue.
type FIXBoolean bool

//Bool converts the FIXBoolean value to bool
func (f FIXBoolean) Bool() bool { return bool(f) }

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
