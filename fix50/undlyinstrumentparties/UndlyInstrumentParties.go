package undlyinstrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50/undlyinstrumentptyssubgrp"
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
	undlyinstrumentptyssubgrp.UndlyInstrumentPtysSubGrp
}

//UndlyInstrumentParties is a fix50 Component
type UndlyInstrumentParties struct {
	//NoUndlyInstrumentParties is a non-required field for UndlyInstrumentParties.
	NoUndlyInstrumentParties []NoUndlyInstrumentParties `fix:"1058,omitempty"`
}

func (m *UndlyInstrumentParties) SetNoUndlyInstrumentParties(v []NoUndlyInstrumentParties) {
	m.NoUndlyInstrumentParties = v
}
