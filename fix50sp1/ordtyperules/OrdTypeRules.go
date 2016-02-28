package ordtyperules

//NoOrdTypeRules is a repeating group in OrdTypeRules
type NoOrdTypeRules struct {
	//OrdType is a non-required field for NoOrdTypeRules.
	OrdType *string `fix:"40"`
}

//Component is a fix50sp1 OrdTypeRules Component
type Component struct {
	//NoOrdTypeRules is a non-required field for OrdTypeRules.
	NoOrdTypeRules []NoOrdTypeRules `fix:"1237,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoOrdTypeRules(v []NoOrdTypeRules) { m.NoOrdTypeRules = v }
