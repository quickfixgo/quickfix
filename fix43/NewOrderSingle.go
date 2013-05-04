package fix43

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
func (m *NewOrderSingle) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) DayBookingInst() (*field.DayBookingInst, error) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) BookingUnit() (*field.BookingUnit, error) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
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
func (m *NewOrderSingle) CashMargin() (*field.CashMargin, error) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
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
func (m *NewOrderSingle) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) QuantityType() (*field.QuantityType, error) {
	f := new(field.QuantityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
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
func (m *NewOrderSingle) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
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
func (m *NewOrderSingle) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
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
func (m *NewOrderSingle) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) FutSettDate2() (*field.FutSettDate2, error) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Price2() (*field.Price2, error) {
	f := new(field.Price2)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) CoveredOrUncovered() (*field.CoveredOrUncovered, error) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) DiscretionOffset() (*field.DiscretionOffset, error) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderSingle) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}
