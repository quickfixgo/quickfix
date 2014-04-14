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

//NewDontKnowTradeBuilder returns an initialized DontKnowTradeBuilder with specified required fields.
func NewDontKnowTradeBuilder(
	dkreason field.DKReason,
	symbol field.Symbol,
	side field.Side,
	orderqty field.OrderQty,
	lastshares field.LastShares,
	lastpx field.LastPx) *DontKnowTradeBuilder {
	builder := new(DontKnowTradeBuilder)
	builder.Body.Set(dkreason)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	builder.Body.Set(lastshares)
	builder.Body.Set(lastpx)
	return builder
}

//OrderID is a non-required field for DontKnowTrade.
func (m *DontKnowTrade) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a non-required field for DontKnowTrade.
func (m *DontKnowTrade) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//DKReason is a required field for DontKnowTrade.
func (m *DontKnowTrade) DKReason() (*field.DKReason, error) {
	f := new(field.DKReason)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for DontKnowTrade.
func (m *DontKnowTrade) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for DontKnowTrade.
func (m *DontKnowTrade) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a required field for DontKnowTrade.
func (m *DontKnowTrade) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//LastShares is a required field for DontKnowTrade.
func (m *DontKnowTrade) LastShares() (*field.LastShares, error) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a required field for DontKnowTrade.
func (m *DontKnowTrade) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for DontKnowTrade.
func (m *DontKnowTrade) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
