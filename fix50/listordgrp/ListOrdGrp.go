package listordgrp

import (
	"github.com/quickfixgo/quickfix/fix50/commissiondata"
	"github.com/quickfixgo/quickfix/fix50/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/peginstructions"
	"github.com/quickfixgo/quickfix/fix50/preallocgrp"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"time"
)

func New(noorders []NoOrders) *ListOrdGrp {
	var m ListOrdGrp
	m.SetNoOrders(noorders)
	return &m
}

//NoOrders is a repeating group in ListOrdGrp
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListSeqNo is a required field for NoOrders.
	ListSeqNo int `fix:"67"`
	//ClOrdLinkID is a non-required field for NoOrders.
	ClOrdLinkID *string `fix:"583"`
	//SettlInstMode is a non-required field for NoOrders.
	SettlInstMode *string `fix:"160"`
	//Parties is a non-required component for NoOrders.
	Parties *parties.Parties
	//TradeOriginationDate is a non-required field for NoOrders.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NoOrders.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NoOrders.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoOrders.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoOrders.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoOrders.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoOrders.
	BookingUnit *string `fix:"590"`
	//AllocID is a non-required field for NoOrders.
	AllocID *string `fix:"70"`
	//PreallocMethod is a non-required field for NoOrders.
	PreallocMethod *string `fix:"591"`
	//PreAllocGrp is a non-required component for NoOrders.
	PreAllocGrp *preallocgrp.PreAllocGrp
	//SettlType is a non-required field for NoOrders.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoOrders.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NoOrders.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NoOrders.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NoOrders.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NoOrders.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NoOrders.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NoOrders.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NoOrders.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp is a non-required component for NoOrders.
	TrdgSesGrp *trdgsesgrp.TrdgSesGrp
	//ProcessCode is a non-required field for NoOrders.
	ProcessCode *string `fix:"81"`
	//Instrument is a required component for NoOrders.
	instrument.Instrument
	//UndInstrmtGrp is a non-required component for NoOrders.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//PrevClosePx is a non-required field for NoOrders.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NoOrders.
	Side string `fix:"54"`
	//SideValueInd is a non-required field for NoOrders.
	SideValueInd *int `fix:"401"`
	//LocateReqd is a non-required field for NoOrders.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a non-required field for NoOrders.
	TransactTime *time.Time `fix:"60"`
	//Stipulations is a non-required component for NoOrders.
	Stipulations *stipulations.Stipulations
	//QtyType is a non-required field for NoOrders.
	QtyType *int `fix:"854"`
	//OrderQtyData is a required component for NoOrders.
	orderqtydata.OrderQtyData
	//OrdType is a non-required field for NoOrders.
	OrdType *string `fix:"40"`
	//PriceType is a non-required field for NoOrders.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoOrders.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NoOrders.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData is a non-required component for NoOrders.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for NoOrders.
	YieldData *yielddata.YieldData
	//Currency is a non-required field for NoOrders.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoOrders.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoOrders.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for NoOrders.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NoOrders.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NoOrders.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NoOrders.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NoOrders.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoOrders.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NoOrders.
	GTBookingInst *int `fix:"427"`
	//CommissionData is a non-required component for NoOrders.
	CommissionData *commissiondata.CommissionData
	//OrderCapacity is a non-required field for NoOrders.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoOrders.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoOrders.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NoOrders.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoOrders.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NoOrders.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for NoOrders.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoOrders.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for NoOrders.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for NoOrders.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoOrders.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NoOrders.
	MaxShow *float64 `fix:"210"`
	//PegInstructions is a non-required component for NoOrders.
	PegInstructions *peginstructions.PegInstructions
	//DiscretionInstructions is a non-required component for NoOrders.
	DiscretionInstructions *discretioninstructions.DiscretionInstructions
	//TargetStrategy is a non-required field for NoOrders.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NoOrders.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NoOrders.
	ParticipationRate *float64 `fix:"849"`
	//Designation is a non-required field for NoOrders.
	Designation *string `fix:"494"`
	//StrategyParametersGrp is a non-required component for NoOrders.
	StrategyParametersGrp *strategyparametersgrp.StrategyParametersGrp
	//MatchIncrement is a non-required field for NoOrders.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for NoOrders.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction is a non-required component for NoOrders.
	DisplayInstruction *displayinstruction.DisplayInstruction
	//PriceProtectionScope is a non-required field for NoOrders.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction is a non-required component for NoOrders.
	TriggeringInstruction *triggeringinstruction.TriggeringInstruction
	//RefOrderID is a non-required field for NoOrders.
	RefOrderID *string `fix:"1080"`
	//RefOrderIDSource is a non-required field for NoOrders.
	RefOrderIDSource *string `fix:"1081"`
	//PreTradeAnonymity is a non-required field for NoOrders.
	PreTradeAnonymity *bool `fix:"1091"`
	//ExDestinationIDSource is a non-required field for NoOrders.
	ExDestinationIDSource *string `fix:"1133"`
}

//ListOrdGrp is a fix50 Component
type ListOrdGrp struct {
	//NoOrders is a required field for ListOrdGrp.
	NoOrders []NoOrders `fix:"73"`
}

func (m *ListOrdGrp) SetNoOrders(v []NoOrders) { m.NoOrders = v }
