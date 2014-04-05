package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type Allocation struct {
	quickfixgo.Message
}

func (m *Allocation) AllocTransType() (*field.AllocTransType, error) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) RefAllocID() (*field.RefAllocID, error) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) AvgPrxPrecision() (*field.AvgPrxPrecision, error) {
	f := new(field.AvgPrxPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) SettlCurrAmt() (*field.SettlCurrAmt, error) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) NoMiscFees() (*field.NoMiscFees, error) {
	f := new(field.NoMiscFees)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Shares() (*field.Shares, error) {
	f := new(field.Shares)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) OpenClose() (*field.OpenClose, error) {
	f := new(field.OpenClose)
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
func (m *Allocation) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}
func (m *Allocation) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
