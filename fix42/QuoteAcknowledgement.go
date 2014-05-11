package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteAcknowledgement msg type = b.
type QuoteAcknowledgement struct {
	message.Message
}

//QuoteAcknowledgementBuilder builds QuoteAcknowledgement messages.
type QuoteAcknowledgementBuilder struct {
	message.MessageBuilder
}

//CreateQuoteAcknowledgementBuilder returns an initialized QuoteAcknowledgementBuilder with specified required fields.
func CreateQuoteAcknowledgementBuilder(
	quoteackstatus *field.QuoteAckStatusField) QuoteAcknowledgementBuilder {
	var builder QuoteAcknowledgementBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("b"))
	builder.Body().Set(quoteackstatus)
	return builder
}

//QuoteReqID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteAckStatus is a required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteAckStatus() (*field.QuoteAckStatusField, errors.MessageRejectError) {
	f := &field.QuoteAckStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteAckStatus reads a QuoteAckStatus from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteAckStatus(f *field.QuoteAckStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRejectReason is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteRejectReason() (*field.QuoteRejectReasonField, errors.MessageRejectError) {
	f := &field.QuoteRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRejectReason reads a QuoteRejectReason from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteRejectReason(f *field.QuoteRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) NoQuoteSets() (*field.NoQuoteSetsField, errors.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetNoQuoteSets(f *field.NoQuoteSetsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
