package undlyinstrumentptyssubgrp

func New() *UndlyInstrumentPtysSubGrp {
	var m UndlyInstrumentPtysSubGrp
	return &m
}

//NoUndlyInstrumentPartySubIDs is a repeating group in UndlyInstrumentPtysSubGrp
type NoUndlyInstrumentPartySubIDs struct {
	//UndlyInstrumentPartySubID is a non-required field for NoUndlyInstrumentPartySubIDs.
	UndlyInstrumentPartySubID *string `fix:"1063"`
	//UndlyInstrumentPartySubIDType is a non-required field for NoUndlyInstrumentPartySubIDs.
	UndlyInstrumentPartySubIDType *int `fix:"1064"`
}

func (m *NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubID(v string) {
	m.UndlyInstrumentPartySubID = &v
}
func (m *NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubIDType(v int) {
	m.UndlyInstrumentPartySubIDType = &v
}

//UndlyInstrumentPtysSubGrp is a fix50 Component
type UndlyInstrumentPtysSubGrp struct {
	//NoUndlyInstrumentPartySubIDs is a non-required field for UndlyInstrumentPtysSubGrp.
	NoUndlyInstrumentPartySubIDs []NoUndlyInstrumentPartySubIDs `fix:"1062,omitempty"`
}

func (m *UndlyInstrumentPtysSubGrp) SetNoUndlyInstrumentPartySubIDs(v []NoUndlyInstrumentPartySubIDs) {
	m.NoUndlyInstrumentPartySubIDs = v
}
