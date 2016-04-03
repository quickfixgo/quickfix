package securitytradingrules

import (
	"github.com/quickfixgo/quickfix/fix50sp1/basetradingrules"
	"github.com/quickfixgo/quickfix/fix50sp1/nestedinstrumentattribute"
	"github.com/quickfixgo/quickfix/fix50sp1/tradingsessionrulesgrp"
)

func New() *SecurityTradingRules {
	var m SecurityTradingRules
	return &m
}

//SecurityTradingRules is a fix50sp1 Component
type SecurityTradingRules struct {
	//BaseTradingRules is a non-required component for SecurityTradingRules.
	BaseTradingRules *basetradingrules.BaseTradingRules
	//TradingSessionRulesGrp is a non-required component for SecurityTradingRules.
	TradingSessionRulesGrp *tradingsessionrulesgrp.TradingSessionRulesGrp
	//NestedInstrumentAttribute is a non-required component for SecurityTradingRules.
	NestedInstrumentAttribute *nestedinstrumentattribute.NestedInstrumentAttribute
}

func (m *SecurityTradingRules) SetBaseTradingRules(v basetradingrules.BaseTradingRules) {
	m.BaseTradingRules = &v
}
func (m *SecurityTradingRules) SetTradingSessionRulesGrp(v tradingsessionrulesgrp.TradingSessionRulesGrp) {
	m.TradingSessionRulesGrp = &v
}
func (m *SecurityTradingRules) SetNestedInstrumentAttribute(v nestedinstrumentattribute.NestedInstrumentAttribute) {
	m.NestedInstrumentAttribute = &v
}
