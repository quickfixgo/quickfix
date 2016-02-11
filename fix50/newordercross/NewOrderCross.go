//Package newordercross msg type = s.
package newordercross

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/peginstructions"
	"github.com/quickfixgo/quickfix/fix50/rootparties"
	"github.com/quickfixgo/quickfix/fix50/sidecrossordmodgrp"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a NewOrderCross FIX Message
type Message struct {
	FIXMsgType string `fix:"s"`
	Header     fixt11.Header
	//CrossID is a required field for NewOrderCross.
	CrossID string `fix:"548"`
	//CrossType is a required field for NewOrderCross.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for NewOrderCross.
	CrossPrioritization int `fix:"550"`
	//SideCrossOrdModGrp Component
	SideCrossOrdModGrp sidecrossordmodgrp.Component
	//Instrument Component
	Instrument instrument.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//SettlType is a non-required field for NewOrderCross.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NewOrderCross.
	SettlDate *string `fix:"64"`
	//HandlInst is a non-required field for NewOrderCross.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NewOrderCross.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderCross.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NewOrderCross.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NewOrderCross.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//ProcessCode is a non-required field for NewOrderCross.
	ProcessCode *string `fix:"81"`
	//PrevClosePx is a non-required field for NewOrderCross.
	PrevClosePx *float64 `fix:"140"`
	//LocateReqd is a non-required field for NewOrderCross.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderCross.
	TransactTime time.Time `fix:"60"`
	//Stipulations Component
	Stipulations stipulations.Component
	//OrdType is a required field for NewOrderCross.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for NewOrderCross.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NewOrderCross.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderCross.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Currency is a non-required field for NewOrderCross.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NewOrderCross.
	ComplianceID *string `fix:"376"`
	//IOIID is a non-required field for NewOrderCross.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NewOrderCross.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NewOrderCross.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NewOrderCross.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NewOrderCross.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NewOrderCross.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NewOrderCross.
	GTBookingInst *int `fix:"427"`
	//MaxShow is a non-required field for NewOrderCross.
	MaxShow *float64 `fix:"210"`
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//TargetStrategy is a non-required field for NewOrderCross.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NewOrderCross.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NewOrderCross.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for NewOrderCross.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderCross.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderCross.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for NewOrderCross.
	Designation *string `fix:"494"`
	//StrategyParametersGrp Component
	StrategyParametersGrp strategyparametersgrp.Component
	//TransBkdTime is a non-required field for NewOrderCross.
	TransBkdTime *time.Time `fix:"483"`
	//RootParties Component
	RootParties rootparties.Component
	//MatchIncrement is a non-required field for NewOrderCross.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for NewOrderCross.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	DisplayInstruction displayinstruction.Component
	//PriceProtectionScope is a non-required field for NewOrderCross.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	TriggeringInstruction triggeringinstruction.Component
	//ExDestinationIDSource is a non-required field for NewOrderCross.
	ExDestinationIDSource *string `fix:"1133"`
	Trailer               fixt11.Trailer
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
	return enum.BeginStringFIX50, "s", r
}
