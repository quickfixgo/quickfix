package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	origclordid *field.OrigClOrdIDField,
	clordid *field.ClOrdIDField,
	symbol *field.SymbolField,
	side *field.SideField) OrderCancelRequestBuilder {
	var builder OrderCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header.Set(field.NewMsgType("F"))
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	return builder
}

//OrigClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelRequest.
func (m OrderCancelRequest) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelRequest.
func (m OrderCancelRequest) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelRequest.
func (m OrderCancelRequest) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelRequest.
func (m OrderCancelRequest) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ClientID() (*field.ClientIDField, errors.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderCancelRequest.
func (m OrderCancelRequest) GetClientID(f *field.ClientIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ExecBroker() (*field.ExecBrokerField, errors.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderCancelRequest.
func (m OrderCancelRequest) GetExecBroker(f *field.ExecBrokerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderCancelRequest.
func (m OrderCancelRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderCancelRequest.
func (m OrderCancelRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from OrderCancelRequest.
func (m OrderCancelRequest) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderCancelRequest.
func (m OrderCancelRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MaturityDay() (*field.MaturityDayField, errors.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from OrderCancelRequest.
func (m OrderCancelRequest) GetMaturityDay(f *field.MaturityDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderCancelRequest.
func (m OrderCancelRequest) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderCancelRequest.
func (m OrderCancelRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderCancelRequest.
func (m OrderCancelRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderCancelRequest.
func (m OrderCancelRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderCancelRequest.
func (m OrderCancelRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from OrderCancelRequest.
func (m OrderCancelRequest) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from OrderCancelRequest.
func (m OrderCancelRequest) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelRequest.
func (m OrderCancelRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
