package message

// Field is the common interface implemented by all field types
type Field interface {
	Tag() Tag

	//the value for this field
	Value() string

	//the string representation of the field (tag=value\001)
	String() string

	//the length of the field in bytes
	Length() int

	//the sum of the bytes in the field
	Total() int
}
