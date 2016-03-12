package derivativesecuritydefinition

import (
	"github.com/quickfixgo/quickfix/fix50sp2/derivativeinstrument"
	"github.com/quickfixgo/quickfix/fix50sp2/derivativeinstrumentattribute"
	"github.com/quickfixgo/quickfix/fix50sp2/marketsegmentgrp"
)

//DerivativeSecurityDefinition is a fix50sp2 Component
type DerivativeSecurityDefinition struct {
	//DerivativeInstrument Component
	derivativeinstrument.DerivativeInstrument
	//DerivativeInstrumentAttribute Component
	derivativeinstrumentattribute.DerivativeInstrumentAttribute
	//MarketSegmentGrp Component
	marketsegmentgrp.MarketSegmentGrp
}
