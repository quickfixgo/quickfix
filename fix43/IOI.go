package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//IOI msg type = 6.
type IOI struct {
	message.Message
}

//IOIBuilder builds IOI messages.
type IOIBuilder struct {
	message.MessageBuilder
}

//CreateIOIBuilder returns an initialized IOIBuilder with specified required fields.
func CreateIOIBuilder(
	ioiid field.IOIid,
	ioitranstype field.IOITransType,
	side field.Side,
	ioiqty field.IOIQty) IOIBuilder {
	var builder IOIBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(ioiid)
	builder.Body.Set(ioitranstype)
	builder.Body.Set(side)
	builder.Body.Set(ioiqty)
	return builder
}

//IOIid is a required field for IOI.
func (m IOI) IOIid() (field.IOIid, error) {
	var f field.IOIid
	err := m.Body.Get(&f)
	return f, err
}

//IOITransType is a required field for IOI.
func (m IOI) IOITransType() (field.IOITransType, error) {
	var f field.IOITransType
	err := m.Body.Get(&f)
	return f, err
}

//IOIRefID is a non-required field for IOI.
func (m IOI) IOIRefID() (field.IOIRefID, error) {
	var f field.IOIRefID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for IOI.
func (m IOI) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for IOI.
func (m IOI) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for IOI.
func (m IOI) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for IOI.
func (m IOI) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for IOI.
func (m IOI) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for IOI.
func (m IOI) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for IOI.
func (m IOI) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for IOI.
func (m IOI) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for IOI.
func (m IOI) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for IOI.
func (m IOI) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for IOI.
func (m IOI) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for IOI.
func (m IOI) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for IOI.
func (m IOI) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for IOI.
func (m IOI) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for IOI.
func (m IOI) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for IOI.
func (m IOI) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for IOI.
func (m IOI) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for IOI.
func (m IOI) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for IOI.
func (m IOI) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for IOI.
func (m IOI) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for IOI.
func (m IOI) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for IOI.
func (m IOI) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for IOI.
func (m IOI) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for IOI.
func (m IOI) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for IOI.
func (m IOI) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for IOI.
func (m IOI) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for IOI.
func (m IOI) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for IOI.
func (m IOI) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for IOI.
func (m IOI) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for IOI.
func (m IOI) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for IOI.
func (m IOI) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for IOI.
func (m IOI) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for IOI.
func (m IOI) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for IOI.
func (m IOI) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//QuantityType is a non-required field for IOI.
func (m IOI) QuantityType() (field.QuantityType, error) {
	var f field.QuantityType
	err := m.Body.Get(&f)
	return f, err
}

//IOIQty is a required field for IOI.
func (m IOI) IOIQty() (field.IOIQty, error) {
	var f field.IOIQty
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for IOI.
func (m IOI) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for IOI.
func (m IOI) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for IOI.
func (m IOI) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ValidUntilTime is a non-required field for IOI.
func (m IOI) ValidUntilTime() (field.ValidUntilTime, error) {
	var f field.ValidUntilTime
	err := m.Body.Get(&f)
	return f, err
}

//IOIQltyInd is a non-required field for IOI.
func (m IOI) IOIQltyInd() (field.IOIQltyInd, error) {
	var f field.IOIQltyInd
	err := m.Body.Get(&f)
	return f, err
}

//IOINaturalFlag is a non-required field for IOI.
func (m IOI) IOINaturalFlag() (field.IOINaturalFlag, error) {
	var f field.IOINaturalFlag
	err := m.Body.Get(&f)
	return f, err
}

//NoIOIQualifiers is a non-required field for IOI.
func (m IOI) NoIOIQualifiers() (field.NoIOIQualifiers, error) {
	var f field.NoIOIQualifiers
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for IOI.
func (m IOI) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for IOI.
func (m IOI) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for IOI.
func (m IOI) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for IOI.
func (m IOI) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//URLLink is a non-required field for IOI.
func (m IOI) URLLink() (field.URLLink, error) {
	var f field.URLLink
	err := m.Body.Get(&f)
	return f, err
}

//NoRoutingIDs is a non-required field for IOI.
func (m IOI) NoRoutingIDs() (field.NoRoutingIDs, error) {
	var f field.NoRoutingIDs
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for IOI.
func (m IOI) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for IOI.
func (m IOI) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for IOI.
func (m IOI) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for IOI.
func (m IOI) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//Benchmark is a non-required field for IOI.
func (m IOI) Benchmark() (field.Benchmark, error) {
	var f field.Benchmark
	err := m.Body.Get(&f)
	return f, err
}
