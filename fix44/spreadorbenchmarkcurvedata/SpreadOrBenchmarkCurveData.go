package spreadorbenchmarkcurvedata

//SpreadOrBenchmarkCurveData is a fix44 Component
type SpreadOrBenchmarkCurveData struct {
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

func (m *SpreadOrBenchmarkCurveData) SetSpread(v float64) { m.Spread = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkCurveCurrency(v string) {
	m.BenchmarkCurveCurrency = &v
}
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkCurveName(v string)  { m.BenchmarkCurveName = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkCurvePoint(v string) { m.BenchmarkCurvePoint = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkPrice(v float64)     { m.BenchmarkPrice = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkPriceType(v int)     { m.BenchmarkPriceType = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkSecurityID(v string) { m.BenchmarkSecurityID = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkSecurityIDSource(v string) {
	m.BenchmarkSecurityIDSource = &v
}
