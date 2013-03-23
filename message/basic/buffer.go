package basic

//Simple implementation of message.Buffer 
type Buffer struct {
	bytes []byte
}

func (buffer Buffer) Bytes() []byte {
	return buffer.bytes
}

func (buffer Buffer) Free() {}
