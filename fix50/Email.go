package fix50

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
	nolinesoftext field.NoLinesOfText) EmailBuilder {
	var builder EmailBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(emailthreadid)
	builder.Body.Set(emailtype)
	builder.Body.Set(subject)
	builder.Body.Set(nolinesoftext)
	return builder
}

//EmailThreadID is a required field for Email.
func (m Email) EmailThreadID() (field.EmailThreadID, errors.MessageRejectError) {
	var f field.EmailThreadID
	err := m.Body.Get(&f)
	return f, err
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

//Subject is a required field for Email.
func (m Email) Subject() (field.Subject, errors.MessageRejectError) {
	var f field.Subject
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSubjectLen is a non-required field for Email.
func (m Email) EncodedSubjectLen() (field.EncodedSubjectLen, errors.MessageRejectError) {
	var f field.EncodedSubjectLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSubject is a non-required field for Email.
func (m Email) EncodedSubject() (field.EncodedSubject, errors.MessageRejectError) {
	var f field.EncodedSubject
	err := m.Body.Get(&f)
	return f, err
}

//NoRoutingIDs is a non-required field for Email.
func (m Email) NoRoutingIDs() (field.NoRoutingIDs, errors.MessageRejectError) {
	var f field.NoRoutingIDs
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for Email.
func (m Email) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for Email.
func (m Email) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for Email.
func (m Email) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
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

//NoLinesOfText is a required field for Email.
func (m Email) NoLinesOfText() (field.NoLinesOfText, errors.MessageRejectError) {
	var f field.NoLinesOfText
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
