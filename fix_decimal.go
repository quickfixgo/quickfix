package quickfix

import "github.com/shopspring/decimal"

//FIXDecimal is a FIX Float Value that implements an arbitrary precision fixed-point decimal.  Implements FieldValue
type FIXDecimal struct {
	decimal.Decimal

	//Scale is the number of digits after the decimal point when Writing the field value as a FIX value
	Scale int32
}

func (d FIXDecimal) Write() []byte {
	return []byte(d.Decimal.StringFixed(d.Scale))
}

func (d *FIXDecimal) Read(bytes []byte) (err error) {
	d.Decimal, err = decimal.NewFromString(string(bytes))
	return
}
