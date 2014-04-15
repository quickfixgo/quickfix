package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityDefinitionRequest msg type = c.
type SecurityDefinitionRequest struct {
	message.Message
}

//SecurityDefinitionRequestBuilder builds SecurityDefinitionRequest messages.
type SecurityDefinitionRequestBuilder struct {
	message.MessageBuilder
}

//CreateSecurityDefinitionRequestBuilder returns an initialized SecurityDefinitionRequestBuilder with specified required fields.
func CreateSecurityDefinitionRequestBuilder(
	securityreqid field.SecurityReqID,
	securityrequesttype field.SecurityRequestType) SecurityDefinitionRequestBuilder {
	var builder SecurityDefinitionRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityrequesttype)
	return builder
}

//SecurityReqID is a required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestType is a required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityRequestType() (field.SecurityRequestType, error) {
	var f field.SecurityRequestType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}
