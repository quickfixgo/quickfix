//Package ordercancelreplacerequest msg type = G.
package ordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp1/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50sp1/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50sp1/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/peginstructions"
	"github.com/quickfixgo/quickfix/fix50sp1/preallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp1/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a OrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"G"`
	Header     fixt11.Header
	//OrderID is a non-required field for OrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for OrderCancelReplaceRequest.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for OrderCancelReplaceRequest.
	TradeDate *string `fix:"75"`
	//OrigClOrdID is a non-required field for OrderCancelReplaceRequest.
	OrigClOrdID *string `fix:"41"`
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
	//AcctIDSource is a non-required field for OrderCancelReplaceRequest.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for OrderCancelReplaceRequest.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for OrderCancelReplaceRequest.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for OrderCancelReplaceRequest.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for OrderCancelReplaceRequest.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for OrderCancelReplaceRequest.
	AllocID *string `fix:"70"`
	//PreAllocGrp Component
	PreAllocGrp preallocgrp.Component
	//SettlType is a non-required field for OrderCancelReplaceRequest.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for OrderCancelReplaceRequest.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for OrderCancelReplaceRequest.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for OrderCancelReplaceRequest.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for OrderCancelReplaceRequest.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for OrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for OrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for OrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for OrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Side is a required field for OrderCancelReplaceRequest.
	Side string `fix:"54"`
	//TransactTime is a required field for OrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//QtyType is a non-required field for OrderCancelReplaceRequest.
	QtyType *int `fix:"854"`
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
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//TargetStrategy is a non-required field for OrderCancelReplaceRequest.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for OrderCancelReplaceRequest.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for OrderCancelReplaceRequest.
	ParticipationRate *float64 `fix:"849"`
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
	//ForexReq is a non-required field for OrderCancelReplaceRequest.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for OrderCancelReplaceRequest.
	BookingType *int `fix:"775"`
	//Text is a non-required field for OrderCancelReplaceRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelReplaceRequest.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for OrderCancelReplaceRequest.
	SettlDate2 *string `fix:"193"`
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
	//StrategyParametersGrp Component
	StrategyParametersGrp strategyparametersgrp.Component
	//ManualOrderIndicator is a non-required field for OrderCancelReplaceRequest.
	ManualOrderIndicator *bool `fix:"1028"`
	//CustDirectedOrder is a non-required field for OrderCancelReplaceRequest.
	CustDirectedOrder *bool `fix:"1029"`
	//ReceivedDeptID is a non-required field for OrderCancelReplaceRequest.
	ReceivedDeptID *string `fix:"1030"`
	//CustOrderHandlingInst is a non-required field for OrderCancelReplaceRequest.
	CustOrderHandlingInst *string `fix:"1031"`
	//OrderHandlingInstSource is a non-required field for OrderCancelReplaceRequest.
	OrderHandlingInstSource *int `fix:"1032"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//MatchIncrement is a non-required field for OrderCancelReplaceRequest.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for OrderCancelReplaceRequest.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	DisplayInstruction displayinstruction.Component
	//PriceProtectionScope is a non-required field for OrderCancelReplaceRequest.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	TriggeringInstruction triggeringinstruction.Component
	//PreTradeAnonymity is a non-required field for OrderCancelReplaceRequest.
	PreTradeAnonymity *bool `fix:"1091"`
	//ExDestinationIDSource is a non-required field for OrderCancelReplaceRequest.
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
	return enum.ApplVerID_FIX50SP1, "G", r
}
