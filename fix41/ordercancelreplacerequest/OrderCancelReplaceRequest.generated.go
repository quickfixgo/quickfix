package ordercancelreplacerequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderCancelReplaceRequest is the fix41 OrderCancelReplaceRequest type, MsgType = G
type OrderCancelReplaceRequest struct {
	fix41.Header
	*quickfix.Body
	fix41.Trailer
	Message *quickfix.Message
}

//FromMessage creates a OrderCancelReplaceRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) OrderCancelReplaceRequest {
	return OrderCancelReplaceRequest{
		Header:  fix41.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix41.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m OrderCancelReplaceRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a OrderCancelReplaceRequest initialized with the required fields for OrderCancelReplaceRequest
func New(origclordid field.OrigClOrdIDField, clordid field.ClOrdIDField, handlinst field.HandlInstField, symbol field.SymbolField, side field.SideField, ordtype field.OrdTypeField) (m OrderCancelReplaceRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix41.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

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
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
func (m OrderCancelReplaceRequest) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m OrderCancelReplaceRequest) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m OrderCancelReplaceRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) SetExecInst(v enum.ExecInst) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) SetHandlInst(v enum.HandlInst) {
	m.Set(field.NewHandlInst(v))
}

//SetIDSource sets IDSource, Tag 22
func (m OrderCancelReplaceRequest) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m OrderCancelReplaceRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetOrdType sets OrdType, Tag 40
func (m OrderCancelReplaceRequest) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetPrice sets Price, Tag 44
func (m OrderCancelReplaceRequest) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetRule80A sets Rule80A, Tag 47
func (m OrderCancelReplaceRequest) SetRule80A(v enum.Rule80A) {
	m.Set(field.NewRule80A(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m OrderCancelReplaceRequest) SetSide(v enum.Side) {
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
func (m OrderCancelReplaceRequest) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m OrderCancelReplaceRequest) SetSettlmntTyp(v enum.SettlmntTyp) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m OrderCancelReplaceRequest) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) SetSymbolSfx(v enum.SymbolSfx) {
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
func (m OrderCancelReplaceRequest) SetOpenClose(v enum.OpenClose) {
	m.Set(field.NewOpenClose(v))
}

//SetStopPx sets StopPx, Tag 99
func (m OrderCancelReplaceRequest) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

//SetExDestination sets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) SetExDestination(v enum.ExDestination) {
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
func (m OrderCancelReplaceRequest) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
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
func (m OrderCancelReplaceRequest) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
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
func (m OrderCancelReplaceRequest) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) SetCoveredOrUncovered(v enum.CoveredOrUncovered) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetCustomerOrFirm sets CustomerOrFirm, Tag 204
func (m OrderCancelReplaceRequest) SetCustomerOrFirm(v enum.CustomerOrFirm) {
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
func (m OrderCancelReplaceRequest) SetMaxShow(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxShow(value, scale))
}

//SetPegDifference sets PegDifference, Tag 211
func (m OrderCancelReplaceRequest) SetPegDifference(value decimal.Decimal, scale int32) {
	m.Set(field.NewPegDifference(value, scale))
}

//GetAccount gets Account, Tag 1
func (m OrderCancelReplaceRequest) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m OrderCancelReplaceRequest) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommission gets Commission, Tag 12
func (m OrderCancelReplaceRequest) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m OrderCancelReplaceRequest) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m OrderCancelReplaceRequest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m OrderCancelReplaceRequest) GetExecInst() (v enum.ExecInst, err quickfix.MessageRejectError) {
	var f field.ExecInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m OrderCancelReplaceRequest) GetHandlInst() (v enum.HandlInst, err quickfix.MessageRejectError) {
	var f field.HandlInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIDSource gets IDSource, Tag 22
func (m OrderCancelReplaceRequest) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m OrderCancelReplaceRequest) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m OrderCancelReplaceRequest) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdType gets OrdType, Tag 40
func (m OrderCancelReplaceRequest) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m OrderCancelReplaceRequest) GetOrigClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPrice gets Price, Tag 44
func (m OrderCancelReplaceRequest) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRule80A gets Rule80A, Tag 47
func (m OrderCancelReplaceRequest) GetRule80A() (v enum.Rule80A, err quickfix.MessageRejectError) {
	var f field.Rule80AField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m OrderCancelReplaceRequest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m OrderCancelReplaceRequest) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m OrderCancelReplaceRequest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m OrderCancelReplaceRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m OrderCancelReplaceRequest) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m OrderCancelReplaceRequest) GetSettlmntTyp() (v enum.SettlmntTyp, err quickfix.MessageRejectError) {
	var f field.SettlmntTypField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m OrderCancelReplaceRequest) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m OrderCancelReplaceRequest) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListID gets ListID, Tag 66
func (m OrderCancelReplaceRequest) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m OrderCancelReplaceRequest) GetExecBroker() (v string, err quickfix.MessageRejectError) {
	var f field.ExecBrokerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOpenClose gets OpenClose, Tag 77
func (m OrderCancelReplaceRequest) GetOpenClose() (v enum.OpenClose, err quickfix.MessageRejectError) {
	var f field.OpenCloseField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStopPx gets StopPx, Tag 99
func (m OrderCancelReplaceRequest) GetStopPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StopPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m OrderCancelReplaceRequest) GetExDestination() (v enum.ExDestination, err quickfix.MessageRejectError) {
	var f field.ExDestinationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m OrderCancelReplaceRequest) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m OrderCancelReplaceRequest) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClientID gets ClientID, Tag 109
func (m OrderCancelReplaceRequest) GetClientID() (v string, err quickfix.MessageRejectError) {
	var f field.ClientIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinQty gets MinQty, Tag 110
func (m OrderCancelReplaceRequest) GetMinQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m OrderCancelReplaceRequest) GetMaxFloor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxFloorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m OrderCancelReplaceRequest) GetLocateReqd() (v bool, err quickfix.MessageRejectError) {
	var f field.LocateReqdField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m OrderCancelReplaceRequest) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m OrderCancelReplaceRequest) GetForexReq() (v bool, err quickfix.MessageRejectError) {
	var f field.ForexReqField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m OrderCancelReplaceRequest) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m OrderCancelReplaceRequest) GetCashOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CashOrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m OrderCancelReplaceRequest) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m OrderCancelReplaceRequest) GetOrderQty2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQty2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m OrderCancelReplaceRequest) GetFutSettDate2() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDate2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m OrderCancelReplaceRequest) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m OrderCancelReplaceRequest) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m OrderCancelReplaceRequest) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m OrderCancelReplaceRequest) GetCoveredOrUncovered() (v enum.CoveredOrUncovered, err quickfix.MessageRejectError) {
	var f field.CoveredOrUncoveredField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCustomerOrFirm gets CustomerOrFirm, Tag 204
func (m OrderCancelReplaceRequest) GetCustomerOrFirm() (v enum.CustomerOrFirm, err quickfix.MessageRejectError) {
	var f field.CustomerOrFirmField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m OrderCancelReplaceRequest) GetMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m OrderCancelReplaceRequest) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m OrderCancelReplaceRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m OrderCancelReplaceRequest) GetMaxShow() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxShowField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegDifference gets PegDifference, Tag 211
func (m OrderCancelReplaceRequest) GetPegDifference() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PegDifferenceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
