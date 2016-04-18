package ratesource

//New returns an initialized RateSource instance
func New() *RateSource {
	var m RateSource
	return &m
}

//NoRateSources is a repeating group in RateSource
type NoRateSources struct {
	//RateSource is a non-required field for NoRateSources.
	RateSource *int `fix:"1446"`
	//RateSourceType is a non-required field for NoRateSources.
	RateSourceType *int `fix:"1447"`
	//ReferencePage is a non-required field for NoRateSources.
	ReferencePage *string `fix:"1448"`
}

//NewNoRateSources returns an initialized NoRateSources instance
func NewNoRateSources() *NoRateSources {
	var m NoRateSources
	return &m
}

func (m *NoRateSources) SetRateSource(v int)       { m.RateSource = &v }
func (m *NoRateSources) SetRateSourceType(v int)   { m.RateSourceType = &v }
func (m *NoRateSources) SetReferencePage(v string) { m.ReferencePage = &v }

//RateSource is a fix50sp2 Component
type RateSource struct {
	//NoRateSources is a non-required field for RateSource.
	NoRateSources []NoRateSources `fix:"1445,omitempty"`
}

func (m *RateSource) SetNoRateSources(v []NoRateSources) { m.NoRateSources = v }
