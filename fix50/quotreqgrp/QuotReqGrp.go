package quotreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/quotqualgrp"
	"github.com/quickfixgo/quickfix/fix50/quotreqlegsgrp"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"time"
)

//NoRelatedSym is a repeating group in QuotReqGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
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
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
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
	//Stipulations Component
	Stipulations stipulations.Component
	//Account is a non-required field for NoRelatedSym.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoRelatedSym.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoRelatedSym.
	AccountType *int `fix:"581"`
	//QuotReqLegsGrp Component
	QuotReqLegsGrp quotreqlegsgrp.Component
	//QuotQualGrp Component
	QuotQualGrp quotqualgrp.Component
	//QuotePriceType is a non-required field for NoRelatedSym.
	QuotePriceType *int `fix:"692"`
	//OrdType is a non-required field for NoRelatedSym.
	OrdType *string `fix:"40"`
	//ValidUntilTime is a non-required field for NoRelatedSym.
	ValidUntilTime *time.Time `fix:"62"`
	//ExpireTime is a non-required field for NoRelatedSym.
	ExpireTime *time.Time `fix:"126"`
	//TransactTime is a non-required field for NoRelatedSym.
	TransactTime *time.Time `fix:"60"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//PriceType is a non-required field for NoRelatedSym.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoRelatedSym.
	Price *float64 `fix:"44"`
	//Price2 is a non-required field for NoRelatedSym.
	Price2 *float64 `fix:"640"`
	//YieldData Component
	YieldData yielddata.Component
	//Parties Component
	Parties parties.Component
}

//Component is a fix50 QuotReqGrp Component
type Component struct {
	//NoRelatedSym is a required field for QuotReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
