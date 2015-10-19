//Package dontknowtrade msg type = Q.
package dontknowtrade

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a DontKnowTrade wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrderID is a non-required field for DontKnowTrade.
func (m Message) OrderID() (*field.OrderIDField, quickfix.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from DontKnowTrade.
func (m Message) GetOrderID(f *field.OrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a non-required field for DontKnowTrade.
func (m Message) ExecID() (*field.ExecIDField, quickfix.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from DontKnowTrade.
func (m Message) GetExecID(f *field.ExecIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DKReason is a required field for DontKnowTrade.
func (m Message) DKReason() (*field.DKReasonField, quickfix.MessageRejectError) {
	f := &field.DKReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDKReason reads a DKReason from DontKnowTrade.
func (m Message) GetDKReason(f *field.DKReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for DontKnowTrade.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from DontKnowTrade.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for DontKnowTrade.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from DontKnowTrade.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a required field for DontKnowTrade.
func (m Message) OrderQty() (*field.OrderQtyField, quickfix.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from DontKnowTrade.
func (m Message) GetOrderQty(f *field.OrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastShares is a required field for DontKnowTrade.
func (m Message) LastShares() (*field.LastSharesField, quickfix.MessageRejectError) {
	f := &field.LastSharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastShares reads a LastShares from DontKnowTrade.
func (m Message) GetLastShares(f *field.LastSharesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a required field for DontKnowTrade.
func (m Message) LastPx() (*field.LastPxField, quickfix.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from DontKnowTrade.
func (m Message) GetLastPx(f *field.LastPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for DontKnowTrade.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from DontKnowTrade.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds DontKnowTrade messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for DontKnowTrade.
func Builder(
	dkreason *field.DKReasonField,
	symbol *field.SymbolField,
	side *field.SideField,
	orderqty *field.OrderQtyField,
	lastshares *field.LastSharesField,
	lastpx *field.LastPxField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.NewMsgType("Q"))
	builder.Body.Set(dkreason)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	builder.Body.Set(lastshares)
	builder.Body.Set(lastpx)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "Q", r
}
