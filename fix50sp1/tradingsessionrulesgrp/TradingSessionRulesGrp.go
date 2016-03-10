package tradingsessionrulesgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/tradingsessionrules"
)

//NoTradingSessionRules is a repeating group in TradingSessionRulesGrp
type NoTradingSessionRules struct {
	//TradingSessionID is a non-required field for NoTradingSessionRules.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessionRules.
	TradingSessionSubID *string `fix:"625"`
	//TradingSessionRules Component
	tradingsessionrules.TradingSessionRules
}

//TradingSessionRulesGrp is a fix50sp1 Component
type TradingSessionRulesGrp struct {
	//NoTradingSessionRules is a non-required field for TradingSessionRulesGrp.
	NoTradingSessionRules []NoTradingSessionRules `fix:"1309,omitempty"`
}

func (m *TradingSessionRulesGrp) SetNoTradingSessionRules(v []NoTradingSessionRules) {
	m.NoTradingSessionRules = v
}
