package fix41

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
	emailthreadid *field.EmailThreadIDField,
	emailtype *field.EmailTypeField,
	subject *field.SubjectField,
	linesoftext *field.LinesOfTextField) EmailBuilder {
	var builder EmailBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header.Set(field.NewMsgType("C"))
	builder.Body.Set(emailthreadid)
	builder.Body.Set(emailtype)
	builder.Body.Set(subject)
	builder.Body.Set(linesoftext)
	return builder
}

//EmailThreadID is a required field for Email.
func (m Email) EmailThreadID() (*field.EmailThreadIDField, errors.MessageRejectError) {
	f := &field.EmailThreadIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEmailThreadID reads a EmailThreadID from Email.
func (m Email) GetEmailThreadID(f *field.EmailThreadIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EmailType is a required field for Email.
func (m Email) EmailType() (*field.EmailTypeField, errors.MessageRejectError) {
	f := &field.EmailTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEmailType reads a EmailType from Email.
func (m Email) GetEmailType(f *field.EmailTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTime is a non-required field for Email.
func (m Email) OrigTime() (*field.OrigTimeField, errors.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from Email.
func (m Email) GetOrigTime(f *field.OrigTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Subject is a required field for Email.
func (m Email) Subject() (*field.SubjectField, errors.MessageRejectError) {
	f := &field.SubjectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubject reads a Subject from Email.
func (m Email) GetSubject(f *field.SubjectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for Email.
func (m Email) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from Email.
func (m Email) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for Email.
func (m Email) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from Email.
func (m Email) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for Email.
func (m Email) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from Email.
func (m Email) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for Email.
func (m Email) LinesOfText() (*field.LinesOfTextField, errors.MessageRejectError) {
	f := &field.LinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from Email.
func (m Email) GetLinesOfText(f *field.LinesOfTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Email.
func (m Email) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Email.
func (m Email) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Email.
func (m Email) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Email.
func (m Email) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}
