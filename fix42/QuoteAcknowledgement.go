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
	builder.Body.Set(quoteackstatus)
	return builder
}

//QuoteReqID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteReqID() (field.QuoteReqID, errors.MessageRejectError) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteID() (field.QuoteID, errors.MessageRejectError) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteAckStatus is a required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteAckStatus() (field.QuoteAckStatus, errors.MessageRejectError) {
	var f field.QuoteAckStatus
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRejectReason is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteRejectReason() (field.QuoteRejectReason, errors.MessageRejectError) {
	var f field.QuoteRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteResponseLevel() (field.QuoteResponseLevel, errors.MessageRejectError) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteSets is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) NoQuoteSets() (field.NoQuoteSets, errors.MessageRejectError) {
	var f field.NoQuoteSets
	err := m.Body.Get(&f)
	return f, err
}
