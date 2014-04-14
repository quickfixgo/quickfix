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

//NewQuoteAcknowledgementBuilder returns an initialized QuoteAcknowledgementBuilder with specified required fields.
func NewQuoteAcknowledgementBuilder(
	quoteackstatus field.QuoteAckStatus) *QuoteAcknowledgementBuilder {
	builder := new(QuoteAcknowledgementBuilder)
	builder.Body.Set(quoteackstatus)
	return builder
}

//QuoteReqID is a non-required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a non-required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteAckStatus is a required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) QuoteAckStatus() (*field.QuoteAckStatus, error) {
	f := new(field.QuoteAckStatus)
	err := m.Body.Get(f)
	return f, err
}

//QuoteRejectReason is a non-required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) QuoteRejectReason() (*field.QuoteRejectReason, error) {
	f := new(field.QuoteRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//QuoteResponseLevel is a non-required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//NoQuoteSets is a non-required field for QuoteAcknowledgement.
func (m *QuoteAcknowledgement) NoQuoteSets() (*field.NoQuoteSets, error) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}
