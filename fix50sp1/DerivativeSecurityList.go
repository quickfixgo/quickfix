package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
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
func CreateDerivativeSecurityListBuilder(
	securityresponseid field.SecurityResponseID) DerivativeSecurityListBuilder {
	var builder DerivativeSecurityListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityresponseid)
	return builder
}

//SecurityReqID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from DerivativeSecurityList.
func (m DerivativeSecurityList) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for DerivativeSecurityList.
func (m DerivativeSecurityList) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from DerivativeSecurityList.
func (m DerivativeSecurityList) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) SecurityRequestResult() (*field.SecurityRequestResult, errors.MessageRejectError) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from DerivativeSecurityList.
func (m DerivativeSecurityList) GetSecurityRequestResult(f *field.SecurityRequestResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSymbol() (*field.UnderlyingSymbol, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSymbol(f *field.UnderlyingSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityID() (*field.UnderlyingSecurityID, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSecurityID(f *field.UnderlyingSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingProduct() (*field.UnderlyingProduct, errors.MessageRejectError) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingProduct(f *field.UnderlyingProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCFICode() (*field.UnderlyingCFICode, errors.MessageRejectError) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCFICode(f *field.UnderlyingCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityType() (*field.UnderlyingSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSecurityType(f *field.UnderlyingSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingIssueDate() (*field.UnderlyingIssueDate, errors.MessageRejectError) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingIssueDate(f *field.UnderlyingIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingFactor() (*field.UnderlyingFactor, errors.MessageRejectError) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingFactor(f *field.UnderlyingFactor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCreditRating() (*field.UnderlyingCreditRating, errors.MessageRejectError) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCreditRating(f *field.UnderlyingCreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, errors.MessageRejectError) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, errors.MessageRejectError) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, errors.MessageRejectError) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, errors.MessageRejectError) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCouponRate() (*field.UnderlyingCouponRate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCouponRate(f *field.UnderlyingCouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingIssuer() (*field.UnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingIssuer(f *field.UnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from DerivativeSecurityList.
func (m DerivativeSecurityList) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from DerivativeSecurityList.
func (m DerivativeSecurityList) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from DerivativeSecurityList.
func (m DerivativeSecurityList) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from DerivativeSecurityList.
func (m DerivativeSecurityList) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCPProgram() (*field.UnderlyingCPProgram, errors.MessageRejectError) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCPProgram(f *field.UnderlyingCPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCPRegType() (*field.UnderlyingCPRegType, errors.MessageRejectError) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCPRegType(f *field.UnderlyingCPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCurrency() (*field.UnderlyingCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCurrency(f *field.UnderlyingCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingQty() (*field.UnderlyingQty, errors.MessageRejectError) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingQty(f *field.UnderlyingQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPx() (*field.UnderlyingPx, errors.MessageRejectError) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingPx(f *field.UnderlyingPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingEndPrice() (*field.UnderlyingEndPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingEndPrice(f *field.UnderlyingEndPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingStartValue() (*field.UnderlyingStartValue, errors.MessageRejectError) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingStartValue(f *field.UnderlyingStartValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingEndValue() (*field.UnderlyingEndValue, errors.MessageRejectError) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingEndValue(f *field.UnderlyingEndValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoUnderlyingStips() (*field.NoUnderlyingStips, errors.MessageRejectError) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoUnderlyingStips(f *field.NoUnderlyingStips) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, errors.MessageRejectError) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSettlementType() (*field.UnderlyingSettlementType, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSettlementType(f *field.UnderlyingSettlementType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCashAmount() (*field.UnderlyingCashAmount, errors.MessageRejectError) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCashAmount(f *field.UnderlyingCashAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCashType() (*field.UnderlyingCashType, errors.MessageRejectError) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCashType(f *field.UnderlyingCashType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, errors.MessageRejectError) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingCapValue() (*field.UnderlyingCapValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingCapValue(f *field.UnderlyingCapValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, errors.MessageRejectError) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingFXRate() (*field.UnderlyingFXRate, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingFXRate(f *field.UnderlyingFXRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, errors.MessageRejectError) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, errors.MessageRejectError) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from DerivativeSecurityList.
func (m DerivativeSecurityList) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) TotNoRelatedSym() (*field.TotNoRelatedSym, errors.MessageRejectError) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from DerivativeSecurityList.
func (m DerivativeSecurityList) GetTotNoRelatedSym(f *field.TotNoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from DerivativeSecurityList.
func (m DerivativeSecurityList) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbol is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSymbol() (*field.DerivativeSymbol, errors.MessageRejectError) {
	f := new(field.DerivativeSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbol reads a DerivativeSymbol from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSymbol(f *field.DerivativeSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSymbolSfx() (*field.DerivativeSymbolSfx, errors.MessageRejectError) {
	f := new(field.DerivativeSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSymbolSfx reads a DerivativeSymbolSfx from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSymbolSfx(f *field.DerivativeSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityID() (*field.DerivativeSecurityID, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityID reads a DerivativeSecurityID from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityID(f *field.DerivativeSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityIDSource() (*field.DerivativeSecurityIDSource, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityIDSource reads a DerivativeSecurityIDSource from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityIDSource(f *field.DerivativeSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeSecurityAltID() (*field.NoDerivativeSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoDerivativeSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeSecurityAltID reads a NoDerivativeSecurityAltID from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoDerivativeSecurityAltID(f *field.NoDerivativeSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProduct is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeProduct() (*field.DerivativeProduct, errors.MessageRejectError) {
	f := new(field.DerivativeProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProduct reads a DerivativeProduct from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeProduct(f *field.DerivativeProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeProductComplex() (*field.DerivativeProductComplex, errors.MessageRejectError) {
	f := new(field.DerivativeProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeProductComplex reads a DerivativeProductComplex from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeProductComplex(f *field.DerivativeProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivFlexProductEligibilityIndicator() (*field.DerivFlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.DerivFlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivFlexProductEligibilityIndicator reads a DerivFlexProductEligibilityIndicator from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivFlexProductEligibilityIndicator(f *field.DerivFlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityGroup() (*field.DerivativeSecurityGroup, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityGroup reads a DerivativeSecurityGroup from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityGroup(f *field.DerivativeSecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCFICode is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeCFICode() (*field.DerivativeCFICode, errors.MessageRejectError) {
	f := new(field.DerivativeCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCFICode reads a DerivativeCFICode from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeCFICode(f *field.DerivativeCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityType() (*field.DerivativeSecurityType, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityType reads a DerivativeSecurityType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityType(f *field.DerivativeSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecuritySubType() (*field.DerivativeSecuritySubType, errors.MessageRejectError) {
	f := new(field.DerivativeSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecuritySubType reads a DerivativeSecuritySubType from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecuritySubType(f *field.DerivativeSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMaturityMonthYear() (*field.DerivativeMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityMonthYear reads a DerivativeMaturityMonthYear from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeMaturityMonthYear(f *field.DerivativeMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMaturityDate() (*field.DerivativeMaturityDate, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityDate reads a DerivativeMaturityDate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeMaturityDate(f *field.DerivativeMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMaturityTime() (*field.DerivativeMaturityTime, errors.MessageRejectError) {
	f := new(field.DerivativeMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMaturityTime reads a DerivativeMaturityTime from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeMaturityTime(f *field.DerivativeMaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSettleOnOpenFlag() (*field.DerivativeSettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.DerivativeSettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettleOnOpenFlag reads a DerivativeSettleOnOpenFlag from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSettleOnOpenFlag(f *field.DerivativeSettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeInstrmtAssignmentMethod() (*field.DerivativeInstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.DerivativeInstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrmtAssignmentMethod reads a DerivativeInstrmtAssignmentMethod from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeInstrmtAssignmentMethod(f *field.DerivativeInstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityStatus() (*field.DerivativeSecurityStatus, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityStatus reads a DerivativeSecurityStatus from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityStatus(f *field.DerivativeSecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeIssueDate() (*field.DerivativeIssueDate, errors.MessageRejectError) {
	f := new(field.DerivativeIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssueDate reads a DerivativeIssueDate from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeIssueDate(f *field.DerivativeIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeInstrRegistry() (*field.DerivativeInstrRegistry, errors.MessageRejectError) {
	f := new(field.DerivativeInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeInstrRegistry reads a DerivativeInstrRegistry from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeInstrRegistry(f *field.DerivativeInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeCountryOfIssue() (*field.DerivativeCountryOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCountryOfIssue reads a DerivativeCountryOfIssue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeCountryOfIssue(f *field.DerivativeCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStateOrProvinceOfIssue() (*field.DerivativeStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStateOrProvinceOfIssue reads a DerivativeStateOrProvinceOfIssue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeStateOrProvinceOfIssue(f *field.DerivativeStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikePrice() (*field.DerivativeStrikePrice, errors.MessageRejectError) {
	f := new(field.DerivativeStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikePrice reads a DerivativeStrikePrice from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeStrikePrice(f *field.DerivativeStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeLocaleOfIssue() (*field.DerivativeLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.DerivativeLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeLocaleOfIssue reads a DerivativeLocaleOfIssue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeLocaleOfIssue(f *field.DerivativeLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikeCurrency() (*field.DerivativeStrikeCurrency, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeCurrency reads a DerivativeStrikeCurrency from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeStrikeCurrency(f *field.DerivativeStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikeMultiplier() (*field.DerivativeStrikeMultiplier, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeMultiplier reads a DerivativeStrikeMultiplier from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeStrikeMultiplier(f *field.DerivativeStrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeStrikeValue() (*field.DerivativeStrikeValue, errors.MessageRejectError) {
	f := new(field.DerivativeStrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeStrikeValue reads a DerivativeStrikeValue from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeStrikeValue(f *field.DerivativeStrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeOptAttribute() (*field.DerivativeOptAttribute, errors.MessageRejectError) {
	f := new(field.DerivativeOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptAttribute reads a DerivativeOptAttribute from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeOptAttribute(f *field.DerivativeOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeContractMultiplier() (*field.DerivativeContractMultiplier, errors.MessageRejectError) {
	f := new(field.DerivativeContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractMultiplier reads a DerivativeContractMultiplier from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeContractMultiplier(f *field.DerivativeContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMinPriceIncrement() (*field.DerivativeMinPriceIncrement, errors.MessageRejectError) {
	f := new(field.DerivativeMinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrement reads a DerivativeMinPriceIncrement from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeMinPriceIncrement(f *field.DerivativeMinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeMinPriceIncrementAmount() (*field.DerivativeMinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.DerivativeMinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeMinPriceIncrementAmount reads a DerivativeMinPriceIncrementAmount from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeMinPriceIncrementAmount(f *field.DerivativeMinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeUnitOfMeasure() (*field.DerivativeUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.DerivativeUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasure reads a DerivativeUnitOfMeasure from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeUnitOfMeasure(f *field.DerivativeUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeUnitOfMeasureQty() (*field.DerivativeUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.DerivativeUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeUnitOfMeasureQty reads a DerivativeUnitOfMeasureQty from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeUnitOfMeasureQty(f *field.DerivativeUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePriceUnitOfMeasure() (*field.DerivativePriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.DerivativePriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasure reads a DerivativePriceUnitOfMeasure from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativePriceUnitOfMeasure(f *field.DerivativePriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePriceUnitOfMeasureQty() (*field.DerivativePriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.DerivativePriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceUnitOfMeasureQty reads a DerivativePriceUnitOfMeasureQty from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativePriceUnitOfMeasureQty(f *field.DerivativePriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeExerciseStyle() (*field.DerivativeExerciseStyle, errors.MessageRejectError) {
	f := new(field.DerivativeExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeExerciseStyle reads a DerivativeExerciseStyle from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeExerciseStyle(f *field.DerivativeExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeOptPayAmount() (*field.DerivativeOptPayAmount, errors.MessageRejectError) {
	f := new(field.DerivativeOptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeOptPayAmount reads a DerivativeOptPayAmount from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeOptPayAmount(f *field.DerivativeOptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeTimeUnit() (*field.DerivativeTimeUnit, errors.MessageRejectError) {
	f := new(field.DerivativeTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeTimeUnit reads a DerivativeTimeUnit from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeTimeUnit(f *field.DerivativeTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityExchange() (*field.DerivativeSecurityExchange, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityExchange reads a DerivativeSecurityExchange from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityExchange(f *field.DerivativeSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePositionLimit() (*field.DerivativePositionLimit, errors.MessageRejectError) {
	f := new(field.DerivativePositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePositionLimit reads a DerivativePositionLimit from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativePositionLimit(f *field.DerivativePositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeNTPositionLimit() (*field.DerivativeNTPositionLimit, errors.MessageRejectError) {
	f := new(field.DerivativeNTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeNTPositionLimit reads a DerivativeNTPositionLimit from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeNTPositionLimit(f *field.DerivativeNTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeIssuer() (*field.DerivativeIssuer, errors.MessageRejectError) {
	f := new(field.DerivativeIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeIssuer reads a DerivativeIssuer from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeIssuer(f *field.DerivativeIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedIssuerLen() (*field.DerivativeEncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuerLen reads a DerivativeEncodedIssuerLen from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeEncodedIssuerLen(f *field.DerivativeEncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedIssuer() (*field.DerivativeEncodedIssuer, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedIssuer reads a DerivativeEncodedIssuer from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeEncodedIssuer(f *field.DerivativeEncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityDesc() (*field.DerivativeSecurityDesc, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityDesc reads a DerivativeSecurityDesc from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityDesc(f *field.DerivativeSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedSecurityDescLen() (*field.DerivativeEncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDescLen reads a DerivativeEncodedSecurityDescLen from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeEncodedSecurityDescLen(f *field.DerivativeEncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeEncodedSecurityDesc() (*field.DerivativeEncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.DerivativeEncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeEncodedSecurityDesc reads a DerivativeEncodedSecurityDesc from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeEncodedSecurityDesc(f *field.DerivativeEncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeContractSettlMonth() (*field.DerivativeContractSettlMonth, errors.MessageRejectError) {
	f := new(field.DerivativeContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeContractSettlMonth reads a DerivativeContractSettlMonth from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeContractSettlMonth(f *field.DerivativeContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeEvents() (*field.NoDerivativeEvents, errors.MessageRejectError) {
	f := new(field.NoDerivativeEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeEvents reads a NoDerivativeEvents from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoDerivativeEvents(f *field.NoDerivativeEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeInstrumentParties() (*field.NoDerivativeInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoDerivativeInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrumentParties reads a NoDerivativeInstrumentParties from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoDerivativeInstrumentParties(f *field.NoDerivativeInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSettlMethod() (*field.DerivativeSettlMethod, errors.MessageRejectError) {
	f := new(field.DerivativeSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSettlMethod reads a DerivativeSettlMethod from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSettlMethod(f *field.DerivativeSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePriceQuoteMethod() (*field.DerivativePriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.DerivativePriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePriceQuoteMethod reads a DerivativePriceQuoteMethod from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativePriceQuoteMethod(f *field.DerivativePriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFuturesValuationMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeFuturesValuationMethod() (*field.DerivativeFuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.DerivativeFuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFuturesValuationMethod reads a DerivativeFuturesValuationMethod from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeFuturesValuationMethod(f *field.DerivativeFuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeListMethod is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeListMethod() (*field.DerivativeListMethod, errors.MessageRejectError) {
	f := new(field.DerivativeListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeListMethod reads a DerivativeListMethod from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeListMethod(f *field.DerivativeListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeCapPrice() (*field.DerivativeCapPrice, errors.MessageRejectError) {
	f := new(field.DerivativeCapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeCapPrice reads a DerivativeCapPrice from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeCapPrice(f *field.DerivativeCapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeFloorPrice() (*field.DerivativeFloorPrice, errors.MessageRejectError) {
	f := new(field.DerivativeFloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeFloorPrice reads a DerivativeFloorPrice from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeFloorPrice(f *field.DerivativeFloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativePutOrCall() (*field.DerivativePutOrCall, errors.MessageRejectError) {
	f := new(field.DerivativePutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativePutOrCall reads a DerivativePutOrCall from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativePutOrCall(f *field.DerivativePutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityXMLLen() (*field.DerivativeSecurityXMLLen, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLLen reads a DerivativeSecurityXMLLen from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityXMLLen(f *field.DerivativeSecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityXML() (*field.DerivativeSecurityXML, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXML reads a DerivativeSecurityXML from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityXML(f *field.DerivativeSecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) DerivativeSecurityXMLSchema() (*field.DerivativeSecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.DerivativeSecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetDerivativeSecurityXMLSchema reads a DerivativeSecurityXMLSchema from DerivativeSecurityList.
func (m DerivativeSecurityList) GetDerivativeSecurityXMLSchema(f *field.DerivativeSecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDerivativeInstrAttrib is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoDerivativeInstrAttrib() (*field.NoDerivativeInstrAttrib, errors.MessageRejectError) {
	f := new(field.NoDerivativeInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDerivativeInstrAttrib reads a NoDerivativeInstrAttrib from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoDerivativeInstrAttrib(f *field.NoDerivativeInstrAttrib) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMarketSegments is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) NoMarketSegments() (*field.NoMarketSegments, errors.MessageRejectError) {
	f := new(field.NoMarketSegments)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMarketSegments reads a NoMarketSegments from DerivativeSecurityList.
func (m DerivativeSecurityList) GetNoMarketSegments(f *field.NoMarketSegments) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from DerivativeSecurityList.
func (m DerivativeSecurityList) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from DerivativeSecurityList.
func (m DerivativeSecurityList) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from DerivativeSecurityList.
func (m DerivativeSecurityList) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for DerivativeSecurityList.
func (m DerivativeSecurityList) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from DerivativeSecurityList.
func (m DerivativeSecurityList) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
