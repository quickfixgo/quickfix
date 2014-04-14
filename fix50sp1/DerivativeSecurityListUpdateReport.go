package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//DerivativeSecurityListUpdateReport msg type = BR.
type DerivativeSecurityListUpdateReport struct {
	message.Message
}

//DerivativeSecurityListUpdateReportBuilder builds DerivativeSecurityListUpdateReport messages.
type DerivativeSecurityListUpdateReportBuilder struct {
	message.MessageBuilder
}

//NewDerivativeSecurityListUpdateReportBuilder returns an initialized DerivativeSecurityListUpdateReportBuilder with specified required fields.
func NewDerivativeSecurityListUpdateReportBuilder() *DerivativeSecurityListUpdateReportBuilder {
	builder := new(DerivativeSecurityListUpdateReportBuilder)
	return builder
}

//SecurityReqID is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseID is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityRequestResult is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//SecurityUpdateAction is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, error) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSymbol() (*field.UnderlyingSymbol, error) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, error) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSecurityID() (*field.UnderlyingSecurityID, error) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, error) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, error) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingProduct() (*field.UnderlyingProduct, error) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCFICode() (*field.UnderlyingCFICode, error) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSecurityType() (*field.UnderlyingSecurityType, error) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, error) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, error) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, error) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, error) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingIssueDate() (*field.UnderlyingIssueDate, error) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, error) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, error) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, error) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingFactor() (*field.UnderlyingFactor, error) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCreditRating() (*field.UnderlyingCreditRating, error) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, error) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, error) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, error) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, error) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, error) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, error) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, error) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, error) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, error) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCouponRate() (*field.UnderlyingCouponRate, error) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, error) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingIssuer() (*field.UnderlyingIssuer, error) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, error) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, error) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, error) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, error) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, error) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCPProgram() (*field.UnderlyingCPProgram, error) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCPRegType() (*field.UnderlyingCPRegType, error) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCurrency() (*field.UnderlyingCurrency, error) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingQty() (*field.UnderlyingQty, error) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPx is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingPx() (*field.UnderlyingPx, error) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, error) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingEndPrice() (*field.UnderlyingEndPrice, error) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingStartValue() (*field.UnderlyingStartValue, error) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, error) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingEndValue() (*field.UnderlyingEndValue, error) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoUnderlyingStips() (*field.NoUnderlyingStips, error) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, error) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSettlementType() (*field.UnderlyingSettlementType, error) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCashAmount() (*field.UnderlyingCashAmount, error) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCashType() (*field.UnderlyingCashType, error) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, error) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, error) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingCapValue() (*field.UnderlyingCapValue, error) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, error) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, error) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, error) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingFXRate() (*field.UnderlyingFXRate, error) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, error) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, error) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, error) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, error) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, error) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, error) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, error) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSymbol() (*field.DerivativeSymbol, error) {
	f := new(field.DerivativeSymbol)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSymbolSfx() (*field.DerivativeSymbolSfx, error) {
	f := new(field.DerivativeSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityID() (*field.DerivativeSecurityID, error) {
	f := new(field.DerivativeSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityIDSource() (*field.DerivativeSecurityIDSource, error) {
	f := new(field.DerivativeSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoDerivativeSecurityAltID() (*field.NoDerivativeSecurityAltID, error) {
	f := new(field.NoDerivativeSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeProduct() (*field.DerivativeProduct, error) {
	f := new(field.DerivativeProduct)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeProductComplex() (*field.DerivativeProductComplex, error) {
	f := new(field.DerivativeProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivFlexProductEligibilityIndicator() (*field.DerivFlexProductEligibilityIndicator, error) {
	f := new(field.DerivFlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityGroup() (*field.DerivativeSecurityGroup, error) {
	f := new(field.DerivativeSecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeCFICode() (*field.DerivativeCFICode, error) {
	f := new(field.DerivativeCFICode)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityType() (*field.DerivativeSecurityType, error) {
	f := new(field.DerivativeSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecuritySubType() (*field.DerivativeSecuritySubType, error) {
	f := new(field.DerivativeSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeMaturityMonthYear() (*field.DerivativeMaturityMonthYear, error) {
	f := new(field.DerivativeMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeMaturityDate() (*field.DerivativeMaturityDate, error) {
	f := new(field.DerivativeMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeMaturityTime() (*field.DerivativeMaturityTime, error) {
	f := new(field.DerivativeMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSettleOnOpenFlag() (*field.DerivativeSettleOnOpenFlag, error) {
	f := new(field.DerivativeSettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeInstrmtAssignmentMethod() (*field.DerivativeInstrmtAssignmentMethod, error) {
	f := new(field.DerivativeInstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityStatus() (*field.DerivativeSecurityStatus, error) {
	f := new(field.DerivativeSecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeIssueDate() (*field.DerivativeIssueDate, error) {
	f := new(field.DerivativeIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeInstrRegistry() (*field.DerivativeInstrRegistry, error) {
	f := new(field.DerivativeInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeCountryOfIssue() (*field.DerivativeCountryOfIssue, error) {
	f := new(field.DerivativeCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeStateOrProvinceOfIssue() (*field.DerivativeStateOrProvinceOfIssue, error) {
	f := new(field.DerivativeStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeStrikePrice() (*field.DerivativeStrikePrice, error) {
	f := new(field.DerivativeStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeLocaleOfIssue() (*field.DerivativeLocaleOfIssue, error) {
	f := new(field.DerivativeLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeStrikeCurrency() (*field.DerivativeStrikeCurrency, error) {
	f := new(field.DerivativeStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeStrikeMultiplier() (*field.DerivativeStrikeMultiplier, error) {
	f := new(field.DerivativeStrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeStrikeValue() (*field.DerivativeStrikeValue, error) {
	f := new(field.DerivativeStrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeOptAttribute() (*field.DerivativeOptAttribute, error) {
	f := new(field.DerivativeOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeContractMultiplier() (*field.DerivativeContractMultiplier, error) {
	f := new(field.DerivativeContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrement() (*field.DerivativeMinPriceIncrement, error) {
	f := new(field.DerivativeMinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrementAmount() (*field.DerivativeMinPriceIncrementAmount, error) {
	f := new(field.DerivativeMinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasure() (*field.DerivativeUnitOfMeasure, error) {
	f := new(field.DerivativeUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasureQty() (*field.DerivativeUnitOfMeasureQty, error) {
	f := new(field.DerivativeUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasure() (*field.DerivativePriceUnitOfMeasure, error) {
	f := new(field.DerivativePriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasureQty() (*field.DerivativePriceUnitOfMeasureQty, error) {
	f := new(field.DerivativePriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeExerciseStyle() (*field.DerivativeExerciseStyle, error) {
	f := new(field.DerivativeExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeOptPayAmount() (*field.DerivativeOptPayAmount, error) {
	f := new(field.DerivativeOptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeTimeUnit() (*field.DerivativeTimeUnit, error) {
	f := new(field.DerivativeTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityExchange() (*field.DerivativeSecurityExchange, error) {
	f := new(field.DerivativeSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativePositionLimit() (*field.DerivativePositionLimit, error) {
	f := new(field.DerivativePositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeNTPositionLimit() (*field.DerivativeNTPositionLimit, error) {
	f := new(field.DerivativeNTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeIssuer() (*field.DerivativeIssuer, error) {
	f := new(field.DerivativeIssuer)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeEncodedIssuerLen() (*field.DerivativeEncodedIssuerLen, error) {
	f := new(field.DerivativeEncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeEncodedIssuer() (*field.DerivativeEncodedIssuer, error) {
	f := new(field.DerivativeEncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityDesc() (*field.DerivativeSecurityDesc, error) {
	f := new(field.DerivativeSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDescLen() (*field.DerivativeEncodedSecurityDescLen, error) {
	f := new(field.DerivativeEncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDesc() (*field.DerivativeEncodedSecurityDesc, error) {
	f := new(field.DerivativeEncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeContractSettlMonth() (*field.DerivativeContractSettlMonth, error) {
	f := new(field.DerivativeContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoDerivativeEvents() (*field.NoDerivativeEvents, error) {
	f := new(field.NoDerivativeEvents)
	err := m.Body.Get(f)
	return f, err
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoDerivativeInstrumentParties() (*field.NoDerivativeInstrumentParties, error) {
	f := new(field.NoDerivativeInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSettlMethod() (*field.DerivativeSettlMethod, error) {
	f := new(field.DerivativeSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativePriceQuoteMethod() (*field.DerivativePriceQuoteMethod, error) {
	f := new(field.DerivativePriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeFuturesValuationMethod() (*field.DerivativeFuturesValuationMethod, error) {
	f := new(field.DerivativeFuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeListMethod() (*field.DerivativeListMethod, error) {
	f := new(field.DerivativeListMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeCapPrice() (*field.DerivativeCapPrice, error) {
	f := new(field.DerivativeCapPrice)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeFloorPrice() (*field.DerivativeFloorPrice, error) {
	f := new(field.DerivativeFloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativePutOrCall() (*field.DerivativePutOrCall, error) {
	f := new(field.DerivativePutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityXMLLen() (*field.DerivativeSecurityXMLLen, error) {
	f := new(field.DerivativeSecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityXML() (*field.DerivativeSecurityXML, error) {
	f := new(field.DerivativeSecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) DerivativeSecurityXMLSchema() (*field.DerivativeSecurityXMLSchema, error) {
	f := new(field.DerivativeSecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoDerivativeInstrAttrib() (*field.NoDerivativeInstrAttrib, error) {
	f := new(field.NoDerivativeInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//NoMarketSegments is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoMarketSegments() (*field.NoMarketSegments, error) {
	f := new(field.NoMarketSegments)
	err := m.Body.Get(f)
	return f, err
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m *DerivativeSecurityListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
