//Package massquote msg type = i.
package massquote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a MassQuote wrapper for the generic Message type
type Message struct {
	message.Message
}

//QuoteReqID is a non-required field for MassQuote.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuote.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for MassQuote.
func (m Message) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuote.
func (m Message) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for MassQuote.
func (m Message) QuoteType() (*field.QuoteTypeField, errors.MessageRejectError) {
	f := &field.QuoteTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from MassQuote.
func (m Message) GetQuoteType(f *field.QuoteTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuote.
func (m Message) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuote.
func (m Message) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MassQuote.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MassQuote.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MassQuote.
func (m Message) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MassQuote.
func (m Message) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for MassQuote.
func (m Message) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from MassQuote.
func (m Message) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MassQuote.
func (m Message) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MassQuote.
func (m Message) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DefBidSize is a non-required field for MassQuote.
func (m Message) DefBidSize() (*field.DefBidSizeField, errors.MessageRejectError) {
	f := &field.DefBidSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefBidSize reads a DefBidSize from MassQuote.
func (m Message) GetDefBidSize(f *field.DefBidSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DefOfferSize is a non-required field for MassQuote.
func (m Message) DefOfferSize() (*field.DefOfferSizeField, errors.MessageRejectError) {
	f := &field.DefOfferSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefOfferSize reads a DefOfferSize from MassQuote.
func (m Message) GetDefOfferSize(f *field.DefOfferSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a required field for MassQuote.
func (m Message) NoQuoteSets() (*field.NoQuoteSetsField, errors.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuote.
func (m Message) GetNoQuoteSets(f *field.NoQuoteSetsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MassQuote messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MassQuote.
func Builder(
	quoteid *field.QuoteIDField,
	noquotesets *field.NoQuoteSetsField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("i"))
	builder.Body().Set(quoteid)
	builder.Body().Set(noquotesets)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "i", r
}
