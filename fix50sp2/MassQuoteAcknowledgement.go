package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	quotestatus *field.QuoteStatusField) MassQuoteAcknowledgementBuilder {
	var builder MassQuoteAcknowledgementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("b"))
	builder.Body.Set(quotestatus)
	return builder
}

//QuoteReqID is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteStatus is a required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteStatus() (*field.QuoteStatusField, errors.MessageRejectError) {
	f := &field.QuoteStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteStatus reads a QuoteStatus from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteStatus(f *field.QuoteStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRejectReason is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteRejectReason() (*field.QuoteRejectReasonField, errors.MessageRejectError) {
	f := &field.QuoteRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRejectReason reads a QuoteRejectReason from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteRejectReason(f *field.QuoteRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteType() (*field.QuoteTypeField, errors.MessageRejectError) {
	f := &field.QuoteTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteType(f *field.QuoteTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) NoQuoteSets() (*field.NoQuoteSetsField, errors.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetNoQuoteSets(f *field.NoQuoteSetsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteCancelType is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteCancelType() (*field.QuoteCancelTypeField, errors.MessageRejectError) {
	f := &field.QuoteCancelTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteCancelType reads a QuoteCancelType from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetQuoteCancelType(f *field.QuoteCancelTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTargetPartyIDs is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) NoTargetPartyIDs() (*field.NoTargetPartyIDsField, errors.MessageRejectError) {
	f := &field.NoTargetPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTargetPartyIDs reads a NoTargetPartyIDs from MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) GetNoTargetPartyIDs(f *field.NoTargetPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
