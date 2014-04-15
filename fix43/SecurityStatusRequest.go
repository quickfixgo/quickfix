package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityStatusRequest msg type = e.
type SecurityStatusRequest struct {
	message.Message
}

//SecurityStatusRequestBuilder builds SecurityStatusRequest messages.
type SecurityStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateSecurityStatusRequestBuilder returns an initialized SecurityStatusRequestBuilder with specified required fields.
func CreateSecurityStatusRequestBuilder(
	securitystatusreqid field.SecurityStatusReqID,
	subscriptionrequesttype field.SubscriptionRequestType) SecurityStatusRequestBuilder {
	var builder SecurityStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securitystatusreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//SecurityStatusReqID is a required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SecurityStatusReqID() (field.SecurityStatusReqID, error) {
	var f field.SecurityStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a required field for SecurityStatusRequest.
func (m SecurityStatusRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityStatusRequest.
func (m SecurityStatusRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}
