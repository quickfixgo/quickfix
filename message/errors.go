package message

import("fmt")

type ParseError struct {
  OrigError string
}

func (e ParseError) Error() string {return fmt.Sprintf("Error parsing message: %s", e.OrigError)}

type FieldConvertError struct {}
func (f *FieldConvertError) Error() string {return "Field Convert Error"}


