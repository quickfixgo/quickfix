package executionreport

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//ExecutionReport is the fix40 ExecutionReport type, MsgType = 8
type ExecutionReport struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ExecutionReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) ExecutionReport {
	return ExecutionReport{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ExecutionReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ExecutionReport initialized with the required fields for ExecutionReport
func New(orderid field.OrderIDField, execid field.ExecIDField, exectranstype field.ExecTransTypeField, ordstatus field.OrdStatusField, symbol field.SymbolField, side field.SideField, orderqty field.OrderQtyField, lastshares field.LastSharesField, lastpx field.LastPxField, cumqty field.CumQtyField, avgpx field.AvgPxField) (m ExecutionReport) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("8"))
	m.Header.Set(field.NewBeginString("FIX.4.0"))
	m.Set(orderid)
	m.Set(execid)
	m.Set(exectranstype)
	m.Set(ordstatus)
	m.Set(symbol)
	m.Set(side)
	m.Set(orderqty)
	m.Set(lastshares)
	m.Set(lastpx)
	m.Set(cumqty)
	m.Set(avgpx)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ExecutionReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "8", r
}

//SetAccount sets Account, Tag 1
func (m ExecutionReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAvgPx sets AvgPx, Tag 6
func (m ExecutionReport) SetAvgPx(v float64) {
	m.Set(field.NewAvgPx(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m ExecutionReport) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m ExecutionReport) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m ExecutionReport) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCumQty sets CumQty, Tag 14
func (m ExecutionReport) SetCumQty(v float64) {
	m.Set(field.NewCumQty(v))
}

//SetCurrency sets Currency, Tag 15
func (m ExecutionReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecID sets ExecID, Tag 17
func (m ExecutionReport) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m ExecutionReport) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetExecRefID sets ExecRefID, Tag 19
func (m ExecutionReport) SetExecRefID(v string) {
	m.Set(field.NewExecRefID(v))
}

//SetExecTransType sets ExecTransType, Tag 20
func (m ExecutionReport) SetExecTransType(v string) {
	m.Set(field.NewExecTransType(v))
}

//SetIDSource sets IDSource, Tag 22
func (m ExecutionReport) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetLastCapacity sets LastCapacity, Tag 29
func (m ExecutionReport) SetLastCapacity(v string) {
	m.Set(field.NewLastCapacity(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m ExecutionReport) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetLastPx sets LastPx, Tag 31
func (m ExecutionReport) SetLastPx(v float64) {
	m.Set(field.NewLastPx(v))
}

//SetLastShares sets LastShares, Tag 32
func (m ExecutionReport) SetLastShares(v float64) {
	m.Set(field.NewLastShares(v))
}

//SetOrderID sets OrderID, Tag 37
func (m ExecutionReport) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m ExecutionReport) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetOrdStatus sets OrdStatus, Tag 39
func (m ExecutionReport) SetOrdStatus(v string) {
	m.Set(field.NewOrdStatus(v))
}

//SetOrdType sets OrdType, Tag 40
func (m ExecutionReport) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetPrice sets Price, Tag 44
func (m ExecutionReport) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetRule80A sets Rule80A, Tag 47
func (m ExecutionReport) SetRule80A(v string) {
	m.Set(field.NewRule80A(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m ExecutionReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m ExecutionReport) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m ExecutionReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m ExecutionReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m ExecutionReport) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m ExecutionReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m ExecutionReport) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m ExecutionReport) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m ExecutionReport) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetListID sets ListID, Tag 66
func (m ExecutionReport) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m ExecutionReport) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m ExecutionReport) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetStopPx sets StopPx, Tag 99
func (m ExecutionReport) SetStopPx(v float64) {
	m.Set(field.NewStopPx(v))
}

//SetOrdRejReason sets OrdRejReason, Tag 103
func (m ExecutionReport) SetOrdRejReason(v int) {
	m.Set(field.NewOrdRejReason(v))
}

//SetIssuer sets Issuer, Tag 106
func (m ExecutionReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m ExecutionReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetClientID sets ClientID, Tag 109
func (m ExecutionReport) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetReportToExch sets ReportToExch, Tag 113
func (m ExecutionReport) SetReportToExch(v bool) {
	m.Set(field.NewReportToExch(v))
}

//SetNetMoney sets NetMoney, Tag 118
func (m ExecutionReport) SetNetMoney(v float64) {
	m.Set(field.NewNetMoney(v))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m ExecutionReport) SetSettlCurrAmt(v float64) {
	m.Set(field.NewSettlCurrAmt(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m ExecutionReport) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m ExecutionReport) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetNoMiscFees sets NoMiscFees, Tag 136
func (m ExecutionReport) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

//GetAccount gets Account, Tag 1
func (m ExecutionReport) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAvgPx gets AvgPx, Tag 6
func (m ExecutionReport) GetAvgPx() (f field.AvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m ExecutionReport) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m ExecutionReport) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m ExecutionReport) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCumQty gets CumQty, Tag 14
func (m ExecutionReport) GetCumQty() (f field.CumQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m ExecutionReport) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecID gets ExecID, Tag 17
func (m ExecutionReport) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m ExecutionReport) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecRefID gets ExecRefID, Tag 19
func (m ExecutionReport) GetExecRefID() (f field.ExecRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecTransType gets ExecTransType, Tag 20
func (m ExecutionReport) GetExecTransType() (f field.ExecTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m ExecutionReport) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastCapacity gets LastCapacity, Tag 29
func (m ExecutionReport) GetLastCapacity() (f field.LastCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m ExecutionReport) GetLastMkt() (f field.LastMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m ExecutionReport) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastShares gets LastShares, Tag 32
func (m ExecutionReport) GetLastShares() (f field.LastSharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m ExecutionReport) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m ExecutionReport) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdStatus gets OrdStatus, Tag 39
func (m ExecutionReport) GetOrdStatus() (f field.OrdStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m ExecutionReport) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m ExecutionReport) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRule80A gets Rule80A, Tag 47
func (m ExecutionReport) GetRule80A() (f field.Rule80AField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m ExecutionReport) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m ExecutionReport) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m ExecutionReport) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m ExecutionReport) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m ExecutionReport) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m ExecutionReport) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m ExecutionReport) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m ExecutionReport) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m ExecutionReport) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m ExecutionReport) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m ExecutionReport) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m ExecutionReport) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m ExecutionReport) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdRejReason gets OrdRejReason, Tag 103
func (m ExecutionReport) GetOrdRejReason() (f field.OrdRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m ExecutionReport) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m ExecutionReport) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m ExecutionReport) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetReportToExch gets ReportToExch, Tag 113
func (m ExecutionReport) GetReportToExch() (f field.ReportToExchField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m ExecutionReport) GetNetMoney() (f field.NetMoneyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m ExecutionReport) GetSettlCurrAmt() (f field.SettlCurrAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m ExecutionReport) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m ExecutionReport) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoMiscFees gets NoMiscFees, Tag 136
func (m ExecutionReport) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasAccount returns true if Account is present, Tag 1
func (m ExecutionReport) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m ExecutionReport) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m ExecutionReport) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m ExecutionReport) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m ExecutionReport) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCumQty returns true if CumQty is present, Tag 14
func (m ExecutionReport) HasCumQty() bool {
	return m.Has(tag.CumQty)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m ExecutionReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m ExecutionReport) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m ExecutionReport) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasExecRefID returns true if ExecRefID is present, Tag 19
func (m ExecutionReport) HasExecRefID() bool {
	return m.Has(tag.ExecRefID)
}

//HasExecTransType returns true if ExecTransType is present, Tag 20
func (m ExecutionReport) HasExecTransType() bool {
	return m.Has(tag.ExecTransType)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m ExecutionReport) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasLastCapacity returns true if LastCapacity is present, Tag 29
func (m ExecutionReport) HasLastCapacity() bool {
	return m.Has(tag.LastCapacity)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m ExecutionReport) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m ExecutionReport) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastShares returns true if LastShares is present, Tag 32
func (m ExecutionReport) HasLastShares() bool {
	return m.Has(tag.LastShares)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m ExecutionReport) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m ExecutionReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdStatus returns true if OrdStatus is present, Tag 39
func (m ExecutionReport) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m ExecutionReport) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasPrice returns true if Price is present, Tag 44
func (m ExecutionReport) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasRule80A returns true if Rule80A is present, Tag 47
func (m ExecutionReport) HasRule80A() bool {
	return m.Has(tag.Rule80A)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m ExecutionReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m ExecutionReport) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m ExecutionReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m ExecutionReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m ExecutionReport) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m ExecutionReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m ExecutionReport) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m ExecutionReport) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m ExecutionReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasListID returns true if ListID is present, Tag 66
func (m ExecutionReport) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m ExecutionReport) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m ExecutionReport) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m ExecutionReport) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasOrdRejReason returns true if OrdRejReason is present, Tag 103
func (m ExecutionReport) HasOrdRejReason() bool {
	return m.Has(tag.OrdRejReason)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m ExecutionReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m ExecutionReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m ExecutionReport) HasClientID() bool {
	return m.Has(tag.ClientID)
}

//HasReportToExch returns true if ReportToExch is present, Tag 113
func (m ExecutionReport) HasReportToExch() bool {
	return m.Has(tag.ReportToExch)
}

//HasNetMoney returns true if NetMoney is present, Tag 118
func (m ExecutionReport) HasNetMoney() bool {
	return m.Has(tag.NetMoney)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m ExecutionReport) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m ExecutionReport) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m ExecutionReport) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasNoMiscFees returns true if NoMiscFees is present, Tag 136
func (m ExecutionReport) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
}

//NoMiscFees is a repeating group element, Tag 136
type NoMiscFees struct {
	quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(v float64) {
	m.Set(field.NewMiscFeeAmt(v))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v string) {
	m.Set(field.NewMiscFeeType(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (f field.MiscFeeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (f field.MiscFeeCurrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (f field.MiscFeeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasMiscFeeAmt returns true if MiscFeeAmt is present, Tag 137
func (m NoMiscFees) HasMiscFeeAmt() bool {
	return m.Has(tag.MiscFeeAmt)
}

//HasMiscFeeCurr returns true if MiscFeeCurr is present, Tag 138
func (m NoMiscFees) HasMiscFeeCurr() bool {
	return m.Has(tag.MiscFeeCurr)
}

//HasMiscFeeType returns true if MiscFeeType is present, Tag 139
func (m NoMiscFees) HasMiscFeeType() bool {
	return m.Has(tag.MiscFeeType)
}

//NoMiscFeesRepeatingGroup is a repeating group, Tag 136
type NoMiscFeesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMiscFeesRepeatingGroup returns an initialized, NoMiscFeesRepeatingGroup
func NewNoMiscFeesRepeatingGroup() NoMiscFeesRepeatingGroup {
	return NoMiscFeesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMiscFees,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MiscFeeAmt), quickfix.GroupElement(tag.MiscFeeCurr), quickfix.GroupElement(tag.MiscFeeType)})}
}

//Add create and append a new NoMiscFees to this group
func (m NoMiscFeesRepeatingGroup) Add() NoMiscFees {
	g := m.RepeatingGroup.Add()
	return NoMiscFees{g}
}

//Get returns the ith NoMiscFees in the NoMiscFeesRepeatinGroup
func (m NoMiscFeesRepeatingGroup) Get(i int) NoMiscFees {
	return NoMiscFees{m.RepeatingGroup.Get(i)}
}
