package strategyparametersgrp

//NoStrategyParameters is a repeating group in StrategyParametersGrp
type NoStrategyParameters struct {
	//StrategyParameterName is a non-required field for NoStrategyParameters.
	StrategyParameterName *string `fix:"958"`
	//StrategyParameterType is a non-required field for NoStrategyParameters.
	StrategyParameterType *int `fix:"959"`
	//StrategyParameterValue is a non-required field for NoStrategyParameters.
	StrategyParameterValue *string `fix:"960"`
}

//Component is a fix50sp1 StrategyParametersGrp Component
type Component struct {
	//NoStrategyParameters is a non-required field for StrategyParametersGrp.
	NoStrategyParameters []NoStrategyParameters `fix:"957,omitempty"`
}

func New() *Component { return new(Component) }
