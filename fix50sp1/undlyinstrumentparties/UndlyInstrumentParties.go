package undlyinstrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50sp1/undlyinstrumentptyssubgrp"
)

//NoUndlyInstrumentParties is a repeating group in UndlyInstrumentParties
type NoUndlyInstrumentParties struct {
	//UndlyInstrumentPartyID is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyID *string `fix:"1059"`
	//UndlyInstrumentPartyIDSource is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyIDSource *string `fix:"1060"`
	//UndlyInstrumentPartyRole is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyRole *int `fix:"1061"`
	//UndlyInstrumentPtysSubGrp Component
	UndlyInstrumentPtysSubGrp undlyinstrumentptyssubgrp.Component
}

//Component is a fix50sp1 UndlyInstrumentParties Component
type Component struct {
	//NoUndlyInstrumentParties is a non-required field for UndlyInstrumentParties.
	NoUndlyInstrumentParties []NoUndlyInstrumentParties `fix:"1058,omitempty"`
}

func New() *Component { return new(Component) }
