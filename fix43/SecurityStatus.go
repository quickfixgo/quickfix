package fix43

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
func CreateSecurityStatusBuilder() SecurityStatusBuilder {
	var builder SecurityStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//SecurityStatusReqID is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityStatusReqID() (field.SecurityStatusReqID, error) {
	var f field.SecurityStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for SecurityStatus.
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

//SecurityIDSource is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for SecurityStatus.
func (m SecurityStatus) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for SecurityStatus.
func (m SecurityStatus) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for SecurityStatus.
func (m SecurityStatus) CFICode() (field.CFICode, error) {
	var f field.CFICode
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

//MaturityDate is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for SecurityStatus.
func (m SecurityStatus) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for SecurityStatus.
func (m SecurityStatus) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for SecurityStatus.
func (m SecurityStatus) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for SecurityStatus.
func (m SecurityStatus) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for SecurityStatus.
func (m SecurityStatus) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for SecurityStatus.
func (m SecurityStatus) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for SecurityStatus.
func (m SecurityStatus) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for SecurityStatus.
func (m SecurityStatus) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for SecurityStatus.
func (m SecurityStatus) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for SecurityStatus.
func (m SecurityStatus) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for SecurityStatus.
func (m SecurityStatus) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for SecurityStatus.
func (m SecurityStatus) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
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

//TradingSessionSubID is a non-required field for SecurityStatus.
func (m SecurityStatus) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
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

//Text is a non-required field for SecurityStatus.
func (m SecurityStatus) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
