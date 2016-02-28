package contamtgrp

//NoContAmts is a repeating group in ContAmtGrp
type NoContAmts struct {
	//ContAmtType is a non-required field for NoContAmts.
	ContAmtType *int `fix:"519"`
	//ContAmtValue is a non-required field for NoContAmts.
	ContAmtValue *float64 `fix:"520"`
	//ContAmtCurr is a non-required field for NoContAmts.
	ContAmtCurr *string `fix:"521"`
}

//Component is a fix50sp2 ContAmtGrp Component
type Component struct {
	//NoContAmts is a non-required field for ContAmtGrp.
	NoContAmts []NoContAmts `fix:"518,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoContAmts(v []NoContAmts) { m.NoContAmts = v }
