package displayinstruction

//Component is a fix50sp1 DisplayInstruction Component
type Component struct {
	//SecondaryDisplayQty is a non-required field for DisplayInstruction.
	SecondaryDisplayQty *float64 `fix:"1082"`
	//DisplayWhen is a non-required field for DisplayInstruction.
	DisplayWhen *string `fix:"1083"`
	//DisplayMethod is a non-required field for DisplayInstruction.
	DisplayMethod *string `fix:"1084"`
	//DisplayLowQty is a non-required field for DisplayInstruction.
	DisplayLowQty *float64 `fix:"1085"`
	//DisplayHighQty is a non-required field for DisplayInstruction.
	DisplayHighQty *float64 `fix:"1086"`
	//DisplayMinIncr is a non-required field for DisplayInstruction.
	DisplayMinIncr *float64 `fix:"1087"`
	//RefreshQty is a non-required field for DisplayInstruction.
	RefreshQty *float64 `fix:"1088"`
	//DisplayQty is a non-required field for DisplayInstruction.
	DisplayQty *float64 `fix:"1138"`
}

func New() *Component { return new(Component) }
