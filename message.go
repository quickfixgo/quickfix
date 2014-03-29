package quickfixgo

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
)

type Message struct {
	Header  FieldMap
	Trailer FieldMap
	Body    FieldMap

	rawMessage []byte
}

func MessageFromParsedBytes(rawMessage []byte) (*Message, error) {
	msg := new(Message)
	msg.Header.init()
	msg.Trailer.init()
	msg.Body.init()
	msg.rawMessage = rawMessage

	var parsedFieldBytes *fieldBytes
	var err error

	//message must start with begin string, body length, msg type
	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.BeginString, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.append(tag.BeginString, parsedFieldBytes)
	}

	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.BodyLength, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.append(tag.BodyLength, parsedFieldBytes)
	}

	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.MsgType, rawMessage); err != nil {
		return nil, err
	} else {
		msg.Header.append(tag.MsgType, parsedFieldBytes)
	}

	var fieldTag tag.Tag
	for {
		fieldTag, parsedFieldBytes, rawMessage, err = extractField(rawMessage)

		switch {
		case err != nil:
			return nil, err
		case fieldTag.IsHeader():
			msg.Header.append(fieldTag, parsedFieldBytes)
		case fieldTag.IsTrailer():
			msg.Trailer.append(fieldTag, parsedFieldBytes)
		default:
			msg.Body.append(fieldTag, parsedFieldBytes)
		}
		if fieldTag == tag.CheckSum {
			break
		}
	}

	length := msg.Header.length() + msg.Body.length() + msg.Trailer.length()
	bodyLength := new(IntValue)
	msg.Header.GetField(tag.BodyLength, bodyLength)
	if bodyLength.Value != length {
		return msg, ParseError{fmt.Sprintf("Incorrect Message Length, expected %d, got %d", bodyLength.Value, length)}
	}

	return msg, nil
}

func (m Message) Bytes() []byte {
	return m.rawMessage
}

func extractSpecificField(expectedTag tag.Tag, buffer []byte) (field *fieldBytes, remBuffer []byte, err error) {
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

func extractField(buffer []byte) (tag tag.Tag, parsedFieldBytes *fieldBytes, remBytes []byte, err error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		err = ParseError{"extractField: No Trailing Delim in " + string(buffer)}
		remBytes = buffer
		return
	}

	tag, parsedFieldBytes, err = parseField(buffer[:endIndex+1])
	return tag, parsedFieldBytes, buffer[(endIndex + 1):], err
}

func (m Message) String() string {
	return string(m.rawMessage)
}

func (m Message) ToBuilder() *MessageBuilder {
	builder := NewMessageBuilder()
	for tag := range m.Header.fields {
		builder.Header.fields[tag] = m.Header.fields[tag]
	}

	for tag := range m.Body.fields {
		builder.Body.fields[tag] = m.Body.fields[tag]
	}

	for tag := range m.Trailer.fields {
		builder.Trailer.fields[tag] = m.Trailer.fields[tag]
	}

	return builder
}

func newCheckSum(value int) *StringField {
	return NewStringField(tag.CheckSum, fmt.Sprintf("%03d", value))
}

func (m *Message) Free() {

}
