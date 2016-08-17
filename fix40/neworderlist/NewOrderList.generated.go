package neworderlist

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//NewOrderList is the fix40 NewOrderList type, MsgType = E
type NewOrderList struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a NewOrderList from a quickfix.Message instance
func FromMessage(m quickfix.Message) NewOrderList {
	return NewOrderList{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m NewOrderList) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a NewOrderList initialized with the required fields for NewOrderList
func New(listid field.ListIDField, listseqno field.ListSeqNoField, listnoords field.ListNoOrdsField, clordid field.ClOrdIDField, handlinst field.HandlInstField, symbol field.SymbolField, side field.SideField, orderqty field.OrderQtyField, ordtype field.OrdTypeField) (m NewOrderList) {
	m.Header = fix40.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("E"))
	m.Set(listid)
	m.Set(listseqno)
	m.Set(listnoords)
	m.Set(clordid)
	m.Set(handlinst)
	m.Set(symbol)
	m.Set(side)
	m.Set(orderqty)
	m.Set(ordtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg NewOrderList, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "E", r
}

//SetAccount sets Account, Tag 1
func (m NewOrderList) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NewOrderList) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m NewOrderList) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NewOrderList) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m NewOrderList) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m NewOrderList) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m NewOrderList) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetIDSource sets IDSource, Tag 22
func (m NewOrderList) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m NewOrderList) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetOrdType sets OrdType, Tag 40
func (m NewOrderList) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetPrice sets Price, Tag 44
func (m NewOrderList) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetRule80A sets Rule80A, Tag 47
func (m NewOrderList) SetRule80A(v string) {
	m.Set(field.NewRule80A(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NewOrderList) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m NewOrderList) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m NewOrderList) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m NewOrderList) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NewOrderList) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m NewOrderList) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m NewOrderList) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NewOrderList) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetListID sets ListID, Tag 66
func (m NewOrderList) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetListSeqNo sets ListSeqNo, Tag 67
func (m NewOrderList) SetListSeqNo(v int) {
	m.Set(field.NewListSeqNo(v))
}

//SetListNoOrds sets ListNoOrds, Tag 68
func (m NewOrderList) SetListNoOrds(v int) {
	m.Set(field.NewListNoOrds(v))
}

//SetListExecInst sets ListExecInst, Tag 69
func (m NewOrderList) SetListExecInst(v string) {
	m.Set(field.NewListExecInst(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m NewOrderList) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NewOrderList) SetProcessCode(v string) {
	m.Set(field.NewProcessCode(v))
}

//SetStopPx sets StopPx, Tag 99
func (m NewOrderList) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

//SetExDestination sets ExDestination, Tag 100
func (m NewOrderList) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetWaveNo sets WaveNo, Tag 105
func (m NewOrderList) SetWaveNo(v string) {
	m.Set(field.NewWaveNo(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NewOrderList) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NewOrderList) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetClientID sets ClientID, Tag 109
func (m NewOrderList) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetMinQty sets MinQty, Tag 110
func (m NewOrderList) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m NewOrderList) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m NewOrderList) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NewOrderList) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m NewOrderList) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NewOrderList) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m NewOrderList) SetPrevClosePx(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrevClosePx(value, scale))
}

//GetAccount gets Account, Tag 1
func (m NewOrderList) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NewOrderList) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m NewOrderList) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m NewOrderList) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NewOrderList) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m NewOrderList) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m NewOrderList) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m NewOrderList) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m NewOrderList) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m NewOrderList) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m NewOrderList) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRule80A gets Rule80A, Tag 47
func (m NewOrderList) GetRule80A() (f field.Rule80AField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NewOrderList) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m NewOrderList) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NewOrderList) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NewOrderList) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NewOrderList) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m NewOrderList) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m NewOrderList) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NewOrderList) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m NewOrderList) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListSeqNo gets ListSeqNo, Tag 67
func (m NewOrderList) GetListSeqNo() (f field.ListSeqNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListNoOrds gets ListNoOrds, Tag 68
func (m NewOrderList) GetListNoOrds() (f field.ListNoOrdsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListExecInst gets ListExecInst, Tag 69
func (m NewOrderList) GetListExecInst() (f field.ListExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m NewOrderList) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NewOrderList) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m NewOrderList) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m NewOrderList) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetWaveNo gets WaveNo, Tag 105
func (m NewOrderList) GetWaveNo() (f field.WaveNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NewOrderList) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NewOrderList) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m NewOrderList) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m NewOrderList) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m NewOrderList) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m NewOrderList) GetLocateReqd() (f field.LocateReqdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NewOrderList) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m NewOrderList) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NewOrderList) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m NewOrderList) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m NewOrderList) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NewOrderList) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NewOrderList) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NewOrderList) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NewOrderList) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m NewOrderList) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m NewOrderList) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m NewOrderList) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NewOrderList) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NewOrderList) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasPrice returns true if Price is present, Tag 44
func (m NewOrderList) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasRule80A returns true if Rule80A is present, Tag 47
func (m NewOrderList) HasRule80A() bool {
	return m.Has(tag.Rule80A)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NewOrderList) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m NewOrderList) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NewOrderList) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m NewOrderList) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m NewOrderList) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m NewOrderList) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m NewOrderList) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NewOrderList) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasListID returns true if ListID is present, Tag 66
func (m NewOrderList) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasListSeqNo returns true if ListSeqNo is present, Tag 67
func (m NewOrderList) HasListSeqNo() bool {
	return m.Has(tag.ListSeqNo)
}

//HasListNoOrds returns true if ListNoOrds is present, Tag 68
func (m NewOrderList) HasListNoOrds() bool {
	return m.Has(tag.ListNoOrds)
}

//HasListExecInst returns true if ListExecInst is present, Tag 69
func (m NewOrderList) HasListExecInst() bool {
	return m.Has(tag.ListExecInst)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m NewOrderList) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NewOrderList) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m NewOrderList) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m NewOrderList) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasWaveNo returns true if WaveNo is present, Tag 105
func (m NewOrderList) HasWaveNo() bool {
	return m.Has(tag.WaveNo)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NewOrderList) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NewOrderList) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m NewOrderList) HasClientID() bool {
	return m.Has(tag.ClientID)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m NewOrderList) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m NewOrderList) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m NewOrderList) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NewOrderList) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m NewOrderList) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NewOrderList) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m NewOrderList) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}
