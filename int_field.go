package quickfixgo

// IntField is the common interface implemented by all int fields
type IntField interface {
  Field

  //field value as an integer
  IntValue() int
}

/*type intFieldBase struct {
  Field
  intValue int
}

func (f *intFieldBase) IntValue() int {return f.intValue}

func NewIntField(tag Tag, value int) IntField {
  return &intFieldBase{New(tag, strconv.Itoa(value)), value}
}

//Converts a generic field to an intfield.
//Check error for convert errors.
func ToIntField(f Field) (IntField, error) {
  value, err:=strconv.Atoi(f.Value())

  if err!=nil {
    return nil,err
  }

  return NewIntField(f.Tag(), value), nil
}
*/
