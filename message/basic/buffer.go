package basic

//Simple implementation of message.Buffer 
type Buffer []byte

func (buffer Buffer) Bytes() []byte {
	return buffer
}

func (buffer Buffer) Free() {}
