package discretioninstructions

func New() *DiscretionInstructions {
	var m DiscretionInstructions
	return &m
}

//DiscretionInstructions is a fix50sp1 Component
type DiscretionInstructions struct {
	//DiscretionInst is a non-required field for DiscretionInstructions.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffsetValue is a non-required field for DiscretionInstructions.
	DiscretionOffsetValue *float64 `fix:"389"`
	//DiscretionMoveType is a non-required field for DiscretionInstructions.
	DiscretionMoveType *int `fix:"841"`
	//DiscretionOffsetType is a non-required field for DiscretionInstructions.
	DiscretionOffsetType *int `fix:"842"`
	//DiscretionLimitType is a non-required field for DiscretionInstructions.
	DiscretionLimitType *int `fix:"843"`
	//DiscretionRoundDirection is a non-required field for DiscretionInstructions.
	DiscretionRoundDirection *int `fix:"844"`
	//DiscretionScope is a non-required field for DiscretionInstructions.
	DiscretionScope *int `fix:"846"`
}

func (m *DiscretionInstructions) SetDiscretionInst(v string)         { m.DiscretionInst = &v }
func (m *DiscretionInstructions) SetDiscretionOffsetValue(v float64) { m.DiscretionOffsetValue = &v }
func (m *DiscretionInstructions) SetDiscretionMoveType(v int)        { m.DiscretionMoveType = &v }
func (m *DiscretionInstructions) SetDiscretionOffsetType(v int)      { m.DiscretionOffsetType = &v }
func (m *DiscretionInstructions) SetDiscretionLimitType(v int)       { m.DiscretionLimitType = &v }
func (m *DiscretionInstructions) SetDiscretionRoundDirection(v int)  { m.DiscretionRoundDirection = &v }
func (m *DiscretionInstructions) SetDiscretionScope(v int)           { m.DiscretionScope = &v }
