package message

import (
	"fmt"
	"github.com/quickfixgo/quickfix/fix"
)

type ParseError struct {
	OrigError string
}

func (e ParseError) Error() string { return fmt.Sprintf("error parsing message: %s", e.OrigError) }

type FieldConvertError struct {
	fix.Tag        //Tag for field that failed conversion
	Value   string //String representation for field
}

func (f FieldConvertError) Error() string { return "Field Convert Error" }

type FieldNotFoundError struct {
	fix.Tag //Tag for field missing
}

func (f FieldNotFoundError) Error() string { return "Field Not Found" }
