package marketsegmentgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/securitytradingrules"
	"github.com/quickfixgo/quickfix/fix50sp1/strikerules"
)

//NoMarketSegments is a repeating group in MarketSegmentGrp
type NoMarketSegments struct {
	//MarketID is a non-required field for NoMarketSegments.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for NoMarketSegments.
	MarketSegmentID *string `fix:"1300"`
	//SecurityTradingRules Component
	securitytradingrules.SecurityTradingRules
	//StrikeRules Component
	strikerules.StrikeRules
}

//MarketSegmentGrp is a fix50sp1 Component
type MarketSegmentGrp struct {
	//NoMarketSegments is a non-required field for MarketSegmentGrp.
	NoMarketSegments []NoMarketSegments `fix:"1310,omitempty"`
}

func (m *MarketSegmentGrp) SetNoMarketSegments(v []NoMarketSegments) { m.NoMarketSegments = v }
