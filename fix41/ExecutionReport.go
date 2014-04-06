package fix41

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ExecutionReport struct {
	quickfix.Message
}

func (m *ExecutionReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecTransType() (*field.ExecTransType, error) {
	f := new(field.ExecTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastShares() (*field.LastShares, error) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRate, error) {
	f := new(field.SettlCurrFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmt, error) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrdRejReason() (*field.OrdRejReason, error) {
	f := new(field.OrdRejReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ReportToExch() (*field.ReportToExch, error) {
	f := new(field.ReportToExch)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecRefID() (*field.ExecRefID, error) {
	f := new(field.ExecRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastCapacity() (*field.LastCapacity, error) {
	f := new(field.LastCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LeavesQty() (*field.LeavesQty, error) {
	f := new(field.LeavesQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) CumQty() (*field.CumQty, error) {
	f := new(field.CumQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Rule80A() (*field.Rule80A, error) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}
