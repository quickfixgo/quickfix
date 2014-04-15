package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelRequest msg type = F.
type OrderCancelRequest struct {
	message.Message
}

//OrderCancelRequestBuilder builds OrderCancelRequest messages.
type OrderCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderCancelRequestBuilder returns an initialized OrderCancelRequestBuilder with specified required fields.
func CreateOrderCancelRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	side field.Side,
	transacttime field.TransactTime) OrderCancelRequestBuilder {
	var builder OrderCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	return builder
}

//OrigClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) OrigClOrdID() (field.OrigClOrdID, error) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ClOrdLinkID() (field.ClOrdLinkID, error) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//OrigOrdModTime is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrigOrdModTime() (field.OrigOrdModTime, error) {
	var f field.OrigOrdModTime
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for OrderCancelRequest.
func (m OrderCancelRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
