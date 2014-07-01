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
	Build() ([]byte, error)
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

func (m messageBuilder) Build() ([]byte, error) {
	m.cook()

	var b bytes.Buffer
	m.header.write(&b)
	m.body.write(&b)
	m.trailer.write(&b)

	return b.Bytes(), nil
}

func (m messageBuilder) cook() {
	bodyLength := m.header.length() + m.body.length() + m.trailer.length()
	m.header.Set(fix.NewIntField(tag.BodyLength, bodyLength))
	checkSum := (m.header.total() + m.body.total() + m.trailer.total()) % 256
	m.trailer.Set(newCheckSum(checkSum))
}
