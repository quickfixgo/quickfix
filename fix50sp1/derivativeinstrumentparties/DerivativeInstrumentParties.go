package derivativeinstrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrumentpartysubidsgrp"
)

//NoDerivativeInstrumentParties is a repeating group in DerivativeInstrumentParties
type NoDerivativeInstrumentParties struct {
	//DerivativeInstrumentPartyID is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyID *string `fix:"1293"`
	//DerivativeInstrumentPartyIDSource is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyIDSource *string `fix:"1294"`
	//DerivativeInstrumentPartyRole is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyRole *int `fix:"1295"`
	//DerivativeInstrumentPartySubIDsGrp Component
	DerivativeInstrumentPartySubIDsGrp derivativeinstrumentpartysubidsgrp.Component
}

//Component is a fix50sp1 DerivativeInstrumentParties Component
type Component struct {
	//NoDerivativeInstrumentParties is a non-required field for DerivativeInstrumentParties.
	NoDerivativeInstrumentParties []NoDerivativeInstrumentParties `fix:"1292,omitempty"`
}

func New() *Component { return new(Component) }
