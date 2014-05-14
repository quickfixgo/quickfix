//Package massquoteacknowledgement msg type = b.
package massquoteacknowledgement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a MassQuoteAcknowledgement wrapper for the generic Message type
type Message struct {
	message.Message
}

//QuoteReqID is a non-required field for MassQuoteAcknowledgement.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuoteAcknowledgement.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for MassQuoteAcknowledgement.
func (m Message) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuoteAcknowledgement.
func (m Message) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteStatus is a required field for MassQuoteAcknowledgement.
func (m Message) QuoteStatus() (*field.QuoteStatusField, errors.MessageRejectError) {
	f := &field.QuoteStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteStatus reads a QuoteStatus from MassQuoteAcknowledgement.
func (m Message) GetQuoteStatus(f *field.QuoteStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRejectReason is a non-required field for MassQuoteAcknowledgement.
func (m Message) QuoteRejectReason() (*field.QuoteRejectReasonField, errors.MessageRejectError) {
	f := &field.QuoteRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRejectReason reads a QuoteRejectReason from MassQuoteAcknowledgement.
func (m Message) GetQuoteRejectReason(f *field.QuoteRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuoteAcknowledgement.
func (m Message) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuoteAcknowledgement.
func (m Message) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for MassQuoteAcknowledgement.
func (m Message) QuoteType() (*field.QuoteTypeField, errors.MessageRejectError) {
	f := &field.QuoteTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from MassQuoteAcknowledgement.
func (m Message) GetQuoteType(f *field.QuoteTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MassQuoteAcknowledgement.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MassQuoteAcknowledgement.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MassQuoteAcknowledgement.
func (m Message) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MassQuoteAcknowledgement.
func (m Message) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for MassQuoteAcknowledgement.
func (m Message) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from MassQuoteAcknowledgement.
func (m Message) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MassQuoteAcknowledgement.
func (m Message) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MassQuoteAcknowledgement.
func (m Message) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MassQuoteAcknowledgement.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MassQuoteAcknowledgement.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MassQuoteAcknowledgement.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MassQuoteAcknowledgement.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MassQuoteAcknowledgement.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MassQuoteAcknowledgement.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a non-required field for MassQuoteAcknowledgement.
func (m Message) NoQuoteSets() (*field.NoQuoteSetsField, errors.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuoteAcknowledgement.
func (m Message) GetNoQuoteSets(f *field.NoQuoteSetsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteCancelType is a non-required field for MassQuoteAcknowledgement.
func (m Message) QuoteCancelType() (*field.QuoteCancelTypeField, errors.MessageRejectError) {
	f := &field.QuoteCancelTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteCancelType reads a QuoteCancelType from MassQuoteAcknowledgement.
func (m Message) GetQuoteCancelType(f *field.QuoteCancelTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MassQuoteAcknowledgement messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MassQuoteAcknowledgement.
func Builder(
	quotestatus *field.QuoteStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("b"))
	builder.Body().Set(quotestatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "b", r
}
