package tradingsessionrulesgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/tradingsessionrules"
)

//NoTradingSessionRules is a repeating group in TradingSessionRulesGrp
type NoTradingSessionRules struct {
	//TradingSessionID is a non-required field for NoTradingSessionRules.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessionRules.
	TradingSessionSubID *string `fix:"625"`
	//TradingSessionRules Component
	TradingSessionRules tradingsessionrules.Component
}

//Component is a fix50sp2 TradingSessionRulesGrp Component
type Component struct {
	//NoTradingSessionRules is a non-required field for TradingSessionRulesGrp.
	NoTradingSessionRules []NoTradingSessionRules `fix:"1309,omitempty"`
}

func New() *Component { return new(Component) }
