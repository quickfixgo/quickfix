package lottyperules

//New returns an initialized LotTypeRules instance
func New() *LotTypeRules {
	var m LotTypeRules
	return &m
}

//NoLotTypeRules is a repeating group in LotTypeRules
type NoLotTypeRules struct {
	//LotType is a non-required field for NoLotTypeRules.
	LotType *string `fix:"1093"`
	//MinLotSize is a non-required field for NoLotTypeRules.
	MinLotSize *float64 `fix:"1231"`
}

//NewNoLotTypeRules returns an initialized NoLotTypeRules instance
func NewNoLotTypeRules() *NoLotTypeRules {
	var m NoLotTypeRules
	return &m
}

func (m *NoLotTypeRules) SetLotType(v string)     { m.LotType = &v }
func (m *NoLotTypeRules) SetMinLotSize(v float64) { m.MinLotSize = &v }

//LotTypeRules is a fix50sp2 Component
type LotTypeRules struct {
	//NoLotTypeRules is a non-required field for LotTypeRules.
	NoLotTypeRules []NoLotTypeRules `fix:"1234,omitempty"`
}

func (m *LotTypeRules) SetNoLotTypeRules(v []NoLotTypeRules) { m.NoLotTypeRules = v }
