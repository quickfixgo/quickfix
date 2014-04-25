package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderStatusRequest msg type = H.
type OrderStatusRequest struct {
	message.Message
}

//OrderStatusRequestBuilder builds OrderStatusRequest messages.
type OrderStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderStatusRequestBuilder returns an initialized OrderStatusRequestBuilder with specified required fields.
func CreateOrderStatusRequestBuilder(
	clordid field.ClOrdID,
	symbol field.Symbol,
	side field.Side) OrderStatusRequestBuilder {
	var builder OrderStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("H"))
	builder.Body.Set(clordid)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	return builder
}

//OrderID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderStatusRequest.
func (m OrderStatusRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderStatusRequest.
func (m OrderStatusRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderStatusRequest.
func (m OrderStatusRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ClientID() (*field.ClientID, errors.MessageRejectError) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderStatusRequest.
func (m OrderStatusRequest) GetClientID(f *field.ClientID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ExecBroker() (*field.ExecBroker, errors.MessageRejectError) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderStatusRequest.
func (m OrderStatusRequest) GetExecBroker(f *field.ExecBroker) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for OrderStatusRequest.
func (m OrderStatusRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderStatusRequest.
func (m OrderStatusRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderStatusRequest.
func (m OrderStatusRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from OrderStatusRequest.
func (m OrderStatusRequest) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderStatusRequest.
func (m OrderStatusRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from OrderStatusRequest.
func (m OrderStatusRequest) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderStatusRequest.
func (m OrderStatusRequest) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderStatusRequest.
func (m OrderStatusRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderStatusRequest.
func (m OrderStatusRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderStatusRequest.
func (m OrderStatusRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderStatusRequest.
func (m OrderStatusRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderStatusRequest.
func (m OrderStatusRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderStatusRequest.
func (m OrderStatusRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}
