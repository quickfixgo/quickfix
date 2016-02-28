package peginstructions

//Component is a fix50sp2 PegInstructions Component
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
	//PegPriceType is a non-required field for PegInstructions.
	PegPriceType *int `fix:"1094"`
	//PegSecurityIDSource is a non-required field for PegInstructions.
	PegSecurityIDSource *string `fix:"1096"`
	//PegSecurityID is a non-required field for PegInstructions.
	PegSecurityID *string `fix:"1097"`
	//PegSymbol is a non-required field for PegInstructions.
	PegSymbol *string `fix:"1098"`
	//PegSecurityDesc is a non-required field for PegInstructions.
	PegSecurityDesc *string `fix:"1099"`
}

func New() *Component { return new(Component) }

func (m *Component) SetPegOffsetValue(v float64)     { m.PegOffsetValue = &v }
func (m *Component) SetPegMoveType(v int)            { m.PegMoveType = &v }
func (m *Component) SetPegOffsetType(v int)          { m.PegOffsetType = &v }
func (m *Component) SetPegLimitType(v int)           { m.PegLimitType = &v }
func (m *Component) SetPegRoundDirection(v int)      { m.PegRoundDirection = &v }
func (m *Component) SetPegScope(v int)               { m.PegScope = &v }
func (m *Component) SetPegPriceType(v int)           { m.PegPriceType = &v }
func (m *Component) SetPegSecurityIDSource(v string) { m.PegSecurityIDSource = &v }
func (m *Component) SetPegSecurityID(v string)       { m.PegSecurityID = &v }
func (m *Component) SetPegSymbol(v string)           { m.PegSymbol = &v }
func (m *Component) SetPegSecurityDesc(v string)     { m.PegSecurityDesc = &v }
