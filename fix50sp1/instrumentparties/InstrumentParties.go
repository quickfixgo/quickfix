package instrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentptyssubgrp"
)

//NoInstrumentParties is a repeating group in InstrumentParties
type NoInstrumentParties struct {
	//InstrumentPartyID is a non-required field for NoInstrumentParties.
	InstrumentPartyID *string `fix:"1019"`
	//InstrumentPartyIDSource is a non-required field for NoInstrumentParties.
	InstrumentPartyIDSource *string `fix:"1050"`
	//InstrumentPartyRole is a non-required field for NoInstrumentParties.
	InstrumentPartyRole *int `fix:"1051"`
	//InstrumentPtysSubGrp is a non-required component for NoInstrumentParties.
	InstrumentPtysSubGrp *instrumentptyssubgrp.InstrumentPtysSubGrp
}

//InstrumentParties is a fix50sp1 Component
type InstrumentParties struct {
	//NoInstrumentParties is a non-required field for InstrumentParties.
	NoInstrumentParties []NoInstrumentParties `fix:"1018,omitempty"`
}

func (m *InstrumentParties) SetNoInstrumentParties(v []NoInstrumentParties) { m.NoInstrumentParties = v }
