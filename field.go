package quickfix

//The FieldValue interface is used to write/extract typed field values to/from raw bytes
type FieldValue interface {
	Write() []byte

	//Parameter to Read is a single length slice of TagValues.  For most fields, only the first
	//TagValue will be required. The capacity of the slice extends from the field to be read through the
	//underlying message length. This can be useful for FieldValues made up of repeating groups.
	Read([]TagValue) error
}

//Field is the interface implemented by all typed Fields in a Message
type Field interface {
	Tag() Tag
	FieldValue
}
