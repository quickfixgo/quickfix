package fix41

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
	emailthreadid field.EmailThreadID,
	emailtype field.EmailType,
	subject field.Subject,
	linesoftext field.LinesOfText) EmailBuilder {
	var builder EmailBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("C"))
	builder.Body.Set(emailthreadid)
	builder.Body.Set(emailtype)
	builder.Body.Set(subject)
	builder.Body.Set(linesoftext)
	return builder
}

//EmailThreadID is a required field for Email.
func (m Email) EmailThreadID() (*field.EmailThreadID, errors.MessageRejectError) {
	f := new(field.EmailThreadID)
	err := m.Body.Get(f)
	return f, err
}

//GetEmailThreadID reads a EmailThreadID from Email.
func (m Email) GetEmailThreadID(f *field.EmailThreadID) errors.MessageRejectError {
	return m.Body.Get(f)
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

//Subject is a required field for Email.
func (m Email) Subject() (*field.Subject, errors.MessageRejectError) {
	f := new(field.Subject)
	err := m.Body.Get(f)
	return f, err
}

//GetSubject reads a Subject from Email.
func (m Email) GetSubject(f *field.Subject) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for Email.
func (m Email) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from Email.
func (m Email) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
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
