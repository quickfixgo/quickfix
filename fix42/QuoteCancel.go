package fix42

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
	quoteid field.QuoteID,
	quotecanceltype field.QuoteCancelType,
	noquoteentries field.NoQuoteEntries) QuoteCancelBuilder {
	var builder QuoteCancelBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quoteid)
	builder.Body.Set(quotecanceltype)
	builder.Body.Set(noquoteentries)
	return builder
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteReqID() (field.QuoteReqID, errors.MessageRejectError) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a required field for QuoteCancel.
func (m QuoteCancel) QuoteID() (field.QuoteID, errors.MessageRejectError) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteCancelType is a required field for QuoteCancel.
func (m QuoteCancel) QuoteCancelType() (field.QuoteCancelType, errors.MessageRejectError) {
	var f field.QuoteCancelType
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteResponseLevel() (field.QuoteResponseLevel, errors.MessageRejectError) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteCancel.
func (m QuoteCancel) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteEntries is a required field for QuoteCancel.
func (m QuoteCancel) NoQuoteEntries() (field.NoQuoteEntries, errors.MessageRejectError) {
	var f field.NoQuoteEntries
	err := m.Body.Get(&f)
	return f, err
}
