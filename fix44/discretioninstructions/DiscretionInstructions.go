package discretioninstructions

//Component is a fix44 DiscretionInstructions Component
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
