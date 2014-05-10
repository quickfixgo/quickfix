package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BR"))
	return builder
}

//SecurityReqID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResultField, errors.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityRequestResult(f *field.SecurityRequestResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityUpdateAction is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateActionField, errors.MessageRejectError) {
	f := &field.SecurityUpdateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityUpdateAction reads a SecurityUpdateAction from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityUpdateAction(f *field.SecurityUpdateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbol() (*field.UnderlyingSymbolField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSymbol(f *field.UnderlyingSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfxField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityID() (*field.UnderlyingSecurityIDField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityID(f *field.UnderlyingSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoUnderlyingSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingProduct() (*field.UnderlyingProductField, errors.MessageRejectError) {
	f := &field.UnderlyingProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingProduct(f *field.UnderlyingProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCFICode() (*field.UnderlyingCFICodeField, errors.MessageRejectError) {
	f := &field.UnderlyingCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCFICode(f *field.UnderlyingCFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityType() (*field.UnderlyingSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityType(f *field.UnderlyingSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYearField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityDate() (*field.UnderlyingMaturityDateField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssueDate() (*field.UnderlyingIssueDateField, errors.MessageRejectError) {
	f := &field.UnderlyingIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingIssueDate(f *field.UnderlyingIssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingRepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTermField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRateField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFactor() (*field.UnderlyingFactorField, errors.MessageRejectError) {
	f := &field.UnderlyingFactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFactor(f *field.UnderlyingFactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCreditRating() (*field.UnderlyingCreditRatingField, errors.MessageRejectError) {
	f := &field.UnderlyingCreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCreditRating(f *field.UnderlyingCreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistryField, errors.MessageRejectError) {
	f := &field.UnderlyingInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDateField, errors.MessageRejectError) {
	f := &field.UnderlyingRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikePrice() (*field.UnderlyingStrikePriceField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingOptAttribute() (*field.UnderlyingOptAttributeField, errors.MessageRejectError) {
	f := &field.UnderlyingOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplierField, errors.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponRate() (*field.UnderlyingCouponRateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCouponRate(f *field.UnderlyingCouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchangeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssuer() (*field.UnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.UnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingIssuer(f *field.UnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPProgram() (*field.UnderlyingCPProgramField, errors.MessageRejectError) {
	f := &field.UnderlyingCPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCPProgram(f *field.UnderlyingCPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPRegType() (*field.UnderlyingCPRegTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCPRegType(f *field.UnderlyingCPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrency() (*field.UnderlyingCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCurrency(f *field.UnderlyingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingQty() (*field.UnderlyingQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingQty(f *field.UnderlyingQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPx() (*field.UnderlyingPxField, errors.MessageRejectError) {
	f := &field.UnderlyingPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPx(f *field.UnderlyingPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingDirtyPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndPrice() (*field.UnderlyingEndPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingEndPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingEndPrice(f *field.UnderlyingEndPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStartValue() (*field.UnderlyingStartValueField, errors.MessageRejectError) {
	f := &field.UnderlyingStartValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStartValue(f *field.UnderlyingStartValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrentValue() (*field.UnderlyingCurrentValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrentValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndValue() (*field.UnderlyingEndValueField, errors.MessageRejectError) {
	f := &field.UnderlyingEndValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingEndValue(f *field.UnderlyingEndValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingStips() (*field.NoUnderlyingStipsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingStipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoUnderlyingStips(f *field.NoUnderlyingStipsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercentField, errors.MessageRejectError) {
	f := &field.UnderlyingAllocationPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlementType() (*field.UnderlyingSettlementTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlementTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSettlementType(f *field.UnderlyingSettlementTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashAmount() (*field.UnderlyingCashAmountField, errors.MessageRejectError) {
	f := &field.UnderlyingCashAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCashAmount(f *field.UnderlyingCashAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashType() (*field.UnderlyingCashTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCashTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCashType(f *field.UnderlyingCashTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingTimeUnit() (*field.UnderlyingTimeUnitField, errors.MessageRejectError) {
	f := &field.UnderlyingTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCapValue() (*field.UnderlyingCapValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCapValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCapValue(f *field.UnderlyingCapValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoUndlyInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlMethod() (*field.UnderlyingSettlMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantityField, errors.MessageRejectError) {
	f := &field.UnderlyingAdjustedQuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRate() (*field.UnderlyingFXRateField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFXRate(f *field.UnderlyingFXRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalcField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityTime() (*field.UnderlyingMaturityTimeField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPutOrCall() (*field.UnderlyingPutOrCallField, errors.MessageRejectError) {
	f := &field.UnderlyingPutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyleField, errors.MessageRejectError) {
	f := &field.UnderlyingExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplierUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingContractMultiplierUnit() (*field.UnderlyingContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplierUnit reads a UnderlyingContractMultiplierUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingContractMultiplierUnit(f *field.UnderlyingContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFlowScheduleType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFlowScheduleType() (*field.UnderlyingFlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingFlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFlowScheduleType reads a UnderlyingFlowScheduleType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFlowScheduleType(f *field.UnderlyingFlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRestructuringType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRestructuringType() (*field.UnderlyingRestructuringTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingRestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRestructuringType reads a UnderlyingRestructuringType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRestructuringType(f *field.UnderlyingRestructuringTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSeniority is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSeniority() (*field.UnderlyingSeniorityField, errors.MessageRejectError) {
	f := &field.UnderlyingSeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSeniority reads a UnderlyingSeniority from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSeniority(f *field.UnderlyingSeniorityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingNotionalPercentageOutstanding() (*field.UnderlyingNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.UnderlyingNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingNotionalPercentageOutstanding reads a UnderlyingNotionalPercentageOutstanding from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingNotionalPercentageOutstanding(f *field.UnderlyingNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingOriginalNotionalPercentageOutstanding() (*field.UnderlyingOriginalNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.UnderlyingOriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOriginalNotionalPercentageOutstanding reads a UnderlyingOriginalNotionalPercentageOutstanding from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingOriginalNotionalPercentageOutstanding(f *field.UnderlyingOriginalNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAttachmentPoint is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAttachmentPoint() (*field.UnderlyingAttachmentPointField, errors.MessageRejectError) {
	f := &field.UnderlyingAttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAttachmentPoint reads a UnderlyingAttachmentPoint from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingAttachmentPoint(f *field.UnderlyingAttachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDetachmentPoint is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingDetachmentPoint() (*field.UnderlyingDetachmentPointField, errors.MessageRejectError) {
	f := &field.UnderlyingDetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDetachmentPoint reads a UnderlyingDetachmentPoint from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingDetachmentPoint(f *field.UnderlyingDetachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbol() (*field.DerivativeSymbolField, errors.MessageRejectError) {
	f := &field.DerivativeSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbol reads a DerivativeSymbol from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSymbol(f *field.DerivativeSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbolSfx() (*field.DerivativeSymbolSfxField, errors.MessageRejectError) {
	f := &field.DerivativeSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbolSfx reads a DerivativeSymbolSfx from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSymbolSfx(f *field.DerivativeSymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityID() (*field.DerivativeSecurityIDField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityID reads a DerivativeSecurityID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityID(f *field.DerivativeSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityIDSource() (*field.DerivativeSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityIDSource reads a DerivativeSecurityIDSource from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityIDSource(f *field.DerivativeSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeSecurityAltID() (*field.NoDerivativeSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoDerivativeSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeSecurityAltID reads a NoDerivativeSecurityAltID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeSecurityAltID(f *field.NoDerivativeSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProduct() (*field.DerivativeProductField, errors.MessageRejectError) {
	f := &field.DerivativeProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProduct reads a DerivativeProduct from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeProduct(f *field.DerivativeProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProductComplex() (*field.DerivativeProductComplexField, errors.MessageRejectError) {
	f := &field.DerivativeProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProductComplex reads a DerivativeProductComplex from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeProductComplex(f *field.DerivativeProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivFlexProductEligibilityIndicator() (*field.DerivFlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.DerivFlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivFlexProductEligibilityIndicator reads a DerivFlexProductEligibilityIndicator from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivFlexProductEligibilityIndicator(f *field.DerivFlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityGroup() (*field.DerivativeSecurityGroupField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityGroup reads a DerivativeSecurityGroup from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityGroup(f *field.DerivativeSecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCFICode() (*field.DerivativeCFICodeField, errors.MessageRejectError) {
	f := &field.DerivativeCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCFICode reads a DerivativeCFICode from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeCFICode(f *field.DerivativeCFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityType() (*field.DerivativeSecurityTypeField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityType reads a DerivativeSecurityType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityType(f *field.DerivativeSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecuritySubType() (*field.DerivativeSecuritySubTypeField, errors.MessageRejectError) {
	f := &field.DerivativeSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecuritySubType reads a DerivativeSecuritySubType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecuritySubType(f *field.DerivativeSecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityMonthYear() (*field.DerivativeMaturityMonthYearField, errors.MessageRejectError) {
	f := &field.DerivativeMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityMonthYear reads a DerivativeMaturityMonthYear from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMaturityMonthYear(f *field.DerivativeMaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityDate() (*field.DerivativeMaturityDateField, errors.MessageRejectError) {
	f := &field.DerivativeMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityDate reads a DerivativeMaturityDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMaturityDate(f *field.DerivativeMaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityTime() (*field.DerivativeMaturityTimeField, errors.MessageRejectError) {
	f := &field.DerivativeMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityTime reads a DerivativeMaturityTime from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMaturityTime(f *field.DerivativeMaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettleOnOpenFlag() (*field.DerivativeSettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.DerivativeSettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettleOnOpenFlag reads a DerivativeSettleOnOpenFlag from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSettleOnOpenFlag(f *field.DerivativeSettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrmtAssignmentMethod() (*field.DerivativeInstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.DerivativeInstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrmtAssignmentMethod reads a DerivativeInstrmtAssignmentMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeInstrmtAssignmentMethod(f *field.DerivativeInstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityStatus() (*field.DerivativeSecurityStatusField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityStatus reads a DerivativeSecurityStatus from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityStatus(f *field.DerivativeSecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssueDate() (*field.DerivativeIssueDateField, errors.MessageRejectError) {
	f := &field.DerivativeIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssueDate reads a DerivativeIssueDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeIssueDate(f *field.DerivativeIssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrRegistry() (*field.DerivativeInstrRegistryField, errors.MessageRejectError) {
	f := &field.DerivativeInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrRegistry reads a DerivativeInstrRegistry from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeInstrRegistry(f *field.DerivativeInstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCountryOfIssue() (*field.DerivativeCountryOfIssueField, errors.MessageRejectError) {
	f := &field.DerivativeCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCountryOfIssue reads a DerivativeCountryOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeCountryOfIssue(f *field.DerivativeCountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStateOrProvinceOfIssue() (*field.DerivativeStateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.DerivativeStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStateOrProvinceOfIssue reads a DerivativeStateOrProvinceOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStateOrProvinceOfIssue(f *field.DerivativeStateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikePrice() (*field.DerivativeStrikePriceField, errors.MessageRejectError) {
	f := &field.DerivativeStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikePrice reads a DerivativeStrikePrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikePrice(f *field.DerivativeStrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeLocaleOfIssue() (*field.DerivativeLocaleOfIssueField, errors.MessageRejectError) {
	f := &field.DerivativeLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeLocaleOfIssue reads a DerivativeLocaleOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeLocaleOfIssue(f *field.DerivativeLocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeCurrency() (*field.DerivativeStrikeCurrencyField, errors.MessageRejectError) {
	f := &field.DerivativeStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeCurrency reads a DerivativeStrikeCurrency from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikeCurrency(f *field.DerivativeStrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeMultiplier() (*field.DerivativeStrikeMultiplierField, errors.MessageRejectError) {
	f := &field.DerivativeStrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeMultiplier reads a DerivativeStrikeMultiplier from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikeMultiplier(f *field.DerivativeStrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeValue() (*field.DerivativeStrikeValueField, errors.MessageRejectError) {
	f := &field.DerivativeStrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeValue reads a DerivativeStrikeValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikeValue(f *field.DerivativeStrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptAttribute() (*field.DerivativeOptAttributeField, errors.MessageRejectError) {
	f := &field.DerivativeOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptAttribute reads a DerivativeOptAttribute from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeOptAttribute(f *field.DerivativeOptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractMultiplier() (*field.DerivativeContractMultiplierField, errors.MessageRejectError) {
	f := &field.DerivativeContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractMultiplier reads a DerivativeContractMultiplier from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeContractMultiplier(f *field.DerivativeContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrement() (*field.DerivativeMinPriceIncrementField, errors.MessageRejectError) {
	f := &field.DerivativeMinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrement reads a DerivativeMinPriceIncrement from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMinPriceIncrement(f *field.DerivativeMinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrementAmount() (*field.DerivativeMinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.DerivativeMinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrementAmount reads a DerivativeMinPriceIncrementAmount from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMinPriceIncrementAmount(f *field.DerivativeMinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasure() (*field.DerivativeUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.DerivativeUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasure reads a DerivativeUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeUnitOfMeasure(f *field.DerivativeUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasureQty() (*field.DerivativeUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.DerivativeUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasureQty reads a DerivativeUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeUnitOfMeasureQty(f *field.DerivativeUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasure() (*field.DerivativePriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.DerivativePriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasure reads a DerivativePriceUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePriceUnitOfMeasure(f *field.DerivativePriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasureQty() (*field.DerivativePriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.DerivativePriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasureQty reads a DerivativePriceUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePriceUnitOfMeasureQty(f *field.DerivativePriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeExerciseStyle() (*field.DerivativeExerciseStyleField, errors.MessageRejectError) {
	f := &field.DerivativeExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeExerciseStyle reads a DerivativeExerciseStyle from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeExerciseStyle(f *field.DerivativeExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptPayAmount() (*field.DerivativeOptPayAmountField, errors.MessageRejectError) {
	f := &field.DerivativeOptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptPayAmount reads a DerivativeOptPayAmount from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeOptPayAmount(f *field.DerivativeOptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeTimeUnit() (*field.DerivativeTimeUnitField, errors.MessageRejectError) {
	f := &field.DerivativeTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeTimeUnit reads a DerivativeTimeUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeTimeUnit(f *field.DerivativeTimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityExchange() (*field.DerivativeSecurityExchangeField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityExchange reads a DerivativeSecurityExchange from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityExchange(f *field.DerivativeSecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePositionLimit() (*field.DerivativePositionLimitField, errors.MessageRejectError) {
	f := &field.DerivativePositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePositionLimit reads a DerivativePositionLimit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePositionLimit(f *field.DerivativePositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeNTPositionLimit() (*field.DerivativeNTPositionLimitField, errors.MessageRejectError) {
	f := &field.DerivativeNTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeNTPositionLimit reads a DerivativeNTPositionLimit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeNTPositionLimit(f *field.DerivativeNTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssuer() (*field.DerivativeIssuerField, errors.MessageRejectError) {
	f := &field.DerivativeIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssuer reads a DerivativeIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeIssuer(f *field.DerivativeIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuerLen() (*field.DerivativeEncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuerLen reads a DerivativeEncodedIssuerLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedIssuerLen(f *field.DerivativeEncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuer() (*field.DerivativeEncodedIssuerField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuer reads a DerivativeEncodedIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedIssuer(f *field.DerivativeEncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityDesc() (*field.DerivativeSecurityDescField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityDesc reads a DerivativeSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityDesc(f *field.DerivativeSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDescLen() (*field.DerivativeEncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDescLen reads a DerivativeEncodedSecurityDescLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedSecurityDescLen(f *field.DerivativeEncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDesc() (*field.DerivativeEncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.DerivativeEncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDesc reads a DerivativeEncodedSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedSecurityDesc(f *field.DerivativeEncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractSettlMonth() (*field.DerivativeContractSettlMonthField, errors.MessageRejectError) {
	f := &field.DerivativeContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractSettlMonth reads a DerivativeContractSettlMonth from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeContractSettlMonth(f *field.DerivativeContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeEvents() (*field.NoDerivativeEventsField, errors.MessageRejectError) {
	f := &field.NoDerivativeEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeEvents reads a NoDerivativeEvents from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeEvents(f *field.NoDerivativeEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrumentParties() (*field.NoDerivativeInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoDerivativeInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrumentParties reads a NoDerivativeInstrumentParties from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeInstrumentParties(f *field.NoDerivativeInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettlMethod() (*field.DerivativeSettlMethodField, errors.MessageRejectError) {
	f := &field.DerivativeSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettlMethod reads a DerivativeSettlMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSettlMethod(f *field.DerivativeSettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceQuoteMethod() (*field.DerivativePriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.DerivativePriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceQuoteMethod reads a DerivativePriceQuoteMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePriceQuoteMethod(f *field.DerivativePriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeValuationMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeValuationMethod() (*field.DerivativeValuationMethodField, errors.MessageRejectError) {
	f := &field.DerivativeValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeValuationMethod reads a DerivativeValuationMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeValuationMethod(f *field.DerivativeValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeListMethod() (*field.DerivativeListMethodField, errors.MessageRejectError) {
	f := &field.DerivativeListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeListMethod reads a DerivativeListMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeListMethod(f *field.DerivativeListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCapPrice() (*field.DerivativeCapPriceField, errors.MessageRejectError) {
	f := &field.DerivativeCapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCapPrice reads a DerivativeCapPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeCapPrice(f *field.DerivativeCapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFloorPrice() (*field.DerivativeFloorPriceField, errors.MessageRejectError) {
	f := &field.DerivativeFloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFloorPrice reads a DerivativeFloorPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeFloorPrice(f *field.DerivativeFloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePutOrCall() (*field.DerivativePutOrCallField, errors.MessageRejectError) {
	f := &field.DerivativePutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePutOrCall reads a DerivativePutOrCall from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePutOrCall(f *field.DerivativePutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLLen() (*field.DerivativeSecurityXMLLenField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLLen reads a DerivativeSecurityXMLLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityXMLLen(f *field.DerivativeSecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXML() (*field.DerivativeSecurityXMLField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXML reads a DerivativeSecurityXML from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityXML(f *field.DerivativeSecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLSchema() (*field.DerivativeSecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.DerivativeSecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLSchema reads a DerivativeSecurityXMLSchema from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityXMLSchema(f *field.DerivativeSecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractMultiplierUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractMultiplierUnit() (*field.DerivativeContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.DerivativeContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractMultiplierUnit reads a DerivativeContractMultiplierUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeContractMultiplierUnit(f *field.DerivativeContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFlowScheduleType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFlowScheduleType() (*field.DerivativeFlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.DerivativeFlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFlowScheduleType reads a DerivativeFlowScheduleType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeFlowScheduleType(f *field.DerivativeFlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrAttrib() (*field.NoDerivativeInstrAttribField, errors.MessageRejectError) {
	f := &field.NoDerivativeInstrAttribField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrAttrib reads a NoDerivativeInstrAttrib from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeInstrAttrib(f *field.NoDerivativeInstrAttribField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMarketSegments is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoMarketSegments() (*field.NoMarketSegmentsField, errors.MessageRejectError) {
	f := &field.NoMarketSegmentsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMarketSegments reads a NoMarketSegments from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoMarketSegments(f *field.NoMarketSegmentsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSymField, errors.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}
