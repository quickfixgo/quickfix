package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Header.Set(field.BuildMsgType("Q"))
	builder.Body.Set(dkreason)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	builder.Body.Set(lastshares)
	builder.Body.Set(lastpx)
	return builder
}

//OrderID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from DontKnowTrade.
func (m DontKnowTrade) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) ExecID() (*field.ExecID, errors.MessageRejectError) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from DontKnowTrade.
func (m DontKnowTrade) GetExecID(f *field.ExecID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DKReason is a required field for DontKnowTrade.
func (m DontKnowTrade) DKReason() (*field.DKReason, errors.MessageRejectError) {
	f := new(field.DKReason)
	err := m.Body.Get(f)
	return f, err
}

//GetDKReason reads a DKReason from DontKnowTrade.
func (m DontKnowTrade) GetDKReason(f *field.DKReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for DontKnowTrade.
func (m DontKnowTrade) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from DontKnowTrade.
func (m DontKnowTrade) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for DontKnowTrade.
func (m DontKnowTrade) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from DontKnowTrade.
func (m DontKnowTrade) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a required field for DontKnowTrade.
func (m DontKnowTrade) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from DontKnowTrade.
func (m DontKnowTrade) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastShares is a required field for DontKnowTrade.
func (m DontKnowTrade) LastShares() (*field.LastShares, errors.MessageRejectError) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}

//GetLastShares reads a LastShares from DontKnowTrade.
func (m DontKnowTrade) GetLastShares(f *field.LastShares) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a required field for DontKnowTrade.
func (m DontKnowTrade) LastPx() (*field.LastPx, errors.MessageRejectError) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from DontKnowTrade.
func (m DontKnowTrade) GetLastPx(f *field.LastPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from DontKnowTrade.
func (m DontKnowTrade) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}
