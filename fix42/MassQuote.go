package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type MassQuote struct {
	quickfix.Message
}

func (m *MassQuote) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) DefBidSize() (*field.DefBidSize, error) {
	f := new(field.DefBidSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) DefOfferSize() (*field.DefOfferSize, error) {
	f := new(field.DefOfferSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) NoQuoteSets() (*field.NoQuoteSets, error) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}
