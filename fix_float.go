package quickfix

import (
	"fmt"
	"regexp"
	"strconv"
)

//FIXFloat is a FIX Float Value, implements FieldValue
type FIXFloat float64

//NewFIXFloat returns an initialized FIXFloat
func NewFIXFloat(val float64) *FIXFloat {
	f := FIXFloat(val)
	return &f
}

func (f *FIXFloat) Read(tv TagValues) (TagValues, error) {
	bytes := tv[0].Value
	val, err := strconv.ParseFloat(string(bytes), 64)
	if err != nil {
		return tv, err
	}

	//strconv allows values like "+100.00", which is not allowed for FIX float types
	if valid, _ := regexp.MatchString("^[0-9.-]+$", string(bytes)); !valid {
		return tv, fmt.Errorf("invalid value %v", string(bytes))
	}

	*f = FIXFloat(val)

	return tv[1:], err
}

func (f FIXFloat) Write() []byte {
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64))
}

func (f FIXFloat) Clone() FieldValue {
	clone := f
	return &clone
}
