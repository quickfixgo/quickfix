package quickfix

//FIXString is a FIX String Value, implements FieldValue
type FIXString string

//NewFIXString returns an initialized FIXString
func NewFIXString(val string) *FIXString {
	s := FIXString(val)
	return &s
}

func (f FIXString) String() string {
	return string(f)
}

func (f *FIXString) Read(tv TagValues) (TagValues, error) {
	*f = FIXString(tv[0].Value)
	return tv[1:], nil
}

func (f FIXString) Write() []byte {
	return []byte(f)
}

func (f FIXString) Clone() FieldValue {
	clone := f
	return &clone
}
