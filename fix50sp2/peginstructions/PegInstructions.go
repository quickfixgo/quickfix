package peginstructions

func New() *PegInstructions {
	var m PegInstructions
	return &m
}

//PegInstructions is a fix50sp2 Component
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

func (m *PegInstructions) SetPegOffsetValue(v float64)     { m.PegOffsetValue = &v }
func (m *PegInstructions) SetPegMoveType(v int)            { m.PegMoveType = &v }
func (m *PegInstructions) SetPegOffsetType(v int)          { m.PegOffsetType = &v }
func (m *PegInstructions) SetPegLimitType(v int)           { m.PegLimitType = &v }
func (m *PegInstructions) SetPegRoundDirection(v int)      { m.PegRoundDirection = &v }
func (m *PegInstructions) SetPegScope(v int)               { m.PegScope = &v }
func (m *PegInstructions) SetPegPriceType(v int)           { m.PegPriceType = &v }
func (m *PegInstructions) SetPegSecurityIDSource(v string) { m.PegSecurityIDSource = &v }
func (m *PegInstructions) SetPegSecurityID(v string)       { m.PegSecurityID = &v }
func (m *PegInstructions) SetPegSymbol(v string)           { m.PegSymbol = &v }
func (m *PegInstructions) SetPegSecurityDesc(v string)     { m.PegSecurityDesc = &v }
