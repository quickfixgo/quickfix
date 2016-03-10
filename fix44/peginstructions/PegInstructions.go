package peginstructions

//PegInstructions is a fix44 Component
type PegInstructions struct {
	//PegOffsetValue is a non-required field for PegInstructions.
	PegOffsetValue *float64 `fix:"211"`
	//PegMoveType is a non-required field for PegInstructions.
	PegMoveType *int `fix:"835"`
	//PegOffsetType is a non-required field for PegInstructions.
	PegOffsetType *int `fix:"836"`
	//PegLimitType is a non-required field for PegInstructions.
	PegLimitType *int `fix:"837"`
	//PegRoundDirection is a non-required field for PegInstructions.
	PegRoundDirection *int `fix:"838"`
	//PegScope is a non-required field for PegInstructions.
	PegScope *int `fix:"840"`
}

func (m *PegInstructions) SetPegOffsetValue(v float64) { m.PegOffsetValue = &v }
func (m *PegInstructions) SetPegMoveType(v int)        { m.PegMoveType = &v }
func (m *PegInstructions) SetPegOffsetType(v int)      { m.PegOffsetType = &v }
func (m *PegInstructions) SetPegLimitType(v int)       { m.PegLimitType = &v }
func (m *PegInstructions) SetPegRoundDirection(v int)  { m.PegRoundDirection = &v }
func (m *PegInstructions) SetPegScope(v int)           { m.PegScope = &v }
