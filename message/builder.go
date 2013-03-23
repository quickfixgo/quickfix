package message

//Buffer is a container for message bytes.  The interface provides a Free() method that will be called after the buffer is consumed.
type Buffer interface {
	//Called when the Buffer can be disposed by the Buffer creator
	Free()
	Bytes() []byte
}

type Builder interface {
	SetHeaderField(Field)
	Build() Buffer
}
