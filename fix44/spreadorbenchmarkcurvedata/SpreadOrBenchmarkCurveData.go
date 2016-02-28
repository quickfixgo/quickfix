package spreadorbenchmarkcurvedata

//Component is a fix44 SpreadOrBenchmarkCurveData Component
type Component struct {
	//Spread is a non-required field for SpreadOrBenchmarkCurveData.
	Spread *float64 `fix:"218"`
	//BenchmarkCurveCurrency is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveCurrency *string `fix:"220"`
	//BenchmarkCurveName is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveName *string `fix:"221"`
	//BenchmarkCurvePoint is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurvePoint *string `fix:"222"`
	//BenchmarkPrice is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkPrice *float64 `fix:"662"`
	//BenchmarkPriceType is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkPriceType *int `fix:"663"`
	//BenchmarkSecurityID is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkSecurityID *string `fix:"699"`
	//BenchmarkSecurityIDSource is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkSecurityIDSource *string `fix:"761"`
}

func New() *Component { return new(Component) }

func (m *Component) SetSpread(v float64)                   { m.Spread = &v }
func (m *Component) SetBenchmarkCurveCurrency(v string)    { m.BenchmarkCurveCurrency = &v }
func (m *Component) SetBenchmarkCurveName(v string)        { m.BenchmarkCurveName = &v }
func (m *Component) SetBenchmarkCurvePoint(v string)       { m.BenchmarkCurvePoint = &v }
func (m *Component) SetBenchmarkPrice(v float64)           { m.BenchmarkPrice = &v }
func (m *Component) SetBenchmarkPriceType(v int)           { m.BenchmarkPriceType = &v }
func (m *Component) SetBenchmarkSecurityID(v string)       { m.BenchmarkSecurityID = &v }
func (m *Component) SetBenchmarkSecurityIDSource(v string) { m.BenchmarkSecurityIDSource = &v }
