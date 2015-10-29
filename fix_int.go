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

//NewFIXInt returns an initialized FIXInt
func NewFIXInt(val int) *FIXInt {
	i := FIXInt(val)
	return &i
}

func (f *FIXInt) Read(tv TagValues) (TagValues, error) {
	i, err := atoi(tv[0].Value)
	if err != nil {
		return tv, err
	}
	*f = FIXInt(i)
	return tv[1:], nil
}

func (f FIXInt) Write() []byte {
	return []byte(strconv.Itoa(int(f)))
}

func (f FIXInt) Clone() FieldValue {
	clone := f
	return &clone
}
