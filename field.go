package quickfix

//FieldValueWriter is an interface for writing field values
type FieldValueWriter interface {
	//Writes out the contents of the FieldValue to a []byte
	Write() []byte
}

//FieldValueReader is an interface for reading field values
type FieldValueReader interface {
	//Parameter to Read is TagValues.  For most fields, only the first TagValue will be required.
	//The length of the slice extends from the TagValue mapped to the field to be read through the
	//following fields. This can be useful for FieldValues made up of repeating groups.
	//
	//The Read function returns the remaining TagValues not processed by the FieldValue. If there was a
	//problem reading the field, an error may be returned
	Read(TagValues) (TagValues, error)
}

//The FieldValue interface is used to write/extract typed field values to/from raw bytes
type FieldValue interface {
	FieldValueWriter
	FieldValueReader

	//Returns a new FieldValue cloned from this one
	Clone() FieldValue
}

//FieldWriter is an interface for a writing a field
type FieldWriter interface {
	Tag() Tag
	FieldValueWriter
}

//Field is the interface implemented by all typed Fields in a Message
type Field interface {
	FieldWriter
	FieldValueReader
}
