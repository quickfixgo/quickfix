package ordercancelreplacerequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderCancelReplaceRequest is the fix41 OrderCancelReplaceRequest type, MsgType = G
type OrderCancelReplaceRequest struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a OrderCancelReplaceRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) OrderCancelReplaceRequest {
	return OrderCancelReplaceRequest{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m OrderCancelReplaceRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a OrderCancelReplaceRequest initialized with the required fields for OrderCancelReplaceRequest
func New(origclordid field.OrigClOrdIDField, clordid field.ClOrdIDField, handlinst field.HandlInstField, symbol field.SymbolField, side field.SideField, ordtype field.OrdTypeField) (m OrderCancelReplaceRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("G"))
	m.Set(origclordid)
	m.Set(clordid)
	m.Set(handlinst)
	m.Set(symbol)
	m.Set(side)
	m.Set(ordtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "G", r
}

//SetAccount sets Account, Tag 1
func (m OrderCancelReplaceRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m OrderCancelReplaceRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m OrderCancelReplaceRequest) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m OrderCancelReplaceRequest) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m OrderCancelReplaceRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetIDSource sets IDSource, Tag 22
func (m OrderCancelReplaceRequest) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m OrderCancelReplaceRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetOrdType sets OrdType, Tag 40
func (m OrderCancelReplaceRequest) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetPrice sets Price, Tag 44
func (m OrderCancelReplaceRequest) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetRule80A sets Rule80A, Tag 47
func (m OrderCancelReplaceRequest) SetRule80A(v string) {
	m.Set(field.NewRule80A(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m OrderCancelReplaceRequest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m OrderCancelReplaceRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m OrderCancelReplaceRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m OrderCancelReplaceRequest) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m OrderCancelReplaceRequest) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m OrderCancelReplaceRequest) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetListID sets ListID, Tag 66
func (m OrderCancelReplaceRequest) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m OrderCancelReplaceRequest) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetOpenClose sets OpenClose, Tag 77
func (m OrderCancelReplaceRequest) SetOpenClose(v string) {
	m.Set(field.NewOpenClose(v))
}

//SetStopPx sets StopPx, Tag 99
func (m OrderCancelReplaceRequest) SetStopPx(v float64) {
	m.Set(field.NewStopPx(v))
}

//SetExDestination sets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m OrderCancelReplaceRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m OrderCancelReplaceRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetClientID sets ClientID, Tag 109
func (m OrderCancelReplaceRequest) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetMinQty sets MinQty, Tag 110
func (m OrderCancelReplaceRequest) SetMinQty(v float64) {
	m.Set(field.NewMinQty(v))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) SetMaxFloor(v float64) {
	m.Set(field.NewMaxFloor(v))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m OrderCancelReplaceRequest) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m OrderCancelReplaceRequest) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m OrderCancelReplaceRequest) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m OrderCancelReplaceRequest) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m OrderCancelReplaceRequest) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) SetOrderQty2(v float64) {
	m.Set(field.NewOrderQty2(v))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m OrderCancelReplaceRequest) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m OrderCancelReplaceRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m OrderCancelReplaceRequest) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) SetCoveredOrUncovered(v int) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetCustomerOrFirm sets CustomerOrFirm, Tag 204
func (m OrderCancelReplaceRequest) SetCustomerOrFirm(v int) {
	m.Set(field.NewCustomerOrFirm(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m OrderCancelReplaceRequest) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m OrderCancelReplaceRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m OrderCancelReplaceRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m OrderCancelReplaceRequest) SetMaxShow(v float64) {
	m.Set(field.NewMaxShow(v))
}

//SetPegDifference sets PegDifference, Tag 211
func (m OrderCancelReplaceRequest) SetPegDifference(v float64) {
	m.Set(field.NewPegDifference(v))
}

//GetAccount gets Account, Tag 1
func (m OrderCancelReplaceRequest) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m OrderCancelReplaceRequest) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m OrderCancelReplaceRequest) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m OrderCancelReplaceRequest) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m OrderCancelReplaceRequest) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m OrderCancelReplaceRequest) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m OrderCancelReplaceRequest) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m OrderCancelReplaceRequest) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) GetOrigClOrdID() (f field.OrigClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m OrderCancelReplaceRequest) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRule80A gets Rule80A, Tag 47
func (m OrderCancelReplaceRequest) GetRule80A() (f field.Rule80AField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m OrderCancelReplaceRequest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m OrderCancelReplaceRequest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m OrderCancelReplaceRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m OrderCancelReplaceRequest) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m OrderCancelReplaceRequest) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m OrderCancelReplaceRequest) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m OrderCancelReplaceRequest) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m OrderCancelReplaceRequest) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOpenClose gets OpenClose, Tag 77
func (m OrderCancelReplaceRequest) GetOpenClose() (f field.OpenCloseField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m OrderCancelReplaceRequest) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m OrderCancelReplaceRequest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m OrderCancelReplaceRequest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m OrderCancelReplaceRequest) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m OrderCancelReplaceRequest) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m OrderCancelReplaceRequest) GetLocateReqd() (f field.LocateReqdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m OrderCancelReplaceRequest) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m OrderCancelReplaceRequest) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m OrderCancelReplaceRequest) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m OrderCancelReplaceRequest) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m OrderCancelReplaceRequest) GetFutSettDate2() (f field.FutSettDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m OrderCancelReplaceRequest) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m OrderCancelReplaceRequest) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) GetCoveredOrUncovered() (f field.CoveredOrUncoveredField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustomerOrFirm gets CustomerOrFirm, Tag 204
func (m OrderCancelReplaceRequest) GetCustomerOrFirm() (f field.CustomerOrFirmField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m OrderCancelReplaceRequest) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m OrderCancelReplaceRequest) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m OrderCancelReplaceRequest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m OrderCancelReplaceRequest) GetMaxShow() (f field.MaxShowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegDifference gets PegDifference, Tag 211
func (m OrderCancelReplaceRequest) GetPegDifference() (f field.PegDifferenceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m OrderCancelReplaceRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m OrderCancelReplaceRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m OrderCancelReplaceRequest) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m OrderCancelReplaceRequest) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m OrderCancelReplaceRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m OrderCancelReplaceRequest) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m OrderCancelReplaceRequest) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m OrderCancelReplaceRequest) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m OrderCancelReplaceRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m OrderCancelReplaceRequest) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m OrderCancelReplaceRequest) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m OrderCancelReplaceRequest) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasPrice returns true if Price is present, Tag 44
func (m OrderCancelReplaceRequest) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasRule80A returns true if Rule80A is present, Tag 47
func (m OrderCancelReplaceRequest) HasRule80A() bool {
	return m.Has(tag.Rule80A)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m OrderCancelReplaceRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m OrderCancelReplaceRequest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m OrderCancelReplaceRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m OrderCancelReplaceRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m OrderCancelReplaceRequest) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m OrderCancelReplaceRequest) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m OrderCancelReplaceRequest) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m OrderCancelReplaceRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasListID returns true if ListID is present, Tag 66
func (m OrderCancelReplaceRequest) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m OrderCancelReplaceRequest) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasOpenClose returns true if OpenClose is present, Tag 77
func (m OrderCancelReplaceRequest) HasOpenClose() bool {
	return m.Has(tag.OpenClose)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m OrderCancelReplaceRequest) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m OrderCancelReplaceRequest) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m OrderCancelReplaceRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m OrderCancelReplaceRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m OrderCancelReplaceRequest) HasClientID() bool {
	return m.Has(tag.ClientID)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m OrderCancelReplaceRequest) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m OrderCancelReplaceRequest) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m OrderCancelReplaceRequest) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m OrderCancelReplaceRequest) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m OrderCancelReplaceRequest) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m OrderCancelReplaceRequest) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m OrderCancelReplaceRequest) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m OrderCancelReplaceRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m OrderCancelReplaceRequest) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m OrderCancelReplaceRequest) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m OrderCancelReplaceRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m OrderCancelReplaceRequest) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m OrderCancelReplaceRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m OrderCancelReplaceRequest) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

//HasCustomerOrFirm returns true if CustomerOrFirm is present, Tag 204
func (m OrderCancelReplaceRequest) HasCustomerOrFirm() bool {
	return m.Has(tag.CustomerOrFirm)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m OrderCancelReplaceRequest) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m OrderCancelReplaceRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m OrderCancelReplaceRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m OrderCancelReplaceRequest) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegDifference returns true if PegDifference is present, Tag 211
func (m OrderCancelReplaceRequest) HasPegDifference() bool {
	return m.Has(tag.PegDifference)
}
