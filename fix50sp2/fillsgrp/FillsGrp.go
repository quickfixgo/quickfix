package fillsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties4"
)

//NoFills is a repeating group in FillsGrp
type NoFills struct {
	//FillExecID is a non-required field for NoFills.
	FillExecID *string `fix:"1363"`
	//FillPx is a non-required field for NoFills.
	FillPx *float64 `fix:"1364"`
	//FillQty is a non-required field for NoFills.
	FillQty *float64 `fix:"1365"`
	//NestedParties4 Component
	NestedParties4 nestedparties4.Component
	//FillLiquidityInd is a non-required field for NoFills.
	FillLiquidityInd *int `fix:"1443"`
}

//Component is a fix50sp2 FillsGrp Component
type Component struct {
	//NoFills is a non-required field for FillsGrp.
	NoFills []NoFills `fix:"1362,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoFills(v []NoFills) { m.NoFills = v }
