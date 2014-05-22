package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
)

type MessageBuilder interface {
	Header() MutableFieldMap
	Trailer() MutableFieldMap
	Body() MutableFieldMap
	Build() (*Message, error)
}

type messageBuilder struct {
	header  fieldMap
	trailer fieldMap
	body    fieldMap
}

func NewMessageBuilder() MessageBuilder {
	m := &messageBuilder{}
	m.header.init(headerFieldOrder)
	m.trailer.init(trailerFieldOrder)
	m.body.init(normalFieldOrder)
	return m
}

func (m messageBuilder) Header() MutableFieldMap  { return m.header }
func (m messageBuilder) Trailer() MutableFieldMap { return m.trailer }
func (m messageBuilder) Body() MutableFieldMap    { return m.body }

func (m messageBuilder) Build() (*Message, error) {
	m.cook()

	var b bytes.Buffer
	m.header.write(&b)
	m.body.write(&b)
	m.trailer.write(&b)

	rawMessage := b.Bytes()
	return parseMessage(rawMessage)
}

func (m messageBuilder) cook() {
	bodyLength := m.header.length() + m.body.length() + m.trailer.length()
	checkSum := (m.header.total() + m.body.total() + m.trailer.total()) % 256
	m.header.Set(fix.NewIntField(tag.BodyLength, bodyLength))
	m.trailer.Set(newCheckSum(checkSum))
}
