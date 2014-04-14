package fix40

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

//NewOrderCancelRequestBuilder returns an initialized OrderCancelRequestBuilder with specified required fields.
func NewOrderCancelRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	cxltype field.CxlType,
	symbol field.Symbol,
	side field.Side,
	orderqty field.OrderQty) *OrderCancelRequestBuilder {
	builder := new(OrderCancelRequestBuilder)
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(cxltype)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	return builder
}

//OrigClOrdID is a required field for OrderCancelRequest.
func (m *OrderCancelRequest) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrderID is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a required field for OrderCancelRequest.
func (m *OrderCancelRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ListID is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//CxlType is a required field for OrderCancelRequest.
func (m *OrderCancelRequest) CxlType() (*field.CxlType, error) {
	f := new(field.CxlType)
	err := m.Body.Get(f)
	return f, err
}

//ClientID is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//ExecBroker is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for OrderCancelRequest.
func (m *OrderCancelRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for OrderCancelRequest.
func (m *OrderCancelRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a required field for OrderCancelRequest.
func (m *OrderCancelRequest) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for OrderCancelRequest.
func (m *OrderCancelRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
