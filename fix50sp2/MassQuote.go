package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MassQuote struct {
	quickfixgo.Message
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
func (m *MassQuote) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuote) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
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
