package peginstructions

//Component is a fix44 PegInstructions Component
type Component struct {
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

func New() *Component { return new(Component) }
