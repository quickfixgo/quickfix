package discretioninstructions

//Component is a fix50sp2 DiscretionInstructions Component
type Component struct {
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

func New() *Component { return new(Component) }

func (m *Component) SetDiscretionInst(v string)         { m.DiscretionInst = &v }
func (m *Component) SetDiscretionOffsetValue(v float64) { m.DiscretionOffsetValue = &v }
func (m *Component) SetDiscretionMoveType(v int)        { m.DiscretionMoveType = &v }
func (m *Component) SetDiscretionOffsetType(v int)      { m.DiscretionOffsetType = &v }
func (m *Component) SetDiscretionLimitType(v int)       { m.DiscretionLimitType = &v }
func (m *Component) SetDiscretionRoundDirection(v int)  { m.DiscretionRoundDirection = &v }
func (m *Component) SetDiscretionScope(v int)           { m.DiscretionScope = &v }
