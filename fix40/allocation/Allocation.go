//Package allocation msg type = J.
package allocation

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//NoOrders is a repeating group in Allocation
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//WaveNo is a non-required field for NoOrders.
	WaveNo *string `fix:"105"`
}

//NoExecs is a repeating group in Allocation
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *int `fix:"17"`
	//LastShares is a non-required field for NoExecs.
	LastShares *int `fix:"32"`
	//LastPx is a non-required field for NoExecs.
	LastPx *float64 `fix:"31"`
	//LastMkt is a non-required field for NoExecs.
	LastMkt *string `fix:"30"`
}

//NoMiscFees is a repeating group in Allocation
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
}

//NoAllocs is a repeating group in Allocation
type NoAllocs struct {
	//AllocAccount is a required field for NoAllocs.
	AllocAccount string `fix:"79"`
	//AllocShares is a required field for NoAllocs.
	AllocShares int `fix:"80"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//ExecBroker is a non-required field for NoAllocs.
	ExecBroker *string `fix:"76"`
	//ClientID is a non-required field for NoAllocs.
	ClientID *string `fix:"109"`
	//Commission is a non-required field for NoAllocs.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NoAllocs.
	CommType *string `fix:"13"`
	//NoDlvyInst is a non-required field for NoAllocs.
	NoDlvyInst *int `fix:"85"`
	//BrokerOfCredit is a non-required field for NoAllocs.
	BrokerOfCredit *string `fix:"92"`
	//DlvyInst is a non-required field for NoAllocs.
	DlvyInst *string `fix:"86"`
}

//Message is a Allocation FIX Message
type Message struct {
	FIXMsgType string `fix:"J"`
	Header     fix40.Header
	//AllocID is a required field for Allocation.
	AllocID int `fix:"70"`
	//AllocTransType is a required field for Allocation.
	AllocTransType string `fix:"71"`
	//RefAllocID is a non-required field for Allocation.
	RefAllocID *int `fix:"72"`
	//NoOrders is a required field for Allocation.
	NoOrders []NoOrders `fix:"73"`
	//NoExecs is a non-required field for Allocation.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//Side is a required field for Allocation.
	Side string `fix:"54"`
	//Symbol is a required field for Allocation.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for Allocation.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Allocation.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for Allocation.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for Allocation.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for Allocation.
	SecurityDesc *string `fix:"107"`
	//Shares is a required field for Allocation.
	Shares int `fix:"53"`
	//AvgPx is a required field for Allocation.
	AvgPx float64 `fix:"6"`
	//Currency is a non-required field for Allocation.
	Currency *string `fix:"15"`
	//AvgPrxPrecision is a non-required field for Allocation.
	AvgPrxPrecision *int `fix:"74"`
	//TradeDate is a required field for Allocation.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for Allocation.
	TransactTime *time.Time `fix:"60"`
	//SettlmntTyp is a non-required field for Allocation.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for Allocation.
	FutSettDate *string `fix:"64"`
	//NetMoney is a non-required field for Allocation.
	NetMoney *float64 `fix:"118"`
	//NoMiscFees is a non-required field for Allocation.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
	//SettlCurrAmt is a non-required field for Allocation.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for Allocation.
	SettlCurrency *string `fix:"120"`
	//OpenClose is a non-required field for Allocation.
	OpenClose *string `fix:"77"`
	//Text is a non-required field for Allocation.
	Text *string `fix:"58"`
	//NoAllocs is a required field for Allocation.
	NoAllocs []NoAllocs `fix:"78"`
	Trailer  fix40.Trailer
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
	return enum.BeginStringFIX40, "J", r
}
