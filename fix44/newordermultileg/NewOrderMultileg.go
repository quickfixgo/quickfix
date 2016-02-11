//Package newordermultileg msg type = AB.
package newordermultileg

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"github.com/quickfixgo/quickfix/fix44/nestedparties3"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/peginstructions"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoAllocs is a repeating group in NewOrderMultileg
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties3 Component
	NestedParties3 nestedparties3.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NoTradingSessions is a repeating group in NewOrderMultileg
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//NoUnderlyings is a repeating group in NewOrderMultileg
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in NewOrderMultileg
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
	//NoLegAllocs is a non-required field for NoLegs.
	NoLegAllocs []NoLegAllocs `fix:"670,omitempty"`
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegPrice is a non-required field for NoLegs.
	LegPrice *float64 `fix:"566"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
}

//NoLegAllocs is a repeating group in NoLegs
type NoLegAllocs struct {
	//LegAllocAccount is a non-required field for NoLegAllocs.
	LegAllocAccount *string `fix:"671"`
	//LegIndividualAllocID is a non-required field for NoLegAllocs.
	LegIndividualAllocID *string `fix:"672"`
	//NestedParties2 Component
	NestedParties2 nestedparties2.Component
	//LegAllocQty is a non-required field for NoLegAllocs.
	LegAllocQty *float64 `fix:"673"`
	//LegAllocAcctIDSource is a non-required field for NoLegAllocs.
	LegAllocAcctIDSource *string `fix:"674"`
	//LegSettlCurrency is a non-required field for NoLegAllocs.
	LegSettlCurrency *string `fix:"675"`
}

//Message is a NewOrderMultileg FIX Message
type Message struct {
	FIXMsgType string `fix:"AB"`
	Header     fix44.Header
	//ClOrdID is a required field for NewOrderMultileg.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NewOrderMultileg.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NewOrderMultileg.
	ClOrdLinkID *string `fix:"583"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for NewOrderMultileg.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NewOrderMultileg.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NewOrderMultileg.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NewOrderMultileg.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NewOrderMultileg.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NewOrderMultileg.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NewOrderMultileg.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NewOrderMultileg.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NewOrderMultileg.
	AllocID *string `fix:"70"`
	//NoAllocs is a non-required field for NewOrderMultileg.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlType is a non-required field for NewOrderMultileg.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NewOrderMultileg.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NewOrderMultileg.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NewOrderMultileg.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NewOrderMultileg.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NewOrderMultileg.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderMultileg.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NewOrderMultileg.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NewOrderMultileg.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NewOrderMultileg.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NewOrderMultileg.
	ProcessCode *string `fix:"81"`
	//Side is a required field for NewOrderMultileg.
	Side string `fix:"54"`
	//Instrument Component
	Instrument instrument.Component
	//NoUnderlyings is a non-required field for NewOrderMultileg.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PrevClosePx is a non-required field for NewOrderMultileg.
	PrevClosePx *float64 `fix:"140"`
	//NoLegs is a required field for NewOrderMultileg.
	NoLegs []NoLegs `fix:"555"`
	//LocateReqd is a non-required field for NewOrderMultileg.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderMultileg.
	TransactTime time.Time `fix:"60"`
	//QtyType is a non-required field for NewOrderMultileg.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a required field for NewOrderMultileg.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for NewOrderMultileg.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NewOrderMultileg.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderMultileg.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for NewOrderMultileg.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NewOrderMultileg.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NewOrderMultileg.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for NewOrderMultileg.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NewOrderMultileg.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NewOrderMultileg.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NewOrderMultileg.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NewOrderMultileg.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NewOrderMultileg.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NewOrderMultileg.
	GTBookingInst *int `fix:"427"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for NewOrderMultileg.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NewOrderMultileg.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NewOrderMultileg.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NewOrderMultileg.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderMultileg.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NewOrderMultileg.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NewOrderMultileg.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NewOrderMultileg.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NewOrderMultileg.
	EncodedText *string `fix:"355"`
	//PositionEffect is a non-required field for NewOrderMultileg.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NewOrderMultileg.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NewOrderMultileg.
	MaxShow *float64 `fix:"210"`
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//TargetStrategy is a non-required field for NewOrderMultileg.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NewOrderMultileg.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NewOrderMultileg.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for NewOrderMultileg.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderMultileg.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderMultileg.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for NewOrderMultileg.
	Designation *string `fix:"494"`
	//MultiLegRptTypeReq is a non-required field for NewOrderMultileg.
	MultiLegRptTypeReq *int `fix:"563"`
	Trailer            fix44.Trailer
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
	return enum.BeginStringFIX44, "AB", r
}
