package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderSingle msg type = D.
type NewOrderSingle struct {
	message.Message
}

//NewOrderSingleBuilder builds NewOrderSingle messages.
type NewOrderSingleBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderSingleBuilder returns an initialized NewOrderSingleBuilder with specified required fields.
func CreateNewOrderSingleBuilder(
	clordid *field.ClOrdIDField,
	handlinst *field.HandlInstField,
	symbol *field.SymbolField,
	side *field.SideField,
	orderqty *field.OrderQtyField,
	ordtype *field.OrdTypeField) NewOrderSingleBuilder {
	var builder NewOrderSingleBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.NewMsgType("D"))
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderSingle.
func (m NewOrderSingle) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from NewOrderSingle.
func (m NewOrderSingle) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ClientID() (*field.ClientIDField, errors.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from NewOrderSingle.
func (m NewOrderSingle) GetClientID(f *field.ClientIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExecBroker() (*field.ExecBrokerField, errors.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from NewOrderSingle.
func (m NewOrderSingle) GetExecBroker(f *field.ExecBrokerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from NewOrderSingle.
func (m NewOrderSingle) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from NewOrderSingle.
func (m NewOrderSingle) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from NewOrderSingle.
func (m NewOrderSingle) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for NewOrderSingle.
func (m NewOrderSingle) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderSingle.
func (m NewOrderSingle) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderSingle.
func (m NewOrderSingle) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderSingle.
func (m NewOrderSingle) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderSingle.
func (m NewOrderSingle) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderSingle.
func (m NewOrderSingle) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderSingle.
func (m NewOrderSingle) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for NewOrderSingle.
func (m NewOrderSingle) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderSingle.
func (m NewOrderSingle) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderSingle.
func (m NewOrderSingle) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderSingle.
func (m NewOrderSingle) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from NewOrderSingle.
func (m NewOrderSingle) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderSingle.
func (m NewOrderSingle) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderSingle.
func (m NewOrderSingle) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderSingle.
func (m NewOrderSingle) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for NewOrderSingle.
func (m NewOrderSingle) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from NewOrderSingle.
func (m NewOrderSingle) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderSingle.
func (m NewOrderSingle) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderSingle.
func (m NewOrderSingle) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a required field for NewOrderSingle.
func (m NewOrderSingle) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from NewOrderSingle.
func (m NewOrderSingle) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderSingle.
func (m NewOrderSingle) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderSingle.
func (m NewOrderSingle) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderSingle.
func (m NewOrderSingle) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderSingle.
func (m NewOrderSingle) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderSingle.
func (m NewOrderSingle) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIid is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IOIid() (*field.IOIidField, errors.MessageRejectError) {
	f := &field.IOIidField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from NewOrderSingle.
func (m NewOrderSingle) GetIOIid(f *field.IOIidField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from NewOrderSingle.
func (m NewOrderSingle) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderSingle.
func (m NewOrderSingle) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderSingle.
func (m NewOrderSingle) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from NewOrderSingle.
func (m NewOrderSingle) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from NewOrderSingle.
func (m NewOrderSingle) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Rule80A() (*field.Rule80AField, errors.MessageRejectError) {
	f := &field.Rule80AField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from NewOrderSingle.
func (m NewOrderSingle) GetRule80A(f *field.Rule80AField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from NewOrderSingle.
func (m NewOrderSingle) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from NewOrderSingle.
func (m NewOrderSingle) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from NewOrderSingle.
func (m NewOrderSingle) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
