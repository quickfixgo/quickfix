package quickfix

//FIXBytes is a generic FIX field value, implements FieldValue.  Enables zero copy read from a FieldMap
type FIXBytes []byte

func (f *FIXBytes) Read(bytes []byte) (err error) {
	*f = FIXBytes(bytes)
	return
}

func (f FIXBytes) Write() []byte {
	return []byte(f)
}
