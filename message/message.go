//message package contains types and functions specific to generic fix messages
package message

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/fix"
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
	var field Field
	var err error

	if field, rawMessage, err = extractSpecificField(tag.BeginString, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.SetField(field)
	}

	if field, rawMessage, err = extractSpecificField(tag.BodyLength, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.SetField(field)
	}

	if field, rawMessage, err = extractSpecificField(tag.MsgType, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.SetField(field)
	}

	for {
		field, rawMessage, err = extractField(rawMessage)

		switch {
		case err != nil:
			return nil, err
		case fix.IsHeader(field.Tag()):
			msg.Header.SetField(field)
		case fix.IsTrailer(field.Tag()):
			msg.Trailer.SetField(field)
		default:
			msg.Body.SetField(field)
		}
		if field.Tag() == tag.CheckSum {
			break
		}
	}

	msg.length = msg.Header.Length() + msg.Body.Length() + msg.Trailer.Length()
	msgLength, _ := msg.Header.IntValue(tag.BodyLength)
	if msgLength != msg.length {
		return msg, ParseError{fmt.Sprintf("Incorrect Message Length, expected %d, got %d", msgLength, msg.length)}
	}

	return msg, nil
}

func extractSpecificField(expectedTag tag.Tag, buffer []byte) (field Field, remBuffer []byte, err error) {
	field, remBuffer, err = extractField(buffer)
	switch {
	case err != nil:
		return
	case field.Tag() != expectedTag:
		return field, remBuffer, ParseError{fmt.Sprintf("extractSpecificField: Fields out of order, expected %d, got %d", expectedTag, field.Tag())}
	}

	return
}

func extractField(buffer []byte) (Field, []byte, error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		return nil, buffer, ParseError{"extractField: No Trailing Delim in " + string(buffer)}
	}

	field, err := ParseField(buffer[:endIndex+1])
	return field, buffer[(endIndex + 1):], err
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
	bodyLength := m.Header.Length() + m.Body.Length() + m.Trailer.Length()
	checkSum := (m.Header.Total() + m.Body.Total() + m.Trailer.Total()) % 256
	m.Header.SetField(NewIntField(tag.BodyLength, bodyLength))
	m.Trailer.setCheckSum(*newCheckSum(checkSum))
}

func newCheckSum(value int) *StringField {
	return NewStringField(tag.CheckSum, fmt.Sprintf("%03d", value))
}

func (message *Message) SetHeaderField(field Field) {
	message.Header.SetField(field)
}
