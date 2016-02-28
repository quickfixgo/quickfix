package lottyperules

//NoLotTypeRules is a repeating group in LotTypeRules
type NoLotTypeRules struct {
	//LotType is a non-required field for NoLotTypeRules.
	LotType *string `fix:"1093"`
	//MinLotSize is a non-required field for NoLotTypeRules.
	MinLotSize *float64 `fix:"1231"`
}

//Component is a fix50sp2 LotTypeRules Component
type Component struct {
	//NoLotTypeRules is a non-required field for LotTypeRules.
	NoLotTypeRules []NoLotTypeRules `fix:"1234,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoLotTypeRules(v []NoLotTypeRules) { m.NoLotTypeRules = v }
