package fix42

import (
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
func (m SecurityStatus) SecurityStatusReqID() (field.SecurityStatusReqID, error) {
	var f field.SecurityStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for SecurityStatus.
func (m SecurityStatus) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityStatus.
func (m SecurityStatus) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for SecurityStatus.
func (m SecurityStatus) IDSource() (field.IDSource, error) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityDay() (field.MaturityDay, error) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for SecurityStatus.
func (m SecurityStatus) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for SecurityStatus.
func (m SecurityStatus) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityStatus.
func (m SecurityStatus) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for SecurityStatus.
func (m SecurityStatus) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for SecurityStatus.
func (m SecurityStatus) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for SecurityStatus.
func (m SecurityStatus) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityStatus.
func (m SecurityStatus) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for SecurityStatus.
func (m SecurityStatus) UnsolicitedIndicator() (field.UnsolicitedIndicator, error) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//SecurityTradingStatus is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityTradingStatus() (field.SecurityTradingStatus, error) {
	var f field.SecurityTradingStatus
	err := m.Body.Get(&f)
	return f, err
}

//FinancialStatus is a non-required field for SecurityStatus.
func (m SecurityStatus) FinancialStatus() (field.FinancialStatus, error) {
	var f field.FinancialStatus
	err := m.Body.Get(&f)
	return f, err
}

//CorporateAction is a non-required field for SecurityStatus.
func (m SecurityStatus) CorporateAction() (field.CorporateAction, error) {
	var f field.CorporateAction
	err := m.Body.Get(&f)
	return f, err
}

//HaltReasonChar is a non-required field for SecurityStatus.
func (m SecurityStatus) HaltReasonChar() (field.HaltReasonChar, error) {
	var f field.HaltReasonChar
	err := m.Body.Get(&f)
	return f, err
}

//InViewOfCommon is a non-required field for SecurityStatus.
func (m SecurityStatus) InViewOfCommon() (field.InViewOfCommon, error) {
	var f field.InViewOfCommon
	err := m.Body.Get(&f)
	return f, err
}

//DueToRelated is a non-required field for SecurityStatus.
func (m SecurityStatus) DueToRelated() (field.DueToRelated, error) {
	var f field.DueToRelated
	err := m.Body.Get(&f)
	return f, err
}

//BuyVolume is a non-required field for SecurityStatus.
func (m SecurityStatus) BuyVolume() (field.BuyVolume, error) {
	var f field.BuyVolume
	err := m.Body.Get(&f)
	return f, err
}

//SellVolume is a non-required field for SecurityStatus.
func (m SecurityStatus) SellVolume() (field.SellVolume, error) {
	var f field.SellVolume
	err := m.Body.Get(&f)
	return f, err
}

//HighPx is a non-required field for SecurityStatus.
func (m SecurityStatus) HighPx() (field.HighPx, error) {
	var f field.HighPx
	err := m.Body.Get(&f)
	return f, err
}

//LowPx is a non-required field for SecurityStatus.
func (m SecurityStatus) LowPx() (field.LowPx, error) {
	var f field.LowPx
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a non-required field for SecurityStatus.
func (m SecurityStatus) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for SecurityStatus.
func (m SecurityStatus) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//Adjustment is a non-required field for SecurityStatus.
func (m SecurityStatus) Adjustment() (field.Adjustment, error) {
	var f field.Adjustment
	err := m.Body.Get(&f)
	return f, err
}
