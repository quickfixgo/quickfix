package derivativeinstrumentpartysubidsgrp

func New() *DerivativeInstrumentPartySubIDsGrp {
	var m DerivativeInstrumentPartySubIDsGrp
	return &m
}

//NoDerivativeInstrumentPartySubIDs is a repeating group in DerivativeInstrumentPartySubIDsGrp
type NoDerivativeInstrumentPartySubIDs struct {
	//DerivativeInstrumentPartySubID is a non-required field for NoDerivativeInstrumentPartySubIDs.
	DerivativeInstrumentPartySubID *string `fix:"1297"`
	//DerivativeInstrumentPartySubIDType is a non-required field for NoDerivativeInstrumentPartySubIDs.
	DerivativeInstrumentPartySubIDType *int `fix:"1298"`
}

func (m *NoDerivativeInstrumentPartySubIDs) SetDerivativeInstrumentPartySubID(v string) {
	m.DerivativeInstrumentPartySubID = &v
}
func (m *NoDerivativeInstrumentPartySubIDs) SetDerivativeInstrumentPartySubIDType(v int) {
	m.DerivativeInstrumentPartySubIDType = &v
}

//DerivativeInstrumentPartySubIDsGrp is a fix50sp1 Component
type DerivativeInstrumentPartySubIDsGrp struct {
	//NoDerivativeInstrumentPartySubIDs is a non-required field for DerivativeInstrumentPartySubIDsGrp.
	NoDerivativeInstrumentPartySubIDs []NoDerivativeInstrumentPartySubIDs `fix:"1296,omitempty"`
}

func (m *DerivativeInstrumentPartySubIDsGrp) SetNoDerivativeInstrumentPartySubIDs(v []NoDerivativeInstrumentPartySubIDs) {
	m.NoDerivativeInstrumentPartySubIDs = v
}
