package derivativesecuritydefinition

import (
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrument"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrumentattribute"
	"github.com/quickfixgo/quickfix/fix50sp1/marketsegmentgrp"
)

func New() *DerivativeSecurityDefinition {
	var m DerivativeSecurityDefinition
	return &m
}

//DerivativeSecurityDefinition is a fix50sp1 Component
type DerivativeSecurityDefinition struct {
	//DerivativeInstrument is a non-required component for DerivativeSecurityDefinition.
	DerivativeInstrument *derivativeinstrument.DerivativeInstrument
	//DerivativeInstrumentAttribute is a non-required component for DerivativeSecurityDefinition.
	DerivativeInstrumentAttribute *derivativeinstrumentattribute.DerivativeInstrumentAttribute
	//MarketSegmentGrp is a non-required component for DerivativeSecurityDefinition.
	MarketSegmentGrp *marketsegmentgrp.MarketSegmentGrp
}

func (m *DerivativeSecurityDefinition) SetDerivativeInstrument(v derivativeinstrument.DerivativeInstrument) {
	m.DerivativeInstrument = &v
}
func (m *DerivativeSecurityDefinition) SetDerivativeInstrumentAttribute(v derivativeinstrumentattribute.DerivativeInstrumentAttribute) {
	m.DerivativeInstrumentAttribute = &v
}
func (m *DerivativeSecurityDefinition) SetMarketSegmentGrp(v marketsegmentgrp.MarketSegmentGrp) {
	m.MarketSegmentGrp = &v
}
