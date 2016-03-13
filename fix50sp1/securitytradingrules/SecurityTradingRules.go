package securitytradingrules

import (
	"github.com/quickfixgo/quickfix/fix50sp1/basetradingrules"
	"github.com/quickfixgo/quickfix/fix50sp1/nestedinstrumentattribute"
	"github.com/quickfixgo/quickfix/fix50sp1/tradingsessionrulesgrp"
)

//SecurityTradingRules is a fix50sp1 Component
type SecurityTradingRules struct {
	//BaseTradingRules Component
	basetradingrules.BaseTradingRules
	//TradingSessionRulesGrp Component
	tradingsessionrulesgrp.TradingSessionRulesGrp
	//NestedInstrumentAttribute Component
	nestedinstrumentattribute.NestedInstrumentAttribute
}
