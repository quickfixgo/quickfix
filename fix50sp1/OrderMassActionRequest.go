package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type OrderMassActionRequest struct {
	quickfixgo.Message
}

func (m *OrderMassActionRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MassActionType() (*field.MassActionType, error) {
	f := new(field.MassActionType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MassActionScope() (*field.MassActionScope, error) {
	f := new(field.MassActionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
