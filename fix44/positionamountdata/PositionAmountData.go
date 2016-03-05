package positionamountdata

//NoPosAmt is a repeating group in PositionAmountData
type NoPosAmt struct {
	//PosAmtType is a non-required field for NoPosAmt.
	PosAmtType *string `fix:"707"`
	//PosAmt is a non-required field for NoPosAmt.
	PosAmt *float64 `fix:"708"`
}

//PositionAmountData is a fix44 Component
type PositionAmountData struct {
	//NoPosAmt is a non-required field for PositionAmountData.
	NoPosAmt []NoPosAmt `fix:"753,omitempty"`
}

func (m *PositionAmountData) SetNoPosAmt(v []NoPosAmt) { m.NoPosAmt = v }
