package mdfullgrp

import (
	"github.com/quickfixgo/quickfix/fix50/parties"
	"time"
)

//NoMDEntries is a repeating group in MDFullGrp
type NoMDEntries struct {
	//MDEntryType is a required field for NoMDEntries.
	MDEntryType string `fix:"269"`
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
	//Text is a non-required field for NoMDEntries.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoMDEntries.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoMDEntries.
	EncodedText *string `fix:"355"`
	//MDPriceLevel is a non-required field for NoMDEntries.
	MDPriceLevel *int `fix:"1023"`
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
	//MDEntryID is a non-required field for NoMDEntries.
	MDEntryID *string `fix:"278"`
	//Parties Component
	Parties parties.Component
	//SecondaryOrderID is a non-required field for NoMDEntries.
	SecondaryOrderID *string `fix:"198"`
	//OrdType is a non-required field for NoMDEntries.
	OrdType *string `fix:"40"`
}

//Component is a fix50 MDFullGrp Component
type Component struct {
	//NoMDEntries is a required field for MDFullGrp.
	NoMDEntries []NoMDEntries `fix:"268"`
}

func New() *Component { return new(Component) }
