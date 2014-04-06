package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type OrderMassActionRequest struct {
	quickfix.Message
}

func (m *OrderMassActionRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, error) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingPx() (*field.UnderlyingPx, error) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, error) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, error) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, error) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, error) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingDetachmentPoint() (*field.UnderlyingDetachmentPoint, error) {
	f := new(field.UnderlyingDetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRate, error) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingFactor() (*field.UnderlyingFactor, error) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCashType() (*field.UnderlyingCashType, error) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, error) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, error) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingStartValue() (*field.UnderlyingStartValue, error) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementType, error) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCapValue() (*field.UnderlyingCapValue, error) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingAttachmentPoint() (*field.UnderlyingAttachmentPoint, error) {
	f := new(field.UnderlyingAttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSymbol() (*field.UnderlyingSymbol, error) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, error) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingOriginalNotionalPercentageOutstanding() (*field.UnderlyingOriginalNotionalPercentageOutstanding, error) {
	f := new(field.UnderlyingOriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, error) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, error) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingFXRate() (*field.UnderlyingFXRate, error) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, error) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, error) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmount, error) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, error) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDate, error) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, error) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCurrency() (*field.UnderlyingCurrency, error) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, error) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, error) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingContractMultiplierUnit() (*field.UnderlyingContractMultiplierUnit, error) {
	f := new(field.UnderlyingContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingQty() (*field.UnderlyingQty, error) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, error) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingIssuer() (*field.UnderlyingIssuer, error) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, error) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSeniority() (*field.UnderlyingSeniority, error) {
	f := new(field.UnderlyingSeniority)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MassActionScope() (*field.MassActionScope, error) {
	f := new(field.MassActionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, error) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoTargetPartyIDs() (*field.NoTargetPartyIDs, error) {
	f := new(field.NoTargetPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, error) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCFICode() (*field.UnderlyingCFICode, error) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, error) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, error) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, error) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, error) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityType, error) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, error) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, error) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingProduct() (*field.UnderlyingProduct, error) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, error) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, error) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, error) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, error) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, error) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingRestructuringType() (*field.UnderlyingRestructuringType, error) {
	f := new(field.UnderlyingRestructuringType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, error) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingFlowScheduleType() (*field.UnderlyingFlowScheduleType, error) {
	f := new(field.UnderlyingFlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) MassActionType() (*field.MassActionType, error) {
	f := new(field.MassActionType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, error) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, error) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgram, error) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, error) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, error) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, error) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegType, error) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingEndPrice() (*field.UnderlyingEndPrice, error) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) NoUnderlyingStips() (*field.NoUnderlyingStips, error) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, error) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingNotionalPercentageOutstanding() (*field.UnderlyingNotionalPercentageOutstanding, error) {
	f := new(field.UnderlyingNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityID, error) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, error) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, error) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRating, error) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassActionRequest) UnderlyingEndValue() (*field.UnderlyingEndValue, error) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}
