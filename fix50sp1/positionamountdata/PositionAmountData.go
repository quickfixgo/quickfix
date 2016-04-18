package positionamountdata

//New returns an initialized PositionAmountData instance
func New() *PositionAmountData {
	var m PositionAmountData
	return &m
}

//NoPosAmt is a repeating group in PositionAmountData
type NoPosAmt struct {
	//PosAmtType is a non-required field for NoPosAmt.
	PosAmtType *string `fix:"707"`
	//PosAmt is a non-required field for NoPosAmt.
	PosAmt *float64 `fix:"708"`
	//PositionCurrency is a non-required field for NoPosAmt.
	PositionCurrency *string `fix:"1055"`
}

//NewNoPosAmt returns an initialized NoPosAmt instance
func NewNoPosAmt() *NoPosAmt {
	var m NoPosAmt
	return &m
}

func (m *NoPosAmt) SetPosAmtType(v string)       { m.PosAmtType = &v }
func (m *NoPosAmt) SetPosAmt(v float64)          { m.PosAmt = &v }
func (m *NoPosAmt) SetPositionCurrency(v string) { m.PositionCurrency = &v }

//PositionAmountData is a fix50sp1 Component
type PositionAmountData struct {
	//NoPosAmt is a non-required field for PositionAmountData.
	NoPosAmt []NoPosAmt `fix:"753,omitempty"`
}

func (m *PositionAmountData) SetNoPosAmt(v []NoPosAmt) { m.NoPosAmt = v }
