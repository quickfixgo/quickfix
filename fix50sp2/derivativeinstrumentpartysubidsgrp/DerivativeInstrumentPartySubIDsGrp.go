package derivativeinstrumentpartysubidsgrp

//NoDerivativeInstrumentPartySubIDs is a repeating group in DerivativeInstrumentPartySubIDsGrp
type NoDerivativeInstrumentPartySubIDs struct {
	//DerivativeInstrumentPartySubID is a non-required field for NoDerivativeInstrumentPartySubIDs.
	DerivativeInstrumentPartySubID *string `fix:"1297"`
	//DerivativeInstrumentPartySubIDType is a non-required field for NoDerivativeInstrumentPartySubIDs.
	DerivativeInstrumentPartySubIDType *int `fix:"1298"`
}

//DerivativeInstrumentPartySubIDsGrp is a fix50sp2 Component
type DerivativeInstrumentPartySubIDsGrp struct {
	//NoDerivativeInstrumentPartySubIDs is a non-required field for DerivativeInstrumentPartySubIDsGrp.
	NoDerivativeInstrumentPartySubIDs []NoDerivativeInstrumentPartySubIDs `fix:"1296,omitempty"`
}

func (m *DerivativeInstrumentPartySubIDsGrp) SetNoDerivativeInstrumentPartySubIDs(v []NoDerivativeInstrumentPartySubIDs) {
	m.NoDerivativeInstrumentPartySubIDs = v
}
