package fix42

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type QuoteCancel struct {
	quickfixgo.Message
}

func (m *QuoteCancel) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) QuoteCancelType() (*field.QuoteCancelType, error) {
	f := new(field.QuoteCancelType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) NoQuoteEntries() (*field.NoQuoteEntries, error) {
	f := new(field.NoQuoteEntries)
	err := m.Body.Get(f)
	return f, err
}
