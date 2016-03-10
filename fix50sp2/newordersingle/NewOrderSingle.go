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
	fixt11.Header
	//ClOrdID is a required field for NewOrderSingle.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NewOrderSingle.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NewOrderSingle.
	ClOrdLinkID *string `fix:"583"`
	//Parties Component
	parties.Parties
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
	preallocgrp.PreAllocGrp
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
	trdgsesgrp.TrdgSesGrp
	//ProcessCode is a non-required field for NewOrderSingle.
	ProcessCode *string `fix:"81"`
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//PrevClosePx is a non-required field for NewOrderSingle.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NewOrderSingle.
	Side string `fix:"54"`
	//LocateReqd is a non-required field for NewOrderSingle.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderSingle.
	TransactTime time.Time `fix:"60"`
	//Stipulations Component
	stipulations.Stipulations
	//QtyType is a non-required field for NewOrderSingle.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//OrdType is a required field for NewOrderSingle.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for NewOrderSingle.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NewOrderSingle.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderSingle.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData Component
	yielddata.YieldData
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
	commissiondata.CommissionData
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
	peginstructions.PegInstructions
	//DiscretionInstructions Component
	discretioninstructions.DiscretionInstructions
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
	strategyparametersgrp.StrategyParametersGrp
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
	trdregtimestamps.TrdRegTimestamps
	//MatchIncrement is a non-required field for NewOrderSingle.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for NewOrderSingle.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	displayinstruction.DisplayInstruction
	//PriceProtectionScope is a non-required field for NewOrderSingle.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	triggeringinstruction.TriggeringInstruction
	//PreTradeAnonymity is a non-required field for NewOrderSingle.
	PreTradeAnonymity *bool `fix:"1091"`
	//RefOrderID is a non-required field for NewOrderSingle.
	RefOrderID *string `fix:"1080"`
	//RefOrderIDSource is a non-required field for NewOrderSingle.
	RefOrderIDSource *string `fix:"1081"`
	//ExDestinationIDSource is a non-required field for NewOrderSingle.
	ExDestinationIDSource *string `fix:"1133"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetClOrdID(v string)                  { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)         { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)              { m.ClOrdLinkID = &v }
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
func (m *Message) SetPrevClosePx(v float64)             { m.PrevClosePx = &v }
func (m *Message) SetSide(v string)                     { m.Side = v }
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
func (m *Message) SetSettlDate2(v string)               { m.SettlDate2 = &v }
func (m *Message) SetOrderQty2(v float64)               { m.OrderQty2 = &v }
func (m *Message) SetPrice2(v float64)                  { m.Price2 = &v }
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
func (m *Message) SetManualOrderIndicator(v bool)       { m.ManualOrderIndicator = &v }
func (m *Message) SetCustDirectedOrder(v bool)          { m.CustDirectedOrder = &v }
func (m *Message) SetReceivedDeptID(v string)           { m.ReceivedDeptID = &v }
func (m *Message) SetCustOrderHandlingInst(v string)    { m.CustOrderHandlingInst = &v }
func (m *Message) SetOrderHandlingInstSource(v int)     { m.OrderHandlingInstSource = &v }
func (m *Message) SetMatchIncrement(v float64)          { m.MatchIncrement = &v }
func (m *Message) SetMaxPriceLevels(v int)              { m.MaxPriceLevels = &v }
func (m *Message) SetPriceProtectionScope(v string)     { m.PriceProtectionScope = &v }
func (m *Message) SetPreTradeAnonymity(v bool)          { m.PreTradeAnonymity = &v }
func (m *Message) SetRefOrderID(v string)               { m.RefOrderID = &v }
func (m *Message) SetRefOrderIDSource(v string)         { m.RefOrderIDSource = &v }
func (m *Message) SetExDestinationIDSource(v string)    { m.ExDestinationIDSource = &v }

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
