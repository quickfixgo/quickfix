package quickfix

import (
	"bytes"
	"fmt"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	"time"
)

//Message is a FIX Message abstraction.
type Message struct {
	Header  FieldMap
	Trailer FieldMap
	Body    FieldMap

	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time

	rawMessage []byte

	//field bytes as they appear in the raw message
	fields []*fieldBytes
}

//parseError is returned when bytes cannot be parsed as a FIX message.
type parseError struct {
	OrigError string
}

func (e parseError) Error() string { return fmt.Sprintf("error parsing message: %s", e.OrigError) }

//ParseMessage constructs a Message from a byte slice wrapping a FIX message.
func ParseMessage(rawMessage []byte) (*Message, error) {
	msg := new(Message)
	msg.Header.init(headerFieldOrder)
	msg.Trailer.init(trailerFieldOrder)
	msg.Body.init(normalFieldOrder)
	msg.rawMessage = rawMessage

	//including required header and trailer fields, minimum of 7 fields can be expected
	//TODO: expose size for priming
	msg.fields = make([]*fieldBytes, 0, 7)

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
		case tag.IsHeader(parsedFieldBytes.Tag):
			msg.Header.append(parsedFieldBytes)
		case tag.IsTrailer(parsedFieldBytes.Tag):
			msg.Trailer.append(parsedFieldBytes)
		default:
			msg.Body.append(parsedFieldBytes)
		}
		if parsedFieldBytes.Tag == tag.CheckSum {
			break
		}
	}

	length := 0
	for _, field := range msg.fields {
		switch field.Tag {
		case tag.BeginString, tag.BodyLength, tag.CheckSum: //tags do not contribute to length
		default:
			length += field.Length()
		}
	}

	bodyLength := new(fix.IntValue)
	msg.Header.GetField(tag.BodyLength, bodyLength)
	if bodyLength.Value != length {
		return msg, parseError{OrigError: fmt.Sprintf("Incorrect Message Length, expected %d, got %d", bodyLength.Value, length)}
	}

	return msg, nil
}

//Bytes returns the raw bytes of the Message
func (m *Message) Bytes() []byte {
	return m.rawMessage
}

//ReverseRoute returns a message builder with routing header fields initialized as the reverse of this message.
func (m *Message) ReverseRoute() MessageBuilder {
	reverseBuilder := NewMessageBuilder()

	copy := func(src fix.Tag, dest fix.Tag) {
		if field := new(fix.StringValue); m.Header.GetField(src, field) == nil {
			if len(field.Value) != 0 {
				reverseBuilder.Header().SetField(dest, field)
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
	if beginString := new(fix.StringValue); m.Header.GetField(tag.BeginString, beginString) == nil {
		if beginString.Value != fix.BeginString_FIX40 {
			copy(tag.OnBehalfOfLocationID, tag.DeliverToLocationID)
			copy(tag.DeliverToLocationID, tag.OnBehalfOfLocationID)
		}
	}

	return reverseBuilder
}

func extractSpecificField(expectedTag fix.Tag, buffer []byte) (field *fieldBytes, remBuffer []byte, err error) {
	field, remBuffer, err = extractField(buffer)
	switch {
	case err != nil:
		return
	case field.Tag != expectedTag:
		err = parseError{OrigError: fmt.Sprintf("extractSpecificField: Fields out of order, expected %d, got %d", expectedTag, field.Tag)}
		return
	}

	return
}

func extractField(buffer []byte) (parsedFieldBytes *fieldBytes, remBytes []byte, err error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		err = parseError{OrigError: "extractField: No Trailing Delim in " + string(buffer)}
		remBytes = buffer
		return
	}

	parsedFieldBytes, err = parseField(buffer[:endIndex+1])
	return parsedFieldBytes, buffer[(endIndex + 1):], err
}

func (m *Message) String() string {
	return string(m.rawMessage)
}

//ToBuilder returns a writable message builder initialized with the fields in the message
//FIXME: not safe with repeated groups
func (m *Message) ToBuilder() MessageBuilder {
	builder := NewMessageBuilder()
	for tag := range m.Header.fieldLookup {
		builder.Header().fieldLookup[tag] = m.Header.fieldLookup[tag]
	}

	for tag := range m.Body.fieldLookup {
		builder.Body().fieldLookup[tag] = m.Body.fieldLookup[tag]
	}

	for tag := range m.Trailer.fieldLookup {
		builder.Trailer().fieldLookup[tag] = m.Trailer.fieldLookup[tag]
	}

	return builder
}

func newCheckSum(value int) *fix.StringField {
	return fix.NewStringField(tag.CheckSum, fmt.Sprintf("%03d", value))
}

//Free is required for Buffer interface FIXME
func (m *Message) Free() {

}
