package fix42

import (
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
func (m QuoteAcknowledgement) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteAckStatus is a required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteAckStatus() (field.QuoteAckStatus, error) {
	var f field.QuoteAckStatus
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRejectReason is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteRejectReason() (field.QuoteRejectReason, error) {
	var f field.QuoteRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) QuoteResponseLevel() (field.QuoteResponseLevel, error) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteSets is a non-required field for QuoteAcknowledgement.
func (m QuoteAcknowledgement) NoQuoteSets() (field.NoQuoteSets, error) {
	var f field.NoQuoteSets
	err := m.Body.Get(&f)
	return f, err
}
