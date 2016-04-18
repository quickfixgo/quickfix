package strategyparametersgrp

//New returns an initialized StrategyParametersGrp instance
func New() *StrategyParametersGrp {
	var m StrategyParametersGrp
	return &m
}

//NoStrategyParameters is a repeating group in StrategyParametersGrp
type NoStrategyParameters struct {
	//StrategyParameterName is a non-required field for NoStrategyParameters.
	StrategyParameterName *string `fix:"958"`
	//StrategyParameterType is a non-required field for NoStrategyParameters.
	StrategyParameterType *int `fix:"959"`
	//StrategyParameterValue is a non-required field for NoStrategyParameters.
	StrategyParameterValue *string `fix:"960"`
}

//NewNoStrategyParameters returns an initialized NoStrategyParameters instance
func NewNoStrategyParameters() *NoStrategyParameters {
	var m NoStrategyParameters
	return &m
}

func (m *NoStrategyParameters) SetStrategyParameterName(v string)  { m.StrategyParameterName = &v }
func (m *NoStrategyParameters) SetStrategyParameterType(v int)     { m.StrategyParameterType = &v }
func (m *NoStrategyParameters) SetStrategyParameterValue(v string) { m.StrategyParameterValue = &v }

//StrategyParametersGrp is a fix50sp1 Component
type StrategyParametersGrp struct {
	//NoStrategyParameters is a non-required field for StrategyParametersGrp.
	NoStrategyParameters []NoStrategyParameters `fix:"957,omitempty"`
}

func (m *StrategyParametersGrp) SetNoStrategyParameters(v []NoStrategyParameters) {
	m.NoStrategyParameters = v
}
