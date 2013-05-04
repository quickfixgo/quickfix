package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SecurityDefinitionRequest struct {
	quickfixgo.Message
}

func (m *SecurityDefinitionRequest) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) SecurityRequestType() (*field.SecurityRequestType, error) {
	f := new(field.SecurityRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionRequest) ExpirationCycle() (*field.ExpirationCycle, error) {
	f := new(field.ExpirationCycle)
	err := m.Body.Get(f)
	return f, err
}
