package fix41

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//DontKnowTrade msg type = Q.
type DontKnowTrade struct {
	message.Message
}

//DontKnowTradeBuilder builds DontKnowTrade messages.
type DontKnowTradeBuilder struct {
	message.MessageBuilder
}

//CreateDontKnowTradeBuilder returns an initialized DontKnowTradeBuilder with specified required fields.
func CreateDontKnowTradeBuilder(
	dkreason field.DKReason,
	symbol field.Symbol,
	side field.Side) DontKnowTradeBuilder {
	var builder DontKnowTradeBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(dkreason)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	return builder
}

//OrderID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//DKReason is a required field for DontKnowTrade.
func (m DontKnowTrade) DKReason() (field.DKReason, error) {
	var f field.DKReason
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for DontKnowTrade.
func (m DontKnowTrade) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for DontKnowTrade.
func (m DontKnowTrade) IDSource() (field.IDSource, error) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for DontKnowTrade.
func (m DontKnowTrade) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for DontKnowTrade.
func (m DontKnowTrade) MaturityDay() (field.MaturityDay, error) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for DontKnowTrade.
func (m DontKnowTrade) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for DontKnowTrade.
func (m DontKnowTrade) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for DontKnowTrade.
func (m DontKnowTrade) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//LastShares is a non-required field for DontKnowTrade.
func (m DontKnowTrade) LastShares() (field.LastShares, error) {
	var f field.LastShares
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a non-required field for DontKnowTrade.
func (m DontKnowTrade) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
