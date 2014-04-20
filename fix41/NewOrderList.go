package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
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
func (m NewOrderList) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from NewOrderList.
func (m NewOrderList) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for NewOrderList.
func (m NewOrderList) WaveNo() (*field.WaveNo, errors.MessageRejectError) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from NewOrderList.
func (m NewOrderList) GetWaveNo(f *field.WaveNo) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListSeqNo is a required field for NewOrderList.
func (m NewOrderList) ListSeqNo() (*field.ListSeqNo, errors.MessageRejectError) {
	f := new(field.ListSeqNo)
	err := m.Body.Get(f)
	return f, err
}

//GetListSeqNo reads a ListSeqNo from NewOrderList.
func (m NewOrderList) GetListSeqNo(f *field.ListSeqNo) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListNoOrds is a required field for NewOrderList.
func (m NewOrderList) ListNoOrds() (*field.ListNoOrds, errors.MessageRejectError) {
	f := new(field.ListNoOrds)
	err := m.Body.Get(f)
	return f, err
}

//GetListNoOrds reads a ListNoOrds from NewOrderList.
func (m NewOrderList) GetListNoOrds(f *field.ListNoOrds) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInst() (*field.ListExecInst, errors.MessageRejectError) {
	f := new(field.ListExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInst reads a ListExecInst from NewOrderList.
func (m NewOrderList) GetListExecInst(f *field.ListExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for NewOrderList.
func (m NewOrderList) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from NewOrderList.
func (m NewOrderList) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for NewOrderList.
func (m NewOrderList) ClientID() (*field.ClientID, errors.MessageRejectError) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from NewOrderList.
func (m NewOrderList) GetClientID(f *field.ClientID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for NewOrderList.
func (m NewOrderList) ExecBroker() (*field.ExecBroker, errors.MessageRejectError) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from NewOrderList.
func (m NewOrderList) GetExecBroker(f *field.ExecBroker) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for NewOrderList.
func (m NewOrderList) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from NewOrderList.
func (m NewOrderList) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for NewOrderList.
func (m NewOrderList) SettlmntTyp() (*field.SettlmntTyp, errors.MessageRejectError) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from NewOrderList.
func (m NewOrderList) GetSettlmntTyp(f *field.SettlmntTyp) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for NewOrderList.
func (m NewOrderList) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from NewOrderList.
func (m NewOrderList) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for NewOrderList.
func (m NewOrderList) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderList.
func (m NewOrderList) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderList.
func (m NewOrderList) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderList.
func (m NewOrderList) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderList.
func (m NewOrderList) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderList.
func (m NewOrderList) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderList.
func (m NewOrderList) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderList.
func (m NewOrderList) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderList.
func (m NewOrderList) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderList.
func (m NewOrderList) ProcessCode() (*field.ProcessCode, errors.MessageRejectError) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderList.
func (m NewOrderList) GetProcessCode(f *field.ProcessCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for NewOrderList.
func (m NewOrderList) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderList.
func (m NewOrderList) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderList.
func (m NewOrderList) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderList.
func (m NewOrderList) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderList.
func (m NewOrderList) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderList.
func (m NewOrderList) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for NewOrderList.
func (m NewOrderList) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from NewOrderList.
func (m NewOrderList) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for NewOrderList.
func (m NewOrderList) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from NewOrderList.
func (m NewOrderList) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for NewOrderList.
func (m NewOrderList) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from NewOrderList.
func (m NewOrderList) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for NewOrderList.
func (m NewOrderList) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from NewOrderList.
func (m NewOrderList) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for NewOrderList.
func (m NewOrderList) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from NewOrderList.
func (m NewOrderList) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for NewOrderList.
func (m NewOrderList) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from NewOrderList.
func (m NewOrderList) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for NewOrderList.
func (m NewOrderList) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from NewOrderList.
func (m NewOrderList) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for NewOrderList.
func (m NewOrderList) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from NewOrderList.
func (m NewOrderList) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderList.
func (m NewOrderList) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderList.
func (m NewOrderList) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderList.
func (m NewOrderList) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderList.
func (m NewOrderList) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderList.
func (m NewOrderList) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderList.
func (m NewOrderList) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for NewOrderList.
func (m NewOrderList) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from NewOrderList.
func (m NewOrderList) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderList.
func (m NewOrderList) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderList.
func (m NewOrderList) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a required field for NewOrderList.
func (m NewOrderList) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from NewOrderList.
func (m NewOrderList) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderList.
func (m NewOrderList) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderList.
func (m NewOrderList) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderList.
func (m NewOrderList) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderList.
func (m NewOrderList) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderList.
func (m NewOrderList) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderList.
func (m NewOrderList) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for NewOrderList.
func (m NewOrderList) PegDifference() (*field.PegDifference, errors.MessageRejectError) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from NewOrderList.
func (m NewOrderList) GetPegDifference(f *field.PegDifference) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderList.
func (m NewOrderList) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderList.
func (m NewOrderList) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderList.
func (m NewOrderList) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderList.
func (m NewOrderList) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderList.
func (m NewOrderList) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderList.
func (m NewOrderList) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for NewOrderList.
func (m NewOrderList) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from NewOrderList.
func (m NewOrderList) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for NewOrderList.
func (m NewOrderList) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from NewOrderList.
func (m NewOrderList) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for NewOrderList.
func (m NewOrderList) Rule80A() (*field.Rule80A, errors.MessageRejectError) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from NewOrderList.
func (m NewOrderList) GetRule80A(f *field.Rule80A) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for NewOrderList.
func (m NewOrderList) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from NewOrderList.
func (m NewOrderList) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for NewOrderList.
func (m NewOrderList) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from NewOrderList.
func (m NewOrderList) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for NewOrderList.
func (m NewOrderList) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from NewOrderList.
func (m NewOrderList) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for NewOrderList.
func (m NewOrderList) FutSettDate2() (*field.FutSettDate2, errors.MessageRejectError) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from NewOrderList.
func (m NewOrderList) GetFutSettDate2(f *field.FutSettDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for NewOrderList.
func (m NewOrderList) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from NewOrderList.
func (m NewOrderList) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenClose is a non-required field for NewOrderList.
func (m NewOrderList) OpenClose() (*field.OpenClose, errors.MessageRejectError) {
	f := new(field.OpenClose)
	err := m.Body.Get(f)
	return f, err
}

//GetOpenClose reads a OpenClose from NewOrderList.
func (m NewOrderList) GetOpenClose(f *field.OpenClose) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for NewOrderList.
func (m NewOrderList) CoveredOrUncovered() (*field.CoveredOrUncovered, errors.MessageRejectError) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from NewOrderList.
func (m NewOrderList) GetCoveredOrUncovered(f *field.CoveredOrUncovered) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustomerOrFirm is a non-required field for NewOrderList.
func (m NewOrderList) CustomerOrFirm() (*field.CustomerOrFirm, errors.MessageRejectError) {
	f := new(field.CustomerOrFirm)
	err := m.Body.Get(f)
	return f, err
}

//GetCustomerOrFirm reads a CustomerOrFirm from NewOrderList.
func (m NewOrderList) GetCustomerOrFirm(f *field.CustomerOrFirm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for NewOrderList.
func (m NewOrderList) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from NewOrderList.
func (m NewOrderList) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}
