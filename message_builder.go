package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
)

type MessageBuilder struct {
	Header  FieldMap
	Trailer FieldMap
	Body    FieldMap
}

func NewMessageBuilder() *MessageBuilder {
	header := fieldMap{}
	header.init(headerFieldOrder)
	body := fieldMap{}
	body.init(normalFieldOrder)
	trailer := fieldMap{}
	trailer.init(normalFieldOrder)

	return &MessageBuilder{Header: header, Body: body, Trailer: trailer}
}

func (m MessageBuilder) Build() ([]byte, error) {
	m.cook()

	var b bytes.Buffer
	m.Header.write(&b)
	m.Body.write(&b)
	m.Trailer.write(&b)

	return b.Bytes(), nil
}

func (m MessageBuilder) cook() {
	bodyLength := m.Header.length() + m.Body.length() + m.Trailer.length()
	m.Header.Set(fix.NewIntField(tag.BodyLength, bodyLength))
	checkSum := (m.Header.total() + m.Body.total() + m.Trailer.total()) % 256
	m.Trailer.Set(newCheckSum(checkSum))
}
