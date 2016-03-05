package trdcaprptsidegrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/clrinstgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp2/contamtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/settldetails"
	"github.com/quickfixgo/quickfix/fix50sp2/sidetrdregts"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/tradereportorderdetail"
	"github.com/quickfixgo/quickfix/fix50sp2/trdallocgrp"
)

//NoSides is a repeating group in TrdCapRptSideGrp
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for NoSides.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoSides.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoSides.
	AccountType *int `fix:"581"`
	//ProcessCode is a non-required field for NoSides.
	ProcessCode *string `fix:"81"`
	//OddLot is a non-required field for NoSides.
	OddLot *bool `fix:"575"`
	//ClrInstGrp Component
	clrinstgrp.ClrInstGrp
	//TradeInputSource is a non-required field for NoSides.
	TradeInputSource *string `fix:"578"`
	//TradeInputDevice is a non-required field for NoSides.
	TradeInputDevice *string `fix:"579"`
	//ComplianceID is a non-required field for NoSides.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoSides.
	SolicitedFlag *bool `fix:"377"`
	//CustOrderCapacity is a non-required field for NoSides.
	CustOrderCapacity *int `fix:"582"`
	//TradingSessionID is a non-required field for NoSides.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoSides.
	TradingSessionSubID *string `fix:"625"`
	//TimeBracket is a non-required field for NoSides.
	TimeBracket *string `fix:"943"`
	//CommissionData Component
	commissiondata.CommissionData
	//NumDaysInterest is a non-required field for NoSides.
	NumDaysInterest *int `fix:"157"`
	//ExDate is a non-required field for NoSides.
	ExDate *string `fix:"230"`
	//AccruedInterestRate is a non-required field for NoSides.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for NoSides.
	AccruedInterestAmt *float64 `fix:"159"`
	//InterestAtMaturity is a non-required field for NoSides.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for NoSides.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for NoSides.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for NoSides.
	EndCash *float64 `fix:"922"`
	//Concession is a non-required field for NoSides.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for NoSides.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for NoSides.
	NetMoney *float64 `fix:"118"`
	//SettlCurrAmt is a non-required field for NoSides.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrFxRate is a non-required field for NoSides.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoSides.
	SettlCurrFxRateCalc *string `fix:"156"`
	//PositionEffect is a non-required field for NoSides.
	PositionEffect *string `fix:"77"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
	//SideMultiLegReportingType is a non-required field for NoSides.
	SideMultiLegReportingType *int `fix:"752"`
	//ContAmtGrp Component
	contamtgrp.ContAmtGrp
	//Stipulations Component
	stipulations.Stipulations
	//MiscFeesGrp Component
	miscfeesgrp.MiscFeesGrp
	//ExchangeRule is a non-required field for NoSides.
	ExchangeRule *string `fix:"825"`
	//TradeAllocIndicator is a non-required field for NoSides.
	TradeAllocIndicator *int `fix:"826"`
	//PreallocMethod is a non-required field for NoSides.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NoSides.
	AllocID *string `fix:"70"`
	//TrdAllocGrp Component
	trdallocgrp.TrdAllocGrp
	//SideLastQty is a non-required field for NoSides.
	SideLastQty *int `fix:"1009"`
	//SideTradeReportID is a non-required field for NoSides.
	SideTradeReportID *string `fix:"1005"`
	//SideFillStationCd is a non-required field for NoSides.
	SideFillStationCd *string `fix:"1006"`
	//SideReasonCd is a non-required field for NoSides.
	SideReasonCd *string `fix:"1007"`
	//RptSeq is a non-required field for NoSides.
	RptSeq *int `fix:"83"`
	//SideTrdSubTyp is a non-required field for NoSides.
	SideTrdSubTyp *int `fix:"1008"`
	//SideTrdRegTS Component
	sidetrdregts.SideTrdRegTS
	//SideGrossTradeAmt is a non-required field for NoSides.
	SideGrossTradeAmt *float64 `fix:"1072"`
	//AggressorIndicator is a non-required field for NoSides.
	AggressorIndicator *bool `fix:"1057"`
	//ExchangeSpecialInstructions is a non-required field for NoSides.
	ExchangeSpecialInstructions *string `fix:"1139"`
	//NetGrossInd is a non-required field for NoSides.
	NetGrossInd *int `fix:"430"`
	//SideCurrency is a non-required field for NoSides.
	SideCurrency *string `fix:"1154"`
	//SideSettlCurrency is a non-required field for NoSides.
	SideSettlCurrency *string `fix:"1155"`
	//SettlDetails Component
	settldetails.SettlDetails
	//OrderCategory is a non-required field for NoSides.
	OrderCategory *string `fix:"1115"`
	//TradeReportOrderDetail Component
	tradereportorderdetail.TradeReportOrderDetail
	//SideExecID is a non-required field for NoSides.
	SideExecID *string `fix:"1427"`
	//OrderDelay is a non-required field for NoSides.
	OrderDelay *int `fix:"1428"`
	//OrderDelayUnit is a non-required field for NoSides.
	OrderDelayUnit *int `fix:"1429"`
	//SideLiquidityInd is a non-required field for NoSides.
	SideLiquidityInd *int `fix:"1444"`
}

//TrdCapRptSideGrp is a fix50sp2 Component
type TrdCapRptSideGrp struct {
	//NoSides is a required field for TrdCapRptSideGrp.
	NoSides []NoSides `fix:"552"`
}

func (m *TrdCapRptSideGrp) SetNoSides(v []NoSides) { m.NoSides = v }
