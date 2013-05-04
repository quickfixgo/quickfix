package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type Allocation struct {
	quickfixgo.Message
}

func (m *Allocation) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AllocTransType() (*field.AllocTransType, error) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AllocType() (*field.AllocType, error) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) RefAllocID() (*field.RefAllocID, error) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AllocLinkID() (*field.AllocLinkID, error) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AllocLinkType() (*field.AllocLinkType, error) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) BookingRefID() (*field.BookingRefID, error) {
	f := new(field.BookingRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AvgPrxPrecision() (*field.AvgPrxPrecision, error) {
	f := new(field.AvgPrxPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmt, error) {
	f := new(field.TotalAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}
