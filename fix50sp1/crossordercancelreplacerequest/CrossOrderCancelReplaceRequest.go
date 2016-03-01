//Package crossordercancelreplacerequest msg type = t.
package crossordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50sp1/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/peginstructions"
	"github.com/quickfixgo/quickfix/fix50sp1/rootparties"
	"github.com/quickfixgo/quickfix/fix50sp1/sidecrossordmodgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CrossOrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"t"`
	Header     fixt11.Header
	//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//CrossID is a required field for CrossOrderCancelReplaceRequest.
	CrossID string `fix:"548"`
	//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
	OrigCrossID string `fix:"551"`
	//CrossType is a required field for CrossOrderCancelReplaceRequest.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
	CrossPrioritization int `fix:"550"`
	//SideCrossOrdModGrp Component
	SideCrossOrdModGrp sidecrossordmodgrp.Component
	//Instrument Component
	Instrument instrument.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//SettlType is a non-required field for CrossOrderCancelReplaceRequest.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for CrossOrderCancelReplaceRequest.
	SettlDate *string `fix:"64"`
	//HandlInst is a non-required field for CrossOrderCancelReplaceRequest.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
	ProcessCode *string `fix:"81"`
	//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
	PrevClosePx *float64 `fix:"140"`
	//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for CrossOrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//Stipulations Component
	Stipulations stipulations.Component
	//OrdType is a required field for CrossOrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
	PriceType *int `fix:"423"`
	//Price is a non-required field for CrossOrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Currency is a non-required field for CrossOrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//IOIID is a non-required field for CrossOrderCancelReplaceRequest.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
	MaxShow *float64 `fix:"210"`
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//TargetStrategy is a non-required field for CrossOrderCancelReplaceRequest.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for CrossOrderCancelReplaceRequest.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for CrossOrderCancelReplaceRequest.
	Designation *string `fix:"494"`
	//StrategyParametersGrp Component
	StrategyParametersGrp strategyparametersgrp.Component
	//HostCrossID is a non-required field for CrossOrderCancelReplaceRequest.
	HostCrossID *string `fix:"961"`
	//TransBkdTime is a non-required field for CrossOrderCancelReplaceRequest.
	TransBkdTime *time.Time `fix:"483"`
	//RootParties Component
	RootParties rootparties.Component
	//MatchIncrement is a non-required field for CrossOrderCancelReplaceRequest.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for CrossOrderCancelReplaceRequest.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	DisplayInstruction displayinstruction.Component
	//PriceProtectionScope is a non-required field for CrossOrderCancelReplaceRequest.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	TriggeringInstruction triggeringinstruction.Component
	//ExDestinationIDSource is a non-required field for CrossOrderCancelReplaceRequest.
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
	return enum.ApplVerID_FIX50SP1, "t", r
}
