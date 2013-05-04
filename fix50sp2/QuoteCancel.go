package fix50sp2

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
func (m *QuoteCancel) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) QuoteMsgID() (*field.QuoteMsgID, error) {
	f := new(field.QuoteMsgID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteCancel) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}
