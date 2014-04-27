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
	securityreqid field.SecurityReqID,
	securitylistrequesttype field.SecurityListRequestType) DerivativeSecurityListRequestBuilder {
	var builder DerivativeSecurityListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.BuildMsgType("z"))
	builder.Body.Set(securityreqid)
	builder.Body.Set(securitylistrequesttype)
	return builder
}

//SecurityReqID is a required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListRequestType is a required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecurityListRequestType() (*field.SecurityListRequestType, errors.MessageRejectError) {
	f := new(field.SecurityListRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListRequestType reads a SecurityListRequestType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSecurityListRequestType(f *field.SecurityListRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSymbol() (*field.UnderlyingSymbol, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSymbol(f *field.UnderlyingSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityID, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityID(f *field.UnderlyingSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingProduct() (*field.UnderlyingProduct, errors.MessageRejectError) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingProduct(f *field.UnderlyingProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCFICode() (*field.UnderlyingCFICode, errors.MessageRejectError) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCFICode(f *field.UnderlyingCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityType(f *field.UnderlyingSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDate, errors.MessageRejectError) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingIssueDate(f *field.UnderlyingIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFactor() (*field.UnderlyingFactor, errors.MessageRejectError) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingFactor(f *field.UnderlyingFactor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRating, errors.MessageRejectError) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCreditRating(f *field.UnderlyingCreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, errors.MessageRejectError) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, errors.MessageRejectError) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, errors.MessageRejectError) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, errors.MessageRejectError) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCouponRate(f *field.UnderlyingCouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingIssuer() (*field.UnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingIssuer(f *field.UnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgram, errors.MessageRejectError) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCPProgram(f *field.UnderlyingCPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegType, errors.MessageRejectError) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCPRegType(f *field.UnderlyingCPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCurrency() (*field.UnderlyingCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCurrency(f *field.UnderlyingCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingQty() (*field.UnderlyingQty, errors.MessageRejectError) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingQty(f *field.UnderlyingQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPx() (*field.UnderlyingPx, errors.MessageRejectError) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPx(f *field.UnderlyingPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingEndPrice() (*field.UnderlyingEndPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingEndPrice(f *field.UnderlyingEndPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStartValue() (*field.UnderlyingStartValue, errors.MessageRejectError) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingStartValue(f *field.UnderlyingStartValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingEndValue() (*field.UnderlyingEndValue, errors.MessageRejectError) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingEndValue(f *field.UnderlyingEndValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUnderlyingStips() (*field.NoUnderlyingStips, errors.MessageRejectError) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoUnderlyingStips(f *field.NoUnderlyingStips) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, errors.MessageRejectError) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementType, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSettlementType(f *field.UnderlyingSettlementType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmount, errors.MessageRejectError) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCashAmount(f *field.UnderlyingCashAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCashType() (*field.UnderlyingCashType, errors.MessageRejectError) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCashType(f *field.UnderlyingCashType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, errors.MessageRejectError) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCapValue() (*field.UnderlyingCapValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingCapValue(f *field.UnderlyingCapValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, errors.MessageRejectError) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFXRate() (*field.UnderlyingFXRate, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingFXRate(f *field.UnderlyingFXRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, errors.MessageRejectError) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, errors.MessageRejectError) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSymbol() (*field.DerivativeSymbol, errors.MessageRejectError) {
	f := new(field.DerivativeSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbol reads a DerivativeSymbol from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSymbol(f *field.DerivativeSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSymbolSfx() (*field.DerivativeSymbolSfx, errors.MessageRejectError) {
	f := new(field.DerivativeSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbolSfx reads a DerivativeSymbolSfx from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSymbolSfx(f *field.DerivativeSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityID() (*field.DerivativeSecurityID, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityID reads a DerivativeSecurityID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityID(f *field.DerivativeSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityIDSource() (*field.DerivativeSecurityIDSource, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityIDSource reads a DerivativeSecurityIDSource from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityIDSource(f *field.DerivativeSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeSecurityAltID() (*field.NoDerivativeSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoDerivativeSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeSecurityAltID reads a NoDerivativeSecurityAltID from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoDerivativeSecurityAltID(f *field.NoDerivativeSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProduct is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeProduct() (*field.DerivativeProduct, errors.MessageRejectError) {
	f := new(field.DerivativeProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProduct reads a DerivativeProduct from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeProduct(f *field.DerivativeProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeProductComplex() (*field.DerivativeProductComplex, errors.MessageRejectError) {
	f := new(field.DerivativeProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProductComplex reads a DerivativeProductComplex from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeProductComplex(f *field.DerivativeProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivFlexProductEligibilityIndicator() (*field.DerivFlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.DerivFlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivFlexProductEligibilityIndicator reads a DerivFlexProductEligibilityIndicator from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivFlexProductEligibilityIndicator(f *field.DerivFlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityGroup() (*field.DerivativeSecurityGroup, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityGroup reads a DerivativeSecurityGroup from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityGroup(f *field.DerivativeSecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCFICode() (*field.DerivativeCFICode, errors.MessageRejectError) {
	f := new(field.DerivativeCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCFICode reads a DerivativeCFICode from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeCFICode(f *field.DerivativeCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityType() (*field.DerivativeSecurityType, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityType reads a DerivativeSecurityType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityType(f *field.DerivativeSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecuritySubType() (*field.DerivativeSecuritySubType, errors.MessageRejectError) {
	f := new(field.DerivativeSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecuritySubType reads a DerivativeSecuritySubType from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecuritySubType(f *field.DerivativeSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityMonthYear() (*field.DerivativeMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityMonthYear reads a DerivativeMaturityMonthYear from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMaturityMonthYear(f *field.DerivativeMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityDate() (*field.DerivativeMaturityDate, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityDate reads a DerivativeMaturityDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMaturityDate(f *field.DerivativeMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityTime() (*field.DerivativeMaturityTime, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityTime reads a DerivativeMaturityTime from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMaturityTime(f *field.DerivativeMaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSettleOnOpenFlag() (*field.DerivativeSettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.DerivativeSettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettleOnOpenFlag reads a DerivativeSettleOnOpenFlag from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSettleOnOpenFlag(f *field.DerivativeSettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeInstrmtAssignmentMethod() (*field.DerivativeInstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.DerivativeInstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrmtAssignmentMethod reads a DerivativeInstrmtAssignmentMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeInstrmtAssignmentMethod(f *field.DerivativeInstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityStatus() (*field.DerivativeSecurityStatus, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityStatus reads a DerivativeSecurityStatus from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityStatus(f *field.DerivativeSecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeIssueDate() (*field.DerivativeIssueDate, errors.MessageRejectError) {
	f := new(field.DerivativeIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssueDate reads a DerivativeIssueDate from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeIssueDate(f *field.DerivativeIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeInstrRegistry() (*field.DerivativeInstrRegistry, errors.MessageRejectError) {
	f := new(field.DerivativeInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrRegistry reads a DerivativeInstrRegistry from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeInstrRegistry(f *field.DerivativeInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCountryOfIssue() (*field.DerivativeCountryOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCountryOfIssue reads a DerivativeCountryOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeCountryOfIssue(f *field.DerivativeCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStateOrProvinceOfIssue() (*field.DerivativeStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStateOrProvinceOfIssue reads a DerivativeStateOrProvinceOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStateOrProvinceOfIssue(f *field.DerivativeStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikePrice() (*field.DerivativeStrikePrice, errors.MessageRejectError) {
	f := new(field.DerivativeStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikePrice reads a DerivativeStrikePrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikePrice(f *field.DerivativeStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeLocaleOfIssue() (*field.DerivativeLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeLocaleOfIssue reads a DerivativeLocaleOfIssue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeLocaleOfIssue(f *field.DerivativeLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeCurrency() (*field.DerivativeStrikeCurrency, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeCurrency reads a DerivativeStrikeCurrency from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikeCurrency(f *field.DerivativeStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeMultiplier() (*field.DerivativeStrikeMultiplier, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeMultiplier reads a DerivativeStrikeMultiplier from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikeMultiplier(f *field.DerivativeStrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeValue() (*field.DerivativeStrikeValue, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeValue reads a DerivativeStrikeValue from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeStrikeValue(f *field.DerivativeStrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeOptAttribute() (*field.DerivativeOptAttribute, errors.MessageRejectError) {
	f := new(field.DerivativeOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptAttribute reads a DerivativeOptAttribute from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeOptAttribute(f *field.DerivativeOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeContractMultiplier() (*field.DerivativeContractMultiplier, errors.MessageRejectError) {
	f := new(field.DerivativeContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractMultiplier reads a DerivativeContractMultiplier from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeContractMultiplier(f *field.DerivativeContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMinPriceIncrement() (*field.DerivativeMinPriceIncrement, errors.MessageRejectError) {
	f := new(field.DerivativeMinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrement reads a DerivativeMinPriceIncrement from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrement(f *field.DerivativeMinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMinPriceIncrementAmount() (*field.DerivativeMinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.DerivativeMinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrementAmount reads a DerivativeMinPriceIncrementAmount from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeMinPriceIncrementAmount(f *field.DerivativeMinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeUnitOfMeasure() (*field.DerivativeUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.DerivativeUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasure reads a DerivativeUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasure(f *field.DerivativeUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeUnitOfMeasureQty() (*field.DerivativeUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.DerivativeUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasureQty reads a DerivativeUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeUnitOfMeasureQty(f *field.DerivativeUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceUnitOfMeasure() (*field.DerivativePriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.DerivativePriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasure reads a DerivativePriceUnitOfMeasure from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasure(f *field.DerivativePriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceUnitOfMeasureQty() (*field.DerivativePriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.DerivativePriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasureQty reads a DerivativePriceUnitOfMeasureQty from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePriceUnitOfMeasureQty(f *field.DerivativePriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeExerciseStyle() (*field.DerivativeExerciseStyle, errors.MessageRejectError) {
	f := new(field.DerivativeExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeExerciseStyle reads a DerivativeExerciseStyle from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeExerciseStyle(f *field.DerivativeExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeOptPayAmount() (*field.DerivativeOptPayAmount, errors.MessageRejectError) {
	f := new(field.DerivativeOptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptPayAmount reads a DerivativeOptPayAmount from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeOptPayAmount(f *field.DerivativeOptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeTimeUnit() (*field.DerivativeTimeUnit, errors.MessageRejectError) {
	f := new(field.DerivativeTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeTimeUnit reads a DerivativeTimeUnit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeTimeUnit(f *field.DerivativeTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityExchange() (*field.DerivativeSecurityExchange, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityExchange reads a DerivativeSecurityExchange from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityExchange(f *field.DerivativeSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePositionLimit() (*field.DerivativePositionLimit, errors.MessageRejectError) {
	f := new(field.DerivativePositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePositionLimit reads a DerivativePositionLimit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePositionLimit(f *field.DerivativePositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeNTPositionLimit() (*field.DerivativeNTPositionLimit, errors.MessageRejectError) {
	f := new(field.DerivativeNTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeNTPositionLimit reads a DerivativeNTPositionLimit from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeNTPositionLimit(f *field.DerivativeNTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeIssuer() (*field.DerivativeIssuer, errors.MessageRejectError) {
	f := new(field.DerivativeIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssuer reads a DerivativeIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeIssuer(f *field.DerivativeIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedIssuerLen() (*field.DerivativeEncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuerLen reads a DerivativeEncodedIssuerLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuerLen(f *field.DerivativeEncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedIssuer() (*field.DerivativeEncodedIssuer, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuer reads a DerivativeEncodedIssuer from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedIssuer(f *field.DerivativeEncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityDesc() (*field.DerivativeSecurityDesc, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityDesc reads a DerivativeSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityDesc(f *field.DerivativeSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedSecurityDescLen() (*field.DerivativeEncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDescLen reads a DerivativeEncodedSecurityDescLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDescLen(f *field.DerivativeEncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedSecurityDesc() (*field.DerivativeEncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDesc reads a DerivativeEncodedSecurityDesc from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeEncodedSecurityDesc(f *field.DerivativeEncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeContractSettlMonth() (*field.DerivativeContractSettlMonth, errors.MessageRejectError) {
	f := new(field.DerivativeContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractSettlMonth reads a DerivativeContractSettlMonth from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeContractSettlMonth(f *field.DerivativeContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeEvents() (*field.NoDerivativeEvents, errors.MessageRejectError) {
	f := new(field.NoDerivativeEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeEvents reads a NoDerivativeEvents from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoDerivativeEvents(f *field.NoDerivativeEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeInstrumentParties() (*field.NoDerivativeInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoDerivativeInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrumentParties reads a NoDerivativeInstrumentParties from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetNoDerivativeInstrumentParties(f *field.NoDerivativeInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSettlMethod() (*field.DerivativeSettlMethod, errors.MessageRejectError) {
	f := new(field.DerivativeSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettlMethod reads a DerivativeSettlMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSettlMethod(f *field.DerivativeSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceQuoteMethod() (*field.DerivativePriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.DerivativePriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceQuoteMethod reads a DerivativePriceQuoteMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePriceQuoteMethod(f *field.DerivativePriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeFuturesValuationMethod() (*field.DerivativeFuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.DerivativeFuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFuturesValuationMethod reads a DerivativeFuturesValuationMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeFuturesValuationMethod(f *field.DerivativeFuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeListMethod() (*field.DerivativeListMethod, errors.MessageRejectError) {
	f := new(field.DerivativeListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeListMethod reads a DerivativeListMethod from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeListMethod(f *field.DerivativeListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCapPrice() (*field.DerivativeCapPrice, errors.MessageRejectError) {
	f := new(field.DerivativeCapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCapPrice reads a DerivativeCapPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeCapPrice(f *field.DerivativeCapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeFloorPrice() (*field.DerivativeFloorPrice, errors.MessageRejectError) {
	f := new(field.DerivativeFloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFloorPrice reads a DerivativeFloorPrice from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeFloorPrice(f *field.DerivativeFloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePutOrCall() (*field.DerivativePutOrCall, errors.MessageRejectError) {
	f := new(field.DerivativePutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePutOrCall reads a DerivativePutOrCall from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativePutOrCall(f *field.DerivativePutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXMLLen() (*field.DerivativeSecurityXMLLen, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLLen reads a DerivativeSecurityXMLLen from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLLen(f *field.DerivativeSecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXML() (*field.DerivativeSecurityXML, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXML reads a DerivativeSecurityXML from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXML(f *field.DerivativeSecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXMLSchema() (*field.DerivativeSecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLSchema reads a DerivativeSecurityXMLSchema from DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) GetDerivativeSecurityXMLSchema(f *field.DerivativeSecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}
