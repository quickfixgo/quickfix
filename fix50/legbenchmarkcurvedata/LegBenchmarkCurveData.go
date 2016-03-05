package legbenchmarkcurvedata

//LegBenchmarkCurveData is a fix50 Component
type LegBenchmarkCurveData struct {
	//LegBenchmarkCurveCurrency is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkCurveCurrency *string `fix:"676"`
	//LegBenchmarkCurveName is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkCurveName *string `fix:"677"`
	//LegBenchmarkCurvePoint is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkCurvePoint *string `fix:"678"`
	//LegBenchmarkPrice is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkPrice *float64 `fix:"679"`
	//LegBenchmarkPriceType is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkPriceType *int `fix:"680"`
}

func (m *LegBenchmarkCurveData) SetLegBenchmarkCurveCurrency(v string) {
	m.LegBenchmarkCurveCurrency = &v
}
func (m *LegBenchmarkCurveData) SetLegBenchmarkCurveName(v string)  { m.LegBenchmarkCurveName = &v }
func (m *LegBenchmarkCurveData) SetLegBenchmarkCurvePoint(v string) { m.LegBenchmarkCurvePoint = &v }
func (m *LegBenchmarkCurveData) SetLegBenchmarkPrice(v float64)     { m.LegBenchmarkPrice = &v }
func (m *LegBenchmarkCurveData) SetLegBenchmarkPriceType(v int)     { m.LegBenchmarkPriceType = &v }
