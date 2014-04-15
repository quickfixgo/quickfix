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

//CreateDerivativeSecurityListUpdateReportBuilder returns an initialized DerivativeSecurityListUpdateReportBuilder with specified required fields.
func CreateDerivativeSecurityListUpdateReportBuilder() DerivativeSecurityListUpdateReportBuilder {
	var builder DerivativeSecurityListUpdateReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//SecurityReqID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityResponseID() (field.SecurityResponseID, error) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityRequestResult() (field.SecurityRequestResult, error) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//SecurityUpdateAction is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityUpdateAction() (field.SecurityUpdateAction, error) {
	var f field.SecurityUpdateAction
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbol() (field.UnderlyingSymbol, error) {
	var f field.UnderlyingSymbol
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbolSfx() (field.UnderlyingSymbolSfx, error) {
	var f field.UnderlyingSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityID() (field.UnderlyingSecurityID, error) {
	var f field.UnderlyingSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityIDSource() (field.UnderlyingSecurityIDSource, error) {
	var f field.UnderlyingSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingSecurityAltID() (field.NoUnderlyingSecurityAltID, error) {
	var f field.NoUnderlyingSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingProduct() (field.UnderlyingProduct, error) {
	var f field.UnderlyingProduct
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCFICode() (field.UnderlyingCFICode, error) {
	var f field.UnderlyingCFICode
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityType() (field.UnderlyingSecurityType, error) {
	var f field.UnderlyingSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecuritySubType() (field.UnderlyingSecuritySubType, error) {
	var f field.UnderlyingSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityMonthYear() (field.UnderlyingMaturityMonthYear, error) {
	var f field.UnderlyingMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityDate() (field.UnderlyingMaturityDate, error) {
	var f field.UnderlyingMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponPaymentDate() (field.UnderlyingCouponPaymentDate, error) {
	var f field.UnderlyingCouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssueDate() (field.UnderlyingIssueDate, error) {
	var f field.UnderlyingIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepoCollateralSecurityType() (field.UnderlyingRepoCollateralSecurityType, error) {
	var f field.UnderlyingRepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseTerm() (field.UnderlyingRepurchaseTerm, error) {
	var f field.UnderlyingRepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseRate() (field.UnderlyingRepurchaseRate, error) {
	var f field.UnderlyingRepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFactor() (field.UnderlyingFactor, error) {
	var f field.UnderlyingFactor
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCreditRating() (field.UnderlyingCreditRating, error) {
	var f field.UnderlyingCreditRating
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingInstrRegistry() (field.UnderlyingInstrRegistry, error) {
	var f field.UnderlyingInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCountryOfIssue() (field.UnderlyingCountryOfIssue, error) {
	var f field.UnderlyingCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStateOrProvinceOfIssue() (field.UnderlyingStateOrProvinceOfIssue, error) {
	var f field.UnderlyingStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingLocaleOfIssue() (field.UnderlyingLocaleOfIssue, error) {
	var f field.UnderlyingLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRedemptionDate() (field.UnderlyingRedemptionDate, error) {
	var f field.UnderlyingRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikePrice() (field.UnderlyingStrikePrice, error) {
	var f field.UnderlyingStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikeCurrency() (field.UnderlyingStrikeCurrency, error) {
	var f field.UnderlyingStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingOptAttribute() (field.UnderlyingOptAttribute, error) {
	var f field.UnderlyingOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingContractMultiplier() (field.UnderlyingContractMultiplier, error) {
	var f field.UnderlyingContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponRate() (field.UnderlyingCouponRate, error) {
	var f field.UnderlyingCouponRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityExchange() (field.UnderlyingSecurityExchange, error) {
	var f field.UnderlyingSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssuer() (field.UnderlyingIssuer, error) {
	var f field.UnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuerLen() (field.EncodedUnderlyingIssuerLen, error) {
	var f field.EncodedUnderlyingIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuer() (field.EncodedUnderlyingIssuer, error) {
	var f field.EncodedUnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityDesc() (field.UnderlyingSecurityDesc, error) {
	var f field.UnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDescLen() (field.EncodedUnderlyingSecurityDescLen, error) {
	var f field.EncodedUnderlyingSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDesc() (field.EncodedUnderlyingSecurityDesc, error) {
	var f field.EncodedUnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPProgram() (field.UnderlyingCPProgram, error) {
	var f field.UnderlyingCPProgram
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPRegType() (field.UnderlyingCPRegType, error) {
	var f field.UnderlyingCPRegType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrency() (field.UnderlyingCurrency, error) {
	var f field.UnderlyingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingQty() (field.UnderlyingQty, error) {
	var f field.UnderlyingQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPx() (field.UnderlyingPx, error) {
	var f field.UnderlyingPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingDirtyPrice() (field.UnderlyingDirtyPrice, error) {
	var f field.UnderlyingDirtyPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndPrice() (field.UnderlyingEndPrice, error) {
	var f field.UnderlyingEndPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStartValue() (field.UnderlyingStartValue, error) {
	var f field.UnderlyingStartValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrentValue() (field.UnderlyingCurrentValue, error) {
	var f field.UnderlyingCurrentValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndValue() (field.UnderlyingEndValue, error) {
	var f field.UnderlyingEndValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingStips() (field.NoUnderlyingStips, error) {
	var f field.NoUnderlyingStips
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAllocationPercent() (field.UnderlyingAllocationPercent, error) {
	var f field.UnderlyingAllocationPercent
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlementType() (field.UnderlyingSettlementType, error) {
	var f field.UnderlyingSettlementType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashAmount() (field.UnderlyingCashAmount, error) {
	var f field.UnderlyingCashAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashType() (field.UnderlyingCashType, error) {
	var f field.UnderlyingCashType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasure() (field.UnderlyingUnitOfMeasure, error) {
	var f field.UnderlyingUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingTimeUnit() (field.UnderlyingTimeUnit, error) {
	var f field.UnderlyingTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCapValue() (field.UnderlyingCapValue, error) {
	var f field.UnderlyingCapValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUndlyInstrumentParties() (field.NoUndlyInstrumentParties, error) {
	var f field.NoUndlyInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlMethod() (field.UnderlyingSettlMethod, error) {
	var f field.UnderlyingSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAdjustedQuantity() (field.UnderlyingAdjustedQuantity, error) {
	var f field.UnderlyingAdjustedQuantity
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRate() (field.UnderlyingFXRate, error) {
	var f field.UnderlyingFXRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRateCalc() (field.UnderlyingFXRateCalc, error) {
	var f field.UnderlyingFXRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityTime() (field.UnderlyingMaturityTime, error) {
	var f field.UnderlyingMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPutOrCall() (field.UnderlyingPutOrCall, error) {
	var f field.UnderlyingPutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingExerciseStyle() (field.UnderlyingExerciseStyle, error) {
	var f field.UnderlyingExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasureQty() (field.UnderlyingUnitOfMeasureQty, error) {
	var f field.UnderlyingUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasure() (field.UnderlyingPriceUnitOfMeasure, error) {
	var f field.UnderlyingPriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasureQty() (field.UnderlyingPriceUnitOfMeasureQty, error) {
	var f field.UnderlyingPriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbol() (field.DerivativeSymbol, error) {
	var f field.DerivativeSymbol
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbolSfx() (field.DerivativeSymbolSfx, error) {
	var f field.DerivativeSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityID() (field.DerivativeSecurityID, error) {
	var f field.DerivativeSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityIDSource() (field.DerivativeSecurityIDSource, error) {
	var f field.DerivativeSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeSecurityAltID() (field.NoDerivativeSecurityAltID, error) {
	var f field.NoDerivativeSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProduct() (field.DerivativeProduct, error) {
	var f field.DerivativeProduct
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProductComplex() (field.DerivativeProductComplex, error) {
	var f field.DerivativeProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivFlexProductEligibilityIndicator() (field.DerivFlexProductEligibilityIndicator, error) {
	var f field.DerivFlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityGroup() (field.DerivativeSecurityGroup, error) {
	var f field.DerivativeSecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCFICode() (field.DerivativeCFICode, error) {
	var f field.DerivativeCFICode
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityType() (field.DerivativeSecurityType, error) {
	var f field.DerivativeSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecuritySubType() (field.DerivativeSecuritySubType, error) {
	var f field.DerivativeSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityMonthYear() (field.DerivativeMaturityMonthYear, error) {
	var f field.DerivativeMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityDate() (field.DerivativeMaturityDate, error) {
	var f field.DerivativeMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityTime() (field.DerivativeMaturityTime, error) {
	var f field.DerivativeMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettleOnOpenFlag() (field.DerivativeSettleOnOpenFlag, error) {
	var f field.DerivativeSettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrmtAssignmentMethod() (field.DerivativeInstrmtAssignmentMethod, error) {
	var f field.DerivativeInstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityStatus() (field.DerivativeSecurityStatus, error) {
	var f field.DerivativeSecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssueDate() (field.DerivativeIssueDate, error) {
	var f field.DerivativeIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrRegistry() (field.DerivativeInstrRegistry, error) {
	var f field.DerivativeInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCountryOfIssue() (field.DerivativeCountryOfIssue, error) {
	var f field.DerivativeCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStateOrProvinceOfIssue() (field.DerivativeStateOrProvinceOfIssue, error) {
	var f field.DerivativeStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikePrice() (field.DerivativeStrikePrice, error) {
	var f field.DerivativeStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeLocaleOfIssue() (field.DerivativeLocaleOfIssue, error) {
	var f field.DerivativeLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeCurrency() (field.DerivativeStrikeCurrency, error) {
	var f field.DerivativeStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeMultiplier() (field.DerivativeStrikeMultiplier, error) {
	var f field.DerivativeStrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeValue() (field.DerivativeStrikeValue, error) {
	var f field.DerivativeStrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptAttribute() (field.DerivativeOptAttribute, error) {
	var f field.DerivativeOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractMultiplier() (field.DerivativeContractMultiplier, error) {
	var f field.DerivativeContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrement() (field.DerivativeMinPriceIncrement, error) {
	var f field.DerivativeMinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrementAmount() (field.DerivativeMinPriceIncrementAmount, error) {
	var f field.DerivativeMinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasure() (field.DerivativeUnitOfMeasure, error) {
	var f field.DerivativeUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasureQty() (field.DerivativeUnitOfMeasureQty, error) {
	var f field.DerivativeUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasure() (field.DerivativePriceUnitOfMeasure, error) {
	var f field.DerivativePriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasureQty() (field.DerivativePriceUnitOfMeasureQty, error) {
	var f field.DerivativePriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeExerciseStyle() (field.DerivativeExerciseStyle, error) {
	var f field.DerivativeExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptPayAmount() (field.DerivativeOptPayAmount, error) {
	var f field.DerivativeOptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeTimeUnit() (field.DerivativeTimeUnit, error) {
	var f field.DerivativeTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityExchange() (field.DerivativeSecurityExchange, error) {
	var f field.DerivativeSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePositionLimit() (field.DerivativePositionLimit, error) {
	var f field.DerivativePositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeNTPositionLimit() (field.DerivativeNTPositionLimit, error) {
	var f field.DerivativeNTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssuer() (field.DerivativeIssuer, error) {
	var f field.DerivativeIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuerLen() (field.DerivativeEncodedIssuerLen, error) {
	var f field.DerivativeEncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuer() (field.DerivativeEncodedIssuer, error) {
	var f field.DerivativeEncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityDesc() (field.DerivativeSecurityDesc, error) {
	var f field.DerivativeSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDescLen() (field.DerivativeEncodedSecurityDescLen, error) {
	var f field.DerivativeEncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDesc() (field.DerivativeEncodedSecurityDesc, error) {
	var f field.DerivativeEncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractSettlMonth() (field.DerivativeContractSettlMonth, error) {
	var f field.DerivativeContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeEvents() (field.NoDerivativeEvents, error) {
	var f field.NoDerivativeEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrumentParties() (field.NoDerivativeInstrumentParties, error) {
	var f field.NoDerivativeInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettlMethod() (field.DerivativeSettlMethod, error) {
	var f field.DerivativeSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceQuoteMethod() (field.DerivativePriceQuoteMethod, error) {
	var f field.DerivativePriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFuturesValuationMethod() (field.DerivativeFuturesValuationMethod, error) {
	var f field.DerivativeFuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeListMethod() (field.DerivativeListMethod, error) {
	var f field.DerivativeListMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCapPrice() (field.DerivativeCapPrice, error) {
	var f field.DerivativeCapPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFloorPrice() (field.DerivativeFloorPrice, error) {
	var f field.DerivativeFloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePutOrCall() (field.DerivativePutOrCall, error) {
	var f field.DerivativePutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLLen() (field.DerivativeSecurityXMLLen, error) {
	var f field.DerivativeSecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXML() (field.DerivativeSecurityXML, error) {
	var f field.DerivativeSecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLSchema() (field.DerivativeSecurityXMLSchema, error) {
	var f field.DerivativeSecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrAttrib() (field.NoDerivativeInstrAttrib, error) {
	var f field.NoDerivativeInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//NoMarketSegments is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoMarketSegments() (field.NoMarketSegments, error) {
	var f field.NoMarketSegments
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) TotNoRelatedSym() (field.TotNoRelatedSym, error) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
