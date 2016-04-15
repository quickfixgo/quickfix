package undlyinstrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/undlyinstrumentptyssubgrp"
)

func New() *UndlyInstrumentParties {
	var m UndlyInstrumentParties
	return &m
}

//NoUndlyInstrumentParties is a repeating group in UndlyInstrumentParties
type NoUndlyInstrumentParties struct {
	//UnderlyingInstrumentPartyID is a non-required field for NoUndlyInstrumentParties.
	UnderlyingInstrumentPartyID *string `fix:"1059"`
	//UnderlyingInstrumentPartyIDSource is a non-required field for NoUndlyInstrumentParties.
	UnderlyingInstrumentPartyIDSource *string `fix:"1060"`
	//UnderlyingInstrumentPartyRole is a non-required field for NoUndlyInstrumentParties.
	UnderlyingInstrumentPartyRole *int `fix:"1061"`
	//UndlyInstrumentPtysSubGrp is a non-required component for NoUndlyInstrumentParties.
	UndlyInstrumentPtysSubGrp *undlyinstrumentptyssubgrp.UndlyInstrumentPtysSubGrp
}

func (m *NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyID(v string) {
	m.UnderlyingInstrumentPartyID = &v
}
func (m *NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyIDSource(v string) {
	m.UnderlyingInstrumentPartyIDSource = &v
}
func (m *NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyRole(v int) {
	m.UnderlyingInstrumentPartyRole = &v
}
func (m *NoUndlyInstrumentParties) SetUndlyInstrumentPtysSubGrp(v undlyinstrumentptyssubgrp.UndlyInstrumentPtysSubGrp) {
	m.UndlyInstrumentPtysSubGrp = &v
}

//UndlyInstrumentParties is a fix50sp2 Component
type UndlyInstrumentParties struct {
	//NoUndlyInstrumentParties is a non-required field for UndlyInstrumentParties.
	NoUndlyInstrumentParties []NoUndlyInstrumentParties `fix:"1058,omitempty"`
}

func (m *UndlyInstrumentParties) SetNoUndlyInstrumentParties(v []NoUndlyInstrumentParties) {
	m.NoUndlyInstrumentParties = v
}
