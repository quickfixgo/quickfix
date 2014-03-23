package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
)

type SessionNotFound interface {
	error
}

type ParseError struct {
	OrigError string
}

func (e ParseError) Error() string { return fmt.Sprintf("Error parsing message: %s", e.OrigError) }

type FieldConvertError struct {
	tag.Tag        //Tag for field that failed conversion
	Value   string //String representation for field
}

func (f FieldConvertError) Error() string { return "Field Convert Error" }

type FieldNotFoundError struct {
	tag.Tag //Tag for field missing
}

func (f FieldNotFoundError) Error() string { return "Field Not Found" }
