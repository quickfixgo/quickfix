package fix40

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
	side field.Side,
	orderqty field.OrderQty,
	lastshares field.LastShares,
	lastpx field.LastPx) DontKnowTradeBuilder {
	var builder DontKnowTradeBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(dkreason)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	builder.Body.Set(lastshares)
	builder.Body.Set(lastpx)
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

//Side is a required field for DontKnowTrade.
func (m DontKnowTrade) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a required field for DontKnowTrade.
func (m DontKnowTrade) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//LastShares is a required field for DontKnowTrade.
func (m DontKnowTrade) LastShares() (field.LastShares, error) {
	var f field.LastShares
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a required field for DontKnowTrade.
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
