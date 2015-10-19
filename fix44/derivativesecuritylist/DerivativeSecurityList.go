//Package derivativesecuritylist msg type = AA.
package derivativesecuritylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a DerivativeSecurityList wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SecurityReqID is a required field for DerivativeSecurityList.
func (m Message) SecurityReqID() (*field.SecurityReqIDField, quickfix.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from DerivativeSecurityList.
func (m Message) GetSecurityReqID(f *field.SecurityReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for DerivativeSecurityList.
func (m Message) SecurityResponseID() (*field.SecurityResponseIDField, quickfix.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from DerivativeSecurityList.
func (m Message) GetSecurityResponseID(f *field.SecurityResponseIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a required field for DerivativeSecurityList.
func (m Message) SecurityRequestResult() (*field.SecurityRequestResultField, quickfix.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from DerivativeSecurityList.
func (m Message) GetSecurityRequestResult(f *field.SecurityRequestResultField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSymbol() (*field.UnderlyingSymbolField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from DerivativeSecurityList.
func (m Message) GetUnderlyingSymbol(f *field.UnderlyingSymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfxField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from DerivativeSecurityList.
func (m Message) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSecurityID() (*field.UnderlyingSecurityIDField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from DerivativeSecurityList.
func (m Message) GetUnderlyingSecurityID(f *field.UnderlyingSecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from DerivativeSecurityList.
func (m Message) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityList.
func (m Message) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from DerivativeSecurityList.
func (m Message) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingProduct() (*field.UnderlyingProductField, quickfix.MessageRejectError) {
	f := &field.UnderlyingProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from DerivativeSecurityList.
func (m Message) GetUnderlyingProduct(f *field.UnderlyingProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCFICode() (*field.UnderlyingCFICodeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from DerivativeSecurityList.
func (m Message) GetUnderlyingCFICode(f *field.UnderlyingCFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSecurityType() (*field.UnderlyingSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from DerivativeSecurityList.
func (m Message) GetUnderlyingSecurityType(f *field.UnderlyingSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from DerivativeSecurityList.
func (m Message) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.UnderlyingMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from DerivativeSecurityList.
func (m Message) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingMaturityDate() (*field.UnderlyingMaturityDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from DerivativeSecurityList.
func (m Message) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from DerivativeSecurityList.
func (m Message) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingIssueDate() (*field.UnderlyingIssueDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from DerivativeSecurityList.
func (m Message) GetUnderlyingIssueDate(f *field.UnderlyingIssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from DerivativeSecurityList.
func (m Message) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from DerivativeSecurityList.
func (m Message) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from DerivativeSecurityList.
func (m Message) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingFactor() (*field.UnderlyingFactorField, quickfix.MessageRejectError) {
	f := &field.UnderlyingFactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from DerivativeSecurityList.
func (m Message) GetUnderlyingFactor(f *field.UnderlyingFactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCreditRating() (*field.UnderlyingCreditRatingField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from DerivativeSecurityList.
func (m Message) GetUnderlyingCreditRating(f *field.UnderlyingCreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistryField, quickfix.MessageRejectError) {
	f := &field.UnderlyingInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from DerivativeSecurityList.
func (m Message) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from DerivativeSecurityList.
func (m Message) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from DerivativeSecurityList.
func (m Message) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from DerivativeSecurityList.
func (m Message) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from DerivativeSecurityList.
func (m Message) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingStrikePrice() (*field.UnderlyingStrikePriceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from DerivativeSecurityList.
func (m Message) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from DerivativeSecurityList.
func (m Message) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingOptAttribute() (*field.UnderlyingOptAttributeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from DerivativeSecurityList.
func (m Message) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from DerivativeSecurityList.
func (m Message) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCouponRate() (*field.UnderlyingCouponRateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from DerivativeSecurityList.
func (m Message) GetUnderlyingCouponRate(f *field.UnderlyingCouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from DerivativeSecurityList.
func (m Message) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingIssuer() (*field.UnderlyingIssuerField, quickfix.MessageRejectError) {
	f := &field.UnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from DerivativeSecurityList.
func (m Message) GetUnderlyingIssuer(f *field.UnderlyingIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityList.
func (m Message) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from DerivativeSecurityList.
func (m Message) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m Message) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from DerivativeSecurityList.
func (m Message) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDescField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from DerivativeSecurityList.
func (m Message) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityList.
func (m Message) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from DerivativeSecurityList.
func (m Message) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m Message) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from DerivativeSecurityList.
func (m Message) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCPProgram() (*field.UnderlyingCPProgramField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from DerivativeSecurityList.
func (m Message) GetUnderlyingCPProgram(f *field.UnderlyingCPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCPRegType() (*field.UnderlyingCPRegTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from DerivativeSecurityList.
func (m Message) GetUnderlyingCPRegType(f *field.UnderlyingCPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCurrency() (*field.UnderlyingCurrencyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from DerivativeSecurityList.
func (m Message) GetUnderlyingCurrency(f *field.UnderlyingCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingQty() (*field.UnderlyingQtyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from DerivativeSecurityList.
func (m Message) GetUnderlyingQty(f *field.UnderlyingQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingPx() (*field.UnderlyingPxField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from DerivativeSecurityList.
func (m Message) GetUnderlyingPx(f *field.UnderlyingPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPriceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingDirtyPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from DerivativeSecurityList.
func (m Message) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingEndPrice() (*field.UnderlyingEndPriceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingEndPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from DerivativeSecurityList.
func (m Message) GetUnderlyingEndPrice(f *field.UnderlyingEndPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingStartValue() (*field.UnderlyingStartValueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStartValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from DerivativeSecurityList.
func (m Message) GetUnderlyingStartValue(f *field.UnderlyingStartValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingCurrentValue() (*field.UnderlyingCurrentValueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCurrentValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from DerivativeSecurityList.
func (m Message) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityList.
func (m Message) UnderlyingEndValue() (*field.UnderlyingEndValueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingEndValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from DerivativeSecurityList.
func (m Message) GetUnderlyingEndValue(f *field.UnderlyingEndValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityList.
func (m Message) NoUnderlyingStips() (*field.NoUnderlyingStipsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingStipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from DerivativeSecurityList.
func (m Message) GetNoUnderlyingStips(f *field.NoUnderlyingStipsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityList.
func (m Message) TotNoRelatedSym() (*field.TotNoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from DerivativeSecurityList.
func (m Message) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for DerivativeSecurityList.
func (m Message) LastFragment() (*field.LastFragmentField, quickfix.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from DerivativeSecurityList.
func (m Message) GetLastFragment(f *field.LastFragmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for DerivativeSecurityList.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from DerivativeSecurityList.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds DerivativeSecurityList messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for DerivativeSecurityList.
func Builder(
	securityreqid *field.SecurityReqIDField,
	securityresponseid *field.SecurityResponseIDField,
	securityrequestresult *field.SecurityRequestResultField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AA"))
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityrequestresult)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "AA", r
}
