package trdcaprptsidegrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/clrinstgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp1/contamtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/settldetails"
	"github.com/quickfixgo/quickfix/fix50sp1/sidetrdregts"
	"github.com/quickfixgo/quickfix/fix50sp1/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/trdallocgrp"
	"time"
)

func New(nosides []NoSides) *TrdCapRptSideGrp {
	var m TrdCapRptSideGrp
	m.SetNoSides(nosides)
	return &m
}

//NoSides is a repeating group in TrdCapRptSideGrp
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//OrderID is a non-required field for NoSides.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoSides.
	SecondaryOrderID *string `fix:"198"`
	//ClOrdID is a non-required field for NoSides.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoSides.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoSides.
	ListID *string `fix:"66"`
	//Parties is a non-required component for NoSides.
	Parties *parties.Parties
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
	//ClrInstGrp is a non-required component for NoSides.
	ClrInstGrp *clrinstgrp.ClrInstGrp
	//TradeInputSource is a non-required field for NoSides.
	TradeInputSource *string `fix:"578"`
	//TradeInputDevice is a non-required field for NoSides.
	TradeInputDevice *string `fix:"579"`
	//OrderInputDevice is a non-required field for NoSides.
	OrderInputDevice *string `fix:"821"`
	//ComplianceID is a non-required field for NoSides.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoSides.
	SolicitedFlag *bool `fix:"377"`
	//OrderCapacity is a non-required field for NoSides.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoSides.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoSides.
	CustOrderCapacity *int `fix:"582"`
	//OrdType is a non-required field for NoSides.
	OrdType *string `fix:"40"`
	//ExecInst is a non-required field for NoSides.
	ExecInst *string `fix:"18"`
	//TransBkdTime is a non-required field for NoSides.
	TransBkdTime *time.Time `fix:"483"`
	//TradingSessionID is a non-required field for NoSides.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoSides.
	TradingSessionSubID *string `fix:"625"`
	//TimeBracket is a non-required field for NoSides.
	TimeBracket *string `fix:"943"`
	//CommissionData is a non-required component for NoSides.
	CommissionData *commissiondata.CommissionData
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
	//ContAmtGrp is a non-required component for NoSides.
	ContAmtGrp *contamtgrp.ContAmtGrp
	//Stipulations is a non-required component for NoSides.
	Stipulations *stipulations.Stipulations
	//MiscFeesGrp is a non-required component for NoSides.
	MiscFeesGrp *miscfeesgrp.MiscFeesGrp
	//ExchangeRule is a non-required field for NoSides.
	ExchangeRule *string `fix:"825"`
	//TradeAllocIndicator is a non-required field for NoSides.
	TradeAllocIndicator *int `fix:"826"`
	//PreallocMethod is a non-required field for NoSides.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NoSides.
	AllocID *string `fix:"70"`
	//TrdAllocGrp is a non-required component for NoSides.
	TrdAllocGrp *trdallocgrp.TrdAllocGrp
	//SideQty is a non-required field for NoSides.
	SideQty *int `fix:"1009"`
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
	//SideTrdRegTS is a non-required component for NoSides.
	SideTrdRegTS *sidetrdregts.SideTrdRegTS
	//ExecRefID is a non-required field for NoSides.
	ExecRefID *string `fix:"19"`
	//LotType is a non-required field for NoSides.
	LotType *string `fix:"1093"`
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
	//SettlDetails is a non-required component for NoSides.
	SettlDetails *settldetails.SettlDetails
}

//TrdCapRptSideGrp is a fix50sp1 Component
type TrdCapRptSideGrp struct {
	//NoSides is a required field for TrdCapRptSideGrp.
	NoSides []NoSides `fix:"552"`
}

func (m *TrdCapRptSideGrp) SetNoSides(v []NoSides) { m.NoSides = v }
