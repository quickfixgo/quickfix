package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ExecutionReport struct {
	quickfixgo.Message
}

func (m *ExecutionReport) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) CrossID() (*field.CrossID, error) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrigCrossID() (*field.OrigCrossID, error) {
	f := new(field.OrigCrossID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) CrossType() (*field.CrossType, error) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecRefID() (*field.ExecRefID, error) {
	f := new(field.ExecRefID)
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
func (m *ExecutionReport) WorkingIndicator() (*field.WorkingIndicator, error) {
	f := new(field.WorkingIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrdRejReason() (*field.OrdRejReason, error) {
	f := new(field.OrdRejReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) DayBookingInst() (*field.DayBookingInst, error) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) BookingUnit() (*field.BookingUnit, error) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) CashMargin() (*field.CashMargin, error) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) QuantityType() (*field.QuantityType, error) {
	f := new(field.QuantityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) DiscretionOffset() (*field.DiscretionOffset, error) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Rule80A() (*field.Rule80A, error) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) UnderlyingLastQty() (*field.UnderlyingLastQty, error) {
	f := new(field.UnderlyingLastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) UnderlyingLastPx() (*field.UnderlyingLastPx, error) {
	f := new(field.UnderlyingLastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastCapacity() (*field.LastCapacity, error) {
	f := new(field.LastCapacity)
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
func (m *ExecutionReport) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) DayOrderQty() (*field.DayOrderQty, error) {
	f := new(field.DayOrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) DayCumQty() (*field.DayCumQty, error) {
	f := new(field.DayCumQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) DayAvgPx() (*field.DayAvgPx, error) {
	f := new(field.DayAvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ReportToExch() (*field.ReportToExch, error) {
	f := new(field.ReportToExch)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExDate() (*field.ExDate, error) {
	f := new(field.ExDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TradedFlatSwitch() (*field.TradedFlatSwitch, error) {
	f := new(field.TradedFlatSwitch)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) BasisFeatureDate() (*field.BasisFeatureDate, error) {
	f := new(field.BasisFeatureDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) BasisFeaturePrice() (*field.BasisFeaturePrice, error) {
	f := new(field.BasisFeaturePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmt, error) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRate, error) {
	f := new(field.SettlCurrFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) FutSettDate2() (*field.FutSettDate2, error) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) LastForwardPoints2() (*field.LastForwardPoints2, error) {
	f := new(field.LastForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) TransBkdTime() (*field.TransBkdTime, error) {
	f := new(field.TransBkdTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecValuationPoint() (*field.ExecValuationPoint, error) {
	f := new(field.ExecValuationPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecPriceType() (*field.ExecPriceType, error) {
	f := new(field.ExecPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) ExecPriceAdjustment() (*field.ExecPriceAdjustment, error) {
	f := new(field.ExecPriceAdjustment)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PriorityIndicator() (*field.PriorityIndicator, error) {
	f := new(field.PriorityIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionReport) PriceImprovement() (*field.PriceImprovement, error) {
	f := new(field.PriceImprovement)
	err := m.Body.Get(f)
	return f, err
}
