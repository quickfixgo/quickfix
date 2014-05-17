package quickfix

//Buffer is a container for message bytes.  The interface provides a free() method that will be called after the buffer is consumed.
type buffer interface {
	//Called when the Buffer can be disposed by the Buffer creator
	free()
	Bytes() []byte
}
