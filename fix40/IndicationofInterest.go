package fix40

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

//Issuer is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
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

//IOIOthSvc is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) IOIOthSvc() (*field.IOIOthSvc, error) {
	f := new(field.IOIOthSvc)
	err := m.Body.Get(f)
	return f, err
}

//IOINaturalFlag is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) IOINaturalFlag() (*field.IOINaturalFlag, error) {
	f := new(field.IOINaturalFlag)
	err := m.Body.Get(f)
	return f, err
}

//IOIQualifier is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) IOIQualifier() (*field.IOIQualifier, error) {
	f := new(field.IOIQualifier)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for IndicationofInterest.
func (m *IndicationofInterest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
