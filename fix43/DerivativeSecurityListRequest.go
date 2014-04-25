package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateDerivativeSecurityListRequestBuilder returns an initialized DerivativeSecurityListRequestBuilder with specified required fields.
func CreateDerivativeSecurityListRequestBuilder(
	securityreqid field.SecurityReqID,
	securitylistrequesttype field.SecurityListRequestType) DerivativeSecurityListRequestBuilder {
	var builder DerivativeSecurityListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
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
