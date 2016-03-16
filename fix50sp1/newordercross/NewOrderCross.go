//Package newordercross msg type = s.
package newordercross

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

//Message is a NewOrderCross FIX Message
type Message struct {
	FIXMsgType string `fix:"s"`
	fixt11.Header
	//CrossID is a required field for NewOrderCross.
	CrossID string `fix:"548"`
	//CrossType is a required field for NewOrderCross.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for NewOrderCross.
	CrossPrioritization int `fix:"550"`
	//SideCrossOrdModGrp is a required component for NewOrderCross.
	sidecrossordmodgrp.SideCrossOrdModGrp
	//Instrument is a required component for NewOrderCross.
	instrument.Instrument
	//UndInstrmtGrp is a non-required component for NewOrderCross.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for NewOrderCross.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
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
	//TrdgSesGrp is a non-required component for NewOrderCross.
	TrdgSesGrp *trdgsesgrp.TrdgSesGrp
	//ProcessCode is a non-required field for NewOrderCross.
	ProcessCode *string `fix:"81"`
	//PrevClosePx is a non-required field for NewOrderCross.
	PrevClosePx *float64 `fix:"140"`
	//LocateReqd is a non-required field for NewOrderCross.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderCross.
	TransactTime time.Time `fix:"60"`
	//Stipulations is a non-required component for NewOrderCross.
	Stipulations *stipulations.Stipulations
	//OrdType is a required field for NewOrderCross.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for NewOrderCross.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NewOrderCross.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderCross.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData is a non-required component for NewOrderCross.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for NewOrderCross.
	YieldData *yielddata.YieldData
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
	//PegInstructions is a non-required component for NewOrderCross.
	PegInstructions *peginstructions.PegInstructions
	//DiscretionInstructions is a non-required component for NewOrderCross.
	DiscretionInstructions *discretioninstructions.DiscretionInstructions
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
	//StrategyParametersGrp is a non-required component for NewOrderCross.
	StrategyParametersGrp *strategyparametersgrp.StrategyParametersGrp
	//TransBkdTime is a non-required field for NewOrderCross.
	TransBkdTime *time.Time `fix:"483"`
	//RootParties is a non-required component for NewOrderCross.
	RootParties *rootparties.RootParties
	//MatchIncrement is a non-required field for NewOrderCross.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for NewOrderCross.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction is a non-required component for NewOrderCross.
	DisplayInstruction *displayinstruction.DisplayInstruction
	//PriceProtectionScope is a non-required field for NewOrderCross.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction is a non-required component for NewOrderCross.
	TriggeringInstruction *triggeringinstruction.TriggeringInstruction
	//ExDestinationIDSource is a non-required field for NewOrderCross.
	ExDestinationIDSource *string `fix:"1133"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetCrossID(v string)          { m.CrossID = v }
func (m *Message) SetCrossType(v int)           { m.CrossType = v }
func (m *Message) SetCrossPrioritization(v int) { m.CrossPrioritization = v }
func (m *Message) SetSideCrossOrdModGrp(v sidecrossordmodgrp.SideCrossOrdModGrp) {
	m.SideCrossOrdModGrp = v
}
func (m *Message) SetInstrument(v instrument.Instrument)          { m.Instrument = v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *Message) SetSettlType(v string)                          { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                          { m.SettlDate = &v }
func (m *Message) SetHandlInst(v string)                          { m.HandlInst = &v }
func (m *Message) SetExecInst(v string)                           { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                            { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                          { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)                      { m.ExDestination = &v }
func (m *Message) SetTrdgSesGrp(v trdgsesgrp.TrdgSesGrp)          { m.TrdgSesGrp = &v }
func (m *Message) SetProcessCode(v string)                        { m.ProcessCode = &v }
func (m *Message) SetPrevClosePx(v float64)                       { m.PrevClosePx = &v }
func (m *Message) SetLocateReqd(v bool)                           { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)                    { m.TransactTime = v }
func (m *Message) SetStipulations(v stipulations.Stipulations)    { m.Stipulations = &v }
func (m *Message) SetOrdType(v string)                            { m.OrdType = v }
func (m *Message) SetPriceType(v int)                             { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                             { m.Price = &v }
func (m *Message) SetStopPx(v float64)                            { m.StopPx = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetYieldData(v yielddata.YieldData)                   { m.YieldData = &v }
func (m *Message) SetCurrency(v string)                                 { m.Currency = &v }
func (m *Message) SetComplianceID(v string)                             { m.ComplianceID = &v }
func (m *Message) SetIOIID(v string)                                    { m.IOIID = &v }
func (m *Message) SetQuoteID(v string)                                  { m.QuoteID = &v }
func (m *Message) SetTimeInForce(v string)                              { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)                         { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)                               { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)                            { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)                               { m.GTBookingInst = &v }
func (m *Message) SetMaxShow(v float64)                                 { m.MaxShow = &v }
func (m *Message) SetPegInstructions(v peginstructions.PegInstructions) { m.PegInstructions = &v }
func (m *Message) SetDiscretionInstructions(v discretioninstructions.DiscretionInstructions) {
	m.DiscretionInstructions = &v
}
func (m *Message) SetTargetStrategy(v int)              { m.TargetStrategy = &v }
func (m *Message) SetTargetStrategyParameters(v string) { m.TargetStrategyParameters = &v }
func (m *Message) SetParticipationRate(v float64)       { m.ParticipationRate = &v }
func (m *Message) SetCancellationRights(v string)       { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)    { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                 { m.RegistID = &v }
func (m *Message) SetDesignation(v string)              { m.Designation = &v }
func (m *Message) SetStrategyParametersGrp(v strategyparametersgrp.StrategyParametersGrp) {
	m.StrategyParametersGrp = &v
}
func (m *Message) SetTransBkdTime(v time.Time)              { m.TransBkdTime = &v }
func (m *Message) SetRootParties(v rootparties.RootParties) { m.RootParties = &v }
func (m *Message) SetMatchIncrement(v float64)              { m.MatchIncrement = &v }
func (m *Message) SetMaxPriceLevels(v int)                  { m.MaxPriceLevels = &v }
func (m *Message) SetDisplayInstruction(v displayinstruction.DisplayInstruction) {
	m.DisplayInstruction = &v
}
func (m *Message) SetPriceProtectionScope(v string) { m.PriceProtectionScope = &v }
func (m *Message) SetTriggeringInstruction(v triggeringinstruction.TriggeringInstruction) {
	m.TriggeringInstruction = &v
}
func (m *Message) SetExDestinationIDSource(v string) { m.ExDestinationIDSource = &v }

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
	return enum.ApplVerID_FIX50SP1, "s", r
}
