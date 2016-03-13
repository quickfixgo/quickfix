package derivativesecuritydefinition

import (
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrument"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrumentattribute"
	"github.com/quickfixgo/quickfix/fix50sp1/marketsegmentgrp"
)

//DerivativeSecurityDefinition is a fix50sp1 Component
type DerivativeSecurityDefinition struct {
	//DerivativeInstrument Component
	derivativeinstrument.DerivativeInstrument
	//DerivativeInstrumentAttribute Component
	derivativeinstrumentattribute.DerivativeInstrumentAttribute
	//MarketSegmentGrp Component
	marketsegmentgrp.MarketSegmentGrp
}
