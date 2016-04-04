package fillsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties4"
)

func New() *FillsGrp {
	var m FillsGrp
	return &m
}

//NoFills is a repeating group in FillsGrp
type NoFills struct {
	//FillExecID is a non-required field for NoFills.
	FillExecID *string `fix:"1363"`
	//FillPx is a non-required field for NoFills.
	FillPx *float64 `fix:"1364"`
	//FillQty is a non-required field for NoFills.
	FillQty *float64 `fix:"1365"`
	//NestedParties4 is a non-required component for NoFills.
	NestedParties4 *nestedparties4.NestedParties4
	//FillLiquidityInd is a non-required field for NoFills.
	FillLiquidityInd *int `fix:"1443"`
}

func (m *NoFills) SetFillExecID(v string)                            { m.FillExecID = &v }
func (m *NoFills) SetFillPx(v float64)                               { m.FillPx = &v }
func (m *NoFills) SetFillQty(v float64)                              { m.FillQty = &v }
func (m *NoFills) SetNestedParties4(v nestedparties4.NestedParties4) { m.NestedParties4 = &v }
func (m *NoFills) SetFillLiquidityInd(v int)                         { m.FillLiquidityInd = &v }

//FillsGrp is a fix50sp2 Component
type FillsGrp struct {
	//NoFills is a non-required field for FillsGrp.
	NoFills []NoFills `fix:"1362,omitempty"`
}

func (m *FillsGrp) SetNoFills(v []NoFills) { m.NoFills = v }
