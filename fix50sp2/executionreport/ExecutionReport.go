//Package executionreport msg type = 8.
package executionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp2/contamtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/contragrp"
	"github.com/quickfixgo/quickfix/fix50sp2/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50sp2/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50sp2/fillsgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtlegexecgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/peginstructions"
	"github.com/quickfixgo/quickfix/fix50sp2/preallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/ratesource"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp2/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a ExecutionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"8"`
	Header     fixt11.Header
	//OrderID is a required field for ExecutionReport.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for ExecutionReport.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for ExecutionReport.
	SecondaryClOrdID *string `fix:"526"`
	//SecondaryExecID is a non-required field for ExecutionReport.
	SecondaryExecID *string `fix:"527"`
	//ClOrdID is a non-required field for ExecutionReport.
	ClOrdID *string `fix:"11"`
	//OrigClOrdID is a non-required field for ExecutionReport.
	OrigClOrdID *string `fix:"41"`
	//ClOrdLinkID is a non-required field for ExecutionReport.
	ClOrdLinkID *string `fix:"583"`
	//QuoteRespID is a non-required field for ExecutionReport.
	QuoteRespID *string `fix:"693"`
	//OrdStatusReqID is a non-required field for ExecutionReport.
	OrdStatusReqID *string `fix:"790"`
	//MassStatusReqID is a non-required field for ExecutionReport.
	MassStatusReqID *string `fix:"584"`
	//TotNumReports is a non-required field for ExecutionReport.
	TotNumReports *int `fix:"911"`
	//LastRptRequested is a non-required field for ExecutionReport.
	LastRptRequested *bool `fix:"912"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for ExecutionReport.
	TradeOriginationDate *string `fix:"229"`
	//ContraGrp Component
	ContraGrp contragrp.Component
	//ListID is a non-required field for ExecutionReport.
	ListID *string `fix:"66"`
	//CrossID is a non-required field for ExecutionReport.
	CrossID *string `fix:"548"`
	//OrigCrossID is a non-required field for ExecutionReport.
	OrigCrossID *string `fix:"551"`
	//CrossType is a non-required field for ExecutionReport.
	CrossType *int `fix:"549"`
	//ExecID is a required field for ExecutionReport.
	ExecID string `fix:"17"`
	//ExecRefID is a non-required field for ExecutionReport.
	ExecRefID *string `fix:"19"`
	//ExecType is a required field for ExecutionReport.
	ExecType string `fix:"150"`
	//OrdStatus is a required field for ExecutionReport.
	OrdStatus string `fix:"39"`
	//WorkingIndicator is a non-required field for ExecutionReport.
	WorkingIndicator *bool `fix:"636"`
	//OrdRejReason is a non-required field for ExecutionReport.
	OrdRejReason *int `fix:"103"`
	//ExecRestatementReason is a non-required field for ExecutionReport.
	ExecRestatementReason *int `fix:"378"`
	//Account is a non-required field for ExecutionReport.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for ExecutionReport.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for ExecutionReport.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for ExecutionReport.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for ExecutionReport.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for ExecutionReport.
	PreallocMethod *string `fix:"591"`
	//SettlType is a non-required field for ExecutionReport.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for ExecutionReport.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for ExecutionReport.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for ExecutionReport.
	ClearingFeeIndicator *string `fix:"635"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Side is a required field for ExecutionReport.
	Side string `fix:"54"`
	//Stipulations Component
	Stipulations stipulations.Component
	//QtyType is a non-required field for ExecutionReport.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a non-required field for ExecutionReport.
	OrdType *string `fix:"40"`
	//PriceType is a non-required field for ExecutionReport.
	PriceType *int `fix:"423"`
	//Price is a non-required field for ExecutionReport.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for ExecutionReport.
	StopPx *float64 `fix:"99"`
	//PegInstructions Component
	PegInstructions peginstructions.Component
	//DiscretionInstructions Component
	DiscretionInstructions discretioninstructions.Component
	//PeggedPrice is a non-required field for ExecutionReport.
	PeggedPrice *float64 `fix:"839"`
	//DiscretionPrice is a non-required field for ExecutionReport.
	DiscretionPrice *float64 `fix:"845"`
	//TargetStrategy is a non-required field for ExecutionReport.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for ExecutionReport.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for ExecutionReport.
	ParticipationRate *float64 `fix:"849"`
	//TargetStrategyPerformance is a non-required field for ExecutionReport.
	TargetStrategyPerformance *float64 `fix:"850"`
	//Currency is a non-required field for ExecutionReport.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for ExecutionReport.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for ExecutionReport.
	SolicitedFlag *bool `fix:"377"`
	//TimeInForce is a non-required field for ExecutionReport.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for ExecutionReport.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for ExecutionReport.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for ExecutionReport.
	ExpireTime *time.Time `fix:"126"`
	//ExecInst is a non-required field for ExecutionReport.
	ExecInst *string `fix:"18"`
	//OrderCapacity is a non-required field for ExecutionReport.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for ExecutionReport.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for ExecutionReport.
	CustOrderCapacity *int `fix:"582"`
	//LastQty is a non-required field for ExecutionReport.
	LastQty *float64 `fix:"32"`
	//UnderlyingLastQty is a non-required field for ExecutionReport.
	UnderlyingLastQty *float64 `fix:"652"`
	//LastPx is a non-required field for ExecutionReport.
	LastPx *float64 `fix:"31"`
	//UnderlyingLastPx is a non-required field for ExecutionReport.
	UnderlyingLastPx *float64 `fix:"651"`
	//LastParPx is a non-required field for ExecutionReport.
	LastParPx *float64 `fix:"669"`
	//LastSpotRate is a non-required field for ExecutionReport.
	LastSpotRate *float64 `fix:"194"`
	//LastForwardPoints is a non-required field for ExecutionReport.
	LastForwardPoints *float64 `fix:"195"`
	//LastMkt is a non-required field for ExecutionReport.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for ExecutionReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for ExecutionReport.
	TradingSessionSubID *string `fix:"625"`
	//TimeBracket is a non-required field for ExecutionReport.
	TimeBracket *string `fix:"943"`
	//LastCapacity is a non-required field for ExecutionReport.
	LastCapacity *string `fix:"29"`
	//LeavesQty is a required field for ExecutionReport.
	LeavesQty float64 `fix:"151"`
	//CumQty is a required field for ExecutionReport.
	CumQty float64 `fix:"14"`
	//AvgPx is a non-required field for ExecutionReport.
	AvgPx *float64 `fix:"6"`
	//DayOrderQty is a non-required field for ExecutionReport.
	DayOrderQty *float64 `fix:"424"`
	//DayCumQty is a non-required field for ExecutionReport.
	DayCumQty *float64 `fix:"425"`
	//DayAvgPx is a non-required field for ExecutionReport.
	DayAvgPx *float64 `fix:"426"`
	//GTBookingInst is a non-required field for ExecutionReport.
	GTBookingInst *int `fix:"427"`
	//TradeDate is a non-required field for ExecutionReport.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for ExecutionReport.
	TransactTime *time.Time `fix:"60"`
	//ReportToExch is a non-required field for ExecutionReport.
	ReportToExch *bool `fix:"113"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//GrossTradeAmt is a non-required field for ExecutionReport.
	GrossTradeAmt *float64 `fix:"381"`
	//NumDaysInterest is a non-required field for ExecutionReport.
	NumDaysInterest *int `fix:"157"`
	//ExDate is a non-required field for ExecutionReport.
	ExDate *string `fix:"230"`
	//AccruedInterestRate is a non-required field for ExecutionReport.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for ExecutionReport.
	AccruedInterestAmt *float64 `fix:"159"`
	//InterestAtMaturity is a non-required field for ExecutionReport.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for ExecutionReport.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for ExecutionReport.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for ExecutionReport.
	EndCash *float64 `fix:"922"`
	//TradedFlatSwitch is a non-required field for ExecutionReport.
	TradedFlatSwitch *bool `fix:"258"`
	//BasisFeatureDate is a non-required field for ExecutionReport.
	BasisFeatureDate *string `fix:"259"`
	//BasisFeaturePrice is a non-required field for ExecutionReport.
	BasisFeaturePrice *float64 `fix:"260"`
	//Concession is a non-required field for ExecutionReport.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for ExecutionReport.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for ExecutionReport.
	NetMoney *float64 `fix:"118"`
	//SettlCurrAmt is a non-required field for ExecutionReport.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for ExecutionReport.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for ExecutionReport.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
	SettlCurrFxRateCalc *string `fix:"156"`
	//HandlInst is a non-required field for ExecutionReport.
	HandlInst *string `fix:"21"`
	//MinQty is a non-required field for ExecutionReport.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for ExecutionReport.
	MaxFloor *float64 `fix:"111"`
	//PositionEffect is a non-required field for ExecutionReport.
	PositionEffect *string `fix:"77"`
	//MaxShow is a non-required field for ExecutionReport.
	MaxShow *float64 `fix:"210"`
	//BookingType is a non-required field for ExecutionReport.
	BookingType *int `fix:"775"`
	//Text is a non-required field for ExecutionReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ExecutionReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ExecutionReport.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for ExecutionReport.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for ExecutionReport.
	OrderQty2 *float64 `fix:"192"`
	//LastForwardPoints2 is a non-required field for ExecutionReport.
	LastForwardPoints2 *float64 `fix:"641"`
	//MultiLegReportingType is a non-required field for ExecutionReport.
	MultiLegReportingType *string `fix:"442"`
	//CancellationRights is a non-required field for ExecutionReport.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for ExecutionReport.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for ExecutionReport.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for ExecutionReport.
	Designation *string `fix:"494"`
	//TransBkdTime is a non-required field for ExecutionReport.
	TransBkdTime *time.Time `fix:"483"`
	//ExecValuationPoint is a non-required field for ExecutionReport.
	ExecValuationPoint *time.Time `fix:"515"`
	//ExecPriceType is a non-required field for ExecutionReport.
	ExecPriceType *string `fix:"484"`
	//ExecPriceAdjustment is a non-required field for ExecutionReport.
	ExecPriceAdjustment *float64 `fix:"485"`
	//PriorityIndicator is a non-required field for ExecutionReport.
	PriorityIndicator *int `fix:"638"`
	//PriceImprovement is a non-required field for ExecutionReport.
	PriceImprovement *float64 `fix:"639"`
	//LastLiquidityInd is a non-required field for ExecutionReport.
	LastLiquidityInd *int `fix:"851"`
	//ContAmtGrp Component
	ContAmtGrp contamtgrp.Component
	//InstrmtLegExecGrp Component
	InstrmtLegExecGrp instrmtlegexecgrp.Component
	//CopyMsgIndicator is a non-required field for ExecutionReport.
	CopyMsgIndicator *bool `fix:"797"`
	//MiscFeesGrp Component
	MiscFeesGrp miscfeesgrp.Component
	//StrategyParametersGrp Component
	StrategyParametersGrp strategyparametersgrp.Component
	//HostCrossID is a non-required field for ExecutionReport.
	HostCrossID *string `fix:"961"`
	//ManualOrderIndicator is a non-required field for ExecutionReport.
	ManualOrderIndicator *bool `fix:"1028"`
	//CustDirectedOrder is a non-required field for ExecutionReport.
	CustDirectedOrder *bool `fix:"1029"`
	//ReceivedDeptID is a non-required field for ExecutionReport.
	ReceivedDeptID *string `fix:"1030"`
	//CustOrderHandlingInst is a non-required field for ExecutionReport.
	CustOrderHandlingInst *string `fix:"1031"`
	//OrderHandlingInstSource is a non-required field for ExecutionReport.
	OrderHandlingInstSource *int `fix:"1032"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//AggressorIndicator is a non-required field for ExecutionReport.
	AggressorIndicator *bool `fix:"1057"`
	//CalculatedCcyLastQty is a non-required field for ExecutionReport.
	CalculatedCcyLastQty *float64 `fix:"1056"`
	//LastSwapPoints is a non-required field for ExecutionReport.
	LastSwapPoints *float64 `fix:"1071"`
	//MatchType is a non-required field for ExecutionReport.
	MatchType *string `fix:"574"`
	//OrderCategory is a non-required field for ExecutionReport.
	OrderCategory *string `fix:"1115"`
	//LotType is a non-required field for ExecutionReport.
	LotType *string `fix:"1093"`
	//PriceProtectionScope is a non-required field for ExecutionReport.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	TriggeringInstruction triggeringinstruction.Component
	//PeggedRefPrice is a non-required field for ExecutionReport.
	PeggedRefPrice *float64 `fix:"1095"`
	//PreTradeAnonymity is a non-required field for ExecutionReport.
	PreTradeAnonymity *bool `fix:"1091"`
	//MatchIncrement is a non-required field for ExecutionReport.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for ExecutionReport.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	DisplayInstruction displayinstruction.Component
	//Volatility is a non-required field for ExecutionReport.
	Volatility *float64 `fix:"1188"`
	//TimeToExpiration is a non-required field for ExecutionReport.
	TimeToExpiration *float64 `fix:"1189"`
	//RiskFreeRate is a non-required field for ExecutionReport.
	RiskFreeRate *float64 `fix:"1190"`
	//PriceDelta is a non-required field for ExecutionReport.
	PriceDelta *float64 `fix:"811"`
	//TrdMatchID is a non-required field for ExecutionReport.
	TrdMatchID *string `fix:"880"`
	//AllocID is a non-required field for ExecutionReport.
	AllocID *string `fix:"70"`
	//PreAllocGrp Component
	PreAllocGrp preallocgrp.Component
	//TotNoFills is a non-required field for ExecutionReport.
	TotNoFills *int `fix:"1361"`
	//LastFragment is a non-required field for ExecutionReport.
	LastFragment *bool `fix:"893"`
	//FillsGrp Component
	FillsGrp fillsgrp.Component
	//DividendYield is a non-required field for ExecutionReport.
	DividendYield *float64 `fix:"1380"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	//RateSource Component
	RateSource ratesource.Component
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
	return enum.ApplVerID_FIX50SP2, "8", r
}
