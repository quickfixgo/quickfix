package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type NewOrderSingle struct {
	quickfixgo.Message
}

func (m *NewOrderSingle) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) IOIid() (*field.IOIid, error) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Rule80A() (*field.Rule80A, error) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
