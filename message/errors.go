package message

import("fmt")

type ParseError struct {
  OrigError string
}

func (e ParseError) Error() string {return fmt.Sprintf("Error parsing message: %s", e.OrigError)}

type FieldConvertError struct {
  Tag //Tag for field that failed conversion
  Value string //String representation for field
}
func (f FieldConvertError) Error() string {return "Field Convert Error"}

type FieldNotFoundError struct {
  Tag //Tag for field missing 
}
func (f FieldNotFoundError) Error() string {return "Field Not Found"}
