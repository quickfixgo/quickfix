package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SecurityTypes struct {
	quickfixgo.Message
}

func (m *SecurityTypes) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) SecurityResponseType() (*field.SecurityResponseType, error) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) TotNoSecurityTypes() (*field.TotNoSecurityTypes, error) {
	f := new(field.TotNoSecurityTypes)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
