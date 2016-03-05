package trdinstrmtleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
	"github.com/quickfixgo/quickfix/fix50sp1/tradecaplegunderlyingsgrp"
)

//NoLegs is a repeating group in TrdInstrmtLegGrp
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations Component
	legstipulations.LegStipulations
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties Component
	nestedparties.NestedParties
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegLastPx is a non-required field for NoLegs.
	LegLastPx *float64 `fix:"637"`
	//LegReportID is a non-required field for NoLegs.
	LegReportID *string `fix:"990"`
	//LegSettlCurrency is a non-required field for NoLegs.
	LegSettlCurrency *string `fix:"675"`
	//LegLastForwardPoints is a non-required field for NoLegs.
	LegLastForwardPoints *float64 `fix:"1073"`
	//LegCalculatedCcyLastQty is a non-required field for NoLegs.
	LegCalculatedCcyLastQty *float64 `fix:"1074"`
	//LegGrossTradeAmt is a non-required field for NoLegs.
	LegGrossTradeAmt *float64 `fix:"1075"`
	//LegNumber is a non-required field for NoLegs.
	LegNumber *int `fix:"1152"`
	//TradeCapLegUnderlyingsGrp Component
	tradecaplegunderlyingsgrp.TradeCapLegUnderlyingsGrp
	//LegVolatility is a non-required field for NoLegs.
	LegVolatility *float64 `fix:"1379"`
	//LegDividendYield is a non-required field for NoLegs.
	LegDividendYield *float64 `fix:"1381"`
	//LegCurrencyRatio is a non-required field for NoLegs.
	LegCurrencyRatio *float64 `fix:"1383"`
	//LegExecInst is a non-required field for NoLegs.
	LegExecInst *string `fix:"1384"`
	//LegLastQty is a non-required field for NoLegs.
	LegLastQty *float64 `fix:"1418"`
}

//TrdInstrmtLegGrp is a fix50sp1 Component
type TrdInstrmtLegGrp struct {
	//NoLegs is a non-required field for TrdInstrmtLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *TrdInstrmtLegGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
