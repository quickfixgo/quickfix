//Package newordermultileg msg type = AB.
package newordermultileg

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp1/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50sp1/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/legordgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/peginstructions"
	"github.com/quickfixgo/quickfix/fix50sp1/preallocmleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a NewOrderMultileg FIX Message
type Message struct {
	FIXMsgType string `fix:"AB"`
	fixt11.Header
	//ClOrdID is a required field for NewOrderMultileg.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NewOrderMultileg.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NewOrderMultileg.
	ClOrdLinkID *string `fix:"583"`
	//Parties is a non-required component for NewOrderMultileg.
	Parties *parties.Parties
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
	//PreAllocMlegGrp is a non-required component for NewOrderMultileg.
	PreAllocMlegGrp *preallocmleggrp.PreAllocMlegGrp
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
	//TrdgSesGrp is a non-required component for NewOrderMultileg.
	TrdgSesGrp *trdgsesgrp.TrdgSesGrp
	//ProcessCode is a non-required field for NewOrderMultileg.
	ProcessCode *string `fix:"81"`
	//Side is a required field for NewOrderMultileg.
	Side string `fix:"54"`
	//Instrument is a non-required component for NewOrderMultileg.
	Instrument *instrument.Instrument
	//UndInstrmtGrp is a non-required component for NewOrderMultileg.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//PrevClosePx is a non-required field for NewOrderMultileg.
	PrevClosePx *float64 `fix:"140"`
	//LegOrdGrp is a non-required component for NewOrderMultileg.
	LegOrdGrp *legordgrp.LegOrdGrp
	//LocateReqd is a non-required field for NewOrderMultileg.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderMultileg.
	TransactTime time.Time `fix:"60"`
	//QtyType is a non-required field for NewOrderMultileg.
	QtyType *int `fix:"854"`
	//OrderQtyData is a non-required component for NewOrderMultileg.
	OrderQtyData *orderqtydata.OrderQtyData
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
	//CommissionData is a non-required component for NewOrderMultileg.
	CommissionData *commissiondata.CommissionData
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
	//PegInstructions is a non-required component for NewOrderMultileg.
	PegInstructions *peginstructions.PegInstructions
	//DiscretionInstructions is a non-required component for NewOrderMultileg.
	DiscretionInstructions *discretioninstructions.DiscretionInstructions
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
	//StrategyParametersGrp is a non-required component for NewOrderMultileg.
	StrategyParametersGrp *strategyparametersgrp.StrategyParametersGrp
	//SwapPoints is a non-required field for NewOrderMultileg.
	SwapPoints *float64 `fix:"1069"`
	//MatchIncrement is a non-required field for NewOrderMultileg.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for NewOrderMultileg.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction is a non-required component for NewOrderMultileg.
	DisplayInstruction *displayinstruction.DisplayInstruction
	//PriceProtectionScope is a non-required field for NewOrderMultileg.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction is a non-required component for NewOrderMultileg.
	TriggeringInstruction *triggeringinstruction.TriggeringInstruction
	//RefOrderID is a non-required field for NewOrderMultileg.
	RefOrderID *string `fix:"1080"`
	//RefOrderIDSource is a non-required field for NewOrderMultileg.
	RefOrderIDSource *string `fix:"1081"`
	//PreTradeAnonymity is a non-required field for NewOrderMultileg.
	PreTradeAnonymity *bool `fix:"1091"`
	//ExDestinationIDSource is a non-required field for NewOrderMultileg.
	ExDestinationIDSource *string `fix:"1133"`
	//MultilegModel is a non-required field for NewOrderMultileg.
	MultilegModel *int `fix:"1377"`
	//MultilegPriceMethod is a non-required field for NewOrderMultileg.
	MultilegPriceMethod *int `fix:"1378"`
	//RiskFreeRate is a non-required field for NewOrderMultileg.
	RiskFreeRate *float64 `fix:"1190"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized NewOrderMultileg instance
func New(clordid string, side string, transacttime time.Time, ordtype string) *Message {
	var m Message
	m.SetClOrdID(clordid)
	m.SetSide(side)
	m.SetTransactTime(transacttime)
	m.SetOrdType(ordtype)
	return &m
}

func (m *Message) SetClOrdID(v string)                                  { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)                         { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)                              { m.ClOrdLinkID = &v }
func (m *Message) SetParties(v parties.Parties)                         { m.Parties = &v }
func (m *Message) SetTradeOriginationDate(v string)                     { m.TradeOriginationDate = &v }
func (m *Message) SetTradeDate(v string)                                { m.TradeDate = &v }
func (m *Message) SetAccount(v string)                                  { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                                { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                                 { m.AccountType = &v }
func (m *Message) SetDayBookingInst(v string)                           { m.DayBookingInst = &v }
func (m *Message) SetBookingUnit(v string)                              { m.BookingUnit = &v }
func (m *Message) SetPreallocMethod(v string)                           { m.PreallocMethod = &v }
func (m *Message) SetAllocID(v string)                                  { m.AllocID = &v }
func (m *Message) SetPreAllocMlegGrp(v preallocmleggrp.PreAllocMlegGrp) { m.PreAllocMlegGrp = &v }
func (m *Message) SetSettlType(v string)                                { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                                { m.SettlDate = &v }
func (m *Message) SetCashMargin(v string)                               { m.CashMargin = &v }
func (m *Message) SetClearingFeeIndicator(v string)                     { m.ClearingFeeIndicator = &v }
func (m *Message) SetHandlInst(v string)                                { m.HandlInst = &v }
func (m *Message) SetExecInst(v string)                                 { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                                  { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                                { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)                            { m.ExDestination = &v }
func (m *Message) SetTrdgSesGrp(v trdgsesgrp.TrdgSesGrp)                { m.TrdgSesGrp = &v }
func (m *Message) SetProcessCode(v string)                              { m.ProcessCode = &v }
func (m *Message) SetSide(v string)                                     { m.Side = v }
func (m *Message) SetInstrument(v instrument.Instrument)                { m.Instrument = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)       { m.UndInstrmtGrp = &v }
func (m *Message) SetPrevClosePx(v float64)                             { m.PrevClosePx = &v }
func (m *Message) SetLegOrdGrp(v legordgrp.LegOrdGrp)                   { m.LegOrdGrp = &v }
func (m *Message) SetLocateReqd(v bool)                                 { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)                          { m.TransactTime = v }
func (m *Message) SetQtyType(v int)                                     { m.QtyType = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData)          { m.OrderQtyData = &v }
func (m *Message) SetOrdType(v string)                                  { m.OrdType = v }
func (m *Message) SetPriceType(v int)                                   { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                                   { m.Price = &v }
func (m *Message) SetStopPx(v float64)                                  { m.StopPx = &v }
func (m *Message) SetCurrency(v string)                                 { m.Currency = &v }
func (m *Message) SetComplianceID(v string)                             { m.ComplianceID = &v }
func (m *Message) SetSolicitedFlag(v bool)                              { m.SolicitedFlag = &v }
func (m *Message) SetIOIID(v string)                                    { m.IOIID = &v }
func (m *Message) SetQuoteID(v string)                                  { m.QuoteID = &v }
func (m *Message) SetTimeInForce(v string)                              { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)                         { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)                               { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)                            { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)                               { m.GTBookingInst = &v }
func (m *Message) SetCommissionData(v commissiondata.CommissionData)    { m.CommissionData = &v }
func (m *Message) SetOrderCapacity(v string)                            { m.OrderCapacity = &v }
func (m *Message) SetOrderRestrictions(v string)                        { m.OrderRestrictions = &v }
func (m *Message) SetCustOrderCapacity(v int)                           { m.CustOrderCapacity = &v }
func (m *Message) SetForexReq(v bool)                                   { m.ForexReq = &v }
func (m *Message) SetSettlCurrency(v string)                            { m.SettlCurrency = &v }
func (m *Message) SetBookingType(v int)                                 { m.BookingType = &v }
func (m *Message) SetText(v string)                                     { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                              { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                              { m.EncodedText = &v }
func (m *Message) SetPositionEffect(v string)                           { m.PositionEffect = &v }
func (m *Message) SetCoveredOrUncovered(v int)                          { m.CoveredOrUncovered = &v }
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
func (m *Message) SetMultiLegRptTypeReq(v int)          { m.MultiLegRptTypeReq = &v }
func (m *Message) SetStrategyParametersGrp(v strategyparametersgrp.StrategyParametersGrp) {
	m.StrategyParametersGrp = &v
}
func (m *Message) SetSwapPoints(v float64)     { m.SwapPoints = &v }
func (m *Message) SetMatchIncrement(v float64) { m.MatchIncrement = &v }
func (m *Message) SetMaxPriceLevels(v int)     { m.MaxPriceLevels = &v }
func (m *Message) SetDisplayInstruction(v displayinstruction.DisplayInstruction) {
	m.DisplayInstruction = &v
}
func (m *Message) SetPriceProtectionScope(v string) { m.PriceProtectionScope = &v }
func (m *Message) SetTriggeringInstruction(v triggeringinstruction.TriggeringInstruction) {
	m.TriggeringInstruction = &v
}
func (m *Message) SetRefOrderID(v string)            { m.RefOrderID = &v }
func (m *Message) SetRefOrderIDSource(v string)      { m.RefOrderIDSource = &v }
func (m *Message) SetPreTradeAnonymity(v bool)       { m.PreTradeAnonymity = &v }
func (m *Message) SetExDestinationIDSource(v string) { m.ExDestinationIDSource = &v }
func (m *Message) SetMultilegModel(v int)            { m.MultilegModel = &v }
func (m *Message) SetMultilegPriceMethod(v int)      { m.MultilegPriceMethod = &v }
func (m *Message) SetRiskFreeRate(v float64)         { m.RiskFreeRate = &v }

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
	return enum.ApplVerID_FIX50SP1, "AB", r
}
