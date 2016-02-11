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
