package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type OrderCancelReplaceRequest struct {
	quickfix.Message
}

func (m *OrderCancelReplaceRequest) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DisplayWhen() (*field.DisplayWhen, error) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerPriceType() (*field.TriggerPriceType, error) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DiscretionMoveType() (*field.DiscretionMoveType, error) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerNewPrice() (*field.TriggerNewPrice, error) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CustOrderHandlingInst() (*field.CustOrderHandlingInst, error) {
	f := new(field.CustOrderHandlingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DisplayQty() (*field.DisplayQty, error) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DiscretionScope() (*field.DiscretionScope, error) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ReceivedDeptID() (*field.ReceivedDeptID, error) {
	f := new(field.ReceivedDeptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DayBookingInst() (*field.DayBookingInst, error) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegSecurityDesc() (*field.PegSecurityDesc, error) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncovered, error) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecondaryDisplayQty() (*field.SecondaryDisplayQty, error) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegSecurityIDSource() (*field.PegSecurityIDSource, error) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegMoveType() (*field.PegMoveType, error) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrderHandlingInstSource() (*field.OrderHandlingInstSource, error) {
	f := new(field.OrderHandlingInstSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerPrice() (*field.TriggerPrice, error) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerNewQty() (*field.TriggerNewQty, error) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegOffsetValue() (*field.PegOffsetValue, error) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerSymbol() (*field.TriggerSymbol, error) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerPriceDirection() (*field.TriggerPriceDirection, error) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrigOrdModTime() (*field.OrigOrdModTime, error) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DiscretionOffsetType() (*field.DiscretionOffsetType, error) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerAction() (*field.TriggerAction, error) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CashMargin() (*field.CashMargin, error) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegRoundDirection() (*field.PegRoundDirection, error) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TargetStrategyParameters() (*field.TargetStrategyParameters, error) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegPriceType() (*field.PegPriceType, error) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CommCurrency() (*field.CommCurrency, error) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RefreshQty() (*field.RefreshQty, error) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MatchIncrement() (*field.MatchIncrement, error) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SettlDate2() (*field.SettlDate2, error) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DisplayLowQty() (*field.DisplayLowQty, error) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DisplayMinIncr() (*field.DisplayMinIncr, error) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerType() (*field.TriggerType, error) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, error) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CustDirectedOrder() (*field.CustDirectedOrder, error) {
	f := new(field.CustDirectedOrder)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegOffsetType() (*field.PegOffsetType, error) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerTradingSessionID() (*field.TriggerTradingSessionID, error) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DisplayHighQty() (*field.DisplayHighQty, error) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegScope() (*field.PegScope, error) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DiscretionRoundDirection() (*field.DiscretionRoundDirection, error) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegLimitType() (*field.PegLimitType, error) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DiscretionLimitType() (*field.DiscretionLimitType, error) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Price2() (*field.Price2, error) {
	f := new(field.Price2)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, error) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoStrategyParameters() (*field.NoStrategyParameters, error) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegSymbol() (*field.PegSymbol, error) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PriceProtectionScope() (*field.PriceProtectionScope, error) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ManualOrderIndicator() (*field.ManualOrderIndicator, error) {
	f := new(field.ManualOrderIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BookingUnit() (*field.BookingUnit, error) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) FundRenewWaiv() (*field.FundRenewWaiv, error) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerSecurityID() (*field.TriggerSecurityID, error) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerOrderType() (*field.TriggerOrderType, error) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PreTradeAnonymity() (*field.PreTradeAnonymity, error) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, error) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DiscretionOffsetValue() (*field.DiscretionOffsetValue, error) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TriggerSecurityDesc() (*field.TriggerSecurityDesc, error) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PegSecurityID() (*field.PegSecurityID, error) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, error) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ParticipationRate() (*field.ParticipationRate, error) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) MaxPriceLevels() (*field.MaxPriceLevels, error) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) TargetStrategy() (*field.TargetStrategy, error) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DisplayMethod() (*field.DisplayMethod, error) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderCancelReplaceRequest) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}
