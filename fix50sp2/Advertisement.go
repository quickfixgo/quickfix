package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type Advertisement struct {
	quickfixgo.Message
}

func (m *Advertisement) AdvId() (*field.AdvId, error) {
	f := new(field.AdvId)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) AdvTransType() (*field.AdvTransType, error) {
	f := new(field.AdvTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) AdvRefID() (*field.AdvRefID, error) {
	f := new(field.AdvRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) AdvSide() (*field.AdvSide, error) {
	f := new(field.AdvSide)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
