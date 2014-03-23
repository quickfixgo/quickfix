package quickfixgo

//Simple implementation of message.Buffer
type BasicBuffer []byte

func (buffer BasicBuffer) Bytes() []byte {
	return buffer
}

func (buffer BasicBuffer) Free() {}
