package quickfix

//FieldValueWriter is an interface for writing field values
type FieldValueWriter interface {
	//Writes out the contents of the FieldValue to a []byte
	Write() []byte
}

//FieldValueReader is an interface for reading field values
type FieldValueReader interface {
	//Reads the contents of the []byte into FieldValue.  Returns an error if there are issues in the data processing
	Read([]byte) error
}

//The FieldValue interface is used to write/extract typed field values to/from raw bytes
type FieldValue interface {
	FieldValueWriter
	FieldValueReader
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

//FieldGroupWriter is an interface for writing a FieldGroup
type FieldGroupWriter interface {
	Tag() Tag
	Write() []TagValue
}

//FieldGroupReader is an interface for reading a FieldGroup
type FieldGroupReader interface {
	Tag() Tag
	Read([]TagValue) ([]TagValue, error)
}

//FieldGroup is the interface implemented by all typed Groups in a Message
type FieldGroup interface {
	Tag() Tag
	Write() []TagValue
	Read([]TagValue) ([]TagValue, error)
}
