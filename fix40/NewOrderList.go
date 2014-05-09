package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderList msg type = E.
type NewOrderList struct {
	message.Message
}

//NewOrderListBuilder builds NewOrderList messages.
type NewOrderListBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderListBuilder returns an initialized NewOrderListBuilder with specified required fields.
func CreateNewOrderListBuilder(
	listid *field.ListIDField,
	listseqno *field.ListSeqNoField,
	listnoords *field.ListNoOrdsField,
	clordid *field.ClOrdIDField,
	handlinst *field.HandlInstField,
	symbol *field.SymbolField,
	side *field.SideField,
	orderqty *field.OrderQtyField,
	ordtype *field.OrdTypeField) NewOrderListBuilder {
	var builder NewOrderListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.NewMsgType("E"))
	builder.Body.Set(listid)
	builder.Body.Set(listseqno)
	builder.Body.Set(listnoords)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	builder.Body.Set(ordtype)
	return builder
}

//ListID is a required field for NewOrderList.
func (m NewOrderList) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from NewOrderList.
func (m NewOrderList) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for NewOrderList.
func (m NewOrderList) WaveNo() (*field.WaveNoField, errors.MessageRejectError) {
	f := &field.WaveNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from NewOrderList.
func (m NewOrderList) GetWaveNo(f *field.WaveNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListSeqNo is a required field for NewOrderList.
func (m NewOrderList) ListSeqNo() (*field.ListSeqNoField, errors.MessageRejectError) {
	f := &field.ListSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListSeqNo reads a ListSeqNo from NewOrderList.
func (m NewOrderList) GetListSeqNo(f *field.ListSeqNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListNoOrds is a required field for NewOrderList.
func (m NewOrderList) ListNoOrds() (*field.ListNoOrdsField, errors.MessageRejectError) {
	f := &field.ListNoOrdsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListNoOrds reads a ListNoOrds from NewOrderList.
func (m NewOrderList) GetListNoOrds(f *field.ListNoOrdsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInst() (*field.ListExecInstField, errors.MessageRejectError) {
	f := &field.ListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInst reads a ListExecInst from NewOrderList.
func (m NewOrderList) GetListExecInst(f *field.ListExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for NewOrderList.
func (m NewOrderList) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from NewOrderList.
func (m NewOrderList) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for NewOrderList.
func (m NewOrderList) ClientID() (*field.ClientIDField, errors.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from NewOrderList.
func (m NewOrderList) GetClientID(f *field.ClientIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for NewOrderList.
func (m NewOrderList) ExecBroker() (*field.ExecBrokerField, errors.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from NewOrderList.
func (m NewOrderList) GetExecBroker(f *field.ExecBrokerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for NewOrderList.
func (m NewOrderList) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from NewOrderList.
func (m NewOrderList) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for NewOrderList.
func (m NewOrderList) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from NewOrderList.
func (m NewOrderList) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for NewOrderList.
func (m NewOrderList) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from NewOrderList.
func (m NewOrderList) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for NewOrderList.
func (m NewOrderList) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderList.
func (m NewOrderList) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderList.
func (m NewOrderList) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderList.
func (m NewOrderList) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderList.
func (m NewOrderList) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderList.
func (m NewOrderList) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderList.
func (m NewOrderList) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderList.
func (m NewOrderList) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderList.
func (m NewOrderList) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderList.
func (m NewOrderList) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderList.
func (m NewOrderList) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for NewOrderList.
func (m NewOrderList) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderList.
func (m NewOrderList) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderList.
func (m NewOrderList) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderList.
func (m NewOrderList) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderList.
func (m NewOrderList) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderList.
func (m NewOrderList) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for NewOrderList.
func (m NewOrderList) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from NewOrderList.
func (m NewOrderList) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderList.
func (m NewOrderList) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderList.
func (m NewOrderList) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderList.
func (m NewOrderList) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderList.
func (m NewOrderList) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderList.
func (m NewOrderList) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderList.
func (m NewOrderList) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for NewOrderList.
func (m NewOrderList) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from NewOrderList.
func (m NewOrderList) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderList.
func (m NewOrderList) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderList.
func (m NewOrderList) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a required field for NewOrderList.
func (m NewOrderList) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from NewOrderList.
func (m NewOrderList) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderList.
func (m NewOrderList) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderList.
func (m NewOrderList) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderList.
func (m NewOrderList) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderList.
func (m NewOrderList) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderList.
func (m NewOrderList) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderList.
func (m NewOrderList) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderList.
func (m NewOrderList) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderList.
func (m NewOrderList) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderList.
func (m NewOrderList) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderList.
func (m NewOrderList) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderList.
func (m NewOrderList) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderList.
func (m NewOrderList) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for NewOrderList.
func (m NewOrderList) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from NewOrderList.
func (m NewOrderList) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for NewOrderList.
func (m NewOrderList) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from NewOrderList.
func (m NewOrderList) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for NewOrderList.
func (m NewOrderList) Rule80A() (*field.Rule80AField, errors.MessageRejectError) {
	f := &field.Rule80AField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from NewOrderList.
func (m NewOrderList) GetRule80A(f *field.Rule80AField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for NewOrderList.
func (m NewOrderList) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from NewOrderList.
func (m NewOrderList) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for NewOrderList.
func (m NewOrderList) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from NewOrderList.
func (m NewOrderList) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for NewOrderList.
func (m NewOrderList) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from NewOrderList.
func (m NewOrderList) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
