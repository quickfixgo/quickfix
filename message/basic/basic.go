//Package basic implements very simple implementations of the message 
package basic

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/tag"
)

//Message implements the message.Message and message.Builder interface.
type Message struct {
	MsgHeader  Header
	MsgTrailer Trailer
	MsgBody    FieldMap

	length int
}

func NewMessage() *Message {
	m := new(Message)
	m.MsgHeader.init()
	m.MsgTrailer.init()
	m.MsgBody.init(normalFieldOrder)
	return m
}

func MessageFromParsedBytes(rawMessage []byte) (*Message, error) {
	//message must start with begin string, body length, msg type
	msg := NewMessage()
	var field message.Field
	var err error

	if field, rawMessage, err = extractSpecificField(tag.BeginString, rawMessage); err != nil {
		return nil, err
	} else {
		msg.MsgHeader.SetField(field)
	}

	if field, rawMessage, err = extractSpecificField(tag.BodyLength, rawMessage); err != nil {
		return nil, err
	} else {
		msg.MsgHeader.SetField(field)
	}

	if field, rawMessage, err = extractSpecificField(tag.MsgType, rawMessage); err != nil {
		return nil, err
	} else {
		msg.MsgHeader.SetField(field)
	}

	for {
		field, rawMessage, err = extractField(rawMessage)

		switch {
		case err != nil:
			return nil, err
		case fix.IsHeader(field.Tag()):
			msg.MsgHeader.SetField(field)
		case fix.IsTrailer(field.Tag()):
			msg.MsgTrailer.SetField(field)
		default:
			msg.MsgBody.SetField(field)
		}
		if field.Tag() == tag.CheckSum {
			break
		}
	}

	msg.length = msg.MsgHeader.Length() + msg.MsgBody.Length() + msg.MsgTrailer.Length()
	msgLength, _ := msg.MsgHeader.IntValue(tag.BodyLength)
	if msgLength != msg.length {
		return msg, message.ParseError{fmt.Sprintf("Incorrect Message Length, expected %d, got %d", msgLength, msg.length)}
	}

	return msg, nil
}

func extractSpecificField(expectedTag tag.Tag, buffer []byte) (field message.Field, remBuffer []byte, err error) {
	field, remBuffer, err = extractField(buffer)
	switch {
	case err != nil:
		return
	case field.Tag() != expectedTag:
		return field, remBuffer, message.ParseError{fmt.Sprintf("extractSpecificField: Fields out of order, expected %d, got %d", expectedTag, field.Tag())}
	}

	return
}

func extractField(buffer []byte) (message.Field, []byte, error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		return nil, buffer, message.ParseError{"extractField: No Trailing Delim in " + string(buffer)}
	}

	field, err := ParseField(buffer[:endIndex+1])
	return field, buffer[(endIndex + 1):], err
}

func (m *Message) Header() message.FieldMap {
	return &m.MsgHeader
}

func (m *Message) Trailer() message.FieldMap {
	return &m.MsgTrailer
}

func (m *Message) Body() message.FieldMap {
	return &m.MsgBody
}

func (m *Message) Length() int {
	return m.length
}

func (message *Message) String() string {
	var b bytes.Buffer
	message.MsgHeader.Write(&b)
	message.MsgBody.Write(&b)
	message.MsgTrailer.Write(&b)

	return b.String()
}

func (m *Message) Build() message.Buffer {
	m.cook()

	var b bytes.Buffer
	m.MsgHeader.Write(&b)
	m.MsgBody.Write(&b)
	m.MsgTrailer.Write(&b)

	return Buffer(b.Bytes())
}

func (m *Message) cook() {
	bodyLength := m.MsgHeader.Length() + m.MsgBody.Length() + m.MsgTrailer.Length()
	checkSum := (m.MsgHeader.Total() + m.MsgBody.Total() + m.MsgTrailer.Total()) % 256
	m.MsgHeader.SetField(NewIntField(tag.BodyLength, bodyLength))
	m.MsgTrailer.setCheckSum(newCheckSum(checkSum))
}

func newCheckSum(value int) *StringField {
	return NewStringField(tag.CheckSum, fmt.Sprintf("%03d", value))
}

func (message *Message) SetHeaderField(field message.Field) {
	message.MsgHeader.SetField(field)
}
