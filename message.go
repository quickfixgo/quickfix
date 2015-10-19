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

	//slice of Bytes corresponding to the message body
	bodyBytes []byte

	//field bytes as they appear in the raw message
	fields []fieldBytes
}

//parseError is returned when bytes cannot be parsed as a FIX message.
type parseError struct {
	OrigError string
}

func (e parseError) Error() string { return fmt.Sprintf("error parsing message: %s", e.OrigError) }

//NewMessage returns a newly initialized Message instance
func NewMessage() (m Message) {
	m.Init()
	return
}

//Init initializes the Message instance
func (m *Message) Init() {
	var header, body, trailer fieldMap

	header.init(headerFieldOrder)
	body.init(normalFieldOrder)
	trailer.init(trailerFieldOrder)

	m.Header = header
	m.Body = body
	m.Trailer = trailer
}

//parseMessage constructs a Message from a byte slice wrapping a FIX message.
func parseMessage(rawMessage []byte) (*Message, error) {
	msg := &Message{rawMessage: rawMessage}
	var header, body, trailer fieldMap

	header.init(headerFieldOrder)
	body.init(normalFieldOrder)
	trailer.init(trailerFieldOrder)

	msg.Header = header
	msg.Body = body
	msg.Trailer = trailer

	//allocate fields in one chunk
	fieldCount := 0
	for _, b := range rawMessage {
		if b == '\001' {
			fieldCount++
		}
	}
	msg.fields = make([]fieldBytes, fieldCount)

	fieldIndex := 0
	var err error

	//message must start with begin string, body length, msg type
	if rawMessage, err = extractSpecificField(&msg.fields[fieldIndex], tag.BeginString, rawMessage); err != nil {
		return nil, err
	}

	header.fieldLookup[msg.fields[fieldIndex].Tag] = &msg.fields[fieldIndex]
	fieldIndex++

	parsedFieldBytes := &msg.fields[fieldIndex]
	if rawMessage, err = extractSpecificField(parsedFieldBytes, tag.BodyLength, rawMessage); err != nil {
		return nil, err
	}

	header.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
	fieldIndex++

	parsedFieldBytes = &msg.fields[fieldIndex]
	if rawMessage, err = extractSpecificField(parsedFieldBytes, tag.MsgType, rawMessage); err != nil {
		return nil, err
	}

	header.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
	fieldIndex++

	trailerBytes := []byte{}
	foundBody := false
	for {
		parsedFieldBytes = &msg.fields[fieldIndex]
		rawMessage, err = extractField(parsedFieldBytes, rawMessage)
		if err != nil {
			return nil, err
		}

		switch {
		case tag.IsHeader(parsedFieldBytes.Tag):
			header.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
		case tag.IsTrailer(parsedFieldBytes.Tag):
			trailer.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
		default:
			foundBody = true
			trailerBytes = rawMessage
			body.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
		}
		if parsedFieldBytes.Tag == tag.CheckSum {
			break
		}

		if !foundBody {
			msg.bodyBytes = rawMessage
		}

		fieldIndex++
	}

	//body length would only be larger than trailer if fields out of order
	if len(msg.bodyBytes) > len(trailerBytes) {
		msg.bodyBytes = msg.bodyBytes[:len(msg.bodyBytes)-len(trailerBytes)]
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

//reverseRoute returns a message builder with routing header fields initialized as the reverse of this message.
func (m Message) reverseRoute() Message {
	reverseMsg := NewMessage()

	copy := func(src fix.Tag, dest fix.Tag) {
		if field := new(fix.StringValue); m.Header.GetField(src, field) == nil {
			if len(field.Value) != 0 {
				reverseMsg.Header.SetField(dest, field)
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

	return reverseMsg
}

func extractSpecificField(field *fieldBytes, expectedTag fix.Tag, buffer []byte) (remBuffer []byte, err error) {
	remBuffer, err = extractField(field, buffer)
	switch {
	case err != nil:
		return
	case field.Tag != expectedTag:
		err = parseError{OrigError: fmt.Sprintf("extractSpecificField: Fields out of order, expected %d, got %d", expectedTag, field.Tag)}
		return
	}

	return
}

func extractField(parsedFieldBytes *fieldBytes, buffer []byte) (remBytes []byte, err error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		err = parseError{OrigError: "extractField: No Trailing Delim in " + string(buffer)}
		remBytes = buffer
		return
	}

	err = parsedFieldBytes.parseField(buffer[:endIndex+1])
	return buffer[(endIndex + 1):], err
}

func (m *Message) String() string {
	return string(m.rawMessage)
}

func newCheckSum(value int) *fix.StringField {
	return fix.NewStringField(tag.CheckSum, fmt.Sprintf("%03d", value))
}

func (m *Message) Build() ([]byte, error) {
	m.cook()

	var b bytes.Buffer
	m.Header.write(&b)
	m.Body.write(&b)
	m.Trailer.write(&b)

	m.rawMessage = b.Bytes()

	return m.rawMessage, nil
}

func (m Message) cook() {
	bodyLength := m.Header.length() + m.Body.length() + m.Trailer.length()
	m.Header.Set(fix.NewIntField(tag.BodyLength, bodyLength))
	checkSum := (m.Header.total() + m.Body.total() + m.Trailer.total()) % 256
	m.Trailer.Set(newCheckSum(checkSum))
}
