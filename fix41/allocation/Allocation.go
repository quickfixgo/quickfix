//Package allocation msg type = J.
package allocation

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
	"time"
)

//NoOrders is a repeating group in Allocation
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoOrders.
	SecondaryOrderID *string `fix:"198"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//WaveNo is a non-required field for NoOrders.
	WaveNo *string `fix:"105"`
}

//NoExecs is a repeating group in Allocation
type NoExecs struct {
	//LastShares is a non-required field for NoExecs.
	LastShares *int `fix:"32"`
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
	//LastPx is a non-required field for NoExecs.
	LastPx *float64 `fix:"31"`
	//LastCapacity is a non-required field for NoExecs.
	LastCapacity *string `fix:"29"`
}

//NoAllocs is a repeating group in Allocation
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocShares is a required field for NoAllocs.
	AllocShares int `fix:"80"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//BrokerOfCredit is a non-required field for NoAllocs.
	BrokerOfCredit *string `fix:"92"`
	//NotifyBrokerOfCredit is a non-required field for NoAllocs.
	NotifyBrokerOfCredit *string `fix:"208"`
	//AllocHandlInst is a non-required field for NoAllocs.
	AllocHandlInst *int `fix:"209"`
	//AllocText is a non-required field for NoAllocs.
	AllocText *string `fix:"161"`
	//ExecBroker is a non-required field for NoAllocs.
	ExecBroker *string `fix:"76"`
	//ClientID is a non-required field for NoAllocs.
	ClientID *string `fix:"109"`
	//Commission is a non-required field for NoAllocs.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NoAllocs.
	CommType *string `fix:"13"`
	//AllocAvgPx is a non-required field for NoAllocs.
	AllocAvgPx *float64 `fix:"153"`
	//AllocNetMoney is a non-required field for NoAllocs.
	AllocNetMoney *float64 `fix:"154"`
	//SettlCurrAmt is a non-required field for NoAllocs.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for NoAllocs.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for NoAllocs.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoAllocs.
	SettlCurrFxRateCalc *string `fix:"156"`
	//AccruedInterestAmt is a non-required field for NoAllocs.
	AccruedInterestAmt *float64 `fix:"159"`
	//SettlInstMode is a non-required field for NoAllocs.
	SettlInstMode *string `fix:"160"`
	//NoMiscFees is a non-required field for NoAllocs.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
}

//NoMiscFees is a repeating group in NoAllocs
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
}

//Message is a Allocation FIX Message
type Message struct {
	FIXMsgType string `fix:"J"`
	Header     fix41.Header
	//AllocID is a required field for Allocation.
	AllocID string `fix:"70"`
	//AllocTransType is a required field for Allocation.
	AllocTransType string `fix:"71"`
	//RefAllocID is a non-required field for Allocation.
	RefAllocID *string `fix:"72"`
	//AllocLinkID is a non-required field for Allocation.
	AllocLinkID *string `fix:"196"`
	//AllocLinkType is a non-required field for Allocation.
	AllocLinkType *int `fix:"197"`
	//NoOrders is a non-required field for Allocation.
	NoOrders []NoOrders `fix:"73,omitempty"`
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
	//SecurityType is a non-required field for Allocation.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for Allocation.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for Allocation.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for Allocation.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for Allocation.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for Allocation.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for Allocation.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for Allocation.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for Allocation.
	SecurityDesc *string `fix:"107"`
	//Shares is a required field for Allocation.
	Shares int `fix:"53"`
	//LastMkt is a non-required field for Allocation.
	LastMkt *string `fix:"30"`
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
	//OpenClose is a non-required field for Allocation.
	OpenClose *string `fix:"77"`
	//Text is a non-required field for Allocation.
	Text *string `fix:"58"`
	//NumDaysInterest is a non-required field for Allocation.
	NumDaysInterest *int `fix:"157"`
	//AccruedInterestRate is a non-required field for Allocation.
	AccruedInterestRate *float64 `fix:"158"`
	//NoAllocs is a non-required field for Allocation.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	Trailer  fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX41, "J", r
}
