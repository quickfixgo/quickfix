package quotreqrjctgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/quotqualgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/quotreqlegsgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/yielddata"
	"time"
)

//NoRelatedSym is a repeating group in QuotReqRjctGrp
type NoRelatedSym struct {
	//Instrument is a required component for NoRelatedSym.
	instrument.Instrument
	//FinancingDetails is a non-required component for NoRelatedSym.
	FinancingDetails *financingdetails.FinancingDetails
	//UndInstrmtGrp is a non-required component for NoRelatedSym.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//PrevClosePx is a non-required field for NoRelatedSym.
	PrevClosePx *float64 `fix:"140"`
	//QuoteRequestType is a non-required field for NoRelatedSym.
	QuoteRequestType *int `fix:"303"`
	//QuoteType is a non-required field for NoRelatedSym.
	QuoteType *int `fix:"537"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
	//TradeOriginationDate is a non-required field for NoRelatedSym.
	TradeOriginationDate *string `fix:"229"`
	//Side is a non-required field for NoRelatedSym.
	Side *string `fix:"54"`
	//QtyType is a non-required field for NoRelatedSym.
	QtyType *int `fix:"854"`
	//OrderQtyData is a non-required component for NoRelatedSym.
	OrderQtyData *orderqtydata.OrderQtyData
	//SettlType is a non-required field for NoRelatedSym.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoRelatedSym.
	SettlDate *string `fix:"64"`
	//SettlDate2 is a non-required field for NoRelatedSym.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoRelatedSym.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations is a non-required component for NoRelatedSym.
	Stipulations *stipulations.Stipulations
	//Account is a non-required field for NoRelatedSym.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoRelatedSym.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoRelatedSym.
	AccountType *int `fix:"581"`
	//QuotReqLegsGrp is a non-required component for NoRelatedSym.
	QuotReqLegsGrp *quotreqlegsgrp.QuotReqLegsGrp
	//QuotQualGrp is a non-required component for NoRelatedSym.
	QuotQualGrp *quotqualgrp.QuotQualGrp
	//QuotePriceType is a non-required field for NoRelatedSym.
	QuotePriceType *int `fix:"692"`
	//OrdType is a non-required field for NoRelatedSym.
	OrdType *string `fix:"40"`
	//ExpireTime is a non-required field for NoRelatedSym.
	ExpireTime *time.Time `fix:"126"`
	//TransactTime is a non-required field for NoRelatedSym.
	TransactTime *time.Time `fix:"60"`
	//SpreadOrBenchmarkCurveData is a non-required component for NoRelatedSym.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//PriceType is a non-required field for NoRelatedSym.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoRelatedSym.
	Price *float64 `fix:"44"`
	//Price2 is a non-required field for NoRelatedSym.
	Price2 *float64 `fix:"640"`
	//YieldData is a non-required component for NoRelatedSym.
	YieldData *yielddata.YieldData
	//Parties is a non-required component for NoRelatedSym.
	Parties *parties.Parties
}

//QuotReqRjctGrp is a fix50sp1 Component
type QuotReqRjctGrp struct {
	//NoRelatedSym is a required field for QuotReqRjctGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *QuotReqRjctGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
