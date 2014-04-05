package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type OrderMassCancelReport struct {
	quickfixgo.Message
}

func (m *OrderMassCancelReport) TotalAffectedOrders() (*field.TotalAffectedOrders, error) {
	f := new(field.TotalAffectedOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoAffectedOrders() (*field.NoAffectedOrders, error) {
	f := new(field.NoAffectedOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, error) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, error) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, error) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, error) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MassCancelResponse() (*field.MassCancelResponse, error) {
	f := new(field.MassCancelResponse)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCFICode() (*field.UnderlyingCFICode, error) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, error) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, error) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, error) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, error) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, error) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, error) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCPRegType() (*field.UnderlyingCPRegType, error) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) FuturesValuationMethod() (*field.FuturesValuationMethod, error) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, error) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, error) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSecurityID() (*field.UnderlyingSecurityID, error) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingProduct() (*field.UnderlyingProduct, error) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, error) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, error) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoUnderlyingStips() (*field.NoUnderlyingStips, error) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) OptPayAmount() (*field.OptPayAmount, error) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCashType() (*field.UnderlyingCashType, error) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, error) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingIssuer() (*field.UnderlyingIssuer, error) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, error) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, error) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, error) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, error) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, error) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoNotAffectedOrders() (*field.NoNotAffectedOrders, error) {
	f := new(field.NoNotAffectedOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MassCancelRequestType() (*field.MassCancelRequestType, error) {
	f := new(field.MassCancelRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, error) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingQty() (*field.UnderlyingQty, error) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, error) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, error) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingStartValue() (*field.UnderlyingStartValue, error) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, error) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, error) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, error) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCouponRate() (*field.UnderlyingCouponRate, error) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCapValue() (*field.UnderlyingCapValue, error) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, error) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSecurityType() (*field.UnderlyingSecurityType, error) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingPx() (*field.UnderlyingPx, error) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, error) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, error) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, error) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, error) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MassCancelRejectReason() (*field.MassCancelRejectReason, error) {
	f := new(field.MassCancelRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, error) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MassActionReportID() (*field.MassActionReportID, error) {
	f := new(field.MassActionReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, error) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSettlementType() (*field.UnderlyingSettlementType, error) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingEndValue() (*field.UnderlyingEndValue, error) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingFactor() (*field.UnderlyingFactor, error) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, error) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCashAmount() (*field.UnderlyingCashAmount, error) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingSymbol() (*field.UnderlyingSymbol, error) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCPProgram() (*field.UnderlyingCPProgram, error) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, error) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, error) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCreditRating() (*field.UnderlyingCreditRating, error) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, error) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingIssueDate() (*field.UnderlyingIssueDate, error) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingCurrency() (*field.UnderlyingCurrency, error) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, error) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, error) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, error) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingEndPrice() (*field.UnderlyingEndPrice, error) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) UnderlyingFXRate() (*field.UnderlyingFXRate, error) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}
