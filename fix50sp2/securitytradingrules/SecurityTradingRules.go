package securitytradingrules

import (
	"github.com/quickfixgo/quickfix/fix50sp2/basetradingrules"
	"github.com/quickfixgo/quickfix/fix50sp2/nestedinstrumentattribute"
	"github.com/quickfixgo/quickfix/fix50sp2/tradingsessionrulesgrp"
)

//SecurityTradingRules is a fix50sp2 Component
type SecurityTradingRules struct {
	//BaseTradingRules Component
	basetradingrules.BaseTradingRules
	//TradingSessionRulesGrp Component
	tradingsessionrulesgrp.TradingSessionRulesGrp
	//NestedInstrumentAttribute Component
	nestedinstrumentattribute.NestedInstrumentAttribute
}
