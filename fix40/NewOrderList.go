package fix40

import (
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
	listid field.ListID,
	listseqno field.ListSeqNo,
	listnoords field.ListNoOrds,
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	symbol field.Symbol,
	side field.Side,
	orderqty field.OrderQty,
	ordtype field.OrdType) NewOrderListBuilder {
	var builder NewOrderListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
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
func (m NewOrderList) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//WaveNo is a non-required field for NewOrderList.
func (m NewOrderList) WaveNo() (field.WaveNo, error) {
	var f field.WaveNo
	err := m.Body.Get(&f)
	return f, err
}

//ListSeqNo is a required field for NewOrderList.
func (m NewOrderList) ListSeqNo() (field.ListSeqNo, error) {
	var f field.ListSeqNo
	err := m.Body.Get(&f)
	return f, err
}

//ListNoOrds is a required field for NewOrderList.
func (m NewOrderList) ListNoOrds() (field.ListNoOrds, error) {
	var f field.ListNoOrds
	err := m.Body.Get(&f)
	return f, err
}

//ListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInst() (field.ListExecInst, error) {
	var f field.ListExecInst
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for NewOrderList.
func (m NewOrderList) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClientID is a non-required field for NewOrderList.
func (m NewOrderList) ClientID() (field.ClientID, error) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for NewOrderList.
func (m NewOrderList) ExecBroker() (field.ExecBroker, error) {
	var f field.ExecBroker
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for NewOrderList.
func (m NewOrderList) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for NewOrderList.
func (m NewOrderList) SettlmntTyp() (field.SettlmntTyp, error) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for NewOrderList.
func (m NewOrderList) FutSettDate() (field.FutSettDate, error) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a required field for NewOrderList.
func (m NewOrderList) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for NewOrderList.
func (m NewOrderList) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for NewOrderList.
func (m NewOrderList) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for NewOrderList.
func (m NewOrderList) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for NewOrderList.
func (m NewOrderList) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for NewOrderList.
func (m NewOrderList) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for NewOrderList.
func (m NewOrderList) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for NewOrderList.
func (m NewOrderList) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for NewOrderList.
func (m NewOrderList) IDSource() (field.IDSource, error) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for NewOrderList.
func (m NewOrderList) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for NewOrderList.
func (m NewOrderList) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for NewOrderList.
func (m NewOrderList) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for NewOrderList.
func (m NewOrderList) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for NewOrderList.
func (m NewOrderList) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a required field for NewOrderList.
func (m NewOrderList) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for NewOrderList.
func (m NewOrderList) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for NewOrderList.
func (m NewOrderList) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for NewOrderList.
func (m NewOrderList) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for NewOrderList.
func (m NewOrderList) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for NewOrderList.
func (m NewOrderList) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for NewOrderList.
func (m NewOrderList) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for NewOrderList.
func (m NewOrderList) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for NewOrderList.
func (m NewOrderList) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//Rule80A is a non-required field for NewOrderList.
func (m NewOrderList) Rule80A() (field.Rule80A, error) {
	var f field.Rule80A
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for NewOrderList.
func (m NewOrderList) ForexReq() (field.ForexReq, error) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for NewOrderList.
func (m NewOrderList) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for NewOrderList.
func (m NewOrderList) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
