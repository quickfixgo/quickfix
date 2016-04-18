package marketsegmentgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/securitytradingrules"
	"github.com/quickfixgo/quickfix/fix50sp2/strikerules"
)

//New returns an initialized MarketSegmentGrp instance
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

//NewNoMarketSegments returns an initialized NoMarketSegments instance
func NewNoMarketSegments() *NoMarketSegments {
	var m NoMarketSegments
	return &m
}

func (m *NoMarketSegments) SetMarketID(v string)        { m.MarketID = &v }
func (m *NoMarketSegments) SetMarketSegmentID(v string) { m.MarketSegmentID = &v }
func (m *NoMarketSegments) SetSecurityTradingRules(v securitytradingrules.SecurityTradingRules) {
	m.SecurityTradingRules = &v
}
func (m *NoMarketSegments) SetStrikeRules(v strikerules.StrikeRules) { m.StrikeRules = &v }

//MarketSegmentGrp is a fix50sp2 Component
type MarketSegmentGrp struct {
	//NoMarketSegments is a non-required field for MarketSegmentGrp.
	NoMarketSegments []NoMarketSegments `fix:"1310,omitempty"`
}

func (m *MarketSegmentGrp) SetNoMarketSegments(v []NoMarketSegments) { m.NoMarketSegments = v }
