package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//DerivativeSecurityList msg type = AA.
type DerivativeSecurityList struct {
	message.Message
}

//DerivativeSecurityListBuilder builds DerivativeSecurityList messages.
type DerivativeSecurityListBuilder struct {
	message.MessageBuilder
}

//CreateDerivativeSecurityListBuilder returns an initialized DerivativeSecurityListBuilder with specified required fields.
func CreateDerivativeSecurityListBuilder() DerivativeSecurityListBuilder {
	var builder DerivativeSecurityListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//SecurityReqID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) SecurityResponseID() (field.SecurityResponseID, error) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) SecurityRequestResult() (field.SecurityRequestResult, error) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSymbol() (field.UnderlyingSymbol, error) {
	var f field.UnderlyingSymbol
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSymbolSfx() (field.UnderlyingSymbolSfx, error) {
	var f field.UnderlyingSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityID() (field.UnderlyingSecurityID, error) {
	var f field.UnderlyingSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityIDSource() (field.UnderlyingSecurityIDSource, error) {
	var f field.UnderlyingSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoUnderlyingSecurityAltID() (field.NoUnderlyingSecurityAltID, error) {
	var f field.NoUnderlyingSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingProduct is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingProduct() (field.UnderlyingProduct, error) {
	var f field.UnderlyingProduct
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCFICode() (field.UnderlyingCFICode, error) {
	var f field.UnderlyingCFICode
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityType() (field.UnderlyingSecurityType, error) {
	var f field.UnderlyingSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecuritySubType() (field.UnderlyingSecuritySubType, error) {
	var f field.UnderlyingSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingMaturityMonthYear() (field.UnderlyingMaturityMonthYear, error) {
	var f field.UnderlyingMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingMaturityDate() (field.UnderlyingMaturityDate, error) {
	var f field.UnderlyingMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCouponPaymentDate() (field.UnderlyingCouponPaymentDate, error) {
	var f field.UnderlyingCouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingIssueDate() (field.UnderlyingIssueDate, error) {
	var f field.UnderlyingIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRepoCollateralSecurityType() (field.UnderlyingRepoCollateralSecurityType, error) {
	var f field.UnderlyingRepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRepurchaseTerm() (field.UnderlyingRepurchaseTerm, error) {
	var f field.UnderlyingRepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRepurchaseRate() (field.UnderlyingRepurchaseRate, error) {
	var f field.UnderlyingRepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFactor is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingFactor() (field.UnderlyingFactor, error) {
	var f field.UnderlyingFactor
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCreditRating() (field.UnderlyingCreditRating, error) {
	var f field.UnderlyingCreditRating
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingInstrRegistry() (field.UnderlyingInstrRegistry, error) {
	var f field.UnderlyingInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCountryOfIssue() (field.UnderlyingCountryOfIssue, error) {
	var f field.UnderlyingCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStateOrProvinceOfIssue() (field.UnderlyingStateOrProvinceOfIssue, error) {
	var f field.UnderlyingStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingLocaleOfIssue() (field.UnderlyingLocaleOfIssue, error) {
	var f field.UnderlyingLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRedemptionDate() (field.UnderlyingRedemptionDate, error) {
	var f field.UnderlyingRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStrikePrice() (field.UnderlyingStrikePrice, error) {
	var f field.UnderlyingStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStrikeCurrency() (field.UnderlyingStrikeCurrency, error) {
	var f field.UnderlyingStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingOptAttribute() (field.UnderlyingOptAttribute, error) {
	var f field.UnderlyingOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingContractMultiplier() (field.UnderlyingContractMultiplier, error) {
	var f field.UnderlyingContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCouponRate() (field.UnderlyingCouponRate, error) {
	var f field.UnderlyingCouponRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityExchange() (field.UnderlyingSecurityExchange, error) {
	var f field.UnderlyingSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingIssuer() (field.UnderlyingIssuer, error) {
	var f field.UnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingIssuerLen() (field.EncodedUnderlyingIssuerLen, error) {
	var f field.EncodedUnderlyingIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingIssuer() (field.EncodedUnderlyingIssuer, error) {
	var f field.EncodedUnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityDesc() (field.UnderlyingSecurityDesc, error) {
	var f field.UnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingSecurityDescLen() (field.EncodedUnderlyingSecurityDescLen, error) {
	var f field.EncodedUnderlyingSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingSecurityDesc() (field.EncodedUnderlyingSecurityDesc, error) {
	var f field.EncodedUnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCPProgram() (field.UnderlyingCPProgram, error) {
	var f field.UnderlyingCPProgram
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCPRegType() (field.UnderlyingCPRegType, error) {
	var f field.UnderlyingCPRegType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCurrency() (field.UnderlyingCurrency, error) {
	var f field.UnderlyingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingQty() (field.UnderlyingQty, error) {
	var f field.UnderlyingQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPx is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPx() (field.UnderlyingPx, error) {
	var f field.UnderlyingPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingDirtyPrice() (field.UnderlyingDirtyPrice, error) {
	var f field.UnderlyingDirtyPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingEndPrice() (field.UnderlyingEndPrice, error) {
	var f field.UnderlyingEndPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStartValue() (field.UnderlyingStartValue, error) {
	var f field.UnderlyingStartValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCurrentValue() (field.UnderlyingCurrentValue, error) {
	var f field.UnderlyingCurrentValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingEndValue() (field.UnderlyingEndValue, error) {
	var f field.UnderlyingEndValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoUnderlyingStips() (field.NoUnderlyingStips, error) {
	var f field.NoUnderlyingStips
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingAllocationPercent() (field.UnderlyingAllocationPercent, error) {
	var f field.UnderlyingAllocationPercent
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSettlementType() (field.UnderlyingSettlementType, error) {
	var f field.UnderlyingSettlementType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCashAmount() (field.UnderlyingCashAmount, error) {
	var f field.UnderlyingCashAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCashType() (field.UnderlyingCashType, error) {
	var f field.UnderlyingCashType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingUnitOfMeasure() (field.UnderlyingUnitOfMeasure, error) {
	var f field.UnderlyingUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingTimeUnit() (field.UnderlyingTimeUnit, error) {
	var f field.UnderlyingTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCapValue() (field.UnderlyingCapValue, error) {
	var f field.UnderlyingCapValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoUndlyInstrumentParties() (field.NoUndlyInstrumentParties, error) {
	var f field.NoUndlyInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSettlMethod() (field.UnderlyingSettlMethod, error) {
	var f field.UnderlyingSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingAdjustedQuantity() (field.UnderlyingAdjustedQuantity, error) {
	var f field.UnderlyingAdjustedQuantity
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingFXRate() (field.UnderlyingFXRate, error) {
	var f field.UnderlyingFXRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingFXRateCalc() (field.UnderlyingFXRateCalc, error) {
	var f field.UnderlyingFXRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingMaturityTime() (field.UnderlyingMaturityTime, error) {
	var f field.UnderlyingMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPutOrCall() (field.UnderlyingPutOrCall, error) {
	var f field.UnderlyingPutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingExerciseStyle() (field.UnderlyingExerciseStyle, error) {
	var f field.UnderlyingExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingUnitOfMeasureQty() (field.UnderlyingUnitOfMeasureQty, error) {
	var f field.UnderlyingUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPriceUnitOfMeasure() (field.UnderlyingPriceUnitOfMeasure, error) {
	var f field.UnderlyingPriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPriceUnitOfMeasureQty() (field.UnderlyingPriceUnitOfMeasureQty, error) {
	var f field.UnderlyingPriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplierUnit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingContractMultiplierUnit() (field.UnderlyingContractMultiplierUnit, error) {
	var f field.UnderlyingContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFlowScheduleType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingFlowScheduleType() (field.UnderlyingFlowScheduleType, error) {
	var f field.UnderlyingFlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRestructuringType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRestructuringType() (field.UnderlyingRestructuringType, error) {
	var f field.UnderlyingRestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSeniority is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSeniority() (field.UnderlyingSeniority, error) {
	var f field.UnderlyingSeniority
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingNotionalPercentageOutstanding() (field.UnderlyingNotionalPercentageOutstanding, error) {
	var f field.UnderlyingNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingOriginalNotionalPercentageOutstanding() (field.UnderlyingOriginalNotionalPercentageOutstanding, error) {
	var f field.UnderlyingOriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAttachmentPoint is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingAttachmentPoint() (field.UnderlyingAttachmentPoint, error) {
	var f field.UnderlyingAttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDetachmentPoint is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingDetachmentPoint() (field.UnderlyingDetachmentPoint, error) {
	var f field.UnderlyingDetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) TotNoRelatedSym() (field.TotNoRelatedSym, error) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbol is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSymbol() (field.DerivativeSymbol, error) {
	var f field.DerivativeSymbol
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSymbolSfx() (field.DerivativeSymbolSfx, error) {
	var f field.DerivativeSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityID() (field.DerivativeSecurityID, error) {
	var f field.DerivativeSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityIDSource() (field.DerivativeSecurityIDSource, error) {
	var f field.DerivativeSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeSecurityAltID() (field.NoDerivativeSecurityAltID, error) {
	var f field.NoDerivativeSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProduct is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeProduct() (field.DerivativeProduct, error) {
	var f field.DerivativeProduct
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeProductComplex() (field.DerivativeProductComplex, error) {
	var f field.DerivativeProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivFlexProductEligibilityIndicator() (field.DerivFlexProductEligibilityIndicator, error) {
	var f field.DerivFlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityGroup() (field.DerivativeSecurityGroup, error) {
	var f field.DerivativeSecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCFICode is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeCFICode() (field.DerivativeCFICode, error) {
	var f field.DerivativeCFICode
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityType() (field.DerivativeSecurityType, error) {
	var f field.DerivativeSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecuritySubType() (field.DerivativeSecuritySubType, error) {
	var f field.DerivativeSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMaturityMonthYear() (field.DerivativeMaturityMonthYear, error) {
	var f field.DerivativeMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMaturityDate() (field.DerivativeMaturityDate, error) {
	var f field.DerivativeMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMaturityTime() (field.DerivativeMaturityTime, error) {
	var f field.DerivativeMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSettleOnOpenFlag() (field.DerivativeSettleOnOpenFlag, error) {
	var f field.DerivativeSettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeInstrmtAssignmentMethod() (field.DerivativeInstrmtAssignmentMethod, error) {
	var f field.DerivativeInstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityStatus() (field.DerivativeSecurityStatus, error) {
	var f field.DerivativeSecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeIssueDate() (field.DerivativeIssueDate, error) {
	var f field.DerivativeIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeInstrRegistry() (field.DerivativeInstrRegistry, error) {
	var f field.DerivativeInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeCountryOfIssue() (field.DerivativeCountryOfIssue, error) {
	var f field.DerivativeCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStateOrProvinceOfIssue() (field.DerivativeStateOrProvinceOfIssue, error) {
	var f field.DerivativeStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikePrice() (field.DerivativeStrikePrice, error) {
	var f field.DerivativeStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeLocaleOfIssue() (field.DerivativeLocaleOfIssue, error) {
	var f field.DerivativeLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikeCurrency() (field.DerivativeStrikeCurrency, error) {
	var f field.DerivativeStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikeMultiplier() (field.DerivativeStrikeMultiplier, error) {
	var f field.DerivativeStrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikeValue() (field.DerivativeStrikeValue, error) {
	var f field.DerivativeStrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeOptAttribute() (field.DerivativeOptAttribute, error) {
	var f field.DerivativeOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeContractMultiplier() (field.DerivativeContractMultiplier, error) {
	var f field.DerivativeContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMinPriceIncrement() (field.DerivativeMinPriceIncrement, error) {
	var f field.DerivativeMinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMinPriceIncrementAmount() (field.DerivativeMinPriceIncrementAmount, error) {
	var f field.DerivativeMinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeUnitOfMeasure() (field.DerivativeUnitOfMeasure, error) {
	var f field.DerivativeUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeUnitOfMeasureQty() (field.DerivativeUnitOfMeasureQty, error) {
	var f field.DerivativeUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePriceUnitOfMeasure() (field.DerivativePriceUnitOfMeasure, error) {
	var f field.DerivativePriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePriceUnitOfMeasureQty() (field.DerivativePriceUnitOfMeasureQty, error) {
	var f field.DerivativePriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeExerciseStyle() (field.DerivativeExerciseStyle, error) {
	var f field.DerivativeExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeOptPayAmount() (field.DerivativeOptPayAmount, error) {
	var f field.DerivativeOptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeTimeUnit() (field.DerivativeTimeUnit, error) {
	var f field.DerivativeTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityExchange() (field.DerivativeSecurityExchange, error) {
	var f field.DerivativeSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePositionLimit() (field.DerivativePositionLimit, error) {
	var f field.DerivativePositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeNTPositionLimit() (field.DerivativeNTPositionLimit, error) {
	var f field.DerivativeNTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeIssuer() (field.DerivativeIssuer, error) {
	var f field.DerivativeIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedIssuerLen() (field.DerivativeEncodedIssuerLen, error) {
	var f field.DerivativeEncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedIssuer() (field.DerivativeEncodedIssuer, error) {
	var f field.DerivativeEncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityDesc() (field.DerivativeSecurityDesc, error) {
	var f field.DerivativeSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedSecurityDescLen() (field.DerivativeEncodedSecurityDescLen, error) {
	var f field.DerivativeEncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedSecurityDesc() (field.DerivativeEncodedSecurityDesc, error) {
	var f field.DerivativeEncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeContractSettlMonth() (field.DerivativeContractSettlMonth, error) {
	var f field.DerivativeContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeEvents() (field.NoDerivativeEvents, error) {
	var f field.NoDerivativeEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeInstrumentParties() (field.NoDerivativeInstrumentParties, error) {
	var f field.NoDerivativeInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSettlMethod() (field.DerivativeSettlMethod, error) {
	var f field.DerivativeSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePriceQuoteMethod() (field.DerivativePriceQuoteMethod, error) {
	var f field.DerivativePriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeValuationMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeValuationMethod() (field.DerivativeValuationMethod, error) {
	var f field.DerivativeValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeListMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeListMethod() (field.DerivativeListMethod, error) {
	var f field.DerivativeListMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeCapPrice() (field.DerivativeCapPrice, error) {
	var f field.DerivativeCapPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeFloorPrice() (field.DerivativeFloorPrice, error) {
	var f field.DerivativeFloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePutOrCall() (field.DerivativePutOrCall, error) {
	var f field.DerivativePutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityXMLLen() (field.DerivativeSecurityXMLLen, error) {
	var f field.DerivativeSecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityXML() (field.DerivativeSecurityXML, error) {
	var f field.DerivativeSecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityXMLSchema() (field.DerivativeSecurityXMLSchema, error) {
	var f field.DerivativeSecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractMultiplierUnit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeContractMultiplierUnit() (field.DerivativeContractMultiplierUnit, error) {
	var f field.DerivativeContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFlowScheduleType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeFlowScheduleType() (field.DerivativeFlowScheduleType, error) {
	var f field.DerivativeFlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeInstrAttrib() (field.NoDerivativeInstrAttrib, error) {
	var f field.NoDerivativeInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//NoMarketSegments is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoMarketSegments() (field.NoMarketSegments, error) {
	var f field.NoMarketSegments
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//SecurityReportID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) SecurityReportID() (field.SecurityReportID, error) {
	var f field.SecurityReportID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}
