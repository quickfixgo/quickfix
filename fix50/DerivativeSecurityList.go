package fix50

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

//NewDerivativeSecurityListBuilder returns an initialized DerivativeSecurityListBuilder with specified required fields.
func NewDerivativeSecurityListBuilder(
	securityreqid field.SecurityReqID,
	securityresponseid field.SecurityResponseID,
	securityrequestresult field.SecurityRequestResult) *DerivativeSecurityListBuilder {
	builder := new(DerivativeSecurityListBuilder)
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityrequestresult)
	return builder
}

//SecurityReqID is a required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseID is a required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityRequestResult is a required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSymbol() (*field.UnderlyingSymbol, error) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, error) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSecurityID() (*field.UnderlyingSecurityID, error) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, error) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, error) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingProduct is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingProduct() (*field.UnderlyingProduct, error) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCFICode() (*field.UnderlyingCFICode, error) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSecurityType() (*field.UnderlyingSecurityType, error) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, error) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, error) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, error) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, error) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingIssueDate() (*field.UnderlyingIssueDate, error) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, error) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, error) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, error) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFactor is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingFactor() (*field.UnderlyingFactor, error) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCreditRating() (*field.UnderlyingCreditRating, error) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, error) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, error) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, error) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, error) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, error) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, error) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, error) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, error) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, error) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCouponRate() (*field.UnderlyingCouponRate, error) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, error) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingIssuer() (*field.UnderlyingIssuer, error) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, error) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, error) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, error) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, error) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, error) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCPProgram() (*field.UnderlyingCPProgram, error) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCPRegType() (*field.UnderlyingCPRegType, error) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCurrency() (*field.UnderlyingCurrency, error) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingQty is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingQty() (*field.UnderlyingQty, error) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPx is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingPx() (*field.UnderlyingPx, error) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, error) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingEndPrice() (*field.UnderlyingEndPrice, error) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingStartValue() (*field.UnderlyingStartValue, error) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, error) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingEndValue() (*field.UnderlyingEndValue, error) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) NoUnderlyingStips() (*field.NoUnderlyingStips, error) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, error) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSettlementType() (*field.UnderlyingSettlementType, error) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCashAmount() (*field.UnderlyingCashAmount, error) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashType is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCashType() (*field.UnderlyingCashType, error) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, error) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, error) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingCapValue() (*field.UnderlyingCapValue, error) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, error) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, error) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, error) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingFXRate() (*field.UnderlyingFXRate, error) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, error) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for DerivativeSecurityList.
func (m *DerivativeSecurityList) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
