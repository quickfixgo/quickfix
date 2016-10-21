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

//FIXInt is a FIX Int Value, implements FieldValue
type FIXInt int

//Int converts the FIXInt value to int
func (f FIXInt) Int() int { return int(f) }

func (f *FIXInt) Read(bytes []byte) error {
	i, err := atoi(bytes)
	if err != nil {
		return err
	}
	*f = FIXInt(i)
	return nil
}

func (f FIXInt) Write() []byte {
	return strconv.AppendInt(nil, int64(f), 10)
}
