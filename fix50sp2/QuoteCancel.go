package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Header.Set(field.BuildMsgType("Z"))
	builder.Body.Set(quotecanceltype)
	return builder
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteCancel.
func (m QuoteCancel) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteCancel.
func (m QuoteCancel) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteCancelType is a required field for QuoteCancel.
func (m QuoteCancel) QuoteCancelType() (*field.QuoteCancelType, errors.MessageRejectError) {
	f := new(field.QuoteCancelType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteCancelType reads a QuoteCancelType from QuoteCancel.
func (m QuoteCancel) GetQuoteCancelType(f *field.QuoteCancelType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteResponseLevel() (*field.QuoteResponseLevel, errors.MessageRejectError) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from QuoteCancel.
func (m QuoteCancel) GetQuoteResponseLevel(f *field.QuoteResponseLevel) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for QuoteCancel.
func (m QuoteCancel) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from QuoteCancel.
func (m QuoteCancel) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for QuoteCancel.
func (m QuoteCancel) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from QuoteCancel.
func (m QuoteCancel) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for QuoteCancel.
func (m QuoteCancel) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from QuoteCancel.
func (m QuoteCancel) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for QuoteCancel.
func (m QuoteCancel) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from QuoteCancel.
func (m QuoteCancel) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteCancel.
func (m QuoteCancel) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteCancel.
func (m QuoteCancel) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for QuoteCancel.
func (m QuoteCancel) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from QuoteCancel.
func (m QuoteCancel) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteEntries is a non-required field for QuoteCancel.
func (m QuoteCancel) NoQuoteEntries() (*field.NoQuoteEntries, errors.MessageRejectError) {
	f := new(field.NoQuoteEntries)
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteEntries reads a NoQuoteEntries from QuoteCancel.
func (m QuoteCancel) GetNoQuoteEntries(f *field.NoQuoteEntries) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteMsgID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteMsgID() (*field.QuoteMsgID, errors.MessageRejectError) {
	f := new(field.QuoteMsgID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteMsgID reads a QuoteMsgID from QuoteCancel.
func (m QuoteCancel) GetQuoteMsgID(f *field.QuoteMsgID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteType() (*field.QuoteType, errors.MessageRejectError) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from QuoteCancel.
func (m QuoteCancel) GetQuoteType(f *field.QuoteType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTargetPartyIDs is a non-required field for QuoteCancel.
func (m QuoteCancel) NoTargetPartyIDs() (*field.NoTargetPartyIDs, errors.MessageRejectError) {
	f := new(field.NoTargetPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTargetPartyIDs reads a NoTargetPartyIDs from QuoteCancel.
func (m QuoteCancel) GetNoTargetPartyIDs(f *field.NoTargetPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
