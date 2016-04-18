package quotreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/quotqualgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/quotreqlegsgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/ratesource"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"time"
)

//New returns an initialized QuotReqGrp instance
func New(norelatedsym []NoRelatedSym) *QuotReqGrp {
	var m QuotReqGrp
	m.SetNoRelatedSym(norelatedsym)
	return &m
}

//NoRelatedSym is a repeating group in QuotReqGrp
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
	//ValidUntilTime is a non-required field for NoRelatedSym.
	ValidUntilTime *time.Time `fix:"62"`
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
	//MinQty is a non-required field for NoRelatedSym.
	MinQty *float64 `fix:"110"`
	//SettlCurrency is a non-required field for NoRelatedSym.
	SettlCurrency *string `fix:"120"`
	//RateSource is a non-required component for NoRelatedSym.
	RateSource *ratesource.RateSource
}

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym(instrument instrument.Instrument) *NoRelatedSym {
	var m NoRelatedSym
	m.SetInstrument(instrument)
	return &m
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument) { m.Instrument = v }
func (m *NoRelatedSym) SetFinancingDetails(v financingdetails.FinancingDetails) {
	m.FinancingDetails = &v
}
func (m *NoRelatedSym) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)    { m.UndInstrmtGrp = &v }
func (m *NoRelatedSym) SetPrevClosePx(v float64)                          { m.PrevClosePx = &v }
func (m *NoRelatedSym) SetQuoteRequestType(v int)                         { m.QuoteRequestType = &v }
func (m *NoRelatedSym) SetQuoteType(v int)                                { m.QuoteType = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)                      { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string)                   { m.TradingSessionSubID = &v }
func (m *NoRelatedSym) SetTradeOriginationDate(v string)                  { m.TradeOriginationDate = &v }
func (m *NoRelatedSym) SetSide(v string)                                  { m.Side = &v }
func (m *NoRelatedSym) SetQtyType(v int)                                  { m.QtyType = &v }
func (m *NoRelatedSym) SetOrderQtyData(v orderqtydata.OrderQtyData)       { m.OrderQtyData = &v }
func (m *NoRelatedSym) SetSettlType(v string)                             { m.SettlType = &v }
func (m *NoRelatedSym) SetSettlDate(v string)                             { m.SettlDate = &v }
func (m *NoRelatedSym) SetSettlDate2(v string)                            { m.SettlDate2 = &v }
func (m *NoRelatedSym) SetOrderQty2(v float64)                            { m.OrderQty2 = &v }
func (m *NoRelatedSym) SetCurrency(v string)                              { m.Currency = &v }
func (m *NoRelatedSym) SetStipulations(v stipulations.Stipulations)       { m.Stipulations = &v }
func (m *NoRelatedSym) SetAccount(v string)                               { m.Account = &v }
func (m *NoRelatedSym) SetAcctIDSource(v int)                             { m.AcctIDSource = &v }
func (m *NoRelatedSym) SetAccountType(v int)                              { m.AccountType = &v }
func (m *NoRelatedSym) SetQuotReqLegsGrp(v quotreqlegsgrp.QuotReqLegsGrp) { m.QuotReqLegsGrp = &v }
func (m *NoRelatedSym) SetQuotQualGrp(v quotqualgrp.QuotQualGrp)          { m.QuotQualGrp = &v }
func (m *NoRelatedSym) SetQuotePriceType(v int)                           { m.QuotePriceType = &v }
func (m *NoRelatedSym) SetOrdType(v string)                               { m.OrdType = &v }
func (m *NoRelatedSym) SetValidUntilTime(v time.Time)                     { m.ValidUntilTime = &v }
func (m *NoRelatedSym) SetExpireTime(v time.Time)                         { m.ExpireTime = &v }
func (m *NoRelatedSym) SetTransactTime(v time.Time)                       { m.TransactTime = &v }
func (m *NoRelatedSym) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *NoRelatedSym) SetPriceType(v int)                    { m.PriceType = &v }
func (m *NoRelatedSym) SetPrice(v float64)                    { m.Price = &v }
func (m *NoRelatedSym) SetPrice2(v float64)                   { m.Price2 = &v }
func (m *NoRelatedSym) SetYieldData(v yielddata.YieldData)    { m.YieldData = &v }
func (m *NoRelatedSym) SetParties(v parties.Parties)          { m.Parties = &v }
func (m *NoRelatedSym) SetMinQty(v float64)                   { m.MinQty = &v }
func (m *NoRelatedSym) SetSettlCurrency(v string)             { m.SettlCurrency = &v }
func (m *NoRelatedSym) SetRateSource(v ratesource.RateSource) { m.RateSource = &v }

//QuotReqGrp is a fix50sp2 Component
type QuotReqGrp struct {
	//NoRelatedSym is a required field for QuotReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *QuotReqGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
