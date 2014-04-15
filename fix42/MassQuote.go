package fix42

import (
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
	quoteid field.QuoteID,
	noquotesets field.NoQuoteSets) MassQuoteBuilder {
	var builder MassQuoteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quoteid)
	builder.Body.Set(noquotesets)
	return builder
}

//QuoteReqID is a non-required field for MassQuote.
func (m MassQuote) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a required field for MassQuote.
func (m MassQuote) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for MassQuote.
func (m MassQuote) QuoteResponseLevel() (field.QuoteResponseLevel, error) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//DefBidSize is a non-required field for MassQuote.
func (m MassQuote) DefBidSize() (field.DefBidSize, error) {
	var f field.DefBidSize
	err := m.Body.Get(&f)
	return f, err
}

//DefOfferSize is a non-required field for MassQuote.
func (m MassQuote) DefOfferSize() (field.DefOfferSize, error) {
	var f field.DefOfferSize
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteSets is a required field for MassQuote.
func (m MassQuote) NoQuoteSets() (field.NoQuoteSets, error) {
	var f field.NoQuoteSets
	err := m.Body.Get(&f)
	return f, err
}
