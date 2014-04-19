package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
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
func (m DerivativeSecurityListUpdateReport) SecurityReqID() (field.SecurityReqID, errors.MessageRejectError) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityResponseID() (field.SecurityResponseID, errors.MessageRejectError) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityRequestResult() (field.SecurityRequestResult, errors.MessageRejectError) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//SecurityUpdateAction is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityUpdateAction() (field.SecurityUpdateAction, errors.MessageRejectError) {
	var f field.SecurityUpdateAction
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbol() (field.UnderlyingSymbol, errors.MessageRejectError) {
	var f field.UnderlyingSymbol
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbolSfx() (field.UnderlyingSymbolSfx, errors.MessageRejectError) {
	var f field.UnderlyingSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityID() (field.UnderlyingSecurityID, errors.MessageRejectError) {
	var f field.UnderlyingSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityIDSource() (field.UnderlyingSecurityIDSource, errors.MessageRejectError) {
	var f field.UnderlyingSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingSecurityAltID() (field.NoUnderlyingSecurityAltID, errors.MessageRejectError) {
	var f field.NoUnderlyingSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingProduct() (field.UnderlyingProduct, errors.MessageRejectError) {
	var f field.UnderlyingProduct
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCFICode() (field.UnderlyingCFICode, errors.MessageRejectError) {
	var f field.UnderlyingCFICode
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityType() (field.UnderlyingSecurityType, errors.MessageRejectError) {
	var f field.UnderlyingSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecuritySubType() (field.UnderlyingSecuritySubType, errors.MessageRejectError) {
	var f field.UnderlyingSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityMonthYear() (field.UnderlyingMaturityMonthYear, errors.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityDate() (field.UnderlyingMaturityDate, errors.MessageRejectError) {
	var f field.UnderlyingMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponPaymentDate() (field.UnderlyingCouponPaymentDate, errors.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssueDate() (field.UnderlyingIssueDate, errors.MessageRejectError) {
	var f field.UnderlyingIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepoCollateralSecurityType() (field.UnderlyingRepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseTerm() (field.UnderlyingRepurchaseTerm, errors.MessageRejectError) {
	var f field.UnderlyingRepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseRate() (field.UnderlyingRepurchaseRate, errors.MessageRejectError) {
	var f field.UnderlyingRepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFactor() (field.UnderlyingFactor, errors.MessageRejectError) {
	var f field.UnderlyingFactor
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCreditRating() (field.UnderlyingCreditRating, errors.MessageRejectError) {
	var f field.UnderlyingCreditRating
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingInstrRegistry() (field.UnderlyingInstrRegistry, errors.MessageRejectError) {
	var f field.UnderlyingInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCountryOfIssue() (field.UnderlyingCountryOfIssue, errors.MessageRejectError) {
	var f field.UnderlyingCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStateOrProvinceOfIssue() (field.UnderlyingStateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingLocaleOfIssue() (field.UnderlyingLocaleOfIssue, errors.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRedemptionDate() (field.UnderlyingRedemptionDate, errors.MessageRejectError) {
	var f field.UnderlyingRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikePrice() (field.UnderlyingStrikePrice, errors.MessageRejectError) {
	var f field.UnderlyingStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikeCurrency() (field.UnderlyingStrikeCurrency, errors.MessageRejectError) {
	var f field.UnderlyingStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingOptAttribute() (field.UnderlyingOptAttribute, errors.MessageRejectError) {
	var f field.UnderlyingOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingContractMultiplier() (field.UnderlyingContractMultiplier, errors.MessageRejectError) {
	var f field.UnderlyingContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponRate() (field.UnderlyingCouponRate, errors.MessageRejectError) {
	var f field.UnderlyingCouponRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityExchange() (field.UnderlyingSecurityExchange, errors.MessageRejectError) {
	var f field.UnderlyingSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssuer() (field.UnderlyingIssuer, errors.MessageRejectError) {
	var f field.UnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuerLen() (field.EncodedUnderlyingIssuerLen, errors.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuer() (field.EncodedUnderlyingIssuer, errors.MessageRejectError) {
	var f field.EncodedUnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityDesc() (field.UnderlyingSecurityDesc, errors.MessageRejectError) {
	var f field.UnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDescLen() (field.EncodedUnderlyingSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDesc() (field.EncodedUnderlyingSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPProgram() (field.UnderlyingCPProgram, errors.MessageRejectError) {
	var f field.UnderlyingCPProgram
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPRegType() (field.UnderlyingCPRegType, errors.MessageRejectError) {
	var f field.UnderlyingCPRegType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrency() (field.UnderlyingCurrency, errors.MessageRejectError) {
	var f field.UnderlyingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingQty() (field.UnderlyingQty, errors.MessageRejectError) {
	var f field.UnderlyingQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPx() (field.UnderlyingPx, errors.MessageRejectError) {
	var f field.UnderlyingPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingDirtyPrice() (field.UnderlyingDirtyPrice, errors.MessageRejectError) {
	var f field.UnderlyingDirtyPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndPrice() (field.UnderlyingEndPrice, errors.MessageRejectError) {
	var f field.UnderlyingEndPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStartValue() (field.UnderlyingStartValue, errors.MessageRejectError) {
	var f field.UnderlyingStartValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrentValue() (field.UnderlyingCurrentValue, errors.MessageRejectError) {
	var f field.UnderlyingCurrentValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndValue() (field.UnderlyingEndValue, errors.MessageRejectError) {
	var f field.UnderlyingEndValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingStips() (field.NoUnderlyingStips, errors.MessageRejectError) {
	var f field.NoUnderlyingStips
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAllocationPercent() (field.UnderlyingAllocationPercent, errors.MessageRejectError) {
	var f field.UnderlyingAllocationPercent
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlementType() (field.UnderlyingSettlementType, errors.MessageRejectError) {
	var f field.UnderlyingSettlementType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashAmount() (field.UnderlyingCashAmount, errors.MessageRejectError) {
	var f field.UnderlyingCashAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashType() (field.UnderlyingCashType, errors.MessageRejectError) {
	var f field.UnderlyingCashType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasure() (field.UnderlyingUnitOfMeasure, errors.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingTimeUnit() (field.UnderlyingTimeUnit, errors.MessageRejectError) {
	var f field.UnderlyingTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCapValue() (field.UnderlyingCapValue, errors.MessageRejectError) {
	var f field.UnderlyingCapValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUndlyInstrumentParties() (field.NoUndlyInstrumentParties, errors.MessageRejectError) {
	var f field.NoUndlyInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlMethod() (field.UnderlyingSettlMethod, errors.MessageRejectError) {
	var f field.UnderlyingSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAdjustedQuantity() (field.UnderlyingAdjustedQuantity, errors.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantity
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRate() (field.UnderlyingFXRate, errors.MessageRejectError) {
	var f field.UnderlyingFXRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRateCalc() (field.UnderlyingFXRateCalc, errors.MessageRejectError) {
	var f field.UnderlyingFXRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityTime() (field.UnderlyingMaturityTime, errors.MessageRejectError) {
	var f field.UnderlyingMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPutOrCall() (field.UnderlyingPutOrCall, errors.MessageRejectError) {
	var f field.UnderlyingPutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingExerciseStyle() (field.UnderlyingExerciseStyle, errors.MessageRejectError) {
	var f field.UnderlyingExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasureQty() (field.UnderlyingUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasure() (field.UnderlyingPriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasureQty() (field.UnderlyingPriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbol() (field.DerivativeSymbol, errors.MessageRejectError) {
	var f field.DerivativeSymbol
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbolSfx() (field.DerivativeSymbolSfx, errors.MessageRejectError) {
	var f field.DerivativeSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityID() (field.DerivativeSecurityID, errors.MessageRejectError) {
	var f field.DerivativeSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityIDSource() (field.DerivativeSecurityIDSource, errors.MessageRejectError) {
	var f field.DerivativeSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeSecurityAltID() (field.NoDerivativeSecurityAltID, errors.MessageRejectError) {
	var f field.NoDerivativeSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProduct() (field.DerivativeProduct, errors.MessageRejectError) {
	var f field.DerivativeProduct
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProductComplex() (field.DerivativeProductComplex, errors.MessageRejectError) {
	var f field.DerivativeProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivFlexProductEligibilityIndicator() (field.DerivFlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.DerivFlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityGroup() (field.DerivativeSecurityGroup, errors.MessageRejectError) {
	var f field.DerivativeSecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCFICode() (field.DerivativeCFICode, errors.MessageRejectError) {
	var f field.DerivativeCFICode
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityType() (field.DerivativeSecurityType, errors.MessageRejectError) {
	var f field.DerivativeSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecuritySubType() (field.DerivativeSecuritySubType, errors.MessageRejectError) {
	var f field.DerivativeSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityMonthYear() (field.DerivativeMaturityMonthYear, errors.MessageRejectError) {
	var f field.DerivativeMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityDate() (field.DerivativeMaturityDate, errors.MessageRejectError) {
	var f field.DerivativeMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityTime() (field.DerivativeMaturityTime, errors.MessageRejectError) {
	var f field.DerivativeMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettleOnOpenFlag() (field.DerivativeSettleOnOpenFlag, errors.MessageRejectError) {
	var f field.DerivativeSettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrmtAssignmentMethod() (field.DerivativeInstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.DerivativeInstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityStatus() (field.DerivativeSecurityStatus, errors.MessageRejectError) {
	var f field.DerivativeSecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssueDate() (field.DerivativeIssueDate, errors.MessageRejectError) {
	var f field.DerivativeIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrRegistry() (field.DerivativeInstrRegistry, errors.MessageRejectError) {
	var f field.DerivativeInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCountryOfIssue() (field.DerivativeCountryOfIssue, errors.MessageRejectError) {
	var f field.DerivativeCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStateOrProvinceOfIssue() (field.DerivativeStateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.DerivativeStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikePrice() (field.DerivativeStrikePrice, errors.MessageRejectError) {
	var f field.DerivativeStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeLocaleOfIssue() (field.DerivativeLocaleOfIssue, errors.MessageRejectError) {
	var f field.DerivativeLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeCurrency() (field.DerivativeStrikeCurrency, errors.MessageRejectError) {
	var f field.DerivativeStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeMultiplier() (field.DerivativeStrikeMultiplier, errors.MessageRejectError) {
	var f field.DerivativeStrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeValue() (field.DerivativeStrikeValue, errors.MessageRejectError) {
	var f field.DerivativeStrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptAttribute() (field.DerivativeOptAttribute, errors.MessageRejectError) {
	var f field.DerivativeOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractMultiplier() (field.DerivativeContractMultiplier, errors.MessageRejectError) {
	var f field.DerivativeContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrement() (field.DerivativeMinPriceIncrement, errors.MessageRejectError) {
	var f field.DerivativeMinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrementAmount() (field.DerivativeMinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.DerivativeMinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasure() (field.DerivativeUnitOfMeasure, errors.MessageRejectError) {
	var f field.DerivativeUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasureQty() (field.DerivativeUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.DerivativeUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasure() (field.DerivativePriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.DerivativePriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasureQty() (field.DerivativePriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.DerivativePriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeExerciseStyle() (field.DerivativeExerciseStyle, errors.MessageRejectError) {
	var f field.DerivativeExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptPayAmount() (field.DerivativeOptPayAmount, errors.MessageRejectError) {
	var f field.DerivativeOptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeTimeUnit() (field.DerivativeTimeUnit, errors.MessageRejectError) {
	var f field.DerivativeTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityExchange() (field.DerivativeSecurityExchange, errors.MessageRejectError) {
	var f field.DerivativeSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePositionLimit() (field.DerivativePositionLimit, errors.MessageRejectError) {
	var f field.DerivativePositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeNTPositionLimit() (field.DerivativeNTPositionLimit, errors.MessageRejectError) {
	var f field.DerivativeNTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssuer() (field.DerivativeIssuer, errors.MessageRejectError) {
	var f field.DerivativeIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuerLen() (field.DerivativeEncodedIssuerLen, errors.MessageRejectError) {
	var f field.DerivativeEncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuer() (field.DerivativeEncodedIssuer, errors.MessageRejectError) {
	var f field.DerivativeEncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityDesc() (field.DerivativeSecurityDesc, errors.MessageRejectError) {
	var f field.DerivativeSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDescLen() (field.DerivativeEncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.DerivativeEncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDesc() (field.DerivativeEncodedSecurityDesc, errors.MessageRejectError) {
	var f field.DerivativeEncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractSettlMonth() (field.DerivativeContractSettlMonth, errors.MessageRejectError) {
	var f field.DerivativeContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeEvents() (field.NoDerivativeEvents, errors.MessageRejectError) {
	var f field.NoDerivativeEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrumentParties() (field.NoDerivativeInstrumentParties, errors.MessageRejectError) {
	var f field.NoDerivativeInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettlMethod() (field.DerivativeSettlMethod, errors.MessageRejectError) {
	var f field.DerivativeSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceQuoteMethod() (field.DerivativePriceQuoteMethod, errors.MessageRejectError) {
	var f field.DerivativePriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFuturesValuationMethod() (field.DerivativeFuturesValuationMethod, errors.MessageRejectError) {
	var f field.DerivativeFuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeListMethod() (field.DerivativeListMethod, errors.MessageRejectError) {
	var f field.DerivativeListMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCapPrice() (field.DerivativeCapPrice, errors.MessageRejectError) {
	var f field.DerivativeCapPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFloorPrice() (field.DerivativeFloorPrice, errors.MessageRejectError) {
	var f field.DerivativeFloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePutOrCall() (field.DerivativePutOrCall, errors.MessageRejectError) {
	var f field.DerivativePutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLLen() (field.DerivativeSecurityXMLLen, errors.MessageRejectError) {
	var f field.DerivativeSecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXML() (field.DerivativeSecurityXML, errors.MessageRejectError) {
	var f field.DerivativeSecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLSchema() (field.DerivativeSecurityXMLSchema, errors.MessageRejectError) {
	var f field.DerivativeSecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrAttrib() (field.NoDerivativeInstrAttrib, errors.MessageRejectError) {
	var f field.NoDerivativeInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//NoMarketSegments is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoMarketSegments() (field.NoMarketSegments, errors.MessageRejectError) {
	var f field.NoMarketSegments
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) TotNoRelatedSym() (field.TotNoRelatedSym, errors.MessageRejectError) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) LastFragment() (field.LastFragment, errors.MessageRejectError) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplID() (field.ApplID, errors.MessageRejectError) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplSeqNum() (field.ApplSeqNum, errors.MessageRejectError) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplLastSeqNum() (field.ApplLastSeqNum, errors.MessageRejectError) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplResendFlag() (field.ApplResendFlag, errors.MessageRejectError) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
