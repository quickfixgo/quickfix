package instrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentptyssubgrp"
)

//NoInstrumentParties is a repeating group in InstrumentParties
type NoInstrumentParties struct {
	//InstrumentPartyID is a non-required field for NoInstrumentParties.
	InstrumentPartyID *string `fix:"1019"`
	//InstrumentPartyIDSource is a non-required field for NoInstrumentParties.
	InstrumentPartyIDSource *string `fix:"1050"`
	//InstrumentPartyRole is a non-required field for NoInstrumentParties.
	InstrumentPartyRole *int `fix:"1051"`
	//InstrumentPtysSubGrp Component
	InstrumentPtysSubGrp instrumentptyssubgrp.Component
}

//Component is a fix50 InstrumentParties Component
type Component struct {
	//NoInstrumentParties is a non-required field for InstrumentParties.
	NoInstrumentParties []NoInstrumentParties `fix:"1018,omitempty"`
}

func New() *Component { return new(Component) }
