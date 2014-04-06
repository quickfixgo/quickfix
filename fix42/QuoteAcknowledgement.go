package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type QuoteAcknowledgement struct {
	quickfix.Message
}

func (m *QuoteAcknowledgement) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteAcknowledgement) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteAcknowledgement) QuoteAckStatus() (*field.QuoteAckStatus, error) {
	f := new(field.QuoteAckStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteAcknowledgement) QuoteRejectReason() (*field.QuoteRejectReason, error) {
	f := new(field.QuoteRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteAcknowledgement) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteAcknowledgement) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteAcknowledgement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteAcknowledgement) NoQuoteSets() (*field.NoQuoteSets, error) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}
