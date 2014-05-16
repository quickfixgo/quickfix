package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
)

type MessageBuilder interface {
	Header() FieldMap
	Trailer() FieldMap
	Body() FieldMap
	Build() (*Message, error)
}

type messageBuilder struct {
	header  FieldMap
	trailer FieldMap
	body    FieldMap
}

func NewMessageBuilder() MessageBuilder {
	m := &messageBuilder{}
	m.header.init(headerFieldOrder)
	m.trailer.init(trailerFieldOrder)
	m.body.init(normalFieldOrder)
	return m
}

func (m messageBuilder) Header() FieldMap  { return m.header }
func (m messageBuilder) Trailer() FieldMap { return m.trailer }
func (m messageBuilder) Body() FieldMap    { return m.body }

func (m messageBuilder) Build() (*Message, error) {
	m.cook()

	var b bytes.Buffer
	m.header.Write(&b)
	m.body.Write(&b)
	m.trailer.Write(&b)

	rawMessage := b.Bytes()
	return ParseMessage(rawMessage)
}

func (m messageBuilder) cook() {
	bodyLength := m.header.length() + m.body.length() + m.trailer.length()
	checkSum := (m.header.total() + m.body.total() + m.trailer.total()) % 256
	m.header.Set(fix.NewIntField(tag.BodyLength, bodyLength))
	m.trailer.Set(newCheckSum(checkSum))
}
