package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Buffer is a container for message bytes.  The interface provides a Free() method that will be called after the buffer is consumed.
type Buffer interface {
	//Called when the Buffer can be disposed by the Buffer creator
	Free()
	Bytes() []byte
}

type MessageBuilder struct {
	Header
	Trailer
	Body FieldMapBuilder
}

func NewMessageBuilder() *MessageBuilder {
	m := new(MessageBuilder)
	m.Header.init()
	m.Trailer.init()
	m.Body.init(normalFieldOrder)
	return m
}

func (m MessageBuilder) Build() (*Message, error) {
	m.cook()

	var b bytes.Buffer
	m.Header.Write(&b)
	m.Body.Write(&b)
	m.Trailer.Write(&b)

	rawMessage := b.Bytes()
	return MessageFromParsedBytes(rawMessage)
}

func (m MessageBuilder) cook() {
	bodyLength := m.Header.length() + m.Body.length() + m.Trailer.length()
	checkSum := (m.Header.total() + m.Body.total() + m.Trailer.total()) % 256
	m.Header.Set(field.NewIntField(tag.BodyLength, bodyLength))
	m.Trailer.setCheckSum(newCheckSum(checkSum))
}
