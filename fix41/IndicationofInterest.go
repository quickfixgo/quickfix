package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateIndicationofInterestBuilder returns an initialized IndicationofInterestBuilder with specified required fields.
func CreateIndicationofInterestBuilder(
	ioiid field.IOIid,
	ioitranstype field.IOITransType,
	symbol field.Symbol,
	side field.Side,
	ioishares field.IOIShares) IndicationofInterestBuilder {
	var builder IndicationofInterestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(ioiid)
	builder.Body.Set(ioitranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(ioishares)
	return builder
}

//IOIid is a required field for IndicationofInterest.
func (m IndicationofInterest) IOIid() (field.IOIid, errors.MessageRejectError) {
	var f field.IOIid
	err := m.Body.Get(&f)
	return f, err
}

//IOITransType is a required field for IndicationofInterest.
func (m IndicationofInterest) IOITransType() (field.IOITransType, errors.MessageRejectError) {
	var f field.IOITransType
	err := m.Body.Get(&f)
	return f, err
}

//IOIRefID is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIRefID() (field.IOIRefID, errors.MessageRejectError) {
	var f field.IOIRefID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for IndicationofInterest.
func (m IndicationofInterest) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IDSource() (field.IDSource, errors.MessageRejectError) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for IndicationofInterest.
func (m IndicationofInterest) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for IndicationofInterest.
func (m IndicationofInterest) MaturityDay() (field.MaturityDay, errors.MessageRejectError) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for IndicationofInterest.
func (m IndicationofInterest) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for IndicationofInterest.
func (m IndicationofInterest) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for IndicationofInterest.
func (m IndicationofInterest) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for IndicationofInterest.
func (m IndicationofInterest) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//IOIShares is a required field for IndicationofInterest.
func (m IndicationofInterest) IOIShares() (field.IOIShares, errors.MessageRejectError) {
	var f field.IOIShares
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Price() (field.Price, errors.MessageRejectError) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ValidUntilTime is a non-required field for IndicationofInterest.
func (m IndicationofInterest) ValidUntilTime() (field.ValidUntilTime, errors.MessageRejectError) {
	var f field.ValidUntilTime
	err := m.Body.Get(&f)
	return f, err
}

//IOIQltyInd is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIQltyInd() (field.IOIQltyInd, errors.MessageRejectError) {
	var f field.IOIQltyInd
	err := m.Body.Get(&f)
	return f, err
}

//IOIOthSvc is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIOthSvc() (field.IOIOthSvc, errors.MessageRejectError) {
	var f field.IOIOthSvc
	err := m.Body.Get(&f)
	return f, err
}

//IOINaturalFlag is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOINaturalFlag() (field.IOINaturalFlag, errors.MessageRejectError) {
	var f field.IOINaturalFlag
	err := m.Body.Get(&f)
	return f, err
}

//NoIOIQualifiers is a non-required field for IndicationofInterest.
func (m IndicationofInterest) NoIOIQualifiers() (field.NoIOIQualifiers, errors.MessageRejectError) {
	var f field.NoIOIQualifiers
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for IndicationofInterest.
func (m IndicationofInterest) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//URLLink is a non-required field for IndicationofInterest.
func (m IndicationofInterest) URLLink() (field.URLLink, errors.MessageRejectError) {
	var f field.URLLink
	err := m.Body.Get(&f)
	return f, err
}
