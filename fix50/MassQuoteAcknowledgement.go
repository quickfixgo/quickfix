package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MassQuoteAcknowledgement msg type = b.
type MassQuoteAcknowledgement struct {
	message.Message
}

//MassQuoteAcknowledgementBuilder builds MassQuoteAcknowledgement messages.
type MassQuoteAcknowledgementBuilder struct {
	message.MessageBuilder
}

//CreateMassQuoteAcknowledgementBuilder returns an initialized MassQuoteAcknowledgementBuilder with specified required fields.
func CreateMassQuoteAcknowledgementBuilder(
	quotestatus field.QuoteStatus) MassQuoteAcknowledgementBuilder {
	var builder MassQuoteAcknowledgementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quotestatus)
	return builder
}

//QuoteReqID is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteStatus is a required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteStatus() (*field.QuoteStatus, errors.MessageRejectError) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteStatus reads a QuoteStatus from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteStatus(f *field.QuoteStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRejectReason is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteRejectReason() (*field.QuoteRejectReason, errors.MessageRejectError) {
	f := new(field.QuoteRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRejectReason reads a QuoteRejectReason from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteRejectReason(f *field.QuoteRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteResponseLevel() (*field.QuoteResponseLevel, errors.MessageRejectError) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteResponseLevel(f *field.QuoteResponseLevel) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteType() (*field.QuoteType, errors.MessageRejectError) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteType(f *field.QuoteType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) NoQuoteSets() (*field.NoQuoteSets, errors.MessageRejectError) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetNoQuoteSets(f *field.NoQuoteSets) errors.MessageRejectError {
	return m.Body.Get(f)
}
