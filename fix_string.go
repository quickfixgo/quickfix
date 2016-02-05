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

func (f *FIXString) Read(bytes []byte) (err error) {
	*f = FIXString(bytes)
	return
}

func (f FIXString) Write() []byte {
	return []byte(f)
}
