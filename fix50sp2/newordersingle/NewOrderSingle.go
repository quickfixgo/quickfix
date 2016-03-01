//Package newordersingle msg type = D.
package newordersingle

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp2/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50sp2/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/peginstructions"
	"github.com/quickfixgo/quickfix/fix50sp2/preallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp2/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a NewOrderSingle FIX Message
type Message struct {
	FIXMsgType string `fix:"D"`
	Header     fixt11.Header
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
	//TradeDate is a non-required field for NewOrderSingle.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NewOrderSingle.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NewOrderSingle.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NewOrderSingle.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NewOrderSingle.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NewOrderSingle.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NewOrderSingle.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NewOrderSingle.
	AllocID *string `fix:"70"`
	//PreAllocGrp Component
	PreAllocGrp preallocgrp.Component
	//SettlType is a non-required field for NewOrderSingle.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NewOrderSingle.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NewOrderSingle.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NewOrderSingle.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NewOrderSingle.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NewOrderSingle.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderSingle.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NewOrderSingle.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NewOrderSingle.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//ProcessCode is a non-required field for NewOrderSingle.
	ProcessCode *string `fix:"81"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
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
	//QtyType is a non-required field for NewOrderSingle.
	QtyType *int `fix:"854"`
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
	//IOIID is a non-required field for NewOrderSingle.
	IOIID *string `fix:"23"`
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
	//ForexReq is a non-required field for NewOrderSingle.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderSingle.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NewOrderSingle.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NewOrderSingle.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NewOrderSingle.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NewOrderSingle.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for NewOrderSingle.
	SettlDate2 *string `fix:"193"`
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
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//TargetStrategy is a non-required field for NewOrderSingle.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NewOrderSingle.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NewOrderSingle.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for NewOrderSingle.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderSingle.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderSingle.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for NewOrderSingle.
	Designation *string `fix:"494"`
	//StrategyParametersGrp Component
	StrategyParametersGrp strategyparametersgrp.Component
	//ManualOrderIndicator is a non-required field for NewOrderSingle.
	ManualOrderIndicator *bool `fix:"1028"`
	//CustDirectedOrder is a non-required field for NewOrderSingle.
	CustDirectedOrder *bool `fix:"1029"`
	//ReceivedDeptID is a non-required field for NewOrderSingle.
	ReceivedDeptID *string `fix:"1030"`
	//CustOrderHandlingInst is a non-required field for NewOrderSingle.
	CustOrderHandlingInst *string `fix:"1031"`
	//OrderHandlingInstSource is a non-required field for NewOrderSingle.
	OrderHandlingInstSource *int `fix:"1032"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//MatchIncrement is a non-required field for NewOrderSingle.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for NewOrderSingle.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	DisplayInstruction displayinstruction.Component
	//PriceProtectionScope is a non-required field for NewOrderSingle.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	TriggeringInstruction triggeringinstruction.Component
	//PreTradeAnonymity is a non-required field for NewOrderSingle.
	PreTradeAnonymity *bool `fix:"1091"`
	//RefOrderID is a non-required field for NewOrderSingle.
	RefOrderID *string `fix:"1080"`
	//RefOrderIDSource is a non-required field for NewOrderSingle.
	RefOrderIDSource *string `fix:"1081"`
	//ExDestinationIDSource is a non-required field for NewOrderSingle.
	ExDestinationIDSource *string `fix:"1133"`
	Trailer               fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "D", r
}
