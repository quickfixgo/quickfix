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
	fixt11.Header
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
	parties.Parties
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
	preallocmleggrp.PreAllocMlegGrp
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
	trdgsesgrp.TrdgSesGrp
	//ProcessCode is a non-required field for MultilegOrderCancelReplace.
	ProcessCode *string `fix:"81"`
	//Side is a required field for MultilegOrderCancelReplace.
	Side string `fix:"54"`
	//Instrument Component
	instrument.Instrument
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//PrevClosePx is a non-required field for MultilegOrderCancelReplace.
	PrevClosePx *float64 `fix:"140"`
	//LegOrdGrp Component
	legordgrp.LegOrdGrp
	//LocateReqd is a non-required field for MultilegOrderCancelReplace.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for MultilegOrderCancelReplace.
	TransactTime time.Time `fix:"60"`
	//QtyType is a non-required field for MultilegOrderCancelReplace.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
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
	commissiondata.CommissionData
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
	peginstructions.PegInstructions
	//DiscretionInstructions Component
	discretioninstructions.DiscretionInstructions
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
	strategyparametersgrp.StrategyParametersGrp
	//MatchIncrement is a non-required field for MultilegOrderCancelReplace.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for MultilegOrderCancelReplace.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	displayinstruction.DisplayInstruction
	//PriceProtectionScope is a non-required field for MultilegOrderCancelReplace.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	triggeringinstruction.TriggeringInstruction
	//PreTradeAnonymity is a non-required field for MultilegOrderCancelReplace.
	PreTradeAnonymity *bool `fix:"1091"`
	//ExDestinationIDSource is a non-required field for MultilegOrderCancelReplace.
	ExDestinationIDSource *string `fix:"1133"`
	//SwapPoints is a non-required field for MultilegOrderCancelReplace.
	SwapPoints *float64 `fix:"1069"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)                  { m.OrderID = &v }
func (m *Message) SetOrigClOrdID(v string)              { m.OrigClOrdID = v }
func (m *Message) SetClOrdID(v string)                  { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)         { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)              { m.ClOrdLinkID = &v }
func (m *Message) SetOrigOrdModTime(v time.Time)        { m.OrigOrdModTime = &v }
func (m *Message) SetTradeOriginationDate(v string)     { m.TradeOriginationDate = &v }
func (m *Message) SetTradeDate(v string)                { m.TradeDate = &v }
func (m *Message) SetAccount(v string)                  { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                 { m.AccountType = &v }
func (m *Message) SetDayBookingInst(v string)           { m.DayBookingInst = &v }
func (m *Message) SetBookingUnit(v string)              { m.BookingUnit = &v }
func (m *Message) SetPreallocMethod(v string)           { m.PreallocMethod = &v }
func (m *Message) SetAllocID(v string)                  { m.AllocID = &v }
func (m *Message) SetSettlType(v string)                { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                { m.SettlDate = &v }
func (m *Message) SetCashMargin(v string)               { m.CashMargin = &v }
func (m *Message) SetClearingFeeIndicator(v string)     { m.ClearingFeeIndicator = &v }
func (m *Message) SetHandlInst(v string)                { m.HandlInst = &v }
func (m *Message) SetExecInst(v string)                 { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                  { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)            { m.ExDestination = &v }
func (m *Message) SetProcessCode(v string)              { m.ProcessCode = &v }
func (m *Message) SetSide(v string)                     { m.Side = v }
func (m *Message) SetPrevClosePx(v float64)             { m.PrevClosePx = &v }
func (m *Message) SetLocateReqd(v bool)                 { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)          { m.TransactTime = v }
func (m *Message) SetQtyType(v int)                     { m.QtyType = &v }
func (m *Message) SetOrdType(v string)                  { m.OrdType = v }
func (m *Message) SetPriceType(v int)                   { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                   { m.Price = &v }
func (m *Message) SetStopPx(v float64)                  { m.StopPx = &v }
func (m *Message) SetCurrency(v string)                 { m.Currency = &v }
func (m *Message) SetComplianceID(v string)             { m.ComplianceID = &v }
func (m *Message) SetSolicitedFlag(v bool)              { m.SolicitedFlag = &v }
func (m *Message) SetIOIID(v string)                    { m.IOIID = &v }
func (m *Message) SetQuoteID(v string)                  { m.QuoteID = &v }
func (m *Message) SetTimeInForce(v string)              { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)         { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)               { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)            { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)               { m.GTBookingInst = &v }
func (m *Message) SetOrderCapacity(v string)            { m.OrderCapacity = &v }
func (m *Message) SetOrderRestrictions(v string)        { m.OrderRestrictions = &v }
func (m *Message) SetCustOrderCapacity(v int)           { m.CustOrderCapacity = &v }
func (m *Message) SetForexReq(v bool)                   { m.ForexReq = &v }
func (m *Message) SetSettlCurrency(v string)            { m.SettlCurrency = &v }
func (m *Message) SetBookingType(v int)                 { m.BookingType = &v }
func (m *Message) SetText(v string)                     { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)              { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)              { m.EncodedText = &v }
func (m *Message) SetPositionEffect(v string)           { m.PositionEffect = &v }
func (m *Message) SetCoveredOrUncovered(v int)          { m.CoveredOrUncovered = &v }
func (m *Message) SetMaxShow(v float64)                 { m.MaxShow = &v }
func (m *Message) SetTargetStrategy(v int)              { m.TargetStrategy = &v }
func (m *Message) SetTargetStrategyParameters(v string) { m.TargetStrategyParameters = &v }
func (m *Message) SetParticipationRate(v float64)       { m.ParticipationRate = &v }
func (m *Message) SetCancellationRights(v string)       { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)    { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                 { m.RegistID = &v }
func (m *Message) SetDesignation(v string)              { m.Designation = &v }
func (m *Message) SetMultiLegRptTypeReq(v int)          { m.MultiLegRptTypeReq = &v }
func (m *Message) SetMatchIncrement(v float64)          { m.MatchIncrement = &v }
func (m *Message) SetMaxPriceLevels(v int)              { m.MaxPriceLevels = &v }
func (m *Message) SetPriceProtectionScope(v string)     { m.PriceProtectionScope = &v }
func (m *Message) SetPreTradeAnonymity(v bool)          { m.PreTradeAnonymity = &v }
func (m *Message) SetExDestinationIDSource(v string)    { m.ExDestinationIDSource = &v }
func (m *Message) SetSwapPoints(v float64)              { m.SwapPoints = &v }

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
	return enum.BeginStringFIX50, "AC", r
}
