//Package massquote msg type = i.
package massquote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a MassQuote wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//QuoteReqID is a non-required field for MassQuote.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, quickfix.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuote.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for MassQuote.
func (m Message) QuoteID() (*field.QuoteIDField, quickfix.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuote.
func (m Message) GetQuoteID(f *field.QuoteIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for MassQuote.
func (m Message) QuoteType() (*field.QuoteTypeField, quickfix.MessageRejectError) {
	f := &field.QuoteTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from MassQuote.
func (m Message) GetQuoteType(f *field.QuoteTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuote.
func (m Message) QuoteResponseLevel() (*field.QuoteResponseLevelField, quickfix.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuote.
func (m Message) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MassQuote.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MassQuote.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MassQuote.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MassQuote.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MassQuote.
func (m Message) AccountType() (*field.AccountTypeField, quickfix.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MassQuote.
func (m Message) GetAccountType(f *field.AccountTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DefBidSize is a non-required field for MassQuote.
func (m Message) DefBidSize() (*field.DefBidSizeField, quickfix.MessageRejectError) {
	f := &field.DefBidSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefBidSize reads a DefBidSize from MassQuote.
func (m Message) GetDefBidSize(f *field.DefBidSizeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DefOfferSize is a non-required field for MassQuote.
func (m Message) DefOfferSize() (*field.DefOfferSizeField, quickfix.MessageRejectError) {
	f := &field.DefOfferSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefOfferSize reads a DefOfferSize from MassQuote.
func (m Message) GetDefOfferSize(f *field.DefOfferSizeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a required field for MassQuote.
func (m Message) NoQuoteSets() (*field.NoQuoteSetsField, quickfix.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuote.
func (m Message) GetNoQuoteSets(f *field.NoQuoteSetsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for MassQuote.
func New(
	quoteid *field.QuoteIDField,
	noquotesets *field.NoQuoteSetsField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX43))
	builder.Header.Set(field.NewMsgType("i"))
	builder.Body.Set(quoteid)
	builder.Body.Set(noquotesets)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX43, "i", r
}
