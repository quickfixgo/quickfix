//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/peginstructions"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoOrders is a repeating group in NewOrderList
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListSeqNo is a required field for NoOrders.
	ListSeqNo int `fix:"67"`
	//ClOrdLinkID is a non-required field for NoOrders.
	ClOrdLinkID *string `fix:"583"`
	//SettlInstMode is a non-required field for NoOrders.
	SettlInstMode *string `fix:"160"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for NoOrders.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NoOrders.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NoOrders.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoOrders.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoOrders.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoOrders.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoOrders.
	BookingUnit *string `fix:"590"`
	//AllocID is a non-required field for NoOrders.
	AllocID *string `fix:"70"`
	//PreallocMethod is a non-required field for NoOrders.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for NoOrders.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlType is a non-required field for NoOrders.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoOrders.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NoOrders.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NoOrders.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NoOrders.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NoOrders.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NoOrders.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NoOrders.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NoOrders.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NoOrders.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NoOrders.
	ProcessCode *string `fix:"81"`
	//Instrument Component
	Instrument instrument.Component
	//NoUnderlyings is a non-required field for NoOrders.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PrevClosePx is a non-required field for NoOrders.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NoOrders.
	Side string `fix:"54"`
	//SideValueInd is a non-required field for NoOrders.
	SideValueInd *int `fix:"401"`
	//LocateReqd is a non-required field for NoOrders.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a non-required field for NoOrders.
	TransactTime *time.Time `fix:"60"`
	//Stipulations Component
	Stipulations stipulations.Component
	//QtyType is a non-required field for NoOrders.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a non-required field for NoOrders.
	OrdType *string `fix:"40"`
	//PriceType is a non-required field for NoOrders.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoOrders.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NoOrders.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Currency is a non-required field for NoOrders.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoOrders.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoOrders.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for NoOrders.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NoOrders.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NoOrders.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NoOrders.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NoOrders.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoOrders.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NoOrders.
	GTBookingInst *int `fix:"427"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for NoOrders.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoOrders.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoOrders.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NoOrders.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoOrders.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NoOrders.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for NoOrders.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoOrders.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for NoOrders.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for NoOrders.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoOrders.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NoOrders.
	MaxShow *float64 `fix:"210"`
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//TargetStrategy is a non-required field for NoOrders.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NoOrders.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NoOrders.
	ParticipationRate *float64 `fix:"849"`
	//Designation is a non-required field for NoOrders.
	Designation *string `fix:"494"`
}

//NoAllocs is a repeating group in NoOrders
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NoTradingSessions is a repeating group in NoOrders
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//NoUnderlyings is a repeating group in NoOrders
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//Message is a NewOrderList FIX Message
type Message struct {
	FIXMsgType string `fix:"E"`
	Header     fix44.Header
	//ListID is a required field for NewOrderList.
	ListID string `fix:"66"`
	//BidID is a non-required field for NewOrderList.
	BidID *string `fix:"390"`
	//ClientBidID is a non-required field for NewOrderList.
	ClientBidID *string `fix:"391"`
	//ProgRptReqs is a non-required field for NewOrderList.
	ProgRptReqs *int `fix:"414"`
	//BidType is a required field for NewOrderList.
	BidType int `fix:"394"`
	//ProgPeriodInterval is a non-required field for NewOrderList.
	ProgPeriodInterval *int `fix:"415"`
	//CancellationRights is a non-required field for NewOrderList.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderList.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderList.
	RegistID *string `fix:"513"`
	//ListExecInstType is a non-required field for NewOrderList.
	ListExecInstType *string `fix:"433"`
	//ListExecInst is a non-required field for NewOrderList.
	ListExecInst *string `fix:"69"`
	//EncodedListExecInstLen is a non-required field for NewOrderList.
	EncodedListExecInstLen *int `fix:"352"`
	//EncodedListExecInst is a non-required field for NewOrderList.
	EncodedListExecInst *string `fix:"353"`
	//AllowableOneSidednessPct is a non-required field for NewOrderList.
	AllowableOneSidednessPct *float64 `fix:"765"`
	//AllowableOneSidednessValue is a non-required field for NewOrderList.
	AllowableOneSidednessValue *float64 `fix:"766"`
	//AllowableOneSidednessCurr is a non-required field for NewOrderList.
	AllowableOneSidednessCurr *string `fix:"767"`
	//TotNoOrders is a required field for NewOrderList.
	TotNoOrders int `fix:"68"`
	//LastFragment is a non-required field for NewOrderList.
	LastFragment *bool `fix:"893"`
	//NoOrders is a required field for NewOrderList.
	NoOrders []NoOrders `fix:"73"`
	Trailer  fix44.Trailer
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
	return enum.BeginStringFIX44, "E", r
}
