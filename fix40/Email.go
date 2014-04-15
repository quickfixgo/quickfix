package fix40

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Email msg type = C.
type Email struct {
	message.Message
}

//EmailBuilder builds Email messages.
type EmailBuilder struct {
	message.MessageBuilder
}

//CreateEmailBuilder returns an initialized EmailBuilder with specified required fields.
func CreateEmailBuilder(
	emailtype field.EmailType,
	linesoftext field.LinesOfText,
	text field.Text) EmailBuilder {
	var builder EmailBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(emailtype)
	builder.Body.Set(linesoftext)
	builder.Body.Set(text)
	return builder
}

//EmailType is a required field for Email.
func (m Email) EmailType() (field.EmailType, error) {
	var f field.EmailType
	err := m.Body.Get(&f)
	return f, err
}

//OrigTime is a non-required field for Email.
func (m Email) OrigTime() (field.OrigTime, error) {
	var f field.OrigTime
	err := m.Body.Get(&f)
	return f, err
}

//RelatdSym is a non-required field for Email.
func (m Email) RelatdSym() (field.RelatdSym, error) {
	var f field.RelatdSym
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for Email.
func (m Email) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for Email.
func (m Email) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//LinesOfText is a required field for Email.
func (m Email) LinesOfText() (field.LinesOfText, error) {
	var f field.LinesOfText
	err := m.Body.Get(&f)
	return f, err
}

//Text is a required field for Email.
func (m Email) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//RawDataLength is a non-required field for Email.
func (m Email) RawDataLength() (field.RawDataLength, error) {
	var f field.RawDataLength
	err := m.Body.Get(&f)
	return f, err
}

//RawData is a non-required field for Email.
func (m Email) RawData() (field.RawData, error) {
	var f field.RawData
	err := m.Body.Get(&f)
	return f, err
}
