package contamtgrp

//New returns an initialized ContAmtGrp instance
func New() *ContAmtGrp {
	var m ContAmtGrp
	return &m
}

//NoContAmts is a repeating group in ContAmtGrp
type NoContAmts struct {
	//ContAmtType is a non-required field for NoContAmts.
	ContAmtType *int `fix:"519"`
	//ContAmtValue is a non-required field for NoContAmts.
	ContAmtValue *float64 `fix:"520"`
	//ContAmtCurr is a non-required field for NoContAmts.
	ContAmtCurr *string `fix:"521"`
}

//NewNoContAmts returns an initialized NoContAmts instance
func NewNoContAmts() *NoContAmts {
	var m NoContAmts
	return &m
}

func (m *NoContAmts) SetContAmtType(v int)      { m.ContAmtType = &v }
func (m *NoContAmts) SetContAmtValue(v float64) { m.ContAmtValue = &v }
func (m *NoContAmts) SetContAmtCurr(v string)   { m.ContAmtCurr = &v }

//ContAmtGrp is a fix50 Component
type ContAmtGrp struct {
	//NoContAmts is a non-required field for ContAmtGrp.
	NoContAmts []NoContAmts `fix:"518,omitempty"`
}

func (m *ContAmtGrp) SetNoContAmts(v []NoContAmts) { m.NoContAmts = v }
