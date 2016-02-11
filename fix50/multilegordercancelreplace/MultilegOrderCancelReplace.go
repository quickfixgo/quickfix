//Package multilegordercancelreplace msg type = AC.
package multilegordercancelreplace

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/commissiondata"
	"github.com/quickfixgo/quickfix/fix50/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/legordgrp"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/peginstructions"
	"github.com/quickfixgo/quickfix/fix50/preallocmleggrp"
	"github.com/quickfixgo/quickfix/fix50/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a MultilegOrderCancelReplace FIX Message
type Message struct {
	FIXMsgType string `fix:"AC"`
	Header     fixt11.Header
	//OrderID is a non-required field for MultilegOrderCancelReplace.
	OrderID *string `fix:"37"`
	//OrigClOrdID is a required field for MultilegOrderCancelReplace.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for MultilegOrderCancelReplace.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplace.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for MultilegOrderCancelReplace.
	ClOrdLinkID *string `fix:"583"`
	//OrigOrdModTime is a non-required field for MultilegOrderCancelReplace.
	OrigOrdModTime *time.Time `fix:"586"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for MultilegOrderCancelReplace.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for MultilegOrderCancelReplace.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for MultilegOrderCancelReplace.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for MultilegOrderCancelReplace.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for MultilegOrderCancelReplace.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for MultilegOrderCancelReplace.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for MultilegOrderCancelReplace.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for MultilegOrderCancelReplace.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for MultilegOrderCancelReplace.
	AllocID *string `fix:"70"`
	//PreAllocMlegGrp Component
	PreAllocMlegGrp preallocmleggrp.Component
	//SettlType is a non-required field for MultilegOrderCancelReplace.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for MultilegOrderCancelReplace.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for MultilegOrderCancelReplace.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplace.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for MultilegOrderCancelReplace.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for MultilegOrderCancelReplace.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for MultilegOrderCancelReplace.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for MultilegOrderCancelReplace.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for MultilegOrderCancelReplace.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//ProcessCode is a non-required field for MultilegOrderCancelReplace.
	ProcessCode *string `fix:"81"`
	//Side is a required field for MultilegOrderCancelReplace.
	Side string `fix:"54"`
	//Instrument Component
	Instrument instrument.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//PrevClosePx is a non-required field for MultilegOrderCancelReplace.
	PrevClosePx *float64 `fix:"140"`
	//LegOrdGrp Component
	LegOrdGrp legordgrp.Component
	//LocateReqd is a non-required field for MultilegOrderCancelReplace.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for MultilegOrderCancelReplace.
	TransactTime time.Time `fix:"60"`
	//QtyType is a non-required field for MultilegOrderCancelReplace.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a required field for MultilegOrderCancelReplace.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for MultilegOrderCancelReplace.
	PriceType *int `fix:"423"`
	//Price is a non-required field for MultilegOrderCancelReplace.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for MultilegOrderCancelReplace.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for MultilegOrderCancelReplace.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for MultilegOrderCancelReplace.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for MultilegOrderCancelReplace.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for MultilegOrderCancelReplace.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for MultilegOrderCancelReplace.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for MultilegOrderCancelReplace.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for MultilegOrderCancelReplace.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for MultilegOrderCancelReplace.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for MultilegOrderCancelReplace.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for MultilegOrderCancelReplace.
	GTBookingInst *int `fix:"427"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for MultilegOrderCancelReplace.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for MultilegOrderCancelReplace.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for MultilegOrderCancelReplace.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for MultilegOrderCancelReplace.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for MultilegOrderCancelReplace.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for MultilegOrderCancelReplace.
	BookingType *int `fix:"775"`
	//Text is a non-required field for MultilegOrderCancelReplace.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for MultilegOrderCancelReplace.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for MultilegOrderCancelReplace.
	EncodedText *string `fix:"355"`
	//PositionEffect is a non-required field for MultilegOrderCancelReplace.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplace.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for MultilegOrderCancelReplace.
	MaxShow *float64 `fix:"210"`
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//TargetStrategy is a non-required field for MultilegOrderCancelReplace.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for MultilegOrderCancelReplace.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for MultilegOrderCancelReplace.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for MultilegOrderCancelReplace.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplace.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for MultilegOrderCancelReplace.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for MultilegOrderCancelReplace.
	Designation *string `fix:"494"`
	//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplace.
	MultiLegRptTypeReq *int `fix:"563"`
	//StrategyParametersGrp Component
	StrategyParametersGrp strategyparametersgrp.Component
	//MatchIncrement is a non-required field for MultilegOrderCancelReplace.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for MultilegOrderCancelReplace.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	DisplayInstruction displayinstruction.Component
	//PriceProtectionScope is a non-required field for MultilegOrderCancelReplace.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	TriggeringInstruction triggeringinstruction.Component
	//PreTradeAnonymity is a non-required field for MultilegOrderCancelReplace.
	PreTradeAnonymity *bool `fix:"1091"`
	//ExDestinationIDSource is a non-required field for MultilegOrderCancelReplace.
	ExDestinationIDSource *string `fix:"1133"`
	//SwapPoints is a non-required field for MultilegOrderCancelReplace.
	SwapPoints *float64 `fix:"1069"`
	Trailer    fixt11.Trailer
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
	return enum.BeginStringFIX50, "AC", r
}
