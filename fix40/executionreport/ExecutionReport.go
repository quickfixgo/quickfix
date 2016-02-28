//Package executionreport msg type = 8.
package executionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//NoMiscFees is a repeating group in ExecutionReport
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
}

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }

//Message is a ExecutionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"8"`
	Header     fix40.Header
	//OrderID is a required field for ExecutionReport.
	OrderID string `fix:"37"`
	//ClOrdID is a non-required field for ExecutionReport.
	ClOrdID *string `fix:"11"`
	//ClientID is a non-required field for ExecutionReport.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for ExecutionReport.
	ExecBroker *string `fix:"76"`
	//ListID is a non-required field for ExecutionReport.
	ListID *string `fix:"66"`
	//ExecID is a required field for ExecutionReport.
	ExecID int `fix:"17"`
	//ExecTransType is a required field for ExecutionReport.
	ExecTransType string `fix:"20"`
	//ExecRefID is a non-required field for ExecutionReport.
	ExecRefID *int `fix:"19"`
	//OrdStatus is a required field for ExecutionReport.
	OrdStatus string `fix:"39"`
	//OrdRejReason is a non-required field for ExecutionReport.
	OrdRejReason *int `fix:"103"`
	//Account is a non-required field for ExecutionReport.
	Account *string `fix:"1"`
	//SettlmntTyp is a non-required field for ExecutionReport.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for ExecutionReport.
	FutSettDate *string `fix:"64"`
	//Symbol is a required field for ExecutionReport.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for ExecutionReport.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for ExecutionReport.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for ExecutionReport.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for ExecutionReport.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for ExecutionReport.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for ExecutionReport.
	Side string `fix:"54"`
	//OrderQty is a required field for ExecutionReport.
	OrderQty int `fix:"38"`
	//OrdType is a non-required field for ExecutionReport.
	OrdType *string `fix:"40"`
	//Price is a non-required field for ExecutionReport.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for ExecutionReport.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for ExecutionReport.
	Currency *string `fix:"15"`
	//TimeInForce is a non-required field for ExecutionReport.
	TimeInForce *string `fix:"59"`
	//ExpireTime is a non-required field for ExecutionReport.
	ExpireTime *time.Time `fix:"126"`
	//ExecInst is a non-required field for ExecutionReport.
	ExecInst *string `fix:"18"`
	//Rule80A is a non-required field for ExecutionReport.
	Rule80A *string `fix:"47"`
	//LastShares is a required field for ExecutionReport.
	LastShares int `fix:"32"`
	//LastPx is a required field for ExecutionReport.
	LastPx float64 `fix:"31"`
	//LastMkt is a non-required field for ExecutionReport.
	LastMkt *string `fix:"30"`
	//LastCapacity is a non-required field for ExecutionReport.
	LastCapacity *string `fix:"29"`
	//CumQty is a required field for ExecutionReport.
	CumQty int `fix:"14"`
	//AvgPx is a required field for ExecutionReport.
	AvgPx float64 `fix:"6"`
	//TradeDate is a non-required field for ExecutionReport.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for ExecutionReport.
	TransactTime *time.Time `fix:"60"`
	//ReportToExch is a non-required field for ExecutionReport.
	ReportToExch *string `fix:"113"`
	//Commission is a non-required field for ExecutionReport.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for ExecutionReport.
	CommType *string `fix:"13"`
	//NoMiscFees is a non-required field for ExecutionReport.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
	//NetMoney is a non-required field for ExecutionReport.
	NetMoney *float64 `fix:"118"`
	//SettlCurrAmt is a non-required field for ExecutionReport.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for ExecutionReport.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for ExecutionReport.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)          { m.OrderID = v }
func (m *Message) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *Message) SetClientID(v string)         { m.ClientID = &v }
func (m *Message) SetExecBroker(v string)       { m.ExecBroker = &v }
func (m *Message) SetListID(v string)           { m.ListID = &v }
func (m *Message) SetExecID(v int)              { m.ExecID = v }
func (m *Message) SetExecTransType(v string)    { m.ExecTransType = v }
func (m *Message) SetExecRefID(v int)           { m.ExecRefID = &v }
func (m *Message) SetOrdStatus(v string)        { m.OrdStatus = v }
func (m *Message) SetOrdRejReason(v int)        { m.OrdRejReason = &v }
func (m *Message) SetAccount(v string)          { m.Account = &v }
func (m *Message) SetSettlmntTyp(v string)      { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)      { m.FutSettDate = &v }
func (m *Message) SetSymbol(v string)           { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)        { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)       { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)         { m.IDSource = &v }
func (m *Message) SetIssuer(v string)           { m.Issuer = &v }
func (m *Message) SetSecurityDesc(v string)     { m.SecurityDesc = &v }
func (m *Message) SetSide(v string)             { m.Side = v }
func (m *Message) SetOrderQty(v int)            { m.OrderQty = v }
func (m *Message) SetOrdType(v string)          { m.OrdType = &v }
func (m *Message) SetPrice(v float64)           { m.Price = &v }
func (m *Message) SetStopPx(v float64)          { m.StopPx = &v }
func (m *Message) SetCurrency(v string)         { m.Currency = &v }
func (m *Message) SetTimeInForce(v string)      { m.TimeInForce = &v }
func (m *Message) SetExpireTime(v time.Time)    { m.ExpireTime = &v }
func (m *Message) SetExecInst(v string)         { m.ExecInst = &v }
func (m *Message) SetRule80A(v string)          { m.Rule80A = &v }
func (m *Message) SetLastShares(v int)          { m.LastShares = v }
func (m *Message) SetLastPx(v float64)          { m.LastPx = v }
func (m *Message) SetLastMkt(v string)          { m.LastMkt = &v }
func (m *Message) SetLastCapacity(v string)     { m.LastCapacity = &v }
func (m *Message) SetCumQty(v int)              { m.CumQty = v }
func (m *Message) SetAvgPx(v float64)           { m.AvgPx = v }
func (m *Message) SetTradeDate(v string)        { m.TradeDate = &v }
func (m *Message) SetTransactTime(v time.Time)  { m.TransactTime = &v }
func (m *Message) SetReportToExch(v string)     { m.ReportToExch = &v }
func (m *Message) SetCommission(v float64)      { m.Commission = &v }
func (m *Message) SetCommType(v string)         { m.CommType = &v }
func (m *Message) SetNoMiscFees(v []NoMiscFees) { m.NoMiscFees = v }
func (m *Message) SetNetMoney(v float64)        { m.NetMoney = &v }
func (m *Message) SetSettlCurrAmt(v float64)    { m.SettlCurrAmt = &v }
func (m *Message) SetSettlCurrency(v string)    { m.SettlCurrency = &v }
func (m *Message) SetText(v string)             { m.Text = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX40, "8", r
}
