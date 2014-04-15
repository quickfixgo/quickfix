package message

import (
	"bytes"
	"github.com/quickfixgo/quickfix/fix/tag"
)

type MessageBuilder struct {
	Header
	Trailer
	Body FieldMapBuilder
}

func CreateMessageBuilder() MessageBuilder {
	var m MessageBuilder
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
	m.Header.Set(NewIntField(tag.BodyLength, bodyLength))
	m.Trailer.setCheckSum(newCheckSum(checkSum))
}
