package fillsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties4"
)

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
}

//FillsGrp is a fix50sp1 Component
type FillsGrp struct {
	//NoFills is a non-required field for FillsGrp.
	NoFills []NoFills `fix:"1362,omitempty"`
}

func (m *FillsGrp) SetNoFills(v []NoFills) { m.NoFills = v }
