package quickfix

import (
	"bytes"
	"fmt"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Message is a FIX Message abstraction.
type Message struct {
	Header  FieldMap
	Trailer FieldMap
	Body    FieldMap

	rawMessage []byte

	//field bytes as they appear in the raw message
	fields []*fieldBytes
}

//MessageFromParsedBytes constructs a Message from a byte slice wrapping a FIX message
func MessageFromParsedBytes(rawMessage []byte) (*Message, error) {
	msg := new(Message)
	msg.Header.init()
	msg.Trailer.init()
	msg.Body.init()
	msg.rawMessage = rawMessage

	//BeginString, BodyLength, MsgType, CheckSum - at least 4
	msg.fields = make([]*fieldBytes, 0, 4)

	var parsedFieldBytes *fieldBytes
	var err error

	//message must start with begin string, body length, msg type
	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.BeginString, rawMessage); err != nil {
		return nil, err
	}

	msg.Header.append(parsedFieldBytes)
	msg.fields = append(msg.fields, parsedFieldBytes)

	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.BodyLength, rawMessage); err != nil {
		return nil, err
	}

	msg.Header.append(parsedFieldBytes)
	msg.fields = append(msg.fields, parsedFieldBytes)

	if parsedFieldBytes, rawMessage, err = extractSpecificField(tag.MsgType, rawMessage); err != nil {
		return nil, err
	}

	msg.Header.append(parsedFieldBytes)
	msg.fields = append(msg.fields, parsedFieldBytes)

	for {
		parsedFieldBytes, rawMessage, err = extractField(rawMessage)
		if err != nil {
			return nil, err
		}

		msg.fields = append(msg.fields, parsedFieldBytes)

		switch {
		case parsedFieldBytes.Tag.IsHeader():
			msg.Header.append(parsedFieldBytes)
		case parsedFieldBytes.Tag.IsTrailer():
			msg.Trailer.append(parsedFieldBytes)
		default:
			msg.Body.append(parsedFieldBytes)
		}
		if parsedFieldBytes.Tag == tag.CheckSum {
			break
		}
	}

	length := msg.Header.length() + msg.Body.length() + msg.Trailer.length()
	bodyLength := new(field.IntValue)
	msg.Header.GetField(tag.BodyLength, bodyLength)
	if bodyLength.Value != length {
		return msg, ParseError{fmt.Sprintf("Incorrect Message Length, expected %d, got %d", bodyLength.Value, length)}
	}

	return msg, nil
}

//Bytes returns the raw bytes of the Message
func (m Message) Bytes() []byte {
	return m.rawMessage
}

//ReverseRoute returns a message builder with routing header fields initialized as the reverse of this message.
func (m Message) ReverseRoute() *MessageBuilder {
	reverseBuilder := NewMessageBuilder()

	copy := func(src tag.Tag, dest tag.Tag) {
		if field := new(field.StringValue); m.Header.GetField(src, field) == nil {
			if len(field.Value) != 0 {
				reverseBuilder.Header.SetField(dest, field)
			}
		}
	}

	copy(tag.SenderCompID, tag.TargetCompID)
	copy(tag.SenderSubID, tag.TargetSubID)
	copy(tag.SenderLocationID, tag.TargetLocationID)

	copy(tag.TargetCompID, tag.SenderCompID)
	copy(tag.TargetSubID, tag.SenderSubID)
	copy(tag.TargetLocationID, tag.SenderLocationID)

	copy(tag.OnBehalfOfCompID, tag.DeliverToCompID)
	copy(tag.OnBehalfOfSubID, tag.DeliverToSubID)
	copy(tag.DeliverToCompID, tag.OnBehalfOfCompID)
	copy(tag.DeliverToSubID, tag.OnBehalfOfSubID)

	//tags added in 4.1
	if beginString := new(field.StringValue); m.Header.GetField(tag.BeginString, beginString) == nil {
		if beginString.Value != BeginString_FIX40 {
			copy(tag.OnBehalfOfLocationID, tag.DeliverToLocationID)
			copy(tag.DeliverToLocationID, tag.OnBehalfOfLocationID)
		}
	}

	return reverseBuilder
}

func extractSpecificField(expectedTag tag.Tag, buffer []byte) (field *fieldBytes, remBuffer []byte, err error) {
	field, remBuffer, err = extractField(buffer)
	switch {
	case err != nil:
		return
	case field.Tag != expectedTag:
		err = ParseError{fmt.Sprintf("extractSpecificField: Fields out of order, expected %d, got %d", expectedTag, field.Tag)}
		return
	}

	return
}

func extractField(buffer []byte) (parsedFieldBytes *fieldBytes, remBytes []byte, err error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		err = ParseError{"extractField: No Trailing Delim in " + string(buffer)}
		remBytes = buffer
		return
	}

	parsedFieldBytes, err = parseField(buffer[:endIndex+1])
	return parsedFieldBytes, buffer[(endIndex + 1):], err
}

func (m Message) String() string {
	return string(m.rawMessage)
}

//ToBuilder returns a writable message builder initialized with the fields in the message
//FIXME: not safe with repeated groups
func (m Message) ToBuilder() *MessageBuilder {
	builder := NewMessageBuilder()
	for tag := range m.Header.lookup {
		builder.Header.fields[tag] = m.Header.lookup[tag]
	}

	for tag := range m.Body.lookup {
		builder.Body.fields[tag] = m.Body.lookup[tag]
	}

	for tag := range m.Trailer.lookup {
		builder.Trailer.fields[tag] = m.Trailer.lookup[tag]
	}

	return builder
}

func newCheckSum(value int) *field.StringField {
	return field.NewStringField(tag.CheckSum, fmt.Sprintf("%03d", value))
}

//Free is required for Buffer interface FIXME
func (m Message) Free() {

}
