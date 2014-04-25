package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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
	quoteackstatus field.QuoteAckStatus) QuoteAcknowledgementBuilder {
	var builder QuoteAcknowledgementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("b"))
	builder.Body.Set(quoteackstatus)
	return builder
}

//QuoteReqID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteAckStatus is a required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteAckStatus() (*field.QuoteAckStatus, errors.MessageRejectError) {
	f := new(field.QuoteAckStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteAckStatus reads a QuoteAckStatus from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteAckStatus(f *field.QuoteAckStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRejectReason is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteRejectReason() (*field.QuoteRejectReason, errors.MessageRejectError) {
	f := new(field.QuoteRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRejectReason reads a QuoteRejectReason from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteRejectReason(f *field.QuoteRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteResponseLevel() (*field.QuoteResponseLevel, errors.MessageRejectError) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetQuoteResponseLevel(f *field.QuoteResponseLevel) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) NoQuoteSets() (*field.NoQuoteSets, errors.MessageRejectError) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from QuoteAcknowledgement.
func (m QuoteAcknowledgement) GetNoQuoteSets(f *field.NoQuoteSets) errors.MessageRejectError {
	return m.Body.Get(f)
}
