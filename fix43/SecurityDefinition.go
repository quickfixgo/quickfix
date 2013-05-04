package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SecurityDefinition struct {
	quickfixgo.Message
}

func (m *SecurityDefinition) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) SecurityResponseType() (*field.SecurityResponseType, error) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) RoundLot() (*field.RoundLot, error) {
	f := new(field.RoundLot)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinition) MinTradeVol() (*field.MinTradeVol, error) {
	f := new(field.MinTradeVol)
	err := m.Body.Get(f)
	return f, err
}
