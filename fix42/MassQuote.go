package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MassQuote msg type = i.
type MassQuote struct {
	message.Message
}

//MassQuoteBuilder builds MassQuote messages.
type MassQuoteBuilder struct {
	message.MessageBuilder
}

//CreateMassQuoteBuilder returns an initialized MassQuoteBuilder with specified required fields.
func CreateMassQuoteBuilder(
	quoteid *field.QuoteIDField,
	noquotesets *field.NoQuoteSetsField) MassQuoteBuilder {
	var builder MassQuoteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header.Set(field.NewMsgType("i"))
	builder.Body.Set(quoteid)
	builder.Body.Set(noquotesets)
	return builder
}

//QuoteReqID is a non-required field for MassQuote.
func (m MassQuote) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuote.
func (m MassQuote) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for MassQuote.
func (m MassQuote) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuote.
func (m MassQuote) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuote.
func (m MassQuote) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuote.
func (m MassQuote) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DefBidSize is a non-required field for MassQuote.
func (m MassQuote) DefBidSize() (*field.DefBidSizeField, errors.MessageRejectError) {
	f := &field.DefBidSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefBidSize reads a DefBidSize from MassQuote.
func (m MassQuote) GetDefBidSize(f *field.DefBidSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DefOfferSize is a non-required field for MassQuote.
func (m MassQuote) DefOfferSize() (*field.DefOfferSizeField, errors.MessageRejectError) {
	f := &field.DefOfferSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefOfferSize reads a DefOfferSize from MassQuote.
func (m MassQuote) GetDefOfferSize(f *field.DefOfferSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a required field for MassQuote.
func (m MassQuote) NoQuoteSets() (*field.NoQuoteSetsField, errors.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuote.
func (m MassQuote) GetNoQuoteSets(f *field.NoQuoteSetsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
