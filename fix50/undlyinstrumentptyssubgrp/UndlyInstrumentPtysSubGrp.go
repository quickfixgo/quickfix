package undlyinstrumentptyssubgrp

//NoUndlyInstrumentPartySubIDs is a repeating group in UndlyInstrumentPtysSubGrp
type NoUndlyInstrumentPartySubIDs struct {
	//UndlyInstrumentPartySubID is a non-required field for NoUndlyInstrumentPartySubIDs.
	UndlyInstrumentPartySubID *string `fix:"1063"`
	//UndlyInstrumentPartySubIDType is a non-required field for NoUndlyInstrumentPartySubIDs.
	UndlyInstrumentPartySubIDType *int `fix:"1064"`
}

//Component is a fix50 UndlyInstrumentPtysSubGrp Component
type Component struct {
	//NoUndlyInstrumentPartySubIDs is a non-required field for UndlyInstrumentPtysSubGrp.
	NoUndlyInstrumentPartySubIDs []NoUndlyInstrumentPartySubIDs `fix:"1062,omitempty"`
}

func New() *Component { return new(Component) }
