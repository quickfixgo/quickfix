package fix41

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderSingle msg type = D.
type NewOrderSingle struct {
	message.Message
}

//NewOrderSingleBuilder builds NewOrderSingle messages.
type NewOrderSingleBuilder struct {
	message.MessageBuilder
}

//NewNewOrderSingleBuilder returns an initialized NewOrderSingleBuilder with specified required fields.
func NewNewOrderSingleBuilder(
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	symbol field.Symbol,
	side field.Side,
	ordtype field.OrdType) *NewOrderSingleBuilder {
	builder := new(NewOrderSingleBuilder)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderSingle.
func (m *NewOrderSingle) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClientID is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//ExecBroker is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//SettlmntTyp is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//HandlInst is a required field for NewOrderSingle.
func (m *NewOrderSingle) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//ExecInst is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//MinQty is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//MaxFloor is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//ExDestination is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//ProcessCode is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for NewOrderSingle.
func (m *NewOrderSingle) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDay is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//PrevClosePx is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for NewOrderSingle.
func (m *NewOrderSingle) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//LocateReqd is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a required field for NewOrderSingle.
func (m *NewOrderSingle) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//StopPx is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//IOIid is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) IOIid() (*field.IOIid, error) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//TimeInForce is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//Rule80A is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) Rule80A() (*field.Rule80A, error) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}

//ForexReq is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate2 is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) FutSettDate2() (*field.FutSettDate2, error) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty2 is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//OpenClose is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) OpenClose() (*field.OpenClose, error) {
	f := new(field.OpenClose)
	err := m.Body.Get(f)
	return f, err
}

//CoveredOrUncovered is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) CoveredOrUncovered() (*field.CoveredOrUncovered, error) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//CustomerOrFirm is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) CustomerOrFirm() (*field.CustomerOrFirm, error) {
	f := new(field.CustomerOrFirm)
	err := m.Body.Get(f)
	return f, err
}

//MaxShow is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//PegDifference is a non-required field for NewOrderSingle.
func (m *NewOrderSingle) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}
