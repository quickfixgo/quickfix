package instrumentptyssubgrp

func New() *InstrumentPtysSubGrp {
	var m InstrumentPtysSubGrp
	return &m
}

//NoInstrumentPartySubIDs is a repeating group in InstrumentPtysSubGrp
type NoInstrumentPartySubIDs struct {
	//InstrumentPartySubID is a non-required field for NoInstrumentPartySubIDs.
	InstrumentPartySubID *string `fix:"1053"`
	//InstrumentPartySubIDType is a non-required field for NoInstrumentPartySubIDs.
	InstrumentPartySubIDType *int `fix:"1054"`
}

func (m *NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string)  { m.InstrumentPartySubID = &v }
func (m *NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) { m.InstrumentPartySubIDType = &v }

//InstrumentPtysSubGrp is a fix50sp1 Component
type InstrumentPtysSubGrp struct {
	//NoInstrumentPartySubIDs is a non-required field for InstrumentPtysSubGrp.
	NoInstrumentPartySubIDs []NoInstrumentPartySubIDs `fix:"1052,omitempty"`
}

func (m *InstrumentPtysSubGrp) SetNoInstrumentPartySubIDs(v []NoInstrumentPartySubIDs) {
	m.NoInstrumentPartySubIDs = v
}
