package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteCancel msg type = Z.
type QuoteCancel struct {
	message.Message
}

//QuoteCancelBuilder builds QuoteCancel messages.
type QuoteCancelBuilder struct {
	message.MessageBuilder
}

//NewQuoteCancelBuilder returns an initialized QuoteCancelBuilder with specified required fields.
func NewQuoteCancelBuilder(
	quotecanceltype field.QuoteCancelType) *QuoteCancelBuilder {
	builder := new(QuoteCancelBuilder)
	builder.Body.Set(quotecanceltype)
	return builder
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m *QuoteCancel) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a non-required field for QuoteCancel.
func (m *QuoteCancel) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteCancelType is a required field for QuoteCancel.
func (m *QuoteCancel) QuoteCancelType() (*field.QuoteCancelType, error) {
	f := new(field.QuoteCancelType)
	err := m.Body.Get(f)
	return f, err
}

//QuoteResponseLevel is a non-required field for QuoteCancel.
func (m *QuoteCancel) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteCancel.
func (m *QuoteCancel) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for QuoteCancel.
func (m *QuoteCancel) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for QuoteCancel.
func (m *QuoteCancel) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for QuoteCancel.
func (m *QuoteCancel) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteCancel.
func (m *QuoteCancel) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteCancel.
func (m *QuoteCancel) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//NoQuoteEntries is a non-required field for QuoteCancel.
func (m *QuoteCancel) NoQuoteEntries() (*field.NoQuoteEntries, error) {
	f := new(field.NoQuoteEntries)
	err := m.Body.Get(f)
	return f, err
}

//QuoteMsgID is a non-required field for QuoteCancel.
func (m *QuoteCancel) QuoteMsgID() (*field.QuoteMsgID, error) {
	f := new(field.QuoteMsgID)
	err := m.Body.Get(f)
	return f, err
}
