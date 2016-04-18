package ordtyperules

//New returns an initialized OrdTypeRules instance
func New() *OrdTypeRules {
	var m OrdTypeRules
	return &m
}

//NoOrdTypeRules is a repeating group in OrdTypeRules
type NoOrdTypeRules struct {
	//OrdType is a non-required field for NoOrdTypeRules.
	OrdType *string `fix:"40"`
}

//NewNoOrdTypeRules returns an initialized NoOrdTypeRules instance
func NewNoOrdTypeRules() *NoOrdTypeRules {
	var m NoOrdTypeRules
	return &m
}

func (m *NoOrdTypeRules) SetOrdType(v string) { m.OrdType = &v }

//OrdTypeRules is a fix50sp2 Component
type OrdTypeRules struct {
	//NoOrdTypeRules is a non-required field for OrdTypeRules.
	NoOrdTypeRules []NoOrdTypeRules `fix:"1237,omitempty"`
}

func (m *OrdTypeRules) SetNoOrdTypeRules(v []NoOrdTypeRules) { m.NoOrdTypeRules = v }
