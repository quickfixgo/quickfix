package fix41

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

//NewEmailBuilder returns an initialized EmailBuilder with specified required fields.
func NewEmailBuilder(
	emailthreadid field.EmailThreadID,
	emailtype field.EmailType,
	subject field.Subject,
	linesoftext field.LinesOfText) *EmailBuilder {
	builder := new(EmailBuilder)
	builder.Body.Set(emailthreadid)
	builder.Body.Set(emailtype)
	builder.Body.Set(subject)
	builder.Body.Set(linesoftext)
	return builder
}

//EmailThreadID is a required field for Email.
func (m *Email) EmailThreadID() (*field.EmailThreadID, error) {
	f := new(field.EmailThreadID)
	err := m.Body.Get(f)
	return f, err
}

//EmailType is a required field for Email.
func (m *Email) EmailType() (*field.EmailType, error) {
	f := new(field.EmailType)
	err := m.Body.Get(f)
	return f, err
}

//OrigTime is a non-required field for Email.
func (m *Email) OrigTime() (*field.OrigTime, error) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}

//Subject is a required field for Email.
func (m *Email) Subject() (*field.Subject, error) {
	f := new(field.Subject)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for Email.
func (m *Email) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//OrderID is a non-required field for Email.
func (m *Email) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for Email.
func (m *Email) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//LinesOfText is a required field for Email.
func (m *Email) LinesOfText() (*field.LinesOfText, error) {
	f := new(field.LinesOfText)
	err := m.Body.Get(f)
	return f, err
}

//RawDataLength is a non-required field for Email.
func (m *Email) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//RawData is a non-required field for Email.
func (m *Email) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
