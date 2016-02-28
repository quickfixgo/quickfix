package undlyinstrumentptyssubgrp

//NoUndlyInstrumentPartySubIDs is a repeating group in UndlyInstrumentPtysSubGrp
type NoUndlyInstrumentPartySubIDs struct {
	//UnderlyingInstrumentPartySubID is a non-required field for NoUndlyInstrumentPartySubIDs.
	UnderlyingInstrumentPartySubID *string `fix:"1063"`
	//UnderlyingInstrumentPartySubIDType is a non-required field for NoUndlyInstrumentPartySubIDs.
	UnderlyingInstrumentPartySubIDType *int `fix:"1064"`
}

//Component is a fix50sp2 UndlyInstrumentPtysSubGrp Component
type Component struct {
	//NoUndlyInstrumentPartySubIDs is a non-required field for UndlyInstrumentPtysSubGrp.
	NoUndlyInstrumentPartySubIDs []NoUndlyInstrumentPartySubIDs `fix:"1062,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoUndlyInstrumentPartySubIDs(v []NoUndlyInstrumentPartySubIDs) {
	m.NoUndlyInstrumentPartySubIDs = v
}
