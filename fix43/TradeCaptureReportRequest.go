package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReportRequest msg type = AD.
type TradeCaptureReportRequest struct {
	message.Message
}

//TradeCaptureReportRequestBuilder builds TradeCaptureReportRequest messages.
type TradeCaptureReportRequestBuilder struct {
	message.MessageBuilder
}

//CreateTradeCaptureReportRequestBuilder returns an initialized TradeCaptureReportRequestBuilder with specified required fields.
func CreateTradeCaptureReportRequestBuilder(
	traderequestid field.TradeRequestID,
	traderequesttype field.TradeRequestType) TradeCaptureReportRequestBuilder {
	var builder TradeCaptureReportRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(traderequestid)
	builder.Body.Set(traderequesttype)
	return builder
}

//TradeRequestID is a required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeRequestID() (field.TradeRequestID, error) {
	var f field.TradeRequestID
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestType is a required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeRequestType() (field.TradeRequestType, error) {
	var f field.TradeRequestType
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MatchStatus() (field.MatchStatus, error) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//NoDates is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoDates() (field.NoDates, error) {
	var f field.NoDates
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TradeInputSource is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeInputSource() (field.TradeInputSource, error) {
	var f field.TradeInputSource
	err := m.Body.Get(&f)
	return f, err
}

//TradeInputDevice is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeInputDevice() (field.TradeInputDevice, error) {
	var f field.TradeInputDevice
	err := m.Body.Get(&f)
	return f, err
}
