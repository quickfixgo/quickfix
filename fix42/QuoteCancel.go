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
	builder.Header.Set(field.BuildMsgType("Z"))
	builder.Body.Set(quoteid)
	builder.Body.Set(quotecanceltype)
	builder.Body.Set(noquoteentries)
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

//QuoteID is a required field for QuoteCancel.
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

//NoQuoteEntries is a required field for QuoteCancel.
func (m QuoteCancel) NoQuoteEntries() (*field.NoQuoteEntries, errors.MessageRejectError) {
	f := new(field.NoQuoteEntries)
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteEntries reads a NoQuoteEntries from QuoteCancel.
func (m QuoteCancel) GetNoQuoteEntries(f *field.NoQuoteEntries) errors.MessageRejectError {
	return m.Body.Get(f)
}
