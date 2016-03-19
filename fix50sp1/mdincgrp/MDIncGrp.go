package mdincgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/secsizesgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/statsindgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/yielddata"
	"time"
)

func New(nomdentries []NoMDEntries) *MDIncGrp {
	var m MDIncGrp
	m.SetNoMDEntries(nomdentries)
	return &m
}

//NoMDEntries is a repeating group in MDIncGrp
type NoMDEntries struct {
	//MDUpdateAction is a required field for NoMDEntries.
	MDUpdateAction string `fix:"279"`
	//DeleteReason is a non-required field for NoMDEntries.
	DeleteReason *string `fix:"285"`
	//MDEntryType is a non-required field for NoMDEntries.
	MDEntryType *string `fix:"269"`
	//MDEntryID is a non-required field for NoMDEntries.
	MDEntryID *string `fix:"278"`
	//MDEntryRefID is a non-required field for NoMDEntries.
	MDEntryRefID *string `fix:"280"`
	//Instrument is a non-required component for NoMDEntries.
	Instrument *instrument.Instrument
	//UndInstrmtGrp is a non-required component for NoMDEntries.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for NoMDEntries.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//FinancialStatus is a non-required field for NoMDEntries.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for NoMDEntries.
	CorporateAction *string `fix:"292"`
	//MDEntryPx is a non-required field for NoMDEntries.
	MDEntryPx *float64 `fix:"270"`
	//Currency is a non-required field for NoMDEntries.
	Currency *string `fix:"15"`
	//MDEntrySize is a non-required field for NoMDEntries.
	MDEntrySize *float64 `fix:"271"`
	//MDEntryDate is a non-required field for NoMDEntries.
	MDEntryDate *string `fix:"272"`
	//MDEntryTime is a non-required field for NoMDEntries.
	MDEntryTime *string `fix:"273"`
	//TickDirection is a non-required field for NoMDEntries.
	TickDirection *string `fix:"274"`
	//MDMkt is a non-required field for NoMDEntries.
	MDMkt *string `fix:"275"`
	//TradingSessionID is a non-required field for NoMDEntries.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoMDEntries.
	TradingSessionSubID *string `fix:"625"`
	//QuoteCondition is a non-required field for NoMDEntries.
	QuoteCondition *string `fix:"276"`
	//TradeCondition is a non-required field for NoMDEntries.
	TradeCondition *string `fix:"277"`
	//MDEntryOriginator is a non-required field for NoMDEntries.
	MDEntryOriginator *string `fix:"282"`
	//LocationID is a non-required field for NoMDEntries.
	LocationID *string `fix:"283"`
	//DeskID is a non-required field for NoMDEntries.
	DeskID *string `fix:"284"`
	//OpenCloseSettlFlag is a non-required field for NoMDEntries.
	OpenCloseSettlFlag *string `fix:"286"`
	//TimeInForce is a non-required field for NoMDEntries.
	TimeInForce *string `fix:"59"`
	//ExpireDate is a non-required field for NoMDEntries.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoMDEntries.
	ExpireTime *time.Time `fix:"126"`
	//MinQty is a non-required field for NoMDEntries.
	MinQty *float64 `fix:"110"`
	//ExecInst is a non-required field for NoMDEntries.
	ExecInst *string `fix:"18"`
	//SellerDays is a non-required field for NoMDEntries.
	SellerDays *int `fix:"287"`
	//OrderID is a non-required field for NoMDEntries.
	OrderID *string `fix:"37"`
	//QuoteEntryID is a non-required field for NoMDEntries.
	QuoteEntryID *string `fix:"299"`
	//MDEntryBuyer is a non-required field for NoMDEntries.
	MDEntryBuyer *string `fix:"288"`
	//MDEntrySeller is a non-required field for NoMDEntries.
	MDEntrySeller *string `fix:"289"`
	//NumberOfOrders is a non-required field for NoMDEntries.
	NumberOfOrders *int `fix:"346"`
	//MDEntryPositionNo is a non-required field for NoMDEntries.
	MDEntryPositionNo *int `fix:"290"`
	//Scope is a non-required field for NoMDEntries.
	Scope *string `fix:"546"`
	//PriceDelta is a non-required field for NoMDEntries.
	PriceDelta *float64 `fix:"811"`
	//NetChgPrevDay is a non-required field for NoMDEntries.
	NetChgPrevDay *float64 `fix:"451"`
	//Text is a non-required field for NoMDEntries.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoMDEntries.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoMDEntries.
	EncodedText *string `fix:"355"`
	//OrderCapacity is a non-required field for NoMDEntries.
	OrderCapacity *string `fix:"528"`
	//MDOriginType is a non-required field for NoMDEntries.
	MDOriginType *int `fix:"1024"`
	//HighPx is a non-required field for NoMDEntries.
	HighPx *float64 `fix:"332"`
	//LowPx is a non-required field for NoMDEntries.
	LowPx *float64 `fix:"333"`
	//TradeVolume is a non-required field for NoMDEntries.
	TradeVolume *float64 `fix:"1020"`
	//SettlType is a non-required field for NoMDEntries.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoMDEntries.
	SettlDate *string `fix:"64"`
	//MDQuoteType is a non-required field for NoMDEntries.
	MDQuoteType *int `fix:"1070"`
	//RptSeq is a non-required field for NoMDEntries.
	RptSeq *int `fix:"83"`
	//DealingCapacity is a non-required field for NoMDEntries.
	DealingCapacity *float64 `fix:"1048"`
	//MDEntrySpotRate is a non-required field for NoMDEntries.
	MDEntrySpotRate *float64 `fix:"1026"`
	//MDEntryForwardPoints is a non-required field for NoMDEntries.
	MDEntryForwardPoints *float64 `fix:"1027"`
	//MDPriceLevel is a non-required field for NoMDEntries.
	MDPriceLevel *int `fix:"1023"`
	//Parties is a non-required component for NoMDEntries.
	Parties *parties.Parties
	//SecondaryOrderID is a non-required field for NoMDEntries.
	SecondaryOrderID *string `fix:"198"`
	//OrdType is a non-required field for NoMDEntries.
	OrdType *string `fix:"40"`
	//MDSubBookType is a non-required field for NoMDEntries.
	MDSubBookType *int `fix:"1173"`
	//MarketDepth is a non-required field for NoMDEntries.
	MarketDepth *int `fix:"264"`
	//PriceType is a non-required field for NoMDEntries.
	PriceType *int `fix:"423"`
	//YieldData is a non-required component for NoMDEntries.
	YieldData *yielddata.YieldData
	//SpreadOrBenchmarkCurveData is a non-required component for NoMDEntries.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//SecSizesGrp is a non-required component for NoMDEntries.
	SecSizesGrp *secsizesgrp.SecSizesGrp
	//LotType is a non-required field for NoMDEntries.
	LotType *string `fix:"1093"`
	//SecurityTradingStatus is a non-required field for NoMDEntries.
	SecurityTradingStatus *int `fix:"326"`
	//HaltReasonChar is a non-required field for NoMDEntries.
	HaltReasonChar *string `fix:"327"`
	//TrdType is a non-required field for NoMDEntries.
	TrdType *int `fix:"828"`
	//MatchType is a non-required field for NoMDEntries.
	MatchType *string `fix:"574"`
	//TradeID is a non-required field for NoMDEntries.
	TradeID *string `fix:"1003"`
	//TransBkdTime is a non-required field for NoMDEntries.
	TransBkdTime *time.Time `fix:"483"`
	//TransactTime is a non-required field for NoMDEntries.
	TransactTime *time.Time `fix:"60"`
	//StatsIndGrp is a non-required component for NoMDEntries.
	StatsIndGrp *statsindgrp.StatsIndGrp
}

//MDIncGrp is a fix50sp1 Component
type MDIncGrp struct {
	//NoMDEntries is a required field for MDIncGrp.
	NoMDEntries []NoMDEntries `fix:"268"`
}

func (m *MDIncGrp) SetNoMDEntries(v []NoMDEntries) { m.NoMDEntries = v }
