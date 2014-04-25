package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradingSessionStatus msg type = h.
type TradingSessionStatus struct {
	message.Message
}

//TradingSessionStatusBuilder builds TradingSessionStatus messages.
type TradingSessionStatusBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionStatusBuilder returns an initialized TradingSessionStatusBuilder with specified required fields.
func CreateTradingSessionStatusBuilder(
	tradingsessionid field.TradingSessionID,
	tradsesstatus field.TradSesStatus) TradingSessionStatusBuilder {
	var builder TradingSessionStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("h"))
	builder.Body.Set(tradingsessionid)
	builder.Body.Set(tradsesstatus)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesReqID() (*field.TradSesReqID, errors.MessageRejectError) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesReqID(f *field.TradSesReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionStatus.
func (m TradingSessionStatus) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from TradingSessionStatus.
func (m TradingSessionStatus) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMethod() (*field.TradSesMethod, errors.MessageRejectError) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesMethod(f *field.TradSesMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMode() (*field.TradSesMode, errors.MessageRejectError) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesMode(f *field.TradSesMode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, errors.MessageRejectError) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from TradingSessionStatus.
func (m TradingSessionStatus) GetUnsolicitedIndicator(f *field.UnsolicitedIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStatus is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStatus() (*field.TradSesStatus, errors.MessageRejectError) {
	f := new(field.TradSesStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStatus reads a TradSesStatus from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesStatus(f *field.TradSesStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStatusRejReason is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStatusRejReason() (*field.TradSesStatusRejReason, errors.MessageRejectError) {
	f := new(field.TradSesStatusRejReason)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStatusRejReason reads a TradSesStatusRejReason from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesStatusRejReason(f *field.TradSesStatusRejReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStartTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStartTime() (*field.TradSesStartTime, errors.MessageRejectError) {
	f := new(field.TradSesStartTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStartTime reads a TradSesStartTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesStartTime(f *field.TradSesStartTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesOpenTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesOpenTime() (*field.TradSesOpenTime, errors.MessageRejectError) {
	f := new(field.TradSesOpenTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesOpenTime reads a TradSesOpenTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesOpenTime(f *field.TradSesOpenTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesPreCloseTime() (*field.TradSesPreCloseTime, errors.MessageRejectError) {
	f := new(field.TradSesPreCloseTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesPreCloseTime reads a TradSesPreCloseTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesPreCloseTime(f *field.TradSesPreCloseTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesCloseTime() (*field.TradSesCloseTime, errors.MessageRejectError) {
	f := new(field.TradSesCloseTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesCloseTime reads a TradSesCloseTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesCloseTime(f *field.TradSesCloseTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesEndTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesEndTime() (*field.TradSesEndTime, errors.MessageRejectError) {
	f := new(field.TradSesEndTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesEndTime reads a TradSesEndTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesEndTime(f *field.TradSesEndTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalVolumeTraded is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TotalVolumeTraded() (*field.TotalVolumeTraded, errors.MessageRejectError) {
	f := new(field.TotalVolumeTraded)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalVolumeTraded reads a TotalVolumeTraded from TradingSessionStatus.
func (m TradingSessionStatus) GetTotalVolumeTraded(f *field.TotalVolumeTraded) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradingSessionStatus.
func (m TradingSessionStatus) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from TradingSessionStatus.
func (m TradingSessionStatus) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from TradingSessionStatus.
func (m TradingSessionStatus) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from TradingSessionStatus.
func (m TradingSessionStatus) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from TradingSessionStatus.
func (m TradingSessionStatus) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from TradingSessionStatus.
func (m TradingSessionStatus) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from TradingSessionStatus.
func (m TradingSessionStatus) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from TradingSessionStatus.
func (m TradingSessionStatus) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from TradingSessionStatus.
func (m TradingSessionStatus) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from TradingSessionStatus.
func (m TradingSessionStatus) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from TradingSessionStatus.
func (m TradingSessionStatus) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from TradingSessionStatus.
func (m TradingSessionStatus) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from TradingSessionStatus.
func (m TradingSessionStatus) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from TradingSessionStatus.
func (m TradingSessionStatus) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from TradingSessionStatus.
func (m TradingSessionStatus) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from TradingSessionStatus.
func (m TradingSessionStatus) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from TradingSessionStatus.
func (m TradingSessionStatus) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from TradingSessionStatus.
func (m TradingSessionStatus) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from TradingSessionStatus.
func (m TradingSessionStatus) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from TradingSessionStatus.
func (m TradingSessionStatus) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from TradingSessionStatus.
func (m TradingSessionStatus) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from TradingSessionStatus.
func (m TradingSessionStatus) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from TradingSessionStatus.
func (m TradingSessionStatus) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from TradingSessionStatus.
func (m TradingSessionStatus) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from TradingSessionStatus.
func (m TradingSessionStatus) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from TradingSessionStatus.
func (m TradingSessionStatus) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from TradingSessionStatus.
func (m TradingSessionStatus) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from TradingSessionStatus.
func (m TradingSessionStatus) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from TradingSessionStatus.
func (m TradingSessionStatus) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from TradingSessionStatus.
func (m TradingSessionStatus) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from TradingSessionStatus.
func (m TradingSessionStatus) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from TradingSessionStatus.
func (m TradingSessionStatus) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from TradingSessionStatus.
func (m TradingSessionStatus) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from TradingSessionStatus.
func (m TradingSessionStatus) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from TradingSessionStatus.
func (m TradingSessionStatus) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from TradingSessionStatus.
func (m TradingSessionStatus) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from TradingSessionStatus.
func (m TradingSessionStatus) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from TradingSessionStatus.
func (m TradingSessionStatus) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from TradingSessionStatus.
func (m TradingSessionStatus) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from TradingSessionStatus.
func (m TradingSessionStatus) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from TradingSessionStatus.
func (m TradingSessionStatus) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from TradingSessionStatus.
func (m TradingSessionStatus) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from TradingSessionStatus.
func (m TradingSessionStatus) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from TradingSessionStatus.
func (m TradingSessionStatus) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from TradingSessionStatus.
func (m TradingSessionStatus) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from TradingSessionStatus.
func (m TradingSessionStatus) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from TradingSessionStatus.
func (m TradingSessionStatus) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from TradingSessionStatus.
func (m TradingSessionStatus) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from TradingSessionStatus.
func (m TradingSessionStatus) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from TradingSessionStatus.
func (m TradingSessionStatus) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from TradingSessionStatus.
func (m TradingSessionStatus) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from TradingSessionStatus.
func (m TradingSessionStatus) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from TradingSessionStatus.
func (m TradingSessionStatus) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) OptPayAmount() (*field.OptPayAmount, errors.MessageRejectError) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from TradingSessionStatus.
func (m TradingSessionStatus) GetOptPayAmount(f *field.OptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from TradingSessionStatus.
func (m TradingSessionStatus) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from TradingSessionStatus.
func (m TradingSessionStatus) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from TradingSessionStatus.
func (m TradingSessionStatus) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from TradingSessionStatus.
func (m TradingSessionStatus) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from TradingSessionStatus.
func (m TradingSessionStatus) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from TradingSessionStatus.
func (m TradingSessionStatus) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from TradingSessionStatus.
func (m TradingSessionStatus) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) FuturesValuationMethod() (*field.FuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from TradingSessionStatus.
func (m TradingSessionStatus) GetFuturesValuationMethod(f *field.FuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from TradingSessionStatus.
func (m TradingSessionStatus) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from TradingSessionStatus.
func (m TradingSessionStatus) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesEvent is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesEvent() (*field.TradSesEvent, errors.MessageRejectError) {
	f := new(field.TradSesEvent)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesEvent reads a TradSesEvent from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesEvent(f *field.TradSesEvent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from TradingSessionStatus.
func (m TradingSessionStatus) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from TradingSessionStatus.
func (m TradingSessionStatus) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from TradingSessionStatus.
func (m TradingSessionStatus) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from TradingSessionStatus.
func (m TradingSessionStatus) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
