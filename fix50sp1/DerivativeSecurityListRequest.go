package fix50sp1

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

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, error) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, error) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, error) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, error) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, error) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, error) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
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

//MarketID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSymbol() (*field.DerivativeSymbol, error) {
	f := new(field.DerivativeSymbol)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSymbolSfx() (*field.DerivativeSymbolSfx, error) {
	f := new(field.DerivativeSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityID() (*field.DerivativeSecurityID, error) {
	f := new(field.DerivativeSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityIDSource() (*field.DerivativeSecurityIDSource, error) {
	f := new(field.DerivativeSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) NoDerivativeSecurityAltID() (*field.NoDerivativeSecurityAltID, error) {
	f := new(field.NoDerivativeSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeProduct is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeProduct() (*field.DerivativeProduct, error) {
	f := new(field.DerivativeProduct)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeProductComplex() (*field.DerivativeProductComplex, error) {
	f := new(field.DerivativeProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivFlexProductEligibilityIndicator() (*field.DerivFlexProductEligibilityIndicator, error) {
	f := new(field.DerivFlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityGroup() (*field.DerivativeSecurityGroup, error) {
	f := new(field.DerivativeSecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeCFICode() (*field.DerivativeCFICode, error) {
	f := new(field.DerivativeCFICode)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityType() (*field.DerivativeSecurityType, error) {
	f := new(field.DerivativeSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecuritySubType() (*field.DerivativeSecuritySubType, error) {
	f := new(field.DerivativeSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeMaturityMonthYear() (*field.DerivativeMaturityMonthYear, error) {
	f := new(field.DerivativeMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeMaturityDate() (*field.DerivativeMaturityDate, error) {
	f := new(field.DerivativeMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeMaturityTime() (*field.DerivativeMaturityTime, error) {
	f := new(field.DerivativeMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSettleOnOpenFlag() (*field.DerivativeSettleOnOpenFlag, error) {
	f := new(field.DerivativeSettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeInstrmtAssignmentMethod() (*field.DerivativeInstrmtAssignmentMethod, error) {
	f := new(field.DerivativeInstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityStatus() (*field.DerivativeSecurityStatus, error) {
	f := new(field.DerivativeSecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeIssueDate() (*field.DerivativeIssueDate, error) {
	f := new(field.DerivativeIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeInstrRegistry() (*field.DerivativeInstrRegistry, error) {
	f := new(field.DerivativeInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeCountryOfIssue() (*field.DerivativeCountryOfIssue, error) {
	f := new(field.DerivativeCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeStateOrProvinceOfIssue() (*field.DerivativeStateOrProvinceOfIssue, error) {
	f := new(field.DerivativeStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeStrikePrice() (*field.DerivativeStrikePrice, error) {
	f := new(field.DerivativeStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeLocaleOfIssue() (*field.DerivativeLocaleOfIssue, error) {
	f := new(field.DerivativeLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeStrikeCurrency() (*field.DerivativeStrikeCurrency, error) {
	f := new(field.DerivativeStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeStrikeMultiplier() (*field.DerivativeStrikeMultiplier, error) {
	f := new(field.DerivativeStrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeStrikeValue() (*field.DerivativeStrikeValue, error) {
	f := new(field.DerivativeStrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeOptAttribute() (*field.DerivativeOptAttribute, error) {
	f := new(field.DerivativeOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeContractMultiplier() (*field.DerivativeContractMultiplier, error) {
	f := new(field.DerivativeContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeMinPriceIncrement() (*field.DerivativeMinPriceIncrement, error) {
	f := new(field.DerivativeMinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeMinPriceIncrementAmount() (*field.DerivativeMinPriceIncrementAmount, error) {
	f := new(field.DerivativeMinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeUnitOfMeasure() (*field.DerivativeUnitOfMeasure, error) {
	f := new(field.DerivativeUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeUnitOfMeasureQty() (*field.DerivativeUnitOfMeasureQty, error) {
	f := new(field.DerivativeUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativePriceUnitOfMeasure() (*field.DerivativePriceUnitOfMeasure, error) {
	f := new(field.DerivativePriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativePriceUnitOfMeasureQty() (*field.DerivativePriceUnitOfMeasureQty, error) {
	f := new(field.DerivativePriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeExerciseStyle() (*field.DerivativeExerciseStyle, error) {
	f := new(field.DerivativeExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeOptPayAmount() (*field.DerivativeOptPayAmount, error) {
	f := new(field.DerivativeOptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeTimeUnit() (*field.DerivativeTimeUnit, error) {
	f := new(field.DerivativeTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityExchange() (*field.DerivativeSecurityExchange, error) {
	f := new(field.DerivativeSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativePositionLimit() (*field.DerivativePositionLimit, error) {
	f := new(field.DerivativePositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeNTPositionLimit() (*field.DerivativeNTPositionLimit, error) {
	f := new(field.DerivativeNTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeIssuer() (*field.DerivativeIssuer, error) {
	f := new(field.DerivativeIssuer)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeEncodedIssuerLen() (*field.DerivativeEncodedIssuerLen, error) {
	f := new(field.DerivativeEncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeEncodedIssuer() (*field.DerivativeEncodedIssuer, error) {
	f := new(field.DerivativeEncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityDesc() (*field.DerivativeSecurityDesc, error) {
	f := new(field.DerivativeSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeEncodedSecurityDescLen() (*field.DerivativeEncodedSecurityDescLen, error) {
	f := new(field.DerivativeEncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeEncodedSecurityDesc() (*field.DerivativeEncodedSecurityDesc, error) {
	f := new(field.DerivativeEncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeContractSettlMonth() (*field.DerivativeContractSettlMonth, error) {
	f := new(field.DerivativeContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) NoDerivativeEvents() (*field.NoDerivativeEvents, error) {
	f := new(field.NoDerivativeEvents)
	err := m.Body.Get(f)
	return f, err
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) NoDerivativeInstrumentParties() (*field.NoDerivativeInstrumentParties, error) {
	f := new(field.NoDerivativeInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSettlMethod() (*field.DerivativeSettlMethod, error) {
	f := new(field.DerivativeSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativePriceQuoteMethod() (*field.DerivativePriceQuoteMethod, error) {
	f := new(field.DerivativePriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeFuturesValuationMethod() (*field.DerivativeFuturesValuationMethod, error) {
	f := new(field.DerivativeFuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeListMethod() (*field.DerivativeListMethod, error) {
	f := new(field.DerivativeListMethod)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeCapPrice() (*field.DerivativeCapPrice, error) {
	f := new(field.DerivativeCapPrice)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeFloorPrice() (*field.DerivativeFloorPrice, error) {
	f := new(field.DerivativeFloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativePutOrCall() (*field.DerivativePutOrCall, error) {
	f := new(field.DerivativePutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityXMLLen() (*field.DerivativeSecurityXMLLen, error) {
	f := new(field.DerivativeSecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityXML() (*field.DerivativeSecurityXML, error) {
	f := new(field.DerivativeSecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListRequest.
func (m *DerivativeSecurityListRequest) DerivativeSecurityXMLSchema() (*field.DerivativeSecurityXMLSchema, error) {
	f := new(field.DerivativeSecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}
