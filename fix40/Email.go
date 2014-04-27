package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.BuildMsgType("C"))
	builder.Body.Set(emailtype)
	builder.Body.Set(linesoftext)
	builder.Body.Set(text)
	return builder
}

//EmailType is a required field for Email.
func (m Email) EmailType() (*field.EmailType, errors.MessageRejectError) {
	f := new(field.EmailType)
	err := m.Body.Get(f)
	return f, err
}

//GetEmailType reads a EmailType from Email.
func (m Email) GetEmailType(f *field.EmailType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTime is a non-required field for Email.
func (m Email) OrigTime() (*field.OrigTime, errors.MessageRejectError) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from Email.
func (m Email) GetOrigTime(f *field.OrigTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RelatdSym is a non-required field for Email.
func (m Email) RelatdSym() (*field.RelatdSym, errors.MessageRejectError) {
	f := new(field.RelatdSym)
	err := m.Body.Get(f)
	return f, err
}

//GetRelatdSym reads a RelatdSym from Email.
func (m Email) GetRelatdSym(f *field.RelatdSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for Email.
func (m Email) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from Email.
func (m Email) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for Email.
func (m Email) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from Email.
func (m Email) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for Email.
func (m Email) LinesOfText() (*field.LinesOfText, errors.MessageRejectError) {
	f := new(field.LinesOfText)
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from Email.
func (m Email) GetLinesOfText(f *field.LinesOfText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a required field for Email.
func (m Email) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Email.
func (m Email) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Email.
func (m Email) RawDataLength() (*field.RawDataLength, errors.MessageRejectError) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Email.
func (m Email) GetRawDataLength(f *field.RawDataLength) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Email.
func (m Email) RawData() (*field.RawData, errors.MessageRejectError) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Email.
func (m Email) GetRawData(f *field.RawData) errors.MessageRejectError {
	return m.Body.Get(f)
}
