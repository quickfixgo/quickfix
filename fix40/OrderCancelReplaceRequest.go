package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelReplaceRequest msg type = G.
type OrderCancelReplaceRequest struct {
	message.Message
}

//OrderCancelReplaceRequestBuilder builds OrderCancelReplaceRequest messages.
type OrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderCancelReplaceRequestBuilder returns an initialized OrderCancelReplaceRequestBuilder with specified required fields.
func CreateOrderCancelReplaceRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	symbol field.Symbol,
	side field.Side,
	orderqty field.OrderQty,
	ordtype field.OrdType) OrderCancelReplaceRequestBuilder {
	var builder OrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("G"))
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClientID() (*field.ClientID, errors.MessageRejectError) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClientID(f *field.ClientID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExecBroker() (*field.ExecBroker, errors.MessageRejectError) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExecBroker(f *field.ExecBroker) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdID, errors.MessageRejectError) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrigClOrdID(f *field.OrigClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTyp, errors.MessageRejectError) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlmntTyp(f *field.SettlmntTyp) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Rule80A() (*field.Rule80A, errors.MessageRejectError) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRule80A(f *field.Rule80A) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}
