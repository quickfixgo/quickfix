package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//DerivativeSecurityListRequest msg type = z.
type DerivativeSecurityListRequest struct {
	message.Message
}

//DerivativeSecurityListRequestBuilder builds DerivativeSecurityListRequest messages.
type DerivativeSecurityListRequestBuilder struct {
	message.MessageBuilder
}

//NewDerivativeSecurityListRequestBuilder returns an initialized DerivativeSecurityListRequestBuilder with specified required fields.
func NewDerivativeSecurityListRequestBuilder(
	securityreqid field.SecurityReqID,
	securitylistrequesttype field.SecurityListRequestType) *DerivativeSecurityListRequestBuilder {
	builder := new(DerivativeSecurityListRequestBuilder)
	builder.Body.Set(securityreqid)
	builder.Body.Set(securitylistrequesttype)
	return builder
}

//SecurityReqID is a required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListRequestType is a required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) SecurityListRequestType() (*field.SecurityListRequestType, error) {
	f := new(field.SecurityListRequestType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSymbol() (*field.UnderlyingSymbol, error) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, error) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityID, error) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, error) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, error) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingProduct() (*field.UnderlyingProduct, error) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCFICode() (*field.UnderlyingCFICode, error) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityType, error) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, error) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, error) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, error) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, error) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDate, error) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, error) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, error) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, error) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingFactor() (*field.UnderlyingFactor, error) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRating, error) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, error) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, error) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, error) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, error) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, error) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, error) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, error) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, error) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, error) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRate, error) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, error) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingIssuer() (*field.UnderlyingIssuer, error) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, error) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, error) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, error) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, error) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, error) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgram, error) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegType, error) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCurrency() (*field.UnderlyingCurrency, error) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingQty is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingQty() (*field.UnderlyingQty, error) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPx is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingPx() (*field.UnderlyingPx, error) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, error) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingEndPrice() (*field.UnderlyingEndPrice, error) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingStartValue() (*field.UnderlyingStartValue, error) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, error) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingEndValue() (*field.UnderlyingEndValue, error) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) NoUnderlyingStips() (*field.NoUnderlyingStips, error) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, error) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementType, error) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmount, error) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCashType() (*field.UnderlyingCashType, error) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, error) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, error) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingCapValue() (*field.UnderlyingCapValue, error) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, error) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, error) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, error) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingFXRate() (*field.UnderlyingFXRate, error) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, error) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
