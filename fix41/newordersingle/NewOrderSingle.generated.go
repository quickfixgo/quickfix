package newordersingle

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//NewOrderSingle is the fix41 NewOrderSingle type, MsgType = D
type NewOrderSingle struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a NewOrderSingle from a quickfix.Message instance
func FromMessage(m quickfix.Message) NewOrderSingle {
	return NewOrderSingle{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m NewOrderSingle) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a NewOrderSingle initialized with the required fields for NewOrderSingle
func New(clordid field.ClOrdIDField, handlinst field.HandlInstField, symbol field.SymbolField, side field.SideField, ordtype field.OrdTypeField) (m NewOrderSingle) {
	m.Header = fix41.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("D"))
	m.Set(clordid)
	m.Set(handlinst)
	m.Set(symbol)
	m.Set(side)
	m.Set(ordtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg NewOrderSingle, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "D", r
}

//SetAccount sets Account, Tag 1
func (m NewOrderSingle) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NewOrderSingle) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m NewOrderSingle) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NewOrderSingle) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m NewOrderSingle) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m NewOrderSingle) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m NewOrderSingle) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetIDSource sets IDSource, Tag 22
func (m NewOrderSingle) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetIOIid sets IOIid, Tag 23
func (m NewOrderSingle) SetIOIid(v string) {
	m.Set(field.NewIOIid(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m NewOrderSingle) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetOrdType sets OrdType, Tag 40
func (m NewOrderSingle) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetPrice sets Price, Tag 44
func (m NewOrderSingle) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetRule80A sets Rule80A, Tag 47
func (m NewOrderSingle) SetRule80A(v string) {
	m.Set(field.NewRule80A(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NewOrderSingle) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m NewOrderSingle) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m NewOrderSingle) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m NewOrderSingle) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NewOrderSingle) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m NewOrderSingle) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m NewOrderSingle) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NewOrderSingle) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m NewOrderSingle) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetOpenClose sets OpenClose, Tag 77
func (m NewOrderSingle) SetOpenClose(v string) {
	m.Set(field.NewOpenClose(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NewOrderSingle) SetProcessCode(v string) {
	m.Set(field.NewProcessCode(v))
}

//SetStopPx sets StopPx, Tag 99
func (m NewOrderSingle) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

//SetExDestination sets ExDestination, Tag 100
func (m NewOrderSingle) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NewOrderSingle) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NewOrderSingle) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetClientID sets ClientID, Tag 109
func (m NewOrderSingle) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetMinQty sets MinQty, Tag 110
func (m NewOrderSingle) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m NewOrderSingle) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m NewOrderSingle) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m NewOrderSingle) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NewOrderSingle) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m NewOrderSingle) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NewOrderSingle) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m NewOrderSingle) SetPrevClosePx(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrevClosePx(value, scale))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m NewOrderSingle) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NewOrderSingle) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m NewOrderSingle) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m NewOrderSingle) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NewOrderSingle) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NewOrderSingle) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NewOrderSingle) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m NewOrderSingle) SetCoveredOrUncovered(v int) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetCustomerOrFirm sets CustomerOrFirm, Tag 204
func (m NewOrderSingle) SetCustomerOrFirm(v int) {
	m.Set(field.NewCustomerOrFirm(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m NewOrderSingle) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NewOrderSingle) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NewOrderSingle) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m NewOrderSingle) SetMaxShow(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxShow(value, scale))
}

//SetPegDifference sets PegDifference, Tag 211
func (m NewOrderSingle) SetPegDifference(value decimal.Decimal, scale int32) {
	m.Set(field.NewPegDifference(value, scale))
}

//GetAccount gets Account, Tag 1
func (m NewOrderSingle) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NewOrderSingle) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m NewOrderSingle) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m NewOrderSingle) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NewOrderSingle) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m NewOrderSingle) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m NewOrderSingle) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m NewOrderSingle) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIid gets IOIid, Tag 23
func (m NewOrderSingle) GetIOIid() (f field.IOIidField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m NewOrderSingle) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m NewOrderSingle) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m NewOrderSingle) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRule80A gets Rule80A, Tag 47
func (m NewOrderSingle) GetRule80A() (f field.Rule80AField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NewOrderSingle) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m NewOrderSingle) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NewOrderSingle) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NewOrderSingle) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NewOrderSingle) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m NewOrderSingle) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m NewOrderSingle) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NewOrderSingle) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m NewOrderSingle) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOpenClose gets OpenClose, Tag 77
func (m NewOrderSingle) GetOpenClose() (f field.OpenCloseField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NewOrderSingle) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m NewOrderSingle) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m NewOrderSingle) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NewOrderSingle) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NewOrderSingle) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m NewOrderSingle) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m NewOrderSingle) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m NewOrderSingle) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m NewOrderSingle) GetLocateReqd() (f field.LocateReqdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m NewOrderSingle) GetQuoteID() (f field.QuoteIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NewOrderSingle) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m NewOrderSingle) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NewOrderSingle) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m NewOrderSingle) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m NewOrderSingle) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NewOrderSingle) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m NewOrderSingle) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m NewOrderSingle) GetFutSettDate2() (f field.FutSettDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NewOrderSingle) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NewOrderSingle) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NewOrderSingle) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m NewOrderSingle) GetCoveredOrUncovered() (f field.CoveredOrUncoveredField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustomerOrFirm gets CustomerOrFirm, Tag 204
func (m NewOrderSingle) GetCustomerOrFirm() (f field.CustomerOrFirmField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m NewOrderSingle) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NewOrderSingle) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NewOrderSingle) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m NewOrderSingle) GetMaxShow() (f field.MaxShowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegDifference gets PegDifference, Tag 211
func (m NewOrderSingle) GetPegDifference() (f field.PegDifferenceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m NewOrderSingle) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NewOrderSingle) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NewOrderSingle) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NewOrderSingle) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NewOrderSingle) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m NewOrderSingle) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m NewOrderSingle) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m NewOrderSingle) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasIOIid returns true if IOIid is present, Tag 23
func (m NewOrderSingle) HasIOIid() bool {
	return m.Has(tag.IOIid)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NewOrderSingle) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NewOrderSingle) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasPrice returns true if Price is present, Tag 44
func (m NewOrderSingle) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasRule80A returns true if Rule80A is present, Tag 47
func (m NewOrderSingle) HasRule80A() bool {
	return m.Has(tag.Rule80A)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NewOrderSingle) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m NewOrderSingle) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NewOrderSingle) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m NewOrderSingle) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m NewOrderSingle) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m NewOrderSingle) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m NewOrderSingle) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NewOrderSingle) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m NewOrderSingle) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasOpenClose returns true if OpenClose is present, Tag 77
func (m NewOrderSingle) HasOpenClose() bool {
	return m.Has(tag.OpenClose)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NewOrderSingle) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m NewOrderSingle) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m NewOrderSingle) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NewOrderSingle) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NewOrderSingle) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m NewOrderSingle) HasClientID() bool {
	return m.Has(tag.ClientID)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m NewOrderSingle) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m NewOrderSingle) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m NewOrderSingle) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m NewOrderSingle) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NewOrderSingle) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m NewOrderSingle) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NewOrderSingle) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m NewOrderSingle) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m NewOrderSingle) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NewOrderSingle) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m NewOrderSingle) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m NewOrderSingle) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NewOrderSingle) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NewOrderSingle) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NewOrderSingle) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m NewOrderSingle) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

//HasCustomerOrFirm returns true if CustomerOrFirm is present, Tag 204
func (m NewOrderSingle) HasCustomerOrFirm() bool {
	return m.Has(tag.CustomerOrFirm)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m NewOrderSingle) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NewOrderSingle) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NewOrderSingle) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m NewOrderSingle) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegDifference returns true if PegDifference is present, Tag 211
func (m NewOrderSingle) HasPegDifference() bool {
	return m.Has(tag.PegDifference)
}
