package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CrossOrderCancelRequest msg type = u.
type CrossOrderCancelRequest struct {
	message.Message
}

//CrossOrderCancelRequestBuilder builds CrossOrderCancelRequest messages.
type CrossOrderCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateCrossOrderCancelRequestBuilder returns an initialized CrossOrderCancelRequestBuilder with specified required fields.
func CreateCrossOrderCancelRequestBuilder(
	crossid field.CrossID,
	origcrossid field.OrigCrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	transacttime field.TransactTime) CrossOrderCancelRequestBuilder {
	var builder CrossOrderCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(crossid)
	builder.Body.Set(origcrossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(transacttime)
	return builder
}

//OrderID is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OrderID() (field.OrderID, errors.MessageRejectError) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//CrossID is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CrossID() (field.CrossID, errors.MessageRejectError) {
	var f field.CrossID
	err := m.Body.Get(&f)
	return f, err
}

//OrigCrossID is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OrigCrossID() (field.OrigCrossID, errors.MessageRejectError) {
	var f field.OrigCrossID
	err := m.Body.Get(&f)
	return f, err
}

//CrossType is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CrossType() (field.CrossType, errors.MessageRejectError) {
	var f field.CrossType
	err := m.Body.Get(&f)
	return f, err
}

//CrossPrioritization is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CrossPrioritization() (field.CrossPrioritization, errors.MessageRejectError) {
	var f field.CrossPrioritization
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoSides() (field.NoSides, errors.MessageRejectError) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}
