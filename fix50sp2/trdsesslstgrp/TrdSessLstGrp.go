package trdsesslstgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/tradingsessionrules"
	"time"
)

//NoTradingSessions is a repeating group in TrdSessLstGrp
type NoTradingSessions struct {
	//TradingSessionID is a required field for NoTradingSessions.
	TradingSessionID string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
	//SecurityExchange is a non-required field for NoTradingSessions.
	SecurityExchange *string `fix:"207"`
	//TradSesMethod is a non-required field for NoTradingSessions.
	TradSesMethod *int `fix:"338"`
	//TradSesMode is a non-required field for NoTradingSessions.
	TradSesMode *int `fix:"339"`
	//UnsolicitedIndicator is a non-required field for NoTradingSessions.
	UnsolicitedIndicator *bool `fix:"325"`
	//TradSesStatus is a required field for NoTradingSessions.
	TradSesStatus int `fix:"340"`
	//TradSesStatusRejReason is a non-required field for NoTradingSessions.
	TradSesStatusRejReason *int `fix:"567"`
	//TradSesStartTime is a non-required field for NoTradingSessions.
	TradSesStartTime *time.Time `fix:"341"`
	//TradSesOpenTime is a non-required field for NoTradingSessions.
	TradSesOpenTime *time.Time `fix:"342"`
	//TradSesPreCloseTime is a non-required field for NoTradingSessions.
	TradSesPreCloseTime *time.Time `fix:"343"`
	//TradSesCloseTime is a non-required field for NoTradingSessions.
	TradSesCloseTime *time.Time `fix:"344"`
	//TradSesEndTime is a non-required field for NoTradingSessions.
	TradSesEndTime *time.Time `fix:"345"`
	//TotalVolumeTraded is a non-required field for NoTradingSessions.
	TotalVolumeTraded *float64 `fix:"387"`
	//Text is a non-required field for NoTradingSessions.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoTradingSessions.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoTradingSessions.
	EncodedText *string `fix:"355"`
	//MarketID is a non-required field for NoTradingSessions.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for NoTradingSessions.
	MarketSegmentID *string `fix:"1300"`
	//TradingSessionDesc is a non-required field for NoTradingSessions.
	TradingSessionDesc *string `fix:"1326"`
	//TradingSessionRules Component
	TradingSessionRules tradingsessionrules.Component
	//TransactTime is a non-required field for NoTradingSessions.
	TransactTime *time.Time `fix:"60"`
	//TradSesUpdateAction is a non-required field for NoTradingSessions.
	TradSesUpdateAction *string `fix:"1327"`
}

//Component is a fix50sp2 TrdSessLstGrp Component
type Component struct {
	//NoTradingSessions is a required field for TrdSessLstGrp.
	NoTradingSessions []NoTradingSessions `fix:"386"`
}

func New() *Component { return new(Component) }
