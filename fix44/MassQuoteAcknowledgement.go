package fix44

import (
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
func (m MassQuoteAcknowledgement) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteStatus is a required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteStatus() (field.QuoteStatus, error) {
	var f field.QuoteStatus
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRejectReason is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteRejectReason() (field.QuoteRejectReason, error) {
	var f field.QuoteRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteResponseLevel() (field.QuoteResponseLevel, error) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//QuoteType is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) QuoteType() (field.QuoteType, error) {
	var f field.QuoteType
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteSets is a non-required field for MassQuoteAcknowledgement.
func (m MassQuoteAcknowledgement) NoQuoteSets() (field.NoQuoteSets, error) {
	var f field.NoQuoteSets
	err := m.Body.Get(&f)
	return f, err
}
