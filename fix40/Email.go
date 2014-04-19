package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
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
func (m Email) EmailType() (field.EmailType, errors.MessageRejectError) {
	var f field.EmailType
	err := m.Body.Get(&f)
	return f, err
}

//OrigTime is a non-required field for Email.
func (m Email) OrigTime() (field.OrigTime, errors.MessageRejectError) {
	var f field.OrigTime
	err := m.Body.Get(&f)
	return f, err
}

//RelatdSym is a non-required field for Email.
func (m Email) RelatdSym() (field.RelatdSym, errors.MessageRejectError) {
	var f field.RelatdSym
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for Email.
func (m Email) OrderID() (field.OrderID, errors.MessageRejectError) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for Email.
func (m Email) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//LinesOfText is a required field for Email.
func (m Email) LinesOfText() (field.LinesOfText, errors.MessageRejectError) {
	var f field.LinesOfText
	err := m.Body.Get(&f)
	return f, err
}

//Text is a required field for Email.
func (m Email) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//RawDataLength is a non-required field for Email.
func (m Email) RawDataLength() (field.RawDataLength, errors.MessageRejectError) {
	var f field.RawDataLength
	err := m.Body.Get(&f)
	return f, err
}

//RawData is a non-required field for Email.
func (m Email) RawData() (field.RawData, errors.MessageRejectError) {
	var f field.RawData
	err := m.Body.Get(&f)
	return f, err
}
