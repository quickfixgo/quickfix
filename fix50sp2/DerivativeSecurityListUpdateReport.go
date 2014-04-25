package fix50sp2

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
	builder.Header.Set(field.BuildMsgType("BR"))
	return builder
}

//SecurityReqID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResult, errors.MessageRejectError) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityRequestResult(f *field.SecurityRequestResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityUpdateAction is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, errors.MessageRejectError) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityUpdateAction reads a SecurityUpdateAction from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetSecurityUpdateAction(f *field.SecurityUpdateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbol() (*field.UnderlyingSymbol, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSymbol(f *field.UnderlyingSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityID() (*field.UnderlyingSecurityID, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityID(f *field.UnderlyingSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingProduct() (*field.UnderlyingProduct, errors.MessageRejectError) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingProduct(f *field.UnderlyingProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCFICode() (*field.UnderlyingCFICode, errors.MessageRejectError) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCFICode(f *field.UnderlyingCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityType() (*field.UnderlyingSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityType(f *field.UnderlyingSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssueDate() (*field.UnderlyingIssueDate, errors.MessageRejectError) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingIssueDate(f *field.UnderlyingIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFactor() (*field.UnderlyingFactor, errors.MessageRejectError) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFactor(f *field.UnderlyingFactor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCreditRating() (*field.UnderlyingCreditRating, errors.MessageRejectError) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCreditRating(f *field.UnderlyingCreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, errors.MessageRejectError) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, errors.MessageRejectError) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, errors.MessageRejectError) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, errors.MessageRejectError) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCouponRate() (*field.UnderlyingCouponRate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCouponRate(f *field.UnderlyingCouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingIssuer() (*field.UnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingIssuer(f *field.UnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPProgram() (*field.UnderlyingCPProgram, errors.MessageRejectError) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCPProgram(f *field.UnderlyingCPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCPRegType() (*field.UnderlyingCPRegType, errors.MessageRejectError) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCPRegType(f *field.UnderlyingCPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrency() (*field.UnderlyingCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCurrency(f *field.UnderlyingCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingQty() (*field.UnderlyingQty, errors.MessageRejectError) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingQty(f *field.UnderlyingQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPx() (*field.UnderlyingPx, errors.MessageRejectError) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPx(f *field.UnderlyingPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndPrice() (*field.UnderlyingEndPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingEndPrice(f *field.UnderlyingEndPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingStartValue() (*field.UnderlyingStartValue, errors.MessageRejectError) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingStartValue(f *field.UnderlyingStartValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingEndValue() (*field.UnderlyingEndValue, errors.MessageRejectError) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingEndValue(f *field.UnderlyingEndValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUnderlyingStips() (*field.NoUnderlyingStips, errors.MessageRejectError) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoUnderlyingStips(f *field.NoUnderlyingStips) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, errors.MessageRejectError) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlementType() (*field.UnderlyingSettlementType, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSettlementType(f *field.UnderlyingSettlementType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashAmount() (*field.UnderlyingCashAmount, errors.MessageRejectError) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCashAmount(f *field.UnderlyingCashAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCashType() (*field.UnderlyingCashType, errors.MessageRejectError) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCashType(f *field.UnderlyingCashType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, errors.MessageRejectError) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingCapValue() (*field.UnderlyingCapValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingCapValue(f *field.UnderlyingCapValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, errors.MessageRejectError) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRate() (*field.UnderlyingFXRate, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFXRate(f *field.UnderlyingFXRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, errors.MessageRejectError) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, errors.MessageRejectError) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplierUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingContractMultiplierUnit() (*field.UnderlyingContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.UnderlyingContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplierUnit reads a UnderlyingContractMultiplierUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingContractMultiplierUnit(f *field.UnderlyingContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFlowScheduleType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingFlowScheduleType() (*field.UnderlyingFlowScheduleType, errors.MessageRejectError) {
	f := new(field.UnderlyingFlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFlowScheduleType reads a UnderlyingFlowScheduleType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingFlowScheduleType(f *field.UnderlyingFlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRestructuringType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingRestructuringType() (*field.UnderlyingRestructuringType, errors.MessageRejectError) {
	f := new(field.UnderlyingRestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRestructuringType reads a UnderlyingRestructuringType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingRestructuringType(f *field.UnderlyingRestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSeniority is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingSeniority() (*field.UnderlyingSeniority, errors.MessageRejectError) {
	f := new(field.UnderlyingSeniority)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSeniority reads a UnderlyingSeniority from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingSeniority(f *field.UnderlyingSeniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingNotionalPercentageOutstanding() (*field.UnderlyingNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.UnderlyingNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingNotionalPercentageOutstanding reads a UnderlyingNotionalPercentageOutstanding from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingNotionalPercentageOutstanding(f *field.UnderlyingNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingOriginalNotionalPercentageOutstanding() (*field.UnderlyingOriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.UnderlyingOriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOriginalNotionalPercentageOutstanding reads a UnderlyingOriginalNotionalPercentageOutstanding from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingOriginalNotionalPercentageOutstanding(f *field.UnderlyingOriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAttachmentPoint is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingAttachmentPoint() (*field.UnderlyingAttachmentPoint, errors.MessageRejectError) {
	f := new(field.UnderlyingAttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAttachmentPoint reads a UnderlyingAttachmentPoint from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingAttachmentPoint(f *field.UnderlyingAttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDetachmentPoint is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) UnderlyingDetachmentPoint() (*field.UnderlyingDetachmentPoint, errors.MessageRejectError) {
	f := new(field.UnderlyingDetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDetachmentPoint reads a UnderlyingDetachmentPoint from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetUnderlyingDetachmentPoint(f *field.UnderlyingDetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbol() (*field.DerivativeSymbol, errors.MessageRejectError) {
	f := new(field.DerivativeSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbol reads a DerivativeSymbol from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSymbol(f *field.DerivativeSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSymbolSfx() (*field.DerivativeSymbolSfx, errors.MessageRejectError) {
	f := new(field.DerivativeSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbolSfx reads a DerivativeSymbolSfx from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSymbolSfx(f *field.DerivativeSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityID() (*field.DerivativeSecurityID, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityID reads a DerivativeSecurityID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityID(f *field.DerivativeSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityIDSource() (*field.DerivativeSecurityIDSource, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityIDSource reads a DerivativeSecurityIDSource from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityIDSource(f *field.DerivativeSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeSecurityAltID() (*field.NoDerivativeSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoDerivativeSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeSecurityAltID reads a NoDerivativeSecurityAltID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeSecurityAltID(f *field.NoDerivativeSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProduct is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProduct() (*field.DerivativeProduct, errors.MessageRejectError) {
	f := new(field.DerivativeProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProduct reads a DerivativeProduct from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeProduct(f *field.DerivativeProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeProductComplex() (*field.DerivativeProductComplex, errors.MessageRejectError) {
	f := new(field.DerivativeProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProductComplex reads a DerivativeProductComplex from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeProductComplex(f *field.DerivativeProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivFlexProductEligibilityIndicator() (*field.DerivFlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.DerivFlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivFlexProductEligibilityIndicator reads a DerivFlexProductEligibilityIndicator from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivFlexProductEligibilityIndicator(f *field.DerivFlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityGroup() (*field.DerivativeSecurityGroup, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityGroup reads a DerivativeSecurityGroup from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityGroup(f *field.DerivativeSecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCFICode() (*field.DerivativeCFICode, errors.MessageRejectError) {
	f := new(field.DerivativeCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCFICode reads a DerivativeCFICode from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeCFICode(f *field.DerivativeCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityType() (*field.DerivativeSecurityType, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityType reads a DerivativeSecurityType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityType(f *field.DerivativeSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecuritySubType() (*field.DerivativeSecuritySubType, errors.MessageRejectError) {
	f := new(field.DerivativeSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecuritySubType reads a DerivativeSecuritySubType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecuritySubType(f *field.DerivativeSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityMonthYear() (*field.DerivativeMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityMonthYear reads a DerivativeMaturityMonthYear from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMaturityMonthYear(f *field.DerivativeMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityDate() (*field.DerivativeMaturityDate, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityDate reads a DerivativeMaturityDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMaturityDate(f *field.DerivativeMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMaturityTime() (*field.DerivativeMaturityTime, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityTime reads a DerivativeMaturityTime from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMaturityTime(f *field.DerivativeMaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettleOnOpenFlag() (*field.DerivativeSettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.DerivativeSettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettleOnOpenFlag reads a DerivativeSettleOnOpenFlag from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSettleOnOpenFlag(f *field.DerivativeSettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrmtAssignmentMethod() (*field.DerivativeInstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.DerivativeInstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrmtAssignmentMethod reads a DerivativeInstrmtAssignmentMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeInstrmtAssignmentMethod(f *field.DerivativeInstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityStatus() (*field.DerivativeSecurityStatus, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityStatus reads a DerivativeSecurityStatus from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityStatus(f *field.DerivativeSecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssueDate() (*field.DerivativeIssueDate, errors.MessageRejectError) {
	f := new(field.DerivativeIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssueDate reads a DerivativeIssueDate from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeIssueDate(f *field.DerivativeIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeInstrRegistry() (*field.DerivativeInstrRegistry, errors.MessageRejectError) {
	f := new(field.DerivativeInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrRegistry reads a DerivativeInstrRegistry from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeInstrRegistry(f *field.DerivativeInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCountryOfIssue() (*field.DerivativeCountryOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCountryOfIssue reads a DerivativeCountryOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeCountryOfIssue(f *field.DerivativeCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStateOrProvinceOfIssue() (*field.DerivativeStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStateOrProvinceOfIssue reads a DerivativeStateOrProvinceOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStateOrProvinceOfIssue(f *field.DerivativeStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikePrice() (*field.DerivativeStrikePrice, errors.MessageRejectError) {
	f := new(field.DerivativeStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikePrice reads a DerivativeStrikePrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikePrice(f *field.DerivativeStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeLocaleOfIssue() (*field.DerivativeLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeLocaleOfIssue reads a DerivativeLocaleOfIssue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeLocaleOfIssue(f *field.DerivativeLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeCurrency() (*field.DerivativeStrikeCurrency, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeCurrency reads a DerivativeStrikeCurrency from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikeCurrency(f *field.DerivativeStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeMultiplier() (*field.DerivativeStrikeMultiplier, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeMultiplier reads a DerivativeStrikeMultiplier from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikeMultiplier(f *field.DerivativeStrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeStrikeValue() (*field.DerivativeStrikeValue, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeValue reads a DerivativeStrikeValue from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeStrikeValue(f *field.DerivativeStrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptAttribute() (*field.DerivativeOptAttribute, errors.MessageRejectError) {
	f := new(field.DerivativeOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptAttribute reads a DerivativeOptAttribute from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeOptAttribute(f *field.DerivativeOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractMultiplier() (*field.DerivativeContractMultiplier, errors.MessageRejectError) {
	f := new(field.DerivativeContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractMultiplier reads a DerivativeContractMultiplier from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeContractMultiplier(f *field.DerivativeContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrement() (*field.DerivativeMinPriceIncrement, errors.MessageRejectError) {
	f := new(field.DerivativeMinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrement reads a DerivativeMinPriceIncrement from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMinPriceIncrement(f *field.DerivativeMinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeMinPriceIncrementAmount() (*field.DerivativeMinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.DerivativeMinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrementAmount reads a DerivativeMinPriceIncrementAmount from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeMinPriceIncrementAmount(f *field.DerivativeMinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasure() (*field.DerivativeUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.DerivativeUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasure reads a DerivativeUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeUnitOfMeasure(f *field.DerivativeUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeUnitOfMeasureQty() (*field.DerivativeUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.DerivativeUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasureQty reads a DerivativeUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeUnitOfMeasureQty(f *field.DerivativeUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasure() (*field.DerivativePriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.DerivativePriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasure reads a DerivativePriceUnitOfMeasure from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePriceUnitOfMeasure(f *field.DerivativePriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceUnitOfMeasureQty() (*field.DerivativePriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.DerivativePriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasureQty reads a DerivativePriceUnitOfMeasureQty from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePriceUnitOfMeasureQty(f *field.DerivativePriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeExerciseStyle() (*field.DerivativeExerciseStyle, errors.MessageRejectError) {
	f := new(field.DerivativeExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeExerciseStyle reads a DerivativeExerciseStyle from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeExerciseStyle(f *field.DerivativeExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeOptPayAmount() (*field.DerivativeOptPayAmount, errors.MessageRejectError) {
	f := new(field.DerivativeOptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptPayAmount reads a DerivativeOptPayAmount from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeOptPayAmount(f *field.DerivativeOptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeTimeUnit() (*field.DerivativeTimeUnit, errors.MessageRejectError) {
	f := new(field.DerivativeTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeTimeUnit reads a DerivativeTimeUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeTimeUnit(f *field.DerivativeTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityExchange() (*field.DerivativeSecurityExchange, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityExchange reads a DerivativeSecurityExchange from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityExchange(f *field.DerivativeSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePositionLimit() (*field.DerivativePositionLimit, errors.MessageRejectError) {
	f := new(field.DerivativePositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePositionLimit reads a DerivativePositionLimit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePositionLimit(f *field.DerivativePositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeNTPositionLimit() (*field.DerivativeNTPositionLimit, errors.MessageRejectError) {
	f := new(field.DerivativeNTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeNTPositionLimit reads a DerivativeNTPositionLimit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeNTPositionLimit(f *field.DerivativeNTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeIssuer() (*field.DerivativeIssuer, errors.MessageRejectError) {
	f := new(field.DerivativeIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssuer reads a DerivativeIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeIssuer(f *field.DerivativeIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuerLen() (*field.DerivativeEncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuerLen reads a DerivativeEncodedIssuerLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedIssuerLen(f *field.DerivativeEncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedIssuer() (*field.DerivativeEncodedIssuer, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuer reads a DerivativeEncodedIssuer from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedIssuer(f *field.DerivativeEncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityDesc() (*field.DerivativeSecurityDesc, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityDesc reads a DerivativeSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityDesc(f *field.DerivativeSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDescLen() (*field.DerivativeEncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDescLen reads a DerivativeEncodedSecurityDescLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedSecurityDescLen(f *field.DerivativeEncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeEncodedSecurityDesc() (*field.DerivativeEncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDesc reads a DerivativeEncodedSecurityDesc from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeEncodedSecurityDesc(f *field.DerivativeEncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractSettlMonth() (*field.DerivativeContractSettlMonth, errors.MessageRejectError) {
	f := new(field.DerivativeContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractSettlMonth reads a DerivativeContractSettlMonth from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeContractSettlMonth(f *field.DerivativeContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeEvents() (*field.NoDerivativeEvents, errors.MessageRejectError) {
	f := new(field.NoDerivativeEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeEvents reads a NoDerivativeEvents from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeEvents(f *field.NoDerivativeEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrumentParties() (*field.NoDerivativeInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoDerivativeInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrumentParties reads a NoDerivativeInstrumentParties from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeInstrumentParties(f *field.NoDerivativeInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSettlMethod() (*field.DerivativeSettlMethod, errors.MessageRejectError) {
	f := new(field.DerivativeSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettlMethod reads a DerivativeSettlMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSettlMethod(f *field.DerivativeSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePriceQuoteMethod() (*field.DerivativePriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.DerivativePriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceQuoteMethod reads a DerivativePriceQuoteMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePriceQuoteMethod(f *field.DerivativePriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeValuationMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeValuationMethod() (*field.DerivativeValuationMethod, errors.MessageRejectError) {
	f := new(field.DerivativeValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeValuationMethod reads a DerivativeValuationMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeValuationMethod(f *field.DerivativeValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeListMethod() (*field.DerivativeListMethod, errors.MessageRejectError) {
	f := new(field.DerivativeListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeListMethod reads a DerivativeListMethod from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeListMethod(f *field.DerivativeListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeCapPrice() (*field.DerivativeCapPrice, errors.MessageRejectError) {
	f := new(field.DerivativeCapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCapPrice reads a DerivativeCapPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeCapPrice(f *field.DerivativeCapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFloorPrice() (*field.DerivativeFloorPrice, errors.MessageRejectError) {
	f := new(field.DerivativeFloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFloorPrice reads a DerivativeFloorPrice from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeFloorPrice(f *field.DerivativeFloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativePutOrCall() (*field.DerivativePutOrCall, errors.MessageRejectError) {
	f := new(field.DerivativePutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePutOrCall reads a DerivativePutOrCall from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativePutOrCall(f *field.DerivativePutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLLen() (*field.DerivativeSecurityXMLLen, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLLen reads a DerivativeSecurityXMLLen from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityXMLLen(f *field.DerivativeSecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXML() (*field.DerivativeSecurityXML, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXML reads a DerivativeSecurityXML from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityXML(f *field.DerivativeSecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeSecurityXMLSchema() (*field.DerivativeSecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLSchema reads a DerivativeSecurityXMLSchema from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeSecurityXMLSchema(f *field.DerivativeSecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractMultiplierUnit is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeContractMultiplierUnit() (*field.DerivativeContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.DerivativeContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractMultiplierUnit reads a DerivativeContractMultiplierUnit from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeContractMultiplierUnit(f *field.DerivativeContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFlowScheduleType is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) DerivativeFlowScheduleType() (*field.DerivativeFlowScheduleType, errors.MessageRejectError) {
	f := new(field.DerivativeFlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFlowScheduleType reads a DerivativeFlowScheduleType from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetDerivativeFlowScheduleType(f *field.DerivativeFlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoDerivativeInstrAttrib() (*field.NoDerivativeInstrAttrib, errors.MessageRejectError) {
	f := new(field.NoDerivativeInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrAttrib reads a NoDerivativeInstrAttrib from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoDerivativeInstrAttrib(f *field.NoDerivativeInstrAttrib) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMarketSegments is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoMarketSegments() (*field.NoMarketSegments, errors.MessageRejectError) {
	f := new(field.NoMarketSegments)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMarketSegments reads a NoMarketSegments from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoMarketSegments(f *field.NoMarketSegments) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSym, errors.MessageRejectError) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetTotNoRelatedSym(f *field.TotNoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from DerivativeSecurityListUpdateReport.
func (m DerivativeSecurityListUpdateReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}
