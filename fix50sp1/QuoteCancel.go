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

//CreateQuoteCancelBuilder returns an initialized QuoteCancelBuilder with specified required fields.
func CreateQuoteCancelBuilder(
	quotecanceltype field.QuoteCancelType) QuoteCancelBuilder {
	var builder QuoteCancelBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quotecanceltype)
	return builder
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteCancelType is a required field for QuoteCancel.
func (m QuoteCancel) QuoteCancelType() (field.QuoteCancelType, error) {
	var f field.QuoteCancelType
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteResponseLevel() (field.QuoteResponseLevel, error) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteCancel.
func (m QuoteCancel) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for QuoteCancel.
func (m QuoteCancel) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for QuoteCancel.
func (m QuoteCancel) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for QuoteCancel.
func (m QuoteCancel) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteCancel.
func (m QuoteCancel) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteCancel.
func (m QuoteCancel) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteEntries is a non-required field for QuoteCancel.
func (m QuoteCancel) NoQuoteEntries() (field.NoQuoteEntries, error) {
	var f field.NoQuoteEntries
	err := m.Body.Get(&f)
	return f, err
}

//QuoteMsgID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteMsgID() (field.QuoteMsgID, error) {
	var f field.QuoteMsgID
	err := m.Body.Get(&f)
	return f, err
}
