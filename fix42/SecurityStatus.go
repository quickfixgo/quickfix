package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityStatus msg type = f.
type SecurityStatus struct {
	message.Message
}

//SecurityStatusBuilder builds SecurityStatus messages.
type SecurityStatusBuilder struct {
	message.MessageBuilder
}

//CreateSecurityStatusBuilder returns an initialized SecurityStatusBuilder with specified required fields.
func CreateSecurityStatusBuilder(
	symbol field.Symbol) SecurityStatusBuilder {
	var builder SecurityStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(symbol)
	return builder
}

//SecurityStatusReqID is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityStatusReqID() (*field.SecurityStatusReqID, errors.MessageRejectError) {
	f := new(field.SecurityStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatusReqID reads a SecurityStatusReqID from SecurityStatus.
func (m SecurityStatus) GetSecurityStatusReqID(f *field.SecurityStatusReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for SecurityStatus.
func (m SecurityStatus) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from SecurityStatus.
func (m SecurityStatus) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for SecurityStatus.
func (m SecurityStatus) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from SecurityStatus.
func (m SecurityStatus) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from SecurityStatus.
func (m SecurityStatus) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for SecurityStatus.
func (m SecurityStatus) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from SecurityStatus.
func (m SecurityStatus) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SecurityStatus.
func (m SecurityStatus) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from SecurityStatus.
func (m SecurityStatus) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from SecurityStatus.
func (m SecurityStatus) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for SecurityStatus.
func (m SecurityStatus) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from SecurityStatus.
func (m SecurityStatus) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from SecurityStatus.
func (m SecurityStatus) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for SecurityStatus.
func (m SecurityStatus) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from SecurityStatus.
func (m SecurityStatus) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for SecurityStatus.
func (m SecurityStatus) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from SecurityStatus.
func (m SecurityStatus) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for SecurityStatus.
func (m SecurityStatus) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from SecurityStatus.
func (m SecurityStatus) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from SecurityStatus.
func (m SecurityStatus) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for SecurityStatus.
func (m SecurityStatus) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from SecurityStatus.
func (m SecurityStatus) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from SecurityStatus.
func (m SecurityStatus) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from SecurityStatus.
func (m SecurityStatus) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from SecurityStatus.
func (m SecurityStatus) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from SecurityStatus.
func (m SecurityStatus) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from SecurityStatus.
func (m SecurityStatus) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for SecurityStatus.
func (m SecurityStatus) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from SecurityStatus.
func (m SecurityStatus) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityStatus.
func (m SecurityStatus) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityStatus.
func (m SecurityStatus) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for SecurityStatus.
func (m SecurityStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, errors.MessageRejectError) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from SecurityStatus.
func (m SecurityStatus) GetUnsolicitedIndicator(f *field.UnsolicitedIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityTradingStatus is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityTradingStatus() (*field.SecurityTradingStatus, errors.MessageRejectError) {
	f := new(field.SecurityTradingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityTradingStatus reads a SecurityTradingStatus from SecurityStatus.
func (m SecurityStatus) GetSecurityTradingStatus(f *field.SecurityTradingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FinancialStatus is a non-required field for SecurityStatus.
func (m SecurityStatus) FinancialStatus() (*field.FinancialStatus, errors.MessageRejectError) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetFinancialStatus reads a FinancialStatus from SecurityStatus.
func (m SecurityStatus) GetFinancialStatus(f *field.FinancialStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for SecurityStatus.
func (m SecurityStatus) CorporateAction() (*field.CorporateAction, errors.MessageRejectError) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from SecurityStatus.
func (m SecurityStatus) GetCorporateAction(f *field.CorporateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HaltReasonChar is a non-required field for SecurityStatus.
func (m SecurityStatus) HaltReasonChar() (*field.HaltReasonChar, errors.MessageRejectError) {
	f := new(field.HaltReasonChar)
	err := m.Body.Get(f)
	return f, err
}

//GetHaltReasonChar reads a HaltReasonChar from SecurityStatus.
func (m SecurityStatus) GetHaltReasonChar(f *field.HaltReasonChar) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InViewOfCommon is a non-required field for SecurityStatus.
func (m SecurityStatus) InViewOfCommon() (*field.InViewOfCommon, errors.MessageRejectError) {
	f := new(field.InViewOfCommon)
	err := m.Body.Get(f)
	return f, err
}

//GetInViewOfCommon reads a InViewOfCommon from SecurityStatus.
func (m SecurityStatus) GetInViewOfCommon(f *field.InViewOfCommon) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DueToRelated is a non-required field for SecurityStatus.
func (m SecurityStatus) DueToRelated() (*field.DueToRelated, errors.MessageRejectError) {
	f := new(field.DueToRelated)
	err := m.Body.Get(f)
	return f, err
}

//GetDueToRelated reads a DueToRelated from SecurityStatus.
func (m SecurityStatus) GetDueToRelated(f *field.DueToRelated) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BuyVolume is a non-required field for SecurityStatus.
func (m SecurityStatus) BuyVolume() (*field.BuyVolume, errors.MessageRejectError) {
	f := new(field.BuyVolume)
	err := m.Body.Get(f)
	return f, err
}

//GetBuyVolume reads a BuyVolume from SecurityStatus.
func (m SecurityStatus) GetBuyVolume(f *field.BuyVolume) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SellVolume is a non-required field for SecurityStatus.
func (m SecurityStatus) SellVolume() (*field.SellVolume, errors.MessageRejectError) {
	f := new(field.SellVolume)
	err := m.Body.Get(f)
	return f, err
}

//GetSellVolume reads a SellVolume from SecurityStatus.
func (m SecurityStatus) GetSellVolume(f *field.SellVolume) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HighPx is a non-required field for SecurityStatus.
func (m SecurityStatus) HighPx() (*field.HighPx, errors.MessageRejectError) {
	f := new(field.HighPx)
	err := m.Body.Get(f)
	return f, err
}

//GetHighPx reads a HighPx from SecurityStatus.
func (m SecurityStatus) GetHighPx(f *field.HighPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LowPx is a non-required field for SecurityStatus.
func (m SecurityStatus) LowPx() (*field.LowPx, errors.MessageRejectError) {
	f := new(field.LowPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLowPx reads a LowPx from SecurityStatus.
func (m SecurityStatus) GetLowPx(f *field.LowPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for SecurityStatus.
func (m SecurityStatus) LastPx() (*field.LastPx, errors.MessageRejectError) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from SecurityStatus.
func (m SecurityStatus) GetLastPx(f *field.LastPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SecurityStatus.
func (m SecurityStatus) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SecurityStatus.
func (m SecurityStatus) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Adjustment is a non-required field for SecurityStatus.
func (m SecurityStatus) Adjustment() (*field.Adjustment, errors.MessageRejectError) {
	f := new(field.Adjustment)
	err := m.Body.Get(f)
	return f, err
}

//GetAdjustment reads a Adjustment from SecurityStatus.
func (m SecurityStatus) GetAdjustment(f *field.Adjustment) errors.MessageRejectError {
	return m.Body.Get(f)
}
