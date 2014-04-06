package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type MultilegOrderCancelReplace struct {
	quickfix.Message
}

func (m *MultilegOrderCancelReplace) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoStrategyParameters() (*field.NoStrategyParameters, error) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DisplayMethod() (*field.DisplayMethod, error) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegLimitType() (*field.PegLimitType, error) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DiscretionScope() (*field.DiscretionScope, error) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RefreshQty() (*field.RefreshQty, error) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RiskFreeRate() (*field.RiskFreeRate, error) {
	f := new(field.RiskFreeRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CashMargin() (*field.CashMargin, error) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerType() (*field.TriggerType, error) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DiscretionOffsetValue() (*field.DiscretionOffsetValue, error) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerSecurityID() (*field.TriggerSecurityID, error) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, error) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) IOIID() (*field.IOIID, error) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TargetStrategy() (*field.TargetStrategy, error) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MultilegPriceMethod() (*field.MultilegPriceMethod, error) {
	f := new(field.MultilegPriceMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrigOrdModTime() (*field.OrigOrdModTime, error) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) FundRenewWaiv() (*field.FundRenewWaiv, error) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerAction() (*field.TriggerAction, error) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegSecurityID() (*field.PegSecurityID, error) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegSecurityDesc() (*field.PegSecurityDesc, error) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DisplayHighQty() (*field.DisplayHighQty, error) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerPriceDirection() (*field.TriggerPriceDirection, error) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) BookingUnit() (*field.BookingUnit, error) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ParticipationRate() (*field.ParticipationRate, error) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DisplayMinIncr() (*field.DisplayMinIncr, error) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegOffsetValue() (*field.PegOffsetValue, error) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerPrice() (*field.TriggerPrice, error) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegSecurityIDSource() (*field.PegSecurityIDSource, error) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegRoundDirection() (*field.PegRoundDirection, error) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DiscretionLimitType() (*field.DiscretionLimitType, error) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegMoveType() (*field.PegMoveType, error) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DiscretionOffsetType() (*field.DiscretionOffsetType, error) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DisplayQty() (*field.DisplayQty, error) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegOffsetType() (*field.PegOffsetType, error) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegPriceType() (*field.PegPriceType, error) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TargetStrategyParameters() (*field.TargetStrategyParameters, error) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerPriceType() (*field.TriggerPriceType, error) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerOrderType() (*field.TriggerOrderType, error) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CommCurrency() (*field.CommCurrency, error) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerSecurityDesc() (*field.TriggerSecurityDesc, error) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PreTradeAnonymity() (*field.PreTradeAnonymity, error) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MatchIncrement() (*field.MatchIncrement, error) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegScope() (*field.PegScope, error) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PegSymbol() (*field.PegSymbol, error) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DiscretionRoundDirection() (*field.DiscretionRoundDirection, error) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerTradingSessionID() (*field.TriggerTradingSessionID, error) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SwapPoints() (*field.SwapPoints, error) {
	f := new(field.SwapPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DiscretionMoveType() (*field.DiscretionMoveType, error) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DisplayLowQty() (*field.DisplayLowQty, error) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerSymbol() (*field.TriggerSymbol, error) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerNewQty() (*field.TriggerNewQty, error) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MultilegModel() (*field.MultilegModel, error) {
	f := new(field.MultilegModel)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CoveredOrUncovered() (*field.CoveredOrUncovered, error) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DayBookingInst() (*field.DayBookingInst, error) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MaxPriceLevels() (*field.MaxPriceLevels, error) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, error) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) DisplayWhen() (*field.DisplayWhen, error) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SecondaryDisplayQty() (*field.SecondaryDisplayQty, error) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerNewPrice() (*field.TriggerNewPrice, error) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) MultiLegRptTypeReq() (*field.MultiLegRptTypeReq, error) {
	f := new(field.MultiLegRptTypeReq)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, error) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *MultilegOrderCancelReplace) PriceProtectionScope() (*field.PriceProtectionScope, error) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}
