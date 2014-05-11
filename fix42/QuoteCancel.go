package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	quoteid *field.QuoteIDField,
	quotecanceltype *field.QuoteCancelTypeField,
	noquoteentries *field.NoQuoteEntriesField) QuoteCancelBuilder {
	var builder QuoteCancelBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("Z"))
	builder.Body().Set(quoteid)
	builder.Body().Set(quotecanceltype)
	builder.Body().Set(noquoteentries)
	return builder
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteCancel.
func (m QuoteCancel) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for QuoteCancel.
func (m QuoteCancel) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteCancel.
func (m QuoteCancel) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteCancelType is a required field for QuoteCancel.
func (m QuoteCancel) QuoteCancelType() (*field.QuoteCancelTypeField, errors.MessageRejectError) {
	f := &field.QuoteCancelTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteCancelType reads a QuoteCancelType from QuoteCancel.
func (m QuoteCancel) GetQuoteCancelType(f *field.QuoteCancelTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for QuoteCancel.
func (m QuoteCancel) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from QuoteCancel.
func (m QuoteCancel) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteCancel.
func (m QuoteCancel) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteCancel.
func (m QuoteCancel) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteEntries is a required field for QuoteCancel.
func (m QuoteCancel) NoQuoteEntries() (*field.NoQuoteEntriesField, errors.MessageRejectError) {
	f := &field.NoQuoteEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteEntries reads a NoQuoteEntries from QuoteCancel.
func (m QuoteCancel) GetNoQuoteEntries(f *field.NoQuoteEntriesField) errors.MessageRejectError {
	return m.Body.Get(f)
}
