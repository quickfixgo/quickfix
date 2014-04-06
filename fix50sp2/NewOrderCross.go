package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type NewOrderCross struct {
	quickfix.Message
}

func (m *NewOrderCross) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegScope() (*field.PegScope, error) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DiscretionMoveType() (*field.DiscretionMoveType, error) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoStrategyParameters() (*field.NoStrategyParameters, error) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DisplayHighQty() (*field.DisplayHighQty, error) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerSecurityDesc() (*field.TriggerSecurityDesc, error) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TargetStrategy() (*field.TargetStrategy, error) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerSymbol() (*field.TriggerSymbol, error) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegRoundDirection() (*field.PegRoundDirection, error) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CrossID() (*field.CrossID, error) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DiscretionOffsetValue() (*field.DiscretionOffsetValue, error) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DisplayMethod() (*field.DisplayMethod, error) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DisplayLowQty() (*field.DisplayLowQty, error) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerPriceDirection() (*field.TriggerPriceDirection, error) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerSecurityID() (*field.TriggerSecurityID, error) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerNewQty() (*field.TriggerNewQty, error) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DiscretionLimitType() (*field.DiscretionLimitType, error) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DisplayQty() (*field.DisplayQty, error) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegOffsetType() (*field.PegOffsetType, error) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegSecurityID() (*field.PegSecurityID, error) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TransBkdTime() (*field.TransBkdTime, error) {
	f := new(field.TransBkdTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, error) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegSecurityIDSource() (*field.PegSecurityIDSource, error) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DiscretionOffsetType() (*field.DiscretionOffsetType, error) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DiscretionRoundDirection() (*field.DiscretionRoundDirection, error) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerType() (*field.TriggerType, error) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CrossType() (*field.CrossType, error) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegPriceType() (*field.PegPriceType, error) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DisplayMinIncr() (*field.DisplayMinIncr, error) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RefreshQty() (*field.RefreshQty, error) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoSides() (*field.NoSides, error) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaxPriceLevels() (*field.MaxPriceLevels, error) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecondaryDisplayQty() (*field.SecondaryDisplayQty, error) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegMoveType() (*field.PegMoveType, error) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CrossPrioritization() (*field.CrossPrioritization, error) {
	f := new(field.CrossPrioritization)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegSymbol() (*field.PegSymbol, error) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MatchIncrement() (*field.MatchIncrement, error) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerTradingSessionID() (*field.TriggerTradingSessionID, error) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TargetStrategyParameters() (*field.TargetStrategyParameters, error) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerAction() (*field.TriggerAction, error) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) IOIID() (*field.IOIID, error) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ParticipationRate() (*field.ParticipationRate, error) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DiscretionScope() (*field.DiscretionScope, error) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerNewPrice() (*field.TriggerNewPrice, error) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegSecurityDesc() (*field.PegSecurityDesc, error) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, error) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerPriceType() (*field.TriggerPriceType, error) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, error) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegLimitType() (*field.PegLimitType, error) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoRootPartyIDs() (*field.NoRootPartyIDs, error) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PriceProtectionScope() (*field.PriceProtectionScope, error) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerPrice() (*field.TriggerPrice, error) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TriggerOrderType() (*field.TriggerOrderType, error) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PegOffsetValue() (*field.PegOffsetValue, error) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) DisplayWhen() (*field.DisplayWhen, error) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
