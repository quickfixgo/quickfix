package marketsegmentgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/securitytradingrules"
	"github.com/quickfixgo/quickfix/fix50sp1/strikerules"
)

func New() *MarketSegmentGrp {
	var m MarketSegmentGrp
	return &m
}

//NoMarketSegments is a repeating group in MarketSegmentGrp
type NoMarketSegments struct {
	//MarketID is a non-required field for NoMarketSegments.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for NoMarketSegments.
	MarketSegmentID *string `fix:"1300"`
	//SecurityTradingRules is a non-required component for NoMarketSegments.
	SecurityTradingRules *securitytradingrules.SecurityTradingRules
	//StrikeRules is a non-required component for NoMarketSegments.
	StrikeRules *strikerules.StrikeRules
}

//MarketSegmentGrp is a fix50sp1 Component
type MarketSegmentGrp struct {
	//NoMarketSegments is a non-required field for MarketSegmentGrp.
	NoMarketSegments []NoMarketSegments `fix:"1310,omitempty"`
}

func (m *MarketSegmentGrp) SetNoMarketSegments(v []NoMarketSegments) { m.NoMarketSegments = v }
