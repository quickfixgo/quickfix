package ratesource

//NoRateSources is a repeating group in RateSource
type NoRateSources struct {
	//RateSource is a non-required field for NoRateSources.
	RateSource *int `fix:"1446"`
	//RateSourceType is a non-required field for NoRateSources.
	RateSourceType *int `fix:"1447"`
	//ReferencePage is a non-required field for NoRateSources.
	ReferencePage *string `fix:"1448"`
}

//Component is a fix50sp2 RateSource Component
type Component struct {
	//NoRateSources is a non-required field for RateSource.
	NoRateSources []NoRateSources `fix:"1445,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRateSources(v []NoRateSources) { m.NoRateSources = v }
