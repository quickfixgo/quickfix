package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type SecurityTypes struct {
	message.Message
}

func (m *SecurityTypes) SecurityResponseType() (*field.SecurityResponseType, error) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityTypes) TotalNumSecurityTypes() (*field.TotalNumSecurityTypes, error) {
	f := new(field.TotalNumSecurityTypes)
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
func (m *SecurityTypes) NoSecurityTypes() (*field.NoSecurityTypes, error) {
	f := new(field.NoSecurityTypes)
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
