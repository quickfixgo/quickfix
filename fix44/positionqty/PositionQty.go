package positionqty

import (
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
)

func New() *PositionQty {
	var m PositionQty
	return &m
}

//NoPositions is a repeating group in PositionQty
type NoPositions struct {
	//PosType is a non-required field for NoPositions.
	PosType *string `fix:"703"`
	//LongQty is a non-required field for NoPositions.
	LongQty *float64 `fix:"704"`
	//ShortQty is a non-required field for NoPositions.
	ShortQty *float64 `fix:"705"`
	//PosQtyStatus is a non-required field for NoPositions.
	PosQtyStatus *int `fix:"706"`
	//NestedParties is a non-required component for NoPositions.
	NestedParties *nestedparties.NestedParties
}

//PositionQty is a fix44 Component
type PositionQty struct {
	//NoPositions is a non-required field for PositionQty.
	NoPositions []NoPositions `fix:"702,omitempty"`
}

func (m *PositionQty) SetNoPositions(v []NoPositions) { m.NoPositions = v }
