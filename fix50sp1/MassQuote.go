package fix50sp1

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

//NewMassQuoteBuilder returns an initialized MassQuoteBuilder with specified required fields.
func NewMassQuoteBuilder(
	quoteid field.QuoteID,
	noquotesets field.NoQuoteSets) *MassQuoteBuilder {
	builder := new(MassQuoteBuilder)
	builder.Body.Set(quoteid)
	builder.Body.Set(noquotesets)
	return builder
}

//QuoteReqID is a non-required field for MassQuote.
func (m *MassQuote) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a required field for MassQuote.
func (m *MassQuote) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteType is a non-required field for MassQuote.
func (m *MassQuote) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//QuoteResponseLevel is a non-required field for MassQuote.
func (m *MassQuote) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for MassQuote.
func (m *MassQuote) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for MassQuote.
func (m *MassQuote) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for MassQuote.
func (m *MassQuote) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for MassQuote.
func (m *MassQuote) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//DefBidSize is a non-required field for MassQuote.
func (m *MassQuote) DefBidSize() (*field.DefBidSize, error) {
	f := new(field.DefBidSize)
	err := m.Body.Get(f)
	return f, err
}

//DefOfferSize is a non-required field for MassQuote.
func (m *MassQuote) DefOfferSize() (*field.DefOfferSize, error) {
	f := new(field.DefOfferSize)
	err := m.Body.Get(f)
	return f, err
}

//NoQuoteSets is a required field for MassQuote.
func (m *MassQuote) NoQuoteSets() (*field.NoQuoteSets, error) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}
