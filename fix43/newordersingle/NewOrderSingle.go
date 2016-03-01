//Package newordersingle msg type = D.
package newordersingle

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
	"github.com/quickfixgo/quickfix/fix43/stipulations"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
	"time"
)

//NoAllocs is a repeating group in NewOrderSingle
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

//NoTradingSessions is a repeating group in NewOrderSingle
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//Message is a NewOrderSingle FIX Message
type Message struct {
	FIXMsgType string `fix:"D"`
	Header     fix43.Header
	//ClOrdID is a required field for NewOrderSingle.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NewOrderSingle.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NewOrderSingle.
	ClOrdLinkID *string `fix:"583"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for NewOrderSingle.
	TradeOriginationDate *string `fix:"229"`
	//Account is a non-required field for NewOrderSingle.
	Account *string `fix:"1"`
	//AccountType is a non-required field for NewOrderSingle.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NewOrderSingle.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NewOrderSingle.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NewOrderSingle.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for NewOrderSingle.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for NewOrderSingle.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NewOrderSingle.
	FutSettDate *string `fix:"64"`
	//CashMargin is a non-required field for NewOrderSingle.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NewOrderSingle.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a required field for NewOrderSingle.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for NewOrderSingle.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderSingle.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NewOrderSingle.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NewOrderSingle.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NewOrderSingle.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NewOrderSingle.
	ProcessCode *string `fix:"81"`
	//Instrument Component
	Instrument instrument.Component
	//PrevClosePx is a non-required field for NewOrderSingle.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NewOrderSingle.
	Side string `fix:"54"`
	//LocateReqd is a non-required field for NewOrderSingle.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderSingle.
	TransactTime time.Time `fix:"60"`
	//Stipulations Component
	Stipulations stipulations.Component
	//QuantityType is a non-required field for NewOrderSingle.
	QuantityType *int `fix:"465"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a required field for NewOrderSingle.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for NewOrderSingle.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NewOrderSingle.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderSingle.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Currency is a non-required field for NewOrderSingle.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NewOrderSingle.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NewOrderSingle.
	SolicitedFlag *bool `fix:"377"`
	//IOIid is a non-required field for NewOrderSingle.
	IOIid *string `fix:"23"`
	//QuoteID is a non-required field for NewOrderSingle.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NewOrderSingle.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NewOrderSingle.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NewOrderSingle.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NewOrderSingle.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NewOrderSingle.
	GTBookingInst *int `fix:"427"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for NewOrderSingle.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NewOrderSingle.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NewOrderSingle.
	CustOrderCapacity *int `fix:"582"`
	//Rule80A is a non-required field for NewOrderSingle.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for NewOrderSingle.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderSingle.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NewOrderSingle.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NewOrderSingle.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NewOrderSingle.
	EncodedText *string `fix:"355"`
	//FutSettDate2 is a non-required field for NewOrderSingle.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NewOrderSingle.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for NewOrderSingle.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for NewOrderSingle.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NewOrderSingle.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NewOrderSingle.
	MaxShow *float64 `fix:"210"`
	//PegDifference is a non-required field for NewOrderSingle.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for NewOrderSingle.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for NewOrderSingle.
	DiscretionOffset *float64 `fix:"389"`
	//CancellationRights is a non-required field for NewOrderSingle.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderSingle.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderSingle.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for NewOrderSingle.
	Designation *string `fix:"494"`
	//AccruedInterestRate is a non-required field for NewOrderSingle.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for NewOrderSingle.
	AccruedInterestAmt *float64 `fix:"159"`
	//NetMoney is a non-required field for NewOrderSingle.
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
	return enum.BeginStringFIX43, "D", r
}
