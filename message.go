package quickfix

import (
	"bytes"
	"fmt"
	"github.com/quickfixgo/quickfix/enum"
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
	m.Header.init(headerFieldOrder)
	m.Body.init(normalFieldOrder)
	m.Trailer.init(trailerFieldOrder)
}

//parseMessage constructs a Message from a byte slice wrapping a FIX message.
func parseMessage(rawMessage []byte) (Message, error) {
	msg := NewMessage()
	msg.rawMessage = rawMessage

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
	if rawMessage, err = extractSpecificField(&msg.fields[fieldIndex], tagBeginString, rawMessage); err != nil {
		return msg, err
	}

	msg.Header.fieldLookup[msg.fields[fieldIndex].Tag] = &msg.fields[fieldIndex]
	fieldIndex++

	parsedFieldBytes := &msg.fields[fieldIndex]
	if rawMessage, err = extractSpecificField(parsedFieldBytes, tagBodyLength, rawMessage); err != nil {
		return msg, err
	}

	msg.Header.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
	fieldIndex++

	parsedFieldBytes = &msg.fields[fieldIndex]
	if rawMessage, err = extractSpecificField(parsedFieldBytes, tagMsgType, rawMessage); err != nil {
		return msg, err
	}

	msg.Header.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
	fieldIndex++

	trailerBytes := []byte{}
	foundBody := false
	for {
		parsedFieldBytes = &msg.fields[fieldIndex]
		rawMessage, err = extractField(parsedFieldBytes, rawMessage)
		if err != nil {
			return msg, err
		}

		switch {
		case parsedFieldBytes.Tag.IsHeader():
			msg.Header.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
		case parsedFieldBytes.Tag.IsTrailer():
			msg.Trailer.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
		default:
			foundBody = true
			trailerBytes = rawMessage
			msg.Body.fieldLookup[parsedFieldBytes.Tag] = parsedFieldBytes
		}
		if parsedFieldBytes.Tag == tagCheckSum {
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
		case tagBeginString, tagBodyLength, tagCheckSum: //tags do not contribute to length
		default:
			length += field.Length()
		}
	}

	bodyLength := new(IntValue)
	msg.Header.GetField(tagBodyLength, bodyLength)
	if bodyLength.Value != length {
		return msg, parseError{OrigError: fmt.Sprintf("Incorrect Message Length, expected %d, got %d", bodyLength.Value, length)}
	}

	return msg, nil
}

//reverseRoute returns a message builder with routing header fields initialized as the reverse of this message.
func (m Message) reverseRoute() Message {
	reverseMsg := NewMessage()

	copy := func(src Tag, dest Tag) {
		if field := new(StringValue); m.Header.GetField(src, field) == nil {
			if len(field.Value) != 0 {
				reverseMsg.Header.SetField(dest, field)
			}
		}
	}

	copy(tagSenderCompID, tagTargetCompID)
	copy(tagSenderSubID, tagTargetSubID)
	copy(tagSenderLocationID, tagTargetLocationID)

	copy(tagTargetCompID, tagSenderCompID)
	copy(tagTargetSubID, tagSenderSubID)
	copy(tagTargetLocationID, tagSenderLocationID)

	copy(tagOnBehalfOfCompID, tagDeliverToCompID)
	copy(tagOnBehalfOfSubID, tagDeliverToSubID)
	copy(tagDeliverToCompID, tagOnBehalfOfCompID)
	copy(tagDeliverToSubID, tagOnBehalfOfSubID)

	//tags added in 4.1
	if beginString := new(StringValue); m.Header.GetField(tagBeginString, beginString) == nil {
		if beginString.Value != enum.BeginStringFIX40 {
			copy(tagOnBehalfOfLocationID, tagDeliverToLocationID)
			copy(tagDeliverToLocationID, tagOnBehalfOfLocationID)
		}
	}

	return reverseMsg
}

func extractSpecificField(field *fieldBytes, expectedTag Tag, buffer []byte) (remBuffer []byte, err error) {
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

func newCheckSum(value int) *StringField {
	return NewStringField(tagCheckSum, fmt.Sprintf("%03d", value))
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
	m.Header.Set(NewIntField(tagBodyLength, bodyLength))
	checkSum := (m.Header.total() + m.Body.total() + m.Trailer.total()) % 256
	m.Trailer.Set(newCheckSum(checkSum))
}
