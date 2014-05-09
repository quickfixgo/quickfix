package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//DerivativeSecurityListRequest msg type = z.
type DerivativeSecurityListRequest struct {
	message.Message
}

//DerivativeSecurityListRequestBuilder builds DerivativeSecurityListRequest messages.
type DerivativeSecurityListRequestBuilder struct {
	message.MessageBuilder
}

//CreateDerivativeSecurityListRequestBuilder returns an initialized DerivativeSecurityListRequestBuilder with specified required fields.
func CreateDerivativeSecurityListRequestBuilder(
	securityreqid *field.SecurityReqIDField,
	securitylistrequesttype *field.SecurityListRequestTypeField) DerivativeSecurityListRequestBuilder {
	var builder DerivativeSecurityListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("z"))
	builder.Body.Set(securityreqid)
	builder.Body.Set(securitylistrequesttype)
	return builder
}

//SecurityReqID is a required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListRequestType is a required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecurityListRequestType() (*field.SecurityListRequestTypeField, errors.MessageRejectError) {
	f := &field.SecurityListRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListRequestType reads a SecurityListRequestType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSecurityListRequestType(f *field.SecurityListRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSymbol() (*field.UnderlyingSymbolField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSymbol(f *field.UnderlyingSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfxField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityIDField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityID(f *field.UnderlyingSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoUnderlyingSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingProduct() (*field.UnderlyingProductField, errors.MessageRejectError) {
	f := &field.UnderlyingProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingProduct(f *field.UnderlyingProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCFICode() (*field.UnderlyingCFICodeField, errors.MessageRejectError) {
	f := &field.UnderlyingCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCFICode(f *field.UnderlyingCFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityType(f *field.UnderlyingSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYearField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDateField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDateField, errors.MessageRejectError) {
	f := &field.UnderlyingIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingIssueDate(f *field.UnderlyingIssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingRepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTermField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRateField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFactor() (*field.UnderlyingFactorField, errors.MessageRejectError) {
	f := &field.UnderlyingFactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingFactor(f *field.UnderlyingFactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRatingField, errors.MessageRejectError) {
	f := &field.UnderlyingCreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCreditRating(f *field.UnderlyingCreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistryField, errors.MessageRejectError) {
	f := &field.UnderlyingInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDateField, errors.MessageRejectError) {
	f := &field.UnderlyingRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePriceField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttributeField, errors.MessageRejectError) {
	f := &field.UnderlyingOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplierField, errors.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCouponRate(f *field.UnderlyingCouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchangeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingIssuer() (*field.UnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.UnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingIssuer(f *field.UnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgramField, errors.MessageRejectError) {
	f := &field.UnderlyingCPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCPProgram(f *field.UnderlyingCPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCPRegType(f *field.UnderlyingCPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCurrency() (*field.UnderlyingCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCurrency(f *field.UnderlyingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingQty() (*field.UnderlyingQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingQty(f *field.UnderlyingQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPx() (*field.UnderlyingPxField, errors.MessageRejectError) {
	f := &field.UnderlyingPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPx(f *field.UnderlyingPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingDirtyPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingEndPrice() (*field.UnderlyingEndPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingEndPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingEndPrice(f *field.UnderlyingEndPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStartValue() (*field.UnderlyingStartValueField, errors.MessageRejectError) {
	f := &field.UnderlyingStartValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStartValue(f *field.UnderlyingStartValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrentValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingEndValue() (*field.UnderlyingEndValueField, errors.MessageRejectError) {
	f := &field.UnderlyingEndValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingEndValue(f *field.UnderlyingEndValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUnderlyingStips() (*field.NoUnderlyingStipsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingStipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoUnderlyingStips(f *field.NoUnderlyingStipsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercentField, errors.MessageRejectError) {
	f := &field.UnderlyingAllocationPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlementTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSettlementType(f *field.UnderlyingSettlementTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmountField, errors.MessageRejectError) {
	f := &field.UnderlyingCashAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCashAmount(f *field.UnderlyingCashAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCashType() (*field.UnderlyingCashTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCashTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCashType(f *field.UnderlyingCashTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnitField, errors.MessageRejectError) {
	f := &field.UnderlyingTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCapValue() (*field.UnderlyingCapValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCapValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCapValue(f *field.UnderlyingCapValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoUndlyInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantityField, errors.MessageRejectError) {
	f := &field.UnderlyingAdjustedQuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFXRate() (*field.UnderlyingFXRateField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingFXRate(f *field.UnderlyingFXRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalcField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityTime() (*field.UnderlyingMaturityTimeField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPutOrCall() (*field.UnderlyingPutOrCallField, errors.MessageRejectError) {
	f := &field.UnderlyingPutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyleField, errors.MessageRejectError) {
	f := &field.UnderlyingExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSymbol() (*field.DerivativeSymbolField, errors.MessageRejectError) {
	f := &field.DerivativeSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbol reads a DerivativeSymbol from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSymbol(f *field.DerivativeSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSymbolSfx() (*field.DerivativeSymbolSfxField, errors.MessageRejectError) {
	f := &field.DerivativeSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbolSfx reads a DerivativeSymbolSfx from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSymbolSfx(f *field.DerivativeSymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityID() (*field.DerivativeSecurityIDField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityID reads a DerivativeSecurityID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityID(f *field.DerivativeSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityIDSource() (*field.DerivativeSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityIDSource reads a DerivativeSecurityIDSource from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityIDSource(f *field.DerivativeSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeSecurityAltID() (*field.NoDerivativeSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoDerivativeSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeSecurityAltID reads a NoDerivativeSecurityAltID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoDerivativeSecurityAltID(f *field.NoDerivativeSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProduct is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeProduct() (*field.DerivativeProductField, errors.MessageRejectError) {
	f := &field.DerivativeProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProduct reads a DerivativeProduct from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeProduct(f *field.DerivativeProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeProductComplex() (*field.DerivativeProductComplexField, errors.MessageRejectError) {
	f := &field.DerivativeProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProductComplex reads a DerivativeProductComplex from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeProductComplex(f *field.DerivativeProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivFlexProductEligibilityIndicator() (*field.DerivFlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.DerivFlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivFlexProductEligibilityIndicator reads a DerivFlexProductEligibilityIndicator from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivFlexProductEligibilityIndicator(f *field.DerivFlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityGroup() (*field.DerivativeSecurityGroupField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityGroup reads a DerivativeSecurityGroup from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityGroup(f *field.DerivativeSecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCFICode() (*field.DerivativeCFICodeField, errors.MessageRejectError) {
	f := &field.DerivativeCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCFICode reads a DerivativeCFICode from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeCFICode(f *field.DerivativeCFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityType() (*field.DerivativeSecurityTypeField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityType reads a DerivativeSecurityType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityType(f *field.DerivativeSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecuritySubType() (*field.DerivativeSecuritySubTypeField, errors.MessageRejectError) {
	f := &field.DerivativeSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecuritySubType reads a DerivativeSecuritySubType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecuritySubType(f *field.DerivativeSecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityMonthYear() (*field.DerivativeMaturityMonthYearField, errors.MessageRejectError) {
	f := &field.DerivativeMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityMonthYear reads a DerivativeMaturityMonthYear from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMaturityMonthYear(f *field.DerivativeMaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityDate() (*field.DerivativeMaturityDateField, errors.MessageRejectError) {
	f := &field.DerivativeMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityDate reads a DerivativeMaturityDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMaturityDate(f *field.DerivativeMaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityTime() (*field.DerivativeMaturityTimeField, errors.MessageRejectError) {
	f := &field.DerivativeMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityTime reads a DerivativeMaturityTime from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMaturityTime(f *field.DerivativeMaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSettleOnOpenFlag() (*field.DerivativeSettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.DerivativeSettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettleOnOpenFlag reads a DerivativeSettleOnOpenFlag from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSettleOnOpenFlag(f *field.DerivativeSettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeInstrmtAssignmentMethod() (*field.DerivativeInstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.DerivativeInstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrmtAssignmentMethod reads a DerivativeInstrmtAssignmentMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeInstrmtAssignmentMethod(f *field.DerivativeInstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityStatus() (*field.DerivativeSecurityStatusField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityStatus reads a DerivativeSecurityStatus from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityStatus(f *field.DerivativeSecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeIssueDate() (*field.DerivativeIssueDateField, errors.MessageRejectError) {
	f := &field.DerivativeIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssueDate reads a DerivativeIssueDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeIssueDate(f *field.DerivativeIssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeInstrRegistry() (*field.DerivativeInstrRegistryField, errors.MessageRejectError) {
	f := &field.DerivativeInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrRegistry reads a DerivativeInstrRegistry from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeInstrRegistry(f *field.DerivativeInstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCountryOfIssue() (*field.DerivativeCountryOfIssueField, errors.MessageRejectError) {
	f := &field.DerivativeCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCountryOfIssue reads a DerivativeCountryOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeCountryOfIssue(f *field.DerivativeCountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStateOrProvinceOfIssue() (*field.DerivativeStateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.DerivativeStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStateOrProvinceOfIssue reads a DerivativeStateOrProvinceOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStateOrProvinceOfIssue(f *field.DerivativeStateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikePrice() (*field.DerivativeStrikePriceField, errors.MessageRejectError) {
	f := &field.DerivativeStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikePrice reads a DerivativeStrikePrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikePrice(f *field.DerivativeStrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeLocaleOfIssue() (*field.DerivativeLocaleOfIssueField, errors.MessageRejectError) {
	f := &field.DerivativeLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeLocaleOfIssue reads a DerivativeLocaleOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeLocaleOfIssue(f *field.DerivativeLocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeCurrency() (*field.DerivativeStrikeCurrencyField, errors.MessageRejectError) {
	f := &field.DerivativeStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeCurrency reads a DerivativeStrikeCurrency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikeCurrency(f *field.DerivativeStrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeMultiplier() (*field.DerivativeStrikeMultiplierField, errors.MessageRejectError) {
	f := &field.DerivativeStrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeMultiplier reads a DerivativeStrikeMultiplier from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikeMultiplier(f *field.DerivativeStrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeValue() (*field.DerivativeStrikeValueField, errors.MessageRejectError) {
	f := &field.DerivativeStrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeValue reads a DerivativeStrikeValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikeValue(f *field.DerivativeStrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeOptAttribute() (*field.DerivativeOptAttributeField, errors.MessageRejectError) {
	f := &field.DerivativeOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptAttribute reads a DerivativeOptAttribute from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeOptAttribute(f *field.DerivativeOptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeContractMultiplier() (*field.DerivativeContractMultiplierField, errors.MessageRejectError) {
	f := &field.DerivativeContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractMultiplier reads a DerivativeContractMultiplier from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeContractMultiplier(f *field.DerivativeContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMinPriceIncrement() (*field.DerivativeMinPriceIncrementField, errors.MessageRejectError) {
	f := &field.DerivativeMinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrement reads a DerivativeMinPriceIncrement from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrement(f *field.DerivativeMinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMinPriceIncrementAmount() (*field.DerivativeMinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.DerivativeMinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrementAmount reads a DerivativeMinPriceIncrementAmount from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrementAmount(f *field.DerivativeMinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeUnitOfMeasure() (*field.DerivativeUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.DerivativeUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasure reads a DerivativeUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasure(f *field.DerivativeUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeUnitOfMeasureQty() (*field.DerivativeUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.DerivativeUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasureQty reads a DerivativeUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasureQty(f *field.DerivativeUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceUnitOfMeasure() (*field.DerivativePriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.DerivativePriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasure reads a DerivativePriceUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasure(f *field.DerivativePriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceUnitOfMeasureQty() (*field.DerivativePriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.DerivativePriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasureQty reads a DerivativePriceUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasureQty(f *field.DerivativePriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeExerciseStyle() (*field.DerivativeExerciseStyleField, errors.MessageRejectError) {
	f := &field.DerivativeExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeExerciseStyle reads a DerivativeExerciseStyle from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeExerciseStyle(f *field.DerivativeExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeOptPayAmount() (*field.DerivativeOptPayAmountField, errors.MessageRejectError) {
	f := &field.DerivativeOptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptPayAmount reads a DerivativeOptPayAmount from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeOptPayAmount(f *field.DerivativeOptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeTimeUnit() (*field.DerivativeTimeUnitField, errors.MessageRejectError) {
	f := &field.DerivativeTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeTimeUnit reads a DerivativeTimeUnit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeTimeUnit(f *field.DerivativeTimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityExchange() (*field.DerivativeSecurityExchangeField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityExchange reads a DerivativeSecurityExchange from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityExchange(f *field.DerivativeSecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePositionLimit() (*field.DerivativePositionLimitField, errors.MessageRejectError) {
	f := &field.DerivativePositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePositionLimit reads a DerivativePositionLimit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePositionLimit(f *field.DerivativePositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeNTPositionLimit() (*field.DerivativeNTPositionLimitField, errors.MessageRejectError) {
	f := &field.DerivativeNTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeNTPositionLimit reads a DerivativeNTPositionLimit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeNTPositionLimit(f *field.DerivativeNTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeIssuer() (*field.DerivativeIssuerField, errors.MessageRejectError) {
	f := &field.DerivativeIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssuer reads a DerivativeIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeIssuer(f *field.DerivativeIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedIssuerLen() (*field.DerivativeEncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuerLen reads a DerivativeEncodedIssuerLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuerLen(f *field.DerivativeEncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedIssuer() (*field.DerivativeEncodedIssuerField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuer reads a DerivativeEncodedIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuer(f *field.DerivativeEncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityDesc() (*field.DerivativeSecurityDescField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityDesc reads a DerivativeSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityDesc(f *field.DerivativeSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedSecurityDescLen() (*field.DerivativeEncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDescLen reads a DerivativeEncodedSecurityDescLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDescLen(f *field.DerivativeEncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedSecurityDesc() (*field.DerivativeEncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDesc reads a DerivativeEncodedSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDesc(f *field.DerivativeEncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeContractSettlMonth() (*field.DerivativeContractSettlMonthField, errors.MessageRejectError) {
	f := &field.DerivativeContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractSettlMonth reads a DerivativeContractSettlMonth from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeContractSettlMonth(f *field.DerivativeContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeEvents() (*field.NoDerivativeEventsField, errors.MessageRejectError) {
	f := &field.NoDerivativeEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeEvents reads a NoDerivativeEvents from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoDerivativeEvents(f *field.NoDerivativeEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeInstrumentParties() (*field.NoDerivativeInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoDerivativeInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrumentParties reads a NoDerivativeInstrumentParties from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoDerivativeInstrumentParties(f *field.NoDerivativeInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSettlMethod() (*field.DerivativeSettlMethodField, errors.MessageRejectError) {
	f := &field.DerivativeSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettlMethod reads a DerivativeSettlMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSettlMethod(f *field.DerivativeSettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceQuoteMethod() (*field.DerivativePriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.DerivativePriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceQuoteMethod reads a DerivativePriceQuoteMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePriceQuoteMethod(f *field.DerivativePriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeFuturesValuationMethod() (*field.DerivativeFuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.DerivativeFuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFuturesValuationMethod reads a DerivativeFuturesValuationMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeFuturesValuationMethod(f *field.DerivativeFuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeListMethod() (*field.DerivativeListMethodField, errors.MessageRejectError) {
	f := &field.DerivativeListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeListMethod reads a DerivativeListMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeListMethod(f *field.DerivativeListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCapPrice() (*field.DerivativeCapPriceField, errors.MessageRejectError) {
	f := &field.DerivativeCapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCapPrice reads a DerivativeCapPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeCapPrice(f *field.DerivativeCapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeFloorPrice() (*field.DerivativeFloorPriceField, errors.MessageRejectError) {
	f := &field.DerivativeFloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFloorPrice reads a DerivativeFloorPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeFloorPrice(f *field.DerivativeFloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePutOrCall() (*field.DerivativePutOrCallField, errors.MessageRejectError) {
	f := &field.DerivativePutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePutOrCall reads a DerivativePutOrCall from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePutOrCall(f *field.DerivativePutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXMLLen() (*field.DerivativeSecurityXMLLenField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLLen reads a DerivativeSecurityXMLLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLLen(f *field.DerivativeSecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXML() (*field.DerivativeSecurityXMLField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXML reads a DerivativeSecurityXML from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXML(f *field.DerivativeSecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXMLSchema() (*field.DerivativeSecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLSchema reads a DerivativeSecurityXMLSchema from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLSchema(f *field.DerivativeSecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}
