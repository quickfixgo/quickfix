package fix50

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type MassQuoteAcknowledgement struct {
	quickfix.Message
}

func (m *MassQuoteAcknowledgement) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) NoQuoteSets() (*field.NoQuoteSets, error) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) QuoteResponseLevel() (*field.QuoteResponseLevel, error) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) QuoteStatus() (*field.QuoteStatus, error) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) QuoteRejectReason() (*field.QuoteRejectReason, error) {
	f := new(field.QuoteRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *MassQuoteAcknowledgement) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
