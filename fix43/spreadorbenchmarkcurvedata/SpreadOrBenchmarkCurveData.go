package spreadorbenchmarkcurvedata

//SpreadOrBenchmarkCurveData is a fix43 Component
type SpreadOrBenchmarkCurveData struct {
	//Spread is a non-required field for SpreadOrBenchmarkCurveData.
	Spread *float64 `fix:"218"`
	//BenchmarkCurveCurrency is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveCurrency *string `fix:"220"`
	//BenchmarkCurveName is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveName *string `fix:"221"`
	//BenchmarkCurvePoint is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurvePoint *string `fix:"222"`
}

func (m *SpreadOrBenchmarkCurveData) SetSpread(v float64) { m.Spread = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkCurveCurrency(v string) {
	m.BenchmarkCurveCurrency = &v
}
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkCurveName(v string)  { m.BenchmarkCurveName = &v }
func (m *SpreadOrBenchmarkCurveData) SetBenchmarkCurvePoint(v string) { m.BenchmarkCurvePoint = &v }
