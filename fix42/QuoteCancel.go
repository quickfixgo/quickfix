package fix42

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
	quoteid field.QuoteID,
	quotecanceltype field.QuoteCancelType,
	noquoteentries field.NoQuoteEntries) *QuoteCancelBuilder {
	builder := new(QuoteCancelBuilder)
	builder.Body.Set(quoteid)
	builder.Body.Set(quotecanceltype)
	builder.Body.Set(noquoteentries)
	return builder
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m *QuoteCancel) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a required field for QuoteCancel.
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

//TradingSessionID is a non-required field for QuoteCancel.
func (m *QuoteCancel) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//NoQuoteEntries is a required field for QuoteCancel.
func (m *QuoteCancel) NoQuoteEntries() (*field.NoQuoteEntries, error) {
	f := new(field.NoQuoteEntries)
	err := m.Body.Get(f)
	return f, err
}
