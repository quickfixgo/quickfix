package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//IndicationofInterest msg type = 6.
type IndicationofInterest struct {
	message.Message
}

//IndicationofInterestBuilder builds IndicationofInterest messages.
type IndicationofInterestBuilder struct {
	message.MessageBuilder
}

//NewIndicationofInterestBuilder returns an initialized IndicationofInterestBuilder with specified required fields.
func NewIndicationofInterestBuilder(
	ioiid field.IOIid,
	ioitranstype field.IOITransType,
	symbol field.Symbol,
	side field.Side,
	ioishares field.IOIShares) *IndicationofInterestBuilder {
	builder := new(IndicationofInterestBuilder)
	builder.Body.Set(ioiid)
	builder.Body.Set(ioitranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(ioishares)
	return builder
}

//IOIid is a required field for IndicationofInterest.
func (m *IndicationofInterest) IOIid() (*field.IOIid, error) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}

//IOITransType is a required field for IndicationofInterest.
func (m *IndicationofInterest) IOITransType() (*field.IOITransType, error) {
	f := new(field.IOITransType)
	err := m.Body.Get(f)
	return f, err
}

//IOIRefID is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) IOIRefID() (*field.IOIRefID, error) {
	f := new(field.IOIRefID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for IndicationofInterest.
func (m *IndicationofInterest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDay is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for IndicationofInterest.
func (m *IndicationofInterest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//IOIShares is a required field for IndicationofInterest.
func (m *IndicationofInterest) IOIShares() (*field.IOIShares, error) {
	f := new(field.IOIShares)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//ValidUntilTime is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//IOIQltyInd is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) IOIQltyInd() (*field.IOIQltyInd, error) {
	f := new(field.IOIQltyInd)
	err := m.Body.Get(f)
	return f, err
}

//IOINaturalFlag is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) IOINaturalFlag() (*field.IOINaturalFlag, error) {
	f := new(field.IOINaturalFlag)
	err := m.Body.Get(f)
	return f, err
}

//NoIOIQualifiers is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) NoIOIQualifiers() (*field.NoIOIQualifiers, error) {
	f := new(field.NoIOIQualifiers)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//URLLink is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}

//NoRoutingIDs is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//SpreadToBenchmark is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) SpreadToBenchmark() (*field.SpreadToBenchmark, error) {
	f := new(field.SpreadToBenchmark)
	err := m.Body.Get(f)
	return f, err
}

//Benchmark is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) Benchmark() (*field.Benchmark, error) {
	f := new(field.Benchmark)
	err := m.Body.Get(f)
	return f, err
}
