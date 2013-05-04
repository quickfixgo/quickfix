package quickfixgo

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/field"
	"github.com/cbusbey/quickfixgo/tag"
)

type Message struct {
	Header
	Trailer
	Body FieldMap

	length int
}

func NewMessage() *Message {
	m := new(Message)
	m.Header.init()
	m.Trailer.init()
	m.Body.init(normalFieldOrder)
	return m
}

func MessageFromParsedBytes(rawMessage []byte) (*Message, error) {
	//message must start with begin string, body length, msg type
	msg := NewMessage()
	var parsedFieldBytes *fieldValue
	var err error

	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.BeginString, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.fields[tag.BeginString] = parsedFieldBytes
	}

	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.BodyLength, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.fields[tag.BodyLength] = parsedFieldBytes
	}

	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.MsgType, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.fields[tag.MsgType] = parsedFieldBytes
	}

	var fieldTag tag.Tag
	for {
		fieldTag, parsedFieldBytes, rawMessage, err = extractField(rawMessage)

		switch {
		case err != nil:
			return nil, err
		case fieldTag.IsHeader():
			msg.Header.fields[fieldTag] = parsedFieldBytes
		case fieldTag.IsTrailer():
			msg.Trailer.fields[fieldTag] = parsedFieldBytes
		default:
			msg.Body.fields[fieldTag] = parsedFieldBytes
		}
		if fieldTag == tag.CheckSum {
			break
		}
	}

	msg.length = msg.Header.length() + msg.Body.length() + msg.Trailer.length()
	bodyLength := new(field.BodyLength)
	msg.Header.Get(bodyLength)
	if bodyLength.Value != msg.length {
		return msg, ParseError{fmt.Sprintf("Incorrect Message Length, expected %d, got %d", bodyLength.Value, msg.length)}
	}

	return msg, nil
}

func extractSpecificField(expectedTag tag.Tag, buffer []byte) (field *fieldValue, remBuffer []byte, err error) {
	var tag tag.Tag
	tag, field, remBuffer, err = extractField(buffer)
	switch {
	case err != nil:
		return
	case tag != expectedTag:
		err = ParseError{fmt.Sprintf("extractSpecificField: Fields out of order, expected %d, got %d", expectedTag, tag)}
		return
	}

	return
}

func extractField(buffer []byte) (tag tag.Tag, parsedFieldBytes *fieldValue, remBytes []byte, err error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		err = ParseError{"extractField: No Trailing Delim in " + string(buffer)}
		remBytes = buffer
		return
	}

	tag, parsedFieldBytes, err = parseField(buffer[:endIndex+1])
	return tag, parsedFieldBytes, buffer[(endIndex + 1):], err
}

func (m *Message) Length() int {
	return m.length
}

func (message *Message) String() string {
	var b bytes.Buffer
	message.Header.Write(&b)
	message.Body.Write(&b)
	message.Trailer.Write(&b)

	return b.String()
}

func (m *Message) Build() Buffer {
	m.cook()

	var b bytes.Buffer
	m.Header.Write(&b)
	m.Body.Write(&b)
	m.Trailer.Write(&b)

	return BasicBuffer(b.Bytes())
}

func (m *Message) cook() {
	bodyLength := m.Header.length() + m.Body.length() + m.Trailer.length()
	checkSum := (m.Header.total() + m.Body.total() + m.Trailer.total()) % 256
	m.Header.SetField(field.NewIntField(tag.BodyLength, bodyLength))
	m.Trailer.setCheckSum(newCheckSum(checkSum))
}

func newCheckSum(value int) *field.StringField {
	return field.NewStringField(tag.CheckSum, fmt.Sprintf("%03d", value))
}

func (message *Message) SetHeaderField(field FieldConverter) {
	message.Header.SetField(field)
}
