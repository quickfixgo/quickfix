package quickfix

//Simple implementation of message.Buffer
type basicBuffer []byte

func (buffer basicBuffer) Bytes() []byte {
	return buffer
}

func (buffer basicBuffer) free() {}
