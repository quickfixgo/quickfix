package fix50sp1

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
	nolinesoftext field.NoLinesOfText) *EmailBuilder {
	builder := new(EmailBuilder)
	builder.Body.Set(emailthreadid)
	builder.Body.Set(emailtype)
	builder.Body.Set(subject)
	builder.Body.Set(nolinesoftext)
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

//EncodedSubjectLen is a non-required field for Email.
func (m *Email) EncodedSubjectLen() (*field.EncodedSubjectLen, error) {
	f := new(field.EncodedSubjectLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSubject is a non-required field for Email.
func (m *Email) EncodedSubject() (*field.EncodedSubject, error) {
	f := new(field.EncodedSubject)
	err := m.Body.Get(f)
	return f, err
}

//NoRoutingIDs is a non-required field for Email.
func (m *Email) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for Email.
func (m *Email) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for Email.
func (m *Email) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for Email.
func (m *Email) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
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

//NoLinesOfText is a required field for Email.
func (m *Email) NoLinesOfText() (*field.NoLinesOfText, error) {
	f := new(field.NoLinesOfText)
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
