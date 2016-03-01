//Package ordercancelreplacerequest msg type = G.
package ordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
	"time"
)

//NoAllocs is a repeating group in OrderCancelReplaceRequest
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NoTradingSessions is a repeating group in OrderCancelReplaceRequest
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//Message is a OrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"G"`
	Header     fix43.Header
	//OrderID is a non-required field for OrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for OrderCancelReplaceRequest.
	TradeOriginationDate *string `fix:"229"`
	//OrigClOrdID is a required field for OrderCancelReplaceRequest.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for OrderCancelReplaceRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderCancelReplaceRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for OrderCancelReplaceRequest.
	ClOrdLinkID *string `fix:"583"`
	//ListID is a non-required field for OrderCancelReplaceRequest.
	ListID *string `fix:"66"`
	//OrigOrdModTime is a non-required field for OrderCancelReplaceRequest.
	OrigOrdModTime *time.Time `fix:"586"`
	//Account is a non-required field for OrderCancelReplaceRequest.
	Account *string `fix:"1"`
	//AccountType is a non-required field for OrderCancelReplaceRequest.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for OrderCancelReplaceRequest.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for OrderCancelReplaceRequest.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for OrderCancelReplaceRequest.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for OrderCancelReplaceRequest.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for OrderCancelReplaceRequest.
	FutSettDate *string `fix:"64"`
	//CashMargin is a non-required field for OrderCancelReplaceRequest.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for OrderCancelReplaceRequest.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a required field for OrderCancelReplaceRequest.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for OrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for OrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for OrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for OrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for OrderCancelReplaceRequest.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//Instrument Component
	Instrument instrument.Component
	//Side is a required field for OrderCancelReplaceRequest.
	Side string `fix:"54"`
	//TransactTime is a required field for OrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//QuantityType is a non-required field for OrderCancelReplaceRequest.
	QuantityType *int `fix:"465"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a required field for OrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for OrderCancelReplaceRequest.
	PriceType *int `fix:"423"`
	//Price is a non-required field for OrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for OrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//PegDifference is a non-required field for OrderCancelReplaceRequest.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for OrderCancelReplaceRequest.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for OrderCancelReplaceRequest.
	DiscretionOffset *float64 `fix:"389"`
	//ComplianceID is a non-required field for OrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for OrderCancelReplaceRequest.
	SolicitedFlag *bool `fix:"377"`
	//Currency is a non-required field for OrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//TimeInForce is a non-required field for OrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for OrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for OrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for OrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for OrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for OrderCancelReplaceRequest.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for OrderCancelReplaceRequest.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for OrderCancelReplaceRequest.
	CustOrderCapacity *int `fix:"582"`
	//Rule80A is a non-required field for OrderCancelReplaceRequest.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for OrderCancelReplaceRequest.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for OrderCancelReplaceRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelReplaceRequest.
	EncodedText *string `fix:"355"`
	//FutSettDate2 is a non-required field for OrderCancelReplaceRequest.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for OrderCancelReplaceRequest.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for OrderCancelReplaceRequest.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for OrderCancelReplaceRequest.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for OrderCancelReplaceRequest.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for OrderCancelReplaceRequest.
	MaxShow *float64 `fix:"210"`
	//LocateReqd is a non-required field for OrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//CancellationRights is a non-required field for OrderCancelReplaceRequest.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for OrderCancelReplaceRequest.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for OrderCancelReplaceRequest.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for OrderCancelReplaceRequest.
	Designation *string `fix:"494"`
	//AccruedInterestRate is a non-required field for OrderCancelReplaceRequest.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for OrderCancelReplaceRequest.
	AccruedInterestAmt *float64 `fix:"159"`
	//NetMoney is a non-required field for OrderCancelReplaceRequest.
	NetMoney *float64 `fix:"118"`
	Trailer  fix43.Trailer
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
	return enum.BeginStringFIX43, "G", r
}
