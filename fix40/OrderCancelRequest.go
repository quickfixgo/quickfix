package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateOrderCancelRequestBuilder returns an initialized OrderCancelRequestBuilder with specified required fields.
func CreateOrderCancelRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	cxltype field.CxlType,
	symbol field.Symbol,
	side field.Side,
	orderqty field.OrderQty) OrderCancelRequestBuilder {
	var builder OrderCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("F"))
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(cxltype)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	return builder
}

//OrigClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) OrigClOrdID() (*field.OrigClOrdID, errors.MessageRejectError) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelRequest.
func (m OrderCancelRequest) GetOrigClOrdID(f *field.OrigClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelRequest.
func (m OrderCancelRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelRequest.
func (m OrderCancelRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelRequest.
func (m OrderCancelRequest) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CxlType is a required field for OrderCancelRequest.
func (m OrderCancelRequest) CxlType() (*field.CxlType, errors.MessageRejectError) {
	f := new(field.CxlType)
	err := m.Body.Get(f)
	return f, err
}

//GetCxlType reads a CxlType from OrderCancelRequest.
func (m OrderCancelRequest) GetCxlType(f *field.CxlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ClientID() (*field.ClientID, errors.MessageRejectError) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderCancelRequest.
func (m OrderCancelRequest) GetClientID(f *field.ClientID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ExecBroker() (*field.ExecBroker, errors.MessageRejectError) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderCancelRequest.
func (m OrderCancelRequest) GetExecBroker(f *field.ExecBroker) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderCancelRequest.
func (m OrderCancelRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderCancelRequest.
func (m OrderCancelRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from OrderCancelRequest.
func (m OrderCancelRequest) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderCancelRequest.
func (m OrderCancelRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderCancelRequest.
func (m OrderCancelRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from OrderCancelRequest.
func (m OrderCancelRequest) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelRequest.
func (m OrderCancelRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}
